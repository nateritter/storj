// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pointerdb.proto

package pb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type RedundancyScheme_SchemeType int32

const (
	RedundancyScheme_INVALID RedundancyScheme_SchemeType = 0
	RedundancyScheme_RS      RedundancyScheme_SchemeType = 1
)

var RedundancyScheme_SchemeType_name = map[int32]string{
	0: "INVALID",
	1: "RS",
}

var RedundancyScheme_SchemeType_value = map[string]int32{
	"INVALID": 0,
	"RS":      1,
}

func (x RedundancyScheme_SchemeType) String() string {
	return proto.EnumName(RedundancyScheme_SchemeType_name, int32(x))
}

func (RedundancyScheme_SchemeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_75fef806d28fc810, []int{0, 0}
}

type Pointer_DataType int32

const (
	Pointer_INLINE Pointer_DataType = 0
	Pointer_REMOTE Pointer_DataType = 1
)

var Pointer_DataType_name = map[int32]string{
	0: "INLINE",
	1: "REMOTE",
}

var Pointer_DataType_value = map[string]int32{
	"INLINE": 0,
	"REMOTE": 1,
}

func (x Pointer_DataType) String() string {
	return proto.EnumName(Pointer_DataType_name, int32(x))
}

func (Pointer_DataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_75fef806d28fc810, []int{3, 0}
}

type RedundancyScheme struct {
	Type RedundancyScheme_SchemeType `protobuf:"varint,1,opt,name=type,proto3,enum=pointerdb.RedundancyScheme_SchemeType" json:"type,omitempty"`
	// these values apply to RS encoding
	MinReq               int32    `protobuf:"varint,2,opt,name=min_req,json=minReq,proto3" json:"min_req,omitempty"`
	Total                int32    `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	RepairThreshold      int32    `protobuf:"varint,4,opt,name=repair_threshold,json=repairThreshold,proto3" json:"repair_threshold,omitempty"`
	SuccessThreshold     int32    `protobuf:"varint,5,opt,name=success_threshold,json=successThreshold,proto3" json:"success_threshold,omitempty"`
	ErasureShareSize     int32    `protobuf:"varint,6,opt,name=erasure_share_size,json=erasureShareSize,proto3" json:"erasure_share_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedundancyScheme) Reset()         { *m = RedundancyScheme{} }
func (m *RedundancyScheme) String() string { return proto.CompactTextString(m) }
func (*RedundancyScheme) ProtoMessage()    {}
func (*RedundancyScheme) Descriptor() ([]byte, []int) {
	return fileDescriptor_75fef806d28fc810, []int{0}
}
func (m *RedundancyScheme) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedundancyScheme.Unmarshal(m, b)
}
func (m *RedundancyScheme) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedundancyScheme.Marshal(b, m, deterministic)
}
func (m *RedundancyScheme) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedundancyScheme.Merge(m, src)
}
func (m *RedundancyScheme) XXX_Size() int {
	return xxx_messageInfo_RedundancyScheme.Size(m)
}
func (m *RedundancyScheme) XXX_DiscardUnknown() {
	xxx_messageInfo_RedundancyScheme.DiscardUnknown(m)
}

var xxx_messageInfo_RedundancyScheme proto.InternalMessageInfo

func (m *RedundancyScheme) GetType() RedundancyScheme_SchemeType {
	if m != nil {
		return m.Type
	}
	return RedundancyScheme_INVALID
}

func (m *RedundancyScheme) GetMinReq() int32 {
	if m != nil {
		return m.MinReq
	}
	return 0
}

func (m *RedundancyScheme) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *RedundancyScheme) GetRepairThreshold() int32 {
	if m != nil {
		return m.RepairThreshold
	}
	return 0
}

func (m *RedundancyScheme) GetSuccessThreshold() int32 {
	if m != nil {
		return m.SuccessThreshold
	}
	return 0
}

func (m *RedundancyScheme) GetErasureShareSize() int32 {
	if m != nil {
		return m.ErasureShareSize
	}
	return 0
}

type RemotePiece struct {
	PieceNum             int32      `protobuf:"varint,1,opt,name=piece_num,json=pieceNum,proto3" json:"piece_num,omitempty"`
	NodeId               NodeID     `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3,customtype=NodeID" json:"node_id"`
	Hash                 *PieceHash `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RemotePiece) Reset()         { *m = RemotePiece{} }
func (m *RemotePiece) String() string { return proto.CompactTextString(m) }
func (*RemotePiece) ProtoMessage()    {}
func (*RemotePiece) Descriptor() ([]byte, []int) {
	return fileDescriptor_75fef806d28fc810, []int{1}
}
func (m *RemotePiece) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemotePiece.Unmarshal(m, b)
}
func (m *RemotePiece) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemotePiece.Marshal(b, m, deterministic)
}
func (m *RemotePiece) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemotePiece.Merge(m, src)
}
func (m *RemotePiece) XXX_Size() int {
	return xxx_messageInfo_RemotePiece.Size(m)
}
func (m *RemotePiece) XXX_DiscardUnknown() {
	xxx_messageInfo_RemotePiece.DiscardUnknown(m)
}

var xxx_messageInfo_RemotePiece proto.InternalMessageInfo

func (m *RemotePiece) GetPieceNum() int32 {
	if m != nil {
		return m.PieceNum
	}
	return 0
}

func (m *RemotePiece) GetHash() *PieceHash {
	if m != nil {
		return m.Hash
	}
	return nil
}

type RemoteSegment struct {
	Redundancy           *RedundancyScheme `protobuf:"bytes,1,opt,name=redundancy,proto3" json:"redundancy,omitempty"`
	RootPieceId          PieceID           `protobuf:"bytes,2,opt,name=root_piece_id,json=rootPieceId,proto3,customtype=PieceID" json:"root_piece_id"`
	RemotePieces         []*RemotePiece    `protobuf:"bytes,3,rep,name=remote_pieces,json=remotePieces,proto3" json:"remote_pieces,omitempty"`
	MerkleRoot           []byte            `protobuf:"bytes,4,opt,name=merkle_root,json=merkleRoot,proto3" json:"merkle_root,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RemoteSegment) Reset()         { *m = RemoteSegment{} }
func (m *RemoteSegment) String() string { return proto.CompactTextString(m) }
func (*RemoteSegment) ProtoMessage()    {}
func (*RemoteSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_75fef806d28fc810, []int{2}
}
func (m *RemoteSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteSegment.Unmarshal(m, b)
}
func (m *RemoteSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteSegment.Marshal(b, m, deterministic)
}
func (m *RemoteSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteSegment.Merge(m, src)
}
func (m *RemoteSegment) XXX_Size() int {
	return xxx_messageInfo_RemoteSegment.Size(m)
}
func (m *RemoteSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteSegment.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteSegment proto.InternalMessageInfo

func (m *RemoteSegment) GetRedundancy() *RedundancyScheme {
	if m != nil {
		return m.Redundancy
	}
	return nil
}

func (m *RemoteSegment) GetRemotePieces() []*RemotePiece {
	if m != nil {
		return m.RemotePieces
	}
	return nil
}

func (m *RemoteSegment) GetMerkleRoot() []byte {
	if m != nil {
		return m.MerkleRoot
	}
	return nil
}

type Pointer struct {
	Type                 Pointer_DataType `protobuf:"varint,1,opt,name=type,proto3,enum=pointerdb.Pointer_DataType" json:"type,omitempty"`
	InlineSegment        []byte           `protobuf:"bytes,3,opt,name=inline_segment,json=inlineSegment,proto3" json:"inline_segment,omitempty"`
	Remote               *RemoteSegment   `protobuf:"bytes,4,opt,name=remote,proto3" json:"remote,omitempty"`
	SegmentSize          int64            `protobuf:"varint,5,opt,name=segment_size,json=segmentSize,proto3" json:"segment_size,omitempty"`
	CreationDate         time.Time        `protobuf:"bytes,6,opt,name=creation_date,json=creationDate,proto3,stdtime" json:"creation_date"`
	ExpirationDate       time.Time        `protobuf:"bytes,7,opt,name=expiration_date,json=expirationDate,proto3,stdtime" json:"expiration_date"`
	Metadata             []byte           `protobuf:"bytes,8,opt,name=metadata,proto3" json:"metadata,omitempty"`
	LastRepaired         time.Time        `protobuf:"bytes,9,opt,name=last_repaired,json=lastRepaired,proto3,stdtime" json:"last_repaired"`
	RepairCount          int32            `protobuf:"varint,10,opt,name=repair_count,json=repairCount,proto3" json:"repair_count,omitempty"`
	PieceHashesVerified  bool             `protobuf:"varint,11,opt,name=piece_hashes_verified,json=pieceHashesVerified,proto3" json:"piece_hashes_verified,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Pointer) Reset()         { *m = Pointer{} }
func (m *Pointer) String() string { return proto.CompactTextString(m) }
func (*Pointer) ProtoMessage()    {}
func (*Pointer) Descriptor() ([]byte, []int) {
	return fileDescriptor_75fef806d28fc810, []int{3}
}
func (m *Pointer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pointer.Unmarshal(m, b)
}
func (m *Pointer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pointer.Marshal(b, m, deterministic)
}
func (m *Pointer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pointer.Merge(m, src)
}
func (m *Pointer) XXX_Size() int {
	return xxx_messageInfo_Pointer.Size(m)
}
func (m *Pointer) XXX_DiscardUnknown() {
	xxx_messageInfo_Pointer.DiscardUnknown(m)
}

var xxx_messageInfo_Pointer proto.InternalMessageInfo

func (m *Pointer) GetType() Pointer_DataType {
	if m != nil {
		return m.Type
	}
	return Pointer_INLINE
}

func (m *Pointer) GetInlineSegment() []byte {
	if m != nil {
		return m.InlineSegment
	}
	return nil
}

func (m *Pointer) GetRemote() *RemoteSegment {
	if m != nil {
		return m.Remote
	}
	return nil
}

func (m *Pointer) GetSegmentSize() int64 {
	if m != nil {
		return m.SegmentSize
	}
	return 0
}

func (m *Pointer) GetCreationDate() time.Time {
	if m != nil {
		return m.CreationDate
	}
	return time.Time{}
}

func (m *Pointer) GetExpirationDate() time.Time {
	if m != nil {
		return m.ExpirationDate
	}
	return time.Time{}
}

func (m *Pointer) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Pointer) GetLastRepaired() time.Time {
	if m != nil {
		return m.LastRepaired
	}
	return time.Time{}
}

func (m *Pointer) GetRepairCount() int32 {
	if m != nil {
		return m.RepairCount
	}
	return 0
}

func (m *Pointer) GetPieceHashesVerified() bool {
	if m != nil {
		return m.PieceHashesVerified
	}
	return false
}

// ListResponse is a response message for the List rpc call
type ListResponse struct {
	Items                []*ListResponse_Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	More                 bool                 `protobuf:"varint,2,opt,name=more,proto3" json:"more,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_75fef806d28fc810, []int{4}
}
func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetItems() []*ListResponse_Item {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ListResponse) GetMore() bool {
	if m != nil {
		return m.More
	}
	return false
}

type ListResponse_Item struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Pointer              *Pointer `protobuf:"bytes,2,opt,name=pointer,proto3" json:"pointer,omitempty"`
	IsPrefix             bool     `protobuf:"varint,3,opt,name=is_prefix,json=isPrefix,proto3" json:"is_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse_Item) Reset()         { *m = ListResponse_Item{} }
func (m *ListResponse_Item) String() string { return proto.CompactTextString(m) }
func (*ListResponse_Item) ProtoMessage()    {}
func (*ListResponse_Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_75fef806d28fc810, []int{4, 0}
}
func (m *ListResponse_Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse_Item.Unmarshal(m, b)
}
func (m *ListResponse_Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse_Item.Marshal(b, m, deterministic)
}
func (m *ListResponse_Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse_Item.Merge(m, src)
}
func (m *ListResponse_Item) XXX_Size() int {
	return xxx_messageInfo_ListResponse_Item.Size(m)
}
func (m *ListResponse_Item) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse_Item.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse_Item proto.InternalMessageInfo

func (m *ListResponse_Item) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ListResponse_Item) GetPointer() *Pointer {
	if m != nil {
		return m.Pointer
	}
	return nil
}

func (m *ListResponse_Item) GetIsPrefix() bool {
	if m != nil {
		return m.IsPrefix
	}
	return false
}

func init() {
	proto.RegisterEnum("pointerdb.RedundancyScheme_SchemeType", RedundancyScheme_SchemeType_name, RedundancyScheme_SchemeType_value)
	proto.RegisterEnum("pointerdb.Pointer_DataType", Pointer_DataType_name, Pointer_DataType_value)
	proto.RegisterType((*RedundancyScheme)(nil), "pointerdb.RedundancyScheme")
	proto.RegisterType((*RemotePiece)(nil), "pointerdb.RemotePiece")
	proto.RegisterType((*RemoteSegment)(nil), "pointerdb.RemoteSegment")
	proto.RegisterType((*Pointer)(nil), "pointerdb.Pointer")
	proto.RegisterType((*ListResponse)(nil), "pointerdb.ListResponse")
	proto.RegisterType((*ListResponse_Item)(nil), "pointerdb.ListResponse.Item")
}

func init() { proto.RegisterFile("pointerdb.proto", fileDescriptor_75fef806d28fc810) }

var fileDescriptor_75fef806d28fc810 = []byte{
	// 797 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x36, 0x6d, 0xfd, 0xd0, 0x43, 0xc9, 0x56, 0xb6, 0x7f, 0x84, 0x52, 0x40, 0x0e, 0x81, 0xb4,
	0x2e, 0x1a, 0xd0, 0x05, 0x73, 0x6b, 0x4e, 0x75, 0x65, 0xa0, 0x04, 0x1c, 0xd5, 0x58, 0x19, 0x39,
	0xf4, 0x42, 0xac, 0xc4, 0xb1, 0xb8, 0xa8, 0xc8, 0x65, 0x76, 0x57, 0x45, 0xec, 0xa7, 0xe8, 0x53,
	0xf4, 0x01, 0x7a, 0xef, 0xbd, 0xcf, 0xd0, 0x43, 0xf2, 0x2a, 0xc5, 0xee, 0x92, 0x92, 0xd2, 0x00,
	0x05, 0x72, 0xb1, 0x77, 0x66, 0xbe, 0xf9, 0xe1, 0x37, 0xdf, 0x08, 0x4e, 0x6b, 0xc1, 0x2b, 0x8d,
	0x32, 0x5f, 0xc4, 0xb5, 0x14, 0x5a, 0x90, 0xe3, 0xad, 0x63, 0x3c, 0x59, 0x09, 0xb1, 0x5a, 0xe3,
	0x85, 0x0d, 0x2c, 0x36, 0x77, 0x17, 0x9a, 0x97, 0xa8, 0x34, 0x2b, 0x6b, 0x87, 0x1d, 0xc3, 0x4a,
	0xac, 0x44, 0xf3, 0x1e, 0x08, 0x99, 0xa3, 0x54, 0xce, 0x8a, 0xfe, 0x38, 0x84, 0x11, 0xc5, 0x7c,
	0x53, 0xe5, 0xac, 0x5a, 0xde, 0xcf, 0x97, 0x05, 0x96, 0x48, 0xbe, 0x87, 0x8e, 0xbe, 0xaf, 0x31,
	0xf4, 0xce, 0xbc, 0xf3, 0x93, 0xe4, 0xab, 0x78, 0xd7, 0xfa, 0xbf, 0xd0, 0xd8, 0xfd, 0xbb, 0xbd,
	0xaf, 0x91, 0xda, 0x1c, 0xf2, 0x05, 0xf4, 0x4b, 0x5e, 0x65, 0x12, 0x5f, 0x87, 0x87, 0x67, 0xde,
	0x79, 0x97, 0xf6, 0x4a, 0x5e, 0x51, 0x7c, 0x4d, 0x3e, 0x85, 0xae, 0x16, 0x9a, 0xad, 0xc3, 0x23,
	0xeb, 0x76, 0x06, 0xf9, 0x06, 0x46, 0x12, 0x6b, 0xc6, 0x65, 0xa6, 0x0b, 0x89, 0xaa, 0x10, 0xeb,
	0x3c, 0xec, 0x58, 0xc0, 0xa9, 0xf3, 0xdf, 0xb6, 0x6e, 0xf2, 0x2d, 0x3c, 0x52, 0x9b, 0xe5, 0x12,
	0x95, 0xda, 0xc3, 0x76, 0x2d, 0x76, 0xd4, 0x04, 0x76, 0xe0, 0x67, 0x40, 0x50, 0x32, 0xb5, 0x91,
	0x98, 0xa9, 0x82, 0x99, 0xbf, 0xfc, 0x01, 0xc3, 0x9e, 0x43, 0x37, 0x91, 0xb9, 0x09, 0xcc, 0xf9,
	0x03, 0x46, 0x4f, 0x00, 0x76, 0x1f, 0x42, 0x02, 0xe8, 0xa7, 0xb3, 0x57, 0x3f, 0x5c, 0xa7, 0xd3,
	0xd1, 0x01, 0xe9, 0xc1, 0x21, 0x9d, 0x8f, 0xbc, 0xe8, 0x01, 0x02, 0x8a, 0xa5, 0xd0, 0x78, 0xc3,
	0x71, 0x89, 0xe4, 0x31, 0x1c, 0xd7, 0xe6, 0x91, 0x55, 0x9b, 0xd2, 0xf2, 0xd4, 0xa5, 0xbe, 0x75,
	0xcc, 0x36, 0x25, 0xf9, 0x1a, 0xfa, 0x95, 0xc8, 0x31, 0xe3, 0xb9, 0xe5, 0x60, 0x70, 0x79, 0xf2,
	0xf7, 0xdb, 0xc9, 0xc1, 0x3f, 0x6f, 0x27, 0xbd, 0x99, 0xc8, 0x31, 0x9d, 0xd2, 0x9e, 0x09, 0xa7,
	0x39, 0x79, 0x0a, 0x9d, 0x82, 0xa9, 0xc2, 0x52, 0x12, 0x24, 0x8f, 0xe2, 0x66, 0x35, 0xb6, 0xc5,
	0x4f, 0x4c, 0x15, 0xd4, 0x86, 0xa3, 0x77, 0x1e, 0x0c, 0x5d, 0xf3, 0x39, 0xae, 0x4a, 0xac, 0x34,
	0x79, 0x01, 0x20, 0xb7, 0xab, 0xb0, 0xfd, 0x83, 0xe4, 0xf1, 0xff, 0xec, 0x89, 0xee, 0xc1, 0xc9,
	0x73, 0x18, 0x4a, 0x21, 0x74, 0xe6, 0x3e, 0x60, 0x3b, 0xe4, 0x69, 0x33, 0x64, 0xdf, 0xb6, 0x4f,
	0xa7, 0x34, 0x30, 0x28, 0x67, 0xe4, 0xe4, 0x05, 0x0c, 0xa5, 0x1d, 0xc1, 0xa5, 0xa9, 0xf0, 0xe8,
	0xec, 0xe8, 0x3c, 0x48, 0x3e, 0x7f, 0xaf, 0xe9, 0x96, 0x1f, 0x3a, 0x90, 0x3b, 0x43, 0x91, 0x09,
	0x04, 0x25, 0xca, 0x5f, 0xd7, 0x98, 0x99, 0x92, 0x76, 0xc1, 0x03, 0x0a, 0xce, 0x45, 0x85, 0xd0,
	0xd1, 0x9f, 0x1d, 0xe8, 0xdf, 0xb8, 0x42, 0xe4, 0xe2, 0x3d, 0xf5, 0xed, 0x7f, 0x55, 0x83, 0x88,
	0xa7, 0x4c, 0xb3, 0x3d, 0xc9, 0x3d, 0x85, 0x13, 0x5e, 0xad, 0x79, 0x85, 0x99, 0x72, 0xf4, 0x58,
	0x3e, 0x07, 0x74, 0xe8, 0xbc, 0x2d, 0x67, 0xdf, 0x41, 0xcf, 0x0d, 0x65, 0xfb, 0x07, 0x49, 0xf8,
	0xc1, 0xe8, 0x0d, 0x92, 0x36, 0x38, 0xf2, 0x04, 0x06, 0x4d, 0x45, 0x27, 0x1f, 0x23, 0xb6, 0x23,
	0x1a, 0x34, 0x3e, 0xa3, 0x1c, 0x92, 0xc2, 0x70, 0x29, 0x91, 0x69, 0x2e, 0xaa, 0x2c, 0x67, 0xda,
	0x49, 0x2c, 0x48, 0xc6, 0xb1, 0x3b, 0xc9, 0xb8, 0x3d, 0xc9, 0xf8, 0xb6, 0x3d, 0xc9, 0x4b, 0xdf,
	0xf0, 0xfc, 0xfb, 0xbb, 0x89, 0x47, 0x07, 0x6d, 0xea, 0x94, 0x69, 0x24, 0x2f, 0xe1, 0x14, 0xdf,
	0xd4, 0x5c, 0xee, 0x15, 0xeb, 0x7f, 0x44, 0xb1, 0x93, 0x5d, 0xb2, 0x2d, 0x37, 0x06, 0xbf, 0x44,
	0xcd, 0x72, 0xa6, 0x59, 0xe8, 0x5b, 0x3e, 0xb6, 0xb6, 0x99, 0x7a, 0xcd, 0x94, 0xce, 0xdc, 0x89,
	0x61, 0x1e, 0x1e, 0x7f, 0xcc, 0xd4, 0x26, 0x95, 0x36, 0x99, 0x86, 0xa3, 0xe6, 0x80, 0x97, 0x62,
	0x53, 0xe9, 0x10, 0xec, 0x2d, 0x04, 0xce, 0xf7, 0xa3, 0x71, 0x91, 0x04, 0x3e, 0x73, 0x52, 0x33,
	0x62, 0x46, 0x95, 0xfd, 0x86, 0x92, 0xdf, 0x71, 0xcc, 0xc3, 0xe0, 0xcc, 0x3b, 0xf7, 0xe9, 0x27,
	0x75, 0x2b, 0x77, 0x54, 0xaf, 0x9a, 0x50, 0x14, 0x81, 0xdf, 0x6e, 0x99, 0x00, 0xf4, 0xd2, 0xd9,
	0x75, 0x3a, 0xbb, 0x1a, 0x1d, 0x98, 0x37, 0xbd, 0x7a, 0xf9, 0xf3, 0xed, 0xd5, 0xc8, 0x8b, 0xfe,
	0xf2, 0x60, 0x70, 0xcd, 0xcd, 0x2c, 0xaa, 0x16, 0x95, 0x42, 0x92, 0x40, 0x97, 0x6b, 0x2c, 0x55,
	0xe8, 0x59, 0x6d, 0x7e, 0xb9, 0xb7, 0xe0, 0x7d, 0x5c, 0x9c, 0x6a, 0x2c, 0xa9, 0x83, 0x12, 0x02,
	0x9d, 0x52, 0x48, 0xb4, 0x37, 0xe0, 0x53, 0xfb, 0x1e, 0x23, 0x74, 0x0c, 0xc4, 0xc4, 0x6a, 0xa6,
	0x0b, 0xab, 0xc4, 0x63, 0x6a, 0xdf, 0xe4, 0x19, 0xf4, 0x9b, 0xaa, 0x36, 0x25, 0x48, 0xc8, 0x87,
	0x02, 0xa5, 0x2d, 0xc4, 0xfc, 0x4c, 0x70, 0x95, 0xd5, 0x12, 0xef, 0xf8, 0x1b, 0xab, 0x4a, 0x9f,
	0xfa, 0x5c, 0xdd, 0x58, 0xfb, 0xb2, 0xf3, 0xcb, 0x61, 0xbd, 0x58, 0xf4, 0x2c, 0xd9, 0xcf, 0xff,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0x76, 0xa2, 0x36, 0x45, 0xe1, 0x05, 0x00, 0x00,
}
