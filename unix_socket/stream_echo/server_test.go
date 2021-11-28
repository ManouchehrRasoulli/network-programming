package stream_echo

import (
	"bytes"
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"testing"
)

func TestServer(t *testing.T) {
	dir, err := ioutil.TempDir("", "echo_unix")
	require.NoError(t, err)

	defer func() {
		if err = os.RemoveAll(dir); err != nil {
			t.Error(err)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	socket := filepath.Join(dir, fmt.Sprintf("%d.socket", os.Getpid()))
	rAdd, err := streamEchoServer(ctx, "unix", socket)
	require.NoError(t, err)

	err = os.Chmod(socket, os.ModeSocket|0666)
	require.NoError(t, err)

	// client
	t.Log(rAdd.String())
	conn, err := net.Dial("unix", rAdd.String())
	require.NoError(t, err)
	defer func() {
		_ = conn.Close()
	}()

	msg := []byte("ping")
	for i := 0; i < 3; i++ {
		_, err := conn.Write(msg)
		require.NoError(t, err)
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	require.NoError(t, err)

	expected := bytes.Repeat(msg, 3)
	require.Equal(t, expected, buf[:n])
	t.Log(string(buf[:n]))
}
