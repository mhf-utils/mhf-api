package models

type Item struct {
	ID          int    `json:"id"`
	Index       int    `json:"index"`
	Name        string `json:"name"`
	Unk00       byte   `json:"unk00"`
	Unk01       byte   `json:"unk01"`
	Rarity      int    `json:"rarity"`
	MaxStack    byte   `json:"max_stack"`
	Unk04       byte   `json:"unk04"`
	IconName    string `json:"icon_name"`
	IconColor   byte   `json:"icon_color"`
	Unk07       byte   `json:"unk07"`
	Unk08       int16  `json:"unk08"`
	Unk0A       int16  `json:"unk0a"`
	BuyPrice    int32  `json:"buy_price"`
	SellPrice   int32  `json:"sell_price"`
	Type        int16  `json:"type"` // 04 = Consumable, 05 = Placeable
	Unk16       int16  `json:"unk16"`
	Unk18       int16  `json:"unk18"`
	Unk1A       byte   `json:"unk1a"`
	Unk1B       byte   `json:"unk1b"`
	Unk1C       int16  `json:"unk1c"`
	Unk1E       byte   `json:"unk1e"`
	Unk1F       byte   `json:"unk1f"`
	Unk20       int16  `json:"unk20"`
	Unk22       int16  `json:"unk22"`
	Description string `json:"description"`
	Source      string `json:"source"`
}
