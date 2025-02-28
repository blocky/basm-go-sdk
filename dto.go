package basm

import (
	"encoding/json"

	"github.com/mailru/easyjson"
)

func marshal(v easyjson.Marshaler) ([]byte, error) {
	return easyjson.Marshal(v)
}

func unmarshal(data []byte, v easyjson.Unmarshaler) error {
	return easyjson.Unmarshal(data, v)
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

type HTTPRequestResult struct {
	IsOk  bool              `json:"ok"`
	Value HTTPRequestOutput `json:"value"`
	Error string            `json:"error"`
}

type VerifyAttestationInput struct {
	EnclaveAttestedKey    json.RawMessage `json:"enclave_attested_app_public_key"`
	TransitiveAttestation json.RawMessage `json:"transitive_attestation"`
	AcceptableMeasures    json.RawMessage `json:"acceptable_measurements"`
}

type VerifyAttestationResult struct {
	IsOk  bool                    `json:"ok"`
	Value VerifyAttestationOutput `json:"value"`
	Error string                  `json:"error"`
}

type VerifyAttestationOutput struct {
	RawClaims []byte `json:"raw_claims"`
}
