// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: protobuffs/recommender.proto

package protobuffs

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

type NewRatingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MovieId int32 `protobuf:"varint,2,opt,name=movie_id,json=movieId,proto3" json:"movie_id,omitempty"`
	Rating  int32 `protobuf:"varint,3,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *NewRatingRequest) Reset() {
	*x = NewRatingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuffs_recommender_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewRatingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewRatingRequest) ProtoMessage() {}

func (x *NewRatingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuffs_recommender_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewRatingRequest.ProtoReflect.Descriptor instead.
func (*NewRatingRequest) Descriptor() ([]byte, []int) {
	return file_protobuffs_recommender_proto_rawDescGZIP(), []int{0}
}

func (x *NewRatingRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *NewRatingRequest) GetMovieId() int32 {
	if x != nil {
		return x.MovieId
	}
	return 0
}

func (x *NewRatingRequest) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type RelevantMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *RelevantMovieRequest) Reset() {
	*x = RelevantMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuffs_recommender_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelevantMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelevantMovieRequest) ProtoMessage() {}

func (x *RelevantMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuffs_recommender_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelevantMovieRequest.ProtoReflect.Descriptor instead.
func (*RelevantMovieRequest) Descriptor() ([]byte, []int) {
	return file_protobuffs_recommender_proto_rawDescGZIP(), []int{1}
}

func (x *RelevantMovieRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RelevantMovieRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *RelevantMovieRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type SimilarMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MovieId int32 `protobuf:"varint,1,opt,name=movie_id,json=movieId,proto3" json:"movie_id,omitempty"`
	Offset  int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit   int32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *SimilarMovieRequest) Reset() {
	*x = SimilarMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuffs_recommender_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimilarMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimilarMovieRequest) ProtoMessage() {}

func (x *SimilarMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuffs_recommender_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimilarMovieRequest.ProtoReflect.Descriptor instead.
func (*SimilarMovieRequest) Descriptor() ([]byte, []int) {
	return file_protobuffs_recommender_proto_rawDescGZIP(), []int{2}
}

func (x *SimilarMovieRequest) GetMovieId() int32 {
	if x != nil {
		return x.MovieId
	}
	return 0
}

func (x *SimilarMovieRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *SimilarMovieRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type RelevantSimilarMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MovieId int32 `protobuf:"varint,2,opt,name=movie_id,json=movieId,proto3" json:"movie_id,omitempty"`
	Offset  int32 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit   int32 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *RelevantSimilarMovieRequest) Reset() {
	*x = RelevantSimilarMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuffs_recommender_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelevantSimilarMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelevantSimilarMovieRequest) ProtoMessage() {}

func (x *RelevantSimilarMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuffs_recommender_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelevantSimilarMovieRequest.ProtoReflect.Descriptor instead.
func (*RelevantSimilarMovieRequest) Descriptor() ([]byte, []int) {
	return file_protobuffs_recommender_proto_rawDescGZIP(), []int{3}
}

func (x *RelevantSimilarMovieRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RelevantSimilarMovieRequest) GetMovieId() int32 {
	if x != nil {
		return x.MovieId
	}
	return 0
}

func (x *RelevantSimilarMovieRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *RelevantSimilarMovieRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type MovieRecommendation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MovieRecommendation) Reset() {
	*x = MovieRecommendation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuffs_recommender_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieRecommendation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieRecommendation) ProtoMessage() {}

func (x *MovieRecommendation) ProtoReflect() protoreflect.Message {
	mi := &file_protobuffs_recommender_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieRecommendation.ProtoReflect.Descriptor instead.
func (*MovieRecommendation) Descriptor() ([]byte, []int) {
	return file_protobuffs_recommender_proto_rawDescGZIP(), []int{4}
}

func (x *MovieRecommendation) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuffs_recommender_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuffs_recommender_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_protobuffs_recommender_proto_rawDescGZIP(), []int{5}
}

func (x *EmptyResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type RecommendationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recommendations []*MovieRecommendation `protobuf:"bytes,1,rep,name=recommendations,proto3" json:"recommendations,omitempty"`
}

func (x *RecommendationResponse) Reset() {
	*x = RecommendationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuffs_recommender_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecommendationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendationResponse) ProtoMessage() {}

func (x *RecommendationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuffs_recommender_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendationResponse.ProtoReflect.Descriptor instead.
func (*RecommendationResponse) Descriptor() ([]byte, []int) {
	return file_protobuffs_recommender_proto_rawDescGZIP(), []int{6}
}

func (x *RecommendationResponse) GetRecommendations() []*MovieRecommendation {
	if x != nil {
		return x.Recommendations
	}
	return nil
}

var File_protobuffs_recommender_proto protoreflect.FileDescriptor

var file_protobuffs_recommender_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x66, 0x73, 0x2f, 0x72, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e,
	0x0a, 0x10, 0x4e, 0x65, 0x77, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x5d,
	0x0a, 0x14, 0x52, 0x65, 0x6c, 0x65, 0x76, 0x61, 0x6e, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x5e, 0x0a,
	0x13, 0x53, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x7f, 0x0a,
	0x1b, 0x52, 0x65, 0x6c, 0x65, 0x76, 0x61, 0x6e, 0x74, 0x53, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x25,
	0x0a, 0x13, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x58, 0x0a, 0x16, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3e, 0x0a, 0x0f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x32, 0xa3, 0x02, 0x0a, 0x0b, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x32, 0x0a, 0x0d, 0x53, 0x61, 0x76, 0x65, 0x4e, 0x65, 0x77, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x11, 0x2e, 0x4e, 0x65, 0x77, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x15, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x76, 0x61, 0x6e,
	0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x15, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x53, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12,
	0x14, 0x2e, 0x53, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56,
	0x0a, 0x1d, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x65, 0x76,
	0x61, 0x6e, 0x74, 0x53, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12,
	0x1c, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x76, 0x61, 0x6e, 0x74, 0x53, 0x69, 0x6d, 0x69, 0x6c, 0x61,
	0x72, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x75, 0x6d, 0x6f, 0x66, 0x66, 0x2f, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x2d, 0x74, 0x68, 0x65, 0x61, 0x74, 0x65, 0x72, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x66, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuffs_recommender_proto_rawDescOnce sync.Once
	file_protobuffs_recommender_proto_rawDescData = file_protobuffs_recommender_proto_rawDesc
)

func file_protobuffs_recommender_proto_rawDescGZIP() []byte {
	file_protobuffs_recommender_proto_rawDescOnce.Do(func() {
		file_protobuffs_recommender_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuffs_recommender_proto_rawDescData)
	})
	return file_protobuffs_recommender_proto_rawDescData
}

var file_protobuffs_recommender_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protobuffs_recommender_proto_goTypes = []interface{}{
	(*NewRatingRequest)(nil),            // 0: NewRatingRequest
	(*RelevantMovieRequest)(nil),        // 1: RelevantMovieRequest
	(*SimilarMovieRequest)(nil),         // 2: SimilarMovieRequest
	(*RelevantSimilarMovieRequest)(nil), // 3: RelevantSimilarMovieRequest
	(*MovieRecommendation)(nil),         // 4: MovieRecommendation
	(*EmptyResponse)(nil),               // 5: EmptyResponse
	(*RecommendationResponse)(nil),      // 6: RecommendationResponse
}
var file_protobuffs_recommender_proto_depIdxs = []int32{
	4, // 0: RecommendationResponse.recommendations:type_name -> MovieRecommendation
	0, // 1: Recommender.SaveNewRating:input_type -> NewRatingRequest
	1, // 2: Recommender.RecommendMovie:input_type -> RelevantMovieRequest
	2, // 3: Recommender.RecommendSimilarMovie:input_type -> SimilarMovieRequest
	3, // 4: Recommender.RecommendRelevantSimilarMovie:input_type -> RelevantSimilarMovieRequest
	5, // 5: Recommender.SaveNewRating:output_type -> EmptyResponse
	6, // 6: Recommender.RecommendMovie:output_type -> RecommendationResponse
	6, // 7: Recommender.RecommendSimilarMovie:output_type -> RecommendationResponse
	6, // 8: Recommender.RecommendRelevantSimilarMovie:output_type -> RecommendationResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protobuffs_recommender_proto_init() }
func file_protobuffs_recommender_proto_init() {
	if File_protobuffs_recommender_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuffs_recommender_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewRatingRequest); i {
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
		file_protobuffs_recommender_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelevantMovieRequest); i {
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
		file_protobuffs_recommender_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimilarMovieRequest); i {
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
		file_protobuffs_recommender_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelevantSimilarMovieRequest); i {
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
		file_protobuffs_recommender_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieRecommendation); i {
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
		file_protobuffs_recommender_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
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
		file_protobuffs_recommender_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecommendationResponse); i {
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
			RawDescriptor: file_protobuffs_recommender_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuffs_recommender_proto_goTypes,
		DependencyIndexes: file_protobuffs_recommender_proto_depIdxs,
		MessageInfos:      file_protobuffs_recommender_proto_msgTypes,
	}.Build()
	File_protobuffs_recommender_proto = out.File
	file_protobuffs_recommender_proto_rawDesc = nil
	file_protobuffs_recommender_proto_goTypes = nil
	file_protobuffs_recommender_proto_depIdxs = nil
}
