package controllers

import (
	"io"
	"mhf-api/core"
	"mhf-api/server/common"
	"mhf-api/server/models"
	"mhf-api/utils/binary"
	"mhf-api/utils/logger"
	"mhf-api/utils/pointers"
	"net/http"
)

type ControllerItem struct {
	log         *logger.Logger
	binary_file *binary.BinaryFile
}

func NewControllerItem(log *logger.Logger, binary_file *binary.BinaryFile) *ControllerItem {
	return &ControllerItem{
		log,
		binary_file,
	}
}

func (controller *ControllerItem) List(res http.ResponseWriter, req *http.Request) {
	extractFunc := func(index int) models.Item {
		return controller.getEntryByIndex(index)
	}
	core.Paginate(res, req, pointers.ItemFrom, pointers.ItemTo, pointers.ItemLength, extractFunc)
}

func (controller *ControllerItem) Read(res http.ResponseWriter, req *http.Request) {
	extractFunc := func(id int) models.Item {
		return core.Read(id, controller.getEntryByIndex)
	}
	core.ReadItem(res, req, extractFunc)
}

func (controller *ControllerItem) getEntry() models.Item {
	var entry models.Item
	entry.Unk00, _ = controller.binary_file.ReadByte()
	entry.Unk01, _ = controller.binary_file.ReadByte()
	rarity, _ := controller.binary_file.ReadByte()
	entry.Rarity = common.GetRarity(rarity)
	entry.MaxStack, _ = controller.binary_file.ReadByte()
	entry.Unk04, _ = controller.binary_file.ReadByte()
	iconId, _ := controller.binary_file.ReadByte()
	entry.IconName = common.GetIconName(iconId)
	entry.IconColor, _ = controller.binary_file.ReadByte()
	entry.Unk07, _ = controller.binary_file.ReadByte()
	entry.Unk08, _ = controller.binary_file.ReadInt16()
	entry.Unk0A, _ = controller.binary_file.ReadInt16()
	entry.BuyPrice, _ = controller.binary_file.ReadInt32()
	entry.SellPrice, _ = controller.binary_file.ReadInt32()
	entry.Type, _ = controller.binary_file.ReadInt16() // 04 = Consumable, 05 = Placeable
	entry.Unk16, _ = controller.binary_file.ReadInt16()
	entry.Unk18, _ = controller.binary_file.ReadInt16()
	entry.Unk1A, _ = controller.binary_file.ReadByte()
	entry.Unk1B, _ = controller.binary_file.ReadByte()
	entry.Unk1C, _ = controller.binary_file.ReadInt16()
	entry.Unk1E, _ = controller.binary_file.ReadByte()
	entry.Unk1F, _ = controller.binary_file.ReadByte()
	entry.Unk20, _ = controller.binary_file.ReadInt16()
	entry.Unk22, _ = controller.binary_file.ReadInt16()

	return entry
}

func (controller *ControllerItem) getEntryByIndex(index int) models.Item {
	cursor_from := int(pointers.ItemFrom) + index*pointers.ItemLength
	controller.binary_file.BaseStream.Seek(int64(cursor_from), io.SeekStart)
	entry := controller.getEntry()
	entry.Index = index
	entry.ID = cursor_from
	entry.Name = controller.getName(index, pointers.ItemName)
	entry.Description = controller.getDescriptions(index, pointers.ItemDesc)
	entry.Source = controller.getSource(index, pointers.ItemSource)

	return entry
}

func (controller *ControllerItem) getName(index int, pointer int64) string {
	cursor := int(pointer) + index*0x4
	controller.binary_file.BaseStream.Seek(int64(cursor), io.SeekStart)
	return controller.binary_file.ReadStringFromPointer()
}

func (controller *ControllerItem) getDescriptions(index int, pointer int64) string {
	cursor := int(pointer) + index*0x4
	controller.binary_file.BaseStream.Seek(int64(cursor), io.SeekStart)
	return controller.binary_file.ReadStringFromPointer()
}

func (controller *ControllerItem) getSource(index int, pointer int64) string {
	cursor := int(pointer) + index*0x4
	controller.binary_file.BaseStream.Seek(int64(cursor), io.SeekStart)
	return controller.binary_file.ReadStringFromPointer()
}
