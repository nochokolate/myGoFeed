// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feed.proto

package feed

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type OK struct {
	OK                   bool     `protobuf:"varint,1,opt,name=OK,proto3" json:"OK,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OK) Reset()         { *m = OK{} }
func (m *OK) String() string { return proto.CompactTextString(m) }
func (*OK) ProtoMessage()    {}
func (*OK) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7a672c1337cb5ac, []int{0}
}

func (m *OK) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OK.Unmarshal(m, b)
}
func (m *OK) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OK.Marshal(b, m, deterministic)
}
func (m *OK) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OK.Merge(m, src)
}
func (m *OK) XXX_Size() int {
	return xxx_messageInfo_OK.Size(m)
}
func (m *OK) XXX_DiscardUnknown() {
	xxx_messageInfo_OK.DiscardUnknown(m)
}

var xxx_messageInfo_OK proto.InternalMessageInfo

func (m *OK) GetOK() bool {
	if m != nil {
		return m.OK
	}
	return false
}

type Feed struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Start                int64    `protobuf:"varint,2,opt,name=Start,proto3" json:"Start,omitempty"`
	End                  int64    `protobuf:"varint,3,opt,name=End,proto3" json:"End,omitempty"`
	StartTime            string   `protobuf:"bytes,4,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime              string   `protobuf:"bytes,5,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	Keywords             string   `protobuf:"bytes,6,opt,name=Keywords,proto3" json:"Keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Feed) Reset()         { *m = Feed{} }
func (m *Feed) String() string { return proto.CompactTextString(m) }
func (*Feed) ProtoMessage()    {}
func (*Feed) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7a672c1337cb5ac, []int{1}
}

func (m *Feed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Feed.Unmarshal(m, b)
}
func (m *Feed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Feed.Marshal(b, m, deterministic)
}
func (m *Feed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Feed.Merge(m, src)
}
func (m *Feed) XXX_Size() int {
	return xxx_messageInfo_Feed.Size(m)
}
func (m *Feed) XXX_DiscardUnknown() {
	xxx_messageInfo_Feed.DiscardUnknown(m)
}

var xxx_messageInfo_Feed proto.InternalMessageInfo

func (m *Feed) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Feed) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Feed) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *Feed) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *Feed) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *Feed) GetKeywords() string {
	if m != nil {
		return m.Keywords
	}
	return ""
}

type FeedGroup struct {
	FeedNames            []string `protobuf:"bytes,1,rep,name=FeedNames,proto3" json:"FeedNames,omitempty"`
	Nums                 int64    `protobuf:"varint,2,opt,name=Nums,proto3" json:"Nums,omitempty"`
	StartTime            string   `protobuf:"bytes,3,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime              string   `protobuf:"bytes,4,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	Keywords             string   `protobuf:"bytes,5,opt,name=Keywords,proto3" json:"Keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedGroup) Reset()         { *m = FeedGroup{} }
func (m *FeedGroup) String() string { return proto.CompactTextString(m) }
func (*FeedGroup) ProtoMessage()    {}
func (*FeedGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7a672c1337cb5ac, []int{2}
}

func (m *FeedGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedGroup.Unmarshal(m, b)
}
func (m *FeedGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedGroup.Marshal(b, m, deterministic)
}
func (m *FeedGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedGroup.Merge(m, src)
}
func (m *FeedGroup) XXX_Size() int {
	return xxx_messageInfo_FeedGroup.Size(m)
}
func (m *FeedGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedGroup.DiscardUnknown(m)
}

var xxx_messageInfo_FeedGroup proto.InternalMessageInfo

func (m *FeedGroup) GetFeedNames() []string {
	if m != nil {
		return m.FeedNames
	}
	return nil
}

func (m *FeedGroup) GetNums() int64 {
	if m != nil {
		return m.Nums
	}
	return 0
}

func (m *FeedGroup) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *FeedGroup) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *FeedGroup) GetKeywords() string {
	if m != nil {
		return m.Keywords
	}
	return ""
}

type Story struct {
	Title                string   `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Link                 string   `protobuf:"bytes,2,opt,name=Link,proto3" json:"Link,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	PubTime              string   `protobuf:"bytes,4,opt,name=PubTime,proto3" json:"PubTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Story) Reset()         { *m = Story{} }
func (m *Story) String() string { return proto.CompactTextString(m) }
func (*Story) ProtoMessage()    {}
func (*Story) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7a672c1337cb5ac, []int{3}
}

func (m *Story) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Story.Unmarshal(m, b)
}
func (m *Story) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Story.Marshal(b, m, deterministic)
}
func (m *Story) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Story.Merge(m, src)
}
func (m *Story) XXX_Size() int {
	return xxx_messageInfo_Story.Size(m)
}
func (m *Story) XXX_DiscardUnknown() {
	xxx_messageInfo_Story.DiscardUnknown(m)
}

var xxx_messageInfo_Story proto.InternalMessageInfo

func (m *Story) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Story) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *Story) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Story) GetPubTime() string {
	if m != nil {
		return m.PubTime
	}
	return ""
}

type Status struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7a672c1337cb5ac, []int{4}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*OK)(nil), "feed.OK")
	proto.RegisterType((*Feed)(nil), "feed.Feed")
	proto.RegisterType((*FeedGroup)(nil), "feed.FeedGroup")
	proto.RegisterType((*Story)(nil), "feed.Story")
	proto.RegisterType((*Status)(nil), "feed.Status")
}

func init() {
	proto.RegisterFile("feed.proto", fileDescriptor_d7a672c1337cb5ac)
}

var fileDescriptor_d7a672c1337cb5ac = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xdb, 0xaa, 0xd3, 0x50,
	0x10, 0x3d, 0x3b, 0x49, 0x2f, 0x99, 0x53, 0x3c, 0xb2, 0xed, 0x43, 0x28, 0x22, 0x21, 0xa0, 0x44,
	0x0a, 0xb5, 0xa8, 0x3f, 0x50, 0xb0, 0xf6, 0x21, 0x62, 0x25, 0xe9, 0x0f, 0x34, 0xdd, 0x53, 0x0c,
	0xe6, 0x52, 0x92, 0x1d, 0xa4, 0x9f, 0xe1, 0x8b, 0x1f, 0xea, 0x17, 0xc8, 0x4c, 0xd2, 0x8b, 0x2d,
	0x2d, 0xbe, 0x84, 0xb5, 0x66, 0xcd, 0x9e, 0x59, 0x33, 0x43, 0x00, 0xb6, 0x88, 0x6a, 0xb2, 0x2b,
	0x0b, 0x5d, 0x48, 0x8b, 0xb0, 0x37, 0x04, 0x63, 0x19, 0xc8, 0x67, 0xf4, 0x75, 0x84, 0x2b, 0xfc,
	0x7e, 0x68, 0x2c, 0x03, 0xef, 0xb7, 0x00, 0xeb, 0x33, 0xa2, 0x92, 0x12, 0xac, 0xaf, 0xeb, 0x0c,
	0x59, 0xb2, 0x43, 0xc6, 0x72, 0x08, 0x9d, 0x48, 0xaf, 0x4b, 0xed, 0x18, 0xae, 0xf0, 0xcd, 0xb0,
	0x21, 0xf2, 0x39, 0x98, 0xf3, 0x5c, 0x39, 0x26, 0xc7, 0x08, 0xca, 0x97, 0x60, 0xb3, 0xb4, 0x4a,
	0x32, 0x74, 0x2c, 0x2e, 0x70, 0x0a, 0x48, 0x07, 0x7a, 0xf3, 0x5c, 0xb1, 0xd6, 0x61, 0xed, 0x40,
	0xe5, 0x08, 0xfa, 0x01, 0xee, 0x7f, 0x16, 0xa5, 0xaa, 0x9c, 0x2e, 0x4b, 0x47, 0xee, 0xfd, 0x12,
	0x60, 0x93, 0xb1, 0x45, 0x59, 0xd4, 0x3b, 0xea, 0x40, 0x84, 0x5c, 0x55, 0x8e, 0x70, 0x4d, 0xea,
	0x70, 0x0c, 0xb0, 0xf7, 0x3a, 0xab, 0x5a, 0x9b, 0x8c, 0xff, 0xf5, 0x64, 0xde, 0xf1, 0x64, 0xdd,
	0xf6, 0xd4, 0xb9, 0xf0, 0x94, 0xd1, 0x3e, 0x8a, 0x72, 0x4f, 0x8b, 0x59, 0x25, 0x3a, 0x3d, 0x6c,
	0xab, 0x21, 0x64, 0xe3, 0x4b, 0x92, 0xff, 0x60, 0x1b, 0x76, 0xc8, 0x58, 0xba, 0xf0, 0xf8, 0x09,
	0xab, 0x4d, 0x99, 0xec, 0x74, 0x52, 0xe4, 0xad, 0x91, 0xf3, 0x10, 0x59, 0xf9, 0x56, 0xc7, 0xe7,
	0x56, 0x5a, 0xea, 0xbd, 0x82, 0x6e, 0xa4, 0xd7, 0xba, 0xae, 0xa8, 0xdf, 0xa6, 0xa8, 0x73, 0xcd,
	0xfd, 0xcc, 0xb0, 0x21, 0xef, 0xff, 0x18, 0x60, 0x86, 0x51, 0x24, 0xc7, 0xf0, 0xb4, 0x40, 0x3d,
	0x4b, 0xd3, 0xd3, 0x46, 0xfa, 0x13, 0xbe, 0xff, 0x32, 0x18, 0x41, 0x83, 0x48, 0xf2, 0x1e, 0xa6,
	0x42, 0xbe, 0x81, 0xde, 0x02, 0x35, 0x9f, 0xfc, 0x4c, 0x1a, 0x3d, 0x36, 0x98, 0xc7, 0xe3, 0xbc,
	0x29, 0x0c, 0xda, 0xbc, 0xe6, 0x02, 0x4f, 0xa7, 0x64, 0x0e, 0x5c, 0xbf, 0x78, 0x0d, 0xbd, 0x99,
	0x52, 0x57, 0x95, 0x07, 0x87, 0x3c, 0x9a, 0xc4, 0x7b, 0x90, 0xef, 0x60, 0xd0, 0xa6, 0xdd, 0x28,
	0x7c, 0xf9, 0xe0, 0x2d, 0xd8, 0x11, 0xe6, 0xfc, 0xa2, 0xba, 0x57, 0xd9, 0x17, 0xf2, 0x23, 0xbc,
	0x88, 0xea, 0x98, 0x76, 0x1b, 0xe3, 0x2c, 0x57, 0x21, 0x6e, 0x4b, 0xac, 0xbe, 0xdf, 0x19, 0xd4,
	0x17, 0x53, 0x21, 0xc7, 0x74, 0xa3, 0x14, 0x35, 0xfe, 0x47, 0x8b, 0xb8, 0xcb, 0xff, 0xd4, 0x87,
	0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x70, 0x21, 0x3b, 0x61, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RSSClient is the client API for RSS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RSSClient interface {
	GetAllFeedNames(ctx context.Context, in *OK, opts ...grpc.CallOption) (RSS_GetAllFeedNamesClient, error)
	GetFeed(ctx context.Context, in *Feed, opts ...grpc.CallOption) (RSS_GetFeedClient, error)
	GetFeedGroup(ctx context.Context, in *FeedGroup, opts ...grpc.CallOption) (RSS_GetFeedGroupClient, error)
	AddFeed(ctx context.Context, in *Feed, opts ...grpc.CallOption) (*Status, error)
	AddFeedGroup(ctx context.Context, in *FeedGroup, opts ...grpc.CallOption) (*Status, error)
	SendFeeds(ctx context.Context, opts ...grpc.CallOption) (RSS_SendFeedsClient, error)
	SubscribeAndRefresh(ctx context.Context, opts ...grpc.CallOption) (RSS_SubscribeAndRefreshClient, error)
	DeleteFeeds(ctx context.Context, opts ...grpc.CallOption) (RSS_DeleteFeedsClient, error)
}

type rSSClient struct {
	cc grpc.ClientConnInterface
}

func NewRSSClient(cc grpc.ClientConnInterface) RSSClient {
	return &rSSClient{cc}
}

func (c *rSSClient) GetAllFeedNames(ctx context.Context, in *OK, opts ...grpc.CallOption) (RSS_GetAllFeedNamesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RSS_serviceDesc.Streams[0], "/feed.RSS/GetAllFeedNames", opts...)
	if err != nil {
		return nil, err
	}
	x := &rSSGetAllFeedNamesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RSS_GetAllFeedNamesClient interface {
	Recv() (*Feed, error)
	grpc.ClientStream
}

type rSSGetAllFeedNamesClient struct {
	grpc.ClientStream
}

func (x *rSSGetAllFeedNamesClient) Recv() (*Feed, error) {
	m := new(Feed)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rSSClient) GetFeed(ctx context.Context, in *Feed, opts ...grpc.CallOption) (RSS_GetFeedClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RSS_serviceDesc.Streams[1], "/feed.RSS/GetFeed", opts...)
	if err != nil {
		return nil, err
	}
	x := &rSSGetFeedClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RSS_GetFeedClient interface {
	Recv() (*Story, error)
	grpc.ClientStream
}

type rSSGetFeedClient struct {
	grpc.ClientStream
}

func (x *rSSGetFeedClient) Recv() (*Story, error) {
	m := new(Story)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rSSClient) GetFeedGroup(ctx context.Context, in *FeedGroup, opts ...grpc.CallOption) (RSS_GetFeedGroupClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RSS_serviceDesc.Streams[2], "/feed.RSS/GetFeedGroup", opts...)
	if err != nil {
		return nil, err
	}
	x := &rSSGetFeedGroupClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RSS_GetFeedGroupClient interface {
	Recv() (*Story, error)
	grpc.ClientStream
}

type rSSGetFeedGroupClient struct {
	grpc.ClientStream
}

func (x *rSSGetFeedGroupClient) Recv() (*Story, error) {
	m := new(Story)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rSSClient) AddFeed(ctx context.Context, in *Feed, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/feed.RSS/AddFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rSSClient) AddFeedGroup(ctx context.Context, in *FeedGroup, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/feed.RSS/AddFeedGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rSSClient) SendFeeds(ctx context.Context, opts ...grpc.CallOption) (RSS_SendFeedsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RSS_serviceDesc.Streams[3], "/feed.RSS/SendFeeds", opts...)
	if err != nil {
		return nil, err
	}
	x := &rSSSendFeedsClient{stream}
	return x, nil
}

type RSS_SendFeedsClient interface {
	Send(*Feed) error
	CloseAndRecv() (*Status, error)
	grpc.ClientStream
}

type rSSSendFeedsClient struct {
	grpc.ClientStream
}

func (x *rSSSendFeedsClient) Send(m *Feed) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rSSSendFeedsClient) CloseAndRecv() (*Status, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Status)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rSSClient) SubscribeAndRefresh(ctx context.Context, opts ...grpc.CallOption) (RSS_SubscribeAndRefreshClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RSS_serviceDesc.Streams[4], "/feed.RSS/SubscribeAndRefresh", opts...)
	if err != nil {
		return nil, err
	}
	x := &rSSSubscribeAndRefreshClient{stream}
	return x, nil
}

type RSS_SubscribeAndRefreshClient interface {
	Send(*Feed) error
	Recv() (*Story, error)
	grpc.ClientStream
}

type rSSSubscribeAndRefreshClient struct {
	grpc.ClientStream
}

func (x *rSSSubscribeAndRefreshClient) Send(m *Feed) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rSSSubscribeAndRefreshClient) Recv() (*Story, error) {
	m := new(Story)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rSSClient) DeleteFeeds(ctx context.Context, opts ...grpc.CallOption) (RSS_DeleteFeedsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RSS_serviceDesc.Streams[5], "/feed.RSS/DeleteFeeds", opts...)
	if err != nil {
		return nil, err
	}
	x := &rSSDeleteFeedsClient{stream}
	return x, nil
}

type RSS_DeleteFeedsClient interface {
	Send(*Feed) error
	CloseAndRecv() (*Status, error)
	grpc.ClientStream
}

type rSSDeleteFeedsClient struct {
	grpc.ClientStream
}

func (x *rSSDeleteFeedsClient) Send(m *Feed) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rSSDeleteFeedsClient) CloseAndRecv() (*Status, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Status)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RSSServer is the server API for RSS service.
type RSSServer interface {
	GetAllFeedNames(*OK, RSS_GetAllFeedNamesServer) error
	GetFeed(*Feed, RSS_GetFeedServer) error
	GetFeedGroup(*FeedGroup, RSS_GetFeedGroupServer) error
	AddFeed(context.Context, *Feed) (*Status, error)
	AddFeedGroup(context.Context, *FeedGroup) (*Status, error)
	SendFeeds(RSS_SendFeedsServer) error
	SubscribeAndRefresh(RSS_SubscribeAndRefreshServer) error
	DeleteFeeds(RSS_DeleteFeedsServer) error
}

// UnimplementedRSSServer can be embedded to have forward compatible implementations.
type UnimplementedRSSServer struct {
}

func (*UnimplementedRSSServer) GetAllFeedNames(req *OK, srv RSS_GetAllFeedNamesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllFeedNames not implemented")
}
func (*UnimplementedRSSServer) GetFeed(req *Feed, srv RSS_GetFeedServer) error {
	return status.Errorf(codes.Unimplemented, "method GetFeed not implemented")
}
func (*UnimplementedRSSServer) GetFeedGroup(req *FeedGroup, srv RSS_GetFeedGroupServer) error {
	return status.Errorf(codes.Unimplemented, "method GetFeedGroup not implemented")
}
func (*UnimplementedRSSServer) AddFeed(ctx context.Context, req *Feed) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFeed not implemented")
}
func (*UnimplementedRSSServer) AddFeedGroup(ctx context.Context, req *FeedGroup) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFeedGroup not implemented")
}
func (*UnimplementedRSSServer) SendFeeds(srv RSS_SendFeedsServer) error {
	return status.Errorf(codes.Unimplemented, "method SendFeeds not implemented")
}
func (*UnimplementedRSSServer) SubscribeAndRefresh(srv RSS_SubscribeAndRefreshServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeAndRefresh not implemented")
}
func (*UnimplementedRSSServer) DeleteFeeds(srv RSS_DeleteFeedsServer) error {
	return status.Errorf(codes.Unimplemented, "method DeleteFeeds not implemented")
}

func RegisterRSSServer(s *grpc.Server, srv RSSServer) {
	s.RegisterService(&_RSS_serviceDesc, srv)
}

func _RSS_GetAllFeedNames_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(OK)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RSSServer).GetAllFeedNames(m, &rSSGetAllFeedNamesServer{stream})
}

type RSS_GetAllFeedNamesServer interface {
	Send(*Feed) error
	grpc.ServerStream
}

type rSSGetAllFeedNamesServer struct {
	grpc.ServerStream
}

func (x *rSSGetAllFeedNamesServer) Send(m *Feed) error {
	return x.ServerStream.SendMsg(m)
}

func _RSS_GetFeed_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Feed)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RSSServer).GetFeed(m, &rSSGetFeedServer{stream})
}

type RSS_GetFeedServer interface {
	Send(*Story) error
	grpc.ServerStream
}

type rSSGetFeedServer struct {
	grpc.ServerStream
}

func (x *rSSGetFeedServer) Send(m *Story) error {
	return x.ServerStream.SendMsg(m)
}

func _RSS_GetFeedGroup_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FeedGroup)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RSSServer).GetFeedGroup(m, &rSSGetFeedGroupServer{stream})
}

type RSS_GetFeedGroupServer interface {
	Send(*Story) error
	grpc.ServerStream
}

type rSSGetFeedGroupServer struct {
	grpc.ServerStream
}

func (x *rSSGetFeedGroupServer) Send(m *Story) error {
	return x.ServerStream.SendMsg(m)
}

func _RSS_AddFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Feed)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RSSServer).AddFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feed.RSS/AddFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RSSServer).AddFeed(ctx, req.(*Feed))
	}
	return interceptor(ctx, in, info, handler)
}

func _RSS_AddFeedGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RSSServer).AddFeedGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feed.RSS/AddFeedGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RSSServer).AddFeedGroup(ctx, req.(*FeedGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _RSS_SendFeeds_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RSSServer).SendFeeds(&rSSSendFeedsServer{stream})
}

type RSS_SendFeedsServer interface {
	SendAndClose(*Status) error
	Recv() (*Feed, error)
	grpc.ServerStream
}

type rSSSendFeedsServer struct {
	grpc.ServerStream
}

func (x *rSSSendFeedsServer) SendAndClose(m *Status) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rSSSendFeedsServer) Recv() (*Feed, error) {
	m := new(Feed)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RSS_SubscribeAndRefresh_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RSSServer).SubscribeAndRefresh(&rSSSubscribeAndRefreshServer{stream})
}

type RSS_SubscribeAndRefreshServer interface {
	Send(*Story) error
	Recv() (*Feed, error)
	grpc.ServerStream
}

type rSSSubscribeAndRefreshServer struct {
	grpc.ServerStream
}

func (x *rSSSubscribeAndRefreshServer) Send(m *Story) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rSSSubscribeAndRefreshServer) Recv() (*Feed, error) {
	m := new(Feed)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RSS_DeleteFeeds_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RSSServer).DeleteFeeds(&rSSDeleteFeedsServer{stream})
}

type RSS_DeleteFeedsServer interface {
	SendAndClose(*Status) error
	Recv() (*Feed, error)
	grpc.ServerStream
}

type rSSDeleteFeedsServer struct {
	grpc.ServerStream
}

func (x *rSSDeleteFeedsServer) SendAndClose(m *Status) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rSSDeleteFeedsServer) Recv() (*Feed, error) {
	m := new(Feed)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _RSS_serviceDesc = grpc.ServiceDesc{
	ServiceName: "feed.RSS",
	HandlerType: (*RSSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFeed",
			Handler:    _RSS_AddFeed_Handler,
		},
		{
			MethodName: "AddFeedGroup",
			Handler:    _RSS_AddFeedGroup_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllFeedNames",
			Handler:       _RSS_GetAllFeedNames_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetFeed",
			Handler:       _RSS_GetFeed_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetFeedGroup",
			Handler:       _RSS_GetFeedGroup_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendFeeds",
			Handler:       _RSS_SendFeeds_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SubscribeAndRefresh",
			Handler:       _RSS_SubscribeAndRefresh_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeleteFeeds",
			Handler:       _RSS_DeleteFeeds_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "feed.proto",
}
