package basm

import (
	"errors"
	"runtime"
)

//go:wasmimport env bufferLog
func _hostFuncBufferLog(ptr, size uint32)

func hostFuncBufferLog(msg string) {
	msgData := []byte(msg)
	inOffset, inLen := bytesToOffsetSize(msgData)
	_hostFuncBufferLog(inOffset, inLen)
	runtime.KeepAlive(msgData)
}

//go:wasmimport env httpRequest
func _hostFuncHTTPRequest(offset, size uint32) uint64

func hostFuncHTTPRequest(
	input httpRequestInput,
) (
	httpRequestOutput,
	error,
) {
	inputData, err := marshal(input)
	if err != nil {
		msg := "marshaling input data: " + err.Error()
		return httpRequestOutput{}, errors.New(msg)
	}

	inOffset, inSize := bytesToOffsetSize(inputData)
	outPtr := _hostFuncHTTPRequest(inOffset, inSize)
	runtime.KeepAlive(inputData)
	resultData := bytesFromFatPtr(outPtr)

	value, err := readHostResult[httpRequestOutput](resultData)
	if err != nil {
		msg := "reading host fn result: " + err.Error()
		return httpRequestOutput{}, errors.New(msg)
	}
	return value, nil
}

//go:wasmimport env verifyAttestation
func _hostFuncVerifyAttestation(offset, size uint32) uint64

func hostFuncVerifyAttestation(
	input verifyAttestationInput,
) (verifyAttestationOutput, error) {
	inputData, err := marshal(input)
	if err != nil {
		msg := "marshaling input data: " + err.Error()
		return verifyAttestationOutput{}, errors.New(msg)
	}

	inOffset, inSize := bytesToOffsetSize(inputData)
	resultPtr := _hostFuncVerifyAttestation(inOffset, inSize)
	runtime.KeepAlive(inputData)
	resultData := bytesFromFatPtr(resultPtr)

	output, err := readHostResult[verifyAttestationOutput](resultData)
	if err != nil {
		msg := "reading host fn result: " + err.Error()
		return verifyAttestationOutput{}, errors.New(msg)
	}
	return output, nil
}
