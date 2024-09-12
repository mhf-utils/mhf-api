package models

type Melee struct {
	ID              int       `json:"id"`
	Index           int       `json:"index"`
	Name            string    `json:"name"`
	Model           string    `json:"model"`
	Rarity          int       `json:"rarity"`
	Type            string    `json:"type"`
	ZennyCost       int32     `json:"zenny_cost"`
	SharpnessId     int8      `json:"sharpness_id"`
	SharpnessMax    int8      `json:"sharpness_max"`
	Sharpness       Sharpness `json:"sharpness"`
	RawDamage       int16     `json:"raw_damage"`
	Defense         int16     `json:"defense"`
	Affinity        int8      `json:"affinity"`
	Element         string    `json:"element"`
	EleDamage       byte      `json:"ele_damage"`
	Ailment         string    `json:"ailment"`
	AilDamage       byte      `json:"ail_damage"`
	Slots           byte      `json:"slots"`
	WeaponAttribute byte      `json:"weapon_attribute"`
	Unk4            byte      `json:"unk4"`
	UpgradePath     int16     `json:"upgrade_path"`
	Unk5            int16     `json:"unk5"`
	EqType          int16     `json:"eq_type"`
	Length          int32     `json:"length"`
	WeaponType      int32     `json:"weapon_type"`
	VisualEffects   int16     `json:"visual_effects"`
	Unk11           int16     `json:"unk11"`
	Unk12           byte      `json:"unk12"`
	Unk13           byte      `json:"unk13"`
	Unk14           byte      `json:"unk14"`
	Unk15           byte      `json:"unk15"`
	Unk16           int32     `json:"unk16"`
	Zenith_skill    int32     `json:"zenith_skill"`
	Description1    string    `json:"description1"`
	Description2    string    `json:"description2"`
	Description3    string    `json:"description3"`
	MhfY            string    `json:"mhf_y"`
}

type Sharpness struct {
	ID      int   `json:"id"`
	Index   int   `json:"index"`
	Red     int16 `json:"red"`
	Orange  int16 `json:"orange"`
	Yellow  int16 `json:"yellow"`
	Green   int16 `json:"green"`
	Blue    int16 `json:"blue"`
	White   int16 `json:"white"`
	Purple  int16 `json:"purple"`
	SkyBlue int16 `json:"sky_blue"`
}
