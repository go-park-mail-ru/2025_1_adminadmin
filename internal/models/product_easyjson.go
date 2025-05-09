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

func easyjsonCf3f67efDecodeGithubComGoParkMailRu20251AdminadminInternalModels(in *jlexer.Lexer, out *WorkingMode) {
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
		case "from":
			out.From = int(in.Int())
		case "to":
			out.To = int(in.Int())
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
func easyjsonCf3f67efEncodeGithubComGoParkMailRu20251AdminadminInternalModels(out *jwriter.Writer, in WorkingMode) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"from\":"
		out.RawString(prefix[1:])
		out.Int(int(in.From))
	}
	{
		const prefix string = ",\"to\":"
		out.RawString(prefix)
		out.Int(int(in.To))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v WorkingMode) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCf3f67efEncodeGithubComGoParkMailRu20251AdminadminInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v WorkingMode) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCf3f67efEncodeGithubComGoParkMailRu20251AdminadminInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *WorkingMode) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCf3f67efDecodeGithubComGoParkMailRu20251AdminadminInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *WorkingMode) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCf3f67efDecodeGithubComGoParkMailRu20251AdminadminInternalModels(l, v)
}
func easyjsonCf3f67efDecodeGithubComGoParkMailRu20251AdminadminInternalModels1(in *jlexer.Lexer, out *ReviewUser) {
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
		case "id":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.Id).UnmarshalText(data))
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
func easyjsonCf3f67efEncodeGithubComGoParkMailRu20251AdminadminInternalModels1(out *jwriter.Writer, in ReviewUser) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.RawText((in.Id).MarshalText())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ReviewUser) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCf3f67efEncodeGithubComGoParkMailRu20251AdminadminInternalModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ReviewUser) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCf3f67efEncodeGithubComGoParkMailRu20251AdminadminInternalModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ReviewUser) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCf3f67efDecodeGithubComGoParkMailRu20251AdminadminInternalModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ReviewUser) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCf3f67efDecodeGithubComGoParkMailRu20251AdminadminInternalModels1(l, v)
}
func easyjsonCf3f67efDecodeGithubComGoParkMailRu20251AdminadminInternalModels2(in *jlexer.Lexer, out *ReviewInReq) {
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
		case "review_text":
			out.ReviewText = string(in.String())
		case "rating":
			out.Rating = int(in.Int())
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
func easyjsonCf3f67efEncodeGithubComGoParkMailRu20251AdminadminInternalModels2(out *jwriter.Writer, in ReviewInReq) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ReviewText != "" {
		const prefix string = ",\"review_text\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.ReviewText))
	}
	{
		const prefix string = ",\"rating\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Rating))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ReviewInReq) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCf3f67efEncodeGithubComGoParkMailRu20251AdminadminInternalModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ReviewInReq) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCf3f67efEncodeGithubComGoParkMailRu20251AdminadminInternalModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ReviewInReq) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCf3f67efDecodeGithubComGoParkMailRu20251AdminadminInternalModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ReviewInReq) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCf3f67efDecodeGithubComGoParkMailRu20251AdminadminInternalModels2(l, v)
}
func easyjsonCf3f67efDecodeGithubComGoParkMailRu20251AdminadminInternalModels3(in *jlexer.Lexer, out *Review) {
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
		case "id":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.Id).UnmarshalText(data))
			}
		case "user":
			out.User = string(in.String())
		case "user_pic_path":
			out.UserPic = string(in.String())
		case "review_text":
			out.ReviewText = string(in.String())
		case "rating":
			out.Rating = int(in.Int())
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
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
func easyjsonCf3f67efEncodeGithubComGoParkMailRu20251AdminadminInternalModels3(out *jwriter.Writer, in Review) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.RawText((in.Id).MarshalText())
	}
	{
		const prefix string = ",\"user\":"
		out.RawString(prefix)
		out.String(string(in.User))
	}
	{
		const prefix string = ",\"user_pic_path\":"
		out.RawString(prefix)
		out.String(string(in.UserPic))
	}
	if in.ReviewText != "" {
		const prefix string = ",\"review_text\":"
		out.RawString(prefix)
		out.String(string(in.ReviewText))
	}
	{
		const prefix string = ",\"rating\":"
		out.RawString(prefix)
		out.Int(int(in.Rating))
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Review) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCf3f67efEncodeGithubComGoParkMailRu20251AdminadminInternalModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Review) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCf3f67efEncodeGithubComGoParkMailRu20251AdminadminInternalModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Review) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCf3f67efDecodeGithubComGoParkMailRu20251AdminadminInternalModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Review) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCf3f67efDecodeGithubComGoParkMailRu20251AdminadminInternalModels3(l, v)
}
func easyjsonCf3f67efDecodeGithubComGoParkMailRu20251AdminadminInternalModels4(in *jlexer.Lexer, out *RestaurantFull) {
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
		case "id":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.Id).UnmarshalText(data))
			}
		case "name":
			out.Name = string(in.String())
		case "banner_url":
			out.BannerURL = string(in.String())
		case "address":
			out.Address = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "rating":
			out.Rating = float64(in.Float64())
		case "rating_count":
			out.RatingCount = int(in.Int())
		case "working_mode":
			(out.WorkingMode).UnmarshalEasyJSON(in)
		case "delivery_time":
			(out.DeliveryTime).UnmarshalEasyJSON(in)
		case "tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				in.Delim('[')
				if out.Tags == nil {
					if !in.IsDelim(']') {
						out.Tags = make([]string, 0, 4)
					} else {
						out.Tags = []string{}
					}
				} else {
					out.Tags = (out.Tags)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Tags = append(out.Tags, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "categories":
			if in.IsNull() {
				in.Skip()
				out.Categories = nil
			} else {
				in.Delim('[')
				if out.Categories == nil {
					if !in.IsDelim(']') {
						out.Categories = make([]Category, 0, 1)
					} else {
						out.Categories = []Category{}
					}
				} else {
					out.Categories = (out.Categories)[:0]
				}
				for !in.IsDelim(']') {
					var v2 Category
					(v2).UnmarshalEasyJSON(in)
					out.Categories = append(out.Categories, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "reviews":
			if in.IsNull() {
				in.Skip()
				out.Reviews = nil
			} else {
				in.Delim('[')
				if out.Reviews == nil {
					if !in.IsDelim(']') {
						out.Reviews = make([]Review, 0, 0)
					} else {
						out.Reviews = []Review{}
					}
				} else {
					out.Reviews = (out.Reviews)[:0]
				}
				for !in.IsDelim(']') {
					var v3 Review
					(v3).UnmarshalEasyJSON(in)
					out.Reviews = append(out.Reviews, v3)
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
func easyjsonCf3f67efEncodeGithubComGoParkMailRu20251AdminadminInternalModels4(out *jwriter.Writer, in RestaurantFull) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.RawText((in.Id).MarshalText())
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"banner_url\":"
		out.RawString(prefix)
		out.String(string(in.BannerURL))
	}
	{
		const prefix string = ",\"address\":"
		out.RawString(prefix)
		out.String(string(in.Address))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"rating\":"
		out.RawString(prefix)
		out.Float64(float64(in.Rating))
	}
	{
		const prefix string = ",\"rating_count\":"
		out.RawString(prefix)
		out.Int(int(in.RatingCount))
	}
	{
		const prefix string = ",\"working_mode\":"
		out.RawString(prefix)
		(in.WorkingMode).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"delivery_time\":"
		out.RawString(prefix)
		(in.DeliveryTime).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"tags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v4, v5 := range in.Tags {
				if v4 > 0 {
					out.RawByte(',')
				}
				out.String(string(v5))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"categories\":"
		out.RawString(prefix)
		if in.Categories == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v6, v7 := range in.Categories {
				if v6 > 0 {
					out.RawByte(',')
				}
				(v7).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"reviews\":"
		out.RawString(prefix)
		if in.Reviews == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Reviews {
				if v8 > 0 {
					out.RawByte(',')
				}
				(v9).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RestaurantFull) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCf3f67efEncodeGithubComGoParkMailRu20251AdminadminInternalModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RestaurantFull) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCf3f67efEncodeGithubComGoParkMailRu20251AdminadminInternalModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RestaurantFull) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCf3f67efDecodeGithubComGoParkMailRu20251AdminadminInternalModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RestaurantFull) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCf3f67efDecodeGithubComGoParkMailRu20251AdminadminInternalModels4(l, v)
}
func easyjsonCf3f67efDecodeGithubComGoParkMailRu20251AdminadminInternalModels5(in *jlexer.Lexer, out *Product) {
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
		case "id":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.Id).UnmarshalText(data))
			}
		case "name":
			out.Name = string(in.String())
		case "price":
			out.Price = float64(in.Float64())
		case "image_url":
			out.ImageURL = string(in.String())
		case "weight":
			out.Weight = int(in.Int())
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
func easyjsonCf3f67efEncodeGithubComGoParkMailRu20251AdminadminInternalModels5(out *jwriter.Writer, in Product) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.RawText((in.Id).MarshalText())
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"price\":"
		out.RawString(prefix)
		out.Float64(float64(in.Price))
	}
	{
		const prefix string = ",\"image_url\":"
		out.RawString(prefix)
		out.String(string(in.ImageURL))
	}
	{
		const prefix string = ",\"weight\":"
		out.RawString(prefix)
		out.Int(int(in.Weight))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Product) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCf3f67efEncodeGithubComGoParkMailRu20251AdminadminInternalModels5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Product) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCf3f67efEncodeGithubComGoParkMailRu20251AdminadminInternalModels5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Product) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCf3f67efDecodeGithubComGoParkMailRu20251AdminadminInternalModels5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Product) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCf3f67efDecodeGithubComGoParkMailRu20251AdminadminInternalModels5(l, v)
}
func easyjsonCf3f67efDecodeGithubComGoParkMailRu20251AdminadminInternalModels6(in *jlexer.Lexer, out *DeliveryTime) {
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
		case "from":
			out.From = int(in.Int())
		case "to":
			out.To = int(in.Int())
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
func easyjsonCf3f67efEncodeGithubComGoParkMailRu20251AdminadminInternalModels6(out *jwriter.Writer, in DeliveryTime) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"from\":"
		out.RawString(prefix[1:])
		out.Int(int(in.From))
	}
	{
		const prefix string = ",\"to\":"
		out.RawString(prefix)
		out.Int(int(in.To))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DeliveryTime) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCf3f67efEncodeGithubComGoParkMailRu20251AdminadminInternalModels6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DeliveryTime) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCf3f67efEncodeGithubComGoParkMailRu20251AdminadminInternalModels6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DeliveryTime) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCf3f67efDecodeGithubComGoParkMailRu20251AdminadminInternalModels6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DeliveryTime) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCf3f67efDecodeGithubComGoParkMailRu20251AdminadminInternalModels6(l, v)
}
func easyjsonCf3f67efDecodeGithubComGoParkMailRu20251AdminadminInternalModels7(in *jlexer.Lexer, out *Category) {
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
		case "name":
			out.Name = string(in.String())
		case "products":
			if in.IsNull() {
				in.Skip()
				out.Products = nil
			} else {
				in.Delim('[')
				if out.Products == nil {
					if !in.IsDelim(']') {
						out.Products = make([]Product, 0, 1)
					} else {
						out.Products = []Product{}
					}
				} else {
					out.Products = (out.Products)[:0]
				}
				for !in.IsDelim(']') {
					var v10 Product
					(v10).UnmarshalEasyJSON(in)
					out.Products = append(out.Products, v10)
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
func easyjsonCf3f67efEncodeGithubComGoParkMailRu20251AdminadminInternalModels7(out *jwriter.Writer, in Category) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"products\":"
		out.RawString(prefix)
		if in.Products == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Products {
				if v11 > 0 {
					out.RawByte(',')
				}
				(v12).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Category) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCf3f67efEncodeGithubComGoParkMailRu20251AdminadminInternalModels7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Category) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCf3f67efEncodeGithubComGoParkMailRu20251AdminadminInternalModels7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Category) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCf3f67efDecodeGithubComGoParkMailRu20251AdminadminInternalModels7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Category) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCf3f67efDecodeGithubComGoParkMailRu20251AdminadminInternalModels7(l, v)
}
