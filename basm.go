package basm

// WriteToHost takes a byte slice and writes it to the shared memory of the host's
// WebAssembly runtime. It returns a packed offset/size pair in a format
// compatible with WebAssembly numeric types. The host is expected to free the
// memory when it is done with it.
func WriteToHost(data []byte) uint64 {
	return leakToSharedMem(data)
}

// ReadFromHost takes a packed offset/size pair in a format compatible with
// WebAssembly numeric types and returns the byte slice that was written to the
// shared memory of the host's WebAssembly runtime. The host is expected to free
// the memory when the guest function is complete.
func ReadFromHost(inputPtr uint64) []byte {
	return bytesFromFatPtr(inputPtr)
}

// Log writes a message to the logs returned to the user with each invocation.
func Log(msg string) {
	hostFuncBufferLog(msg)
}

type HTTPRequestInput struct {
	Method  string
	URL     string
	Headers map[string][]string
	Body    []byte
}

type HTTPRequestOutput struct {
	StatusCode int
	Body       []byte
	Headers    map[string][]string
}

// HTTPRequest uses the host's HTTP client to make a request to the given URL.
func HTTPRequest(req HTTPRequestInput) (HTTPRequestOutput, error) {
	resp, err := hostFuncHTTPRequest(
		fromPublicHTTPInput(req),
	)
	return toPublicHTTPOutput(resp), err
}

type VerifyAttestationInput struct {
	EnclaveAttestedKey    []byte
	TransitiveAttestation []byte
	AcceptableMeasures    []byte
}

type VerifyAttestationOutput struct {
	RawClaims []byte
}

// VerifyAttestation uses the host's attestation verification functionality to
// verify a transitive attestation from a Blocky attestation service.
func VerifyAttestation(
	input VerifyAttestationInput,
) (VerifyAttestationOutput, error) {
	output, err := hostFuncVerifyAttestation(
		fromPublicVerifyAttestationInput(input),
	)
	return toPublicVerifyAttestationOutput(output), err
}
