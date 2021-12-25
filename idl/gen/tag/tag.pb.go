// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: idl/tag.proto

package tag

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

type BaseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrNo  int64  `protobuf:"varint,1,opt,name=err_no,json=errNo,proto3" json:"err_no,omitempty"`
	ErrMsg string `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
}

func (x *BaseResp) Reset() {
	*x = BaseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResp) ProtoMessage() {}

func (x *BaseResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResp.ProtoReflect.Descriptor instead.
func (*BaseResp) Descriptor() ([]byte, []int) {
	return file_idl_tag_proto_rawDescGZIP(), []int{0}
}

func (x *BaseResp) GetErrNo() int64 {
	if x != nil {
		return x.ErrNo
	}
	return 0
}

func (x *BaseResp) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Content   string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	UpdatedAt int64  `protobuf:"varint,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"` // 对于用户来说，这是用户最后一次使用这个tag的时间，对于电影来说，这是电影最后一次被打这个tag的时间
	NTimes    int64  `protobuf:"varint,4,opt,name=n_times,json=nTimes,proto3" json:"n_times,omitempty"`          // 对于用户来说，这是用户使用这个tag的次数，对于电影来说，这是电影被打上这个tag的次数
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_idl_tag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_idl_tag_proto_rawDescGZIP(), []int{1}
}

func (x *Tag) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Tag) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Tag) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Tag) GetNTimes() int64 {
	if x != nil {
		return x.NTimes
	}
	return 0
}

type CreateTagReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MovieId    string `protobuf:"bytes,2,opt,name=movie_id,json=movieId,proto3" json:"movie_id,omitempty"`
	TagContent string `protobuf:"bytes,3,opt,name=tag_content,json=tagContent,proto3" json:"tag_content,omitempty"`
}

func (x *CreateTagReq) Reset() {
	*x = CreateTagReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_tag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTagReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTagReq) ProtoMessage() {}

func (x *CreateTagReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_tag_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTagReq.ProtoReflect.Descriptor instead.
func (*CreateTagReq) Descriptor() ([]byte, []int) {
	return file_idl_tag_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTagReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateTagReq) GetMovieId() string {
	if x != nil {
		return x.MovieId
	}
	return ""
}

func (x *CreateTagReq) GetTagContent() string {
	if x != nil {
		return x.TagContent
	}
	return ""
}

type CreateTagResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResp `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
}

func (x *CreateTagResp) Reset() {
	*x = CreateTagResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_tag_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTagResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTagResp) ProtoMessage() {}

func (x *CreateTagResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_tag_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTagResp.ProtoReflect.Descriptor instead.
func (*CreateTagResp) Descriptor() ([]byte, []int) {
	return file_idl_tag_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTagResp) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

type QueryMovieTagReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MovieId string `protobuf:"bytes,2,opt,name=movie_id,json=movieId,proto3" json:"movie_id,omitempty"`
}

func (x *QueryMovieTagReq) Reset() {
	*x = QueryMovieTagReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_tag_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMovieTagReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMovieTagReq) ProtoMessage() {}

func (x *QueryMovieTagReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_tag_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMovieTagReq.ProtoReflect.Descriptor instead.
func (*QueryMovieTagReq) Descriptor() ([]byte, []int) {
	return file_idl_tag_proto_rawDescGZIP(), []int{4}
}

func (x *QueryMovieTagReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *QueryMovieTagReq) GetMovieId() string {
	if x != nil {
		return x.MovieId
	}
	return ""
}

type QueryMovieTagResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResp `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
	Tags     []*Tag    `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *QueryMovieTagResp) Reset() {
	*x = QueryMovieTagResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_tag_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMovieTagResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMovieTagResp) ProtoMessage() {}

func (x *QueryMovieTagResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_tag_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMovieTagResp.ProtoReflect.Descriptor instead.
func (*QueryMovieTagResp) Descriptor() ([]byte, []int) {
	return file_idl_tag_proto_rawDescGZIP(), []int{5}
}

func (x *QueryMovieTagResp) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *QueryMovieTagResp) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

type QueryUserTagCloudReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	NTags  int64  `protobuf:"varint,2,opt,name=n_tags,json=nTags,proto3" json:"n_tags,omitempty"` // 查询k大频率的tag
}

func (x *QueryUserTagCloudReq) Reset() {
	*x = QueryUserTagCloudReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_tag_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryUserTagCloudReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUserTagCloudReq) ProtoMessage() {}

func (x *QueryUserTagCloudReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_tag_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryUserTagCloudReq.ProtoReflect.Descriptor instead.
func (*QueryUserTagCloudReq) Descriptor() ([]byte, []int) {
	return file_idl_tag_proto_rawDescGZIP(), []int{6}
}

func (x *QueryUserTagCloudReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *QueryUserTagCloudReq) GetNTags() int64 {
	if x != nil {
		return x.NTags
	}
	return 0
}

type QueryUserTagCloudResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResp `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
	Tags     []*Tag    `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *QueryUserTagCloudResp) Reset() {
	*x = QueryUserTagCloudResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_tag_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryUserTagCloudResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUserTagCloudResp) ProtoMessage() {}

func (x *QueryUserTagCloudResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_tag_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryUserTagCloudResp.ProtoReflect.Descriptor instead.
func (*QueryUserTagCloudResp) Descriptor() ([]byte, []int) {
	return file_idl_tag_proto_rawDescGZIP(), []int{7}
}

func (x *QueryUserTagCloudResp) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *QueryUserTagCloudResp) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

type QueryTagRecordsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Page   int64  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Offset int64  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *QueryTagRecordsReq) Reset() {
	*x = QueryTagRecordsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_tag_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryTagRecordsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryTagRecordsReq) ProtoMessage() {}

func (x *QueryTagRecordsReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_tag_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryTagRecordsReq.ProtoReflect.Descriptor instead.
func (*QueryTagRecordsReq) Descriptor() ([]byte, []int) {
	return file_idl_tag_proto_rawDescGZIP(), []int{8}
}

func (x *QueryTagRecordsReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *QueryTagRecordsReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *QueryTagRecordsReq) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type QueryTagRecordsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResp `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
	Tags     []*Tag    `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	NRecords int64     `protobuf:"varint,3,opt,name=n_records,json=nRecords,proto3" json:"n_records,omitempty"`
}

func (x *QueryTagRecordsResp) Reset() {
	*x = QueryTagRecordsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_tag_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryTagRecordsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryTagRecordsResp) ProtoMessage() {}

func (x *QueryTagRecordsResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_tag_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryTagRecordsResp.ProtoReflect.Descriptor instead.
func (*QueryTagRecordsResp) Descriptor() ([]byte, []int) {
	return file_idl_tag_proto_rawDescGZIP(), []int{9}
}

func (x *QueryTagRecordsResp) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *QueryTagRecordsResp) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *QueryTagRecordsResp) GetNRecords() int64 {
	if x != nil {
		return x.NRecords
	}
	return 0
}

type QueryMovieTopNTagsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MovieId string `protobuf:"bytes,1,opt,name=movie_id,json=movieId,proto3" json:"movie_id,omitempty"`
	NTags   int64  `protobuf:"varint,2,opt,name=n_tags,json=nTags,proto3" json:"n_tags,omitempty"`
}

func (x *QueryMovieTopNTagsReq) Reset() {
	*x = QueryMovieTopNTagsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_tag_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMovieTopNTagsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMovieTopNTagsReq) ProtoMessage() {}

func (x *QueryMovieTopNTagsReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_tag_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMovieTopNTagsReq.ProtoReflect.Descriptor instead.
func (*QueryMovieTopNTagsReq) Descriptor() ([]byte, []int) {
	return file_idl_tag_proto_rawDescGZIP(), []int{10}
}

func (x *QueryMovieTopNTagsReq) GetMovieId() string {
	if x != nil {
		return x.MovieId
	}
	return ""
}

func (x *QueryMovieTopNTagsReq) GetNTags() int64 {
	if x != nil {
		return x.NTags
	}
	return 0
}

type QueryMovieTopNTagsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResp `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
	Tags     []*Tag    `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *QueryMovieTopNTagsResp) Reset() {
	*x = QueryMovieTopNTagsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_tag_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMovieTopNTagsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMovieTopNTagsResp) ProtoMessage() {}

func (x *QueryMovieTopNTagsResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_tag_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMovieTopNTagsResp.ProtoReflect.Descriptor instead.
func (*QueryMovieTopNTagsResp) Descriptor() ([]byte, []int) {
	return file_idl_tag_proto_rawDescGZIP(), []int{11}
}

func (x *QueryMovieTopNTagsResp) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *QueryMovieTopNTagsResp) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_idl_tag_proto protoreflect.FileDescriptor

var file_idl_tag_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x69, 0x64, 0x6c, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x74, 0x61, 0x67, 0x22, 0x3a, 0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x15, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x5f, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x4e, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x5f, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67,
	0x22, 0x67, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x22, 0x63, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x74, 0x61, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x74, 0x61, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x3b,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x2a, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x46, 0x0a, 0x10, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x61,
	0x67, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x22, 0x46, 0x0a, 0x14, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x61, 0x67, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6e, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x6e, 0x54, 0x61, 0x67, 0x73, 0x22, 0x61, 0x0a, 0x15, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x1c, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x74, 0x61, 0x67, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x59, 0x0a,
	0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x61, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x7c, 0x0a, 0x13, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x54, 0x61, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x2a, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x74, 0x61, 0x67, 0x2e,
	0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x5f, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6e, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x49, 0x0a, 0x15, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x54, 0x6f, 0x70, 0x4e, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x71, 0x12,
	0x19, 0x0a, 0x08, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6e, 0x5f,
	0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6e, 0x54, 0x61, 0x67,
	0x73, 0x22, 0x62, 0x0a, 0x16, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x54,
	0x6f, 0x70, 0x4e, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x09, 0x62,
	0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x74, 0x61, 0x67, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x62,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x54, 0x61, 0x67, 0x52,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x32, 0xeb, 0x02, 0x0a, 0x0a, 0x54, 0x61, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x67, 0x12, 0x11, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x67, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0d, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x54, 0x61, 0x67, 0x12, 0x15, 0x2e, 0x74, 0x61,
	0x67, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x54, 0x61, 0x67, 0x52,
	0x65, 0x71, 0x1a, 0x16, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x11,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x12, 0x19, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x61, 0x67, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x74,
	0x61, 0x67, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0f, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x54, 0x61, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x17, 0x2e,
	0x74, 0x61, 0x67, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x61, 0x67, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x54, 0x61, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x12, 0x4f, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x54, 0x6f, 0x70, 0x4e, 0x54, 0x61, 0x67, 0x73, 0x12, 0x1a, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x54, 0x6f, 0x70, 0x4e, 0x54, 0x61, 0x67,
	0x73, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x54, 0x6f, 0x70, 0x4e, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x74, 0x61, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_tag_proto_rawDescOnce sync.Once
	file_idl_tag_proto_rawDescData = file_idl_tag_proto_rawDesc
)

func file_idl_tag_proto_rawDescGZIP() []byte {
	file_idl_tag_proto_rawDescOnce.Do(func() {
		file_idl_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_tag_proto_rawDescData)
	})
	return file_idl_tag_proto_rawDescData
}

var file_idl_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_idl_tag_proto_goTypes = []interface{}{
	(*BaseResp)(nil),               // 0: tag.BaseResp
	(*Tag)(nil),                    // 1: tag.Tag
	(*CreateTagReq)(nil),           // 2: tag.CreateTagReq
	(*CreateTagResp)(nil),          // 3: tag.CreateTagResp
	(*QueryMovieTagReq)(nil),       // 4: tag.QueryMovieTagReq
	(*QueryMovieTagResp)(nil),      // 5: tag.QueryMovieTagResp
	(*QueryUserTagCloudReq)(nil),   // 6: tag.QueryUserTagCloudReq
	(*QueryUserTagCloudResp)(nil),  // 7: tag.QueryUserTagCloudResp
	(*QueryTagRecordsReq)(nil),     // 8: tag.QueryTagRecordsReq
	(*QueryTagRecordsResp)(nil),    // 9: tag.QueryTagRecordsResp
	(*QueryMovieTopNTagsReq)(nil),  // 10: tag.QueryMovieTopNTagsReq
	(*QueryMovieTopNTagsResp)(nil), // 11: tag.QueryMovieTopNTagsResp
}
var file_idl_tag_proto_depIdxs = []int32{
	0,  // 0: tag.CreateTagResp.base_resp:type_name -> tag.BaseResp
	0,  // 1: tag.QueryMovieTagResp.base_resp:type_name -> tag.BaseResp
	1,  // 2: tag.QueryMovieTagResp.tags:type_name -> tag.Tag
	0,  // 3: tag.QueryUserTagCloudResp.base_resp:type_name -> tag.BaseResp
	1,  // 4: tag.QueryUserTagCloudResp.tags:type_name -> tag.Tag
	0,  // 5: tag.QueryTagRecordsResp.base_resp:type_name -> tag.BaseResp
	1,  // 6: tag.QueryTagRecordsResp.tags:type_name -> tag.Tag
	0,  // 7: tag.QueryMovieTopNTagsResp.base_resp:type_name -> tag.BaseResp
	1,  // 8: tag.QueryMovieTopNTagsResp.tags:type_name -> tag.Tag
	2,  // 9: tag.TagService.CreateTag:input_type -> tag.CreateTagReq
	4,  // 10: tag.TagService.QueryMovieTag:input_type -> tag.QueryMovieTagReq
	6,  // 11: tag.TagService.QueryUserTagCloud:input_type -> tag.QueryUserTagCloudReq
	8,  // 12: tag.TagService.QueryTagRecords:input_type -> tag.QueryTagRecordsReq
	10, // 13: tag.TagService.QueryMovieTopNTags:input_type -> tag.QueryMovieTopNTagsReq
	3,  // 14: tag.TagService.CreateTag:output_type -> tag.CreateTagResp
	5,  // 15: tag.TagService.QueryMovieTag:output_type -> tag.QueryMovieTagResp
	7,  // 16: tag.TagService.QueryUserTagCloud:output_type -> tag.QueryUserTagCloudResp
	9,  // 17: tag.TagService.QueryTagRecords:output_type -> tag.QueryTagRecordsResp
	11, // 18: tag.TagService.QueryMovieTopNTags:output_type -> tag.QueryMovieTopNTagsResp
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_idl_tag_proto_init() }
func file_idl_tag_proto_init() {
	if File_idl_tag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResp); i {
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
		file_idl_tag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
		file_idl_tag_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTagReq); i {
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
		file_idl_tag_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTagResp); i {
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
		file_idl_tag_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMovieTagReq); i {
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
		file_idl_tag_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMovieTagResp); i {
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
		file_idl_tag_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryUserTagCloudReq); i {
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
		file_idl_tag_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryUserTagCloudResp); i {
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
		file_idl_tag_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryTagRecordsReq); i {
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
		file_idl_tag_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryTagRecordsResp); i {
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
		file_idl_tag_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMovieTopNTagsReq); i {
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
		file_idl_tag_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMovieTopNTagsResp); i {
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
			RawDescriptor: file_idl_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_tag_proto_goTypes,
		DependencyIndexes: file_idl_tag_proto_depIdxs,
		MessageInfos:      file_idl_tag_proto_msgTypes,
	}.Build()
	File_idl_tag_proto = out.File
	file_idl_tag_proto_rawDesc = nil
	file_idl_tag_proto_goTypes = nil
	file_idl_tag_proto_depIdxs = nil
}
