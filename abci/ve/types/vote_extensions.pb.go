// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: connect/abci/v2/vote_extensions.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// OracleVoteExtension defines the vote extension structure for oracle prices.
type OracleVoteExtension struct {
	// Prices defines a map of id(CurrencyPair) -> price.Bytes() . i.e. 1 ->
	// 0x123.. (bytes). Notice the `id` function is determined by the
	// `CurrencyPairIDStrategy` used in the VoteExtensionHandler.
	Prices       map[uint64][]byte `protobuf:"bytes,1,rep,name=prices,proto3" json:"prices,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SanctionList []string          `protobuf:"bytes,2,rep,name=sanction_list,json=sanctionList,proto3" json:"sanction_list,omitempty"`
}

func (m *OracleVoteExtension) Reset()         { *m = OracleVoteExtension{} }
func (m *OracleVoteExtension) String() string { return proto.CompactTextString(m) }
func (*OracleVoteExtension) ProtoMessage()    {}
func (*OracleVoteExtension) Descriptor() ([]byte, []int) {
	return fileDescriptor_185ec0708d9f4b6a, []int{0}
}
func (m *OracleVoteExtension) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OracleVoteExtension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OracleVoteExtension.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OracleVoteExtension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OracleVoteExtension.Merge(m, src)
}
func (m *OracleVoteExtension) XXX_Size() int {
	return m.Size()
}
func (m *OracleVoteExtension) XXX_DiscardUnknown() {
	xxx_messageInfo_OracleVoteExtension.DiscardUnknown(m)
}

var xxx_messageInfo_OracleVoteExtension proto.InternalMessageInfo

func (m *OracleVoteExtension) GetPrices() map[uint64][]byte {
	if m != nil {
		return m.Prices
	}
	return nil
}

func (m *OracleVoteExtension) GetSanctionList() []string {
	if m != nil {
		return m.SanctionList
	}
	return nil
}

func init() {
	proto.RegisterType((*OracleVoteExtension)(nil), "connect.abci.v2.OracleVoteExtension")
	proto.RegisterMapType((map[uint64][]byte)(nil), "connect.abci.v2.OracleVoteExtension.PricesEntry")
}

func init() {
	proto.RegisterFile("connect/abci/v2/vote_extensions.proto", fileDescriptor_185ec0708d9f4b6a)
}

var fileDescriptor_185ec0708d9f4b6a = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x31, 0x4b, 0xc3, 0x40,
	0x14, 0xc7, 0x7b, 0x89, 0x16, 0xbc, 0x56, 0x94, 0xe8, 0x10, 0x1c, 0x8e, 0xa0, 0x08, 0x19, 0xf4,
	0x4e, 0xe2, 0xa2, 0x8e, 0x42, 0xc5, 0x41, 0x50, 0x32, 0x38, 0xb8, 0x94, 0xe4, 0x78, 0xe8, 0xd1,
	0xf4, 0x2e, 0xe4, 0x5e, 0x0f, 0xf3, 0x2d, 0xfc, 0x46, 0xae, 0x8e, 0x1d, 0x1d, 0x25, 0xf9, 0x22,
	0xd2, 0x9a, 0x80, 0x88, 0xdb, 0x7b, 0x7f, 0x7e, 0xbf, 0xe5, 0x47, 0x8f, 0xa5, 0xd1, 0x1a, 0x24,
	0x8a, 0x2c, 0x97, 0x4a, 0xb8, 0x44, 0x38, 0x83, 0x30, 0x85, 0x57, 0x04, 0x6d, 0x95, 0xd1, 0x96,
	0x97, 0x95, 0x41, 0x13, 0xec, 0x74, 0x18, 0x5f, 0x61, 0xdc, 0x25, 0x87, 0xef, 0x84, 0xee, 0xdd,
	0x57, 0x99, 0x2c, 0xe0, 0xd1, 0x20, 0x4c, 0x7a, 0x3e, 0xb8, 0xa5, 0xc3, 0xb2, 0x52, 0x12, 0x6c,
	0x48, 0x22, 0x3f, 0x1e, 0x25, 0x67, 0xfc, 0x8f, 0xc9, 0xff, 0xb1, 0xf8, 0xc3, 0x5a, 0x99, 0x68,
	0xac, 0xea, 0xb4, 0xf3, 0x83, 0x23, 0xba, 0x6d, 0x33, 0x2d, 0x51, 0x19, 0x3d, 0x2d, 0x94, 0xc5,
	0xd0, 0x8b, 0xfc, 0x78, 0x2b, 0x1d, 0xf7, 0xe3, 0x9d, 0xb2, 0x78, 0x70, 0x49, 0x47, 0xbf, 0xdc,
	0x60, 0x97, 0xfa, 0x33, 0xa8, 0x43, 0x12, 0x91, 0x78, 0x23, 0x5d, 0x9d, 0xc1, 0x3e, 0xdd, 0x74,
	0x59, 0xb1, 0x80, 0xd0, 0x8b, 0x48, 0x3c, 0x4e, 0x7f, 0x9e, 0x2b, 0xef, 0x82, 0x5c, 0xdf, 0x7c,
	0x34, 0x8c, 0x2c, 0x1b, 0x46, 0xbe, 0x1a, 0x46, 0xde, 0x5a, 0x36, 0x58, 0xb6, 0x6c, 0xf0, 0xd9,
	0xb2, 0xc1, 0xd3, 0xc9, 0xb3, 0xc2, 0x97, 0x45, 0xce, 0xa5, 0x99, 0x0b, 0x3b, 0x53, 0xe5, 0xe9,
	0x1c, 0x9c, 0xe8, 0x3b, 0xb9, 0xa4, 0x4b, 0x05, 0x02, 0xeb, 0x12, 0x6c, 0x3e, 0x5c, 0x17, 0x3a,
	0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x10, 0x94, 0xab, 0x43, 0x4a, 0x01, 0x00, 0x00,
}

func (m *OracleVoteExtension) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OracleVoteExtension) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OracleVoteExtension) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SanctionList) > 0 {
		for iNdEx := len(m.SanctionList) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.SanctionList[iNdEx])
			copy(dAtA[i:], m.SanctionList[iNdEx])
			i = encodeVarintVoteExtensions(dAtA, i, uint64(len(m.SanctionList[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Prices) > 0 {
		for k := range m.Prices {
			v := m.Prices[k]
			baseI := i
			if len(v) > 0 {
				i -= len(v)
				copy(dAtA[i:], v)
				i = encodeVarintVoteExtensions(dAtA, i, uint64(len(v)))
				i--
				dAtA[i] = 0x12
			}
			i = encodeVarintVoteExtensions(dAtA, i, uint64(k))
			i--
			dAtA[i] = 0x8
			i = encodeVarintVoteExtensions(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintVoteExtensions(dAtA []byte, offset int, v uint64) int {
	offset -= sovVoteExtensions(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OracleVoteExtension) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Prices) > 0 {
		for k, v := range m.Prices {
			_ = k
			_ = v
			l = 0
			if len(v) > 0 {
				l = 1 + len(v) + sovVoteExtensions(uint64(len(v)))
			}
			mapEntrySize := 1 + sovVoteExtensions(uint64(k)) + l
			n += mapEntrySize + 1 + sovVoteExtensions(uint64(mapEntrySize))
		}
	}
	if len(m.SanctionList) > 0 {
		for _, s := range m.SanctionList {
			l = len(s)
			n += 1 + l + sovVoteExtensions(uint64(l))
		}
	}
	return n
}

func sovVoteExtensions(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVoteExtensions(x uint64) (n int) {
	return sovVoteExtensions(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OracleVoteExtension) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVoteExtensions
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OracleVoteExtension: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OracleVoteExtension: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoteExtensions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthVoteExtensions
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVoteExtensions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Prices == nil {
				m.Prices = make(map[uint64][]byte)
			}
			var mapkey uint64
			mapvalue := []byte{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowVoteExtensions
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowVoteExtensions
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapbyteLen uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowVoteExtensions
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapbyteLen |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intMapbyteLen := int(mapbyteLen)
					if intMapbyteLen < 0 {
						return ErrInvalidLengthVoteExtensions
					}
					postbytesIndex := iNdEx + intMapbyteLen
					if postbytesIndex < 0 {
						return ErrInvalidLengthVoteExtensions
					}
					if postbytesIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = make([]byte, mapbyteLen)
					copy(mapvalue, dAtA[iNdEx:postbytesIndex])
					iNdEx = postbytesIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipVoteExtensions(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthVoteExtensions
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Prices[mapkey] = mapvalue
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SanctionList", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoteExtensions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVoteExtensions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVoteExtensions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SanctionList = append(m.SanctionList, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVoteExtensions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVoteExtensions
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipVoteExtensions(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVoteExtensions
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowVoteExtensions
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowVoteExtensions
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthVoteExtensions
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVoteExtensions
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVoteExtensions
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVoteExtensions        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVoteExtensions          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVoteExtensions = fmt.Errorf("proto: unexpected end of group")
)
