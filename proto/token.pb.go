// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: proto/token.proto

package go_tokenmgmt_grpc

import (
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

type NewToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Domain string `protobuf:"bytes,2,opt,name=Domain,proto3" json:"Domain,omitempty"`
	State  string `protobuf:"bytes,4,opt,name=State,proto3" json:"State,omitempty"`
	Id     string `protobuf:"bytes,3,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *NewToken) Reset() {
	*x = NewToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewToken) ProtoMessage() {}

func (x *NewToken) ProtoReflect() protoreflect.Message {
	mi := &file_proto_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewToken.ProtoReflect.Descriptor instead.
func (*NewToken) Descriptor() ([]byte, []int) {
	return file_proto_token_proto_rawDescGZIP(), []int{0}
}

func (x *NewToken) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewToken) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *NewToken) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *NewToken) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Domain string `protobuf:"bytes,2,opt,name=Domain,proto3" json:"Domain,omitempty"`
	State  string `protobuf:"bytes,4,opt,name=State,proto3" json:"State,omitempty"`
	Id     string `protobuf:"bytes,3,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_proto_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_proto_token_proto_rawDescGZIP(), []int{1}
}

func (x *Token) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Token) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Token) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Token) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EmptyToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EmptyToken) Reset() {
	*x = EmptyToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyToken) ProtoMessage() {}

func (x *EmptyToken) ProtoReflect() protoreflect.Message {
	mi := &file_proto_token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyToken.ProtoReflect.Descriptor instead.
func (*EmptyToken) Descriptor() ([]byte, []int) {
	return file_proto_token_proto_rawDescGZIP(), []int{2}
}

func (x *EmptyToken) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type TokenList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tokens []*Token `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
}

func (x *TokenList) Reset() {
	*x = TokenList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_token_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenList) ProtoMessage() {}

func (x *TokenList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_token_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenList.ProtoReflect.Descriptor instead.
func (*TokenList) Descriptor() ([]byte, []int) {
	return file_proto_token_proto_rawDescGZIP(), []int{3}
}

func (x *TokenList) GetTokens() []*Token {
	if x != nil {
		return x.Tokens
	}
	return nil
}

type TokenInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *TokenInfo) Reset() {
	*x = TokenInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_token_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenInfo) ProtoMessage() {}

func (x *TokenInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_token_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenInfo.ProtoReflect.Descriptor instead.
func (*TokenInfo) Descriptor() ([]byte, []int) {
	return file_proto_token_proto_rawDescGZIP(), []int{4}
}

func (x *TokenInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_proto_token_proto protoreflect.FileDescriptor

var file_proto_token_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x67, 0x6d, 0x74, 0x22, 0x5c,
	0x0a, 0x08, 0x4e, 0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x22, 0x59, 0x0a, 0x05,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x22, 0x26, 0x0a, 0x0a, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x35, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x06,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x06,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0x1b, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x64, 0x32, 0xe6, 0x01, 0x0a, 0x03, 0x54, 0x6b, 0x6e, 0x12, 0x39, 0x0a, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x13, 0x2e,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x4e, 0x65, 0x77, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x1a, 0x10, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x65,
	0x6b, 0x6e, 0x73, 0x12, 0x10, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x67, 0x6d, 0x74, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x10, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x67, 0x6d,
	0x74, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x09, 0x44, 0x72, 0x6f,
	0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x67,
	0x6d, 0x74, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x15, 0x2e, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0a, 0x57, 0x72, 0x69, 0x74, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x13, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x67, 0x6d, 0x74, 0x2e,
	0x4e, 0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x10, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f,
	0x6b, 0x61, 0x72, 0x74, 0x68, 0x69, 0x6b, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x67, 0x6d, 0x74, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x67, 0x6f,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x67, 0x6d, 0x74, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_token_proto_rawDescOnce sync.Once
	file_proto_token_proto_rawDescData = file_proto_token_proto_rawDesc
)

func file_proto_token_proto_rawDescGZIP() []byte {
	file_proto_token_proto_rawDescOnce.Do(func() {
		file_proto_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_token_proto_rawDescData)
	})
	return file_proto_token_proto_rawDescData
}

var file_proto_token_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_token_proto_goTypes = []interface{}{
	(*NewToken)(nil),   // 0: tokenmgmt.NewToken
	(*Token)(nil),      // 1: tokenmgmt.Token
	(*EmptyToken)(nil), // 2: tokenmgmt.EmptyToken
	(*TokenList)(nil),  // 3: tokenmgmt.TokenList
	(*TokenInfo)(nil),  // 4: tokenmgmt.TokenInfo
}
var file_proto_token_proto_depIdxs = []int32{
	1, // 0: tokenmgmt.TokenList.tokens:type_name -> tokenmgmt.Token
	0, // 1: tokenmgmt.Tkn.CreateNewToken:input_type -> tokenmgmt.NewToken
	1, // 2: tokenmgmt.Tkn.GetToekns:input_type -> tokenmgmt.Token
	4, // 3: tokenmgmt.Tkn.DropToken:input_type -> tokenmgmt.TokenInfo
	0, // 4: tokenmgmt.Tkn.WriteToken:input_type -> tokenmgmt.NewToken
	1, // 5: tokenmgmt.Tkn.CreateNewToken:output_type -> tokenmgmt.Token
	1, // 6: tokenmgmt.Tkn.GetToekns:output_type -> tokenmgmt.Token
	2, // 7: tokenmgmt.Tkn.DropToken:output_type -> tokenmgmt.EmptyToken
	1, // 8: tokenmgmt.Tkn.WriteToken:output_type -> tokenmgmt.Token
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_token_proto_init() }
func file_proto_token_proto_init() {
	if File_proto_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewToken); i {
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
		file_proto_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_proto_token_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyToken); i {
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
		file_proto_token_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenList); i {
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
		file_proto_token_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenInfo); i {
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
			RawDescriptor: file_proto_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_token_proto_goTypes,
		DependencyIndexes: file_proto_token_proto_depIdxs,
		MessageInfos:      file_proto_token_proto_msgTypes,
	}.Build()
	File_proto_token_proto = out.File
	file_proto_token_proto_rawDesc = nil
	file_proto_token_proto_goTypes = nil
	file_proto_token_proto_depIdxs = nil
}
