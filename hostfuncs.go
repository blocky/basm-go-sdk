package basm

import (
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

//go:wasmimport env consoleLog
func _hostFuncConsoleLog(ptr, size uint32)

func hostFuncConsoleLog(msg string) {
	msgData := []byte(msg)
	inOffset, inLen := bytesToOffsetSize(msgData)
	_hostFuncConsoleLog(inOffset, inLen)
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
	var zeroReturn httpRequestOutput
	inputData, err := marshal(input)
	if err != nil {
		return zeroReturn, errWrap("marshaling input data: ", err)
	}

	inOffset, inSize := bytesToOffsetSize(inputData)
	outPtr := _hostFuncHTTPRequest(inOffset, inSize)
	runtime.KeepAlive(inputData)
	resultData := bytesFromFatPtr(outPtr)

	var output httpRequestOutput
	err = readHostResult(resultData, &output)
	if err != nil {
		return zeroReturn, errWrap("reading host fn result: ", err)
	}
	return output, nil
}

//go:wasmimport env verifyAttestation
func _hostFuncVerifyAttestation(offset, size uint32) uint64

func hostFuncVerifyAttestation(
	input verifyAttestationInput,
) (verifyAttestationOutput, error) {
	var zeroReturn verifyAttestationOutput
	inputData, err := marshal(input)
	if err != nil {
		return zeroReturn, errWrap("marshaling input data: ", err)
	}

	inOffset, inSize := bytesToOffsetSize(inputData)
	resultPtr := _hostFuncVerifyAttestation(inOffset, inSize)
	runtime.KeepAlive(inputData)
	resultData := bytesFromFatPtr(resultPtr)

	var output verifyAttestationOutput
	err = readHostResult(resultData, &output)
	if err != nil {
		return zeroReturn, errWrap("reading host fn result: ", err)
	}
	return output, nil
}
