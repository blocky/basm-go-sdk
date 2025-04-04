package xbasm

import (
	"errors"

	"github.com/mailru/easyjson"

	"github.com/blocky/basm-go-sdk/basm"
)

type FnCallClaims struct {
	CodeHash    string `json:"hash_of_code"`
	Function    string `json:"function"`
	InputHash   string `json:"hash_of_input"`
	Output      []byte `json:"output"`
	SecretsHash string `json:"hash_of_secrets"`
}

//easyjson:json
type sliceOfSliceOfBytes [][]byte

func ParseFnCallClaims(v basm.MarshaledAttestedObject) (*FnCallClaims, error) {
	var fixedRep sliceOfSliceOfBytes
	expectedCount := 5
	switch err := easyjson.Unmarshal(v, &fixedRep); {
	case err != nil:
		localErr := errors.New("unmarshaling to fixed rep")
		return nil, errors.Join(localErr, err)
	case len(fixedRep) != expectedCount:
		return nil, errors.New("unexpected number of fields")
	}

	c := new(FnCallClaims)
	c.CodeHash = string(fixedRep[0])
	c.Function = string(fixedRep[1])
	c.InputHash = string(fixedRep[2])
	c.Output = fixedRep[3]
	c.SecretsHash = string(fixedRep[4])

	return c, nil
}
