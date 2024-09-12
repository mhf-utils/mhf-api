package binary

import (
	"encoding/binary"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/japanese"
)

type BinaryFile struct {
	BaseStream io.ReadWriteSeeker
}

func NewBinaryFile(file *os.File) *BinaryFile {
	return &BinaryFile{
		BaseStream: file,
	}
}

func GetBinaryFile(path string) *BinaryFile {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Printf("GetBinaryFile: An error occurred while opening the file")
		return nil
	}
	log.Printf("GetBinaryFile: binary file opened successfully")
	return NewBinaryFile(file)
}

func (binary_file *BinaryFile) ReadByte() (byte, error) {
	var b [1]byte
	_, err := binary_file.BaseStream.Read(b[:])
	return b[0], err
}

func (binary_file *BinaryFile) ReadByteSafe() byte {
	value, err := binary_file.ReadByte()
	if err != nil {
		log.Printf("Error reading byte: %v", err)
	}
	return value
}

func (binary_file *BinaryFile) ReadInt32() (int32, error) {
	var b [4]byte
	_, err := binary_file.BaseStream.Read(b[:])
	if err != nil {
		return 0, err
	}
	return int32(binary.LittleEndian.Uint32(b[:])), nil
}

func (binary_file *BinaryFile) ReadInt16() (int16, error) {
	var b [2]byte
	_, err := binary_file.BaseStream.Read(b[:])
	if err != nil {
		return 0, err
	}
	return int16(binary.LittleEndian.Uint16(b[:])), nil
}

func (binary_file *BinaryFile) ReadUInt16() (uint16, error) {
	var b [2]byte
	_, err := binary_file.BaseStream.Read(b[:])
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint16(b[:]), nil
}

func (binary_file *BinaryFile) ReadUInt32() (uint32, error) {
	var b [4]byte
	_, err := binary_file.BaseStream.Read(b[:])
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(b[:]), nil
}

func (binary_file *BinaryFile) ReadBool() (bool, error) {
	b, err := binary_file.ReadByte()
	if err != nil {
		return false, err
	}
	return b != 0, nil
}

func (binary_file *BinaryFile) ReadSByte() (int8, error) {
	var b [1]byte
	_, err := binary_file.BaseStream.Read(b[:])
	if err != nil {
		return 0, err
	}
	return int8(b[0]), nil
}

func (binary_file *BinaryFile) ReadBytesUntilNull() ([]byte, error) {
	var result []byte
	buffer := make([]byte, 1)

	for {
		_, err := binary_file.BaseStream.Read(buffer)
		if err != nil {
			return nil, err
		}
		if buffer[0] == 0x00 {
			break
		}
		result = append(result, buffer[0])
	}

	return result, nil
}

func (binary_file *BinaryFile) Close() error {
	if closer, ok := binary_file.BaseStream.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

func (binary_file *BinaryFile) Sync() error {
	if err := binary_file.BaseStream.(*os.File).Sync(); err != nil {
		log.Printf("Failed to sync file: %v", err)
	}
	return nil
}

func (binary_file *BinaryFile) ReadString() (string, error) {
	var str []byte
	for {
		b, err := binary_file.ReadByte()
		if err != nil {
			return "", err
		}
		if b == 0 {
			break
		}
		str = append(str, b)
	}
	return string(str), nil
}

func (binary_file *BinaryFile) ReadStringFromPointer() string {
	off, _ := binary_file.ReadInt32()
	pos, _ := binary_file.BaseStream.Seek(0, io.SeekCurrent)
	_, _ = binary_file.BaseStream.Seek(int64(off), io.SeekStart)
	str, _ := binary_file.ReadNullTerminatedString(japanese.ShiftJIS.NewDecoder())
	str = strings.ReplaceAll(str, "\n", "<NL>")
	_, _ = binary_file.BaseStream.Seek(pos, io.SeekStart)
	return str
}

func (binary_file *BinaryFile) ReadNullTerminatedString(decoder *encoding.Decoder) (string, error) {
	var result []byte
	for {
		b, err := binary_file.ReadByte()
		if err != nil {
			return "", err
		}
		if b == 0 {
			break
		}
		result = append(result, b)
	}
	decoded, err := decoder.Bytes(result)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}
