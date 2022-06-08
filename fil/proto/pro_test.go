package proto

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRReq_MarshalUnmarshalBinary(t *testing.T) {
	t.Log("test")


	r := RReq{
		Path: "test.txt",
		Seq:  0,
		Size: 512,
	}

	b, err := r.MarshalBinary()

	require.Nil(t, err)

	r2 := RReq{}

	err = r2.UnmarshalBinary(b)
	require.Nil(t, err)


	t.Log(fmt.Sprintf("%v", r2))
	require.Equal(t, r.Seq, r2.Seq)
	require.Equal(t, r.Size, r2.Size)
	require.Equal(t, r.Path, r2.Path)
}

func TestData_MarshalUnmarshalBinary(t *testing.T) {
	d := Data{
		Seq:     1,
		Size:    14,
		Payload: "data test and any things you want",
	}

	b, err := d.MarshalBinary()
	require.Nil(t, err)


	d2 := Data{}
	err = d2.UnmarshalBinary(b)
	require.Nil(t, err)

	t.Log(fmt.Sprintf("%v", d2))
	require.Equal(t, d.Size, d2.Size)
	require.Equal(t, d.Seq, d2.Seq)
	require.Equal(t, d.Payload, d2.Payload)
}
