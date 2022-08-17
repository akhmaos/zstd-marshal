package zstd_marshal

import (
	"github.com/klauspost/compress/zstd"
	"runtime"
)

// newZstdEncoder - for create encoder instance
func newZstdEncoder(encoderConcurrency int) *ZstdEncoder {
	encoder, err := zstd.NewWriter(nil, zstd.WithEncoderConcurrency(encoderConcurrency))

	if err != nil {
		panic(err)
	}

	return &ZstdEncoder{
		encoder,
	}
}

// newZstdDecoder - for create decoder instance
func newZstdDecoder(encoderConcurrency int) *ZstdDecoder {
	decoder, err := zstd.NewReader(nil, zstd.WithDecoderConcurrency(runtime.GOMAXPROCS(encoderConcurrency)))

	if err != nil {
		panic(err)
	}

	return &ZstdDecoder{
		decoder,
	}
}
