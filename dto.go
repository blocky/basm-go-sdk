package basm

import (
	"encoding/json"
	"errors"

	"github.com/mailru/easyjson"
)

func marshal(v any) ([]byte, error) {
	easy, ok := v.(easyjson.Marshaler)
	if !ok {
		return nil, errors.New("value does not implement easyjson.Marshaler")
	}
	return easyjson.Marshal(easy)
}

func unmarshal(data []byte, v any) error {
	easy, ok := v.(easyjson.Unmarshaler)
	if !ok {
		return errors.New("value does not implement easyjson.Unmarshaler")
	}
	return easyjson.Unmarshal(data, easy)
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
	EnclaveAttestedKey    []byte `json:"enclave_attested_app_public_key"`
	TransitiveAttestation []byte `json:"transitive_attestation"`
	AcceptableMeasures    []byte `json:"acceptable_measurements"`
}

type verifyAttestationOutput struct {
	RawClaims []byte `json:"raw_claims"`
}

type result struct {
	IsOK  bool            `json:"ok"`
	Error string          `json:"error"`
	Value json.RawMessage `json:"value"`
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
	var value T
	err = unmarshal(res.Value, &value)
	if err != nil {
		return zeroReturn, errors.New(
			"failed to unmarshal result value to expected type: " + err.Error(),
		)
	}
	return value, nil
}
