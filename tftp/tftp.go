package tftp

const (
	DatagramSize = 512
	BlockSize    = DatagramSize - 4 // block size specify payload without 4 byte for header
)

type OpCode uint64

const (
	OpRRQ OpCode = iota + 1
	_            // read OpWRQ .. write request
	OpData
	OpAck
	OpErr
)

type ErrCode uint16

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
