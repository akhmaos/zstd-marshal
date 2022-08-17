package zstd_marshal

import (
	"encoding/json"
	"github.com/klauspost/compress/zstd"
)

type ZstdDecoder struct {
	*zstd.Decoder
}

// UnmarshalWithConcurrency - convert zstd encoded data to struct with concurrency
func UnmarshalWithConcurrency(data []byte, v interface{}, concurrency int) error {

	// init zstd decoder with set concurrency
	decoder := newZstdDecoderWithConcurrency(concurrency)
	defer decoder.Close()

	decodedData, err := decoder.DecodeAll(data, nil)

	if err != nil {
		return err
	}

	err = json.Unmarshal(decodedData, v)
	if err != nil {
		return err
	}

	return nil
}

// Unmarshal - convert zstd encoded data to struct
func Unmarshal(data []byte, v interface{}) error {

	// init zstd decoder with set concurrency
	decoder := newZstdDecoder()
	defer decoder.Close()

	decodedData, err := decoder.DecodeAll(data, nil)

	if err != nil {
		return err
	}

	err = json.Unmarshal(decodedData, v)
	if err != nil {
		return err
	}

	return nil
}
