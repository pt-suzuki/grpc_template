// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: models/rock_paper_scissors/report.proto

package rock_paper_scissors

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

// 今までの試合数、勝敗と対戦結果の履歴を持つメッセージ型です。
type Report struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumberOfGames int32 `protobuf:"varint,1,opt,name=numberOfGames,proto3" json:"numberOfGames,omitempty"`
	NumberOfWins  int32 `protobuf:"varint,2,opt,name=numberOfWins,proto3" json:"numberOfWins,omitempty"`
	// `repeated`を付けることで配列を表現できます。
	MatchResults []*MatchResult `protobuf:"bytes,3,rep,name=matchResults,proto3" json:"matchResults,omitempty"`
	UserId       string         `protobuf:"bytes,4,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *Report) Reset() {
	*x = Report{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_rock_paper_scissors_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Report) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Report) ProtoMessage() {}

func (x *Report) ProtoReflect() protoreflect.Message {
	mi := &file_models_rock_paper_scissors_report_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Report.ProtoReflect.Descriptor instead.
func (*Report) Descriptor() ([]byte, []int) {
	return file_models_rock_paper_scissors_report_proto_rawDescGZIP(), []int{0}
}

func (x *Report) GetNumberOfGames() int32 {
	if x != nil {
		return x.NumberOfGames
	}
	return 0
}

func (x *Report) GetNumberOfWins() int32 {
	if x != nil {
		return x.NumberOfWins
	}
	return 0
}

func (x *Report) GetMatchResults() []*MatchResult {
	if x != nil {
		return x.MatchResults
	}
	return nil
}

func (x *Report) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_models_rock_paper_scissors_report_proto protoreflect.FileDescriptor

var file_models_rock_paper_scissors_report_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x5f, 0x70, 0x61,
	0x70, 0x65, 0x72, 0x5f, 0x73, 0x63, 0x69, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x2f, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x1a, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x5f, 0x70,
	0x61, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x63, 0x69, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x2f, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa3, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x47, 0x61, 0x6d, 0x65,
	0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x57, 0x69, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f,
	0x66, 0x57, 0x69, 0x6e, 0x73, 0x12, 0x37, 0x0a, 0x0c, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x0c, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x42, 0x4f, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x74, 0x2d, 0x73, 0x75, 0x7a, 0x75, 0x6b, 0x69, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x72, 0x63,
	0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x5f, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5f, 0x73,
	0x63, 0x69, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_rock_paper_scissors_report_proto_rawDescOnce sync.Once
	file_models_rock_paper_scissors_report_proto_rawDescData = file_models_rock_paper_scissors_report_proto_rawDesc
)

func file_models_rock_paper_scissors_report_proto_rawDescGZIP() []byte {
	file_models_rock_paper_scissors_report_proto_rawDescOnce.Do(func() {
		file_models_rock_paper_scissors_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_rock_paper_scissors_report_proto_rawDescData)
	})
	return file_models_rock_paper_scissors_report_proto_rawDescData
}

var file_models_rock_paper_scissors_report_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_rock_paper_scissors_report_proto_goTypes = []interface{}{
	(*Report)(nil),      // 0: models.Report
	(*MatchResult)(nil), // 1: models.MatchResult
}
var file_models_rock_paper_scissors_report_proto_depIdxs = []int32{
	1, // 0: models.Report.matchResults:type_name -> models.MatchResult
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_models_rock_paper_scissors_report_proto_init() }
func file_models_rock_paper_scissors_report_proto_init() {
	if File_models_rock_paper_scissors_report_proto != nil {
		return
	}
	file_models_rock_paper_scissors_match_result_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_rock_paper_scissors_report_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Report); i {
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
			RawDescriptor: file_models_rock_paper_scissors_report_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_rock_paper_scissors_report_proto_goTypes,
		DependencyIndexes: file_models_rock_paper_scissors_report_proto_depIdxs,
		MessageInfos:      file_models_rock_paper_scissors_report_proto_msgTypes,
	}.Build()
	File_models_rock_paper_scissors_report_proto = out.File
	file_models_rock_paper_scissors_report_proto_rawDesc = nil
	file_models_rock_paper_scissors_report_proto_goTypes = nil
	file_models_rock_paper_scissors_report_proto_depIdxs = nil
}
