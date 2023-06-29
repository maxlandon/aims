// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: credential/private.proto

package credential

import (
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	types "github.com/infobloxopen/protoc-gen-gorm/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// PrivateType - The type of a Private credential. Each of these types might
// trigger different validations when using fetching/creating Private in DB.
type PrivateType int32

const (
	// A password is the default type, such that if the Private.Data is
	// empty, the type will not be able to save itself in DB without
	// explicitely specifying either its .Type field, or by creating a
	// credential.BlankPassword first.
	PrivateType_Password      PrivateType = 0
	PrivateType_BlankPassword PrivateType = 1
	// ReplayableHash - A credential.PasswordHash password hash that
	// can be replayed to authenticate to additional services.
	PrivateType_ReplayableHash PrivateType = 2
	// NonReplayableHash - A credential.PasswordHash password hash that
	// cannot be replayed to authenticate to other services. Contrasts
	// with credential.ReplayableHash. The NonReplayableHash.Data is any
	// password hash, such as those recovered from `/etc/passwd` or `/etc/shadow`.
	PrivateType_NonReplayableHash PrivateType = 3
	// NTLMHash - A credential.Private password hash that can be credential.ReplayableHash replayed
	// to authenticate to SMB.  It is composed of two hash hex digests (where the hash bytes are
	// printed as a hexadecimal string where 2 characters represent a byte of the original hash with
	// the high nibble first): (1) {lanManagerHexDigestRegexp, the LAN Manager hash's hex digest} and
	// (2) {ntLanManagerHexDigestRegexp, the NTLM hash's hex digest}.
	PrivateType_NTLMHash PrivateType = 4
	// PostgresMD5 - A credential.Private.PasswordHash password hash that can be
	// credential.Private.ReplayableHash replayed to authenticate to PostgreSQL
	// servers. It is composed of a hexadecimal string of 32 characters prepended
	// by the string 'md5'
	PrivateType_PostgresMD5 PrivateType = 5
	// Key - Usually a cryptographic key. When the credential is of this
	// type, various validations are performed on the formatting, encoding
	// potential encryption of the key data, etc.
	// As well, various methods are provided to get cipher and fingerprints.
	PrivateType_Key PrivateType = 6
	// JWT - When the private credential is a JSON Web Token, it is used
	// in conjunction with a given credential.Public to create a JWT string
	// stored in its .Data field.
	// As well, some other types (eg. credential.Core) allow you to create
	// JWT from their own data, like usernames, public/private key pairs, etc.
	PrivateType_JWT PrivateType = 7
)

// Enum value maps for PrivateType.
var (
	PrivateType_name = map[int32]string{
		0: "Password",
		1: "BlankPassword",
		2: "ReplayableHash",
		3: "NonReplayableHash",
		4: "NTLMHash",
		5: "PostgresMD5",
		6: "Key",
		7: "JWT",
	}
	PrivateType_value = map[string]int32{
		"Password":          0,
		"BlankPassword":     1,
		"ReplayableHash":    2,
		"NonReplayableHash": 3,
		"NTLMHash":          4,
		"PostgresMD5":       5,
		"Key":               6,
		"JWT":               7,
	}
)

func (x PrivateType) Enum() *PrivateType {
	p := new(PrivateType)
	*p = x
	return p
}

func (x PrivateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PrivateType) Descriptor() protoreflect.EnumDescriptor {
	return file_credential_private_proto_enumTypes[0].Descriptor()
}

func (PrivateType) Type() protoreflect.EnumType {
	return &file_credential_private_proto_enumTypes[0]
}

func (x PrivateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PrivateType.Descriptor instead.
func (PrivateType) EnumDescriptor() ([]byte, []int) {
	return file_credential_private_proto_rawDescGZIP(), []int{0}
}

// Private - Any credential that should not be publicly exposed.
type Private struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *types.UUID `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// @gotags: display:"Created at" readonly:"true"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	// @gotags: display:"Updated at" readonly:"true"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	// @gotags: display:"Type" readonly:"true"
	Type PrivateType `protobuf:"varint,10,opt,name=Type,proto3,enum=credential.PrivateType" json:"Type,omitempty"`
	// No data passed into Maltego at this point.
	Data string `protobuf:"bytes,11,opt,name=Data,proto3" json:"Data,omitempty"`
	// @gotags: display:"JTR Format" readonly:"true"
	JTRFormat string `protobuf:"bytes,12,opt,name=JTRFormat,proto3" json:"JTRFormat,omitempty"`
}

func (x *Private) Reset() {
	*x = Private{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credential_private_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Private) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Private) ProtoMessage() {}

func (x *Private) ProtoReflect() protoreflect.Message {
	mi := &file_credential_private_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Private.ProtoReflect.Descriptor instead.
func (*Private) Descriptor() ([]byte, []int) {
	return file_credential_private_proto_rawDescGZIP(), []int{0}
}

func (x *Private) GetId() *types.UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Private) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Private) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Private) GetType() PrivateType {
	if x != nil {
		return x.Type
	}
	return PrivateType_Password
}

func (x *Private) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Private) GetJTRFormat() string {
	if x != nil {
		return x.JTRFormat
	}
	return ""
}

var File_credential_private_proto protoreflect.FileDescriptor

var file_credential_private_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96,
	0x02, 0x0a, 0x07, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x42, 0x0e, 0xba, 0xb9, 0x19, 0x0a, 0x0a, 0x08,
	0x12, 0x04, 0x75, 0x75, 0x69, 0x64, 0x28, 0x01, 0x52, 0x02, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x2b, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17,
	0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1c, 0x0a, 0x09, 0x4a, 0x54, 0x52, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4a, 0x54, 0x52, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x3a,
	0x06, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x2a, 0x8a, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x42, 0x6c, 0x61, 0x6e, 0x6b, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6c,
	0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11,
	0x4e, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x61, 0x73,
	0x68, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x54, 0x4c, 0x4d, 0x48, 0x61, 0x73, 0x68, 0x10,
	0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x4d, 0x44, 0x35,
	0x10, 0x05, 0x12, 0x07, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x10, 0x06, 0x12, 0x07, 0x0a, 0x03, 0x4a,
	0x57, 0x54, 0x10, 0x07, 0x42, 0x99, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x42, 0x0c, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x78, 0x6c, 0x61, 0x6e, 0x64, 0x6f, 0x6e, 0x2f, 0x61, 0x69,
	0x6d, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58,
	0xaa, 0x02, 0x0a, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0xca, 0x02, 0x0a,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0xe2, 0x02, 0x16, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_credential_private_proto_rawDescOnce sync.Once
	file_credential_private_proto_rawDescData = file_credential_private_proto_rawDesc
)

func file_credential_private_proto_rawDescGZIP() []byte {
	file_credential_private_proto_rawDescOnce.Do(func() {
		file_credential_private_proto_rawDescData = protoimpl.X.CompressGZIP(file_credential_private_proto_rawDescData)
	})
	return file_credential_private_proto_rawDescData
}

var file_credential_private_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_credential_private_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_credential_private_proto_goTypes = []interface{}{
	(PrivateType)(0),              // 0: credential.PrivateType
	(*Private)(nil),               // 1: credential.Private
	(*types.UUID)(nil),            // 2: gorm.types.UUID
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_credential_private_proto_depIdxs = []int32{
	2, // 0: credential.Private.Id:type_name -> gorm.types.UUID
	3, // 1: credential.Private.CreatedAt:type_name -> google.protobuf.Timestamp
	3, // 2: credential.Private.UpdatedAt:type_name -> google.protobuf.Timestamp
	0, // 3: credential.Private.Type:type_name -> credential.PrivateType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_credential_private_proto_init() }
func file_credential_private_proto_init() {
	if File_credential_private_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_credential_private_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Private); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_credential_private_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_credential_private_proto_goTypes,
		DependencyIndexes: file_credential_private_proto_depIdxs,
		EnumInfos:         file_credential_private_proto_enumTypes,
		MessageInfos:      file_credential_private_proto_msgTypes,
	}.Build()
	File_credential_private_proto = out.File
	file_credential_private_proto_rawDesc = nil
	file_credential_private_proto_goTypes = nil
	file_credential_private_proto_depIdxs = nil
}
