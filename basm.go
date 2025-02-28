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
