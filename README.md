# Go ztds encoding with marshal


This repository makes it possible to implement data compression without transformations on the side of the application code. Excellent tool for integration with different data transfer protocols

## References
* Zstd realtime compress algorithm -  [zstd](https://github.com/klauspost/compress/tree/master/zstd)

## Installation

```
go get github.com/akhmaos/zstd-marshal
```

## Example
```
import zstdM "zstd-marshal"

type SomeStruct struct {
	Field       string `json:"Field"`
	SecondField string `json:"SecondField"`
}
```

We create some struct and set json tag for each because for parse struct we use json marshaling

### Decode and encode without concurrency
```
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
```

### Decode and encode with concurrency

```
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
```
