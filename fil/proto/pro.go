package proto

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

/*
 *	# read request format
 *
 *	  2byte     2byte      2byte      nbyte  delim
 *	+-----------------------------------------+
 *	|  opc    |  seq   |   size   |   path  |0|
 *	+-----------------------------------------+
 */

type RReq struct {
	Path string // the path which client which to read
	Seq  uint64 // the sequence number which we are reading
	Size uint64 // the size of rreq data till now
}

func (r *RReq) MarshalBinary() ([]byte, error) {
	var capacity = 2 + 2 + 2 + len(r.Path) + 1
	buff := &bytes.Buffer{}
	buff.Grow(capacity)

	if err := binary.Write(buff, binary.BigEndian, OpRR); err != nil {
		return nil, err
	}

	if err := binary.Write(buff, binary.BigEndian, r.Seq); err != nil {
		return nil, err
	}

	if err := binary.Write(buff, binary.BigEndian, r.Size); err != nil {
		return nil, err
	}

	if _, err := buff.WriteString(r.Path); err != nil {
		return nil, err
	}

	if err := buff.WriteByte(0); err != nil{
		return nil, err
	}

	return buff.Bytes(), nil
}

func (r *RReq) UnmarshalBinary(b []byte) error {
	buffer := bytes.NewBuffer(b)

	var op Opcode
	if err := binary.Read(buffer, binary.BigEndian, &op); err != nil {
		return err
	}

	if op != OpRR {
		return fmt.Errorf("invalid read request")
	}

	if err := binary.Read(buffer, binary.BigEndian, &r.Seq); err != nil {
		return err
	}

	if err := binary.Read(buffer, binary.BigEndian, &r.Size); err != nil {
		return err
	}

	if path, err := buffer.ReadString(0); err != nil {
		return err
	} else {
		r.Path = strings.TrimRight(path, "\x00")
	}

	return nil
}
