package models

type Equipment struct {
	ID            int    `json:"id"`
	Index         int    `json:"index"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	ModelIdMale   int16  `json:"model_id_male"`
	ModelIdFemale int16  `json:"model_id_female"`
	IsHelm        bool   `json:"is_helm"`
	IsChest       bool   `json:"is_chest"`
	IsArm         bool   `json:"is_arm"`
	IsWaist       bool   `json:"is_waist"`
	IsLeg         bool   `json:"is_leg"`
	Gender        string `json:"gender"`
	Role          string `json:"role"`
	Bool1         bool   `json:"bool1"`
	IsSPEquip     bool   `json:"is_sp_equip"`
	Bool3         bool   `json:"bool3"`
	Bool4         bool   `json:"bool4"`
	Rarity        int    `json:"rarity"`
	MaxLevel      byte   `json:"max_level"`
	Unk1_1        byte   `json:"unk1_1"`
	Unk1_2        byte   `json:"unk1_2"`
	Unk1_3        byte   `json:"unk1_3"`
	Unk1_4        byte   `json:"unk1_4"`
	Unk2          byte   `json:"unk2"`
	ZennyCost     int32  `json:"zenny_cost"`
	Unk3          int16  `json:"unk3"`
	BaseDefense   int16  `json:"base_defense"`
	FireRes       byte   `json:"fire_res"`
	WaterRes      byte   `json:"water_res"`
	ThunderRes    byte   `json:"thunder_res"`
	DragonRes     byte   `json:"dragon_res"`
	IceRes        byte   `json:"ice_res"`
	Unk3_1        int16  `json:"unk3_1"`
	BaseSlots     byte   `json:"base_slots"`
	MaxSlots      byte   `json:"max_slots"`
	SthEventCrown byte   `json:"sth_event_crown"`
	Unk5          byte   `json:"unk5"`
	Unk6          byte   `json:"unk6"`
	Unk7_1        byte   `json:"unk7_1"`
	Unk7_2        byte   `json:"unk7_2"`
	Unk7_3        byte   `json:"unk7_3"`
	Unk7_4        byte   `json:"unk7_4"`
	Unk8_1        byte   `json:"unk8_1"`
	Unk8_2        byte   `json:"unk8_2"`
	Unk8_3        byte   `json:"unk8_3"`
	Unk8_4        byte   `json:"unk8_4"`
	Unk10         int16  `json:"unk10"`
	SkillId1      string `json:"skill_id1"`
	SkillPts1     byte   `json:"skill_pts1"`
	SkillId2      string `json:"skill_id2"`
	SkillPts2     byte   `json:"skill_pts2"`
	SkillId3      string `json:"skill_id3"`
	SkillPts3     byte   `json:"skill_pts3"`
	SkillId4      string `json:"skill_id4"`
	SkillPts4     byte   `json:"skill_pts4"`
	SkillId5      string `json:"skill_id5"`
	SkillPts5     byte   `json:"skill_pts5"`
	SthHiden      int32  `json:"sth_hiden"`
	Unk12         int32  `json:"unk12"`
	Unk13         byte   `json:"unk13"`
	Unk14         byte   `json:"unk14"`
	Unk15         byte   `json:"unk15"`
	Unk16         byte   `json:"unk16"`
	Unk17         int32  `json:"unk17"`
	Unk18         int16  `json:"unk18"`
	Unk19         int16  `json:"unk19"`
	Description1  string `json:"description1"`
	Description2  string `json:"description2"`
	Description3  string `json:"description3"`
	MhfY          string `json:"mhf_y"`
}

type EquipmentType struct {
	Type        string
	From        int64
	To          int64
	Name        int64
	Description int
}
