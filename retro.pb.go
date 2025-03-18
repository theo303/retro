// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.1
// source: retro.proto

package retro

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

type Sticky struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner         string                 `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	SelectedBy    *string                `protobuf:"bytes,3,opt,name=selectedBy,proto3,oneof" json:"selectedBy,omitempty"`
	X             int64                  `protobuf:"varint,4,opt,name=X,proto3" json:"X,omitempty"`
	Y             int64                  `protobuf:"varint,5,opt,name=Y,proto3" json:"Y,omitempty"`
	Width         int64                  `protobuf:"varint,6,opt,name=Width,proto3" json:"Width,omitempty"`
	Height        int64                  `protobuf:"varint,7,opt,name=Height,proto3" json:"Height,omitempty"`
	Content       string                 `protobuf:"bytes,8,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Sticky) Reset() {
	*x = Sticky{}
	mi := &file_retro_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Sticky) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sticky) ProtoMessage() {}

func (x *Sticky) ProtoReflect() protoreflect.Message {
	mi := &file_retro_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sticky.ProtoReflect.Descriptor instead.
func (*Sticky) Descriptor() ([]byte, []int) {
	return file_retro_proto_rawDescGZIP(), []int{0}
}

func (x *Sticky) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Sticky) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Sticky) GetSelectedBy() string {
	if x != nil && x.SelectedBy != nil {
		return *x.SelectedBy
	}
	return ""
}

func (x *Sticky) GetX() int64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Sticky) GetY() int64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Sticky) GetWidth() int64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Sticky) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Sticky) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	HasSelected   *string                `protobuf:"bytes,2,opt,name=hasSelected,proto3,oneof" json:"hasSelected,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_retro_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_retro_proto_msgTypes[1]
	if x != nil {
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
	return file_retro_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetHasSelected() string {
	if x != nil && x.HasSelected != nil {
		return *x.HasSelected
	}
	return ""
}

type State struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Stickies      []*Sticky              `protobuf:"bytes,1,rep,name=stickies,proto3" json:"stickies,omitempty"`
	Users         map[string]*User       `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *State) Reset() {
	*x = State{}
	mi := &file_retro_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_retro_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use State.ProtoReflect.Descriptor instead.
func (*State) Descriptor() ([]byte, []int) {
	return file_retro_proto_rawDescGZIP(), []int{2}
}

func (x *State) GetStickies() []*Sticky {
	if x != nil {
		return x.Stickies
	}
	return nil
}

func (x *State) GetUsers() map[string]*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type SelectAction struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StickyID      string                 `protobuf:"bytes,1,opt,name=StickyID,proto3" json:"StickyID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SelectAction) Reset() {
	*x = SelectAction{}
	mi := &file_retro_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SelectAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectAction) ProtoMessage() {}

func (x *SelectAction) ProtoReflect() protoreflect.Message {
	mi := &file_retro_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectAction.ProtoReflect.Descriptor instead.
func (*SelectAction) Descriptor() ([]byte, []int) {
	return file_retro_proto_rawDescGZIP(), []int{3}
}

func (x *SelectAction) GetStickyID() string {
	if x != nil {
		return x.StickyID
	}
	return ""
}

type AddAction struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	X             int64                  `protobuf:"varint,1,opt,name=X,proto3" json:"X,omitempty"`
	Y             int64                  `protobuf:"varint,2,opt,name=Y,proto3" json:"Y,omitempty"`
	Width         int64                  `protobuf:"varint,3,opt,name=Width,proto3" json:"Width,omitempty"`
	Height        int64                  `protobuf:"varint,4,opt,name=Height,proto3" json:"Height,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddAction) Reset() {
	*x = AddAction{}
	mi := &file_retro_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAction) ProtoMessage() {}

func (x *AddAction) ProtoReflect() protoreflect.Message {
	mi := &file_retro_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAction.ProtoReflect.Descriptor instead.
func (*AddAction) Descriptor() ([]byte, []int) {
	return file_retro_proto_rawDescGZIP(), []int{4}
}

func (x *AddAction) GetX() int64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *AddAction) GetY() int64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *AddAction) GetWidth() int64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *AddAction) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type MoveAction struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StickyID      string                 `protobuf:"bytes,1,opt,name=StickyID,proto3" json:"StickyID,omitempty"`
	X             int64                  `protobuf:"varint,2,opt,name=X,proto3" json:"X,omitempty"`
	Y             int64                  `protobuf:"varint,3,opt,name=Y,proto3" json:"Y,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MoveAction) Reset() {
	*x = MoveAction{}
	mi := &file_retro_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MoveAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveAction) ProtoMessage() {}

func (x *MoveAction) ProtoReflect() protoreflect.Message {
	mi := &file_retro_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveAction.ProtoReflect.Descriptor instead.
func (*MoveAction) Descriptor() ([]byte, []int) {
	return file_retro_proto_rawDescGZIP(), []int{5}
}

func (x *MoveAction) GetStickyID() string {
	if x != nil {
		return x.StickyID
	}
	return ""
}

func (x *MoveAction) GetX() int64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *MoveAction) GetY() int64 {
	if x != nil {
		return x.Y
	}
	return 0
}

type ResizeAction struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StickyID      string                 `protobuf:"bytes,1,opt,name=StickyID,proto3" json:"StickyID,omitempty"`
	X             int64                  `protobuf:"varint,2,opt,name=X,proto3" json:"X,omitempty"`
	Y             int64                  `protobuf:"varint,3,opt,name=Y,proto3" json:"Y,omitempty"`
	Height        int64                  `protobuf:"varint,4,opt,name=Height,proto3" json:"Height,omitempty"`
	Width         int64                  `protobuf:"varint,5,opt,name=Width,proto3" json:"Width,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResizeAction) Reset() {
	*x = ResizeAction{}
	mi := &file_retro_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResizeAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResizeAction) ProtoMessage() {}

func (x *ResizeAction) ProtoReflect() protoreflect.Message {
	mi := &file_retro_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResizeAction.ProtoReflect.Descriptor instead.
func (*ResizeAction) Descriptor() ([]byte, []int) {
	return file_retro_proto_rawDescGZIP(), []int{6}
}

func (x *ResizeAction) GetStickyID() string {
	if x != nil {
		return x.StickyID
	}
	return ""
}

func (x *ResizeAction) GetX() int64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *ResizeAction) GetY() int64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *ResizeAction) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *ResizeAction) GetWidth() int64 {
	if x != nil {
		return x.Width
	}
	return 0
}

type EditAction struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StickyID      string                 `protobuf:"bytes,1,opt,name=StickyID,proto3" json:"StickyID,omitempty"`
	Content       string                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EditAction) Reset() {
	*x = EditAction{}
	mi := &file_retro_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EditAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditAction) ProtoMessage() {}

func (x *EditAction) ProtoReflect() protoreflect.Message {
	mi := &file_retro_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditAction.ProtoReflect.Descriptor instead.
func (*EditAction) Descriptor() ([]byte, []int) {
	return file_retro_proto_rawDescGZIP(), []int{7}
}

func (x *EditAction) GetStickyID() string {
	if x != nil {
		return x.StickyID
	}
	return ""
}

func (x *EditAction) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type DeleteAction struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StickyID      string                 `protobuf:"bytes,1,opt,name=StickyID,proto3" json:"StickyID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAction) Reset() {
	*x = DeleteAction{}
	mi := &file_retro_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAction) ProtoMessage() {}

func (x *DeleteAction) ProtoReflect() protoreflect.Message {
	mi := &file_retro_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAction.ProtoReflect.Descriptor instead.
func (*DeleteAction) Descriptor() ([]byte, []int) {
	return file_retro_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteAction) GetStickyID() string {
	if x != nil {
		return x.StickyID
	}
	return ""
}

type Action struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Action:
	//
	//	*Action_Select
	//	*Action_Add
	//	*Action_Move
	//	*Action_Resize
	//	*Action_Edit
	//	*Action_Delete
	Action        isAction_Action `protobuf_oneof:"action"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Action) Reset() {
	*x = Action{}
	mi := &file_retro_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_retro_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_retro_proto_rawDescGZIP(), []int{9}
}

func (x *Action) GetAction() isAction_Action {
	if x != nil {
		return x.Action
	}
	return nil
}

func (x *Action) GetSelect() *SelectAction {
	if x != nil {
		if x, ok := x.Action.(*Action_Select); ok {
			return x.Select
		}
	}
	return nil
}

func (x *Action) GetAdd() *AddAction {
	if x != nil {
		if x, ok := x.Action.(*Action_Add); ok {
			return x.Add
		}
	}
	return nil
}

func (x *Action) GetMove() *MoveAction {
	if x != nil {
		if x, ok := x.Action.(*Action_Move); ok {
			return x.Move
		}
	}
	return nil
}

func (x *Action) GetResize() *ResizeAction {
	if x != nil {
		if x, ok := x.Action.(*Action_Resize); ok {
			return x.Resize
		}
	}
	return nil
}

func (x *Action) GetEdit() *EditAction {
	if x != nil {
		if x, ok := x.Action.(*Action_Edit); ok {
			return x.Edit
		}
	}
	return nil
}

func (x *Action) GetDelete() *DeleteAction {
	if x != nil {
		if x, ok := x.Action.(*Action_Delete); ok {
			return x.Delete
		}
	}
	return nil
}

type isAction_Action interface {
	isAction_Action()
}

type Action_Select struct {
	Select *SelectAction `protobuf:"bytes,1,opt,name=select,proto3,oneof"`
}

type Action_Add struct {
	Add *AddAction `protobuf:"bytes,2,opt,name=add,proto3,oneof"`
}

type Action_Move struct {
	Move *MoveAction `protobuf:"bytes,3,opt,name=move,proto3,oneof"`
}

type Action_Resize struct {
	Resize *ResizeAction `protobuf:"bytes,4,opt,name=resize,proto3,oneof"`
}

type Action_Edit struct {
	Edit *EditAction `protobuf:"bytes,5,opt,name=edit,proto3,oneof"`
}

type Action_Delete struct {
	Delete *DeleteAction `protobuf:"bytes,6,opt,name=delete,proto3,oneof"`
}

func (*Action_Select) isAction_Action() {}

func (*Action_Add) isAction_Action() {}

func (*Action_Move) isAction_Action() {}

func (*Action_Resize) isAction_Action() {}

func (*Action_Edit) isAction_Action() {}

func (*Action_Delete) isAction_Action() {}

var File_retro_proto protoreflect.FileDescriptor

var file_retro_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x72, 0x65, 0x74, 0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x72,
	0x65, 0x74, 0x72, 0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x69, 0x63, 0x6b, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0a, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x42, 0x79, 0x88, 0x01, 0x01, 0x12, 0x0c, 0x0a, 0x01, 0x58, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x58, 0x12, 0x0c, 0x0a, 0x01, 0x59, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x01, 0x59, 0x12, 0x14, 0x0a, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06,
	0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x48, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x42, 0x79, 0x22, 0x51, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x68, 0x61, 0x73,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0b, 0x68, 0x61, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x68, 0x61, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x22, 0xa8, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x73, 0x74,
	0x69, 0x63, 0x6b, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72,
	0x65, 0x74, 0x72, 0x6f, 0x2e, 0x53, 0x74, 0x69, 0x63, 0x6b, 0x79, 0x52, 0x08, 0x73, 0x74, 0x69,
	0x63, 0x6b, 0x69, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x6f, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x1a, 0x45, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2a, 0x0a, 0x0c, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x53,
	0x74, 0x69, 0x63, 0x6b, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53,
	0x74, 0x69, 0x63, 0x6b, 0x79, 0x49, 0x44, 0x22, 0x55, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x58, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x01, 0x58, 0x12, 0x0c, 0x0a, 0x01, 0x59, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x59,
	0x12, 0x14, 0x0a, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x44,
	0x0a, 0x0a, 0x4d, 0x6f, 0x76, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x53, 0x74, 0x69, 0x63, 0x6b, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x53, 0x74, 0x69, 0x63, 0x6b, 0x79, 0x49, 0x44, 0x12, 0x0c, 0x0a, 0x01, 0x58, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x01, 0x58, 0x12, 0x0c, 0x0a, 0x01, 0x59, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x01, 0x59, 0x22, 0x74, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x74, 0x69, 0x63, 0x6b, 0x79, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x74, 0x69, 0x63, 0x6b, 0x79, 0x49, 0x44,
	0x12, 0x0c, 0x0a, 0x01, 0x58, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x58, 0x12, 0x0c,
	0x0a, 0x01, 0x59, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x59, 0x12, 0x16, 0x0a, 0x06,
	0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x48, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x22, 0x42, 0x0a, 0x0a, 0x45, 0x64,
	0x69, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x74, 0x69, 0x63,
	0x6b, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x74, 0x69, 0x63,
	0x6b, 0x79, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x2a,
	0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x53, 0x74, 0x69, 0x63, 0x6b, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x53, 0x74, 0x69, 0x63, 0x6b, 0x79, 0x49, 0x44, 0x22, 0x97, 0x02, 0x0a, 0x06, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x6f, 0x2e, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x06, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x12, 0x24, 0x0a, 0x03, 0x61, 0x64, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x03, 0x61, 0x64, 0x64, 0x12, 0x27, 0x0a, 0x04, 0x6d, 0x6f,
	0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x6f,
	0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x04, 0x6d,
	0x6f, 0x76, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x69,
	0x7a, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x73, 0x69,
	0x7a, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x65, 0x64, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x6f, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x04, 0x65, 0x64, 0x69, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x65,
	0x74, 0x72, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x00, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x72, 0x65, 0x74, 0x72, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_retro_proto_rawDescOnce sync.Once
	file_retro_proto_rawDescData = file_retro_proto_rawDesc
)

func file_retro_proto_rawDescGZIP() []byte {
	file_retro_proto_rawDescOnce.Do(func() {
		file_retro_proto_rawDescData = protoimpl.X.CompressGZIP(file_retro_proto_rawDescData)
	})
	return file_retro_proto_rawDescData
}

var file_retro_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_retro_proto_goTypes = []any{
	(*Sticky)(nil),       // 0: retro.Sticky
	(*User)(nil),         // 1: retro.User
	(*State)(nil),        // 2: retro.State
	(*SelectAction)(nil), // 3: retro.SelectAction
	(*AddAction)(nil),    // 4: retro.AddAction
	(*MoveAction)(nil),   // 5: retro.MoveAction
	(*ResizeAction)(nil), // 6: retro.ResizeAction
	(*EditAction)(nil),   // 7: retro.EditAction
	(*DeleteAction)(nil), // 8: retro.DeleteAction
	(*Action)(nil),       // 9: retro.Action
	nil,                  // 10: retro.State.UsersEntry
}
var file_retro_proto_depIdxs = []int32{
	0,  // 0: retro.State.stickies:type_name -> retro.Sticky
	10, // 1: retro.State.users:type_name -> retro.State.UsersEntry
	3,  // 2: retro.Action.select:type_name -> retro.SelectAction
	4,  // 3: retro.Action.add:type_name -> retro.AddAction
	5,  // 4: retro.Action.move:type_name -> retro.MoveAction
	6,  // 5: retro.Action.resize:type_name -> retro.ResizeAction
	7,  // 6: retro.Action.edit:type_name -> retro.EditAction
	8,  // 7: retro.Action.delete:type_name -> retro.DeleteAction
	1,  // 8: retro.State.UsersEntry.value:type_name -> retro.User
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_retro_proto_init() }
func file_retro_proto_init() {
	if File_retro_proto != nil {
		return
	}
	file_retro_proto_msgTypes[0].OneofWrappers = []any{}
	file_retro_proto_msgTypes[1].OneofWrappers = []any{}
	file_retro_proto_msgTypes[9].OneofWrappers = []any{
		(*Action_Select)(nil),
		(*Action_Add)(nil),
		(*Action_Move)(nil),
		(*Action_Resize)(nil),
		(*Action_Edit)(nil),
		(*Action_Delete)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_retro_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_retro_proto_goTypes,
		DependencyIndexes: file_retro_proto_depIdxs,
		MessageInfos:      file_retro_proto_msgTypes,
	}.Build()
	File_retro_proto = out.File
	file_retro_proto_rawDesc = nil
	file_retro_proto_goTypes = nil
	file_retro_proto_depIdxs = nil
}
