// Code generated by protoc-gen-go.
// source: azmo.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	azmo.proto

It has these top-level messages:
	BatchRequest
	Argument
	DeleteRequest
	PutRequest
	NumericRequest
	RangeRequest
	GetRequest
	WatchRequest
	Event
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import db "github.com/azmodb/db/pb"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BatchRequest struct {
	Args []*Argument `protobuf:"bytes,1,rep,name=args" json:"args,omitempty"`
}

func (m *BatchRequest) Reset()                    { *m = BatchRequest{} }
func (m *BatchRequest) String() string            { return proto.CompactTextString(m) }
func (*BatchRequest) ProtoMessage()               {}
func (*BatchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BatchRequest) GetArgs() []*Argument {
	if m != nil {
		return m.Args
	}
	return nil
}

type Argument struct {
	// Types that are valid to be assigned to Command:
	//	*Argument_Increment
	//	*Argument_Decrement
	//	*Argument_Delete
	//	*Argument_Put
	Command isArgument_Command `protobuf_oneof:"Command"`
}

func (m *Argument) Reset()                    { *m = Argument{} }
func (m *Argument) String() string            { return proto.CompactTextString(m) }
func (*Argument) ProtoMessage()               {}
func (*Argument) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isArgument_Command interface {
	isArgument_Command()
}

type Argument_Increment struct {
	Increment *NumericRequest `protobuf:"bytes,1,opt,name=increment,oneof"`
}
type Argument_Decrement struct {
	Decrement *NumericRequest `protobuf:"bytes,2,opt,name=decrement,oneof"`
}
type Argument_Delete struct {
	Delete *DeleteRequest `protobuf:"bytes,3,opt,name=delete,oneof"`
}
type Argument_Put struct {
	Put *PutRequest `protobuf:"bytes,4,opt,name=put,oneof"`
}

func (*Argument_Increment) isArgument_Command() {}
func (*Argument_Decrement) isArgument_Command() {}
func (*Argument_Delete) isArgument_Command()    {}
func (*Argument_Put) isArgument_Command()       {}

func (m *Argument) GetCommand() isArgument_Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *Argument) GetIncrement() *NumericRequest {
	if x, ok := m.GetCommand().(*Argument_Increment); ok {
		return x.Increment
	}
	return nil
}

func (m *Argument) GetDecrement() *NumericRequest {
	if x, ok := m.GetCommand().(*Argument_Decrement); ok {
		return x.Decrement
	}
	return nil
}

func (m *Argument) GetDelete() *DeleteRequest {
	if x, ok := m.GetCommand().(*Argument_Delete); ok {
		return x.Delete
	}
	return nil
}

func (m *Argument) GetPut() *PutRequest {
	if x, ok := m.GetCommand().(*Argument_Put); ok {
		return x.Put
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Argument) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Argument_OneofMarshaler, _Argument_OneofUnmarshaler, _Argument_OneofSizer, []interface{}{
		(*Argument_Increment)(nil),
		(*Argument_Decrement)(nil),
		(*Argument_Delete)(nil),
		(*Argument_Put)(nil),
	}
}

func _Argument_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Argument)
	// Command
	switch x := m.Command.(type) {
	case *Argument_Increment:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Increment); err != nil {
			return err
		}
	case *Argument_Decrement:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Decrement); err != nil {
			return err
		}
	case *Argument_Delete:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Delete); err != nil {
			return err
		}
	case *Argument_Put:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Put); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Argument.Command has unexpected type %T", x)
	}
	return nil
}

func _Argument_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Argument)
	switch tag {
	case 1: // Command.increment
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NumericRequest)
		err := b.DecodeMessage(msg)
		m.Command = &Argument_Increment{msg}
		return true, err
	case 2: // Command.decrement
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NumericRequest)
		err := b.DecodeMessage(msg)
		m.Command = &Argument_Decrement{msg}
		return true, err
	case 3: // Command.delete
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DeleteRequest)
		err := b.DecodeMessage(msg)
		m.Command = &Argument_Delete{msg}
		return true, err
	case 4: // Command.put
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PutRequest)
		err := b.DecodeMessage(msg)
		m.Command = &Argument_Put{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Argument_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Argument)
	// Command
	switch x := m.Command.(type) {
	case *Argument_Increment:
		s := proto.Size(x.Increment)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Argument_Decrement:
		s := proto.Size(x.Decrement)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Argument_Delete:
		s := proto.Size(x.Delete)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Argument_Put:
		s := proto.Size(x.Put)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type DeleteRequest struct {
	Key  []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Prev bool   `protobuf:"varint,2,opt,name=prev" json:"prev,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type PutRequest struct {
	Key       []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value     []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Prev      bool   `protobuf:"varint,3,opt,name=prev" json:"prev,omitempty"`
	Tombstone bool   `protobuf:"varint,4,opt,name=tombstone" json:"tombstone,omitempty"`
}

func (m *PutRequest) Reset()                    { *m = PutRequest{} }
func (m *PutRequest) String() string            { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()               {}
func (*PutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type NumericRequest struct {
	Key   []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value int64  `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
	Prev  bool   `protobuf:"varint,3,opt,name=prev" json:"prev,omitempty"`
}

func (m *NumericRequest) Reset()                    { *m = NumericRequest{} }
func (m *NumericRequest) String() string            { return proto.CompactTextString(m) }
func (*NumericRequest) ProtoMessage()               {}
func (*NumericRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type RangeRequest struct {
	From     []byte `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To       []byte `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Rev      int64  `protobuf:"varint,3,opt,name=rev" json:"rev,omitempty"`
	Versions bool   `protobuf:"varint,4,opt,name=versions" json:"versions,omitempty"`
}

func (m *RangeRequest) Reset()                    { *m = RangeRequest{} }
func (m *RangeRequest) String() string            { return proto.CompactTextString(m) }
func (*RangeRequest) ProtoMessage()               {}
func (*RangeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type GetRequest struct {
	Key      []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Rev      int64  `protobuf:"varint,3,opt,name=rev" json:"rev,omitempty"`
	Versions bool   `protobuf:"varint,4,opt,name=versions" json:"versions,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type WatchRequest struct {
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *WatchRequest) Reset()                    { *m = WatchRequest{} }
func (m *WatchRequest) String() string            { return proto.CompactTextString(m) }
func (*WatchRequest) ProtoMessage()               {}
func (*WatchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type Event struct {
	Record   *db.Record `protobuf:"bytes,1,opt,name=record" json:"record,omitempty"`
	Duration int64      `protobuf:"varint,2,opt,name=duration" json:"duration,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Event) GetRecord() *db.Record {
	if m != nil {
		return m.Record
	}
	return nil
}

func init() {
	proto.RegisterType((*BatchRequest)(nil), "azmo.BatchRequest")
	proto.RegisterType((*Argument)(nil), "azmo.Argument")
	proto.RegisterType((*DeleteRequest)(nil), "azmo.DeleteRequest")
	proto.RegisterType((*PutRequest)(nil), "azmo.PutRequest")
	proto.RegisterType((*NumericRequest)(nil), "azmo.NumericRequest")
	proto.RegisterType((*RangeRequest)(nil), "azmo.RangeRequest")
	proto.RegisterType((*GetRequest)(nil), "azmo.GetRequest")
	proto.RegisterType((*WatchRequest)(nil), "azmo.WatchRequest")
	proto.RegisterType((*Event)(nil), "azmo.Event")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for DB service

type DBClient interface {
	Range(ctx context.Context, in *RangeRequest, opts ...grpc.CallOption) (DB_RangeClient, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Event, error)
	Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (DB_WatchClient, error)
	Batch(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (DB_BatchClient, error)
}

type dBClient struct {
	cc *grpc.ClientConn
}

func NewDBClient(cc *grpc.ClientConn) DBClient {
	return &dBClient{cc}
}

func (c *dBClient) Range(ctx context.Context, in *RangeRequest, opts ...grpc.CallOption) (DB_RangeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DB_serviceDesc.Streams[0], c.cc, "/azmo.DB/Range", opts...)
	if err != nil {
		return nil, err
	}
	x := &dBRangeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DB_RangeClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type dBRangeClient struct {
	grpc.ClientStream
}

func (x *dBRangeClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dBClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := grpc.Invoke(ctx, "/azmo.DB/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (DB_WatchClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DB_serviceDesc.Streams[1], c.cc, "/azmo.DB/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &dBWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DB_WatchClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type dBWatchClient struct {
	grpc.ClientStream
}

func (x *dBWatchClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dBClient) Batch(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (DB_BatchClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DB_serviceDesc.Streams[2], c.cc, "/azmo.DB/Batch", opts...)
	if err != nil {
		return nil, err
	}
	x := &dBBatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DB_BatchClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type dBBatchClient struct {
	grpc.ClientStream
}

func (x *dBBatchClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for DB service

type DBServer interface {
	Range(*RangeRequest, DB_RangeServer) error
	Get(context.Context, *GetRequest) (*Event, error)
	Watch(*WatchRequest, DB_WatchServer) error
	Batch(*BatchRequest, DB_BatchServer) error
}

func RegisterDBServer(s *grpc.Server, srv DBServer) {
	s.RegisterService(&_DB_serviceDesc, srv)
}

func _DB_Range_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RangeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DBServer).Range(m, &dBRangeServer{stream})
}

type DB_RangeServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type dBRangeServer struct {
	grpc.ServerStream
}

func (x *dBRangeServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _DB_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azmo.DB/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DB_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DBServer).Watch(m, &dBWatchServer{stream})
}

type DB_WatchServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type dBWatchServer struct {
	grpc.ServerStream
}

func (x *dBWatchServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _DB_Batch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DBServer).Batch(m, &dBBatchServer{stream})
}

type DB_BatchServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type dBBatchServer struct {
	grpc.ServerStream
}

func (x *dBBatchServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

var _DB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "azmo.DB",
	HandlerType: (*DBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _DB_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Range",
			Handler:       _DB_Range_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Watch",
			Handler:       _DB_Watch_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Batch",
			Handler:       _DB_Batch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("azmo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0x36, 0x4d, 0x5a, 0xdb, 0xb3, 0xb1, 0x2c, 0xe3, 0x5e, 0x94, 0xe2, 0x45, 0x19, 0xbc, 0x10,
	0xc1, 0xac, 0x54, 0x7d, 0x00, 0xe3, 0xca, 0x7a, 0x21, 0x22, 0x73, 0x23, 0x78, 0x65, 0x7e, 0x66,
	0xb3, 0xc1, 0x26, 0x13, 0x27, 0x93, 0x80, 0x3e, 0x94, 0x8f, 0xe3, 0xf3, 0x38, 0x73, 0x32, 0xf9,
	0xa9, 0x3f, 0x15, 0xa1, 0x94, 0x73, 0xce, 0x7c, 0xdf, 0x77, 0x7e, 0xf8, 0x02, 0x10, 0x7d, 0x2b,
	0x44, 0x50, 0x49, 0xa1, 0x04, 0xf1, 0x4c, 0xbc, 0xdd, 0x65, 0xb9, 0xba, 0x6d, 0xe2, 0x20, 0x11,
	0xc5, 0xa5, 0x29, 0xa4, 0xf1, 0xa5, 0xfe, 0x55, 0xe6, 0xbf, 0xc3, 0xd1, 0x3d, 0xf8, 0x61, 0xa4,
	0x92, 0x5b, 0xc6, 0xbf, 0x34, 0xbc, 0x56, 0x84, 0x82, 0x17, 0xc9, 0xac, 0xde, 0x38, 0x3b, 0xf7,
	0xd1, 0xd9, 0x7e, 0x1d, 0xa0, 0xe4, 0x4b, 0x99, 0x35, 0x05, 0x2f, 0x15, 0xc3, 0x37, 0xfa, 0xc3,
	0x81, 0x65, 0x5f, 0x22, 0xcf, 0x61, 0x95, 0x97, 0x89, 0xe4, 0x26, 0xd1, 0x2c, 0x47, 0xb3, 0x2e,
	0x3a, 0xd6, 0x3b, 0x0d, 0x90, 0x79, 0x62, 0x95, 0xdf, 0xdc, 0x61, 0x23, 0xd0, 0xb0, 0x52, 0xde,
	0xb3, 0x66, 0xa7, 0x59, 0x03, 0x90, 0x3c, 0x81, 0x45, 0xca, 0x0f, 0x5c, 0xf1, 0x8d, 0x8b, 0x94,
	0xfb, 0x1d, 0xe5, 0x0a, 0x6b, 0x23, 0xc3, 0x82, 0xc8, 0x43, 0x70, 0xab, 0x46, 0x6d, 0x3c, 0xc4,
	0x9e, 0x77, 0xd8, 0xf7, 0x8d, 0x1a, 0x81, 0xe6, 0x39, 0x5c, 0xc1, 0xdd, 0x57, 0xa2, 0x28, 0xa2,
	0x32, 0xa5, 0x2f, 0xe0, 0xde, 0x91, 0x16, 0x39, 0x07, 0xf7, 0x33, 0xff, 0x8a, 0x6b, 0xf9, 0xcc,
	0x84, 0x84, 0x80, 0x57, 0x49, 0xde, 0xe2, 0xcc, 0x4b, 0x86, 0x31, 0xbd, 0x01, 0x18, 0x65, 0xff,
	0xc0, 0xb9, 0x80, 0x79, 0x1b, 0x1d, 0x1a, 0x8e, 0x24, 0x9f, 0x75, 0xc9, 0xa0, 0xe4, 0x8e, 0x4a,
	0xe4, 0x01, 0xac, 0x94, 0x28, 0xe2, 0x5a, 0x89, 0x92, 0xe3, 0xdc, 0x4b, 0x36, 0x16, 0xe8, 0x5b,
	0x58, 0x1f, 0x5f, 0xe7, 0x5f, 0xbd, 0xdc, 0x13, 0xbd, 0xe8, 0x27, 0xf0, 0x59, 0x54, 0x66, 0xc3,
	0xae, 0x1a, 0x73, 0x23, 0x45, 0x61, 0xc5, 0x30, 0x26, 0x6b, 0x98, 0x29, 0x61, 0xc7, 0xd6, 0x91,
	0xe9, 0xd7, 0xcb, 0xb8, 0xcc, 0x84, 0x64, 0x0b, 0xcb, 0x96, 0xcb, 0x3a, 0x17, 0x65, 0x6d, 0x07,
	0x1e, 0x72, 0x3d, 0x2f, 0x5c, 0xf3, 0x13, 0x77, 0xf9, 0x3f, 0xb5, 0x1d, 0xf8, 0x1f, 0xa6, 0x4e,
	0xfd, 0x4d, 0x8f, 0x5e, 0xc3, 0xfc, 0x75, 0x6b, 0x7c, 0x42, 0x61, 0x21, 0x79, 0x22, 0x64, 0x6a,
	0x0d, 0x09, 0x81, 0xf6, 0x3b, 0xc3, 0x0a, 0xb3, 0x2f, 0xa6, 0x55, 0xda, 0xc8, 0x48, 0x69, 0x6d,
	0x7b, 0xab, 0x21, 0xdf, 0x7f, 0x77, 0x60, 0x76, 0x15, 0x92, 0xc7, 0x30, 0xc7, 0x0b, 0x11, 0xd2,
	0x79, 0x67, 0x7a, 0xae, 0xed, 0x59, 0x57, 0xc3, 0x86, 0x4f, 0x1d, 0xe3, 0x35, 0xbd, 0x2b, 0xb1,
	0x2e, 0x1b, 0xd7, 0x3e, 0xc2, 0x19, 0x45, 0xdc, 0xa1, 0x57, 0x9c, 0x2e, 0xf4, 0xab, 0xa2, 0xc6,
	0x86, 0x53, 0x6c, 0xf8, 0x77, 0x6c, 0xe8, 0x7d, 0x9c, 0x55, 0x71, 0xbc, 0xc0, 0x4f, 0xfa, 0xd9,
	0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x89, 0x82, 0x19, 0x0c, 0x08, 0x04, 0x00, 0x00,
}
