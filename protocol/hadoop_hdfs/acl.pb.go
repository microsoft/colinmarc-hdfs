// Code generated by protoc-gen-go. DO NOT EDIT.
// source: acl.proto

package hadoop_hdfs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AclEntryProto_AclEntryScopeProto int32

const (
	AclEntryProto_ACCESS  AclEntryProto_AclEntryScopeProto = 0
	AclEntryProto_DEFAULT AclEntryProto_AclEntryScopeProto = 1
)

var AclEntryProto_AclEntryScopeProto_name = map[int32]string{
	0: "ACCESS",
	1: "DEFAULT",
}
var AclEntryProto_AclEntryScopeProto_value = map[string]int32{
	"ACCESS":  0,
	"DEFAULT": 1,
}

func (x AclEntryProto_AclEntryScopeProto) Enum() *AclEntryProto_AclEntryScopeProto {
	p := new(AclEntryProto_AclEntryScopeProto)
	*p = x
	return p
}
func (x AclEntryProto_AclEntryScopeProto) String() string {
	return proto.EnumName(AclEntryProto_AclEntryScopeProto_name, int32(x))
}
func (x AclEntryProto_AclEntryScopeProto) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *AclEntryProto_AclEntryScopeProto) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AclEntryProto_AclEntryScopeProto_value, data, "AclEntryProto_AclEntryScopeProto")
	if err != nil {
		return err
	}
	*x = AclEntryProto_AclEntryScopeProto(value)
	return nil
}
func (AclEntryProto_AclEntryScopeProto) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor9, []int{0, 0}
}

type AclEntryProto_AclEntryTypeProto int32

const (
	AclEntryProto_USER  AclEntryProto_AclEntryTypeProto = 0
	AclEntryProto_GROUP AclEntryProto_AclEntryTypeProto = 1
	AclEntryProto_MASK  AclEntryProto_AclEntryTypeProto = 2
	AclEntryProto_OTHER AclEntryProto_AclEntryTypeProto = 3
)

var AclEntryProto_AclEntryTypeProto_name = map[int32]string{
	0: "USER",
	1: "GROUP",
	2: "MASK",
	3: "OTHER",
}
var AclEntryProto_AclEntryTypeProto_value = map[string]int32{
	"USER":  0,
	"GROUP": 1,
	"MASK":  2,
	"OTHER": 3,
}

func (x AclEntryProto_AclEntryTypeProto) Enum() *AclEntryProto_AclEntryTypeProto {
	p := new(AclEntryProto_AclEntryTypeProto)
	*p = x
	return p
}
func (x AclEntryProto_AclEntryTypeProto) String() string {
	return proto.EnumName(AclEntryProto_AclEntryTypeProto_name, int32(x))
}
func (x AclEntryProto_AclEntryTypeProto) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *AclEntryProto_AclEntryTypeProto) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AclEntryProto_AclEntryTypeProto_value, data, "AclEntryProto_AclEntryTypeProto")
	if err != nil {
		return err
	}
	*x = AclEntryProto_AclEntryTypeProto(value)
	return nil
}
func (AclEntryProto_AclEntryTypeProto) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor9, []int{0, 1}
}

type AclEntryProto_FsActionProto int32

const (
	AclEntryProto_NONE          AclEntryProto_FsActionProto = 0
	AclEntryProto_EXECUTE       AclEntryProto_FsActionProto = 1
	AclEntryProto_WRITE         AclEntryProto_FsActionProto = 2
	AclEntryProto_WRITE_EXECUTE AclEntryProto_FsActionProto = 3
	AclEntryProto_READ          AclEntryProto_FsActionProto = 4
	AclEntryProto_READ_EXECUTE  AclEntryProto_FsActionProto = 5
	AclEntryProto_READ_WRITE    AclEntryProto_FsActionProto = 6
	AclEntryProto_PERM_ALL      AclEntryProto_FsActionProto = 7
)

var AclEntryProto_FsActionProto_name = map[int32]string{
	0: "NONE",
	1: "EXECUTE",
	2: "WRITE",
	3: "WRITE_EXECUTE",
	4: "READ",
	5: "READ_EXECUTE",
	6: "READ_WRITE",
	7: "PERM_ALL",
}
var AclEntryProto_FsActionProto_value = map[string]int32{
	"NONE":          0,
	"EXECUTE":       1,
	"WRITE":         2,
	"WRITE_EXECUTE": 3,
	"READ":          4,
	"READ_EXECUTE":  5,
	"READ_WRITE":    6,
	"PERM_ALL":      7,
}

func (x AclEntryProto_FsActionProto) Enum() *AclEntryProto_FsActionProto {
	p := new(AclEntryProto_FsActionProto)
	*p = x
	return p
}
func (x AclEntryProto_FsActionProto) String() string {
	return proto.EnumName(AclEntryProto_FsActionProto_name, int32(x))
}
func (x AclEntryProto_FsActionProto) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *AclEntryProto_FsActionProto) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AclEntryProto_FsActionProto_value, data, "AclEntryProto_FsActionProto")
	if err != nil {
		return err
	}
	*x = AclEntryProto_FsActionProto(value)
	return nil
}
func (AclEntryProto_FsActionProto) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor9, []int{0, 2}
}

type AclEntryProto struct {
	Type             *AclEntryProto_AclEntryTypeProto  `protobuf:"varint,1,req,name=type,enum=hadoop.hdfs.AclEntryProto_AclEntryTypeProto" json:"type,omitempty"`
	Scope            *AclEntryProto_AclEntryScopeProto `protobuf:"varint,2,req,name=scope,enum=hadoop.hdfs.AclEntryProto_AclEntryScopeProto" json:"scope,omitempty"`
	Permissions      *AclEntryProto_FsActionProto      `protobuf:"varint,3,req,name=permissions,enum=hadoop.hdfs.AclEntryProto_FsActionProto" json:"permissions,omitempty"`
	Name             *string                           `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte                            `json:"-"`
}

func (m *AclEntryProto) Reset()                    { *m = AclEntryProto{} }
func (m *AclEntryProto) String() string            { return proto.CompactTextString(m) }
func (*AclEntryProto) ProtoMessage()               {}
func (*AclEntryProto) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *AclEntryProto) GetType() AclEntryProto_AclEntryTypeProto {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *AclEntryProto) GetScope() AclEntryProto_AclEntryScopeProto {
	if m != nil && m.Scope != nil {
		return *m.Scope
	}
	return 0
}

func (m *AclEntryProto) GetPermissions() AclEntryProto_FsActionProto {
	if m != nil && m.Permissions != nil {
		return *m.Permissions
	}
	return 0
}

func (m *AclEntryProto) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type AclStatusProto struct {
	Owner            *string            `protobuf:"bytes,1,req,name=owner" json:"owner,omitempty"`
	Group            *string            `protobuf:"bytes,2,req,name=group" json:"group,omitempty"`
	Sticky           *bool              `protobuf:"varint,3,req,name=sticky" json:"sticky,omitempty"`
	Entries          []*AclEntryProto   `protobuf:"bytes,4,rep,name=entries" json:"entries,omitempty"`
	Permission       *FsPermissionProto `protobuf:"bytes,5,opt,name=permission" json:"permission,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *AclStatusProto) Reset()                    { *m = AclStatusProto{} }
func (m *AclStatusProto) String() string            { return proto.CompactTextString(m) }
func (*AclStatusProto) ProtoMessage()               {}
func (*AclStatusProto) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

func (m *AclStatusProto) GetOwner() string {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return ""
}

func (m *AclStatusProto) GetGroup() string {
	if m != nil && m.Group != nil {
		return *m.Group
	}
	return ""
}

func (m *AclStatusProto) GetSticky() bool {
	if m != nil && m.Sticky != nil {
		return *m.Sticky
	}
	return false
}

func (m *AclStatusProto) GetEntries() []*AclEntryProto {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *AclStatusProto) GetPermission() *FsPermissionProto {
	if m != nil {
		return m.Permission
	}
	return nil
}

type ModifyAclEntriesRequestProto struct {
	Src              *string          `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	AclSpec          []*AclEntryProto `protobuf:"bytes,2,rep,name=aclSpec" json:"aclSpec,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *ModifyAclEntriesRequestProto) Reset()                    { *m = ModifyAclEntriesRequestProto{} }
func (m *ModifyAclEntriesRequestProto) String() string            { return proto.CompactTextString(m) }
func (*ModifyAclEntriesRequestProto) ProtoMessage()               {}
func (*ModifyAclEntriesRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

func (m *ModifyAclEntriesRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

func (m *ModifyAclEntriesRequestProto) GetAclSpec() []*AclEntryProto {
	if m != nil {
		return m.AclSpec
	}
	return nil
}

type ModifyAclEntriesResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ModifyAclEntriesResponseProto) Reset()                    { *m = ModifyAclEntriesResponseProto{} }
func (m *ModifyAclEntriesResponseProto) String() string            { return proto.CompactTextString(m) }
func (*ModifyAclEntriesResponseProto) ProtoMessage()               {}
func (*ModifyAclEntriesResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{3} }

type RemoveAclRequestProto struct {
	Src              *string `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RemoveAclRequestProto) Reset()                    { *m = RemoveAclRequestProto{} }
func (m *RemoveAclRequestProto) String() string            { return proto.CompactTextString(m) }
func (*RemoveAclRequestProto) ProtoMessage()               {}
func (*RemoveAclRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{4} }

func (m *RemoveAclRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

type RemoveAclResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RemoveAclResponseProto) Reset()                    { *m = RemoveAclResponseProto{} }
func (m *RemoveAclResponseProto) String() string            { return proto.CompactTextString(m) }
func (*RemoveAclResponseProto) ProtoMessage()               {}
func (*RemoveAclResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{5} }

type RemoveAclEntriesRequestProto struct {
	Src              *string          `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	AclSpec          []*AclEntryProto `protobuf:"bytes,2,rep,name=aclSpec" json:"aclSpec,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *RemoveAclEntriesRequestProto) Reset()                    { *m = RemoveAclEntriesRequestProto{} }
func (m *RemoveAclEntriesRequestProto) String() string            { return proto.CompactTextString(m) }
func (*RemoveAclEntriesRequestProto) ProtoMessage()               {}
func (*RemoveAclEntriesRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{6} }

func (m *RemoveAclEntriesRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

func (m *RemoveAclEntriesRequestProto) GetAclSpec() []*AclEntryProto {
	if m != nil {
		return m.AclSpec
	}
	return nil
}

type RemoveAclEntriesResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RemoveAclEntriesResponseProto) Reset()                    { *m = RemoveAclEntriesResponseProto{} }
func (m *RemoveAclEntriesResponseProto) String() string            { return proto.CompactTextString(m) }
func (*RemoveAclEntriesResponseProto) ProtoMessage()               {}
func (*RemoveAclEntriesResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{7} }

type RemoveDefaultAclRequestProto struct {
	Src              *string `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RemoveDefaultAclRequestProto) Reset()                    { *m = RemoveDefaultAclRequestProto{} }
func (m *RemoveDefaultAclRequestProto) String() string            { return proto.CompactTextString(m) }
func (*RemoveDefaultAclRequestProto) ProtoMessage()               {}
func (*RemoveDefaultAclRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{8} }

func (m *RemoveDefaultAclRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

type RemoveDefaultAclResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RemoveDefaultAclResponseProto) Reset()                    { *m = RemoveDefaultAclResponseProto{} }
func (m *RemoveDefaultAclResponseProto) String() string            { return proto.CompactTextString(m) }
func (*RemoveDefaultAclResponseProto) ProtoMessage()               {}
func (*RemoveDefaultAclResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{9} }

type SetAclRequestProto struct {
	Src              *string          `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	AclSpec          []*AclEntryProto `protobuf:"bytes,2,rep,name=aclSpec" json:"aclSpec,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *SetAclRequestProto) Reset()                    { *m = SetAclRequestProto{} }
func (m *SetAclRequestProto) String() string            { return proto.CompactTextString(m) }
func (*SetAclRequestProto) ProtoMessage()               {}
func (*SetAclRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{10} }

func (m *SetAclRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

func (m *SetAclRequestProto) GetAclSpec() []*AclEntryProto {
	if m != nil {
		return m.AclSpec
	}
	return nil
}

type SetAclResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *SetAclResponseProto) Reset()                    { *m = SetAclResponseProto{} }
func (m *SetAclResponseProto) String() string            { return proto.CompactTextString(m) }
func (*SetAclResponseProto) ProtoMessage()               {}
func (*SetAclResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{11} }

type GetAclStatusRequestProto struct {
	Src              *string `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetAclStatusRequestProto) Reset()                    { *m = GetAclStatusRequestProto{} }
func (m *GetAclStatusRequestProto) String() string            { return proto.CompactTextString(m) }
func (*GetAclStatusRequestProto) ProtoMessage()               {}
func (*GetAclStatusRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{12} }

func (m *GetAclStatusRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

type GetAclStatusResponseProto struct {
	Result           *AclStatusProto `protobuf:"bytes,1,req,name=result" json:"result,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *GetAclStatusResponseProto) Reset()                    { *m = GetAclStatusResponseProto{} }
func (m *GetAclStatusResponseProto) String() string            { return proto.CompactTextString(m) }
func (*GetAclStatusResponseProto) ProtoMessage()               {}
func (*GetAclStatusResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{13} }

func (m *GetAclStatusResponseProto) GetResult() *AclStatusProto {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*AclEntryProto)(nil), "hadoop.hdfs.AclEntryProto")
	proto.RegisterType((*AclStatusProto)(nil), "hadoop.hdfs.AclStatusProto")
	proto.RegisterType((*ModifyAclEntriesRequestProto)(nil), "hadoop.hdfs.ModifyAclEntriesRequestProto")
	proto.RegisterType((*ModifyAclEntriesResponseProto)(nil), "hadoop.hdfs.ModifyAclEntriesResponseProto")
	proto.RegisterType((*RemoveAclRequestProto)(nil), "hadoop.hdfs.RemoveAclRequestProto")
	proto.RegisterType((*RemoveAclResponseProto)(nil), "hadoop.hdfs.RemoveAclResponseProto")
	proto.RegisterType((*RemoveAclEntriesRequestProto)(nil), "hadoop.hdfs.RemoveAclEntriesRequestProto")
	proto.RegisterType((*RemoveAclEntriesResponseProto)(nil), "hadoop.hdfs.RemoveAclEntriesResponseProto")
	proto.RegisterType((*RemoveDefaultAclRequestProto)(nil), "hadoop.hdfs.RemoveDefaultAclRequestProto")
	proto.RegisterType((*RemoveDefaultAclResponseProto)(nil), "hadoop.hdfs.RemoveDefaultAclResponseProto")
	proto.RegisterType((*SetAclRequestProto)(nil), "hadoop.hdfs.SetAclRequestProto")
	proto.RegisterType((*SetAclResponseProto)(nil), "hadoop.hdfs.SetAclResponseProto")
	proto.RegisterType((*GetAclStatusRequestProto)(nil), "hadoop.hdfs.GetAclStatusRequestProto")
	proto.RegisterType((*GetAclStatusResponseProto)(nil), "hadoop.hdfs.GetAclStatusResponseProto")
	proto.RegisterEnum("hadoop.hdfs.AclEntryProto_AclEntryScopeProto", AclEntryProto_AclEntryScopeProto_name, AclEntryProto_AclEntryScopeProto_value)
	proto.RegisterEnum("hadoop.hdfs.AclEntryProto_AclEntryTypeProto", AclEntryProto_AclEntryTypeProto_name, AclEntryProto_AclEntryTypeProto_value)
	proto.RegisterEnum("hadoop.hdfs.AclEntryProto_FsActionProto", AclEntryProto_FsActionProto_name, AclEntryProto_FsActionProto_value)
}

func init() { proto.RegisterFile("acl.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 605 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0x63, 0x27, 0xad, 0xc7, 0x6d, 0xb5, 0x5d, 0x68, 0x65, 0x4a, 0x81, 0xc8, 0x12, 0x52,
	0x90, 0xda, 0x08, 0x05, 0x78, 0x04, 0xe1, 0xa6, 0xdb, 0x72, 0xe9, 0x25, 0x5a, 0x27, 0x82, 0x07,
	0xa4, 0xca, 0xda, 0x6e, 0x5a, 0x0b, 0xd7, 0x6b, 0xbc, 0x0e, 0x28, 0x2f, 0x7c, 0x0b, 0xdf, 0xc3,
	0x77, 0xf0, 0x21, 0xc8, 0xbb, 0x4e, 0x70, 0x88, 0x68, 0x10, 0x12, 0x2f, 0xd1, 0x5c, 0xce, 0x39,
	0x73, 0x46, 0x13, 0x2f, 0xd8, 0x21, 0x8b, 0xdb, 0x69, 0x26, 0x72, 0x81, 0x9d, 0xab, 0xf0, 0x42,
	0x88, 0xb4, 0x7d, 0x75, 0x31, 0x94, 0xdb, 0x50, 0xfc, 0xea, 0x86, 0xf7, 0xc3, 0x84, 0x35, 0x9f,
	0xc5, 0x24, 0xc9, 0xb3, 0x71, 0x4f, 0x41, 0x5f, 0x82, 0x95, 0x8f, 0x53, 0xee, 0x1a, 0xcd, 0x5a,
	0x6b, 0xbd, 0xb3, 0xdb, 0xae, 0x30, 0xdb, 0x33, 0xc8, 0x69, 0xd6, 0x1f, 0xa7, 0x5c, 0x55, 0xa8,
	0x62, 0xe2, 0x2e, 0xd4, 0x25, 0x13, 0x29, 0x77, 0x6b, 0x4a, 0x62, 0xef, 0x2f, 0x24, 0x82, 0x02,
	0xaf, 0x35, 0x34, 0x17, 0xbf, 0x01, 0x27, 0xe5, 0xd9, 0x75, 0x24, 0x65, 0x24, 0x12, 0xe9, 0x9a,
	0x4a, 0xaa, 0x75, 0x83, 0xd4, 0xa1, 0xf4, 0x59, 0x1e, 0x89, 0x44, 0xab, 0x54, 0xc9, 0x18, 0x83,
	0x95, 0x84, 0xd7, 0xdc, 0xb5, 0x9a, 0x46, 0xcb, 0xa6, 0x2a, 0xf6, 0xf6, 0x00, 0xcf, 0x0f, 0xc7,
	0x00, 0x0d, 0xbf, 0xdb, 0x25, 0x41, 0x80, 0x96, 0xb0, 0x03, 0xcb, 0x07, 0xe4, 0xd0, 0x1f, 0x1c,
	0xf7, 0x91, 0xe1, 0x3d, 0x87, 0x8d, 0xb9, 0x75, 0xf1, 0x0a, 0x58, 0x83, 0x80, 0x50, 0xb4, 0x84,
	0x6d, 0xa8, 0x1f, 0xd1, 0xb3, 0x41, 0x0f, 0x19, 0x45, 0xf1, 0xc4, 0x0f, 0xde, 0xa2, 0x5a, 0x51,
	0x3c, 0xeb, 0xbf, 0x22, 0x14, 0x99, 0xde, 0x57, 0x58, 0x9b, 0xf1, 0x57, 0xa0, 0x4e, 0xcf, 0x4e,
	0x89, 0x1e, 0x43, 0xde, 0x93, 0xee, 0xa0, 0x4f, 0x90, 0x51, 0x50, 0xde, 0xd1, 0xd7, 0x7d, 0x82,
	0x6a, 0x78, 0x03, 0xd6, 0x54, 0x78, 0x3e, 0xe9, 0x9a, 0x05, 0x89, 0x12, 0xff, 0x00, 0x59, 0x18,
	0xc1, 0x6a, 0x11, 0x4d, 0x7b, 0x75, 0xbc, 0x0e, 0xa0, 0x2a, 0x9a, 0xde, 0xc0, 0xab, 0xb0, 0xd2,
	0x23, 0xf4, 0xe4, 0xdc, 0x3f, 0x3e, 0x46, 0xcb, 0xde, 0x77, 0x03, 0xd6, 0x7d, 0x16, 0x07, 0x79,
	0x98, 0x8f, 0xa4, 0x76, 0x70, 0x1b, 0xea, 0xe2, 0x4b, 0xc2, 0x33, 0x75, 0x68, 0x9b, 0xea, 0xa4,
	0xa8, 0x5e, 0x66, 0x62, 0x94, 0xaa, 0xdb, 0xd9, 0x54, 0x27, 0x78, 0x0b, 0x1a, 0x32, 0x8f, 0xd8,
	0xc7, 0xb1, 0xba, 0xc3, 0x0a, 0x2d, 0x33, 0xfc, 0x14, 0x96, 0x79, 0x92, 0x67, 0x11, 0x97, 0xae,
	0xd5, 0x34, 0x5b, 0x4e, 0x67, 0xfb, 0xcf, 0x07, 0xa2, 0x13, 0x28, 0x7e, 0x01, 0xf0, 0xeb, 0x3a,
	0x6e, 0xbd, 0x69, 0xb4, 0x9c, 0xce, 0xfd, 0x19, 0xe2, 0xa1, 0xec, 0x4d, 0x01, 0x9a, 0x5c, 0x61,
	0x78, 0x43, 0xd8, 0x39, 0x11, 0x17, 0xd1, 0x70, 0x5c, 0xea, 0x47, 0x5c, 0x52, 0xfe, 0x69, 0xc4,
	0x65, 0xae, 0x37, 0x43, 0x60, 0xca, 0x8c, 0x95, 0x7b, 0x15, 0x61, 0xe1, 0x33, 0x64, 0x71, 0x90,
	0x72, 0xe6, 0xd6, 0x16, 0xfb, 0x2c, 0xa1, 0xde, 0x03, 0xb8, 0x37, 0x3f, 0x47, 0xa6, 0x22, 0x91,
	0xfa, 0xfe, 0xde, 0x23, 0xd8, 0xa4, 0xfc, 0x5a, 0x7c, 0xe6, 0x3e, 0x8b, 0x6f, 0x76, 0xe0, 0xb9,
	0xb0, 0x55, 0x81, 0x56, 0x45, 0x86, 0xb0, 0x33, 0xed, 0xfc, 0xe7, 0x6d, 0xe6, 0xe7, 0x54, 0x8d,
	0x3c, 0x9e, 0x18, 0x39, 0xe0, 0xc3, 0x70, 0x14, 0xe7, 0x8b, 0x97, 0x9a, 0x4a, 0x56, 0x19, 0x55,
	0xc9, 0x0f, 0x80, 0x03, 0xbe, 0x58, 0xe8, 0x1f, 0x37, 0xda, 0x84, 0x5b, 0x13, 0xf5, 0xea, 0xd0,
	0x5d, 0x70, 0x8f, 0x54, 0x59, 0xff, 0xdb, 0x17, 0xec, 0xd0, 0x83, 0x3b, 0xb3, 0xe8, 0x8a, 0x14,
	0x7e, 0x02, 0x8d, 0x8c, 0xcb, 0x51, 0x9c, 0x2b, 0x86, 0xd3, 0xb9, 0xfb, 0xbb, 0xad, 0xca, 0x07,
	0x45, 0x4b, 0xe8, 0xfe, 0x33, 0x78, 0x28, 0xb2, 0xcb, 0x76, 0x98, 0x86, 0xec, 0x8a, 0xcf, 0x10,
	0xd4, 0x93, 0xcb, 0x44, 0xf9, 0x28, 0xef, 0xdb, 0x3e, 0x8b, 0x15, 0x55, 0x7e, 0x33, 0x8c, 0x9f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x5a, 0xf8, 0x89, 0x4e, 0xae, 0x05, 0x00, 0x00,
}
