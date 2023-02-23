// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: wallet/v1/wallet.proto

package walletv1

import (
	v1 "github.com/orbyc/core/domain/transaction/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/api/httpbody"
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

type GetAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAddressRequest) Reset() {
	*x = GetAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_v1_wallet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAddressRequest) ProtoMessage() {}

func (x *GetAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_v1_wallet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAddressRequest.ProtoReflect.Descriptor instead.
func (*GetAddressRequest) Descriptor() ([]byte, []int) {
	return file_wallet_v1_wallet_proto_rawDescGZIP(), []int{0}
}

type GetAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GetAddressResponse) Reset() {
	*x = GetAddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_v1_wallet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAddressResponse) ProtoMessage() {}

func (x *GetAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_v1_wallet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAddressResponse.ProtoReflect.Descriptor instead.
func (*GetAddressResponse) Descriptor() ([]byte, []int) {
	return file_wallet_v1_wallet_proto_rawDescGZIP(), []int{1}
}

func (x *GetAddressResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type SignTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *v1.Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *SignTransactionRequest) Reset() {
	*x = SignTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_v1_wallet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignTransactionRequest) ProtoMessage() {}

func (x *SignTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_v1_wallet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignTransactionRequest.ProtoReflect.Descriptor instead.
func (*SignTransactionRequest) Descriptor() ([]byte, []int) {
	return file_wallet_v1_wallet_proto_rawDescGZIP(), []int{2}
}

func (x *SignTransactionRequest) GetTransaction() *v1.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

type SignTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SignTransactionResponse) Reset() {
	*x = SignTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_v1_wallet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignTransactionResponse) ProtoMessage() {}

func (x *SignTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_v1_wallet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignTransactionResponse.ProtoReflect.Descriptor instead.
func (*SignTransactionResponse) Descriptor() ([]byte, []int) {
	return file_wallet_v1_wallet_proto_rawDescGZIP(), []int{3}
}

func (x *SignTransactionResponse) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

type PushTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction    *v1.Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Authorizations []string        `protobuf:"bytes,2,rep,name=authorizations,proto3" json:"authorizations,omitempty"`
}

func (x *PushTransactionRequest) Reset() {
	*x = PushTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_v1_wallet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushTransactionRequest) ProtoMessage() {}

func (x *PushTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_v1_wallet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushTransactionRequest.ProtoReflect.Descriptor instead.
func (*PushTransactionRequest) Descriptor() ([]byte, []int) {
	return file_wallet_v1_wallet_proto_rawDescGZIP(), []int{4}
}

func (x *PushTransactionRequest) GetTransaction() *v1.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *PushTransactionRequest) GetAuthorizations() []string {
	if x != nil {
		return x.Authorizations
	}
	return nil
}

type PushTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *PushTransactionResponse) Reset() {
	*x = PushTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_v1_wallet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushTransactionResponse) ProtoMessage() {}

func (x *PushTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_v1_wallet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushTransactionResponse.ProtoReflect.Descriptor instead.
func (*PushTransactionResponse) Descriptor() ([]byte, []int) {
	return file_wallet_v1_wallet_proto_rawDescGZIP(), []int{5}
}

func (x *PushTransactionResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_wallet_v1_wallet_proto protoreflect.FileDescriptor

var file_wallet_v1_wallet_proto_rawDesc = []byte{
	0x0a, 0x16, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x74,
	0x74, 0x70, 0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x13,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x2e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x57, 0x0a, 0x16, 0x53, 0x69, 0x67, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a,
	0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x37, 0x0a, 0x17,
	0x53, 0x69, 0x67, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x7f, 0x0a, 0x16, 0x50, 0x75, 0x73, 0x68, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3d, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26,
	0x0a, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x33, 0x0a, 0x17, 0x50, 0x75, 0x73, 0x68, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xc8, 0x02, 0x0a, 0x0d,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x2e, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a,
	0x12, 0x08, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x67, 0x0a, 0x0f, 0x53, 0x69,
	0x67, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67,
	0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x07, 0x12, 0x05, 0x2f, 0x73,
	0x69, 0x67, 0x6e, 0x12, 0x71, 0x0a, 0x0f, 0x50, 0x75, 0x73, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x1a, 0x0c, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x92, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x62, 0x79, 0x63, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x3b,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x57, 0x58, 0x58, 0xaa, 0x02,
	0x09, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x57, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0a, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_wallet_v1_wallet_proto_rawDescOnce sync.Once
	file_wallet_v1_wallet_proto_rawDescData = file_wallet_v1_wallet_proto_rawDesc
)

func file_wallet_v1_wallet_proto_rawDescGZIP() []byte {
	file_wallet_v1_wallet_proto_rawDescOnce.Do(func() {
		file_wallet_v1_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(file_wallet_v1_wallet_proto_rawDescData)
	})
	return file_wallet_v1_wallet_proto_rawDescData
}

var file_wallet_v1_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_wallet_v1_wallet_proto_goTypes = []interface{}{
	(*GetAddressRequest)(nil),       // 0: wallet.v1.GetAddressRequest
	(*GetAddressResponse)(nil),      // 1: wallet.v1.GetAddressResponse
	(*SignTransactionRequest)(nil),  // 2: wallet.v1.SignTransactionRequest
	(*SignTransactionResponse)(nil), // 3: wallet.v1.SignTransactionResponse
	(*PushTransactionRequest)(nil),  // 4: wallet.v1.PushTransactionRequest
	(*PushTransactionResponse)(nil), // 5: wallet.v1.PushTransactionResponse
	(*v1.Transaction)(nil),          // 6: transaction.v1.Transaction
}
var file_wallet_v1_wallet_proto_depIdxs = []int32{
	6, // 0: wallet.v1.SignTransactionRequest.transaction:type_name -> transaction.v1.Transaction
	6, // 1: wallet.v1.PushTransactionRequest.transaction:type_name -> transaction.v1.Transaction
	0, // 2: wallet.v1.WalletService.GetAddress:input_type -> wallet.v1.GetAddressRequest
	2, // 3: wallet.v1.WalletService.SignTransaction:input_type -> wallet.v1.SignTransactionRequest
	4, // 4: wallet.v1.WalletService.PushTransaction:input_type -> wallet.v1.PushTransactionRequest
	1, // 5: wallet.v1.WalletService.GetAddress:output_type -> wallet.v1.GetAddressResponse
	3, // 6: wallet.v1.WalletService.SignTransaction:output_type -> wallet.v1.SignTransactionResponse
	5, // 7: wallet.v1.WalletService.PushTransaction:output_type -> wallet.v1.PushTransactionResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_wallet_v1_wallet_proto_init() }
func file_wallet_v1_wallet_proto_init() {
	if File_wallet_v1_wallet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wallet_v1_wallet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAddressRequest); i {
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
		file_wallet_v1_wallet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAddressResponse); i {
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
		file_wallet_v1_wallet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignTransactionRequest); i {
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
		file_wallet_v1_wallet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignTransactionResponse); i {
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
		file_wallet_v1_wallet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushTransactionRequest); i {
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
		file_wallet_v1_wallet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushTransactionResponse); i {
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
			RawDescriptor: file_wallet_v1_wallet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wallet_v1_wallet_proto_goTypes,
		DependencyIndexes: file_wallet_v1_wallet_proto_depIdxs,
		MessageInfos:      file_wallet_v1_wallet_proto_msgTypes,
	}.Build()
	File_wallet_v1_wallet_proto = out.File
	file_wallet_v1_wallet_proto_rawDesc = nil
	file_wallet_v1_wallet_proto_goTypes = nil
	file_wallet_v1_wallet_proto_depIdxs = nil
}
