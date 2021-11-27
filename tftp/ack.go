package tftp

import (
	"bytes"
	"encoding/binary"
	"errors"
)

/*
 * ## Data Packet
 * 		2byte      2byte
 *    ------------------------
 *   | OpCode   |  Block #   |
 *   ------------------------
 */

type Ack uint16

func (a Ack) MarshalBinary() ([]byte, error) {
	capacity := 2 + 2
	buff := &bytes.Buffer{}
	buff.Grow(capacity)

	err := binary.Write(buff, binary.BigEndian, OpAck)
	if err != nil {
		return nil, err
	}

	err = binary.Write(buff, binary.BigEndian, a)
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

func (a *Ack) UnmarshalBinary(b []byte) error {
	reader := bytes.NewReader(b)

	var opcode OpCode
	err := binary.Read(reader, binary.BigEndian, &opcode)
	if err != nil {
		return err
	} else if opcode != OpAck {
		return errors.New("invalid ACK")
	}

	return binary.Read(reader, binary.BigEndian, a)
}
