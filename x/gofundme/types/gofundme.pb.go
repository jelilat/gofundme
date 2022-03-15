// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gofundme/gofundme.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type Gofundme struct {
	Id             uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator        string   `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Goal           uint64   `protobuf:"varint,3,opt,name=goal,proto3" json:"goal,omitempty"`
	Start          string   `protobuf:"bytes,4,opt,name=start,proto3" json:"start,omitempty"`
	End            string   `protobuf:"bytes,5,opt,name=end,proto3" json:"end,omitempty"`
	Totaldonations uint64   `protobuf:"varint,6,opt,name=totaldonations,proto3" json:"totaldonations,omitempty"`
	Claim          string   `protobuf:"bytes,7,opt,name=claim,proto3" json:"claim,omitempty"`
	Donation       []string `protobuf:"bytes,8,rep,name=donation,proto3" json:"donation,omitempty"`
	Donor          []string `protobuf:"bytes,9,rep,name=donor,proto3" json:"donor,omitempty"`
}

func (m *Gofundme) Reset()         { *m = Gofundme{} }
func (m *Gofundme) String() string { return proto.CompactTextString(m) }
func (*Gofundme) ProtoMessage()    {}
func (*Gofundme) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0ee447664db6f3a, []int{0}
}
func (m *Gofundme) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Gofundme) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Gofundme.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Gofundme) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gofundme.Merge(m, src)
}
func (m *Gofundme) XXX_Size() int {
	return m.Size()
}
func (m *Gofundme) XXX_DiscardUnknown() {
	xxx_messageInfo_Gofundme.DiscardUnknown(m)
}

var xxx_messageInfo_Gofundme proto.InternalMessageInfo

func (m *Gofundme) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Gofundme) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Gofundme) GetGoal() uint64 {
	if m != nil {
		return m.Goal
	}
	return 0
}

func (m *Gofundme) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *Gofundme) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

func (m *Gofundme) GetTotaldonations() uint64 {
	if m != nil {
		return m.Totaldonations
	}
	return 0
}

func (m *Gofundme) GetClaim() string {
	if m != nil {
		return m.Claim
	}
	return ""
}

func (m *Gofundme) GetDonation() []string {
	if m != nil {
		return m.Donation
	}
	return nil
}

func (m *Gofundme) GetDonor() []string {
	if m != nil {
		return m.Donor
	}
	return nil
}

func init() {
	proto.RegisterType((*Gofundme)(nil), "cosmonaut.gofundme.gofundme.Gofundme")
}

func init() { proto.RegisterFile("gofundme/gofundme.proto", fileDescriptor_b0ee447664db6f3a) }

var fileDescriptor_b0ee447664db6f3a = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xb1, 0x4e, 0x84, 0x40,
	0x14, 0x45, 0x19, 0x60, 0x77, 0x61, 0x8a, 0x8d, 0x99, 0x98, 0xf8, 0xa2, 0xc9, 0x84, 0x58, 0x18,
	0x2a, 0x28, 0xfc, 0x03, 0x1b, 0xad, 0x29, 0xed, 0x66, 0x19, 0x44, 0x12, 0x98, 0xb7, 0x81, 0x47,
	0xa2, 0x7f, 0xe1, 0x67, 0x59, 0x6e, 0x69, 0xb9, 0x81, 0x1f, 0x31, 0x0c, 0x82, 0xc9, 0x76, 0xe7,
	0xcc, 0xbb, 0x73, 0x8b, 0xcb, 0x6f, 0x4a, 0x7c, 0xeb, 0x8d, 0x6e, 0x8a, 0x74, 0x81, 0xe4, 0xd8,
	0x22, 0xa1, 0xb8, 0xcb, 0xb1, 0x6b, 0xd0, 0xa8, 0x9e, 0x92, 0xf5, 0xb2, 0xc0, 0xfd, 0x99, 0xf1,
	0xe0, 0xf9, 0x4f, 0xc4, 0x9e, 0xbb, 0x95, 0x06, 0x16, 0xb1, 0xd8, 0xcf, 0xdc, 0x4a, 0x0b, 0xe0,
	0xbb, 0xbc, 0x2d, 0x14, 0x61, 0x0b, 0x6e, 0xc4, 0xe2, 0x30, 0x5b, 0x54, 0x08, 0xee, 0x97, 0xa8,
	0x6a, 0xf0, 0x6c, 0xd6, 0xb2, 0xb8, 0xe6, 0x9b, 0x8e, 0x54, 0x4b, 0xe0, 0xdb, 0xec, 0x2c, 0xe2,
	0x8a, 0x7b, 0x85, 0xd1, 0xb0, 0xb1, 0x6f, 0x13, 0x8a, 0x07, 0xbe, 0x27, 0x24, 0x55, 0x6b, 0x34,
	0x8a, 0x2a, 0x34, 0x1d, 0x6c, 0x6d, 0xcb, 0xc5, 0xeb, 0xd4, 0x97, 0xd7, 0xaa, 0x6a, 0x60, 0x37,
	0xf7, 0x59, 0x11, 0xb7, 0x3c, 0x58, 0x22, 0x10, 0x44, 0x5e, 0x1c, 0x66, 0xab, 0x4f, 0x3f, 0x34,
	0x1a, 0x6c, 0x21, 0xb4, 0x87, 0x59, 0x9e, 0x5e, 0xbe, 0x07, 0xc9, 0x4e, 0x83, 0x64, 0xe7, 0x41,
	0xb2, 0xaf, 0x51, 0x3a, 0xa7, 0x51, 0x3a, 0x3f, 0xa3, 0x74, 0x5e, 0x93, 0xb2, 0xa2, 0xf7, 0xfe,
	0x90, 0xe4, 0xd8, 0xa4, 0xeb, 0x48, 0xeb, 0x7c, 0xe9, 0xc7, 0x3f, 0xd2, 0xe7, 0xb1, 0xe8, 0x0e,
	0x5b, 0x3b, 0xe8, 0xe3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x81, 0x1e, 0xb6, 0x88, 0x6b, 0x01,
	0x00, 0x00,
}

func (m *Gofundme) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Gofundme) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Gofundme) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Donor) > 0 {
		for iNdEx := len(m.Donor) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Donor[iNdEx])
			copy(dAtA[i:], m.Donor[iNdEx])
			i = encodeVarintGofundme(dAtA, i, uint64(len(m.Donor[iNdEx])))
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.Donation) > 0 {
		for iNdEx := len(m.Donation) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Donation[iNdEx])
			copy(dAtA[i:], m.Donation[iNdEx])
			i = encodeVarintGofundme(dAtA, i, uint64(len(m.Donation[iNdEx])))
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.Claim) > 0 {
		i -= len(m.Claim)
		copy(dAtA[i:], m.Claim)
		i = encodeVarintGofundme(dAtA, i, uint64(len(m.Claim)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Totaldonations != 0 {
		i = encodeVarintGofundme(dAtA, i, uint64(m.Totaldonations))
		i--
		dAtA[i] = 0x30
	}
	if len(m.End) > 0 {
		i -= len(m.End)
		copy(dAtA[i:], m.End)
		i = encodeVarintGofundme(dAtA, i, uint64(len(m.End)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Start) > 0 {
		i -= len(m.Start)
		copy(dAtA[i:], m.Start)
		i = encodeVarintGofundme(dAtA, i, uint64(len(m.Start)))
		i--
		dAtA[i] = 0x22
	}
	if m.Goal != 0 {
		i = encodeVarintGofundme(dAtA, i, uint64(m.Goal))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintGofundme(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintGofundme(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGofundme(dAtA []byte, offset int, v uint64) int {
	offset -= sovGofundme(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Gofundme) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovGofundme(uint64(m.Id))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovGofundme(uint64(l))
	}
	if m.Goal != 0 {
		n += 1 + sovGofundme(uint64(m.Goal))
	}
	l = len(m.Start)
	if l > 0 {
		n += 1 + l + sovGofundme(uint64(l))
	}
	l = len(m.End)
	if l > 0 {
		n += 1 + l + sovGofundme(uint64(l))
	}
	if m.Totaldonations != 0 {
		n += 1 + sovGofundme(uint64(m.Totaldonations))
	}
	l = len(m.Claim)
	if l > 0 {
		n += 1 + l + sovGofundme(uint64(l))
	}
	if len(m.Donation) > 0 {
		for _, s := range m.Donation {
			l = len(s)
			n += 1 + l + sovGofundme(uint64(l))
		}
	}
	if len(m.Donor) > 0 {
		for _, s := range m.Donor {
			l = len(s)
			n += 1 + l + sovGofundme(uint64(l))
		}
	}
	return n
}

func sovGofundme(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGofundme(x uint64) (n int) {
	return sovGofundme(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Gofundme) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGofundme
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
			return fmt.Errorf("proto: Gofundme: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Gofundme: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGofundme
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGofundme
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
				return ErrInvalidLengthGofundme
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGofundme
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Goal", wireType)
			}
			m.Goal = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGofundme
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Goal |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGofundme
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
				return ErrInvalidLengthGofundme
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGofundme
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Start = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGofundme
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
				return ErrInvalidLengthGofundme
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGofundme
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.End = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Totaldonations", wireType)
			}
			m.Totaldonations = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGofundme
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Totaldonations |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claim", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGofundme
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
				return ErrInvalidLengthGofundme
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGofundme
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Claim = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Donation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGofundme
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
				return ErrInvalidLengthGofundme
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGofundme
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Donation = append(m.Donation, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Donor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGofundme
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
				return ErrInvalidLengthGofundme
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGofundme
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Donor = append(m.Donor, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGofundme(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGofundme
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
func skipGofundme(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGofundme
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
					return 0, ErrIntOverflowGofundme
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
					return 0, ErrIntOverflowGofundme
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
				return 0, ErrInvalidLengthGofundme
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGofundme
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGofundme
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGofundme        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGofundme          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGofundme = fmt.Errorf("proto: unexpected end of group")
)