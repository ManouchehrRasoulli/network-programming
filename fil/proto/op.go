package proto

import (
	"bytes"
	"encoding/binary"
)

const (
	DatagramSize = 512
)

type Opcode uint64

const (
	OpLs Opcode = iota
	OpRR
	_ // write request is reserver
	OpAck
	OpErr
)

type ErrCode uint32

const (
	ErrUnknown ErrCode = iota + 1
	ErrNotFound
	ErrAccessViolation
	ErrDiskFull
	ErrIllegalOp
	ErrUnknownID
	ErrFileExists
	ErrNoUser
)

func ReadOp(b []byte) (Opcode, error) {
	r := bytes.NewReader(b)
	var op Opcode
	if err := binary.Read(r, binary.BigEndian, &op); err != nil {
		return OpErr, err
	}

	return op, nil
}