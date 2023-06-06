// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: client.proto

package client_service

import (
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type Client struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName                string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName                 string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Address                  string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	PhoneNumber              string `protobuf:"bytes,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	DrivingLicenseNumber     string `protobuf:"bytes,6,opt,name=driving_license_number,json=drivingLicenseNumber,proto3" json:"driving_license_number,omitempty"`
	PassportNumber           string `protobuf:"bytes,7,opt,name=passport_number,json=passportNumber,proto3" json:"passport_number,omitempty"`
	Photo                    string `protobuf:"bytes,8,opt,name=photo,proto3" json:"photo,omitempty"`
	CreatedAt                string `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt                string `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DrivingNumberGivenPlace  string `protobuf:"bytes,11,opt,name=driving_number_given_place,json=drivingNumberGivenPlace,proto3" json:"driving_number_given_place,omitempty"`
	DrivingNumberGivenDate   string `protobuf:"bytes,12,opt,name=driving_number_given_date,json=drivingNumberGivenDate,proto3" json:"driving_number_given_date,omitempty"`
	DrivingNumberGivenExpire string `protobuf:"bytes,13,opt,name=driving_number_given_expire,json=drivingNumberGivenExpire,proto3" json:"driving_number_given_expire,omitempty"`
	DateOfBirth              string `protobuf:"bytes,14,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	IsBlocked                bool   `protobuf:"varint,15,opt,name=is_blocked,json=isBlocked,proto3" json:"is_blocked,omitempty"`
	Propiska                 string `protobuf:"bytes,16,opt,name=propiska,proto3" json:"propiska,omitempty"`
	PassportPinfl            string `protobuf:"bytes,17,opt,name=passport_pinfl,json=passportPinfl,proto3" json:"passport_pinfl,omitempty"`
}

func (x *Client) Reset() {
	*x = Client{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client) ProtoMessage() {}

func (x *Client) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client.ProtoReflect.Descriptor instead.
func (*Client) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{0}
}

func (x *Client) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Client) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Client) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Client) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Client) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Client) GetDrivingLicenseNumber() string {
	if x != nil {
		return x.DrivingLicenseNumber
	}
	return ""
}

func (x *Client) GetPassportNumber() string {
	if x != nil {
		return x.PassportNumber
	}
	return ""
}

func (x *Client) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *Client) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Client) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Client) GetDrivingNumberGivenPlace() string {
	if x != nil {
		return x.DrivingNumberGivenPlace
	}
	return ""
}

func (x *Client) GetDrivingNumberGivenDate() string {
	if x != nil {
		return x.DrivingNumberGivenDate
	}
	return ""
}

func (x *Client) GetDrivingNumberGivenExpire() string {
	if x != nil {
		return x.DrivingNumberGivenExpire
	}
	return ""
}

func (x *Client) GetDateOfBirth() string {
	if x != nil {
		return x.DateOfBirth
	}
	return ""
}

func (x *Client) GetIsBlocked() bool {
	if x != nil {
		return x.IsBlocked
	}
	return false
}

func (x *Client) GetPropiska() string {
	if x != nil {
		return x.Propiska
	}
	return ""
}

func (x *Client) GetPassportPinfl() string {
	if x != nil {
		return x.PassportPinfl
	}
	return ""
}

type ClientPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ClientPrimaryKey) Reset() {
	*x = ClientPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientPrimaryKey) ProtoMessage() {}

func (x *ClientPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientPrimaryKey.ProtoReflect.Descriptor instead.
func (*ClientPrimaryKey) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{1}
}

func (x *ClientPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName                string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName                 string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Address                  string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	PhoneNumber              string `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	DrivingLicenseNumber     string `protobuf:"bytes,5,opt,name=driving_license_number,json=drivingLicenseNumber,proto3" json:"driving_license_number,omitempty"`
	PassportNumber           string `protobuf:"bytes,6,opt,name=passport_number,json=passportNumber,proto3" json:"passport_number,omitempty"`
	Photo                    string `protobuf:"bytes,7,opt,name=photo,proto3" json:"photo,omitempty"`
	DrivingNumberGivenPlace  string `protobuf:"bytes,8,opt,name=driving_number_given_place,json=drivingNumberGivenPlace,proto3" json:"driving_number_given_place,omitempty"`
	DrivingNumberGivenDate   string `protobuf:"bytes,9,opt,name=driving_number_given_date,json=drivingNumberGivenDate,proto3" json:"driving_number_given_date,omitempty"`
	DrivingNumberGivenExpire string `protobuf:"bytes,10,opt,name=driving_number_given_expire,json=drivingNumberGivenExpire,proto3" json:"driving_number_given_expire,omitempty"`
	DateOfBirth              string `protobuf:"bytes,11,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	IsBlocked                bool   `protobuf:"varint,12,opt,name=is_blocked,json=isBlocked,proto3" json:"is_blocked,omitempty"`
	Propiska                 string `protobuf:"bytes,13,opt,name=propiska,proto3" json:"propiska,omitempty"`
	PassportPinfl            string `protobuf:"bytes,14,opt,name=passport_pinfl,json=passportPinfl,proto3" json:"passport_pinfl,omitempty"`
}

func (x *CreateClientRequest) Reset() {
	*x = CreateClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClientRequest) ProtoMessage() {}

func (x *CreateClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClientRequest.ProtoReflect.Descriptor instead.
func (*CreateClientRequest) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{2}
}

func (x *CreateClientRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreateClientRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CreateClientRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateClientRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *CreateClientRequest) GetDrivingLicenseNumber() string {
	if x != nil {
		return x.DrivingLicenseNumber
	}
	return ""
}

func (x *CreateClientRequest) GetPassportNumber() string {
	if x != nil {
		return x.PassportNumber
	}
	return ""
}

func (x *CreateClientRequest) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *CreateClientRequest) GetDrivingNumberGivenPlace() string {
	if x != nil {
		return x.DrivingNumberGivenPlace
	}
	return ""
}

func (x *CreateClientRequest) GetDrivingNumberGivenDate() string {
	if x != nil {
		return x.DrivingNumberGivenDate
	}
	return ""
}

func (x *CreateClientRequest) GetDrivingNumberGivenExpire() string {
	if x != nil {
		return x.DrivingNumberGivenExpire
	}
	return ""
}

func (x *CreateClientRequest) GetDateOfBirth() string {
	if x != nil {
		return x.DateOfBirth
	}
	return ""
}

func (x *CreateClientRequest) GetIsBlocked() bool {
	if x != nil {
		return x.IsBlocked
	}
	return false
}

func (x *CreateClientRequest) GetPropiska() string {
	if x != nil {
		return x.Propiska
	}
	return ""
}

func (x *CreateClientRequest) GetPassportPinfl() string {
	if x != nil {
		return x.PassportPinfl
	}
	return ""
}

type UpdateClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName                string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName                 string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Address                  string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	PhoneNumber              string `protobuf:"bytes,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	DrivingLicenseNumber     string `protobuf:"bytes,6,opt,name=driving_license_number,json=drivingLicenseNumber,proto3" json:"driving_license_number,omitempty"`
	PassportNumber           string `protobuf:"bytes,7,opt,name=passport_number,json=passportNumber,proto3" json:"passport_number,omitempty"`
	Photo                    string `protobuf:"bytes,8,opt,name=photo,proto3" json:"photo,omitempty"`
	DrivingNumberGivenPlace  string `protobuf:"bytes,9,opt,name=driving_number_given_place,json=drivingNumberGivenPlace,proto3" json:"driving_number_given_place,omitempty"`
	DrivingNumberGivenDate   string `protobuf:"bytes,10,opt,name=driving_number_given_date,json=drivingNumberGivenDate,proto3" json:"driving_number_given_date,omitempty"`
	DrivingNumberGivenExpire string `protobuf:"bytes,11,opt,name=driving_number_given_expire,json=drivingNumberGivenExpire,proto3" json:"driving_number_given_expire,omitempty"`
	DateOfBirth              string `protobuf:"bytes,12,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	IsBlocked                bool   `protobuf:"varint,13,opt,name=is_blocked,json=isBlocked,proto3" json:"is_blocked,omitempty"`
	Propiska                 string `protobuf:"bytes,14,opt,name=propiska,proto3" json:"propiska,omitempty"`
	PassportPinfl            string `protobuf:"bytes,15,opt,name=passport_pinfl,json=passportPinfl,proto3" json:"passport_pinfl,omitempty"`
}

func (x *UpdateClient) Reset() {
	*x = UpdateClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateClient) ProtoMessage() {}

func (x *UpdateClient) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateClient.ProtoReflect.Descriptor instead.
func (*UpdateClient) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateClient) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateClient) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UpdateClient) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UpdateClient) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UpdateClient) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *UpdateClient) GetDrivingLicenseNumber() string {
	if x != nil {
		return x.DrivingLicenseNumber
	}
	return ""
}

func (x *UpdateClient) GetPassportNumber() string {
	if x != nil {
		return x.PassportNumber
	}
	return ""
}

func (x *UpdateClient) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *UpdateClient) GetDrivingNumberGivenPlace() string {
	if x != nil {
		return x.DrivingNumberGivenPlace
	}
	return ""
}

func (x *UpdateClient) GetDrivingNumberGivenDate() string {
	if x != nil {
		return x.DrivingNumberGivenDate
	}
	return ""
}

func (x *UpdateClient) GetDrivingNumberGivenExpire() string {
	if x != nil {
		return x.DrivingNumberGivenExpire
	}
	return ""
}

func (x *UpdateClient) GetDateOfBirth() string {
	if x != nil {
		return x.DateOfBirth
	}
	return ""
}

func (x *UpdateClient) GetIsBlocked() bool {
	if x != nil {
		return x.IsBlocked
	}
	return false
}

func (x *UpdateClient) GetPropiska() string {
	if x != nil {
		return x.Propiska
	}
	return ""
}

func (x *UpdateClient) GetPassportPinfl() string {
	if x != nil {
		return x.PassportPinfl
	}
	return ""
}

type UpdatePatchClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Fields *_struct.Struct `protobuf:"bytes,2,opt,name=fields,proto3" json:"fields,omitempty"`
}

func (x *UpdatePatchClient) Reset() {
	*x = UpdatePatchClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePatchClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePatchClient) ProtoMessage() {}

func (x *UpdatePatchClient) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePatchClient.ProtoReflect.Descriptor instead.
func (*UpdatePatchClient) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePatchClient) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatePatchClient) GetFields() *_struct.Struct {
	if x != nil {
		return x.Fields
	}
	return nil
}

type GetListClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetListClientRequest) Reset() {
	*x = GetListClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListClientRequest) ProtoMessage() {}

func (x *GetListClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListClientRequest.ProtoReflect.Descriptor instead.
func (*GetListClientRequest) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{5}
}

func (x *GetListClientRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListClientRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListClientRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetListClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clients []*Client `protobuf:"bytes,1,rep,name=clients,proto3" json:"clients,omitempty"`
	Total   int64     `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetListClientResponse) Reset() {
	*x = GetListClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListClientResponse) ProtoMessage() {}

func (x *GetListClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListClientResponse.ProtoReflect.Descriptor instead.
func (*GetListClientResponse) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{6}
}

func (x *GetListClientResponse) GetClients() []*Client {
	if x != nil {
		return x.Clients
	}
	return nil
}

func (x *GetListClientResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_client_proto protoreflect.FileDescriptor

var file_client_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x05, 0x0a,
	0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x34, 0x0a, 0x16, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x14, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x3b, 0x0a, 0x1a, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x5f, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x47, 0x69, 0x76, 0x65, 0x6e, 0x50, 0x6c, 0x61, 0x63, 0x65,
	0x12, 0x39, 0x0a, 0x19, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x16, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x47, 0x69, 0x76, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3d, 0x0a, 0x1b, 0x64,
	0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x67, 0x69,
	0x76, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x18, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x47,
	0x69, 0x76, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x1d,
	0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x70, 0x69, 0x73, 0x6b, 0x61, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x70, 0x69, 0x73, 0x6b, 0x61, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x73,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x69, 0x6e, 0x66, 0x6c, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x69, 0x6e, 0x66, 0x6c,
	0x22, 0x22, 0x0a, 0x10, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0xc0, 0x04, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x16, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67,
	0x5f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x4c, 0x69,
	0x63, 0x65, 0x6e, 0x73, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x70,
	0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x3b, 0x0a, 0x1a, 0x64, 0x72,
	0x69, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x67, 0x69, 0x76,
	0x65, 0x6e, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17,
	0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x47, 0x69, 0x76,
	0x65, 0x6e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x19, 0x64, 0x72, 0x69, 0x76, 0x69,
	0x6e, 0x67, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x64, 0x72, 0x69, 0x76,
	0x69, 0x6e, 0x67, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x47, 0x69, 0x76, 0x65, 0x6e, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x3d, 0x0a, 0x1b, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x5f, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x47, 0x69, 0x76, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72,
	0x74, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66,
	0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x69, 0x73, 0x6b, 0x61,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x69, 0x73, 0x6b, 0x61,
	0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x69, 0x6e,
	0x66, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x50, 0x69, 0x6e, 0x66, 0x6c, 0x22, 0xc9, 0x04, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x34, 0x0a, 0x16, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x69, 0x63,
	0x65, 0x6e, 0x73, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x14, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x73, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x3b, 0x0a, 0x1a, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e,
	0x67, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x5f, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x64, 0x72, 0x69, 0x76,
	0x69, 0x6e, 0x67, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x47, 0x69, 0x76, 0x65, 0x6e, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x19, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x47, 0x69, 0x76, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3d,
	0x0a, 0x1b, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x5f, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x18, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x47, 0x69, 0x76, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x12, 0x22, 0x0a,
	0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74,
	0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x69, 0x73, 0x6b, 0x61, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x69, 0x73, 0x6b, 0x61, 0x12, 0x25, 0x0a, 0x0e,
	0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x69, 0x6e, 0x66, 0x6c, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x69,
	0x6e, 0x66, 0x6c, 0x22, 0x54, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74,
	0x63, 0x68, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x5c, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x5f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x30, 0x0a, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x19, 0x5a, 0x17, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_proto_rawDescOnce sync.Once
	file_client_proto_rawDescData = file_client_proto_rawDesc
)

func file_client_proto_rawDescGZIP() []byte {
	file_client_proto_rawDescOnce.Do(func() {
		file_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_proto_rawDescData)
	})
	return file_client_proto_rawDescData
}

var file_client_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_client_proto_goTypes = []interface{}{
	(*Client)(nil),                // 0: client_service.Client
	(*ClientPrimaryKey)(nil),      // 1: client_service.ClientPrimaryKey
	(*CreateClientRequest)(nil),   // 2: client_service.CreateClientRequest
	(*UpdateClient)(nil),          // 3: client_service.UpdateClient
	(*UpdatePatchClient)(nil),     // 4: client_service.UpdatePatchClient
	(*GetListClientRequest)(nil),  // 5: client_service.GetListClientRequest
	(*GetListClientResponse)(nil), // 6: client_service.GetListClientResponse
	(*_struct.Struct)(nil),        // 7: google.protobuf.Struct
}
var file_client_proto_depIdxs = []int32{
	7, // 0: client_service.UpdatePatchClient.fields:type_name -> google.protobuf.Struct
	0, // 1: client_service.GetListClientResponse.clients:type_name -> client_service.Client
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_client_proto_init() }
func file_client_proto_init() {
	if File_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Client); i {
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
		file_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientPrimaryKey); i {
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
		file_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClientRequest); i {
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
		file_client_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateClient); i {
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
		file_client_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePatchClient); i {
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
		file_client_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListClientRequest); i {
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
		file_client_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListClientResponse); i {
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
			RawDescriptor: file_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_client_proto_goTypes,
		DependencyIndexes: file_client_proto_depIdxs,
		MessageInfos:      file_client_proto_msgTypes,
	}.Build()
	File_client_proto = out.File
	file_client_proto_rawDesc = nil
	file_client_proto_goTypes = nil
	file_client_proto_depIdxs = nil
}
