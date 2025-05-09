// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

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

func easyjson9a1cafe1DecodeGithubComGoParkMailRu20251AdminadminInternalModels(in *jlexer.Lexer, out *UpdateUserReq) {
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
		case "description":
			out.Description = string(in.String())
		case "first_name":
			out.FirstName = string(in.String())
		case "last_name":
			out.LastName = string(in.String())
		case "phone_number":
			out.PhoneNumber = string(in.String())
		case "password":
			out.Password = string(in.String())
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
func easyjson9a1cafe1EncodeGithubComGoParkMailRu20251AdminadminInternalModels(out *jwriter.Writer, in UpdateUserReq) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Description != "" {
		const prefix string = ",\"description\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Description))
	}
	if in.FirstName != "" {
		const prefix string = ",\"first_name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FirstName))
	}
	if in.LastName != "" {
		const prefix string = ",\"last_name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LastName))
	}
	if in.PhoneNumber != "" {
		const prefix string = ",\"phone_number\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.PhoneNumber))
	}
	if in.Password != "" {
		const prefix string = ",\"password\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Password))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UpdateUserReq) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9a1cafe1EncodeGithubComGoParkMailRu20251AdminadminInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UpdateUserReq) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9a1cafe1EncodeGithubComGoParkMailRu20251AdminadminInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UpdateUserReq) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9a1cafe1DecodeGithubComGoParkMailRu20251AdminadminInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UpdateUserReq) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9a1cafe1DecodeGithubComGoParkMailRu20251AdminadminInternalModels(l, v)
}
