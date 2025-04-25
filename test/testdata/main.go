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
const enclaveAttestedAppPublicKey = "eyJQbGF0Zm9ybSI6InBsYWluIiwiUGxBdHRlc3RzIjpbImV5SkVZWFJoSWpvaVpYbEthbVJZU2pKYVZqa3daVmhDYkVscWIybGpSRWt4VG0xemVFbHBkMmxhUjBZd1dWTkpOa2xyU1QwaUxDSk5aV0Z6ZFhKbGJXVnVkQ0k2ZXlKUWJHRjBabTl5YlNJNkluQnNZV2x1SWl3aVEyOWtaU0k2SW5Cc1lXbHVJbjE5IiwiZXlKRVlYUmhJam9pVWpCU01WTkdUVEpOYWxwUFZtMXNNMVJyT0hsaU1qVmFWRzFhYm1WVldsbFhTRzk2WkVoYU1HVkZTVDBpTENKTlpXRnpkWEpsYldWdWRDSTZleUpRYkdGMFptOXliU0k2SW5Cc1lXbHVJaXdpUTI5a1pTSTZJbkJzWVdsdUluMTkiLCJleUpFWVhSaElqb2lWMVV3Y21GdFJuUlZNVlpGV1ZaT00xVkZVWGhXU0Vvd1RWVkdlRk5JVm14alNHY3hVa1JuZVU1SVRUMGlMQ0pOWldGemRYSmxiV1Z1ZENJNmV5SlFiR0YwWm05eWJTSTZJbkJzWVdsdUlpd2lRMjlrWlNJNkluQnNZV2x1SW4xOSIsImV5SkVZWFJoSWpvaVZXcFdTVTFzVGpCT2EyTXhUSHBvYm1WRVFsTlNNbTk0V1c1d2JsVlVNR2xtVVQwOUlpd2lUV1ZoYzNWeVpXMWxiblFpT25zaVVHeGhkR1p2Y20waU9pSndiR0ZwYmlJc0lrTnZaR1VpT2lKd2JHRnBiaUo5ZlE9PSIsImV5SkVZWFJoSWpvaVpYaHpTMG8zY2tjMmJFRXZORUp3ZEZaSVdXMVFOemhzYUd3NU0wMU1SRFJWVjFKMU9HZG9jWFZFVVQwaUxDSk5aV0Z6ZFhKbGJXVnVkQ0k2ZXlKUWJHRjBabTl5YlNJNkluQnNZV2x1SWl3aVEyOWtaU0k2SW5Cc1lXbHVJbjE5Il19"

// To regenerate, run the following command:
// cat $BKY_DELPHI_ROOT/internal/testdata/transitive-attested-fn-call-response.json  | jq '.transitive_attestation'
const transitiveAttestedClaims = "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA6AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAJgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIBjYTNkZTQ5OGIyNjJiMjhjNjkyYjMwNThiNGQxMzBmMDJlNWVkODA3NWYzNzEzMWJhN2RlMTE4ZDdkNjRkNmRkZDljYjA5ODYzYzM4NzY1OGZjOGQ3YjEzMDNiYjE2ZDUyYjE2MDhjNmUxNTgzMTVkYmNkYmEzMzcwZTUwY2RhNgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAXYWRkQW5kTG9nV2l0aG91dFNlY3JldHMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgDA1NzllYmNiMzQxYTAwY2M2MDYwNjdiNmExOWUzNDdiNWQ3MGMxZGM1ZTE1ZTUyZjJhNzk3OGMzZDQ4ZTJiZjg2NDdjYmVmNWMwZTRkOTRiMzQ4M2EwZTBiN2RmMjgyNjU4ZmFmNTAxNzk0NzJhN2FkNjY1MDBiYWQxNjllMTYzAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABp7IlJlc3VsdCI6MTEsIkVycm9yIjpudWxsfQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAYTY5ZjczY2NhMjNhOWFjNWM4YjU2N2RjMTg1YTc1NmU5N2M5ODIxNjRmZTI1ODU5ZTBkMWRjYzE0NzVjODBhNjE1YjIxMjNhZjFmNWY5NGMxMWUzZTk0MDJjM2FjNTU4ZjUwMDE5OWQ5NWI2ZDNlMzAxNzU4NTg2MjgxZGNkMjYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQca6plWfFSeUmRFtFpIiqeBno69+W9b+daboBe1FsWcPVVtyBCLOmaPhmwXHabjbZdzG5Dl7Mv1ezRQ6JY2K/0YAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=="

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
