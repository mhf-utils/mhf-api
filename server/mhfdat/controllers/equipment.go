package controllers

import (
	"encoding/json"
	"io"
	"mhf-api/core"
	"mhf-api/server/mhfdat/models"
	"mhf-api/server/mhfdat/shared"
	"mhf-api/utils/binary"
	"mhf-api/utils/logger"
	"mhf-api/utils/pointers"
	"net/http"

	"github.com/gorilla/mux"
)

var equipment_types = map[string]models.EquipmentType{
	"helm": {
		From:        pointers.EquipmentHelmFrom,
		To:          pointers.EquipmentHelmTo,
		Name:        pointers.EquipmentHelmName,
		Description: pointers.EquipmentHelmDesc,
	},
	"chest": {
		From:        pointers.EquipmentChestFrom,
		To:          pointers.EquipmentChestTo,
		Name:        pointers.EquipmentChestName,
		Description: pointers.EquipmentChestDesc,
	},
	"arm": {
		From:        pointers.EquipmentArmFrom,
		To:          pointers.EquipmentArmTo,
		Name:        pointers.EquipmentArmName,
		Description: pointers.EquipmentArmDesc,
	},
	"waist": {
		From:        pointers.EquipmentWaistFrom,
		To:          pointers.EquipmentWaistTo,
		Name:        pointers.EquipmentWaistName,
		Description: pointers.EquipmentWaistDesc,
	},
	"leg": {
		From:        pointers.EquipmentLegFrom,
		To:          pointers.EquipmentLegTo,
		Name:        pointers.EquipmentLegName,
		Description: pointers.EquipmentLegDesc,
	},
}

type ControllerEquipment struct {
	log         *logger.Logger
	binary_file *binary.BinaryFile
}

func NewControllerEquipment(log *logger.Logger, binary_file *binary.BinaryFile) *ControllerEquipment {
	return &ControllerEquipment{
		log,
		binary_file,
	}
}

func (controller *ControllerEquipment) List(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	if params["type"] == "" {
		json.NewEncoder(res).Encode("Query params 'type' is required. Available: [ helm, chest, arm, waist, leg ]")
		return
	}

	equipment_type, exists := equipment_types[params["type"]]
	if !exists {
		json.NewEncoder(res).Encode("Query params 'type' is required. Available: [ helm, chest, arm, waist, leg ]")
		return
	}

	extractFunc := func(index int) models.Equipment {
		return controller.getEntryByIndex(index, models.EquipmentType{
			Type:        params["type"],
			From:        equipment_type.From,
			To:          equipment_type.To,
			Name:        equipment_type.Name,
			Description: equipment_type.Description,
		})
	}
	core.Paginate(res, req, equipment_type.From, equipment_type.To, pointers.EquipmentLength, extractFunc)
}

func (controller *ControllerEquipment) Read(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	if params["type"] == "" {
		json.NewEncoder(res).Encode("Query params 'type' is required. Available: [ helm, chest, arm, waist, leg ]")
		return
	}

	equipment_type, exists := equipment_types[params["type"]]
	if !exists {
		json.NewEncoder(res).Encode("Query params 'type' is required. Available: [ helm, chest, arm, waist, leg ]")
		return
	}

	extractFunc := func(id int) models.Equipment {
		return controller.getEntryByIndex(id, models.EquipmentType{
			Type:        params["type"],
			From:        equipment_type.From,
			To:          equipment_type.To,
			Name:        equipment_type.Name,
			Description: equipment_type.Description,
		})
	}
	core.ReadItem(res, req, extractFunc)
}

func (controller *ControllerEquipment) getEntry(armor_type string) models.Equipment {
	var entry models.Equipment
	entry.Type = armor_type
	entry.ModelIdMale, _ = controller.binary_file.ReadInt16()
	entry.ModelIdFemale, _ = controller.binary_file.ReadInt16()
	bitfield, _ := controller.binary_file.ReadByte()
	entry.IsHelm = armor_type == "helm"
	entry.IsChest = armor_type == "chest"
	entry.IsArm = armor_type == "arm"
	entry.IsWaist = armor_type == "waist"
	entry.IsLeg = armor_type == "leg"
	isMaleEquip := (bitfield & (1 << 0)) != 0
	isFemaleEquip := (bitfield & (1 << 1)) != 0
	entry.Gender = shared.GetGenderName(isMaleEquip, isFemaleEquip)
	isBladeEquip := (bitfield & (1 << 2)) != 0
	isGunnerEquip := (bitfield & (1 << 3)) != 0
	entry.Role = shared.GetRoleName(isBladeEquip, isGunnerEquip)
	entry.Bool1 = (bitfield & (1 << 4)) != 0
	entry.IsSPEquip = (bitfield & (1 << 5)) != 0
	entry.Bool3 = (bitfield & (1 << 6)) != 0
	entry.Bool4 = (bitfield & (1 << 7)) != 0
	rarity, _ := controller.binary_file.ReadByte()
	entry.Rarity = shared.GetRarity(rarity)
	entry.MaxLevel, _ = controller.binary_file.ReadByte()
	entry.Unk1_1, _ = controller.binary_file.ReadByte()
	entry.Unk1_2, _ = controller.binary_file.ReadByte()
	entry.Unk1_3, _ = controller.binary_file.ReadByte()
	entry.Unk1_4, _ = controller.binary_file.ReadByte()
	entry.Unk2, _ = controller.binary_file.ReadByte()
	entry.ZennyCost, _ = controller.binary_file.ReadInt32()
	entry.Unk3, _ = controller.binary_file.ReadInt16()
	entry.BaseDefense, _ = controller.binary_file.ReadInt16()
	entry.FireRes, _ = controller.binary_file.ReadByte()
	entry.WaterRes, _ = controller.binary_file.ReadByte()
	entry.ThunderRes, _ = controller.binary_file.ReadByte()
	entry.DragonRes, _ = controller.binary_file.ReadByte()
	entry.IceRes, _ = controller.binary_file.ReadByte()
	entry.Unk3_1, _ = controller.binary_file.ReadInt16()
	entry.BaseSlots, _ = controller.binary_file.ReadByte()
	entry.MaxSlots, _ = controller.binary_file.ReadByte()
	entry.SthEventCrown, _ = controller.binary_file.ReadByte()
	entry.Unk5, _ = controller.binary_file.ReadByte()
	entry.Unk6, _ = controller.binary_file.ReadByte()
	entry.Unk7_1, _ = controller.binary_file.ReadByte()
	entry.Unk7_2, _ = controller.binary_file.ReadByte()
	entry.Unk7_3, _ = controller.binary_file.ReadByte()
	entry.Unk7_4, _ = controller.binary_file.ReadByte()
	entry.Unk8_1, _ = controller.binary_file.ReadByte()
	entry.Unk8_2, _ = controller.binary_file.ReadByte()
	entry.Unk8_3, _ = controller.binary_file.ReadByte()
	entry.Unk8_4, _ = controller.binary_file.ReadByte()
	entry.Unk10, _ = controller.binary_file.ReadInt16()
	entry.SkillId1, _ = controller.binary_file.ReadString()
	entry.SkillPts1, _ = controller.binary_file.ReadByte()
	entry.SkillId2, _ = controller.binary_file.ReadString()
	entry.SkillPts2, _ = controller.binary_file.ReadByte()
	entry.SkillId3, _ = controller.binary_file.ReadString()
	entry.SkillPts3, _ = controller.binary_file.ReadByte()
	entry.SkillId4, _ = controller.binary_file.ReadString()
	entry.SkillPts4, _ = controller.binary_file.ReadByte()
	entry.SkillId5, _ = controller.binary_file.ReadString()
	entry.SkillPts5, _ = controller.binary_file.ReadByte()
	entry.SthHiden, _ = controller.binary_file.ReadInt32()
	entry.Unk12, _ = controller.binary_file.ReadInt32()
	entry.Unk13, _ = controller.binary_file.ReadByte()
	entry.Unk14, _ = controller.binary_file.ReadByte()
	entry.Unk15, _ = controller.binary_file.ReadByte()
	entry.Unk16, _ = controller.binary_file.ReadByte()
	entry.Unk17, _ = controller.binary_file.ReadInt32()
	entry.Unk18, _ = controller.binary_file.ReadInt16()
	entry.Unk19, _ = controller.binary_file.ReadInt16()

	return entry
}

func (controller *ControllerEquipment) getEntryByIndex(index int, equipment_type models.EquipmentType) models.Equipment {
	cursor_from := int(equipment_type.From) + index*pointers.EquipmentLength
	controller.binary_file.BaseStream.Seek(int64(cursor_from), io.SeekStart)
	entry := controller.getEntry(equipment_type.Type)
	entry.Index = index
	entry.ID = cursor_from
	entry.Name = controller.getName(index, equipment_type.Name)
	entry.Description1,
		entry.Description2,
		entry.Description3,
		entry.MhfY = controller.getDescriptions(index, equipment_type.Description)

	return entry
}

func (controller *ControllerEquipment) getName(index int, pointer int64) string {
	cursor := int(pointer) + index*0x4
	controller.binary_file.BaseStream.Seek(int64(cursor), io.SeekStart)
	return controller.binary_file.ReadStringFromPointer()
}

func (controller *ControllerEquipment) getDescriptions(index int, pointer int) (string, string, string, string) {
	cursor := pointer + index*0x4
	controller.binary_file.BaseStream.Seek(int64(cursor), io.SeekStart)
	Description1 := controller.binary_file.ReadStringFromPointer()
	Description2 := controller.binary_file.ReadStringFromPointer()
	Description3 := controller.binary_file.ReadStringFromPointer()
	MhfY := controller.binary_file.ReadStringFromPointer()
	return Description1, Description2, Description3, MhfY
}
