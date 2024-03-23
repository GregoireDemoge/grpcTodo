// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.1
// source: proto/todo/todo.proto

package todo

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

type GetTodoParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTodoParams) Reset() {
	*x = GetTodoParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todo_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTodoParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodoParams) ProtoMessage() {}

func (x *GetTodoParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodoParams.ProtoReflect.Descriptor instead.
func (*GetTodoParams) Descriptor() ([]byte, []int) {
	return file_proto_todo_todo_proto_rawDescGZIP(), []int{0}
}

type AddTodoParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task string `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *AddTodoParams) Reset() {
	*x = AddTodoParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todo_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTodoParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTodoParams) ProtoMessage() {}

func (x *AddTodoParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTodoParams.ProtoReflect.Descriptor instead.
func (*AddTodoParams) Descriptor() ([]byte, []int) {
	return file_proto_todo_todo_proto_rawDescGZIP(), []int{1}
}

func (x *AddTodoParams) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

type DeleteTodoParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTodoParams) Reset() {
	*x = DeleteTodoParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todo_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTodoParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTodoParams) ProtoMessage() {}

func (x *DeleteTodoParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTodoParams.ProtoReflect.Descriptor instead.
func (*DeleteTodoParams) Descriptor() ([]byte, []int) {
	return file_proto_todo_todo_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteTodoParams) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type MarkTodoParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Completed bool   `protobuf:"varint,2,opt,name=completed,proto3" json:"completed,omitempty"`
}

func (x *MarkTodoParams) Reset() {
	*x = MarkTodoParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todo_todo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkTodoParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkTodoParams) ProtoMessage() {}

func (x *MarkTodoParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_todo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkTodoParams.ProtoReflect.Descriptor instead.
func (*MarkTodoParams) Descriptor() ([]byte, []int) {
	return file_proto_todo_todo_proto_rawDescGZIP(), []int{3}
}

func (x *MarkTodoParams) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MarkTodoParams) GetCompleted() bool {
	if x != nil {
		return x.Completed
	}
	return false
}

type TodoObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Task      string `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	Completed bool   `protobuf:"varint,3,opt,name=completed,proto3" json:"completed,omitempty"`
	Deleted   bool   `protobuf:"varint,4,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *TodoObject) Reset() {
	*x = TodoObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todo_todo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoObject) ProtoMessage() {}

func (x *TodoObject) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_todo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoObject.ProtoReflect.Descriptor instead.
func (*TodoObject) Descriptor() ([]byte, []int) {
	return file_proto_todo_todo_proto_rawDescGZIP(), []int{4}
}

func (x *TodoObject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TodoObject) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

func (x *TodoObject) GetCompleted() bool {
	if x != nil {
		return x.Completed
	}
	return false
}

func (x *TodoObject) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

type TodoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *TodoObject `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *TodoResponse) Reset() {
	*x = TodoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todo_todo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoResponse) ProtoMessage() {}

func (x *TodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_todo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoResponse.ProtoReflect.Descriptor instead.
func (*TodoResponse) Descriptor() ([]byte, []int) {
	return file_proto_todo_todo_proto_rawDescGZIP(), []int{5}
}

func (x *TodoResponse) GetTodo() *TodoObject {
	if x != nil {
		return x.Todo
	}
	return nil
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todo_todo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_todo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_proto_todo_todo_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type MarkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MarkResponse) Reset() {
	*x = MarkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todo_todo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkResponse) ProtoMessage() {}

func (x *MarkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_todo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkResponse.ProtoReflect.Descriptor instead.
func (*MarkResponse) Descriptor() ([]byte, []int) {
	return file_proto_todo_todo_proto_rawDescGZIP(), []int{7}
}

func (x *MarkResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_todo_todo_proto protoreflect.FileDescriptor

var file_proto_todo_todo_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x74, 0x6f, 0x64,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x0f, 0x0a,
	0x0d, 0x67, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x23,
	0x0a, 0x0d, 0x61, 0x64, 0x64, 0x54, 0x6f, 0x64, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x61, 0x73, 0x6b, 0x22, 0x22, 0x0a, 0x10, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64,
	0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3e, 0x0a, 0x0e, 0x6d, 0x61, 0x72, 0x6b, 0x54,
	0x6f, 0x64, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x68, 0x0a, 0x0a, 0x74, 0x6f, 0x64, 0x6f, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x22, 0x34, 0x0a, 0x0c, 0x74, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x2a, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x28, 0x0a, 0x0c, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xf0, 0x01,
	0x0a, 0x0b, 0x74, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x13, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e,
	0x61, 0x64, 0x64, 0x54, 0x6f, 0x64, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x10, 0x2e,
	0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22,
	0x00, 0x12, 0x3c, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12,
	0x16, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64,
	0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x14, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x37, 0x0a, 0x08, 0x67, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x13, 0x2e, 0x74, 0x6f,
	0x64, 0x6f, 0x2e, 0x67, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x1a, 0x12, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x36, 0x0a, 0x08, 0x6d, 0x61, 0x72, 0x6b,
	0x54, 0x6f, 0x64, 0x6f, 0x12, 0x14, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x6d, 0x61, 0x72, 0x6b,
	0x54, 0x6f, 0x64, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x12, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x1d, 0x5a, 0x1b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x6f, 0x64, 0x6f, 0x47, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_todo_todo_proto_rawDescOnce sync.Once
	file_proto_todo_todo_proto_rawDescData = file_proto_todo_todo_proto_rawDesc
)

func file_proto_todo_todo_proto_rawDescGZIP() []byte {
	file_proto_todo_todo_proto_rawDescOnce.Do(func() {
		file_proto_todo_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_todo_todo_proto_rawDescData)
	})
	return file_proto_todo_todo_proto_rawDescData
}

var file_proto_todo_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_todo_todo_proto_goTypes = []interface{}{
	(*GetTodoParams)(nil),    // 0: todo.getTodoParams
	(*AddTodoParams)(nil),    // 1: todo.addTodoParams
	(*DeleteTodoParams)(nil), // 2: todo.deleteTodoParams
	(*MarkTodoParams)(nil),   // 3: todo.markTodoParams
	(*TodoObject)(nil),       // 4: todo.todoObject
	(*TodoResponse)(nil),     // 5: todo.todoResponse
	(*DeleteResponse)(nil),   // 6: todo.deleteResponse
	(*MarkResponse)(nil),     // 7: todo.markResponse
}
var file_proto_todo_todo_proto_depIdxs = []int32{
	4, // 0: todo.todoResponse.todo:type_name -> todo.todoObject
	1, // 1: todo.todoService.addTodo:input_type -> todo.addTodoParams
	2, // 2: todo.todoService.deleteTodo:input_type -> todo.deleteTodoParams
	0, // 3: todo.todoService.getTodos:input_type -> todo.getTodoParams
	3, // 4: todo.todoService.markTodo:input_type -> todo.markTodoParams
	4, // 5: todo.todoService.addTodo:output_type -> todo.todoObject
	6, // 6: todo.todoService.deleteTodo:output_type -> todo.deleteResponse
	5, // 7: todo.todoService.getTodos:output_type -> todo.todoResponse
	7, // 8: todo.todoService.markTodo:output_type -> todo.markResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_todo_todo_proto_init() }
func file_proto_todo_todo_proto_init() {
	if File_proto_todo_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_todo_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTodoParams); i {
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
		file_proto_todo_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTodoParams); i {
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
		file_proto_todo_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTodoParams); i {
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
		file_proto_todo_todo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkTodoParams); i {
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
		file_proto_todo_todo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoObject); i {
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
		file_proto_todo_todo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoResponse); i {
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
		file_proto_todo_todo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
		file_proto_todo_todo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkResponse); i {
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
			RawDescriptor: file_proto_todo_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_todo_todo_proto_goTypes,
		DependencyIndexes: file_proto_todo_todo_proto_depIdxs,
		MessageInfos:      file_proto_todo_todo_proto_msgTypes,
	}.Build()
	File_proto_todo_todo_proto = out.File
	file_proto_todo_todo_proto_rawDesc = nil
	file_proto_todo_todo_proto_goTypes = nil
	file_proto_todo_todo_proto_depIdxs = nil
}
