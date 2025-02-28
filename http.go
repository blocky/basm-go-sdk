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
	resultData := bytesFromFatPtr(outPtr)

	value, err := readHostResult[HTTPRequestOutput](resultData)
	if err != nil {
		msg := "reading host fn result: " + err.Error()
		return HTTPRequestOutput{}, errors.New(msg)
	}
	return value, nil
}
