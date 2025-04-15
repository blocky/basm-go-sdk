package xbasm

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"slices"

	"github.com/blocky/basm-go-sdk/basm"
)

type FnCallClaims struct {
	CodeHash    string `json:"hash_of_code"`
	Function    string `json:"function"`
	InputHash   string `json:"hash_of_input"`
	Output      []byte `json:"output"`
	SecretsHash string `json:"hash_of_secrets"`
}

func ParseFnCallClaims(v basm.MarshaledAttestedObject) (*FnCallClaims, error) {
	expectedCount := 5

	fixedRep, err := decodeSliceOfBytes(v)
	switch {
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

func isNonZero(b []byte) bool {
	isNonZero := func(b byte) bool { return b != 0 }
	return slices.ContainsFunc(b, isNonZero)
}

func decodeUint64(v []byte) (uint64, error) {
	if len(v) != 32 {
		return 0, errors.New("uint64 encoding must contain 32 bytes")
	}

	padding, data := v[:24], v[24:]
	if isNonZero(padding) {
		return 0, fmt.Errorf("padding contains non-zero values")
	}

	return binary.BigEndian.Uint64(data), nil
}

func sliceHeader() []byte {
	return append(make([]byte, 31), 0x20)
}

// decodeBytes decodes a byte slice (in the go sense) from a
// byte slice (in the go sense) that is assumed to be abi encoded
// in the context of a slice of byte slices.
// It is not intended to be used in any other context.
func decodeBytes(abiEncoded []byte) ([]byte, error) {
	// we specify a few names to help understand the layout.
	// note that the '|' is not part of the layout, it is just a visual aid.
	// | head (32 bytes) | tail (padded to a multiple of 32 bytes) |
	//
	// Restricting our view to just the body we have:
	// tail = | data | padding |
	//
	// the value is head is an integer that tells us how many bytes of
	// tail are data
	//
	// note that because the head is 32 bytes and the tail is padded
	// to a multiple of 32 bytes, a valid input must always
	// have a length that is a multiple of 32.
	const headLen = uint64(32)
	abiEncodedLen := uint64(len(abiEncoded))
	switch {
	case abiEncodedLen < headLen:
		return nil, errors.New("not long enough to have a head")
	case abiEncodedLen%32 != 0:
		return nil, fmt.Errorf("invalid length '%d' not 32-byte aligned", abiEncodedLen)
	}

	// unpack the abi encoded data
	head := abiEncoded[:headLen]
	tail := abiEncoded[headLen:]

	// unpack the head
	dataLen, err := decodeUint64(head)
	if err != nil {
		return nil, fmt.Errorf("decoding data length, %w", err)
	}

	// validate the content in the head
	if dataLen > uint64(len(tail)) {
		return nil, fmt.Errorf("length in head is out of range")
	}

	// unpack the tail
	data := tail[:dataLen]
	padding := tail[dataLen:]

	// validate the content in the tail
	switch {
	case len(padding) >= 32:
		return nil, fmt.Errorf("invalid padding length '%d'", len(padding))
	case isNonZero(padding):
		return nil, fmt.Errorf("padding contains non-zero values")
	}

	dst := make([]byte, dataLen)
	copy(dst, data)
	return dst, nil
}

func decodeSliceOfBytes(abiEncoded []byte) ([][]byte, error) {
	// We specify a few names to help understand the layout.
	// Note that the '|' is not part of the layout, it is just a visual aid.
	//
	// Assume that we encoded a slice of k bytes.
	// | head 64 byte | tail (padded to a multiple of 32 bytes) |
	//
	// Restricting our view to just the head we have
	// head = | type (32 bytes) | num elts 32 bytes) |
	//
	// Restricting our view to just the tail we have
	// tail = | offsets (32*k bytes) | elements (each 32-byte aligned) |
	//
	// Restricting our view to just the elements we have
	// elements = | elt1 | elt2 | ... | eltk |
	// where each elt is aligned to 32 bytes.
	//
	// note that because the head is 64 the offsets are 32*k bytes
	// and each element is padded to a multiple of 32 bytes,
	// a valid input must always have a length that is a multiple of 32.

	headLen := uint64(64)
	abiEncodedLen := uint64(len(abiEncoded))
	switch {
	case abiEncodedLen < headLen:
		return nil, errors.New("not long enough to have a head")
	case abiEncodedLen%32 != 0:
		return nil, fmt.Errorf("invalid length '%d' not 32-byte aligned", abiEncodedLen)
	}

	// unpack the abi encoded data
	head := abiEncoded[:headLen]
	tail := abiEncoded[headLen:]
	tailLen := uint64(len(tail))

	// unpack the head
	typeBytes := head[:32]
	eltCountBytes := head[32:]
	eltCount, err := decodeUint64(eltCountBytes)
	if err != nil {
		return nil, fmt.Errorf("decoding element count, %w", err)
	}

	// validate the head data
	offsetsLen := 32 * eltCount
	switch {
	case !bytes.Equal(typeBytes, sliceHeader()):
		return nil, errors.New("not a slice type")
	case offsetsLen > tailLen:
		return nil, fmt.Errorf("tail too short for %d elements", eltCount)
	}

	// unpack the offsets
	offsets := make([]uint64, eltCount+1)
	for i := range eltCount {
		offset, err := decodeUint64(tail[i*32 : (i+1)*32])
		switch {
		case err != nil:
			return nil, fmt.Errorf("decoding offset for index %d, %w", i, err)
		case offset >= tailLen:
			return nil, fmt.Errorf("offset at index %d out of bounds", i)
		}

		offsets[i] = offset
	}
	offsets[eltCount] = tailLen

	// use offsets to read and decode each encoded byte array
	results := make([][]byte, eltCount)
	for i := range eltCount {
		start := offsets[i]
		end := offsets[i+1]
		if start >= end {
			return nil, fmt.Errorf("start %d greater than end %d", start, end)
		}

		results[i], err = decodeBytes(tail[start:end])
		if err != nil {
			return nil, fmt.Errorf("decoding element %d, %w", i, err)
		}
	}

	return results, nil
}
