/*
Copyright 2019 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event_state.proto

package event_state

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EventsRecord_Type int32

const (
	EventsRecord_SENT     EventsRecord_Type = 0
	EventsRecord_ACCEPTED EventsRecord_Type = 1
	EventsRecord_FAILED   EventsRecord_Type = 2
	EventsRecord_RECEIVED EventsRecord_Type = 3
)

var EventsRecord_Type_name = map[int32]string{
	0: "SENT",
	1: "ACCEPTED",
	2: "FAILED",
	3: "RECEIVED",
}

var EventsRecord_Type_value = map[string]int32{
	"SENT":     0,
	"ACCEPTED": 1,
	"FAILED":   2,
	"RECEIVED": 3,
}

func (x EventsRecord_Type) String() string {
	return proto.EnumName(EventsRecord_Type_name, int32(x))
}

func (EventsRecord_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_de3fba9d879b76ae, []int{0, 0}
}

type EventsRecord struct {
	Events               map[string]*timestamp.Timestamp `protobuf:"bytes,1,rep,name=Events,proto3" json:"Events,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Type                 EventsRecord_Type               `protobuf:"varint,4,opt,name=type,proto3,enum=event_state.EventsRecord_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *EventsRecord) Reset()         { *m = EventsRecord{} }
func (m *EventsRecord) String() string { return proto.CompactTextString(m) }
func (*EventsRecord) ProtoMessage()    {}
func (*EventsRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_de3fba9d879b76ae, []int{0}
}

func (m *EventsRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventsRecord.Unmarshal(m, b)
}
func (m *EventsRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventsRecord.Marshal(b, m, deterministic)
}
func (m *EventsRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventsRecord.Merge(m, src)
}
func (m *EventsRecord) XXX_Size() int {
	return xxx_messageInfo_EventsRecord.Size(m)
}
func (m *EventsRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_EventsRecord.DiscardUnknown(m)
}

var xxx_messageInfo_EventsRecord proto.InternalMessageInfo

func (m *EventsRecord) GetEvents() map[string]*timestamp.Timestamp {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *EventsRecord) GetType() EventsRecord_Type {
	if m != nil {
		return m.Type
	}
	return EventsRecord_SENT
}

type RecordReply struct {
	Count                uint64   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecordReply) Reset()         { *m = RecordReply{} }
func (m *RecordReply) String() string { return proto.CompactTextString(m) }
func (*RecordReply) ProtoMessage()    {}
func (*RecordReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_de3fba9d879b76ae, []int{1}
}

func (m *RecordReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordReply.Unmarshal(m, b)
}
func (m *RecordReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordReply.Marshal(b, m, deterministic)
}
func (m *RecordReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordReply.Merge(m, src)
}
func (m *RecordReply) XXX_Size() int {
	return xxx_messageInfo_RecordReply.Size(m)
}
func (m *RecordReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordReply.DiscardUnknown(m)
}

var xxx_messageInfo_RecordReply proto.InternalMessageInfo

func (m *RecordReply) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterEnum("event_state.EventsRecord_Type", EventsRecord_Type_name, EventsRecord_Type_value)
	proto.RegisterType((*EventsRecord)(nil), "event_state.EventsRecord")
	proto.RegisterMapType((map[string]*timestamp.Timestamp)(nil), "event_state.EventsRecord.EventsEntry")
	proto.RegisterType((*RecordReply)(nil), "event_state.RecordReply")
}

func init() { proto.RegisterFile("event_state.proto", fileDescriptor_de3fba9d879b76ae) }

var fileDescriptor_de3fba9d879b76ae = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xbb, 0x69, 0x5a, 0xea, 0xa4, 0x94, 0xb8, 0x78, 0x88, 0x39, 0x68, 0x88, 0x08, 0x39,
	0xa5, 0x12, 0x2f, 0x45, 0xf0, 0x50, 0xda, 0x15, 0x0a, 0x22, 0xb2, 0x46, 0x3d, 0x4a, 0x5a, 0xc7,
	0x22, 0xa6, 0xd9, 0x90, 0x6c, 0x0a, 0xf9, 0x1b, 0xfe, 0x62, 0x49, 0x36, 0x85, 0xf5, 0xd0, 0xdb,
	0x4c, 0xde, 0x7b, 0x79, 0xf3, 0x2d, 0x9c, 0xe2, 0x1e, 0x33, 0xf9, 0x51, 0xca, 0x44, 0x62, 0x98,
	0x17, 0x42, 0x0a, 0x6a, 0x69, 0x9f, 0xdc, 0xcb, 0xad, 0x10, 0xdb, 0x14, 0xa7, 0xad, 0xb4, 0xae,
	0xbe, 0xa6, 0xf2, 0x7b, 0x87, 0xa5, 0x4c, 0x76, 0xb9, 0x72, 0xfb, 0xbf, 0x06, 0x8c, 0x59, 0x13,
	0x28, 0x39, 0x6e, 0x44, 0xf1, 0x49, 0xef, 0x61, 0xa8, 0x76, 0x87, 0x78, 0xfd, 0xc0, 0x8a, 0xae,
	0x43, 0xbd, 0x42, 0xb7, 0x76, 0x0b, 0xcb, 0x64, 0x51, 0xf3, 0x2e, 0x44, 0x23, 0x30, 0x65, 0x9d,
	0xa3, 0x63, 0x7a, 0x24, 0x98, 0x44, 0x17, 0xc7, 0xc3, 0x71, 0x9d, 0x23, 0x6f, 0xbd, 0xee, 0x2b,
	0x58, 0xda, 0xaf, 0xa8, 0x0d, 0xfd, 0x1f, 0xac, 0x1d, 0xe2, 0x91, 0xe0, 0x84, 0x37, 0x23, 0xbd,
	0x81, 0xc1, 0x3e, 0x49, 0x2b, 0x74, 0x0c, 0x8f, 0x04, 0x56, 0xe4, 0x86, 0x8a, 0x2a, 0x3c, 0x50,
	0x85, 0xf1, 0x81, 0x8a, 0x2b, 0xe3, 0x9d, 0x31, 0x23, 0xfe, 0x0c, 0xcc, 0xa6, 0x84, 0x8e, 0xc0,
	0x7c, 0x61, 0x4f, 0xb1, 0xdd, 0xa3, 0x63, 0x18, 0xcd, 0x17, 0x0b, 0xf6, 0x1c, 0xb3, 0xa5, 0x4d,
	0x28, 0xc0, 0xf0, 0x61, 0xbe, 0x7a, 0x64, 0x4b, 0xdb, 0x68, 0x14, 0xce, 0x16, 0x6c, 0xf5, 0xc6,
	0x96, 0x76, 0xdf, 0xbf, 0x02, 0x4b, 0x5d, 0xc9, 0x31, 0x4f, 0x6b, 0x7a, 0x06, 0x83, 0x8d, 0xa8,
	0x32, 0xd9, 0x9e, 0x64, 0x72, 0xb5, 0x44, 0xef, 0x30, 0xd1, 0x81, 0xb0, 0xa0, 0x0c, 0xc6, 0x6a,
	0xee, 0xde, 0xe2, 0xfc, 0x28, 0xbd, 0xeb, 0xfc, 0x93, 0xb4, 0x32, 0xbf, 0xb7, 0x1e, 0xb6, 0x58,
	0xb7, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x89, 0xaa, 0x3f, 0xdc, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventsRecorderClient is the client API for EventsRecorder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventsRecorderClient interface {
	RecordEvents(ctx context.Context, in *EventsRecord, opts ...grpc.CallOption) (*RecordReply, error)
}

type eventsRecorderClient struct {
	cc *grpc.ClientConn
}

func NewEventsRecorderClient(cc *grpc.ClientConn) EventsRecorderClient {
	return &eventsRecorderClient{cc}
}

func (c *eventsRecorderClient) RecordEvents(ctx context.Context, in *EventsRecord, opts ...grpc.CallOption) (*RecordReply, error) {
	out := new(RecordReply)
	err := c.cc.Invoke(ctx, "/event_state.EventsRecorder/RecordEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventsRecorderServer is the server API for EventsRecorder service.
type EventsRecorderServer interface {
	RecordEvents(context.Context, *EventsRecord) (*RecordReply, error)
}

// UnimplementedEventsRecorderServer can be embedded to have forward compatible implementations.
type UnimplementedEventsRecorderServer struct {
}

func (*UnimplementedEventsRecorderServer) RecordEvents(ctx context.Context, req *EventsRecord) (*RecordReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordEvents not implemented")
}

func RegisterEventsRecorderServer(s *grpc.Server, srv EventsRecorderServer) {
	s.RegisterService(&_EventsRecorder_serviceDesc, srv)
}

func _EventsRecorder_RecordEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventsRecord)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsRecorderServer).RecordEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event_state.EventsRecorder/RecordEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsRecorderServer).RecordEvents(ctx, req.(*EventsRecord))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventsRecorder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "event_state.EventsRecorder",
	HandlerType: (*EventsRecorderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecordEvents",
			Handler:    _EventsRecorder_RecordEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "event_state.proto",
}
