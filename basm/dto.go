package basm

import (
	"encoding/json"
	"errors"

	"github.com/mailru/easyjson"
)

func marshal(v easyjson.Marshaler) ([]byte, error) {
	return easyjson.Marshal(v)
}

func unmarshal(data []byte, v easyjson.Unmarshaler) error {
	return easyjson.Unmarshal(data, v)
}

//easyjson:json
type httpRequestInput struct {
	Method  string              `json:"method"`
	URL     string              `json:"url"`
	Headers map[string][]string `json:"headers"`
	Body    []byte              `json:"body"`
}

func fromExportedHTTPInput(in HTTPRequestInput) httpRequestInput {
	return httpRequestInput(in)
}

//easyjson:json
type httpRequestOutput struct {
	StatusCode int                 `json:"status_code"`
	Body       []byte              `json:"body"`
	Headers    map[string][]string `json:"headers"`
}

func toExportedHTTPOutput(out httpRequestOutput) HTTPRequestOutput {
	return HTTPRequestOutput(out)
}

//easyjson:json
type verifyAttestationInput struct {
	EnclaveAttestedKey    string               `json:"enclave_attested_app_public_key"`
	TransitiveAttestation string               `json:"transitive_attestation"`
	AcceptableMeasures    []EnclaveMeasurement `json:"acceptable_measurements"`
}

func fromExportedVerifyAttestationInput(
	in VerifyAttestationInput,
) verifyAttestationInput {
	return verifyAttestationInput{
		EnclaveAttestedKey:    string(in.EnclaveAttestedKey),
		TransitiveAttestation: string(in.TransitiveAttestedClaims),
		AcceptableMeasures:    in.AcceptableMeasures,
	}
}

//easyjson:json
type verifyAttestationOutput struct {
	RawClaims MarshaledAttestedObject `json:"raw_claims"`
}

func toExportedVerifyAttestationOutput(
	out verifyAttestationOutput,
) VerifyAttestationOutput {
	return VerifyAttestationOutput(out)
}

//easyjson:json
type result struct {
	IsOK  bool            `json:"ok"`
	Error string          `json:"error"`
	Value json.RawMessage `json:"value"`
}

func readHostResult[T easyjson.Unmarshaler](data []byte, v T) error {
	var res result
	err := unmarshal(data, &res)
	if err != nil {
		msg := "failed to unmarshal result: " + err.Error()
		return errors.New(msg)
	}
	if !res.IsOK {
		msg := "host fn returned error: " + res.Error
		return errors.New(msg)
	}
	err = unmarshal(res.Value, v)
	if err != nil {
		return errors.New(
			"failed to unmarshal result value to expected type: " + err.Error(),
		)
	}
	return nil
}
