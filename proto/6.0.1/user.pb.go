// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: proto/6.0.1/user.proto

package fenix

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identification
	ID            string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Username      string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Discriminator string `protobuf:"bytes,3,opt,name=discriminator,proto3" json:"discriminator,omitempty"`
	// Social
	Activity *Activity `protobuf:"bytes,4,opt,name=activity,proto3" json:"activity,omitempty"`
	Servers  []string  `protobuf:"bytes,5,rep,name=servers,proto3" json:"servers,omitempty"`
	Friends  []string  `protobuf:"bytes,6,rep,name=friends,proto3" json:"friends,omitempty"`
	// Authentication
	Token    string `protobuf:"bytes,7,opt,name=token,proto3" json:"token,omitempty"`
	Email    string `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`
	Salt     []byte `protobuf:"bytes,9,opt,name=Salt,proto3" json:"Salt,omitempty"`
	Password []byte `protobuf:"bytes,10,opt,name=Password,proto3" json:"Password,omitempty"`
	// Misc
	Settings *UserSettings `protobuf:"bytes,11,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_6_0_1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_6_0_1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_6_0_1_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetDiscriminator() string {
	if x != nil {
		return x.Discriminator
	}
	return ""
}

func (x *User) GetActivity() *Activity {
	if x != nil {
		return x.Activity
	}
	return nil
}

func (x *User) GetServers() []string {
	if x != nil {
		return x.Servers
	}
	return nil
}

func (x *User) GetFriends() []string {
	if x != nil {
		return x.Friends
	}
	return nil
}

func (x *User) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetSalt() []byte {
	if x != nil {
		return x.Salt
	}
	return nil
}

func (x *User) GetPassword() []byte {
	if x != nil {
		return x.Password
	}
	return nil
}

func (x *User) GetSettings() *UserSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

type UserSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserSettings) Reset() {
	*x = UserSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_6_0_1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSettings) ProtoMessage() {}

func (x *UserSettings) ProtoReflect() protoreflect.Message {
	mi := &file_proto_6_0_1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSettings.ProtoReflect.Descriptor instead.
func (*UserSettings) Descriptor() ([]byte, []int) {
	return file_proto_6_0_1_user_proto_rawDescGZIP(), []int{1}
}

type Activity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Activity) Reset() {
	*x = Activity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_6_0_1_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Activity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activity) ProtoMessage() {}

func (x *Activity) ProtoReflect() protoreflect.Message {
	mi := &file_proto_6_0_1_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activity.ProtoReflect.Descriptor instead.
func (*Activity) Descriptor() ([]byte, []int) {
	return file_proto_6_0_1_user_proto_rawDescGZIP(), []int{2}
}

type Authenticate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ID    string `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *Authenticate) Reset() {
	*x = Authenticate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_6_0_1_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Authenticate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authenticate) ProtoMessage() {}

func (x *Authenticate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_6_0_1_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authenticate.ProtoReflect.Descriptor instead.
func (*Authenticate) Descriptor() ([]byte, []int) {
	return file_proto_6_0_1_user_proto_rawDescGZIP(), []int{3}
}

func (x *Authenticate) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Authenticate) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

var File_proto_6_0_1_user_proto protoreflect.FileDescriptor

var file_proto_6_0_1_user_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x36, 0x2e, 0x30, 0x2e, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a,
	0x0d, 0x64, 0x69, 0x73, 0x63, 0x72, 0x69, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x72, 0x69, 0x6d, 0x69, 0x6e, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x61,
	0x6c, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x53, 0x61, 0x6c, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x29, 0x0a, 0x08, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x0a, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x22, 0x34, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x32, 0x26, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x12, 0x1d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0d, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x1a, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x42,
	0x13, 0x5a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x36, 0x2e, 0x30, 0x2e, 0x31, 0x3b, 0x66,
	0x65, 0x6e, 0x69, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_6_0_1_user_proto_rawDescOnce sync.Once
	file_proto_6_0_1_user_proto_rawDescData = file_proto_6_0_1_user_proto_rawDesc
)

func file_proto_6_0_1_user_proto_rawDescGZIP() []byte {
	file_proto_6_0_1_user_proto_rawDescOnce.Do(func() {
		file_proto_6_0_1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_6_0_1_user_proto_rawDescData)
	})
	return file_proto_6_0_1_user_proto_rawDescData
}

var file_proto_6_0_1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_6_0_1_user_proto_goTypes = []interface{}{
	(*User)(nil),         // 0: User
	(*UserSettings)(nil), // 1: UserSettings
	(*Activity)(nil),     // 2: Activity
	(*Authenticate)(nil), // 3: Authenticate
}
var file_proto_6_0_1_user_proto_depIdxs = []int32{
	2, // 0: User.activity:type_name -> Activity
	1, // 1: User.settings:type_name -> UserSettings
	3, // 2: Users.Get:input_type -> Authenticate
	0, // 3: Users.Get:output_type -> User
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_6_0_1_user_proto_init() }
func file_proto_6_0_1_user_proto_init() {
	if File_proto_6_0_1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_6_0_1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_proto_6_0_1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSettings); i {
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
		file_proto_6_0_1_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Activity); i {
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
		file_proto_6_0_1_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Authenticate); i {
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
			RawDescriptor: file_proto_6_0_1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_6_0_1_user_proto_goTypes,
		DependencyIndexes: file_proto_6_0_1_user_proto_depIdxs,
		MessageInfos:      file_proto_6_0_1_user_proto_msgTypes,
	}.Build()
	File_proto_6_0_1_user_proto = out.File
	file_proto_6_0_1_user_proto_rawDesc = nil
	file_proto_6_0_1_user_proto_goTypes = nil
	file_proto_6_0_1_user_proto_depIdxs = nil
}