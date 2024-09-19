package controllers

import (
	"encoding/json"
	"io"
	"mhf-api/core"
	"mhf-api/server/mhfdat/models"
	"mhf-api/utils/binary"
	"mhf-api/utils/logger"
	"mhf-api/utils/pointers"
	"net/http"

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
	params := mux.Vars(req)

	if params["type"] == "" {
		json.NewEncoder(res).Encode("Query params 'type' is required. Available: [  quest_*, quest_**, quest_***, quest_****, quest_*****, quest_****** ]")
		return
	}

	quest_type, exists := quest_types[params["type"]]
	if !exists {
		json.NewEncoder(res).Encode("Query params 'type' is required. Available: [  quest_*, quest_**, quest_***, quest_****, quest_*****, quest_****** ]")
		return
	}

	extractFunc := func(index int) models.Quest {
		return controller.getEntryByIndex(index, models.QuestType{
			Type: params["type"],
			From: quest_type.From,
			To:   quest_type.To,
		})
	}
	core.Paginate(res, req, quest_type.From, quest_type.To, pointers.QuestLength, extractFunc)
}

func (controller *ControllerQuest) Read(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	if params["type"] == "" {
		json.NewEncoder(res).Encode("Query params 'type' is required. Available: [  quest_*, quest_**, quest_***, quest_****, quest_*****, quest_****** ]")
		return
	}

	quest_type, exists := quest_types[params["type"]]
	if !exists {
		json.NewEncoder(res).Encode("Query params 'type' is required. Available: [  quest_*, quest_**, quest_***, quest_****, quest_*****, quest_****** ]")
		return
	}

	extractFunc := func(id int) models.Quest {
		return controller.getEntryByIndex(id, models.QuestType{
			Type: params["type"],
			From: quest_type.From,
			To:   quest_type.To,
		})
	}
	core.ReadItem(res, req, extractFunc)
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
