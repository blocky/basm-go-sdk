package basm

import (
	"errors"
	"runtime"
)

//go:wasmimport env httpRequest
func _hostFuncHTTPRequest(offset, size uint32) uint64

func hostFuncHTTPRequest(
	input HTTPRequestInput,
) (
	HTTPRequestOutput,
	error,
) {
	inputData, err := marshal(input)
	if err != nil {
		msg := "marshaling input data: " + err.Error()
		return HTTPRequestOutput{}, errors.New(msg)
	}

	inOffset, inSize := bytesToOffsetSize(inputData)
	outPtr := _hostFuncHTTPRequest(inOffset, inSize)
	runtime.KeepAlive(inputData)
	outputData := bytesFromFatPtr(outPtr)

	var result HTTPRequestResult
	err = unmarshal(outputData, &result)
	switch {
	case err != nil:
		msg := "unmarshaling output data: " + err.Error()
		return HTTPRequestOutput{}, errors.New(msg)
	case !result.IsOk:
		msg := "host fn returned error: " + result.Error
		return HTTPRequestOutput{}, errors.New(msg)
	}
	return result.Value, nil
}
