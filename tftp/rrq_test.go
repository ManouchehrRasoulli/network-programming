package tftp

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEncodeToBinaryAndDecodeFrom_ReadRequest(t *testing.T) {
	req := ReadRequest{
		FileName: "test.txt",
		Mode:     "octet",
	}

	b, err := req.MarshalBinary()
	require.NoError(t, err)
	
	decReq := ReadRequest{
		FileName: "",
		Mode:     "",
	}


	err = decReq.UnmarshalBinary(b)
	require.NoError(t, err)

	require.Equal(t, req.FileName, decReq.FileName)
	require.Equal(t, req.Mode, decReq.Mode)
}
