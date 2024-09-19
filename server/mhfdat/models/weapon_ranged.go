package models

type Ranged struct {
	ID             int    `json:"id"`
	Index          int    `json:"index"`
	Name           string `json:"name"`
	Model          string `json:"model"`
	Rarity         int    `json:"rarity"`
	MaxSlotsMaybe  byte   `json:"max_slots_maybe"`
	Type           string `json:"type"`
	Unk05          byte   `json:"unk05"`
	Eq             byte   `json:"eq"`
	Unk07          byte   `json:"unk07"`
	Unk08          uint32 `json:"unk08"`
	Unk0C          uint32 `json:"unk0c"`
	Unk10          uint32 `json:"unk10"`
	ZennyCost      uint32 `json:"zenny_cost"`
	RawDamage      uint16 `json:"raw_damage"`
	Defense        uint16 `json:"defense"`
	RecoilMaybe    byte   `json:"recoil_maybe"`
	Slots          byte   `json:"slots"`
	Affinity       byte   `json:"affinity"`
	SortOrderMaybe byte   `json:"sort_order_maybe"`
	Unk20          byte   `json:"unk20"`
	Element        string `json:"element"`
	EleDamage      byte   `json:"ele_damage"`
	Unk23          byte   `json:"unk23"`
	Unk24          uint32 `json:"unk24"`
	Unk2C          uint32 `json:"unk2c"`
	Unk34          uint32 `json:"unk34"`
	Description1   string `json:"description1"`
	Description2   string `json:"description2"`
	Description3   string `json:"description3"`
	MhfY           string `json:"mhf_y"`
}
