package handler

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "EAI CUPINXER"
	result := jsonError(msg)
	require.Equal(t, []byte(`{"message":"EAI CUPINXER"}`), result)
}
