// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: task-api/rest.proto

package task

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

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

// Set der REST Verben
// https://tools.ietf.org/html/rfc7231#section-4
type Link_Method int32

const (
	Link_PUT     Link_Method = 0
	Link_GET     Link_Method = 1
	Link_HEAD    Link_Method = 2
	Link_POST    Link_Method = 3
	Link_PATCH   Link_Method = 4
	Link_DELETE  Link_Method = 5
	Link_TRACE   Link_Method = 6
	Link_OPTIONS Link_Method = 7
)

var Link_Method_name = map[int32]string{
	0: "PUT",
	1: "GET",
	2: "HEAD",
	3: "POST",
	4: "PATCH",
	5: "DELETE",
	6: "TRACE",
	7: "OPTIONS",
}
var Link_Method_value = map[string]int32{
	"PUT":     0,
	"GET":     1,
	"HEAD":    2,
	"POST":    3,
	"PATCH":   4,
	"DELETE":  5,
	"TRACE":   6,
	"OPTIONS": 7,
}

func (x Link_Method) String() string {
	return proto.EnumName(Link_Method_name, int32(x))
}
func (Link_Method) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_rest_461cf020d5392fa0, []int{0, 0}
}

// IANA konforme Links /
// List of official link rels:
// http://www.iana.org/assignments/link-relations/link-relations.xhtml
type Link struct {
	Rel string `protobuf:"bytes,1,opt,name=rel,proto3" json:"rel,omitempty"`
	// HTTP Verb
	Method Link_Method `protobuf:"varint,2,opt,name=method,proto3,enum=task.v1.Link_Method" json:"method,omitempty"`
	// Absolute URI
	Href string `protobuf:"bytes,3,opt,name=href,proto3" json:"href,omitempty"`
	// Der mime type des Links.
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Link) Reset()         { *m = Link{} }
func (m *Link) String() string { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()    {}
func (*Link) Descriptor() ([]byte, []int) {
	return fileDescriptor_rest_461cf020d5392fa0, []int{0}
}
func (m *Link) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Link) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Link.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Link) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Link.Merge(dst, src)
}
func (m *Link) XXX_Size() int {
	return m.Size()
}
func (m *Link) XXX_DiscardUnknown() {
	xxx_messageInfo_Link.DiscardUnknown(m)
}

var xxx_messageInfo_Link proto.InternalMessageInfo

func (m *Link) GetRel() string {
	if m != nil {
		return m.Rel
	}
	return ""
}

func (m *Link) GetMethod() Link_Method {
	if m != nil {
		return m.Method
	}
	return Link_PUT
}

func (m *Link) GetHref() string {
	if m != nil {
		return m.Href
	}
	return ""
}

func (m *Link) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*Link)(nil), "task.v1.Link")
	proto.RegisterEnum("task.v1.Link_Method", Link_Method_name, Link_Method_value)
}
func (m *Link) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Link) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Rel) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRest(dAtA, i, uint64(len(m.Rel)))
		i += copy(dAtA[i:], m.Rel)
	}
	if m.Method != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRest(dAtA, i, uint64(m.Method))
	}
	if len(m.Href) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRest(dAtA, i, uint64(len(m.Href)))
		i += copy(dAtA[i:], m.Href)
	}
	if len(m.Type) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintRest(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintRest(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Link) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Rel)
	if l > 0 {
		n += 1 + l + sovRest(uint64(l))
	}
	if m.Method != 0 {
		n += 1 + sovRest(uint64(m.Method))
	}
	l = len(m.Href)
	if l > 0 {
		n += 1 + l + sovRest(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovRest(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRest(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRest(x uint64) (n int) {
	return sovRest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Link) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRest
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
			return fmt.Errorf("proto: Link: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Link: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRest
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
				return ErrInvalidLengthRest
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			m.Method = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Method |= (Link_Method(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Href", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRest
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
				return ErrInvalidLengthRest
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Href = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRest
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
				return ErrInvalidLengthRest
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRest
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
func skipRest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRest
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
					return 0, ErrIntOverflowRest
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
					return 0, ErrIntOverflowRest
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
				return 0, ErrInvalidLengthRest
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRest
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
				next, err := skipRest(dAtA[start:])
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
	ErrInvalidLengthRest = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRest   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("task-api/rest.proto", fileDescriptor_rest_461cf020d5392fa0) }

var fileDescriptor_rest_461cf020d5392fa0 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0x8f, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x40, 0xdd, 0x66, 0xbb, 0xb1, 0x23, 0xc8, 0x30, 0x7a, 0xc8, 0x29, 0x94, 0x9e, 0x7a, 0xd0,
	0x15, 0xf5, 0x0b, 0x62, 0xbb, 0x58, 0xa1, 0x9a, 0x90, 0xae, 0x17, 0xc1, 0x43, 0xc5, 0x95, 0x96,
	0xaa, 0x09, 0xe9, 0x22, 0xf8, 0x87, 0x5e, 0x04, 0x3f, 0x41, 0xf2, 0x25, 0x65, 0xb6, 0xb9, 0x3d,
	0xde, 0xee, 0x1b, 0x66, 0xe0, 0xc4, 0x2f, 0xb7, 0x9b, 0xf3, 0x65, 0xbd, 0xbe, 0x68, 0xdc, 0xd6,
	0xeb, 0xba, 0xa9, 0x7c, 0x45, 0x31, 0x4b, 0xfd, 0x75, 0x39, 0xfa, 0x15, 0x20, 0xe7, 0xeb, 0xcf,
	0x0d, 0x21, 0x44, 0x8d, 0x7b, 0x4f, 0xc4, 0x50, 0x8c, 0x07, 0x25, 0x23, 0x9d, 0x81, 0xfa, 0x70,
	0x7e, 0x55, 0xbd, 0x26, 0xbd, 0xa1, 0x18, 0x1f, 0x5f, 0x9d, 0xea, 0x2e, 0xd2, 0x1c, 0xe8, 0xfb,
	0xf0, 0x56, 0x76, 0x7f, 0x88, 0x40, 0xae, 0x1a, 0xf7, 0x96, 0x44, 0x61, 0x40, 0x60, 0x76, 0xfe,
	0xbb, 0x76, 0x89, 0xdc, 0x3b, 0xe6, 0xd1, 0x33, 0xa8, 0x7d, 0x49, 0x31, 0x44, 0xc5, 0xa3, 0xc5,
	0x03, 0x86, 0x5b, 0x63, 0x51, 0xd0, 0x21, 0xc8, 0x99, 0xc9, 0xa6, 0xd8, 0x63, 0x2a, 0xf2, 0x85,
	0xc5, 0x88, 0x06, 0xd0, 0x2f, 0x32, 0x3b, 0x99, 0xa1, 0x24, 0x00, 0x35, 0x35, 0x73, 0x63, 0x0d,
	0xf6, 0x59, 0xdb, 0x32, 0x9b, 0x18, 0x54, 0x74, 0x04, 0x71, 0x5e, 0xd8, 0xbb, 0xfc, 0x61, 0x81,
	0xf1, 0x0d, 0xfd, 0xb4, 0xa9, 0xf8, 0x6b, 0x53, 0xf1, 0xdf, 0xa6, 0xe2, 0x49, 0xf2, 0xc6, 0x2f,
	0x2a, 0xdc, 0x7c, 0xbd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x25, 0x11, 0x55, 0x76, 0x0a, 0x01, 0x00,
	0x00,
}
