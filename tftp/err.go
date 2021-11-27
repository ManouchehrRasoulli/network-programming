package tftp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"strings"
)

/*
 * ## Error Packet
 * 		2byte      2byte              n byte				1byte
 *    --------------------------------------------------------
 *   | OpCode   |  Error Code   |  	Message 			   |0|   The packet format
 *   --------------------------------------------------------
 */

type Err struct {
	Error   ErrCode
	Message string
}

func (e Err) MarshalBinary() ([]byte, error) {
	capacity := 2 + 2 + len(e.Message) + 1
	buff := &bytes.Buffer{}
	buff.Grow(capacity)

	err := binary.Write(buff, binary.BigEndian, OpErr)
	if err != nil {
		return nil, err
	}

	err = binary.Write(buff, binary.BigEndian, e.Error)
	if err != nil {
		return nil, err
	}

	_, err = buff.WriteString(e.Message)
	if err != nil {
		return nil, err
	}

	err = buff.WriteByte(0)
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

func (e *Err) UnmarshalBinary(bin []byte) error {
	buffer := bytes.NewBuffer(bin)

	var opcode OpCode
	err := binary.Read(buffer, binary.BigEndian, &opcode)
	if err != nil {
		return err
	} else if opcode != OpErr {
		return errors.New("invalid ERR")
	}

	err = binary.Read(buffer, binary.BigEndian, &e.Error)
	if err != nil {
		return err
	}

	e.Message, err = buffer.ReadString(0)
	if err != nil {
		return err
	}

	e.Message = strings.TrimRight(e.Message, "\x00")
	return nil
}
