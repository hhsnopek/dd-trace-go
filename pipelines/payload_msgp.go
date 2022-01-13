package pipelines

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *groupedStats) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Service":
			z.Service, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Service")
				return
			}
		case "Edge":
			z.Edge, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Edge")
				return
			}
		case "Hash":
			z.Hash, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "Hash")
				return
			}
		case "ParentHash":
			z.ParentHash, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "ParentHash")
				return
			}
		case "EdgeLatency":
			z.EdgeLatency, err = dc.ReadBytes(z.EdgeLatency)
			if err != nil {
				err = msgp.WrapError(err, "EdgeLatency")
				return
			}
		case "PathwayLatency":
			z.PathwayLatency, err = dc.ReadBytes(z.PathwayLatency)
			if err != nil {
				err = msgp.WrapError(err, "PathwayLatency")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *groupedStats) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 6
	// write "Service"
	err = en.Append(0x86, 0xa7, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Service)
	if err != nil {
		err = msgp.WrapError(err, "Service")
		return
	}
	// write "Edge"
	err = en.Append(0xa4, 0x45, 0x64, 0x67, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Edge)
	if err != nil {
		err = msgp.WrapError(err, "Edge")
		return
	}
	// write "Hash"
	err = en.Append(0xac, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x48, 0x61, 0x73, 0x68)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Hash)
	if err != nil {
		err = msgp.WrapError(err, "Hash")
		return
	}
	// write "ParentHash"
	err = en.Append(0xaa, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.ParentHash)
	if err != nil {
		err = msgp.WrapError(err, "ParentHash")
		return
	}
	// write "EdgeLatency"
	err = en.Append(0xab, 0x45, 0x64, 0x67, 0x65, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.EdgeLatency)
	if err != nil {
		err = msgp.WrapError(err, "EdgeLatency")
		return
	}
	// write "PathwayLatency"
	err = en.Append(0xaf, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.PathwayLatency)
	if err != nil {
		err = msgp.WrapError(err, "PathwayLatency")
		return
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *groupedStats) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.Service) + 5 + msgp.StringPrefixSize + len(z.Edge) + 13 + msgp.Uint64Size + 11 + msgp.Uint64Size + 12 + msgp.BytesPrefixSize + len(z.EdgeLatency) + 16 + msgp.BytesPrefixSize + len(z.PathwayLatency)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *statsBucket) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Start":
			z.Start, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "Start")
				return
			}
		case "Duration":
			z.Duration, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "Duration")
				return
			}
		case "Stats":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Stats")
				return
			}
			if cap(z.Stats) >= int(zb0002) {
				z.Stats = (z.Stats)[:zb0002]
			} else {
				z.Stats = make([]groupedStats, zb0002)
			}
			for za0001 := range z.Stats {
				err = z.Stats[za0001].DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Stats", za0001)
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *statsBucket) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Start"
	err = en.Append(0x83, 0xa5, 0x53, 0x74, 0x61, 0x72, 0x74)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Start)
	if err != nil {
		err = msgp.WrapError(err, "Start")
		return
	}
	// write "Duration"
	err = en.Append(0xa8, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Duration)
	if err != nil {
		err = msgp.WrapError(err, "Duration")
		return
	}
	// write "Stats"
	err = en.Append(0xa5, 0x53, 0x74, 0x61, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Stats)))
	if err != nil {
		err = msgp.WrapError(err, "Stats")
		return
	}
	for za0001 := range z.Stats {
		err = z.Stats[za0001].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Stats", za0001)
			return
		}
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *statsBucket) Msgsize() (s int) {
	s = 1 + 6 + msgp.Uint64Size + 9 + msgp.Uint64Size + 6 + msgp.ArrayHeaderSize
	for za0001 := range z.Stats {
		s += z.Stats[za0001].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *statsPayload) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Env":
			z.Env, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Env")
				return
			}
		case "Stats":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Stats")
				return
			}
			if cap(z.Stats) >= int(zb0002) {
				z.Stats = (z.Stats)[:zb0002]
			} else {
				z.Stats = make([]statsBucket, zb0002)
			}
			for za0001 := range z.Stats {
				err = z.Stats[za0001].DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Stats", za0001)
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *statsPayload) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Env"
	err = en.Append(0x82, 0xa3, 0x45, 0x6e, 0x76)
	if err != nil {
		return
	}
	err = en.WriteString(z.Env)
	if err != nil {
		err = msgp.WrapError(err, "Env")
		return
	}
	// write "Stats"
	err = en.Append(0xa5, 0x53, 0x74, 0x61, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Stats)))
	if err != nil {
		err = msgp.WrapError(err, "Stats")
		return
	}
	for za0001 := range z.Stats {
		err = z.Stats[za0001].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Stats", za0001)
			return
		}
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *statsPayload) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(z.Env) + 6 + msgp.ArrayHeaderSize
	for za0001 := range z.Stats {
		s += z.Stats[za0001].Msgsize()
	}
	return
}