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

type httpRequestInput struct {
	Method  string              `json:"method"`
	URL     string              `json:"url"`
	Headers map[string][]string `json:"headers"`
	Body    []byte              `json:"body"`
}

type httpRequestOutput struct {
	StatusCode int                 `json:"status_code"`
	Body       []byte              `json:"body"`
	Headers    map[string][]string `json:"headers"`
}

type verifyAttestationInput struct {
	EnclaveAttestedKey    json.RawMessage `json:"enclave_attested_app_public_key"`
	TransitiveAttestation json.RawMessage `json:"transitive_attestation"`
	AcceptableMeasures    json.RawMessage `json:"acceptable_measurements"`
}

type verifyAttestationOutput struct {
	RawClaims []byte `json:"raw_claims"`
}

type result struct {
	IsOK  bool   `json:"ok"`
	Error string `json:"error"`
	Value any    `json:"value"`
}

func readHostResult[T any](data []byte) (T, error) {
	var zeroReturn T
	var res result
	err := unmarshal(data, &res)
	if err != nil {
		msg := "failed to unmarshal result: " + err.Error()
		return zeroReturn, errors.New(msg)
	}
	if !res.IsOK {
		msg := "host fn returned error: " + res.Error
		return zeroReturn, errors.New(msg)
	}
	value, ok := res.Value.(T)
	if !ok {
		msg := "failed to cast result value to expected type"
		return zeroReturn, errors.New(msg)
	}
	return value, nil
}
