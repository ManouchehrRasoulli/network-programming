package tftp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

/*
 * ## Data Packet
 * 		2byte      2byte        n byte
 *    --------------------------------------------------------
 *   | OpCode   |  Block #   |  	Payload 			     |   The packet format
 *   --------------------------------------------------------
 */

type Data struct {
	BlockNumber uint16
	Payload     io.Reader // this payload will reade to packet payload and forward to client
}

func (d *Data) MarshalBinary() ([]byte, error) {
	b := &bytes.Buffer{}
	b.Grow(DatagramSize)

	d.BlockNumber++

	err := binary.Write(b, binary.BigEndian, OpData)
	if err != nil {
		return nil, err
	}

	err = binary.Write(b, binary.BigEndian, d.BlockNumber)
	if err != nil {
		return nil, err
	}

	_, err = io.CopyN(b, d.Payload, BlockSize)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return b.Bytes(), nil
}

func (d *Data) UnmarshalBinary(b []byte) error {
	if l := len(b); l < 4 || l > DatagramSize {
		return errors.New("invalid DATA")
	}

	buff := bytes.NewBuffer(b)

	var opcode OpCode
	err := binary.Read(buff, binary.BigEndian, &opcode)
	if err != nil {
		return err
	} else if opcode != OpData {
		return errors.New("invalid DATA")
	}

	err = binary.Read(buff, binary.BigEndian, &d.BlockNumber)
	if err != nil {
		return err
	}

	d.Payload = bytes.NewReader(b[10:])
	return nil
}
