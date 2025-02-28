package basm

func WriteToHost(data []byte) uint64 {
	return leakToSharedMem(data)
}

func ReadFromHost(inputPtr uint64) []byte {
	return bytesFromFatPtr(inputPtr)
}

func Log(msg string) {
	hostFuncBufferLog(msg)
}

func HTTPRequest(req HTTPRequestInput) (HTTPRequestOutput, error) {
	return hostFuncHTTPRequest(req)
}

func VerifyAttestation(input VerifyAttestationInput) (VerifyAttestationOutput, error) {
	return hostFuncVerifyAttestation(input)
}
