package processor

import (
	"encoding/binary"
	"errors"
)

type Box struct {
	Size uint32
	Type [4]byte
	Data []byte
}

const SIZE_INDICATOR_SIZE uint32 = 4

func parseBox(boxType [4]byte, data []byte) (Box, []byte, error) {
	dataSize := uint32(len(data))
	boxTypeLen := uint32(len(boxType))
	minimalSize := SIZE_INDICATOR_SIZE + boxTypeLen
	boxTypeString := string(boxType[:])
	if dataSize < minimalSize {
		return Box{}, data, errors.New("Failed to parse " + boxTypeString)
	}

	boxSize := binary.BigEndian.Uint32(data[:SIZE_INDICATOR_SIZE])
	data = data[SIZE_INDICATOR_SIZE:]

	boxTypeFromData := string(data[:len(boxType)])
	if boxTypeString != boxTypeFromData {
		return Box{}, data, errors.New("Failed to parse " + boxTypeString + " Found " + boxTypeFromData)
	}
	data = data[len(boxType):]
	leftBytesToParse := boxSize - SIZE_INDICATOR_SIZE - boxTypeLen

	boxData := data[:leftBytesToParse]
	restOfData := data[leftBytesToParse:]

	return Box{Size: boxSize, Type: boxType, Data: boxData}, restOfData, nil
}

func GetInitializationSegment(file []byte) ([]byte, []byte, error) {
	var BOX_TYPES = [][4]byte{
		{'f', 't', 'y', 'p'},
		{'m', 'o', 'o', 'v'},
	}

	remainingData := file
	var size uint32 = 0
	for _, boxType := range BOX_TYPES {
		box, remaining, err := parseBox(boxType, remainingData)
		if err != nil {
			return remainingData, remaining, err
		}

		size = size + box.Size
		remainingData = remaining
	}

	return file[:size], remainingData, nil
}
