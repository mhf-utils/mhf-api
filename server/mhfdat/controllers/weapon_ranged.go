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

type ControllerWeaponRanged struct {
	log         *logger.Logger
	binary_file *binary.BinaryFile
}

func NewControllerWeaponRanged(log *logger.Logger, binary_file *binary.BinaryFile) *ControllerWeaponRanged {
	return &ControllerWeaponRanged{
		log,
		binary_file,
	}
}

func (controller *ControllerWeaponRanged) List(res http.ResponseWriter, req *http.Request) {
	extractFunc := func(index int) models.Ranged {
		return controller.getEntryByIndex(index)
	}
	core.Paginate(res, req, pointers.WeaponRangedFrom, pointers.WeaponRangedTo, pointers.WeaponRangedLength, extractFunc)
}

func (controller *ControllerWeaponRanged) Read(res http.ResponseWriter, req *http.Request) {
	extractFunc := func(id int) models.Ranged {
		return core.Read(id, controller.getEntryByIndex)
	}
	core.ReadItem(res, req, extractFunc)
}

func (controller *ControllerWeaponRanged) getEntry() models.Ranged {
	var entry models.Ranged
	modelId, _ := controller.binary_file.ReadInt16()
	entry.Model = shared.GetModelIdData(modelId)
	rarity, _ := controller.binary_file.ReadByte()
	entry.Rarity = shared.GetRarity(rarity)
	entry.MaxSlotsMaybe, _ = controller.binary_file.ReadByte()
	classId, _ := controller.binary_file.ReadByte()
	entry.Type = shared.GetType(classId)
	entry.Unk05, _ = controller.binary_file.ReadByte()
	entry.Eq, _ = controller.binary_file.ReadByte()
	entry.Unk07, _ = controller.binary_file.ReadByte()
	entry.Unk08, _ = controller.binary_file.ReadUInt32()
	entry.Unk0C, _ = controller.binary_file.ReadUInt32()
	entry.Unk10, _ = controller.binary_file.ReadUInt32()
	entry.ZennyCost, _ = controller.binary_file.ReadUInt32()
	entry.RawDamage, _ = controller.binary_file.ReadUInt16()
	entry.Defense, _ = controller.binary_file.ReadUInt16()
	entry.RecoilMaybe, _ = controller.binary_file.ReadByte()
	entry.Slots, _ = controller.binary_file.ReadByte()
	entry.Affinity, _ = controller.binary_file.ReadByte()
	entry.SortOrderMaybe, _ = controller.binary_file.ReadByte()
	entry.Unk20, _ = controller.binary_file.ReadByte()
	elementId, _ := controller.binary_file.ReadByte()
	entry.Element = shared.GetElementName(elementId)
	entry.EleDamage, _ = controller.binary_file.ReadByte()
	entry.Unk23, _ = controller.binary_file.ReadByte()
	entry.Unk24, _ = controller.binary_file.ReadUInt32()
	entry.Unk2C, _ = controller.binary_file.ReadUInt32()
	entry.Unk34, _ = controller.binary_file.ReadUInt32()

	return entry
}

func (controller *ControllerWeaponRanged) getEntryByIndex(index int) models.Ranged {
	cursor_from := int(pointers.WeaponRangedFrom) + index*pointers.WeaponRangedLength
	controller.binary_file.BaseStream.Seek(int64(cursor_from), io.SeekStart)
	entry := controller.getEntry()
	entry.Index = index
	entry.ID = cursor_from
	entry.Name = controller.getName(index, pointers.WeaponRangedName)
	entry.Description1,
		entry.Description2,
		entry.Description3,
		entry.MhfY = controller.getDescriptions(index, pointers.WeaponRangedDesc)

	return entry
}

func (controller *ControllerWeaponRanged) getName(index int, pointer int64) string {
	cursor := int(pointer) + index*0x4
	controller.binary_file.BaseStream.Seek(int64(cursor), io.SeekStart)
	return controller.binary_file.ReadStringFromPointer()
}

func (controller *ControllerWeaponRanged) getDescriptions(index int, pointer int64) (string, string, string, string) {
	cursor := int(pointer) + index*0x4
	controller.binary_file.BaseStream.Seek(int64(cursor), io.SeekStart)
	Description1 := controller.binary_file.ReadStringFromPointer()
	Description2 := controller.binary_file.ReadStringFromPointer()
	Description3 := controller.binary_file.ReadStringFromPointer()
	MhfY := controller.binary_file.ReadStringFromPointer()
	return Description1, Description2, Description3, MhfY
}
