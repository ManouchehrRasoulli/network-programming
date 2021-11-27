package tftp

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEncodeToBinaryAndDecodeFrom_Error(t *testing.T) {
	errS := Err{
		Error:   ErrIllegalOp,
		Message: "illegal opcode for request",
	}

	binary, err := errS.MarshalBinary()
	require.NoError(t, err)

	errR := Err{
		Error:   0,
		Message: "",
	}
	err = errR.UnmarshalBinary(binary)
	require.NoError(t, err)

	require.Equal(t, errS.Error, errR.Error)
	require.Equal(t, errS.Message, errR.Message)
}
