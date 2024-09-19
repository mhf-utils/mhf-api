package controllers

import (
	"io"
	"mhf-api/core"
	"mhf-api/server/mhfdat/models"
	"mhf-api/server/mhfdat/shared"
	"mhf-api/utils/binary"
	"mhf-api/utils/logger"
	"mhf-api/utils/pointers"
	"net/http"
)

type ControllerWeaponMelee struct {
	log         *logger.Logger
	binary_file *binary.BinaryFile
}

func NewControllerWeaponMelee(log *logger.Logger, binary_file *binary.BinaryFile) *ControllerWeaponMelee {
	return &ControllerWeaponMelee{
		log,
		binary_file,
	}
}

func (controller *ControllerWeaponMelee) List(res http.ResponseWriter, req *http.Request) {
	extractFunc := func(index int) models.Melee {
		return controller.getEntryByIndex(index)
	}
	core.Paginate(res, req, pointers.WeaponMeleeFrom, pointers.WeaponMeleeTo, pointers.WeaponMeleeLength, extractFunc)
}

func (controller *ControllerWeaponMelee) Read(res http.ResponseWriter, req *http.Request) {
	extractFunc := func(id int) models.Melee {
		return core.Read(id, controller.getEntryByIndex)
	}
	core.ReadItem(res, req, extractFunc)
}

func (controller *ControllerWeaponMelee) getEntry() models.Melee {
	var entry models.Melee
	modelId, _ := controller.binary_file.ReadInt16()
	entry.Model = shared.GetModelIdData(modelId)
	rarity, _ := controller.binary_file.ReadByte()
	entry.Rarity = shared.GetRarity(rarity)
	classId, _ := controller.binary_file.ReadByte()
	entry.Type = shared.GetType(classId)
	entry.ZennyCost, _ = controller.binary_file.ReadInt32()
	sharpnessIdByte, _ := controller.binary_file.ReadByte()
	entry.SharpnessId = int8(sharpnessIdByte)
	sharpnessMaxByte, _ := controller.binary_file.ReadByte()
	entry.SharpnessMax = int8(sharpnessMaxByte)
	RawDamage, _ := controller.binary_file.ReadInt16()
	entry.RawDamage = shared.GetRawDamage(entry.Type, RawDamage)
	entry.Defense, _ = controller.binary_file.ReadInt16()
	affinityByte, _ := controller.binary_file.ReadByte()
	entry.Affinity = int8(affinityByte)
	elementId, _ := controller.binary_file.ReadByte()
	entry.Element = shared.GetElementName(elementId)
	eleDamageByte, _ := controller.binary_file.ReadByte()
	entry.EleDamage = byte(eleDamageByte) * 10
	ailmentId, _ := controller.binary_file.ReadByte()
	entry.Ailment = shared.GetAilmentName(ailmentId)
	ailDamageByte, _ := controller.binary_file.ReadByte()
	entry.AilDamage = byte(ailDamageByte) * 10
	entry.Slots, _ = controller.binary_file.ReadByte()
	entry.WeaponAttribute, _ = controller.binary_file.ReadByte()
	entry.Unk4, _ = controller.binary_file.ReadByte()
	entry.UpgradePath, _ = controller.binary_file.ReadInt16()
	entry.Unk5, _ = controller.binary_file.ReadInt16()
	entry.EqType, _ = controller.binary_file.ReadInt16()
	entry.Length, _ = controller.binary_file.ReadInt32()
	entry.WeaponType, _ = controller.binary_file.ReadInt32()
	entry.VisualEffects, _ = controller.binary_file.ReadInt16()
	entry.Unk11, _ = controller.binary_file.ReadInt16()
	entry.Unk12, _ = controller.binary_file.ReadByte()
	entry.Unk13, _ = controller.binary_file.ReadByte()
	entry.Unk14, _ = controller.binary_file.ReadByte()
	entry.Unk15, _ = controller.binary_file.ReadByte()
	entry.Unk16, _ = controller.binary_file.ReadInt32()
	entry.Zenith_skill, _ = controller.binary_file.ReadInt32()

	return entry
}

func (controller *ControllerWeaponMelee) getEntryByIndex(index int) models.Melee {
	cursor := int(pointers.WeaponMeleeFrom) + index*pointers.WeaponMeleeLength
	controller.binary_file.BaseStream.Seek(int64(cursor), io.SeekStart)
	entry := controller.getEntry()
	entry.Index = index
	entry.ID = cursor
	entry.Name = controller.getName(index, pointers.WeaponMeleeName)
	entry.Sharpness = controller.getSharpness(entry.SharpnessId, shared.GetSharpnessPointerByType(entry.Type))
	entry.Description1,
		entry.Description2,
		entry.Description3,
		entry.MhfY = controller.getDescriptions(index, pointers.WeaponMeleeDesc)

	return entry
}

func (controller *ControllerWeaponMelee) getName(index int, pointer int64) string {
	cursor := int(pointer) + index*0x4
	controller.binary_file.BaseStream.Seek(int64(cursor), io.SeekStart)
	return controller.binary_file.ReadStringFromPointer()
}

func (controller *ControllerWeaponMelee) getDescriptions(index int, pointer int64) (string, string, string, string) {
	cursor := int(pointer) + index*0x4
	controller.binary_file.BaseStream.Seek(int64(cursor), io.SeekStart)
	Description1 := controller.binary_file.ReadStringFromPointer()
	Description2 := controller.binary_file.ReadStringFromPointer()
	Description3 := controller.binary_file.ReadStringFromPointer()
	MhfY := controller.binary_file.ReadStringFromPointer()
	return Description1, Description2, Description3, MhfY
}

func (controller *ControllerWeaponMelee) getSharpness(index int8, pointer int64) models.Sharpness {
	cursor := int(pointer) + int(index)*pointers.WeaponSharpnessLength
	controller.binary_file.BaseStream.Seek(int64(cursor), io.SeekStart)

	var entry models.Sharpness
	entry.ID = cursor
	entry.Index = int(index)
	entry.Red, _ = controller.binary_file.ReadInt16()
	entry.Orange, _ = controller.binary_file.ReadInt16()
	entry.Yellow, _ = controller.binary_file.ReadInt16()
	entry.Green, _ = controller.binary_file.ReadInt16()
	entry.Blue, _ = controller.binary_file.ReadInt16()
	entry.White, _ = controller.binary_file.ReadInt16()
	entry.Purple, _ = controller.binary_file.ReadInt16()
	entry.SkyBlue, _ = controller.binary_file.ReadInt16()
	return entry
}
