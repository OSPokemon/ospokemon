package engine

type Collision byte

const (
	CLSNground Collision = 0x01
	CLSNfluid  Collision = 0x02
	CLSNnone   Collision = 0x04
)
