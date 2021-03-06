package chain

// Code generated by github.com/CovenantSQL/HashStablePack DO NOT EDIT.

import (
	hsp "github.com/CovenantSQL/HashStablePack/marshalhash"
)

// MarshalHash marshals for hash
func (z *DemoHeader) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 3
	o = append(o, 0x83, 0x83)
	if oTemp, err := z.DatabaseID.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x83)
	o = hsp.AppendTime(o, z.Timestamp)
	o = append(o, 0x83)
	o = hsp.AppendUint32(o, z.SequenceID)
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *DemoHeader) Msgsize() (s int) {
	s = 1 + 11 + z.DatabaseID.Msgsize() + 10 + hsp.TimeSize + 11 + hsp.Uint32Size
	return
}
