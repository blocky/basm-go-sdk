package main

import (
	"encoding/json"

	"github.com/blocky/basm-go-sdk/basm"
	"github.com/blocky/basm-go-sdk/x/xbasm"
)

// Test data for verifying SDK functionality.
// Here, we assume that the $BKY_DELPHI_ROOT environment variable is
// set to the root of the blocky/delphi repo on your machine.

// To regenerate, run the following command:
// cat $BKY_DELPHI_ROOT/internal/testdata/enclave-attested-application-public-key-response.json | jq '.enclave_attestation'
const enclaveAttestedAppPublicKey = "eyJwbGF0Zm9ybSI6InBsYWluIiwicGxhdGZvcm1fYXR0ZXN0YXRpb25zIjpbImV5SmtZWFJoSWpvaVpYbEthbVJZU2pKYVZqa3daVmhDYkVscWIybGpSRWt4VG0xemVFbHBkMmxhUjBZd1dWTkpOa2xyU1QwaUxDSnRaV0Z6ZFhKbGJXVnVkQ0k2ZXlKd2JHRjBabTl5YlNJNkluQnNZV2x1SWl3aVkyOWtaU0k2SW5Cc1lXbHVJbjE5IiwiZXlKa1lYUmhJam9pVWtoR05VMHhUbEZqTVdSUlQwVk9XRkl5VVhkaGFsRXlaRzVqZGxZd2MzbE9hbXhLVWxSV1JsWlhUVDBpTENKdFpXRnpkWEpsYldWdWRDSTZleUp3YkdGMFptOXliU0k2SW5Cc1lXbHVJaXdpWTI5a1pTSTZJbkJzWVdsdUluMTkiLCJleUprWVhSaElqb2lWa2hHYVUxclozaFVSVTVJU3pCMFVsWjZTbkZUZVhSb1RtdG9ORkZXUmpGWFZsWjJZbFJPVjFKcmF6MGlMQ0p0WldGemRYSmxiV1Z1ZENJNmV5SndiR0YwWm05eWJTSTZJbkJzWVdsdUlpd2lZMjlrWlNJNkluQnNZV2x1SW4xOSIsImV5SmtZWFJoSWpvaVRsZHdNV1J0WkRGTmEwNXVUakZhVEUwd2FFOWpSVVpXVmpKT2JHTjZNR2xtVVQwOUlpd2liV1ZoYzNWeVpXMWxiblFpT25zaWNHeGhkR1p2Y20waU9pSndiR0ZwYmlJc0ltTnZaR1VpT2lKd2JHRnBiaUo5ZlE9PSIsImV5SmtZWFJoSWpvaVZFOUZaV2hxSzJsNk5XOXBTbHBhUkU4MmQzRTVNbEJETVc1emR6WkxOM1V3WTNFM1pWaGliekJrTkQwaUxDSnRaV0Z6ZFhKbGJXVnVkQ0k2ZXlKd2JHRjBabTl5YlNJNkluQnNZV2x1SWl3aVkyOWtaU0k2SW5Cc1lXbHVJbjE5Il19"

// To regenerate, run the following command:
// cat $BKY_DELPHI_ROOT/internal/testdata/transitive-attested-fn-call-response.json  | jq '.transitive_attestation'
const transitiveAttestedClaims = "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA6AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAJgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIBjYTNkZTQ5OGIyNjJiMjhjNjkyYjMwNThiNGQxMzBmMDJlNWVkODA3NWYzNzEzMWJhN2RlMTE4ZDdkNjRkNmRkZDljYjA5ODYzYzM4NzY1OGZjOGQ3YjEzMDNiYjE2ZDUyYjE2MDhjNmUxNTgzMTVkYmNkYmEzMzcwZTUwY2RhNgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAXYWRkQW5kTG9nV2l0aG91dFNlY3JldHMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgDA1NzllYmNiMzQxYTAwY2M2MDYwNjdiNmExOWUzNDdiNWQ3MGMxZGM1ZTE1ZTUyZjJhNzk3OGMzZDQ4ZTJiZjg2NDdjYmVmNWMwZTRkOTRiMzQ4M2EwZTBiN2RmMjgyNjU4ZmFmNTAxNzk0NzJhN2FkNjY1MDBiYWQxNjllMTYzAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABp7IlJlc3VsdCI6MTEsIkVycm9yIjpudWxsfQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAYTY5ZjczY2NhMjNhOWFjNWM4YjU2N2RjMTg1YTc1NmU5N2M5ODIxNjRmZTI1ODU5ZTBkMWRjYzE0NzVjODBhNjE1YjIxMjNhZjFmNWY5NGMxMWUzZTk0MDJjM2FjNTU4ZjUwMDE5OWQ5NWI2ZDNlMzAxNzU4NTg2MjgxZGNkMjYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQYRD08ransoYdzoVQBoLLvQQ+Mt6hrsiPCrNRIOFuWKKC00n1pcLTPajzz4lJoiyU76M1l1knW442XkloZpWJcoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=="

type Args struct {
	LogValue       string `json:"log_value"`
	LogToHostValue string `json:"log_to_host_value"`
}

type SecretArgs struct {
	BearerToken string `json:"bearer_token"`
}

type Result struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	Value   Output `json:"value,omitempty"`
}

type Output struct {
	RawClaims []byte              `json:"raw_claims,omitempty"`
	Claims    *xbasm.FnCallClaims `json:"claims,omitempty"`
}

//export exampleFunc
func exampleFunc(inputFPtr, secInputFPtr uint64) uint64 {
	inputData := basm.ReadFromHost(inputFPtr)
	var args Args
	err := json.Unmarshal(inputData, &args)
	if err != nil {
		return writeError("could not unmarshal input args: " + err.Error())
	}

	secInputData := basm.ReadFromHost(secInputFPtr)
	var secretArgs SecretArgs
	err = json.Unmarshal(secInputData, &secretArgs)
	if err != nil {
		return writeError("could not unmarshal secret input args: " + err.Error())
	}

	// Use a value from the input args and use the host buffered logging function
	basm.Log(args.LogValue)

	// Use a value from the input args and use the host console logging function
	basm.LogToHost(args.LogToHostValue)

	authenticatedRequest := basm.HTTPRequestInput{
		Method: "GET",
		// requests to httpbin.org/bearer returns 200 if the Authorization header
		// includes a Bearer token.
		// https://httpbin.org/#/Auth/get_bearer
		// Note this endpoint is a self-hosted version of httpbin.
		URL: "https://test-httpbin.onrender.com/bearer",
		Headers: map[string][]string{
			"Authorization": {
				// Use a value from the secret input
				"Bearer " + secretArgs.BearerToken,
			},
		},
	}

	// Use the host http function
	resp, err := basm.HTTPRequest(authenticatedRequest)
	switch {
	case err != nil:
		return writeError("making http request via host: " + err.Error())
	case resp.StatusCode != 200:
		return writeError("received non-200 status code")
	}

	// Use the host attestation verification function
	verifyOutput, err := basm.VerifyAttestation(
		basm.VerifyAttestationInput{
			EnclaveAttestedKey:       enclaveAttestedAppPublicKey,
			TransitiveAttestedClaims: transitiveAttestedClaims,
			AcceptableMeasures: []basm.EnclaveMeasurement{
				{
					// The enclave and transitive attestations were created by
					// a local attestation service server, not a real TEE.
					Platform: "plain",
					Code:     "plain",
				},
			},
		})
	switch {
	case err != nil:
		return writeError("verifying attestation via host: " + err.Error())
	case len(verifyOutput.RawClaims) == 0:
		return writeError("expected attestation claims, got empty")
	}

	claims, err := xbasm.ParseFnCallClaims(verifyOutput.RawClaims)
	if err != nil {
		return writeError("parsing call claims: " + err.Error())
	}

	return writeOutput(Output{
		RawClaims: verifyOutput.RawClaims,
		Claims:    claims,
	})
}

func writeOutput(output Output) uint64 {
	return writeResult(Result{
		Success: true,
		Value:   output,
	})
}

func writeError(err string) uint64 {
	return writeResult(Result{
		Success: false,
		Error:   err,
	})
}

func writeResult(res Result) uint64 {
	data, err := json.Marshal(res)
	if err != nil {
		panic("failed to marshal result: " + err.Error())
	}
	// persist the result data to the host
	return basm.WriteToHost(data)
}

// Required for the tinygo compiler
func main() {}
