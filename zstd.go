package zstd_marshal

import (
	"github.com/klauspost/compress/zstd"
	"runtime"
)

// newZstdEncoder - for create encoder instance
func newZstdEncoder() *zstdEncoder {
	encoder, err := zstd.NewWriter(nil)

	if err != nil {
		panic(err)
	}

	return &zstdEncoder{
		encoder,
	}
}

// newZstdEncoderWithConcurrency - for create encoder instance with concurrency
func newZstdEncoderWithConcurrency(encoderConcurrency int) *zstdEncoder {
	encoder, err := zstd.NewWriter(nil, zstd.WithEncoderConcurrency(encoderConcurrency))

	if err != nil {
		panic(err)
	}

	return &zstdEncoder{
		encoder,
	}
}

// newZstdDecoderWithConcurrency - for create decoder instance with concurrency
func newZstdDecoderWithConcurrency(encoderConcurrency int) *zstdDecoder {
	decoder, err := zstd.NewReader(nil, zstd.WithDecoderConcurrency(runtime.GOMAXPROCS(encoderConcurrency)))

	if err != nil {
		panic(err)
	}

	return &zstdDecoder{
		decoder,
	}
}

// newZstdDecoder - for create decoder instance
func newZstdDecoder() *zstdDecoder {
	decoder, err := zstd.NewReader(nil)

	if err != nil {
		panic(err)
	}

	return &zstdDecoder{
		decoder,
	}
}
