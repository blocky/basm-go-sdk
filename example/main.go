package main

import (
	"encoding/json"
	"strconv"

	"github.com/blocky/basm-go-sdk"
)

type Args struct {
	LogThisValue          string `json:"log_this_value"`
	VerifyThisAttestation struct {
		EnclaveAttestedKey    json.RawMessage `json:"enclave_attested_app_public_key"`
		TransitiveAttestation json.RawMessage `json:"transitive_attestation"`
		AcceptableMeasures    json.RawMessage `json:"acceptable_measurements"`
	} `json:"verify_this_attestation"`
}

type SecretArgs struct {
	BearerToken string `json:"bearer_token"`
}

type Result struct {
	Success bool   `json:"success"`
	Value   Output `json:"value,omitempty"`
	Error   string `json:"error,omitempty"`
}

type Output struct {
	RawClaims []byte `json:"raw_claims"`
}

//export kitchenSink
func kitchenSink(inputFPtr, secInputFPtr uint64) uint64 {
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

	// Use the host logging function
	basm.Log(args.LogThisValue)

	authenticatedRequest := basm.HTTPRequestInput{
		Method: "GET",
		// requests to httpbin.org/bearer returns 200 if the Authorization header
		// includes a Bearer token.
		// https://httpbin.org/#/Auth/get_bearer
		URL: "https://httpbin.org/bearer",
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
		statusStr := strconv.Itoa(resp.StatusCode)
		return writeError("expected status code 200, got '" + statusStr + "'")
	}

	// Use the host attestation verification function
	verifyOutput, err := basm.VerifyAttestation(
		basm.VerifyAttestationInput{
			EnclaveAttestedKey:    args.VerifyThisAttestation.EnclaveAttestedKey,
			TransitiveAttestation: args.VerifyThisAttestation.TransitiveAttestation,
			AcceptableMeasures:    args.VerifyThisAttestation.AcceptableMeasures,
		},
	)
	switch {
	case err != nil:
		return writeError("verifying attestation via host: " + err.Error())
	case len(verifyOutput.RawClaims) == 0:
		return writeError("expected attestation claims, got empty")
	}

	return writeOutput(Output{
		RawClaims: verifyOutput.RawClaims,
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
	return basm.WriteToHost(data)
}

func main() {}
