// Code generated by brick-protoc.
// source: codec.proto
// DO NOT EDIT!

/*
	Package pb is a generated protocol buffer package.

	It is generated from these files:
		codec.proto

	It has these top-level messages:
		Request
		Response
		Error
*/
package pb

import proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Method int32

const (
	InvalidMethod Method = 0
	FlowMethod    Method = 1
	AsyncMethod   Method = 2
)

var Method_name = map[int32]string{
	0: "InvalidMethod",
	1: "FlowMethod",
	2: "AsyncMethod",
}
var Method_value = map[string]int32{
	"InvalidMethod": 0,
	"FlowMethod":    1,
	"AsyncMethod":   2,
}

func (x Method) String() string {
	return proto.EnumName(Method_name, int32(x))
}
func (Method) EnumDescriptor() ([]byte, []int) { return fileDescriptorCodec, []int{0} }

type Request struct {
	Method Method `protobuf:"varint,1,opt,name=method,proto3,enum=pb.Method" json:"method,omitempty"`
	Seq    uint64 `protobuf:"varint,2,opt,name=seq,proto3" json:"seq,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptorCodec, []int{0} }

type Response struct {
	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Seq    uint64 `protobuf:"varint,2,opt,name=seq,proto3" json:"seq,omitempty"`
	Error  *Error `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptorCodec, []int{1} }

type Error struct {
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	ErrCode     int32  `protobuf:"varint,2,opt,name=err_code,json=errCode,proto3" json:"err_code,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptorCodec, []int{2} }

func init() {
	proto.RegisterType((*Request)(nil), "pb.Request")
	proto.RegisterType((*Response)(nil), "pb.Response")
	proto.RegisterType((*Error)(nil), "pb.Error")
	proto.RegisterEnum("pb.Method", Method_name, Method_value)
}
func (m *Request) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Request) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Method != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintCodec(data, i, uint64(m.Method))
	}
	if m.Seq != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintCodec(data, i, uint64(m.Seq))
	}
	return i, nil
}

func (m *Response) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Response) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Method) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintCodec(data, i, uint64(len(m.Method)))
		i += copy(data[i:], m.Method)
	}
	if m.Seq != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintCodec(data, i, uint64(m.Seq))
	}
	if m.Error != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintCodec(data, i, uint64(m.Error.Size()))
		n1, err := m.Error.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *Error) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Error) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Description) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintCodec(data, i, uint64(len(m.Description)))
		i += copy(data[i:], m.Description)
	}
	if m.ErrCode != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintCodec(data, i, uint64(m.ErrCode))
	}
	return i, nil
}

func encodeFixed64Codec(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Codec(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintCodec(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Request) Size() (n int) {
	var l int
	_ = l
	if m.Method != 0 {
		n += 1 + sovCodec(uint64(m.Method))
	}
	if m.Seq != 0 {
		n += 1 + sovCodec(uint64(m.Seq))
	}
	return n
}

func (m *Response) Size() (n int) {
	var l int
	_ = l
	l = len(m.Method)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if m.Seq != 0 {
		n += 1 + sovCodec(uint64(m.Seq))
	}
	if m.Error != nil {
		l = m.Error.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func (m *Error) Size() (n int) {
	var l int
	_ = l
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if m.ErrCode != 0 {
		n += 1 + sovCodec(uint64(m.ErrCode))
	}
	return n
}

func sovCodec(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCodec(x uint64) (n int) {
	return sovCodec(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return errIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			m.Method = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return errIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Method |= (Method(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seq", wireType)
			}
			m.Seq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return errIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Seq |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return errInvalidLengthCodec
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
func (m *Response) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return errIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return errIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return errInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Method = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seq", wireType)
			}
			m.Seq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return errIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Seq |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return errIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return errInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Error == nil {
				m.Error = &Error{}
			}
			if err := m.Error.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return errInvalidLengthCodec
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
func (m *Error) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return errIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Error: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Error: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return errIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return errInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrCode", wireType)
			}
			m.ErrCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return errIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ErrCode |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return errInvalidLengthCodec
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
func skipCodec(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, errIntOverflowCodec
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, errIntOverflowCodec
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, errIntOverflowCodec
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, errInvalidLengthCodec
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, errIntOverflowCodec
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipCodec(data[start:])
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
	errInvalidLengthCodec = fmt.Errorf("proto: negative length found during unmarshaling")
	errIntOverflowCodec   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("codec.proto", fileDescriptorCodec) }

var fileDescriptorCodec = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x90, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x33, 0xa9, 0x49, 0xdb, 0x1b, 0xac, 0xf5, 0x22, 0x12, 0x5d, 0x8c, 0x21, 0xab, 0xe0,
	0x22, 0x42, 0xdd, 0x0a, 0xe2, 0x2f, 0xb8, 0x70, 0x33, 0xe0, 0x5a, 0x4c, 0x72, 0xa9, 0x81, 0x9a,
	0x9b, 0xce, 0x44, 0xc5, 0x37, 0xf1, 0x91, 0xba, 0xec, 0x23, 0xd8, 0xf8, 0x22, 0x92, 0x9f, 0x85,
	0x2e, 0xdc, 0xdd, 0xf3, 0x1d, 0xe6, 0x1c, 0xce, 0x80, 0x97, 0x72, 0x46, 0x69, 0x5c, 0x6a, 0xae,
	0x18, 0xed, 0x32, 0x39, 0xdc, 0x9b, 0xf3, 0x9c, 0x5b, 0x79, 0xd2, 0x5c, 0x9d, 0x13, 0x9e, 0xc3,
	0x50, 0xd1, 0xf2, 0x95, 0x4c, 0x85, 0x21, 0xb8, 0x2f, 0x54, 0x3d, 0x73, 0xe6, 0x8b, 0x40, 0x44,
	0x93, 0x19, 0xc4, 0x65, 0x12, 0xdf, 0xb7, 0x44, 0xf5, 0x0e, 0x4e, 0x61, 0x60, 0x68, 0xe9, 0xdb,
	0x81, 0x88, 0xb6, 0x54, 0x73, 0x86, 0x0f, 0x30, 0x52, 0x64, 0x4a, 0x2e, 0x0c, 0xe1, 0xfe, 0x9f,
	0x84, 0xf1, 0xff, 0xaf, 0xf0, 0x08, 0x1c, 0xd2, 0x9a, 0xb5, 0x3f, 0x08, 0x44, 0xe4, 0xcd, 0xc6,
	0x4d, 0xd5, 0x4d, 0x03, 0x54, 0xc7, 0xc3, 0x6b, 0x70, 0x5a, 0x8d, 0x01, 0x78, 0x19, 0x99, 0x54,
	0xe7, 0x65, 0x95, 0x73, 0xd1, 0x07, 0xff, 0x46, 0x78, 0x00, 0x23, 0xd2, 0xfa, 0xb1, 0xd9, 0xdb,
	0x56, 0x38, 0x6a, 0x48, 0x5a, 0x5f, 0x71, 0x46, 0xc7, 0x67, 0xe0, 0x76, 0x03, 0x70, 0x17, 0xb6,
	0xef, 0x8a, 0xb7, 0xa7, 0x45, 0x9e, 0x75, 0x60, 0x6a, 0xe1, 0x04, 0xe0, 0x76, 0xc1, 0xef, 0xbd,
	0x16, 0xb8, 0x03, 0xde, 0x85, 0xf9, 0x28, 0xd2, 0x1e, 0xd8, 0x97, 0xfe, 0x6a, 0x23, 0xad, 0xf5,
	0x46, 0x5a, 0xab, 0x5a, 0x8a, 0x75, 0x2d, 0xc5, 0x57, 0x2d, 0xc5, 0xe7, 0xb7, 0xb4, 0x12, 0xb7,
	0xfd, 0xbc, 0xd3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x4c, 0xb9, 0xfb, 0x65, 0x01, 0x00,
	0x00,
}