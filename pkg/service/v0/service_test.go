package svc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	v0proto "github.com/cernbox/ocis-canary/pkg/proto/v0/github.com/ocis-canary"
)

func TestCanary_SetCanaryt(t *testing.T) {
	tests := []struct {
		name                 string
		req                  string
		expectedErrorMessage interface{}
	}{
		{"simple", "simple", nil},
		{"UTF", "मिलन", nil},
		{"special char", `%&# /\`, nil},
		{"empty", "", "missing a name"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Canary{}
			req := &v0proto.CanaryRequest{Name: tt.req}
			var rsp = &v0proto.CanaryResponse{}

			err := s.SetCanary(context.Background(), req, rsp)

			if tt.expectedErrorMessage != nil || err != nil {
				assert.EqualError(t, err, tt.expectedErrorMessage.(string))
			} else {
				assert.Equal(t, "Canary "+tt.req, rsp.Message)
			}
		})
	}
}
