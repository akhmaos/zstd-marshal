package zstd_marshal

import (
	"encoding/json"
	"github.com/klauspost/compress/zstd"
)

type ZstdEncoder struct {
	*zstd.Encoder
}

// MarshalWithConcurrency - convert struct to zstd encoded bytes with concurrency
func MarshalWithConcurrency(v interface{}, concurrency int) ([]byte, error) {

	// init zstd encoder with set concurrency
	encoder := newZstdEncoderWithConcurrency(concurrency)
	defer encoder.Close()

	dataBytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	encodedData := encoder.EncodeAll(dataBytes, make([]byte, 0, len(dataBytes)))

	return encodedData, nil
}

// Marshal - convert struct to zstd encoded bytes
func Marshal(v interface{}) ([]byte, error) {

	// init zstd encoder with set concurrency
	encoder := newZstdEncoder()
	defer encoder.Close()

	dataBytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	encodedData := encoder.EncodeAll(dataBytes, make([]byte, 0, len(dataBytes)))

	return encodedData, nil
}
