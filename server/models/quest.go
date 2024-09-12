package models

type Quest struct {
	ID          int   `json:"id"`
	Index       int   `json:"index"`
	QuestId     int16 `json:"quest_id"`
	QuestNumber int16 `json:"quest_number"`
	KeyQuest    byte  `json:"key_quest"`
	UrgentQuest byte  `json:"urgent_quest"`
	Unknown     int16 `json:"unknown"`
}

type QuestType struct {
	Type string
	From int64
	To   int64
}
