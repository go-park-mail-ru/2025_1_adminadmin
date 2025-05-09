// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: proto/auth.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CheckRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Login         string                 `protobuf:"bytes,1,opt,name=Login,proto3" json:"Login,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckRequest) Reset() {
	*x = CheckRequest{}
	mi := &file_proto_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRequest) ProtoMessage() {}

func (x *CheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRequest.ProtoReflect.Descriptor instead.
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{0}
}

func (x *CheckRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

type AddressRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Login         string                 `protobuf:"bytes,1,opt,name=Login,proto3" json:"Login,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddressRequest) Reset() {
	*x = AddressRequest{}
	mi := &file_proto_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressRequest) ProtoMessage() {}

func (x *AddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressRequest.ProtoReflect.Descriptor instead.
func (*AddressRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{1}
}

func (x *AddressRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

type SignInRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Login         string                 `protobuf:"bytes,1,opt,name=Login,proto3" json:"Login,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignInRequest) Reset() {
	*x = SignInRequest{}
	mi := &file_proto_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInRequest) ProtoMessage() {}

func (x *SignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInRequest.ProtoReflect.Descriptor instead.
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{2}
}

func (x *SignInRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *SignInRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SignUpRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Login         string                 `protobuf:"bytes,1,opt,name=Login,proto3" json:"Login,omitempty"`
	FirstName     string                 `protobuf:"bytes,2,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName      string                 `protobuf:"bytes,3,opt,name=LastName,proto3" json:"LastName,omitempty"`
	PhoneNumber   string                 `protobuf:"bytes,4,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
	Password      string                 `protobuf:"bytes,5,opt,name=Password,proto3" json:"Password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignUpRequest) Reset() {
	*x = SignUpRequest{}
	mi := &file_proto_auth_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpRequest) ProtoMessage() {}

func (x *SignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpRequest.ProtoReflect.Descriptor instead.
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{3}
}

func (x *SignUpRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *SignUpRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *SignUpRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *SignUpRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *SignUpRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Login         string                 `protobuf:"bytes,1,opt,name=Login,proto3" json:"Login,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
	FirstName     string                 `protobuf:"bytes,3,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName      string                 `protobuf:"bytes,4,opt,name=LastName,proto3" json:"LastName,omitempty"`
	PhoneNumber   string                 `protobuf:"bytes,5,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
	Password      string                 `protobuf:"bytes,6,opt,name=Password,proto3" json:"Password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	mi := &file_proto_auth_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateUserRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *UpdateUserRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateUserRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UpdateUserRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UpdateUserRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *UpdateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UpdateUserPicRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Login         string                 `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	UserPic       []byte                 `protobuf:"bytes,2,opt,name=user_pic,json=userPic,proto3" json:"user_pic,omitempty"`
	FileExtension string                 `protobuf:"bytes,3,opt,name=file_extension,json=fileExtension,proto3" json:"file_extension,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserPicRequest) Reset() {
	*x = UpdateUserPicRequest{}
	mi := &file_proto_auth_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserPicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserPicRequest) ProtoMessage() {}

func (x *UpdateUserPicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserPicRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserPicRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateUserPicRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *UpdateUserPicRequest) GetUserPic() []byte {
	if x != nil {
		return x.UserPic
	}
	return nil
}

func (x *UpdateUserPicRequest) GetFileExtension() string {
	if x != nil {
		return x.FileExtension
	}
	return ""
}

type DeleteAddressRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAddressRequest) Reset() {
	*x = DeleteAddressRequest{}
	mi := &file_proto_auth_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAddressRequest) ProtoMessage() {}

func (x *DeleteAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAddressRequest.ProtoReflect.Descriptor instead.
func (*DeleteAddressRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteAddressRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Address struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Address       string                 `protobuf:"bytes,2,opt,name=Address,proto3" json:"Address,omitempty"`
	UserId        string                 `protobuf:"bytes,3,opt,name=UserId,proto3" json:"UserId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Address) Reset() {
	*x = Address{}
	mi := &file_proto_auth_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{7}
}

func (x *Address) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Address) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Address) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Login         string                 `protobuf:"bytes,1,opt,name=Login,proto3" json:"Login,omitempty"`
	PhoneNumber   string                 `protobuf:"bytes,2,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
	Id            string                 `protobuf:"bytes,3,opt,name=Id,proto3" json:"Id,omitempty"`
	FirstName     string                 `protobuf:"bytes,4,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName      string                 `protobuf:"bytes,5,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Description   string                 `protobuf:"bytes,6,opt,name=Description,proto3" json:"Description,omitempty"`
	UserPic       string                 `protobuf:"bytes,7,opt,name=UserPic,proto3" json:"UserPic,omitempty"`
	Token         string                 `protobuf:"bytes,8,opt,name=Token,proto3" json:"Token,omitempty"`
	CsrfToken     string                 `protobuf:"bytes,9,opt,name=CsrfToken,proto3" json:"CsrfToken,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	mi := &file_proto_auth_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{8}
}

func (x *UserResponse) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *UserResponse) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *UserResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserResponse) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserResponse) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UserResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UserResponse) GetUserPic() string {
	if x != nil {
		return x.UserPic
	}
	return ""
}

func (x *UserResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UserResponse) GetCsrfToken() string {
	if x != nil {
		return x.CsrfToken
	}
	return ""
}

type AddressListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Addresses     []*Address             `protobuf:"bytes,1,rep,name=Addresses,proto3" json:"Addresses,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddressListResponse) Reset() {
	*x = AddressListResponse{}
	mi := &file_proto_auth_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddressListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressListResponse) ProtoMessage() {}

func (x *AddressListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressListResponse.ProtoReflect.Descriptor instead.
func (*AddressListResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{9}
}

func (x *AddressListResponse) GetAddresses() []*Address {
	if x != nil {
		return x.Addresses
	}
	return nil
}

var File_proto_auth_proto protoreflect.FileDescriptor

const file_proto_auth_proto_rawDesc = "" +
	"\n" +
	"\x10proto/auth.proto\x12\x04auth\x1a\x1bgoogle/protobuf/empty.proto\"$\n" +
	"\fCheckRequest\x12\x14\n" +
	"\x05Login\x18\x01 \x01(\tR\x05Login\"&\n" +
	"\x0eAddressRequest\x12\x14\n" +
	"\x05Login\x18\x01 \x01(\tR\x05Login\"A\n" +
	"\rSignInRequest\x12\x14\n" +
	"\x05Login\x18\x01 \x01(\tR\x05Login\x12\x1a\n" +
	"\bPassword\x18\x02 \x01(\tR\bPassword\"\x9d\x01\n" +
	"\rSignUpRequest\x12\x14\n" +
	"\x05Login\x18\x01 \x01(\tR\x05Login\x12\x1c\n" +
	"\tFirstName\x18\x02 \x01(\tR\tFirstName\x12\x1a\n" +
	"\bLastName\x18\x03 \x01(\tR\bLastName\x12 \n" +
	"\vPhoneNumber\x18\x04 \x01(\tR\vPhoneNumber\x12\x1a\n" +
	"\bPassword\x18\x05 \x01(\tR\bPassword\"\xc3\x01\n" +
	"\x11UpdateUserRequest\x12\x14\n" +
	"\x05Login\x18\x01 \x01(\tR\x05Login\x12 \n" +
	"\vDescription\x18\x02 \x01(\tR\vDescription\x12\x1c\n" +
	"\tFirstName\x18\x03 \x01(\tR\tFirstName\x12\x1a\n" +
	"\bLastName\x18\x04 \x01(\tR\bLastName\x12 \n" +
	"\vPhoneNumber\x18\x05 \x01(\tR\vPhoneNumber\x12\x1a\n" +
	"\bPassword\x18\x06 \x01(\tR\bPassword\"n\n" +
	"\x14UpdateUserPicRequest\x12\x14\n" +
	"\x05login\x18\x01 \x01(\tR\x05login\x12\x19\n" +
	"\buser_pic\x18\x02 \x01(\fR\auserPic\x12%\n" +
	"\x0efile_extension\x18\x03 \x01(\tR\rfileExtension\"&\n" +
	"\x14DeleteAddressRequest\x12\x0e\n" +
	"\x02Id\x18\x01 \x01(\tR\x02Id\"K\n" +
	"\aAddress\x12\x0e\n" +
	"\x02Id\x18\x01 \x01(\tR\x02Id\x12\x18\n" +
	"\aAddress\x18\x02 \x01(\tR\aAddress\x12\x16\n" +
	"\x06UserId\x18\x03 \x01(\tR\x06UserId\"\x80\x02\n" +
	"\fUserResponse\x12\x14\n" +
	"\x05Login\x18\x01 \x01(\tR\x05Login\x12 \n" +
	"\vPhoneNumber\x18\x02 \x01(\tR\vPhoneNumber\x12\x0e\n" +
	"\x02Id\x18\x03 \x01(\tR\x02Id\x12\x1c\n" +
	"\tFirstName\x18\x04 \x01(\tR\tFirstName\x12\x1a\n" +
	"\bLastName\x18\x05 \x01(\tR\bLastName\x12 \n" +
	"\vDescription\x18\x06 \x01(\tR\vDescription\x12\x18\n" +
	"\aUserPic\x18\a \x01(\tR\aUserPic\x12\x14\n" +
	"\x05Token\x18\b \x01(\tR\x05Token\x12\x1c\n" +
	"\tCsrfToken\x18\t \x01(\tR\tCsrfToken\"B\n" +
	"\x13AddressListResponse\x12+\n" +
	"\tAddresses\x18\x01 \x03(\v2\r.auth.AddressR\tAddresses2\xef\x03\n" +
	"\vAuthService\x123\n" +
	"\x06SignIn\x12\x13.auth.SignInRequest\x1a\x12.auth.UserResponse\"\x00\x123\n" +
	"\x06SignUp\x12\x13.auth.SignUpRequest\x1a\x12.auth.UserResponse\"\x00\x121\n" +
	"\x05Check\x12\x12.auth.CheckRequest\x1a\x12.auth.UserResponse\"\x00\x12;\n" +
	"\n" +
	"UpdateUser\x12\x17.auth.UpdateUserRequest\x1a\x12.auth.UserResponse\"\x00\x12A\n" +
	"\rUpdateUserPic\x12\x1a.auth.UpdateUserPicRequest\x1a\x12.auth.UserResponse\"\x00\x12E\n" +
	"\x10GetUserAddresses\x12\x14.auth.AddressRequest\x1a\x19.auth.AddressListResponse\"\x00\x12E\n" +
	"\rDeleteAddress\x12\x1a.auth.DeleteAddressRequest\x1a\x16.google.protobuf.Empty\"\x00\x125\n" +
	"\n" +
	"AddAddress\x12\r.auth.Address\x1a\x16.google.protobuf.Empty\"\x00B'Z%./internal/pkg/auth/delivery/grpc/genb\x06proto3"

var (
	file_proto_auth_proto_rawDescOnce sync.Once
	file_proto_auth_proto_rawDescData []byte
)

func file_proto_auth_proto_rawDescGZIP() []byte {
	file_proto_auth_proto_rawDescOnce.Do(func() {
		file_proto_auth_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_auth_proto_rawDesc), len(file_proto_auth_proto_rawDesc)))
	})
	return file_proto_auth_proto_rawDescData
}

var file_proto_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_auth_proto_goTypes = []any{
	(*CheckRequest)(nil),         // 0: auth.CheckRequest
	(*AddressRequest)(nil),       // 1: auth.AddressRequest
	(*SignInRequest)(nil),        // 2: auth.SignInRequest
	(*SignUpRequest)(nil),        // 3: auth.SignUpRequest
	(*UpdateUserRequest)(nil),    // 4: auth.UpdateUserRequest
	(*UpdateUserPicRequest)(nil), // 5: auth.UpdateUserPicRequest
	(*DeleteAddressRequest)(nil), // 6: auth.DeleteAddressRequest
	(*Address)(nil),              // 7: auth.Address
	(*UserResponse)(nil),         // 8: auth.UserResponse
	(*AddressListResponse)(nil),  // 9: auth.AddressListResponse
	(*emptypb.Empty)(nil),        // 10: google.protobuf.Empty
}
var file_proto_auth_proto_depIdxs = []int32{
	7,  // 0: auth.AddressListResponse.Addresses:type_name -> auth.Address
	2,  // 1: auth.AuthService.SignIn:input_type -> auth.SignInRequest
	3,  // 2: auth.AuthService.SignUp:input_type -> auth.SignUpRequest
	0,  // 3: auth.AuthService.Check:input_type -> auth.CheckRequest
	4,  // 4: auth.AuthService.UpdateUser:input_type -> auth.UpdateUserRequest
	5,  // 5: auth.AuthService.UpdateUserPic:input_type -> auth.UpdateUserPicRequest
	1,  // 6: auth.AuthService.GetUserAddresses:input_type -> auth.AddressRequest
	6,  // 7: auth.AuthService.DeleteAddress:input_type -> auth.DeleteAddressRequest
	7,  // 8: auth.AuthService.AddAddress:input_type -> auth.Address
	8,  // 9: auth.AuthService.SignIn:output_type -> auth.UserResponse
	8,  // 10: auth.AuthService.SignUp:output_type -> auth.UserResponse
	8,  // 11: auth.AuthService.Check:output_type -> auth.UserResponse
	8,  // 12: auth.AuthService.UpdateUser:output_type -> auth.UserResponse
	8,  // 13: auth.AuthService.UpdateUserPic:output_type -> auth.UserResponse
	9,  // 14: auth.AuthService.GetUserAddresses:output_type -> auth.AddressListResponse
	10, // 15: auth.AuthService.DeleteAddress:output_type -> google.protobuf.Empty
	10, // 16: auth.AuthService.AddAddress:output_type -> google.protobuf.Empty
	9,  // [9:17] is the sub-list for method output_type
	1,  // [1:9] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_proto_auth_proto_init() }
func file_proto_auth_proto_init() {
	if File_proto_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_auth_proto_rawDesc), len(file_proto_auth_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_auth_proto_goTypes,
		DependencyIndexes: file_proto_auth_proto_depIdxs,
		MessageInfos:      file_proto_auth_proto_msgTypes,
	}.Build()
	File_proto_auth_proto = out.File
	file_proto_auth_proto_goTypes = nil
	file_proto_auth_proto_depIdxs = nil
}
