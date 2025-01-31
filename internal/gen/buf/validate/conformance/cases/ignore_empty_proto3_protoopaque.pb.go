// Copyright 2023-2024 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: buf/validate/conformance/cases/ignore_empty_proto3.proto

//go:build protoopaque

package cases

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IgnoreEmptyProto3Scalar struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val int32                  `protobuf:"varint,1,opt,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *IgnoreEmptyProto3Scalar) Reset() {
	*x = IgnoreEmptyProto3Scalar{}
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IgnoreEmptyProto3Scalar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IgnoreEmptyProto3Scalar) ProtoMessage() {}

func (x *IgnoreEmptyProto3Scalar) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *IgnoreEmptyProto3Scalar) GetVal() int32 {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return 0
}

func (x *IgnoreEmptyProto3Scalar) SetVal(v int32) {
	x.xxx_hidden_Val = v
}

type IgnoreEmptyProto3Scalar_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val int32
}

func (b0 IgnoreEmptyProto3Scalar_builder) Build() *IgnoreEmptyProto3Scalar {
	m0 := &IgnoreEmptyProto3Scalar{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type IgnoreEmptyProto3OptionalScalar struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val         int32                  `protobuf:"varint,1,opt,name=val,proto3,oneof"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *IgnoreEmptyProto3OptionalScalar) Reset() {
	*x = IgnoreEmptyProto3OptionalScalar{}
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IgnoreEmptyProto3OptionalScalar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IgnoreEmptyProto3OptionalScalar) ProtoMessage() {}

func (x *IgnoreEmptyProto3OptionalScalar) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *IgnoreEmptyProto3OptionalScalar) GetVal() int32 {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return 0
}

func (x *IgnoreEmptyProto3OptionalScalar) SetVal(v int32) {
	x.xxx_hidden_Val = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 1)
}

func (x *IgnoreEmptyProto3OptionalScalar) HasVal() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *IgnoreEmptyProto3OptionalScalar) ClearVal() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Val = 0
}

type IgnoreEmptyProto3OptionalScalar_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *int32
}

func (b0 IgnoreEmptyProto3OptionalScalar_builder) Build() *IgnoreEmptyProto3OptionalScalar {
	m0 := &IgnoreEmptyProto3OptionalScalar{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Val != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 1)
		x.xxx_hidden_Val = *b.Val
	}
	return m0
}

type IgnoreEmptyProto3Message struct {
	state          protoimpl.MessageState        `protogen:"opaque.v1"`
	xxx_hidden_Val *IgnoreEmptyProto3Message_Msg `protobuf:"bytes,1,opt,name=val,proto3,oneof"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *IgnoreEmptyProto3Message) Reset() {
	*x = IgnoreEmptyProto3Message{}
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IgnoreEmptyProto3Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IgnoreEmptyProto3Message) ProtoMessage() {}

func (x *IgnoreEmptyProto3Message) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *IgnoreEmptyProto3Message) GetVal() *IgnoreEmptyProto3Message_Msg {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *IgnoreEmptyProto3Message) SetVal(v *IgnoreEmptyProto3Message_Msg) {
	x.xxx_hidden_Val = v
}

func (x *IgnoreEmptyProto3Message) HasVal() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Val != nil
}

func (x *IgnoreEmptyProto3Message) ClearVal() {
	x.xxx_hidden_Val = nil
}

type IgnoreEmptyProto3Message_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *IgnoreEmptyProto3Message_Msg
}

func (b0 IgnoreEmptyProto3Message_builder) Build() *IgnoreEmptyProto3Message {
	m0 := &IgnoreEmptyProto3Message{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type IgnoreEmptyProto3Oneof struct {
	state         protoimpl.MessageState     `protogen:"opaque.v1"`
	xxx_hidden_O  isIgnoreEmptyProto3Oneof_O `protobuf_oneof:"o"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IgnoreEmptyProto3Oneof) Reset() {
	*x = IgnoreEmptyProto3Oneof{}
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IgnoreEmptyProto3Oneof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IgnoreEmptyProto3Oneof) ProtoMessage() {}

func (x *IgnoreEmptyProto3Oneof) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *IgnoreEmptyProto3Oneof) GetVal() int32 {
	if x != nil {
		if x, ok := x.xxx_hidden_O.(*ignoreEmptyProto3Oneof_Val); ok {
			return x.Val
		}
	}
	return 0
}

func (x *IgnoreEmptyProto3Oneof) SetVal(v int32) {
	x.xxx_hidden_O = &ignoreEmptyProto3Oneof_Val{v}
}

func (x *IgnoreEmptyProto3Oneof) HasO() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_O != nil
}

func (x *IgnoreEmptyProto3Oneof) HasVal() bool {
	if x == nil {
		return false
	}
	_, ok := x.xxx_hidden_O.(*ignoreEmptyProto3Oneof_Val)
	return ok
}

func (x *IgnoreEmptyProto3Oneof) ClearO() {
	x.xxx_hidden_O = nil
}

func (x *IgnoreEmptyProto3Oneof) ClearVal() {
	if _, ok := x.xxx_hidden_O.(*ignoreEmptyProto3Oneof_Val); ok {
		x.xxx_hidden_O = nil
	}
}

const IgnoreEmptyProto3Oneof_O_not_set_case case_IgnoreEmptyProto3Oneof_O = 0
const IgnoreEmptyProto3Oneof_Val_case case_IgnoreEmptyProto3Oneof_O = 1

func (x *IgnoreEmptyProto3Oneof) WhichO() case_IgnoreEmptyProto3Oneof_O {
	if x == nil {
		return IgnoreEmptyProto3Oneof_O_not_set_case
	}
	switch x.xxx_hidden_O.(type) {
	case *ignoreEmptyProto3Oneof_Val:
		return IgnoreEmptyProto3Oneof_Val_case
	default:
		return IgnoreEmptyProto3Oneof_O_not_set_case
	}
}

type IgnoreEmptyProto3Oneof_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Fields of oneof xxx_hidden_O:
	Val *int32
	// -- end of xxx_hidden_O
}

func (b0 IgnoreEmptyProto3Oneof_builder) Build() *IgnoreEmptyProto3Oneof {
	m0 := &IgnoreEmptyProto3Oneof{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Val != nil {
		x.xxx_hidden_O = &ignoreEmptyProto3Oneof_Val{*b.Val}
	}
	return m0
}

type case_IgnoreEmptyProto3Oneof_O protoreflect.FieldNumber

func (x case_IgnoreEmptyProto3Oneof_O) String() string {
	md := file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[3].Descriptor()
	if x == 0 {
		return "not set"
	}
	return protoimpl.X.MessageFieldStringOf(md, protoreflect.FieldNumber(x))
}

type isIgnoreEmptyProto3Oneof_O interface {
	isIgnoreEmptyProto3Oneof_O()
}

type ignoreEmptyProto3Oneof_Val struct {
	Val int32 `protobuf:"varint,1,opt,name=val,proto3,oneof"`
}

func (*ignoreEmptyProto3Oneof_Val) isIgnoreEmptyProto3Oneof_O() {}

type IgnoreEmptyProto3Repeated struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val []int32                `protobuf:"varint,1,rep,packed,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *IgnoreEmptyProto3Repeated) Reset() {
	*x = IgnoreEmptyProto3Repeated{}
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IgnoreEmptyProto3Repeated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IgnoreEmptyProto3Repeated) ProtoMessage() {}

func (x *IgnoreEmptyProto3Repeated) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *IgnoreEmptyProto3Repeated) GetVal() []int32 {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *IgnoreEmptyProto3Repeated) SetVal(v []int32) {
	x.xxx_hidden_Val = v
}

type IgnoreEmptyProto3Repeated_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val []int32
}

func (b0 IgnoreEmptyProto3Repeated_builder) Build() *IgnoreEmptyProto3Repeated {
	m0 := &IgnoreEmptyProto3Repeated{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type IgnoreEmptyProto3Map struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val map[int32]int32        `protobuf:"bytes,1,rep,name=val,proto3" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *IgnoreEmptyProto3Map) Reset() {
	*x = IgnoreEmptyProto3Map{}
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IgnoreEmptyProto3Map) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IgnoreEmptyProto3Map) ProtoMessage() {}

func (x *IgnoreEmptyProto3Map) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *IgnoreEmptyProto3Map) GetVal() map[int32]int32 {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *IgnoreEmptyProto3Map) SetVal(v map[int32]int32) {
	x.xxx_hidden_Val = v
}

type IgnoreEmptyProto3Map_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val map[int32]int32
}

func (b0 IgnoreEmptyProto3Map_builder) Build() *IgnoreEmptyProto3Map {
	m0 := &IgnoreEmptyProto3Map{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type IgnoreEmptyRepeatedItems struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val []int32                `protobuf:"varint,1,rep,packed,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *IgnoreEmptyRepeatedItems) Reset() {
	*x = IgnoreEmptyRepeatedItems{}
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IgnoreEmptyRepeatedItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IgnoreEmptyRepeatedItems) ProtoMessage() {}

func (x *IgnoreEmptyRepeatedItems) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *IgnoreEmptyRepeatedItems) GetVal() []int32 {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *IgnoreEmptyRepeatedItems) SetVal(v []int32) {
	x.xxx_hidden_Val = v
}

type IgnoreEmptyRepeatedItems_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val []int32
}

func (b0 IgnoreEmptyRepeatedItems_builder) Build() *IgnoreEmptyRepeatedItems {
	m0 := &IgnoreEmptyRepeatedItems{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type IgnoreEmptyMapPairs struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val map[string]int32       `protobuf:"bytes,1,rep,name=val,proto3" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *IgnoreEmptyMapPairs) Reset() {
	*x = IgnoreEmptyMapPairs{}
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IgnoreEmptyMapPairs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IgnoreEmptyMapPairs) ProtoMessage() {}

func (x *IgnoreEmptyMapPairs) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *IgnoreEmptyMapPairs) GetVal() map[string]int32 {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *IgnoreEmptyMapPairs) SetVal(v map[string]int32) {
	x.xxx_hidden_Val = v
}

type IgnoreEmptyMapPairs_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val map[string]int32
}

func (b0 IgnoreEmptyMapPairs_builder) Build() *IgnoreEmptyMapPairs {
	m0 := &IgnoreEmptyMapPairs{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type IgnoreEmptyProto3Message_Msg struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val string                 `protobuf:"bytes,1,opt,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *IgnoreEmptyProto3Message_Msg) Reset() {
	*x = IgnoreEmptyProto3Message_Msg{}
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IgnoreEmptyProto3Message_Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IgnoreEmptyProto3Message_Msg) ProtoMessage() {}

func (x *IgnoreEmptyProto3Message_Msg) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *IgnoreEmptyProto3Message_Msg) GetVal() string {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return ""
}

func (x *IgnoreEmptyProto3Message_Msg) SetVal(v string) {
	x.xxx_hidden_Val = v
}

type IgnoreEmptyProto3Message_Msg_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val string
}

func (b0 IgnoreEmptyProto3Message_Msg_builder) Build() *IgnoreEmptyProto3Message_Msg {
	m0 := &IgnoreEmptyProto3Message_Msg{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

var File_buf_validate_conformance_cases_ignore_empty_proto3_proto protoreflect.FileDescriptor

var file_buf_validate_conformance_cases_ignore_empty_proto3_proto_rawDesc = string([]byte{
	0x0a, 0x38, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73,
	0x2f, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x62, 0x75, 0x66, 0x2e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x17, 0x49, 0x67, 0x6e, 0x6f, 0x72,
	0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x53, 0x63, 0x61, 0x6c,
	0x61, 0x72, 0x12, 0x1c, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x0a, 0xba, 0x48, 0x07, 0xd8, 0x01, 0x01, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x03, 0x76, 0x61, 0x6c,
	0x22, 0x4c, 0x0a, 0x1f, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x33, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x63, 0x61,
	0x6c, 0x61, 0x72, 0x12, 0x21, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x0a, 0xba, 0x48, 0x07, 0xd8, 0x01, 0x01, 0x1a, 0x02, 0x20, 0x00, 0x48, 0x00, 0x52, 0x03,
	0x76, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x76, 0x61, 0x6c, 0x22, 0xd4,
	0x01, 0x0a, 0x18, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x33, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x96, 0x01, 0x0a, 0x03,
	0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x62, 0x75, 0x66, 0x2e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x49, 0x67, 0x6e, 0x6f, 0x72,
	0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x42, 0x41, 0xba, 0x48, 0x3e, 0xba, 0x01, 0x38, 0x0a,
	0x1b, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x06, 0x66, 0x6f,
	0x6f, 0x62, 0x61, 0x72, 0x1a, 0x11, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x76, 0x61, 0x6c, 0x20, 0x3d,
	0x3d, 0x20, 0x27, 0x66, 0x6f, 0x6f, 0x27, 0xd8, 0x01, 0x01, 0x48, 0x00, 0x52, 0x03, 0x76, 0x61,
	0x6c, 0x88, 0x01, 0x01, 0x1a, 0x17, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x76,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x42, 0x06, 0x0a,
	0x04, 0x5f, 0x76, 0x61, 0x6c, 0x22, 0x3d, 0x0a, 0x16, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x12,
	0x1e, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0xba, 0x48,
	0x07, 0xd8, 0x01, 0x01, 0x1a, 0x02, 0x20, 0x00, 0x48, 0x00, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x42,
	0x03, 0x0a, 0x01, 0x6f, 0x22, 0x3a, 0x0a, 0x19, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x1d, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x42, 0x0b,
	0xba, 0x48, 0x08, 0xd8, 0x01, 0x01, 0x92, 0x01, 0x02, 0x08, 0x03, 0x52, 0x03, 0x76, 0x61, 0x6c,
	0x22, 0xac, 0x01, 0x0a, 0x14, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x4d, 0x61, 0x70, 0x12, 0x5c, 0x0a, 0x03, 0x76, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63,
	0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x4d, 0x61, 0x70, 0x2e, 0x56, 0x61, 0x6c,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x0b, 0xba, 0x48, 0x08, 0xd8, 0x01, 0x01, 0x9a, 0x01, 0x02,
	0x08, 0x03, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x1a, 0x36, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x3d, 0x0a, 0x18, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x21, 0x0a, 0x03, 0x76,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x42, 0x0f, 0xba, 0x48, 0x0c, 0x92, 0x01, 0x09,
	0x22, 0x07, 0xd8, 0x01, 0x01, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0xb7,
	0x01, 0x0a, 0x13, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x61,
	0x70, 0x50, 0x61, 0x69, 0x72, 0x73, 0x12, 0x68, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63,
	0x61, 0x73, 0x65, 0x73, 0x2e, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x4d, 0x61, 0x70, 0x50, 0x61, 0x69, 0x72, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x42, 0x18, 0xba, 0x48, 0x15, 0x9a, 0x01, 0x12, 0x22, 0x07, 0xd8, 0x01, 0x01, 0x72, 0x02,
	0x10, 0x03, 0x2a, 0x07, 0xd8, 0x01, 0x01, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x03, 0x76, 0x61, 0x6c,
	0x1a, 0x36, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0xaa, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d,
	0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x42,
	0x16, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x33, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2d, 0x67, 0x6f, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x75, 0x66,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73, 0xa2, 0x02, 0x04, 0x42, 0x56,
	0x43, 0x43, 0xaa, 0x02, 0x1e, 0x42, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x61,
	0x73, 0x65, 0x73, 0xca, 0x02, 0x1e, 0x42, 0x75, 0x66, 0x5c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x5c, 0x43,
	0x61, 0x73, 0x65, 0x73, 0xe2, 0x02, 0x2a, 0x42, 0x75, 0x66, 0x5c, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x5c,
	0x43, 0x61, 0x73, 0x65, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x21, 0x42, 0x75, 0x66, 0x3a, 0x3a, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x3a, 0x3a,
	0x43, 0x61, 0x73, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_buf_validate_conformance_cases_ignore_empty_proto3_proto_goTypes = []any{
	(*IgnoreEmptyProto3Scalar)(nil),         // 0: buf.validate.conformance.cases.IgnoreEmptyProto3Scalar
	(*IgnoreEmptyProto3OptionalScalar)(nil), // 1: buf.validate.conformance.cases.IgnoreEmptyProto3OptionalScalar
	(*IgnoreEmptyProto3Message)(nil),        // 2: buf.validate.conformance.cases.IgnoreEmptyProto3Message
	(*IgnoreEmptyProto3Oneof)(nil),          // 3: buf.validate.conformance.cases.IgnoreEmptyProto3Oneof
	(*IgnoreEmptyProto3Repeated)(nil),       // 4: buf.validate.conformance.cases.IgnoreEmptyProto3Repeated
	(*IgnoreEmptyProto3Map)(nil),            // 5: buf.validate.conformance.cases.IgnoreEmptyProto3Map
	(*IgnoreEmptyRepeatedItems)(nil),        // 6: buf.validate.conformance.cases.IgnoreEmptyRepeatedItems
	(*IgnoreEmptyMapPairs)(nil),             // 7: buf.validate.conformance.cases.IgnoreEmptyMapPairs
	(*IgnoreEmptyProto3Message_Msg)(nil),    // 8: buf.validate.conformance.cases.IgnoreEmptyProto3Message.Msg
	nil,                                     // 9: buf.validate.conformance.cases.IgnoreEmptyProto3Map.ValEntry
	nil,                                     // 10: buf.validate.conformance.cases.IgnoreEmptyMapPairs.ValEntry
}
var file_buf_validate_conformance_cases_ignore_empty_proto3_proto_depIdxs = []int32{
	8,  // 0: buf.validate.conformance.cases.IgnoreEmptyProto3Message.val:type_name -> buf.validate.conformance.cases.IgnoreEmptyProto3Message.Msg
	9,  // 1: buf.validate.conformance.cases.IgnoreEmptyProto3Map.val:type_name -> buf.validate.conformance.cases.IgnoreEmptyProto3Map.ValEntry
	10, // 2: buf.validate.conformance.cases.IgnoreEmptyMapPairs.val:type_name -> buf.validate.conformance.cases.IgnoreEmptyMapPairs.ValEntry
	3,  // [3:3] is the sub-list for method output_type
	3,  // [3:3] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_buf_validate_conformance_cases_ignore_empty_proto3_proto_init() }
func file_buf_validate_conformance_cases_ignore_empty_proto3_proto_init() {
	if File_buf_validate_conformance_cases_ignore_empty_proto3_proto != nil {
		return
	}
	file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[1].OneofWrappers = []any{}
	file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[2].OneofWrappers = []any{}
	file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[3].OneofWrappers = []any{
		(*ignoreEmptyProto3Oneof_Val)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_buf_validate_conformance_cases_ignore_empty_proto3_proto_rawDesc), len(file_buf_validate_conformance_cases_ignore_empty_proto3_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_validate_conformance_cases_ignore_empty_proto3_proto_goTypes,
		DependencyIndexes: file_buf_validate_conformance_cases_ignore_empty_proto3_proto_depIdxs,
		MessageInfos:      file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes,
	}.Build()
	File_buf_validate_conformance_cases_ignore_empty_proto3_proto = out.File
	file_buf_validate_conformance_cases_ignore_empty_proto3_proto_goTypes = nil
	file_buf_validate_conformance_cases_ignore_empty_proto3_proto_depIdxs = nil
}
