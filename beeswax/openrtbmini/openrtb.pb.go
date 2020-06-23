// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: beeswax/openrtbmini/openrtb.proto

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

//
// OpenRTB 3.3.1: The top-level bid request object.
type BidRequest struct {
	//
	// Unique ID of the bid request, provided by the exchange.
	Id *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	//
	// Information known or derived about the human user of the device.
	User                 *BidRequest_User `protobuf:"bytes,6,opt,name=user" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BidRequest) Reset()         { *m = BidRequest{} }
func (m *BidRequest) String() string { return proto.CompactTextString(m) }
func (*BidRequest) ProtoMessage()    {}
func (*BidRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_openrtb_b6152564d7ce59be, []int{0}
}
func (m *BidRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BidRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BidRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BidRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidRequest.Merge(dst, src)
}
func (m *BidRequest) XXX_Size() int {
	return m.Size()
}
func (m *BidRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BidRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BidRequest proto.InternalMessageInfo

func (m *BidRequest) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *BidRequest) GetUser() *BidRequest_User {
	if m != nil {
		return m.User
	}
	return nil
}

//
// OpenRTB 3.3.12: contains information known or derived about the human user of the device.
// Note that the user ID is an exchange artifact (refer to the “device” object for hardware or
// platform derived IDs) and may be subject to rotation policies. However, this user ID must be
// stable long enough to serve reasonably as the basis for frequency capping.
// If device ID is used as a proxy for unique user ID, use the device object.
type BidRequest_User struct {
	Ext                  *UserExtensions `protobuf:"bytes,1000,opt,name=ext" json:"ext,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BidRequest_User) Reset()         { *m = BidRequest_User{} }
func (m *BidRequest_User) String() string { return proto.CompactTextString(m) }
func (*BidRequest_User) ProtoMessage()    {}
func (*BidRequest_User) Descriptor() ([]byte, []int) {
	return fileDescriptor_openrtb_b6152564d7ce59be, []int{0, 0}
}
func (m *BidRequest_User) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BidRequest_User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BidRequest_User.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BidRequest_User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidRequest_User.Merge(dst, src)
}
func (m *BidRequest_User) XXX_Size() int {
	return m.Size()
}
func (m *BidRequest_User) XXX_DiscardUnknown() {
	xxx_messageInfo_BidRequest_User.DiscardUnknown(m)
}

var xxx_messageInfo_BidRequest_User proto.InternalMessageInfo

func (m *BidRequest_User) GetExt() *UserExtensions {
	if m != nil {
		return m.Ext
	}
	return nil
}

func init() {
	proto.RegisterType((*BidRequest)(nil), "openrtbmini.BidRequest")
	proto.RegisterType((*BidRequest_User)(nil), "openrtbmini.BidRequest.User")
}
func (m *BidRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BidRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("id")
	} else {
		dAtA[i] = 0xa
		i++
		i = encodeVarintOpenrtb(dAtA, i, uint64(len(*m.Id)))
		i += copy(dAtA[i:], *m.Id)
	}
	if m.User != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintOpenrtb(dAtA, i, uint64(m.User.Size()))
		n1, err := m.User.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *BidRequest_User) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BidRequest_User) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Ext != nil {
		dAtA[i] = 0xc2
		i++
		dAtA[i] = 0x3e
		i++
		i = encodeVarintOpenrtb(dAtA, i, uint64(m.Ext.Size()))
		n2, err := m.Ext.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintOpenrtb(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *BidRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != nil {
		l = len(*m.Id)
		n += 1 + l + sovOpenrtb(uint64(l))
	}
	if m.User != nil {
		l = m.User.Size()
		n += 1 + l + sovOpenrtb(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BidRequest_User) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ext != nil {
		l = m.Ext.Size()
		n += 2 + l + sovOpenrtb(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovOpenrtb(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozOpenrtb(x uint64) (n int) {
	return sovOpenrtb(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BidRequest) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOpenrtb
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
			return fmt.Errorf("proto: BidRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BidRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpenrtb
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
				return ErrInvalidLengthOpenrtb
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Id = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpenrtb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOpenrtb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.User == nil {
				m.User = &BidRequest_User{}
			}
			if err := m.User.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOpenrtb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOpenrtb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("id")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BidRequest_User) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOpenrtb
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
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1000:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ext", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpenrtb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOpenrtb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Ext == nil {
				m.Ext = &UserExtensions{}
			}
			if err := m.Ext.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOpenrtb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOpenrtb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipOpenrtb(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOpenrtb
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
					return 0, ErrIntOverflowOpenrtb
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
					return 0, ErrIntOverflowOpenrtb
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
				return 0, ErrInvalidLengthOpenrtb
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowOpenrtb
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
				next, err := skipOpenrtb(dAtA[start:])
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
	ErrInvalidLengthOpenrtb = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOpenrtb   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("beeswax/openrtbmini/openrtb.proto", fileDescriptor_openrtb_b6152564d7ce59be)
}

var fileDescriptor_openrtb_b6152564d7ce59be = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0x4a, 0x4d, 0x2d,
	0x2e, 0x4f, 0xac, 0xd0, 0xcf, 0x2f, 0x48, 0xcd, 0x2b, 0x2a, 0x49, 0xca, 0xcd, 0xcc, 0xcb, 0x84,
	0xb1, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xb8, 0x91, 0xa4, 0xa4, 0x94, 0xb1, 0xa9, 0x4f,
	0xad, 0x28, 0x49, 0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0x83, 0xe8, 0x50, 0x6a, 0x63, 0xe4, 0xe2, 0x72,
	0xca, 0x4c, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91,
	0x60, 0x54, 0x60, 0xd2, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11, 0x32, 0xe0, 0x62, 0x29, 0x2d, 0x4e,
	0x2d, 0x92, 0x60, 0x53, 0x60, 0xd4, 0xe0, 0x36, 0x92, 0xd1, 0x43, 0x32, 0x4a, 0x0f, 0xa1, 0x4d,
	0x2f, 0xb4, 0x38, 0xb5, 0x28, 0x08, 0xac, 0x52, 0xca, 0x8c, 0x8b, 0x05, 0xc4, 0x13, 0xd2, 0xe3,
	0x62, 0x4e, 0xad, 0x28, 0x91, 0x78, 0xc1, 0x0e, 0xd6, 0x29, 0x8d, 0xa2, 0x13, 0xa4, 0xc0, 0x15,
	0xe6, 0x90, 0xe2, 0x20, 0x90, 0x42, 0x27, 0xc5, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63,
	0x7c, 0xf0, 0x48, 0x8e, 0x91, 0x4b, 0x38, 0x39, 0x3f, 0x57, 0x0f, 0xea, 0x7e, 0x98, 0x56, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x0d, 0x22, 0xe2, 0x01, 0x01, 0x00, 0x00,
}
