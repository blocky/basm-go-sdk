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

// HTTPRequest uses the host's HTTP client to make a request to the given URL.
func HTTPRequest(req HTTPRequestInput) (HTTPRequestOutput, error) {
	resp, err := hostFuncHTTPRequest(httpRequestInput{
		Method:  req.Method,
		URL:     req.URL,
		Headers: req.Headers,
		Body:    req.Body,
	})
	if err != nil {
		return HTTPRequestOutput{}, err
	}
	return HTTPRequestOutput{
		StatusCode: resp.StatusCode,
		Body:       resp.Body,
		Headers:    resp.Headers,
	}, nil
}

type HTTPRequestInput struct {
	Method  string              `json:"method"`
	URL     string              `json:"url"`
	Headers map[string][]string `json:"headers"`
	Body    []byte              `json:"body"`
}

type HTTPRequestOutput struct {
	StatusCode int                 `json:"status_code"`
	Body       []byte              `json:"body"`
	Headers    map[string][]string `json:"headers"`
}

// VerifyAttestation uses the host's attestation verification functionality to
// verify a transitive attestation from a Blocky attestation service.
func VerifyAttestation(
	input VerifyAttestationInput,
) (VerifyAttestationOutput, error) {
	resp, err := hostFuncVerifyAttestation(verifyAttestationInput{
		EnclaveAttestedKey:    input.EnclaveAttestedKey,
		TransitiveAttestation: input.TransitiveAttestation,
		AcceptableMeasures:    input.AcceptableMeasures,
	})
	if err != nil {
		return VerifyAttestationOutput{}, err
	}
	return VerifyAttestationOutput{
		RawClaims: resp.RawClaims,
	}, nil
}

type VerifyAttestationInput struct {
	EnclaveAttestedKey    []byte `json:"enclave_attested_app_public_key"`
	TransitiveAttestation []byte `json:"transitive_attestation"`
	AcceptableMeasures    []byte `json:"acceptable_measurements"`
}

type VerifyAttestationOutput struct {
	RawClaims []byte `json:"raw_claims"`
}
