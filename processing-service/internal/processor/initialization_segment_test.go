package processor

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestParseBox(t *testing.T) {
	tests := []struct {
		name           string
		boxType        [4]byte
		data           []byte
		expectedBox    Box
		expectedRemain []byte
		expectedErr    error
	}{
		{
			name:    "Valid box with remaining data",
			boxType: [4]byte{'f', 't', 'y', 'p'},
			data: createTestData(
				16,
				[4]byte{'f', 't', 'y', 'p'},
				[]byte("testdata"),
				[]byte("remaining"),
			),
			expectedBox: Box{
				Size: 16,
				Type: [4]byte{'f', 't', 'y', 'p'},
				Data: []byte("testdata"),
			},
			expectedRemain: []byte("remaining"),
			expectedErr:    nil,
		},
		{
			name:    "Valid box without remaining data",
			boxType: [4]byte{'m', 'o', 'o', 'v'},
			data: createTestData(
				8,
				[4]byte{'m', 'o', 'o', 'v'},
				[]byte{},
				[]byte{},
			),
			expectedBox: Box{
				Size: 8,
				Type: [4]byte{'m', 'o', 'o', 'v'},
				Data: []byte{},
			},
			expectedRemain: []byte{},
			expectedErr:    nil,
		},
		{
			name:           "Data too small",
			boxType:        [4]byte{'f', 't', 'y', 'p'},
			data:           []byte{0, 0, 0},
			expectedBox:    Box{},
			expectedRemain: []byte{0, 0, 0},
			expectedErr:    errors.New("Failed to parse ftyp"),
		},
		{
			name:    "Invalid box type",
			boxType: [4]byte{'f', 't', 'y', 'p'},
			data: createTestData(
				16,
				[4]byte{'m', 'd', 'a', 't'},
				[]byte("testdata"),
				[]byte{},
			),
			expectedBox:    Box{},
			expectedRemain: append([]byte{'m', 'd', 'a', 't'}, []byte("testdata")...),
			expectedErr:    errors.New("Failed to parse ftyp Found mdat"),
		},
		{
			name:    "Special characters in box type",
			boxType: [4]byte{0xFF, 0xFE, 0xFD, 0xFC},
			data: createTestData(
				18,
				[4]byte{0xFF, 0xFE, 0xFD, 0xFC},
				[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				[]byte("remain"),
			),
			expectedBox: Box{
				Size: 18,
				Type: [4]byte{0xFF, 0xFE, 0xFD, 0xFC},
				Data: []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			expectedRemain: []byte("remain"),
			expectedErr:    nil,
		},
		{
			name:    "Box with exact size",
			boxType: [4]byte{'t', 'e', 's', 't'},
			data: createTestData(
				14,
				[4]byte{'t', 'e', 's', 't'},
				[]byte("abcdef"),
				[]byte{},
			),
			expectedBox: Box{
				Size: 14,
				Type: [4]byte{'t', 'e', 's', 't'},
				Data: []byte("abcdef"),
			},
			expectedRemain: []byte{},
			expectedErr:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			box, remain, err := parseBox(tt.boxType, tt.data)
			if (err == nil && tt.expectedErr != nil) || (err != nil && tt.expectedErr == nil) {
				t.Errorf("Error mismatch: got %v, want %v", err, tt.expectedErr)
			} else if err != nil && tt.expectedErr != nil && err.Error() != tt.expectedErr.Error() {
				t.Errorf("Error message mismatch: got %v, want %v", err, tt.expectedErr)
			}

			if tt.expectedErr != nil {
				return
			}

			if !reflect.DeepEqual(box, tt.expectedBox) {
				t.Errorf("Box mismatch: got %+v, want %+v", box, tt.expectedBox)
			}

			if !bytes.Equal(remain, tt.expectedRemain) {
				t.Errorf("Remaining data mismatch: got %v, want %v", remain, tt.expectedRemain)
			}
		})
	}
}

func createTestData(boxSize uint32, boxType [4]byte, boxData []byte, remainingData []byte) []byte {
	sizeBytes := make([]byte, SIZE_INDICATOR_SIZE)
	binary.BigEndian.PutUint32(sizeBytes, boxSize)
	result := append(sizeBytes, boxType[:]...)
	result = append(result, boxData...)
	result = append(result, remainingData...)

	return result
}

func TestGetInitializationSegment(t *testing.T) {
	tests := []struct {
		name                  string
		file                  []byte
		expectedInitSegment   []byte
		expectedRemainingData []byte
		expectedErr           error
	}{
		{
			name: "Valid initialization segment with mdat after",
			file: createTestFile(
				18, [4]byte{'f', 't', 'y', 'p'}, []byte("majorminor"),
				25, [4]byte{'m', 'o', 'o', 'v'}, []byte("moov_box_contents"),
				35, [4]byte{'m', 'd', 'a', 't'}, []byte("actual_media_data_goes_here"),
			),
			expectedInitSegment: createTestFile(
				18, [4]byte{'f', 't', 'y', 'p'}, []byte("majorminor"),
				25, [4]byte{'m', 'o', 'o', 'v'}, []byte("moov_box_contents"),
			),
			expectedRemainingData: createTestFile(
				35, [4]byte{'m', 'd', 'a', 't'}, []byte("actual_media_data_goes_here"),
			),
			expectedErr: nil,
		},
		{
			name: "No initialization segment (missing ftyp box)",
			file: createTestFile(
				25, [4]byte{'m', 'o', 'o', 'v'}, []byte("moov_box_contents"),
				35, [4]byte{'m', 'd', 'a', 't'}, []byte("actual_media_data_goes_here"),
			),
			expectedInitSegment:   nil,
			expectedRemainingData: nil,
			expectedErr:           fmt.Errorf("Failed to parse ftyp Found moov"),
		},
		{
			name: "Empty file",
			file: []byte{},
			expectedInitSegment:   nil,
			expectedRemainingData: nil,
			expectedErr:           fmt.Errorf("Failed to parse ftyp"),
		},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initSegment, remainingData, err := GetInitializationSegment(tt.file)

			if (err == nil && tt.expectedErr != nil) || (err != nil && tt.expectedErr == nil) {
				t.Errorf("Error mismatch: got %v, want %v", err, tt.expectedErr)
			} else if err != nil && tt.expectedErr != nil && err.Error() != tt.expectedErr.Error() {
				t.Errorf("Error message mismatch: got %v, want %v", err, tt.expectedErr)
			}
			if tt.expectedErr != nil {
				return
			}
			if !bytes.Equal(initSegment, tt.expectedInitSegment) {
				t.Errorf("Initialization segment mismatch:\n  got: %v\n want: %v", initSegment, tt.expectedInitSegment)
			}
			if !bytes.Equal(remainingData, tt.expectedRemainingData) {
				t.Errorf("Remaining data mismatch:\n  got: %v\n want: %v", remainingData, tt.expectedRemainingData)
			}
		})
	}
}

func createTestFile(args ...interface{}) []byte {
	var result []byte
	for i := 0; i < len(args); i += 3 {
		boxSize := args[i].(int)
		boxType := args[i+1].([4]byte)
		boxData := args[i+2].([]byte)
		sizeBytes := make([]byte, SIZE_INDICATOR_SIZE)
		binary.BigEndian.PutUint32(sizeBytes, uint32(boxSize))

		result = append(result, sizeBytes...)
		result = append(result, boxType[:]...)
		result = append(result, boxData...)
	}

	return result
}
