package tftp

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEncodeToBinaryAndDecodeFrom_Data(t *testing.T) {
	message := "test message"
	data := Data{
		BlockNumber: 0,
		Payload:     bytes.NewReader([]byte(message)),
	}

	binary, err := data.MarshalBinary()
	require.NoError(t, err)

	recData := Data{
		BlockNumber: 0,
		Payload:     nil,
	}

	err = recData.UnmarshalBinary(binary)
	require.NoError(t, err)

	require.Equal(t, data.BlockNumber, recData.BlockNumber)

	buff := make([]byte, len(message))
	_, err = recData.Payload.Read(buff)
	require.NoError(t, err)

	require.Equal(t, message, string(buff))
}
