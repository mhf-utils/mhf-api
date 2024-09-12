package common

import (
	"fmt"
)

func GetGenderName(isMaleEquip, isFemaleEquip bool) string {
	if isMaleEquip && isFemaleEquip {
		return "mixte"
	}
	if isMaleEquip {
		return "male"
	}
	if isFemaleEquip {
		return "female"
	}
	return "unknown"
}

func GetRoleName(isBladeEquip, isGunnerEquip bool) string {
	if isBladeEquip && isGunnerEquip {
		return "mixte"
	}
	if isBladeEquip {
		return "sworder"
	}
	if isGunnerEquip {
		return "gunner"
	}
	return "unknown"
}

func GetIconName(n uint8) string {
	icon_map := map[uint8]string{
		0x00: "smoke",
		0x01: "orb",
		0x02: "bomb",
		0x03: "meat",
		0x04: "fish_bait",
		0x05: "fish",
		0x06: "box",
		0x07: "whetstone",
		0x08: "dung",
		0x09: "monster",
		0x0A: "bones",
		0x0B: "binoculars",
		0x0C: "mushroom",
		0x0D: "bugnet",
		0x0E: "pelt",
		0x0F: "leaf",
		0x10: "pickaxe",
		0x11: "barrel",
		0x12: "seed",
		0x13: "bbq_spit",
		0x14: "insect",
		0x15: "trap",
		0x16: "net",
		0x17: "scale",
		0x18: "drink",
		0x19: "egg",
		0x1A: "ammo",
		0x1B: "stone",
		0x1C: "husk",
		0x1D: "map",
		0x1E: "flute",
		0x1F: "fang",
		0x20: "grid",
		0x21: "question_mark",
		0x22: "coin",
		0x23: "sac",
		0x24: "book",
		0x25: "ticket",
		0x26: "knife",
		0x29: "music_sheet",
		0x2B: "jewel",
		0x2C: "house",
		0x2D: "plant",
		0x2F: "mocha",
		0x30: "pot",
		0x31: "boomerang",
		0x32: "coating",
		0x33: "empty_bottle",
		0x34: "carapace",
		0x4F: "sword_crystal",
		0x50: "potion",
		0x52: "fruit",
		0x56: "tower_w_sigil",
		0x5B: "tower_a_sigil",
		0x5C: "mantle",
		0x5D: "armor_sphere",
	}

	if name, ok := icon_map[n]; ok {
		return name
	}
	return "unknown"
}

func GetIconId(name string) uint8 {
	icon_map := map[string]uint8{
		"smoke":         0x00,
		"orb":           0x01,
		"bomb":          0x02,
		"meat":          0x03,
		"fish_bait":     0x04,
		"fish":          0x05,
		"box":           0x06,
		"whetstone":     0x07,
		"dung":          0x08,
		"monster":       0x09,
		"bones":         0x0A,
		"binoculars":    0x0B,
		"mushroom":      0x0C,
		"bugnet":        0x0D,
		"pelt":          0x0E,
		"leaf":          0x0F,
		"pickaxe":       0x10,
		"barrel":        0x11,
		"seed":          0x12,
		"bbq_spit":      0x13,
		"insect":        0x14,
		"trap":          0x15,
		"net":           0x16,
		"scale":         0x17,
		"drink":         0x18,
		"egg":           0x19,
		"ammo":          0x1A,
		"stone":         0x1B,
		"husk":          0x1C,
		"map":           0x1D,
		"flute":         0x1E,
		"fang":          0x1F,
		"grid":          0x20,
		"question_mark": 0x21,
		"coin":          0x22,
		"sac":           0x23,
		"book":          0x24,
		"ticket":        0x25,
		"knife":         0x26,
		"music_sheet":   0x29,
		"jewel":         0x2B,
		"house":         0x2C,
		"plant":         0x2D,
		"mocha":         0x2F,
		"pot":           0x30,
		"boomerang":     0x31,
		"coating":       0x32,
		"empty_bottle":  0x33,
		"carapace":      0x34,
		"sword_crystal": 0x4F,
		"potion":        0x50,
		"fruit":         0x52,
		"tower_w_sigil": 0x56,
		"tower_a_sigil": 0x5B,
		"mantle":        0x5C,
		"armor_sphere":  0x5D,
	}
	if id, ok := icon_map[name]; ok {
		return id
	}
	return 0xFF
}

func GetAilmentName(value uint8) string {
	ailment_map := map[uint8]string{
		0x00: "none",
		0x01: "poison",
		0x02: "paralysis",
		0x03: "sleep",
		0x04: "blast",
	}

	if name, ok := ailment_map[value]; ok {
		return name
	}
	return "unknown"
}

func GetElementName(value uint8) string {
	element_map := map[uint8]string{
		0x00: "none",
		0x01: "fire",
		0x02: "water",
		0x03: "thunder",
		0x04: "dragon",
		0x05: "ice",
		0x06: "flame",
		0x07: "light",
		0x08: "thunder_pole",
		0x09: "tenshou",
		0x0a: "okiko",
		0x0b: "black_flame",
		0x0c: "kanade",
		0x0d: "darkness",
		0x0e: "crimson_demon",
		0x0f: "wind",
		0x10: "sound",
		0x11: "burning_zero",
		0x12: "emperors_roar",
	}

	if name, ok := element_map[value]; ok {
		return name
	}
	return "unknown"
}

func GetModelIdData(modelId int16) string {
	if modelId < 0 || modelId >= 10000 {
		return "Unmapped"
	}
	prefixes := []string{
		"we",
		"wf",
		"wg",
		"wh",
		"wi",
		"wj",
		"wk",
		"wl",
		"wm",
		"wn",
	}
	prefixIndex := modelId / 1000
	offset := modelId % 1000

	return fmt.Sprintf("%s%03d", prefixes[prefixIndex], offset)
}

func GetType(value byte) string {
	type_map := map[byte]string{
		0x00: "great_sword",
		0x01: "heavy_bowgun",
		0x02: "hammer",
		0x03: "lance",
		0x04: "sword_and_shield",
		0x05: "light_bowgun",
		0x06: "dual_blades",
		0x07: "long_sword",
		0x08: "hunting_horn",
		0x09: "gunlance",
		0x0a: "bow",
		0x0b: "tonfa",
		0x0c: "switch_axe",
		0x0d: "magnet_spike",
		0x0e: "unknown",
	}

	if name, ok := type_map[value]; ok {
		return name
	}
	return "unknown"
}

func GetRawDamage(Type string, value int16) int16 {
	raw_map := map[string]float64{
		"great_sword":      4.8,
		"heavy_bowgun":     1.2,
		"hammer":           5.2,
		"lance":            2.3,
		"sword_and_shield": 1.4,
		"light_bowgun":     1.2,
		"dual_blades":      1.4,
		"long_sword":       4.8,
		"hunting_horn":     5.2,
		"gunlance":         2.3,
		"bow":              1.2,
		"tonfa":            1.8,
		"switch_axe":       5.4,
		"magnet_spike":     5.4,
		"unknown":          1.0,
	}

	if coeff, ok := raw_map[Type]; ok {
		return int16(float64(value) * coeff)
	}
	return 0
}

func GetSharpnessPointerByType(Type string) int64 {
	type_map := map[string]int64{
		"great_sword":      int64(0x01902D00),
		"heavy_bowgun":     int64(0x01902D00),
		"hammer":           int64(0x01902500),
		"lance":            int64(0x01901D00),
		"sword_and_shield": int64(0x01901500),
		"light_bowgun":     int64(0x01902D00),
		"dual_blades":      int64(0x01900D00),
		"long_sword":       int64(0x01902D00),
		"hunting_horn":     int64(0x01902500),
		"gunlance":         int64(0x01901D00),
		"bow":              int64(0x01902D00),
		"tonfa":            int64(0x01900500),
		"switch_axe":       int64(0x018FFD00),
		"magnet_spike":     int64(0x018FF500),
	}
	id, _ := type_map[Type]
	return id
}

// SETTERS

func SetElementId(name string) uint8 {
	element_map := map[string]uint8{
		"none":          0x00,
		"fire":          0x01,
		"water":         0x02,
		"thunder":       0x03,
		"dragon":        0x04,
		"ice":           0x05,
		"flame":         0x06,
		"light":         0x07,
		"thunder_pole":  0x08,
		"tenshou":       0x09,
		"okiko":         0x0a,
		"black_flame":   0x0b,
		"kanade":        0x0c,
		"darkness":      0x0d,
		"crimson_demon": 0x0e,
		"wind":          0x0f,
		"sound":         0x10,
		"burning_zero":  0x11,
		"emperors_roar": 0x12,
	}
	if id, ok := element_map[name]; ok {
		return id
	}
	return 0xFF
}

func SetAilmentId(name string) uint8 {
	ailment_map := map[string]uint8{
		"none":      0x00,
		"poison":    0x01,
		"paralysis": 0x02,
		"sleep":     0x03,
		"blast":     0x04,
	}
	if id, ok := ailment_map[name]; ok {
		return id
	}
	return 0xFF
}

func SetTypeId(name string) byte {
	type_map := map[string]byte{
		"great_sword":      0x00,
		"heavy_bowgun":     0x01,
		"hammer":           0x02,
		"lance":            0x03,
		"sword_and_shield": 0x04,
		"light_bowgun":     0x05,
		"dual_blades":      0x06,
		"long_sword":       0x07,
		"hunting_horn":     0x08,
		"gunlance":         0x09,
		"bow":              0x0a,
		"tonfa":            0x0b,
		"switch_axe":       0x0c,
		"magnet_spike":     0x0d,
		"unknown":          0x0e,
	}
	if id, ok := type_map[name]; ok {
		return id
	}
	return 0xFF
}

func SetModelId(model string) int16 {
	if len(model) != 5 {
		return -1
	}
	prefixes := map[string]int16{
		"we": 0,
		"wf": 1,
		"wg": 2,
		"wh": 3,
		"wi": 4,
		"wj": 5,
		"wk": 6,
		"wl": 7,
		"wm": 8,
		"wn": 9,
	}
	prefix := model[:2]
	offset := model[2:]

	var offsetValue int16
	_, err := fmt.Sscanf(offset, "%03d", &offsetValue)
	if err != nil {
		return -1
	}

	prefixValue, ok := prefixes[prefix]
	if !ok {
		return -1
	}

	return prefixValue*1000 + offsetValue
}

func SetRawDamage(Type string, value int16) int16 {
	raw_map := map[string]float64{
		"great_sword":      4.8,
		"heavy_bowgun":     1.2,
		"hammer":           5.2,
		"lance":            2.3,
		"sword_and_shield": 1.4,
		"light_bowgun":     1.2,
		"dual_blades":      1.4,
		"long_sword":       4.8,
		"hunting_horn":     5.2,
		"gunlance":         2.3,
		"bow":              1.2,
		"tonfa":            1.8,
		"switch_axe":       5.4,
		"magnet_spike":     5.4,
		"unknown":          1.0,
	}

	if coeff, ok := raw_map[Type]; ok {
		return int16(float64(value) / coeff)
	}
	return 0
}

func GetRarity(rarity byte) int {
	return int(rarity + 1)
}

func GetEquipType(n uint8) string {
	equip_type_map := map[uint8]string{
		0x00: "leg",
		0x02: "helm",
		0x03: "chest",
		0x04: "arm",
		0x05: "waist",
		0x06: "melee",
		0x07: "ranged",
	}

	if name, ok := equip_type_map[n]; ok {
		return name
	}
	return "unknown"
}

func GetRarityValue(rarity int) byte {
	return byte(rarity - 1)
}

func GetEquipTypeId(name string) uint8 {
	equip_type_map := map[string]uint8{
		"leg":    0x00,
		"helm":   0x02,
		"chest":  0x03,
		"arm":    0x04,
		"waist":  0x05,
		"melee":  0x06,
		"ranged": 0x07,
	}
	if id, ok := equip_type_map[name]; ok {
		return id
	}
	return 0xFF
}
