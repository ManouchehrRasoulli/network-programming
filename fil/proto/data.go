package proto

import (
	"bytes"
	"encoding/binary"
	"strings"
)

/*
 *	# read request format
 *
 *	  2byte     2byte      nbyte         delim
 *	+-----------------------------------------+
 *	|  seq   |   size   |   payload         |0|
 *	+-----------------------------------------+
 */

type Data struct {
	Seq     uint64
	Size    uint64 // the read will end when size of the data is less-than size of the request buffer size
	Payload string
}

func (d *Data) MarshalBinary() ([]byte, error) {
	buff := &bytes.Buffer{}
	var capacity = 2 + 2 + len(d.Payload) + 1
	buff.Grow(capacity)

	if err := binary.Write(buff, binary.BigEndian, d.Seq); err != nil {
		return nil, err
	}

	if err := binary.Write(buff, binary.BigEndian, d.Size); err != nil {
		return nil, err
	}

	if _, err := buff.WriteString(d.Payload); err != nil {
		return nil, err
	}

	if err := buff.WriteByte(0); err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

func (d *Data) UnmarshalBinary(b []byte) error {
	buffer := bytes.NewBuffer(b)

	if err := binary.Read(buffer, binary.BigEndian, &d.Seq); err != nil {
		return err
	}

	if err := binary.Read(buffer, binary.BigEndian, &d.Size); err != nil {
		return err
	}

	if line, err := buffer.ReadString(0); err != nil {
		return err
	} else {
		d.Payload = strings.TrimRight(line, "\x00")
	}

	return nil
}
