// Code generated by protoc-gen-microweb. DO NOT EDIT.
// source: proto.proto

package proto

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/golang/protobuf/jsonpb"
)

type webCanaryHandler struct {
	r chi.Router
	h CanaryHandler
}

func (h *webCanaryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func (h *webCanaryHandler) SetCanary(w http.ResponseWriter, r *http.Request) {

	req := &CanaryRequest{}

	resp := &CanaryResponse{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}

	if err := h.h.SetCanary(
		r.Context(),
		req,
		resp,
	); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, resp)
}

func RegisterCanaryWeb(r chi.Router, i CanaryHandler, middlewares ...func(http.Handler) http.Handler) {
	handler := &webCanaryHandler{
		r: r,
		h: i,
	}

	r.MethodFunc("POST", "/api/v0/canary", handler.SetCanary)
}

// CanaryRequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of CanaryRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var CanaryRequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *CanaryRequest) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := CanaryRequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*CanaryRequest)(nil)

// CanaryRequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of CanaryRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var CanaryRequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *CanaryRequest) UnmarshalJSON(b []byte) error {
	return CanaryRequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*CanaryRequest)(nil)

// CanaryResponseJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of CanaryResponse. This struct is safe to replace or modify but
// should not be done so concurrently.
var CanaryResponseJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *CanaryResponse) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := CanaryResponseJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*CanaryResponse)(nil)

// CanaryResponseJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of CanaryResponse. This struct is safe to replace or modify but
// should not be done so concurrently.
var CanaryResponseJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *CanaryResponse) UnmarshalJSON(b []byte) error {
	return CanaryResponseJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*CanaryResponse)(nil)