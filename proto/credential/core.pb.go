// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: credential/core.proto

package credential

import (
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
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

// Core credential that combines Private, Public, and/or Realm so that credential.Private or
// credential.Public that are gathered from a credential.Realm} are properly scoped when used.
//
// A core credential must always have an Origin, but only needs 1 of Private, Public, or Realm set.
type Core struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	// Origin - The origin of this Credential.
	Origin *Origin `protobuf:"bytes,10,opt,name=Origin,proto3" json:"Origin,omitempty"`
	// Private - The credential.Private either gathered from Realm,
	// or used to credential.ReplayableHash authenticate into the realm.
	Private *Private `protobuf:"bytes,11,opt,name=Private,proto3" json:"Private,omitempty"`
	// Public - The credential.Public gathered from Realm.
	Public *Public `protobuf:"bytes,12,opt,name=Public,proto3" json:"Public,omitempty"`
	// Realm - credential.Realm where Private and/or Public was gathered
	// and/or the credential.Realm to which Private and/or Public can be
	// used to authenticate.
	Realm *Realm `protobuf:"bytes,13,opt,name=Realm,proto3" json:"Realm,omitempty"`
	// repeated Login Logins = 14;
	LoginsCount int32 `protobuf:"varint,8,opt,name=LoginsCount,proto3" json:"LoginsCount,omitempty"`
}

func (x *Core) Reset() {
	*x = Core{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credential_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Core) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Core) ProtoMessage() {}

func (x *Core) ProtoReflect() protoreflect.Message {
	mi := &file_credential_core_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Core.ProtoReflect.Descriptor instead.
func (*Core) Descriptor() ([]byte, []int) {
	return file_credential_core_proto_rawDescGZIP(), []int{0}
}

func (x *Core) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Core) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Core) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Core) GetOrigin() *Origin {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *Core) GetPrivate() *Private {
	if x != nil {
		return x.Private
	}
	return nil
}

func (x *Core) GetPublic() *Public {
	if x != nil {
		return x.Public
	}
	return nil
}

func (x *Core) GetRealm() *Realm {
	if x != nil {
		return x.Realm
	}
	return nil
}

func (x *Core) GetLoginsCount() int32 {
	if x != nil {
		return x.LoginsCount
	}
	return 0
}

var File_credential_core_proto protoreflect.FileDescriptor

var file_credential_core_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f,
	0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x2f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2f, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x2f, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x02, 0x0a,
	0x04, 0x43, 0x6f, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0e, 0xba, 0xb9, 0x19, 0x0a, 0x0a, 0x08, 0x12, 0x04, 0x75, 0x75, 0x69, 0x64, 0x28,
	0x01, 0x52, 0x02, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x38, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x4f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x42, 0x08, 0xba,
	0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12,
	0x2d, 0x0a, 0x07, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x07, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x2a,
	0x0a, 0x06, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x52, 0x06, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x27, 0x0a, 0x05, 0x52, 0x65,
	0x61, 0x6c, 0x6d, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x52, 0x05, 0x52, 0x65,
	0x61, 0x6c, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x73,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x42, 0x8f, 0x01,
	0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x42, 0x09, 0x43, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x78, 0x6c, 0x61, 0x6e,
	0x64, 0x6f, 0x6e, 0x2f, 0x61, 0x69, 0x6d, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa,
	0x02, 0x0a, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0xca, 0x02, 0x0a, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0xe2, 0x02, 0x16, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0a, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_credential_core_proto_rawDescOnce sync.Once
	file_credential_core_proto_rawDescData = file_credential_core_proto_rawDesc
)

func file_credential_core_proto_rawDescGZIP() []byte {
	file_credential_core_proto_rawDescOnce.Do(func() {
		file_credential_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_credential_core_proto_rawDescData)
	})
	return file_credential_core_proto_rawDescData
}

var file_credential_core_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_credential_core_proto_goTypes = []interface{}{
	(*Core)(nil),                  // 0: credential.Core
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*Origin)(nil),                // 2: credential.Origin
	(*Private)(nil),               // 3: credential.Private
	(*Public)(nil),                // 4: credential.Public
	(*Realm)(nil),                 // 5: credential.Realm
}
var file_credential_core_proto_depIdxs = []int32{
	1, // 0: credential.Core.CreatedAt:type_name -> google.protobuf.Timestamp
	1, // 1: credential.Core.UpdatedAt:type_name -> google.protobuf.Timestamp
	2, // 2: credential.Core.Origin:type_name -> credential.Origin
	3, // 3: credential.Core.Private:type_name -> credential.Private
	4, // 4: credential.Core.Public:type_name -> credential.Public
	5, // 5: credential.Core.Realm:type_name -> credential.Realm
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_credential_core_proto_init() }
func file_credential_core_proto_init() {
	if File_credential_core_proto != nil {
		return
	}
	file_credential_origin_proto_init()
	file_credential_private_proto_init()
	file_credential_public_proto_init()
	file_credential_realm_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_credential_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Core); i {
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
			RawDescriptor: file_credential_core_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_credential_core_proto_goTypes,
		DependencyIndexes: file_credential_core_proto_depIdxs,
		MessageInfos:      file_credential_core_proto_msgTypes,
	}.Build()
	File_credential_core_proto = out.File
	file_credential_core_proto_rawDesc = nil
	file_credential_core_proto_goTypes = nil
	file_credential_core_proto_depIdxs = nil
}
