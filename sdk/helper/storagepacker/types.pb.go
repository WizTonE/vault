// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sdk/helper/storagepacker/types.proto

package storagepacker

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

// Item represents an entry that gets inserted into the storage packer
type Item struct {
	// ID must be provided by the caller; the same value, if used with GetItem,
	// can be used to fetch the item. However, when iterating through a bucket,
	// this ID will be an internal ID. In other words, outside of the use-case
	// described above, the caller *must not* rely on this value to be
	// consistent with what they passed in.
	ID string `sentinel:"" protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Message holds the contents of the item
	// Deprecated: Use 'Data' instead
	Message *any.Any `sentinel:"" protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// Data holds the contents of the item. Used in storage packer v2.
	Data                 []byte   `sentinel:"" protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_02e03d2111f28219, []int{0}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Item) GetMessage() *any.Any {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Item) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// Bucket is a construct to hold multiple items within itself. This
// abstraction contains multiple buckets of the same kind within itself and
// shares amont them the items that get inserted. When the bucket as a whole
// gets too big to hold more items, the contained buckets gets pushed out only
// to become independent buckets. Hence, this can grow infinitely in terms of
// storage space for items that get inserted.
type Bucket struct {
	// Key is the storage path where the bucket gets stored
	Key string `sentinel:"" protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Items holds the items contained within this bucket. Used by v1.
	Items []*Item `sentinel:"" protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	// ItemMap stores a mapping of item ID to message. Used by v2.
	ItemMap map[string][]byte `sentinel:"" protobuf:"bytes,3,rep,name=item_map,json=itemMap,proto3" json:"item_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Has this bucket been emptied into sub-buckets? Used by v2.
	// If has_shards is true then item_map should be empty.
	HasShards            bool     `sentinel:"" protobuf:"varint,4,opt,name=has_shards,json=hasShards,proto3" json:"has_shards,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bucket) Reset()         { *m = Bucket{} }
func (m *Bucket) String() string { return proto.CompactTextString(m) }
func (*Bucket) ProtoMessage()    {}
func (*Bucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_02e03d2111f28219, []int{1}
}

func (m *Bucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bucket.Unmarshal(m, b)
}
func (m *Bucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bucket.Marshal(b, m, deterministic)
}
func (m *Bucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bucket.Merge(m, src)
}
func (m *Bucket) XXX_Size() int {
	return xxx_messageInfo_Bucket.Size(m)
}
func (m *Bucket) XXX_DiscardUnknown() {
	xxx_messageInfo_Bucket.DiscardUnknown(m)
}

var xxx_messageInfo_Bucket proto.InternalMessageInfo

func (m *Bucket) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Bucket) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *Bucket) GetItemMap() map[string][]byte {
	if m != nil {
		return m.ItemMap
	}
	return nil
}

func (m *Bucket) GetHasShards() bool {
	if m != nil {
		return m.HasShards
	}
	return false
}

func init() {
	proto.RegisterType((*Item)(nil), "storagepacker.Item")
	proto.RegisterType((*Bucket)(nil), "storagepacker.Bucket")
	proto.RegisterMapType((map[string][]byte)(nil), "storagepacker.Bucket.ItemMapEntry")
}

func init() {
	proto.RegisterFile("sdk/helper/storagepacker/types.proto", fileDescriptor_02e03d2111f28219)
}

var fileDescriptor_02e03d2111f28219 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x3f, 0x6b, 0xeb, 0x30,
	0x14, 0xc5, 0x91, 0x9d, 0xbf, 0x4a, 0xde, 0xe3, 0xa1, 0x97, 0xc1, 0x0d, 0x14, 0x4c, 0xe8, 0xe0,
	0x2e, 0x12, 0x24, 0x14, 0x4a, 0xa0, 0x43, 0x03, 0x1d, 0x3a, 0x74, 0x51, 0xb7, 0x2c, 0xe1, 0xc6,
	0x56, 0x2d, 0x63, 0x3b, 0x12, 0x92, 0x1c, 0xf0, 0xa7, 0xed, 0x57, 0x29, 0xb1, 0x1a, 0x68, 0x4a,
	0xbb, 0x9d, 0xab, 0xfb, 0xbb, 0x47, 0x87, 0x7b, 0xf1, 0x8d, 0xcd, 0x4a, 0x26, 0x45, 0xa5, 0x85,
	0x61, 0xd6, 0x29, 0x03, 0xb9, 0xd0, 0x90, 0x96, 0xc2, 0x30, 0xd7, 0x6a, 0x61, 0xa9, 0x36, 0xca,
	0x29, 0xf2, 0xe7, 0xa2, 0x35, 0xbf, 0xca, 0x95, 0xca, 0x2b, 0xc1, 0xba, 0xe6, 0xbe, 0x79, 0x63,
	0x70, 0x68, 0x3d, 0xb9, 0xd8, 0xe2, 0xde, 0xb3, 0x13, 0x35, 0xf9, 0x8b, 0x83, 0x22, 0x8b, 0x50,
	0x8c, 0x92, 0x31, 0x0f, 0x8a, 0x8c, 0x50, 0x3c, 0xac, 0x85, 0xb5, 0x90, 0x8b, 0x28, 0x88, 0x51,
	0x32, 0x59, 0xce, 0xa8, 0x37, 0xa1, 0x67, 0x13, 0xfa, 0x78, 0x68, 0xf9, 0x19, 0x22, 0x04, 0xf7,
	0x32, 0x70, 0x10, 0x85, 0x31, 0x4a, 0xa6, 0xbc, 0xd3, 0x8b, 0x77, 0x84, 0x07, 0x9b, 0x26, 0x2d,
	0x85, 0x23, 0xff, 0x70, 0x58, 0x8a, 0xf6, 0xd3, 0xff, 0x24, 0xc9, 0x2d, 0xee, 0x17, 0x4e, 0xd4,
	0x36, 0x0a, 0xe2, 0x30, 0x99, 0x2c, 0xff, 0xd3, 0x8b, 0xc8, 0xf4, 0x14, 0x8a, 0x7b, 0x82, 0x3c,
	0xe0, 0xd1, 0x49, 0xec, 0x6a, 0xd0, 0x51, 0xd8, 0xd1, 0x8b, 0x6f, 0xb4, 0xff, 0xa5, 0x1b, 0x7a,
	0x01, 0xfd, 0x74, 0x70, 0xa6, 0xe5, 0xc3, 0xc2, 0x57, 0xe4, 0x1a, 0x63, 0x09, 0x76, 0x67, 0x25,
	0x98, 0xcc, 0x46, 0xbd, 0x18, 0x25, 0x23, 0x3e, 0x96, 0x60, 0x5f, 0xbb, 0x87, 0xf9, 0x1a, 0x4f,
	0xbf, 0xce, 0xfd, 0x10, 0x75, 0x86, 0xfb, 0x47, 0xa8, 0x1a, 0xbf, 0x89, 0x29, 0xf7, 0xc5, 0x3a,
	0xb8, 0x47, 0x9b, 0xbb, 0xed, 0x2a, 0x2f, 0x9c, 0x6c, 0xf6, 0x34, 0x55, 0x35, 0x93, 0x60, 0x65,
	0x91, 0x2a, 0xa3, 0xd9, 0x11, 0x9a, 0xca, 0xb1, 0xdf, 0x4e, 0xb5, 0x1f, 0x74, 0x3b, 0x5c, 0x7d,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x02, 0x3e, 0x9e, 0xa4, 0xcd, 0x01, 0x00, 0x00,
}