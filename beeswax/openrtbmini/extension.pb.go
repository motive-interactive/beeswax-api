// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: beeswax/openrtbmini/extension.proto

package openrtbmini

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Next available tag: 9
type UserExtensions struct {
	// This field identifies the user of the bid request. It will be used
	// for user segment lookup and frequency cap.
	// For WEB request, its value is bito_id (user.id).
	// For APP request, its value is the first field which has value
	// in the following order:
	// device.ifa -> device.dpidsha1 -> device.dpidmd5
	// And the value is prefixed with the names in Enum.Bidrequest.User.UserIdType
	// (Except bito_id - we still use bito as prefix)
	UserId     *string                           `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserIdType *Enums_BidRequest_User_UserIdType `protobuf:"varint,2,opt,name=user_id_type,json=userIdType,enum=openrtbmini.Enums_BidRequest_User_UserIdType" json:"user_id_type,omitempty"`
	// GDPR-compliant hashed user_id
	UserIdHashed                 *string  `protobuf:"bytes,8,opt,name=user_id_hashed,json=userIdHashed" json:"user_id_hashed,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *UserExtensions) Reset()         { *m = UserExtensions{} }
func (m *UserExtensions) String() string { return proto.CompactTextString(m) }
func (*UserExtensions) ProtoMessage()    {}
func (*UserExtensions) Descriptor() ([]byte, []int) {
	return fileDescriptor_extension_dc653c55f5fb1dfa, []int{0}
}

var extRange_UserExtensions = []proto.ExtensionRange{
	{Start: 1000, End: 536870911},
}

func (*UserExtensions) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_UserExtensions
}
func (m *UserExtensions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserExtensions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserExtensions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *UserExtensions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserExtensions.Merge(dst, src)
}
func (m *UserExtensions) XXX_Size() int {
	return m.Size()
}
func (m *UserExtensions) XXX_DiscardUnknown() {
	xxx_messageInfo_UserExtensions.DiscardUnknown(m)
}

var xxx_messageInfo_UserExtensions proto.InternalMessageInfo

func (m *UserExtensions) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *UserExtensions) GetUserIdType() Enums_BidRequest_User_UserIdType {
	if m != nil && m.UserIdType != nil {
		return *m.UserIdType
	}
	return Enums_BidRequest_User_UNKNOWN
}

func (m *UserExtensions) GetUserIdHashed() string {
	if m != nil && m.UserIdHashed != nil {
		return *m.UserIdHashed
	}
	return ""
}

func init() {
	proto.RegisterType((*UserExtensions)(nil), "openrtbmini.UserExtensions")
}
func (m *UserExtensions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserExtensions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.UserId != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintExtension(dAtA, i, uint64(len(*m.UserId)))
		i += copy(dAtA[i:], *m.UserId)
	}
	if m.UserIdType != nil {
		dAtA[i] = 0x10
		i++
		i = encodeVarintExtension(dAtA, i, uint64(*m.UserIdType))
	}
	if m.UserIdHashed != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintExtension(dAtA, i, uint64(len(*m.UserIdHashed)))
		i += copy(dAtA[i:], *m.UserIdHashed)
	}
	n, err := github_com_gogo_protobuf_proto.EncodeInternalExtension(m, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintExtension(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *UserExtensions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UserId != nil {
		l = len(*m.UserId)
		n += 1 + l + sovExtension(uint64(l))
	}
	if m.UserIdType != nil {
		n += 1 + sovExtension(uint64(*m.UserIdType))
	}
	if m.UserIdHashed != nil {
		l = len(*m.UserIdHashed)
		n += 1 + l + sovExtension(uint64(l))
	}
	n += github_com_gogo_protobuf_proto.SizeOfInternalExtension(m)
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovExtension(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozExtension(x uint64) (n int) {
	return sovExtension(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserExtensions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExtension
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UserExtensions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserExtensions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtension
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExtension
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.UserId = &s
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserIdType", wireType)
			}
			var v Enums_BidRequest_User_UserIdType
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtension
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (Enums_BidRequest_User_UserIdType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.UserIdType = &v
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserIdHashed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtension
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExtension
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.UserIdHashed = &s
			iNdEx = postIndex
		default:
			if (fieldNum >= 1000) && (fieldNum < 536870912) {
				var sizeOfWire int
				for {
					sizeOfWire++
					wire >>= 7
					if wire == 0 {
						break
					}
				}
				iNdEx -= sizeOfWire
				skippy, err := skipExtension(dAtA[iNdEx:])
				if err != nil {
					return err
				}
				if skippy < 0 {
					return ErrInvalidLengthExtension
				}
				if (iNdEx + skippy) > l {
					return io.ErrUnexpectedEOF
				}
				github_com_gogo_protobuf_proto.AppendExtension(m, int32(fieldNum), dAtA[iNdEx:iNdEx+skippy])
				iNdEx += skippy
			} else {
				iNdEx = preIndex
				skippy, err := skipExtension(dAtA[iNdEx:])
				if err != nil {
					return err
				}
				if skippy < 0 {
					return ErrInvalidLengthExtension
				}
				if (iNdEx + skippy) > l {
					return io.ErrUnexpectedEOF
				}
				m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
				iNdEx += skippy
			}
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipExtension(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExtension
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
					return 0, ErrIntOverflowExtension
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowExtension
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthExtension
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowExtension
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipExtension(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthExtension = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExtension   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("beeswax/openrtbmini/extension.proto", fileDescriptor_extension_dc653c55f5fb1dfa)
}

var fileDescriptor_extension_dc653c55f5fb1dfa = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0x4a, 0x4d, 0x2d,
	0x2e, 0x4f, 0xac, 0xd0, 0xcf, 0x2f, 0x48, 0xcd, 0x2b, 0x2a, 0x49, 0xca, 0xcd, 0xcc, 0xcb, 0xd4,
	0x4f, 0xad, 0x28, 0x49, 0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x46, 0x92, 0x94, 0xd2, 0xc0, 0xa6, 0x03, 0xca, 0x8e, 0x4f, 0xce, 0xcf, 0xcd, 0x85, 0x69,
	0x53, 0x5a, 0xcd, 0xc8, 0xc5, 0x17, 0x5a, 0x9c, 0x5a, 0xe4, 0x0a, 0x33, 0xae, 0x58, 0x48, 0x9c,
	0x8b, 0xbd, 0xb4, 0x38, 0xb5, 0x28, 0x3e, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x88,
	0x0d, 0xc4, 0xf5, 0x4c, 0x11, 0xf2, 0xe7, 0xe2, 0x81, 0x4a, 0xc4, 0x97, 0x54, 0x16, 0xa4, 0x4a,
	0x30, 0x29, 0x30, 0x6a, 0xf0, 0x19, 0xe9, 0xea, 0x21, 0x59, 0xa2, 0xe7, 0x9a, 0x57, 0x9a, 0x5b,
	0xac, 0xe7, 0x94, 0x99, 0x12, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0xa2, 0x07, 0x32, 0x1c, 0x4c,
	0x78, 0xa6, 0x84, 0x54, 0x16, 0xa4, 0x06, 0x71, 0x95, 0xc2, 0xd9, 0x42, 0x2a, 0x5c, 0x7c, 0x30,
	0x03, 0x33, 0x12, 0x8b, 0x33, 0x52, 0x53, 0x24, 0x38, 0xc0, 0x16, 0xf2, 0x40, 0xd4, 0x78, 0x80,
	0xc5, 0xb4, 0x38, 0x39, 0x5e, 0xb0, 0x0b, 0x34, 0x34, 0x34, 0x34, 0x30, 0x39, 0x29, 0x9e, 0x78,
	0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x5c, 0xc2, 0xc9, 0xf9, 0xb9,
	0x7a, 0x50, 0x9f, 0xc2, 0x1c, 0x01, 0x08, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x83, 0x1c, 0xbe, 0x2d,
	0x01, 0x00, 0x00,
}
