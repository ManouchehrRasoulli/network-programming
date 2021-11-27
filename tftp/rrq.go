package tftp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"strings"
)

/*
 * ## Read request packet format
 * 		2byte      nbyte        1byte        nbyte      1byte
 *    --------------------------------------------------------
 *   | OpCode   |  File name   |  0  |   Mode        |   0   |   The packet format
 *   --------------------------------------------------------
 */

type ReadRequest struct {
	FileName string
	Mode     string
}

func (r ReadRequest) MarshalBinary() ([]byte, error) {
	mode := "octet"
	if r.Mode != "" {
		mode = r.Mode
	}

	// opcode + filename + 0 + mode + 0
	capacity := 2 + 2 + len(r.FileName) + 1 + len(r.Mode) + 1
	b := &bytes.Buffer{}
	b.Grow(capacity)

	err := binary.Write(b, binary.BigEndian, OpRRQ)
	if err != nil {
		return nil, err
	}

	n, err := b.WriteString(r.FileName)
	if err != nil {
		return nil, err
	} else if n != len(r.FileName) {
		return nil, errors.New("can't write file name into binary format")
	}

	err = b.WriteByte(0)
	if err != nil {
		return nil, err
	}

	n, err = b.WriteString(mode)
	if err != nil {
		return nil, err
	} else if n != len(mode) {
		return nil, errors.New("can't write mode into binary format")
	}

	err = b.WriteByte(0)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (r *ReadRequest) UnmarshalBinary(b []byte) error {
	buff := bytes.NewBuffer(b)

	var opCode OpCode
	err := binary.Read(buff, binary.BigEndian, &opCode)
	if err != nil {
		return err
	}

	if opCode != OpRRQ {
		return errors.New("invalid RRQ")
	}

	r.FileName, err = buff.ReadString(0)
	if err != nil {
		return err
	}
	r.FileName = strings.TrimRight(r.FileName, "\x00") // remove zero tail
	if len(r.FileName) == 0 {
		return errors.New("invalid RRQ")
	}

	r.Mode, err = buff.ReadString(0)
	if err != nil {
		return err
	}
	r.Mode = strings.TrimRight(r.Mode, "\x00")
	if len(r.Mode) == 0 {
		return errors.New("invalid RRQ")
	}

	actual := strings.ToLower(r.Mode)
	if actual != "octet" {
		return errors.New("only binary transfer supported")
	}
	
	return nil
}
