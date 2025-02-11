// Code generated by protoc-gen-go. DO NOT EDIT.
// source: IM.File.proto

package IM_File

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import IM_BaseDefine "github.com/shiywang/GoTalk/proto/IM_BaseDefine"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type IMFileLoginReq struct {
	// cmd id:	0x0501
	UserId               *uint32                       `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,omitempty"`
	TaskId               *string                       `protobuf:"bytes,2,req,name=task_id,json=taskId" json:"task_id,omitempty"`
	FileRole             *IM_BaseDefine.ClientFileRole `protobuf:"varint,3,req,name=file_role,json=fileRole,enum=IM.BaseDefine.ClientFileRole" json:"file_role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *IMFileLoginReq) Reset()         { *m = IMFileLoginReq{} }
func (m *IMFileLoginReq) String() string { return proto.CompactTextString(m) }
func (*IMFileLoginReq) ProtoMessage()    {}
func (*IMFileLoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_IM_File_881fa908902f7caa, []int{0}
}
func (m *IMFileLoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMFileLoginReq.Unmarshal(m, b)
}
func (m *IMFileLoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMFileLoginReq.Marshal(b, m, deterministic)
}
func (dst *IMFileLoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMFileLoginReq.Merge(dst, src)
}
func (m *IMFileLoginReq) XXX_Size() int {
	return xxx_messageInfo_IMFileLoginReq.Size(m)
}
func (m *IMFileLoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IMFileLoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_IMFileLoginReq proto.InternalMessageInfo

func (m *IMFileLoginReq) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMFileLoginReq) GetTaskId() string {
	if m != nil && m.TaskId != nil {
		return *m.TaskId
	}
	return ""
}

func (m *IMFileLoginReq) GetFileRole() IM_BaseDefine.ClientFileRole {
	if m != nil && m.FileRole != nil {
		return *m.FileRole
	}
	return IM_BaseDefine.ClientFileRole_CLIENT_REALTIME_SENDER
}

type IMFileLoginRsp struct {
	// cmd id:	0x0502
	ResultCode           *uint32  `protobuf:"varint,1,req,name=result_code,json=resultCode" json:"result_code,omitempty"`
	TaskId               *string  `protobuf:"bytes,2,req,name=task_id,json=taskId" json:"task_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IMFileLoginRsp) Reset()         { *m = IMFileLoginRsp{} }
func (m *IMFileLoginRsp) String() string { return proto.CompactTextString(m) }
func (*IMFileLoginRsp) ProtoMessage()    {}
func (*IMFileLoginRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_IM_File_881fa908902f7caa, []int{1}
}
func (m *IMFileLoginRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMFileLoginRsp.Unmarshal(m, b)
}
func (m *IMFileLoginRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMFileLoginRsp.Marshal(b, m, deterministic)
}
func (dst *IMFileLoginRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMFileLoginRsp.Merge(dst, src)
}
func (m *IMFileLoginRsp) XXX_Size() int {
	return xxx_messageInfo_IMFileLoginRsp.Size(m)
}
func (m *IMFileLoginRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_IMFileLoginRsp.DiscardUnknown(m)
}

var xxx_messageInfo_IMFileLoginRsp proto.InternalMessageInfo

func (m *IMFileLoginRsp) GetResultCode() uint32 {
	if m != nil && m.ResultCode != nil {
		return *m.ResultCode
	}
	return 0
}

func (m *IMFileLoginRsp) GetTaskId() string {
	if m != nil && m.TaskId != nil {
		return *m.TaskId
	}
	return ""
}

type IMFileState struct {
	// cmd id: 	0x0503
	State                *IM_BaseDefine.ClientFileState `protobuf:"varint,1,req,name=state,enum=IM.BaseDefine.ClientFileState" json:"state,omitempty"`
	TaskId               *string                        `protobuf:"bytes,2,req,name=task_id,json=taskId" json:"task_id,omitempty"`
	UserId               *uint32                        `protobuf:"varint,3,req,name=user_id,json=userId" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *IMFileState) Reset()         { *m = IMFileState{} }
func (m *IMFileState) String() string { return proto.CompactTextString(m) }
func (*IMFileState) ProtoMessage()    {}
func (*IMFileState) Descriptor() ([]byte, []int) {
	return fileDescriptor_IM_File_881fa908902f7caa, []int{2}
}
func (m *IMFileState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMFileState.Unmarshal(m, b)
}
func (m *IMFileState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMFileState.Marshal(b, m, deterministic)
}
func (dst *IMFileState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMFileState.Merge(dst, src)
}
func (m *IMFileState) XXX_Size() int {
	return xxx_messageInfo_IMFileState.Size(m)
}
func (m *IMFileState) XXX_DiscardUnknown() {
	xxx_messageInfo_IMFileState.DiscardUnknown(m)
}

var xxx_messageInfo_IMFileState proto.InternalMessageInfo

func (m *IMFileState) GetState() IM_BaseDefine.ClientFileState {
	if m != nil && m.State != nil {
		return *m.State
	}
	return IM_BaseDefine.ClientFileState_CLIENT_FILE_PEER_READY
}

func (m *IMFileState) GetTaskId() string {
	if m != nil && m.TaskId != nil {
		return *m.TaskId
	}
	return ""
}

func (m *IMFileState) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

type IMFilePullDataReq struct {
	// cmd id:	0x0504
	TaskId               *string                         `protobuf:"bytes,1,req,name=task_id,json=taskId" json:"task_id,omitempty"`
	UserId               *uint32                         `protobuf:"varint,2,req,name=user_id,json=userId" json:"user_id,omitempty"`
	TransMode            *IM_BaseDefine.TransferFileType `protobuf:"varint,3,req,name=trans_mode,json=transMode,enum=IM.BaseDefine.TransferFileType" json:"trans_mode,omitempty"`
	Offset               *uint32                         `protobuf:"varint,4,req,name=offset" json:"offset,omitempty"`
	DataSize             *uint32                         `protobuf:"varint,5,req,name=data_size,json=dataSize" json:"data_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *IMFilePullDataReq) Reset()         { *m = IMFilePullDataReq{} }
func (m *IMFilePullDataReq) String() string { return proto.CompactTextString(m) }
func (*IMFilePullDataReq) ProtoMessage()    {}
func (*IMFilePullDataReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_IM_File_881fa908902f7caa, []int{3}
}
func (m *IMFilePullDataReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMFilePullDataReq.Unmarshal(m, b)
}
func (m *IMFilePullDataReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMFilePullDataReq.Marshal(b, m, deterministic)
}
func (dst *IMFilePullDataReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMFilePullDataReq.Merge(dst, src)
}
func (m *IMFilePullDataReq) XXX_Size() int {
	return xxx_messageInfo_IMFilePullDataReq.Size(m)
}
func (m *IMFilePullDataReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IMFilePullDataReq.DiscardUnknown(m)
}

var xxx_messageInfo_IMFilePullDataReq proto.InternalMessageInfo

func (m *IMFilePullDataReq) GetTaskId() string {
	if m != nil && m.TaskId != nil {
		return *m.TaskId
	}
	return ""
}

func (m *IMFilePullDataReq) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMFilePullDataReq) GetTransMode() IM_BaseDefine.TransferFileType {
	if m != nil && m.TransMode != nil {
		return *m.TransMode
	}
	return IM_BaseDefine.TransferFileType_FILE_TYPE_ONLINE
}

func (m *IMFilePullDataReq) GetOffset() uint32 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *IMFilePullDataReq) GetDataSize() uint32 {
	if m != nil && m.DataSize != nil {
		return *m.DataSize
	}
	return 0
}

type IMFilePullDataRsp struct {
	// cmd id: 	0x0505
	ResultCode           *uint32  `protobuf:"varint,1,req,name=result_code,json=resultCode" json:"result_code,omitempty"`
	TaskId               *string  `protobuf:"bytes,2,req,name=task_id,json=taskId" json:"task_id,omitempty"`
	UserId               *uint32  `protobuf:"varint,3,req,name=user_id,json=userId" json:"user_id,omitempty"`
	Offset               *uint32  `protobuf:"varint,4,req,name=offset" json:"offset,omitempty"`
	FileData             []byte   `protobuf:"bytes,5,req,name=file_data,json=fileData" json:"file_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IMFilePullDataRsp) Reset()         { *m = IMFilePullDataRsp{} }
func (m *IMFilePullDataRsp) String() string { return proto.CompactTextString(m) }
func (*IMFilePullDataRsp) ProtoMessage()    {}
func (*IMFilePullDataRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_IM_File_881fa908902f7caa, []int{4}
}
func (m *IMFilePullDataRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMFilePullDataRsp.Unmarshal(m, b)
}
func (m *IMFilePullDataRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMFilePullDataRsp.Marshal(b, m, deterministic)
}
func (dst *IMFilePullDataRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMFilePullDataRsp.Merge(dst, src)
}
func (m *IMFilePullDataRsp) XXX_Size() int {
	return xxx_messageInfo_IMFilePullDataRsp.Size(m)
}
func (m *IMFilePullDataRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_IMFilePullDataRsp.DiscardUnknown(m)
}

var xxx_messageInfo_IMFilePullDataRsp proto.InternalMessageInfo

func (m *IMFilePullDataRsp) GetResultCode() uint32 {
	if m != nil && m.ResultCode != nil {
		return *m.ResultCode
	}
	return 0
}

func (m *IMFilePullDataRsp) GetTaskId() string {
	if m != nil && m.TaskId != nil {
		return *m.TaskId
	}
	return ""
}

func (m *IMFilePullDataRsp) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMFilePullDataRsp) GetOffset() uint32 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *IMFilePullDataRsp) GetFileData() []byte {
	if m != nil {
		return m.FileData
	}
	return nil
}

type IMFileReq struct {
	// cmd id: 	0x0506
	FromUserId           *uint32                         `protobuf:"varint,1,req,name=from_user_id,json=fromUserId" json:"from_user_id,omitempty"`
	ToUserId             *uint32                         `protobuf:"varint,2,req,name=to_user_id,json=toUserId" json:"to_user_id,omitempty"`
	FileName             *string                         `protobuf:"bytes,3,req,name=file_name,json=fileName" json:"file_name,omitempty"`
	FileSize             *uint32                         `protobuf:"varint,4,req,name=file_size,json=fileSize" json:"file_size,omitempty"`
	TransMode            *IM_BaseDefine.TransferFileType `protobuf:"varint,5,req,name=trans_mode,json=transMode,enum=IM.BaseDefine.TransferFileType" json:"trans_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *IMFileReq) Reset()         { *m = IMFileReq{} }
func (m *IMFileReq) String() string { return proto.CompactTextString(m) }
func (*IMFileReq) ProtoMessage()    {}
func (*IMFileReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_IM_File_881fa908902f7caa, []int{5}
}
func (m *IMFileReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMFileReq.Unmarshal(m, b)
}
func (m *IMFileReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMFileReq.Marshal(b, m, deterministic)
}
func (dst *IMFileReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMFileReq.Merge(dst, src)
}
func (m *IMFileReq) XXX_Size() int {
	return xxx_messageInfo_IMFileReq.Size(m)
}
func (m *IMFileReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IMFileReq.DiscardUnknown(m)
}

var xxx_messageInfo_IMFileReq proto.InternalMessageInfo

func (m *IMFileReq) GetFromUserId() uint32 {
	if m != nil && m.FromUserId != nil {
		return *m.FromUserId
	}
	return 0
}

func (m *IMFileReq) GetToUserId() uint32 {
	if m != nil && m.ToUserId != nil {
		return *m.ToUserId
	}
	return 0
}

func (m *IMFileReq) GetFileName() string {
	if m != nil && m.FileName != nil {
		return *m.FileName
	}
	return ""
}

func (m *IMFileReq) GetFileSize() uint32 {
	if m != nil && m.FileSize != nil {
		return *m.FileSize
	}
	return 0
}

func (m *IMFileReq) GetTransMode() IM_BaseDefine.TransferFileType {
	if m != nil && m.TransMode != nil {
		return *m.TransMode
	}
	return IM_BaseDefine.TransferFileType_FILE_TYPE_ONLINE
}

type IMFileRsp struct {
	// cmd id: 	0x0507
	ResultCode           *uint32                         `protobuf:"varint,1,req,name=result_code,json=resultCode" json:"result_code,omitempty"`
	FromUserId           *uint32                         `protobuf:"varint,2,req,name=from_user_id,json=fromUserId" json:"from_user_id,omitempty"`
	ToUserId             *uint32                         `protobuf:"varint,3,req,name=to_user_id,json=toUserId" json:"to_user_id,omitempty"`
	FileName             *string                         `protobuf:"bytes,4,req,name=file_name,json=fileName" json:"file_name,omitempty"`
	TaskId               *string                         `protobuf:"bytes,5,req,name=task_id,json=taskId" json:"task_id,omitempty"`
	IpAddrList           []*IM_BaseDefine.IpAddr         `protobuf:"bytes,6,rep,name=ip_addr_list,json=ipAddrList" json:"ip_addr_list,omitempty"`
	TransMode            *IM_BaseDefine.TransferFileType `protobuf:"varint,7,req,name=trans_mode,json=transMode,enum=IM.BaseDefine.TransferFileType" json:"trans_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *IMFileRsp) Reset()         { *m = IMFileRsp{} }
func (m *IMFileRsp) String() string { return proto.CompactTextString(m) }
func (*IMFileRsp) ProtoMessage()    {}
func (*IMFileRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_IM_File_881fa908902f7caa, []int{6}
}
func (m *IMFileRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMFileRsp.Unmarshal(m, b)
}
func (m *IMFileRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMFileRsp.Marshal(b, m, deterministic)
}
func (dst *IMFileRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMFileRsp.Merge(dst, src)
}
func (m *IMFileRsp) XXX_Size() int {
	return xxx_messageInfo_IMFileRsp.Size(m)
}
func (m *IMFileRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_IMFileRsp.DiscardUnknown(m)
}

var xxx_messageInfo_IMFileRsp proto.InternalMessageInfo

func (m *IMFileRsp) GetResultCode() uint32 {
	if m != nil && m.ResultCode != nil {
		return *m.ResultCode
	}
	return 0
}

func (m *IMFileRsp) GetFromUserId() uint32 {
	if m != nil && m.FromUserId != nil {
		return *m.FromUserId
	}
	return 0
}

func (m *IMFileRsp) GetToUserId() uint32 {
	if m != nil && m.ToUserId != nil {
		return *m.ToUserId
	}
	return 0
}

func (m *IMFileRsp) GetFileName() string {
	if m != nil && m.FileName != nil {
		return *m.FileName
	}
	return ""
}

func (m *IMFileRsp) GetTaskId() string {
	if m != nil && m.TaskId != nil {
		return *m.TaskId
	}
	return ""
}

func (m *IMFileRsp) GetIpAddrList() []*IM_BaseDefine.IpAddr {
	if m != nil {
		return m.IpAddrList
	}
	return nil
}

func (m *IMFileRsp) GetTransMode() IM_BaseDefine.TransferFileType {
	if m != nil && m.TransMode != nil {
		return *m.TransMode
	}
	return IM_BaseDefine.TransferFileType_FILE_TYPE_ONLINE
}

type IMFileNotify struct {
	// cmd id: 	0x0508
	FromUserId           *uint32                         `protobuf:"varint,1,req,name=from_user_id,json=fromUserId" json:"from_user_id,omitempty"`
	ToUserId             *uint32                         `protobuf:"varint,2,req,name=to_user_id,json=toUserId" json:"to_user_id,omitempty"`
	FileName             *string                         `protobuf:"bytes,3,req,name=file_name,json=fileName" json:"file_name,omitempty"`
	FileSize             *uint32                         `protobuf:"varint,4,req,name=file_size,json=fileSize" json:"file_size,omitempty"`
	TaskId               *string                         `protobuf:"bytes,5,req,name=task_id,json=taskId" json:"task_id,omitempty"`
	IpAddrList           []*IM_BaseDefine.IpAddr         `protobuf:"bytes,6,rep,name=ip_addr_list,json=ipAddrList" json:"ip_addr_list,omitempty"`
	TransMode            *IM_BaseDefine.TransferFileType `protobuf:"varint,7,req,name=trans_mode,json=transMode,enum=IM.BaseDefine.TransferFileType" json:"trans_mode,omitempty"`
	OfflineReady         *uint32                         `protobuf:"varint,8,req,name=offline_ready,json=offlineReady" json:"offline_ready,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *IMFileNotify) Reset()         { *m = IMFileNotify{} }
func (m *IMFileNotify) String() string { return proto.CompactTextString(m) }
func (*IMFileNotify) ProtoMessage()    {}
func (*IMFileNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_IM_File_881fa908902f7caa, []int{7}
}
func (m *IMFileNotify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMFileNotify.Unmarshal(m, b)
}
func (m *IMFileNotify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMFileNotify.Marshal(b, m, deterministic)
}
func (dst *IMFileNotify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMFileNotify.Merge(dst, src)
}
func (m *IMFileNotify) XXX_Size() int {
	return xxx_messageInfo_IMFileNotify.Size(m)
}
func (m *IMFileNotify) XXX_DiscardUnknown() {
	xxx_messageInfo_IMFileNotify.DiscardUnknown(m)
}

var xxx_messageInfo_IMFileNotify proto.InternalMessageInfo

func (m *IMFileNotify) GetFromUserId() uint32 {
	if m != nil && m.FromUserId != nil {
		return *m.FromUserId
	}
	return 0
}

func (m *IMFileNotify) GetToUserId() uint32 {
	if m != nil && m.ToUserId != nil {
		return *m.ToUserId
	}
	return 0
}

func (m *IMFileNotify) GetFileName() string {
	if m != nil && m.FileName != nil {
		return *m.FileName
	}
	return ""
}

func (m *IMFileNotify) GetFileSize() uint32 {
	if m != nil && m.FileSize != nil {
		return *m.FileSize
	}
	return 0
}

func (m *IMFileNotify) GetTaskId() string {
	if m != nil && m.TaskId != nil {
		return *m.TaskId
	}
	return ""
}

func (m *IMFileNotify) GetIpAddrList() []*IM_BaseDefine.IpAddr {
	if m != nil {
		return m.IpAddrList
	}
	return nil
}

func (m *IMFileNotify) GetTransMode() IM_BaseDefine.TransferFileType {
	if m != nil && m.TransMode != nil {
		return *m.TransMode
	}
	return IM_BaseDefine.TransferFileType_FILE_TYPE_ONLINE
}

func (m *IMFileNotify) GetOfflineReady() uint32 {
	if m != nil && m.OfflineReady != nil {
		return *m.OfflineReady
	}
	return 0
}

type IMFileHasOfflineReq struct {
	// cmd id: 	0x0509
	UserId               *uint32  `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,omitempty"`
	AttachData           []byte   `protobuf:"bytes,20,opt,name=attach_data,json=attachData" json:"attach_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IMFileHasOfflineReq) Reset()         { *m = IMFileHasOfflineReq{} }
func (m *IMFileHasOfflineReq) String() string { return proto.CompactTextString(m) }
func (*IMFileHasOfflineReq) ProtoMessage()    {}
func (*IMFileHasOfflineReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_IM_File_881fa908902f7caa, []int{8}
}
func (m *IMFileHasOfflineReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMFileHasOfflineReq.Unmarshal(m, b)
}
func (m *IMFileHasOfflineReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMFileHasOfflineReq.Marshal(b, m, deterministic)
}
func (dst *IMFileHasOfflineReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMFileHasOfflineReq.Merge(dst, src)
}
func (m *IMFileHasOfflineReq) XXX_Size() int {
	return xxx_messageInfo_IMFileHasOfflineReq.Size(m)
}
func (m *IMFileHasOfflineReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IMFileHasOfflineReq.DiscardUnknown(m)
}

var xxx_messageInfo_IMFileHasOfflineReq proto.InternalMessageInfo

func (m *IMFileHasOfflineReq) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMFileHasOfflineReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMFileHasOfflineRsp struct {
	// cmd id:	0x050a
	UserId               *uint32                          `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,omitempty"`
	OfflineFileList      []*IM_BaseDefine.OfflineFileInfo `protobuf:"bytes,2,rep,name=offline_file_list,json=offlineFileList" json:"offline_file_list,omitempty"`
	IpAddrList           []*IM_BaseDefine.IpAddr          `protobuf:"bytes,3,rep,name=ip_addr_list,json=ipAddrList" json:"ip_addr_list,omitempty"`
	AttachData           []byte                           `protobuf:"bytes,20,opt,name=attach_data,json=attachData" json:"attach_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *IMFileHasOfflineRsp) Reset()         { *m = IMFileHasOfflineRsp{} }
func (m *IMFileHasOfflineRsp) String() string { return proto.CompactTextString(m) }
func (*IMFileHasOfflineRsp) ProtoMessage()    {}
func (*IMFileHasOfflineRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_IM_File_881fa908902f7caa, []int{9}
}
func (m *IMFileHasOfflineRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMFileHasOfflineRsp.Unmarshal(m, b)
}
func (m *IMFileHasOfflineRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMFileHasOfflineRsp.Marshal(b, m, deterministic)
}
func (dst *IMFileHasOfflineRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMFileHasOfflineRsp.Merge(dst, src)
}
func (m *IMFileHasOfflineRsp) XXX_Size() int {
	return xxx_messageInfo_IMFileHasOfflineRsp.Size(m)
}
func (m *IMFileHasOfflineRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_IMFileHasOfflineRsp.DiscardUnknown(m)
}

var xxx_messageInfo_IMFileHasOfflineRsp proto.InternalMessageInfo

func (m *IMFileHasOfflineRsp) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMFileHasOfflineRsp) GetOfflineFileList() []*IM_BaseDefine.OfflineFileInfo {
	if m != nil {
		return m.OfflineFileList
	}
	return nil
}

func (m *IMFileHasOfflineRsp) GetIpAddrList() []*IM_BaseDefine.IpAddr {
	if m != nil {
		return m.IpAddrList
	}
	return nil
}

func (m *IMFileHasOfflineRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMFileAddOfflineReq struct {
	// cmd id:	0x050b
	FromUserId           *uint32  `protobuf:"varint,1,req,name=from_user_id,json=fromUserId" json:"from_user_id,omitempty"`
	ToUserId             *uint32  `protobuf:"varint,2,req,name=to_user_id,json=toUserId" json:"to_user_id,omitempty"`
	TaskId               *string  `protobuf:"bytes,3,req,name=task_id,json=taskId" json:"task_id,omitempty"`
	FileName             *string  `protobuf:"bytes,4,req,name=file_name,json=fileName" json:"file_name,omitempty"`
	FileSize             *uint32  `protobuf:"varint,5,req,name=file_size,json=fileSize" json:"file_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IMFileAddOfflineReq) Reset()         { *m = IMFileAddOfflineReq{} }
func (m *IMFileAddOfflineReq) String() string { return proto.CompactTextString(m) }
func (*IMFileAddOfflineReq) ProtoMessage()    {}
func (*IMFileAddOfflineReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_IM_File_881fa908902f7caa, []int{10}
}
func (m *IMFileAddOfflineReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMFileAddOfflineReq.Unmarshal(m, b)
}
func (m *IMFileAddOfflineReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMFileAddOfflineReq.Marshal(b, m, deterministic)
}
func (dst *IMFileAddOfflineReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMFileAddOfflineReq.Merge(dst, src)
}
func (m *IMFileAddOfflineReq) XXX_Size() int {
	return xxx_messageInfo_IMFileAddOfflineReq.Size(m)
}
func (m *IMFileAddOfflineReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IMFileAddOfflineReq.DiscardUnknown(m)
}

var xxx_messageInfo_IMFileAddOfflineReq proto.InternalMessageInfo

func (m *IMFileAddOfflineReq) GetFromUserId() uint32 {
	if m != nil && m.FromUserId != nil {
		return *m.FromUserId
	}
	return 0
}

func (m *IMFileAddOfflineReq) GetToUserId() uint32 {
	if m != nil && m.ToUserId != nil {
		return *m.ToUserId
	}
	return 0
}

func (m *IMFileAddOfflineReq) GetTaskId() string {
	if m != nil && m.TaskId != nil {
		return *m.TaskId
	}
	return ""
}

func (m *IMFileAddOfflineReq) GetFileName() string {
	if m != nil && m.FileName != nil {
		return *m.FileName
	}
	return ""
}

func (m *IMFileAddOfflineReq) GetFileSize() uint32 {
	if m != nil && m.FileSize != nil {
		return *m.FileSize
	}
	return 0
}

type IMFileDelOfflineReq struct {
	// cmd id:	0x050c
	FromUserId           *uint32  `protobuf:"varint,1,req,name=from_user_id,json=fromUserId" json:"from_user_id,omitempty"`
	ToUserId             *uint32  `protobuf:"varint,2,req,name=to_user_id,json=toUserId" json:"to_user_id,omitempty"`
	TaskId               *string  `protobuf:"bytes,3,req,name=task_id,json=taskId" json:"task_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IMFileDelOfflineReq) Reset()         { *m = IMFileDelOfflineReq{} }
func (m *IMFileDelOfflineReq) String() string { return proto.CompactTextString(m) }
func (*IMFileDelOfflineReq) ProtoMessage()    {}
func (*IMFileDelOfflineReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_IM_File_881fa908902f7caa, []int{11}
}
func (m *IMFileDelOfflineReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMFileDelOfflineReq.Unmarshal(m, b)
}
func (m *IMFileDelOfflineReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMFileDelOfflineReq.Marshal(b, m, deterministic)
}
func (dst *IMFileDelOfflineReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMFileDelOfflineReq.Merge(dst, src)
}
func (m *IMFileDelOfflineReq) XXX_Size() int {
	return xxx_messageInfo_IMFileDelOfflineReq.Size(m)
}
func (m *IMFileDelOfflineReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IMFileDelOfflineReq.DiscardUnknown(m)
}

var xxx_messageInfo_IMFileDelOfflineReq proto.InternalMessageInfo

func (m *IMFileDelOfflineReq) GetFromUserId() uint32 {
	if m != nil && m.FromUserId != nil {
		return *m.FromUserId
	}
	return 0
}

func (m *IMFileDelOfflineReq) GetToUserId() uint32 {
	if m != nil && m.ToUserId != nil {
		return *m.ToUserId
	}
	return 0
}

func (m *IMFileDelOfflineReq) GetTaskId() string {
	if m != nil && m.TaskId != nil {
		return *m.TaskId
	}
	return ""
}

func init() {
	proto.RegisterType((*IMFileLoginReq)(nil), "IM.File.IMFileLoginReq")
	proto.RegisterType((*IMFileLoginRsp)(nil), "IM.File.IMFileLoginRsp")
	proto.RegisterType((*IMFileState)(nil), "IM.File.IMFileState")
	proto.RegisterType((*IMFilePullDataReq)(nil), "IM.File.IMFilePullDataReq")
	proto.RegisterType((*IMFilePullDataRsp)(nil), "IM.File.IMFilePullDataRsp")
	proto.RegisterType((*IMFileReq)(nil), "IM.File.IMFileReq")
	proto.RegisterType((*IMFileRsp)(nil), "IM.File.IMFileRsp")
	proto.RegisterType((*IMFileNotify)(nil), "IM.File.IMFileNotify")
	proto.RegisterType((*IMFileHasOfflineReq)(nil), "IM.File.IMFileHasOfflineReq")
	proto.RegisterType((*IMFileHasOfflineRsp)(nil), "IM.File.IMFileHasOfflineRsp")
	proto.RegisterType((*IMFileAddOfflineReq)(nil), "IM.File.IMFileAddOfflineReq")
	proto.RegisterType((*IMFileDelOfflineReq)(nil), "IM.File.IMFileDelOfflineReq")
}

func init() { proto.RegisterFile("IM.File.proto", fileDescriptor_IM_File_881fa908902f7caa) }

var fileDescriptor_IM_File_881fa908902f7caa = []byte{
	// 628 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x94, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xc7, 0x65, 0xbb, 0x49, 0xe3, 0x89, 0x93, 0x9f, 0xea, 0xfe, 0x00, 0x8b, 0x7f, 0xb1, 0xcc,
	0x25, 0xa7, 0x1c, 0x22, 0x24, 0x24, 0x0e, 0x48, 0x6d, 0x23, 0x54, 0x57, 0x4d, 0x8b, 0xdc, 0x72,
	0xb6, 0x56, 0xdd, 0x35, 0xac, 0x70, 0xbc, 0xc6, 0xbb, 0x39, 0xa4, 0x07, 0x5e, 0xa4, 0x2f, 0xc0,
	0x33, 0x70, 0xe1, 0xc4, 0x1b, 0xf0, 0x40, 0x68, 0x77, 0x9d, 0x10, 0x1b, 0x9c, 0xb6, 0xaa, 0x04,
	0x9c, 0xe2, 0x9d, 0x71, 0x66, 0xbe, 0xdf, 0xcf, 0xcc, 0x1a, 0x7a, 0xe1, 0x74, 0xf4, 0x9a, 0xa6,
	0x64, 0x94, 0x17, 0x4c, 0x30, 0x77, 0xbb, 0x3c, 0x3e, 0xdc, 0x0d, 0xa7, 0xa3, 0x7d, 0xc4, 0xc9,
	0x84, 0x24, 0x34, 0x2b, 0xb3, 0xc1, 0x27, 0xe8, 0x87, 0x53, 0x99, 0x3e, 0x66, 0xef, 0x68, 0x16,
	0x91, 0x8f, 0xee, 0x03, 0xd8, 0x9e, 0x73, 0x52, 0xc4, 0x14, 0x7b, 0x86, 0x6f, 0x0e, 0x7b, 0x51,
	0x5b, 0x1e, 0x43, 0x2c, 0x13, 0x02, 0xf1, 0x0f, 0x32, 0x61, 0xfa, 0xe6, 0xd0, 0x8e, 0xda, 0xf2,
	0x18, 0x62, 0xf7, 0x25, 0xd8, 0x09, 0x4d, 0x49, 0x5c, 0xb0, 0x94, 0x78, 0x96, 0x6f, 0x0e, 0xfb,
	0xe3, 0x27, 0xa3, 0x6a, 0xb3, 0x83, 0x94, 0x92, 0x4c, 0xc8, 0x3e, 0x11, 0x4b, 0x49, 0xd4, 0x49,
	0xca, 0xa7, 0xe0, 0xa8, 0xda, 0x9f, 0xe7, 0xee, 0x00, 0xba, 0x05, 0xe1, 0xf3, 0x54, 0xc4, 0x17,
	0x0c, 0x93, 0x52, 0x03, 0xe8, 0xd0, 0x01, 0xc3, 0xa4, 0x51, 0x47, 0x30, 0x87, 0xae, 0xae, 0x75,
	0x26, 0x90, 0x20, 0xee, 0x73, 0x68, 0x71, 0xf9, 0xa0, 0x4a, 0xf4, 0xc7, 0x4f, 0x1b, 0x25, 0xa9,
	0xd7, 0x23, 0xfd, 0x72, 0xb3, 0xcb, 0x35, 0x2e, 0xd6, 0x3a, 0x97, 0xe0, 0x8b, 0x01, 0x3b, 0xba,
	0xef, 0x9b, 0x79, 0x9a, 0x4e, 0x90, 0x40, 0x25, 0xc6, 0x65, 0x1d, 0xa3, 0xa9, 0x8e, 0x59, 0xe1,
	0xfb, 0x0a, 0x40, 0x14, 0x28, 0xe3, 0xf1, 0x4c, 0xfa, 0xd6, 0x1c, 0x07, 0x35, 0xd1, 0xe7, 0xf2,
	0x85, 0x84, 0x14, 0xb2, 0xdb, 0xf9, 0x22, 0x27, 0x91, 0xad, 0xfe, 0x32, 0x95, 0x5c, 0xee, 0x43,
	0x9b, 0x25, 0x09, 0x27, 0xc2, 0xdb, 0xd2, 0x75, 0xf5, 0xc9, 0x7d, 0x04, 0x36, 0x46, 0x02, 0xc5,
	0x9c, 0x5e, 0x12, 0xaf, 0xa5, 0x52, 0x1d, 0x19, 0x38, 0xa3, 0x97, 0x24, 0xb8, 0xfa, 0x55, 0xfc,
	0x5d, 0x66, 0xd0, 0x48, 0x69, 0x93, 0x3a, 0xb5, 0x3c, 0x52, 0x91, 0x52, 0xe7, 0xe8, 0xed, 0x90,
	0x52, 0x82, 0x6f, 0x06, 0xd8, 0x5a, 0x9d, 0x44, 0xea, 0x83, 0x93, 0x14, 0x6c, 0x16, 0x57, 0xd7,
	0x13, 0x64, 0xec, 0xad, 0x6e, 0xf2, 0x18, 0x40, 0xb0, 0xb8, 0x8a, 0xb7, 0x23, 0x58, 0x99, 0x5d,
	0xb6, 0xca, 0xd0, 0x4c, 0xf3, 0xb5, 0x75, 0xab, 0x13, 0x34, 0x23, 0xab, 0xa4, 0xa2, 0xa4, 0x25,
	0xaa, 0xa4, 0xa4, 0x54, 0x1b, 0x4d, 0xeb, 0xb6, 0xa3, 0x09, 0xae, 0xcc, 0x95, 0x8f, 0x9b, 0xd0,
	0xad, 0x1b, 0x35, 0xaf, 0x31, 0x6a, 0x6d, 0x32, 0xba, 0x55, 0x33, 0xba, 0x36, 0xba, 0x56, 0x65,
	0x74, 0x2f, 0xc0, 0xa1, 0x79, 0x8c, 0x30, 0x2e, 0xe2, 0x94, 0x72, 0xe1, 0xb5, 0x7d, 0x6b, 0xd8,
	0x1d, 0xdf, 0xab, 0xd9, 0x0c, 0xf3, 0x3d, 0x8c, 0x8b, 0x08, 0xa8, 0xfa, 0x3d, 0xa6, 0x5c, 0xd4,
	0xe8, 0x6c, 0xdf, 0x9a, 0xce, 0x57, 0x13, 0x1c, 0x4d, 0xe7, 0x84, 0x09, 0x9a, 0x2c, 0xfe, 0xe2,
	0xa0, 0xff, 0x39, 0x38, 0xee, 0x33, 0xe8, 0xb1, 0x24, 0x49, 0x69, 0x46, 0xe2, 0x82, 0x20, 0xbc,
	0xf0, 0x3a, 0x4a, 0xb2, 0x53, 0x06, 0x23, 0x19, 0x0b, 0x4e, 0x61, 0x57, 0x03, 0x3c, 0x44, 0xfc,
	0x74, 0x99, 0xd8, 0xf0, 0x29, 0x1f, 0x40, 0x17, 0x09, 0x81, 0x2e, 0xde, 0xeb, 0x6b, 0xf7, 0xbf,
	0x6f, 0x0c, 0x9d, 0x08, 0x74, 0x48, 0x5d, 0xbc, 0xef, 0xc6, 0x6f, 0x2a, 0xf2, 0xbc, 0xb9, 0xe2,
	0x11, 0xec, 0x2c, 0x65, 0x2a, 0xba, 0x0a, 0x92, 0xa9, 0x20, 0xd5, 0x3f, 0xbc, 0x65, 0x39, 0x59,
	0x3c, 0xcc, 0x12, 0x16, 0xfd, 0xc7, 0x7e, 0x06, 0x14, 0xb2, 0x3a, 0x6b, 0xeb, 0xa6, 0xac, 0xaf,
	0xb5, 0xf5, 0x79, 0x65, 0x6b, 0x0f, 0xe3, 0x35, 0x50, 0x77, 0x5d, 0xb8, 0xb5, 0xb5, 0xb1, 0x2a,
	0x6b, 0xb3, 0xf1, 0x26, 0x56, 0x36, 0xb1, 0x55, 0xdd, 0xc4, 0x20, 0x5b, 0x2a, 0x9d, 0x90, 0xf4,
	0x0f, 0x28, 0xdd, 0x37, 0x0f, 0xad, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x46, 0x19, 0xf7, 0xc2,
	0x3a, 0x08, 0x00, 0x00,
}
