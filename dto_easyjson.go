// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package basm

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson56de76c1DecodeGithubComBlockyBasmGoSdk(in *jlexer.Lexer, out *verifyAttestationOutput) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "raw_claims":
			if in.IsNull() {
				in.Skip()
				out.RawClaims = nil
			} else {
				out.RawClaims = in.Bytes()
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson56de76c1EncodeGithubComBlockyBasmGoSdk(out *jwriter.Writer, in verifyAttestationOutput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"raw_claims\":"
		out.RawString(prefix[1:])
		out.Base64Bytes(in.RawClaims)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v verifyAttestationOutput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson56de76c1EncodeGithubComBlockyBasmGoSdk(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v verifyAttestationOutput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson56de76c1EncodeGithubComBlockyBasmGoSdk(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *verifyAttestationOutput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson56de76c1DecodeGithubComBlockyBasmGoSdk(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *verifyAttestationOutput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson56de76c1DecodeGithubComBlockyBasmGoSdk(l, v)
}
func easyjson56de76c1DecodeGithubComBlockyBasmGoSdk1(in *jlexer.Lexer, out *verifyAttestationInput) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "enclave_attested_app_public_key":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.EnclaveAttestedKey).UnmarshalJSON(data))
			}
		case "transitive_attestation":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.TransitiveAttestation).UnmarshalJSON(data))
			}
		case "acceptable_measurements":
			if in.IsNull() {
				in.Skip()
				out.AcceptableMeasures = nil
			} else {
				in.Delim('[')
				if out.AcceptableMeasures == nil {
					if !in.IsDelim(']') {
						out.AcceptableMeasures = make([]EnclaveMeasurement, 0, 2)
					} else {
						out.AcceptableMeasures = []EnclaveMeasurement{}
					}
				} else {
					out.AcceptableMeasures = (out.AcceptableMeasures)[:0]
				}
				for !in.IsDelim(']') {
					var v4 EnclaveMeasurement
					easyjson56de76c1DecodeGithubComBlockyBasmGoSdk2(in, &v4)
					out.AcceptableMeasures = append(out.AcceptableMeasures, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson56de76c1EncodeGithubComBlockyBasmGoSdk1(out *jwriter.Writer, in verifyAttestationInput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"enclave_attested_app_public_key\":"
		out.RawString(prefix[1:])
		out.Raw((in.EnclaveAttestedKey).MarshalJSON())
	}
	{
		const prefix string = ",\"transitive_attestation\":"
		out.RawString(prefix)
		out.Raw((in.TransitiveAttestation).MarshalJSON())
	}
	{
		const prefix string = ",\"acceptable_measurements\":"
		out.RawString(prefix)
		if in.AcceptableMeasures == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.AcceptableMeasures {
				if v5 > 0 {
					out.RawByte(',')
				}
				easyjson56de76c1EncodeGithubComBlockyBasmGoSdk2(out, v6)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v verifyAttestationInput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson56de76c1EncodeGithubComBlockyBasmGoSdk1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v verifyAttestationInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson56de76c1EncodeGithubComBlockyBasmGoSdk1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *verifyAttestationInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson56de76c1DecodeGithubComBlockyBasmGoSdk1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *verifyAttestationInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson56de76c1DecodeGithubComBlockyBasmGoSdk1(l, v)
}
func easyjson56de76c1DecodeGithubComBlockyBasmGoSdk2(in *jlexer.Lexer, out *EnclaveMeasurement) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Platform":
			out.Platform = string(in.String())
		case "Code":
			out.Code = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson56de76c1EncodeGithubComBlockyBasmGoSdk2(out *jwriter.Writer, in EnclaveMeasurement) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Platform\":"
		out.RawString(prefix[1:])
		out.String(string(in.Platform))
	}
	{
		const prefix string = ",\"Code\":"
		out.RawString(prefix)
		out.String(string(in.Code))
	}
	out.RawByte('}')
}
func easyjson56de76c1DecodeGithubComBlockyBasmGoSdk3(in *jlexer.Lexer, out *result) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ok":
			out.IsOK = bool(in.Bool())
		case "error":
			out.Error = string(in.String())
		case "value":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Value).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson56de76c1EncodeGithubComBlockyBasmGoSdk3(out *jwriter.Writer, in result) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ok\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.IsOK))
	}
	{
		const prefix string = ",\"error\":"
		out.RawString(prefix)
		out.String(string(in.Error))
	}
	{
		const prefix string = ",\"value\":"
		out.RawString(prefix)
		out.Raw((in.Value).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson56de76c1EncodeGithubComBlockyBasmGoSdk3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson56de76c1EncodeGithubComBlockyBasmGoSdk3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson56de76c1DecodeGithubComBlockyBasmGoSdk3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson56de76c1DecodeGithubComBlockyBasmGoSdk3(l, v)
}
func easyjson56de76c1DecodeGithubComBlockyBasmGoSdk4(in *jlexer.Lexer, out *httpRequestOutput) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "status_code":
			out.StatusCode = int(in.Int())
		case "body":
			if in.IsNull() {
				in.Skip()
				out.Body = nil
			} else {
				out.Body = in.Bytes()
			}
		case "headers":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Headers = make(map[string][]string)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v8 []string
					if in.IsNull() {
						in.Skip()
						v8 = nil
					} else {
						in.Delim('[')
						if v8 == nil {
							if !in.IsDelim(']') {
								v8 = make([]string, 0, 4)
							} else {
								v8 = []string{}
							}
						} else {
							v8 = (v8)[:0]
						}
						for !in.IsDelim(']') {
							var v9 string
							v9 = string(in.String())
							v8 = append(v8, v9)
							in.WantComma()
						}
						in.Delim(']')
					}
					(out.Headers)[key] = v8
					in.WantComma()
				}
				in.Delim('}')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson56de76c1EncodeGithubComBlockyBasmGoSdk4(out *jwriter.Writer, in httpRequestOutput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"status_code\":"
		out.RawString(prefix[1:])
		out.Int(int(in.StatusCode))
	}
	{
		const prefix string = ",\"body\":"
		out.RawString(prefix)
		out.Base64Bytes(in.Body)
	}
	{
		const prefix string = ",\"headers\":"
		out.RawString(prefix)
		if in.Headers == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v12First := true
			for v12Name, v12Value := range in.Headers {
				if v12First {
					v12First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v12Name))
				out.RawByte(':')
				if v12Value == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
					out.RawString("null")
				} else {
					out.RawByte('[')
					for v13, v14 := range v12Value {
						if v13 > 0 {
							out.RawByte(',')
						}
						out.String(string(v14))
					}
					out.RawByte(']')
				}
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v httpRequestOutput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson56de76c1EncodeGithubComBlockyBasmGoSdk4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v httpRequestOutput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson56de76c1EncodeGithubComBlockyBasmGoSdk4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *httpRequestOutput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson56de76c1DecodeGithubComBlockyBasmGoSdk4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *httpRequestOutput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson56de76c1DecodeGithubComBlockyBasmGoSdk4(l, v)
}
func easyjson56de76c1DecodeGithubComBlockyBasmGoSdk5(in *jlexer.Lexer, out *httpRequestInput) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "method":
			out.Method = string(in.String())
		case "url":
			out.URL = string(in.String())
		case "headers":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Headers = make(map[string][]string)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v15 []string
					if in.IsNull() {
						in.Skip()
						v15 = nil
					} else {
						in.Delim('[')
						if v15 == nil {
							if !in.IsDelim(']') {
								v15 = make([]string, 0, 4)
							} else {
								v15 = []string{}
							}
						} else {
							v15 = (v15)[:0]
						}
						for !in.IsDelim(']') {
							var v16 string
							v16 = string(in.String())
							v15 = append(v15, v16)
							in.WantComma()
						}
						in.Delim(']')
					}
					(out.Headers)[key] = v15
					in.WantComma()
				}
				in.Delim('}')
			}
		case "body":
			if in.IsNull() {
				in.Skip()
				out.Body = nil
			} else {
				out.Body = in.Bytes()
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson56de76c1EncodeGithubComBlockyBasmGoSdk5(out *jwriter.Writer, in httpRequestInput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"method\":"
		out.RawString(prefix[1:])
		out.String(string(in.Method))
	}
	{
		const prefix string = ",\"url\":"
		out.RawString(prefix)
		out.String(string(in.URL))
	}
	{
		const prefix string = ",\"headers\":"
		out.RawString(prefix)
		if in.Headers == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v18First := true
			for v18Name, v18Value := range in.Headers {
				if v18First {
					v18First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v18Name))
				out.RawByte(':')
				if v18Value == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
					out.RawString("null")
				} else {
					out.RawByte('[')
					for v19, v20 := range v18Value {
						if v19 > 0 {
							out.RawByte(',')
						}
						out.String(string(v20))
					}
					out.RawByte(']')
				}
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"body\":"
		out.RawString(prefix)
		out.Base64Bytes(in.Body)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v httpRequestInput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson56de76c1EncodeGithubComBlockyBasmGoSdk5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v httpRequestInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson56de76c1EncodeGithubComBlockyBasmGoSdk5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *httpRequestInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson56de76c1DecodeGithubComBlockyBasmGoSdk5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *httpRequestInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson56de76c1DecodeGithubComBlockyBasmGoSdk5(l, v)
}
