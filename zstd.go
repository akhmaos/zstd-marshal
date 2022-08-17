package zstd_marshal

import (
	"github.com/klauspost/compress/zstd"
	"runtime"
)

// newZstdEncoder - for create encoder instance
func newZstdEncoder() *ZstdEncoder {
	encoder, err := zstd.NewWriter(nil)

	if err != nil {
		panic(err)
	}

	return &ZstdEncoder{
		encoder,
	}
}

// newZstdEncoderWithConcurrency - for create encoder instance with concurrency
func newZstdEncoderWithConcurrency(encoderConcurrency int) *ZstdEncoder {
	encoder, err := zstd.NewWriter(nil, zstd.WithEncoderConcurrency(encoderConcurrency))

	if err != nil {
		panic(err)
	}

	return &ZstdEncoder{
		encoder,
	}
}

// newZstdDecoderWithConcurrency - for create decoder instance with concurrency
func newZstdDecoderWithConcurrency(encoderConcurrency int) *ZstdDecoder {
	decoder, err := zstd.NewReader(nil, zstd.WithDecoderConcurrency(runtime.GOMAXPROCS(encoderConcurrency)))

	if err != nil {
		panic(err)
	}

	return &ZstdDecoder{
		decoder,
	}
}

// newZstdDecoder - for create decoder instance
func newZstdDecoder() *ZstdDecoder {
	decoder, err := zstd.NewReader(nil)

	if err != nil {
		panic(err)
	}

	return &ZstdDecoder{
		decoder,
	}
}
