// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: connect/abci/v2/vote_extensions.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type OracleVoteExtension struct {
	// Prices defines a map of id(CurrencyPair) -> price.Bytes() . i.e. 1 ->
	// 0x123.. (bytes). Notice the `id` function is determined by the
	// `CurrencyPairIDStrategy` used in the VoteExtensionHandler.
	Prices map[uint64][]byte `protobuf:"bytes,1,rep,name=prices,proto3" json:"prices,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// SanctionList contains a list of sanction items (address and block height).
	SanctionList []SanctionItem `protobuf:"bytes,2,rep,name=sanction_list,json=sanctionList,proto3" json:"sanction_list"`
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

func (m *OracleVoteExtension) GetSanctionList() []SanctionItem {
	if m != nil {
		return m.SanctionList
	}
	return nil
}

type SanctionItem struct {
	Address     string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	BlockHeight uint64 `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
}

func (m *SanctionItem) Reset()         { *m = SanctionItem{} }
func (m *SanctionItem) String() string { return proto.CompactTextString(m) }
func (*SanctionItem) ProtoMessage()    {}
func (*SanctionItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_185ec0708d9f4b6a, []int{1}
}
func (m *SanctionItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SanctionItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SanctionItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SanctionItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SanctionItem.Merge(m, src)
}
func (m *SanctionItem) XXX_Size() int {
	return m.Size()
}
func (m *SanctionItem) XXX_DiscardUnknown() {
	xxx_messageInfo_SanctionItem.DiscardUnknown(m)
}

var xxx_messageInfo_SanctionItem proto.InternalMessageInfo

func (m *SanctionItem) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *SanctionItem) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*OracleVoteExtension)(nil), "connect.abci.v2.OracleVoteExtension")
	proto.RegisterMapType((map[uint64][]byte)(nil), "connect.abci.v2.OracleVoteExtension.PricesEntry")
	proto.RegisterType((*SanctionItem)(nil), "connect.abci.v2.SanctionItem")
}

func init() {
	proto.RegisterFile("connect/abci/v2/vote_extensions.proto", fileDescriptor_185ec0708d9f4b6a)
}

var fileDescriptor_185ec0708d9f4b6a = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x33, 0x6d, 0x6f, 0x2f, 0x77, 0x9a, 0xcb, 0xbd, 0xc4, 0x2e, 0x42, 0xc1, 0x58, 0x0b,
	0x42, 0x17, 0x3a, 0x23, 0x71, 0xa3, 0x2e, 0x0b, 0x95, 0x8a, 0x82, 0x12, 0xc1, 0x85, 0x9b, 0x92,
	0x4c, 0x0f, 0xe9, 0xd0, 0x74, 0x26, 0x64, 0xa6, 0xc1, 0xbe, 0x85, 0x8f, 0xd5, 0x65, 0x97, 0xae,
	0x44, 0x5a, 0x1f, 0x44, 0x3a, 0x4d, 0xa0, 0xa8, 0xbb, 0x73, 0xce, 0xfc, 0xdf, 0x37, 0x8b, 0x1f,
	0x1f, 0x31, 0x29, 0x04, 0x30, 0x4d, 0xc3, 0x88, 0x71, 0x9a, 0xfb, 0x34, 0x97, 0x1a, 0x86, 0xf0,
	0xac, 0x41, 0x28, 0x2e, 0x85, 0x22, 0x69, 0x26, 0xb5, 0x74, 0xfe, 0x15, 0x31, 0xb2, 0x89, 0x91,
	0xdc, 0x6f, 0x35, 0x63, 0x19, 0x4b, 0xf3, 0x46, 0x37, 0xd3, 0x36, 0xd6, 0xf9, 0x40, 0x78, 0xef,
	0x2e, 0x0b, 0x59, 0x02, 0x8f, 0x52, 0x43, 0xbf, 0xb4, 0x38, 0x03, 0x5c, 0x4f, 0x33, 0xce, 0x40,
	0xb9, 0xa8, 0x5d, 0xed, 0x36, 0xfc, 0x53, 0xf2, 0xc5, 0x47, 0x7e, 0xa0, 0xc8, 0xbd, 0x41, 0xfa,
	0x42, 0x67, 0xf3, 0xa0, 0xe0, 0x9d, 0x01, 0xfe, 0xab, 0x42, 0xc1, 0x34, 0x97, 0x62, 0x98, 0x70,
	0xa5, 0xdd, 0x8a, 0x11, 0xee, 0x7f, 0x13, 0x3e, 0x14, 0xa9, 0x6b, 0x0d, 0xd3, 0x5e, 0x6d, 0xf1,
	0x76, 0x60, 0x05, 0x76, 0x49, 0xde, 0x72, 0xa5, 0x5b, 0x17, 0xb8, 0xb1, 0xf3, 0x81, 0xf3, 0x1f,
	0x57, 0x27, 0x30, 0x77, 0x51, 0x1b, 0x75, 0x6b, 0xc1, 0x66, 0x74, 0x9a, 0xf8, 0x57, 0x1e, 0x26,
	0x33, 0x70, 0x2b, 0x6d, 0xd4, 0xb5, 0x83, 0xed, 0x72, 0x59, 0x39, 0x47, 0x9d, 0x1b, 0x6c, 0xef,
	0xea, 0x1d, 0x17, 0xff, 0x0e, 0x47, 0xa3, 0x0c, 0x94, 0x32, 0xfc, 0x9f, 0xa0, 0x5c, 0x9d, 0x43,
	0x6c, 0x47, 0x89, 0x64, 0x93, 0xe1, 0x18, 0x78, 0x3c, 0xd6, 0x46, 0x55, 0x0b, 0x1a, 0xe6, 0x36,
	0x30, 0xa7, 0xde, 0xd5, 0x62, 0xe5, 0xa1, 0xe5, 0xca, 0x43, 0xef, 0x2b, 0x0f, 0xbd, 0xac, 0x3d,
	0x6b, 0xb9, 0xf6, 0xac, 0xd7, 0xb5, 0x67, 0x3d, 0x1d, 0xc7, 0x5c, 0x8f, 0x67, 0x11, 0x61, 0x72,
	0x4a, 0xd5, 0x84, 0xa7, 0x27, 0x53, 0xc8, 0x69, 0xd9, 0x57, 0xee, 0x17, 0x95, 0x01, 0xd5, 0xf3,
	0x14, 0x54, 0x54, 0x37, 0x15, 0x9c, 0x7d, 0x06, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x6f, 0x36, 0xe0,
	0xd2, 0x01, 0x00, 0x00,
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
			{
				size, err := m.SanctionList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVoteExtensions(dAtA, i, uint64(size))
			}
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

func (m *SanctionItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SanctionItem) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SanctionItem) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockHeight != 0 {
		i = encodeVarintVoteExtensions(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintVoteExtensions(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
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
		for _, e := range m.SanctionList {
			l = e.Size()
			n += 1 + l + sovVoteExtensions(uint64(l))
		}
	}
	return n
}

func (m *SanctionItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovVoteExtensions(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovVoteExtensions(uint64(m.BlockHeight))
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
			m.SanctionList = append(m.SanctionList, SanctionItem{})
			if err := m.SanctionList[len(m.SanctionList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *SanctionItem) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: SanctionItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SanctionItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoteExtensions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
