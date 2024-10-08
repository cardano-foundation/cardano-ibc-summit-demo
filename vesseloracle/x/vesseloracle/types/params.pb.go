// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vesseloracle/vesseloracle/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
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

// Params defines the parameters for the module.
type Params struct {
	// The minimum number of items in a consolidation window needed for performing outlier detection.
	ConsolidationWindowMinItemCount int32 `protobuf:"varint,1,opt,name=consolidation_window_min_item_count,json=consolidationWindowMinItemCount,proto3" json:"consolidation_window_min_item_count,omitempty"`
	// The maximum number of items in a consolidation window chosen for performing outlier detection. Mostly used to prevent event spamming.
	ConsolidationWindowMaxItemCount int32 `protobuf:"varint,2,opt,name=consolidation_window_max_item_count,json=consolidationWindowMaxItemCount,proto3" json:"consolidation_window_max_item_count,omitempty"`
	// The width of the time interval over which a consolidation is executed.
	ConsolidationWindowIntervalWidth uint64 `protobuf:"varint,3,opt,name=consolidation_window_interval_width,json=consolidationWindowIntervalWidth,proto3" json:"consolidation_window_interval_width,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_06190347da6e3985, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetConsolidationWindowMinItemCount() int32 {
	if m != nil {
		return m.ConsolidationWindowMinItemCount
	}
	return 0
}

func (m *Params) GetConsolidationWindowMaxItemCount() int32 {
	if m != nil {
		return m.ConsolidationWindowMaxItemCount
	}
	return 0
}

func (m *Params) GetConsolidationWindowIntervalWidth() uint64 {
	if m != nil {
		return m.ConsolidationWindowIntervalWidth
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "vesseloracle.vesseloracle.Params")
}

func init() {
	proto.RegisterFile("vesseloracle/vesseloracle/params.proto", fileDescriptor_06190347da6e3985)
}

var fileDescriptor_06190347da6e3985 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2b, 0x4b, 0x2d, 0x2e,
	0x4e, 0xcd, 0xc9, 0x2f, 0x4a, 0x4c, 0xce, 0x49, 0xd5, 0x47, 0xe1, 0x14, 0x24, 0x16, 0x25, 0xe6,
	0x16, 0xeb, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x49, 0x22, 0x4b, 0xe9, 0x21, 0x73, 0xa4, 0x04,
	0x13, 0x73, 0x33, 0xf3, 0xf2, 0xf5, 0xc1, 0x24, 0x44, 0xb5, 0x94, 0x48, 0x7a, 0x7e, 0x7a, 0x3e,
	0x98, 0xa9, 0x0f, 0x62, 0x41, 0x44, 0x95, 0x16, 0x30, 0x71, 0xb1, 0x05, 0x80, 0x0d, 0x15, 0xf2,
	0xe1, 0x52, 0x4e, 0xce, 0xcf, 0x2b, 0xce, 0xcf, 0xc9, 0x4c, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x8b,
	0x2f, 0xcf, 0xcc, 0x4b, 0xc9, 0x2f, 0x8f, 0xcf, 0xcd, 0xcc, 0x8b, 0xcf, 0x2c, 0x49, 0xcd, 0x8d,
	0x4f, 0xce, 0x2f, 0xcd, 0x2b, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x92, 0x47, 0x51, 0x1a,
	0x0e, 0x56, 0xe9, 0x9b, 0x99, 0xe7, 0x59, 0x92, 0x9a, 0xeb, 0x0c, 0x52, 0x86, 0xdb, 0xb4, 0xc4,
	0x0a, 0x64, 0xd3, 0x98, 0x70, 0x9b, 0x96, 0x58, 0x81, 0x30, 0xcd, 0x17, 0x87, 0x69, 0x99, 0x79,
	0x25, 0xa9, 0x45, 0x65, 0x89, 0x39, 0xf1, 0xe5, 0x99, 0x29, 0x25, 0x19, 0x12, 0xcc, 0x0a, 0x8c,
	0x1a, 0x2c, 0x41, 0x0a, 0x58, 0x4c, 0xf3, 0x84, 0x2a, 0x0c, 0x07, 0xa9, 0xb3, 0xd2, 0x7e, 0xb1,
	0x40, 0x9e, 0xb1, 0xeb, 0xf9, 0x06, 0x2d, 0x25, 0x94, 0xd0, 0xad, 0x40, 0x0d, 0x6c, 0x48, 0xb8,
	0x38, 0x59, 0x9f, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13,
	0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x22, 0x3e, 0xdd,
	0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0xe0, 0x60, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xaa, 0x47, 0x50, 0x6c, 0xd4, 0x01, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ConsolidationWindowMinItemCount != that1.ConsolidationWindowMinItemCount {
		return false
	}
	if this.ConsolidationWindowMaxItemCount != that1.ConsolidationWindowMaxItemCount {
		return false
	}
	if this.ConsolidationWindowIntervalWidth != that1.ConsolidationWindowIntervalWidth {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ConsolidationWindowIntervalWidth != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ConsolidationWindowIntervalWidth))
		i--
		dAtA[i] = 0x18
	}
	if m.ConsolidationWindowMaxItemCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ConsolidationWindowMaxItemCount))
		i--
		dAtA[i] = 0x10
	}
	if m.ConsolidationWindowMinItemCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ConsolidationWindowMinItemCount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ConsolidationWindowMinItemCount != 0 {
		n += 1 + sovParams(uint64(m.ConsolidationWindowMinItemCount))
	}
	if m.ConsolidationWindowMaxItemCount != 0 {
		n += 1 + sovParams(uint64(m.ConsolidationWindowMaxItemCount))
	}
	if m.ConsolidationWindowIntervalWidth != 0 {
		n += 1 + sovParams(uint64(m.ConsolidationWindowIntervalWidth))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsolidationWindowMinItemCount", wireType)
			}
			m.ConsolidationWindowMinItemCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConsolidationWindowMinItemCount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsolidationWindowMaxItemCount", wireType)
			}
			m.ConsolidationWindowMaxItemCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConsolidationWindowMaxItemCount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsolidationWindowIntervalWidth", wireType)
			}
			m.ConsolidationWindowIntervalWidth = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConsolidationWindowIntervalWidth |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
