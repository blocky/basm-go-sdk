package basm

import (
	"errors"
	"runtime"
)

//go:wasmimport env verifyAttestation
func _hostFuncVerifyAttestation(offset, size uint32) uint64

func hostFuncVerifyAttestation(
	input VerifyAttestationInput,
) (VerifyAttestationOutput, error) {
	inputData, err := marshal(input)
	if err != nil {
		msg := "marshaling input data: " + err.Error()
		return VerifyAttestationOutput{}, errors.New(msg)
	}

	inOffset, inSize := bytesToOffsetSize(inputData)
	resultPtr := _hostFuncVerifyAttestation(inOffset, inSize)
	runtime.KeepAlive(inputData)
	resultData := bytesFromFatPtr(resultPtr)

	var result VerifyAttestationResult
	err = unmarshal(resultData, &result)
	switch {
	case err != nil:
		msg := "unmarshaling result data: " + err.Error()
		return VerifyAttestationOutput{}, errors.New(msg)
	case !result.IsOk:
		msg := "host fn returned error: " + result.Error
		return VerifyAttestationOutput{}, errors.New(msg)
	}
	return result.Value, nil
}
