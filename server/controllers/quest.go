package controllers

import (
	"encoding/json"
	"io"
	"mhf-api/server/models"
	"mhf-api/utils/binary"
	"mhf-api/utils/logger"
	"mhf-api/utils/pointers"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var quest_types = map[string]models.QuestType{
	"quest_*": {
		From: int64(pointers.QuestOneStarFrom),
		To:   int64(pointers.QuestOneStarTo),
	},
	"quest_**": {
		From: int64(pointers.QuestTwoStarsFrom),
		To:   int64(pointers.QuestTwoStarsTo),
	},
	"quest_***": {
		From: int64(pointers.QuestThreeStarsFrom),
		To:   int64(pointers.QuestThreeStarsTo),
	},
	"quest_****": {
		From: int64(pointers.QuestFourStarsFrom),
		To:   int64(pointers.QuestFourStarsTo),
	},
	"quest_*****": {
		From: int64(pointers.QuestFiveStarsFrom),
		To:   int64(pointers.QuestFiveStarsTo),
	},
	"quest_******": {
		From: int64(pointers.QuestSixStarsFrom),
		To:   int64(pointers.QuestSixStarsTo),
	},
}

type ControllerQuest struct {
	log         *logger.Logger
	binary_file *binary.BinaryFile
}

func NewControllerQuest(log *logger.Logger, binary_file *binary.BinaryFile) *ControllerQuest {
	return &ControllerQuest{
		log,
		binary_file,
	}
}

func (controller *ControllerQuest) List(res http.ResponseWriter, req *http.Request) {
	var entries []models.Quest
	query := req.URL.Query()
	query_quest_type := query.Get("quest_type")
	if query_quest_type == "all" {
		for key, quest_type := range quest_types {
			quests := controller.getCategory(
				key,
				quest_type.From,
				quest_type.To,
			)
			entries = append(entries, quests...)
		}
		json.NewEncoder(res).Encode(entries)
	}

	quest_type, exists := quest_types[query_quest_type]
	if !exists {
		json.NewEncoder(res).Encode("quest_type not provided")
	}

	quests := controller.getCategory(
		query_quest_type,
		quest_type.From,
		quest_type.To,
	)
	json.NewEncoder(res).Encode(quests)
}

func (controller *ControllerQuest) Read(res http.ResponseWriter, req *http.Request) {
	var entry models.Quest
	query := req.URL.Query()
	query_quest_type := query.Get("quest_type")
	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])

	exists := false
	for key := range quest_types {
		if key == query_quest_type {
			exists = true
			break
		}
	}

	if query_quest_type == "" {
		json.NewEncoder(res).Encode("quest_type not provided")
	}
	if !exists {
		json.NewEncoder(res).Encode("quest_type not found")
	}

	for key, quest_type := range quest_types {
		if key == query_quest_type {
			entry = controller.getEntryByIndex(id, models.QuestType{
				Type: key,
				From: quest_type.From,
				To:   quest_type.To,
			})
		}
	}

	json.NewEncoder(res).Encode(entry)
}

func (controller *ControllerQuest) getCategory(
	quest_type_type string,
	quest_type_from int64,
	quest_type_to int64,
) []models.Quest {
	var entries []models.Quest

	count := (quest_type_to-quest_type_from)/pointers.QuestLength + 1
	entries = make([]models.Quest, count)

	for index := 0; index < int(count); index++ {
		entries[index] = controller.getEntryByIndex(index, models.QuestType{
			Type: quest_type_type,
			From: quest_type_from,
			To:   quest_type_to,
		})
	}

	return entries
}

func (controller *ControllerQuest) getEntry() models.Quest {
	var entry models.Quest
	entry.QuestId, _ = controller.binary_file.ReadInt16()
	entry.QuestNumber, _ = controller.binary_file.ReadInt16()
	entry.KeyQuest, _ = controller.binary_file.ReadByte()
	entry.UrgentQuest, _ = controller.binary_file.ReadByte()
	entry.Unknown, _ = controller.binary_file.ReadInt16()

	return entry
}

func (controller *ControllerQuest) getEntryByIndex(index int, quest_type models.QuestType) models.Quest {
	cursor := int(quest_type.From) + index*pointers.QuestLength
	controller.binary_file.BaseStream.Seek(int64(cursor), io.SeekStart)
	entry := controller.getEntry()
	entry.Index = index
	entry.ID = cursor

	return entry
}
