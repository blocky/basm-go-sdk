package main

import (
	"encoding/json"

	"github.com/blocky/basm-go-sdk"
	"github.com/blocky/basm-go-sdk/x/xbasm"
)

// Test data for verifying SDK functionality.

// enclaveAttestedAppPublicKey was created by a local attestation service server via http http://localhost:8080/enclave-attested-application-public-key
var enclaveAttestedAppPublicKey = []byte(`
	"eyJQbGF0Zm9ybSI6InBsYWluIiwiUGxBdHRlc3RzIjpbImV5SkVZWFJoSWpvaVpYbEthbVJZU2pKYVZqa3daVmhDYkVscWIybGpSRWt4VG0xemVFbHBkMmxhUjBZd1dWTkpOa2xyU1QwaUxDSk5aV0Z6ZFhKbGJXVnVkQ0k2ZXlKUWJHRjBabTl5YlNJNkluQnNZV2x1SWl3aVEyOWtaU0k2SW5Cc1lXbHVJbjE5IiwiZXlKRVlYUmhJam9pVXpBMWFXRlhTbEJUTTA1VVdsWnZkbUpXU2pGVWJHeExZbXMxVmxKR1pERk9WWEI0WVhwT2NXTkhNRDBpTENKTlpXRnpkWEpsYldWdWRDSTZleUpRYkdGMFptOXliU0k2SW5Cc1lXbHVJaXdpUTI5a1pTSTZJbkJzWVdsdUluMTkiLCJleUpFWVhSaElqb2lWbFZ3VEV4Nll6VmtWRkp5WlZkMGRFNVhjSGxWUkd4dFpESTVkRTB3V2sxbFZUbHpWbmwwYlZsWVVUMGlMQ0pOWldGemRYSmxiV1Z1ZENJNmV5SlFiR0YwWm05eWJTSTZJbkJzWVdsdUlpd2lRMjlrWlNJNkluQnNZV2x1SW4xOSIsImV5SkVZWFJoSWpvaVVqQmtTbGxyY0VsUk1VNUVaVVZrVVdWRVFrVk9NbWd3VFVVNVdXRjZNR2xtVVQwOUlpd2lUV1ZoYzNWeVpXMWxiblFpT25zaVVHeGhkR1p2Y20waU9pSndiR0ZwYmlJc0lrTnZaR1VpT2lKd2JHRnBiaUo5ZlE9PSIsImV5SkVZWFJoSWpvaVVIRkhlR1prTUd0dWIyOTRaV1ZHZERZcmVEbGlPVzFHVFROWWRtaFJZbXREWnpkbFIzRnBNbXhwT0QwaUxDSk5aV0Z6ZFhKbGJXVnVkQ0k2ZXlKUWJHRjBabTl5YlNJNkluQnNZV2x1SWl3aVEyOWtaU0k2SW5Cc1lXbHVJbjE5Il19"
`)

// transitiveAttestedClaims was created by a local blocky attestation service
// server using the following command, where BASE64_CMD is `base64` on darwin,
// and `base64 -w 0` otherwise. add-and-log-go is a wasm function that adds the
// input values and logs a string. Here we expect a value of 11 to be attested
// in the function output.
//
//	echo '{ "template": { "input": "'$(echo '{"A":3,"B":8}' | $(BASE64_CMD))'", "function": "addAndLog", "code": "'$(cat ./test/live/testdata/add-and-log-go/x.wasm | $(BASE64_CMD))'"}}' \
//	   	| http http://localhost:8080/transitive-attested-fn-call
var transitiveAttestedClaims = []byte(`
	"WyJXeUpPVjA1b1dWZE9hbHBFV21wT1JGcHRXa2RLYlZwcVNtdE9WR3hzV1hwUk1WcHFSVEphYlZKc1drUmpNVTlFVlRGT2VrVjZUbFJyTUU5SFVUVk9WR042VFhwbk5VMUhXWHBaZWxVd1dUSlNhazVFVVRWTmFscHRUMVJhYUU5VVRtMWFiVkV5VFhwVmQwOUhTVEJQVjBWM1RXcHNhVTlYVlhoYVIxbDZUWHBuTlZwcVp6Vk9NbGt6V1dwbmQwMTZVVEJaVjFKc1RXcFJlRnBIU1RGYWFsSnFUV3BuTWs1cWF6QlplbWM5SWl3aVdWZFNhMUZYTld0VVJ6bHVJaXdpVFVSVk0wOVhWbWxaTWtsNlRrUkdhRTFFUW1wWmVsbDNUbXBCTWs0eVNUSlpWRVUxV2xSTk1FNHlTVEZhUkdOM1dYcEdhMWw2Vm14TlZGWnNUbFJLYlUxdFJUTlBWR00wV1hwT2EwNUVhR3hOYlVwdFQwUlpNRTR5VG1sYVYxa3hXWHBDYkU1SFVUVk9SMGw2VGtSbmVsbFVRbXhOUjBreldrZFplVTlFU1RKT1ZHaHRXVmRaTVUxRVJUTlBWRkV6VFcxRk0xbFhVVEpPYWxWM1RVZEthRnBFUlRKUFYxVjRUbXBOUFNJc0ltVjVTbE5hV0U0eFlraFJhVTlxUVhOSmExWjVZMjA1ZVVscWIybGtWelYwV1ZoS2VtRkhSbk5oVnpWdVNVaE9iRmt6U214a1JHOW5TbE5HTTB0Rk1VcFZNVTVLVkd0amNFbHVNRDBpTENKWlZGazFXbXBqZWxreVRtaE5hazVvVDFkR2FrNVhUVFJaYWxVeVRqSlNhazFVWnpGWlZHTXhUbTFWTlU0eVRUVlBSRWw0VG1wU2JWcFVTVEZQUkZVMVdsUkNhMDFYVW1wWmVrVXdUbnBXYWs5RVFtaE9ha1V4V1dwSmVFMXFUbWhhYWtadFRsZFpOVTVIVFhoTlYxVjZXbFJyTUUxRVNtcE5Na1pxVGxSVk5GcHFWWGROUkVVMVQxZFJOVTVYU1RKYVJFNXNUWHBCZUU1NlZUUk9WR2N5VFdwbmVGcEhUbXROYWxrOUlsMD0iLCJ4VVpFb282cHFHVW9zSUIzUnlvSmVQU3QyVUJ0YkJzRE02U3k5U3RnUC90TVRnWkxQSjllOFV6L0xyMlRNSFdzb0J4STNzMSs4VUlsYlpDTWNaMExod0U9Il0="
`)

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
