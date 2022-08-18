package example

import (
	zstdM "github.com/akhmaos/zstd-marshal"
	"runtime"
)

type SomeStruct struct {
	Field       string `json:"Field"`
	SecondField string `json:"SecondField"`
}

func EncodeAndDecode() {

	someData := SomeStruct{
		"Test",
		"Test2",
	}

	encodedDataBytes, err := zstdM.Marshal(someData)

	if err != nil {
		return
	}

	var stForDecode SomeStruct

	err = zstdM.Unmarshal(encodedDataBytes, &stForDecode)
	if err != nil {
		return
	}
}

func EncodeAndDecodeWithConcurrency() {
	someData := SomeStruct{
		"Test",
		"Test2",
	}

	encodedDataBytes, err := zstdM.MarshalWithConcurrency(someData, runtime.GOMAXPROCS(-1))

	if err != nil {
		return
	}

	var stForDecode SomeStruct

	err = zstdM.UnmarshalWithConcurrency(encodedDataBytes, &stForDecode, runtime.GOMAXPROCS(-1))
	if err != nil {
		return
	}
}
