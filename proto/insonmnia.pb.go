// Code generated by protoc-gen-go. DO NOT EDIT.
// source: insonmnia.proto

package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TaskStatusReply_Status int32

const (
	TaskStatusReply_UNKNOWN  TaskStatusReply_Status = 0
	TaskStatusReply_SPOOLING TaskStatusReply_Status = 1
	TaskStatusReply_SPAWNING TaskStatusReply_Status = 2
	TaskStatusReply_RUNNING  TaskStatusReply_Status = 3
	TaskStatusReply_FINISHED TaskStatusReply_Status = 4
	TaskStatusReply_BROKEN   TaskStatusReply_Status = 5
)

var TaskStatusReply_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "SPOOLING",
	2: "SPAWNING",
	3: "RUNNING",
	4: "FINISHED",
	5: "BROKEN",
}
var TaskStatusReply_Status_value = map[string]int32{
	"UNKNOWN":  0,
	"SPOOLING": 1,
	"SPAWNING": 2,
	"RUNNING":  3,
	"FINISHED": 4,
	"BROKEN":   5,
}

func (x TaskStatusReply_Status) String() string {
	return proto.EnumName(TaskStatusReply_Status_name, int32(x))
}
func (TaskStatusReply_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{10, 0} }

type TaskLogsRequest_Type int32

const (
	TaskLogsRequest_STDOUT TaskLogsRequest_Type = 0
	TaskLogsRequest_STDERR TaskLogsRequest_Type = 1
	TaskLogsRequest_BOTH   TaskLogsRequest_Type = 2
)

var TaskLogsRequest_Type_name = map[int32]string{
	0: "STDOUT",
	1: "STDERR",
	2: "BOTH",
}
var TaskLogsRequest_Type_value = map[string]int32{
	"STDOUT": 0,
	"STDERR": 1,
	"BOTH":   2,
}

func (x TaskLogsRequest_Type) String() string {
	return proto.EnumName(TaskLogsRequest_Type_name, int32(x))
}
func (TaskLogsRequest_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{13, 0} }

type PingRequest struct {
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type PingReply struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *PingReply) Reset()                    { *m = PingReply{} }
func (m *PingReply) String() string            { return proto.CompactTextString(m) }
func (*PingReply) ProtoMessage()               {}
func (*PingReply) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *PingReply) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type CPUUsage struct {
	Total uint64 `protobuf:"varint,1,opt,name=total" json:"total,omitempty"`
}

func (m *CPUUsage) Reset()                    { *m = CPUUsage{} }
func (m *CPUUsage) String() string            { return proto.CompactTextString(m) }
func (*CPUUsage) ProtoMessage()               {}
func (*CPUUsage) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *CPUUsage) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type MemoryUsage struct {
	MaxUsage uint64 `protobuf:"varint,1,opt,name=maxUsage" json:"maxUsage,omitempty"`
}

func (m *MemoryUsage) Reset()                    { *m = MemoryUsage{} }
func (m *MemoryUsage) String() string            { return proto.CompactTextString(m) }
func (*MemoryUsage) ProtoMessage()               {}
func (*MemoryUsage) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *MemoryUsage) GetMaxUsage() uint64 {
	if m != nil {
		return m.MaxUsage
	}
	return 0
}

type NetworkUsage struct {
	TxBytes   uint64 `protobuf:"varint,1,opt,name=txBytes" json:"txBytes,omitempty"`
	RxBytes   uint64 `protobuf:"varint,2,opt,name=rxBytes" json:"rxBytes,omitempty"`
	TxPackets uint64 `protobuf:"varint,3,opt,name=txPackets" json:"txPackets,omitempty"`
	RxPackets uint64 `protobuf:"varint,4,opt,name=rxPackets" json:"rxPackets,omitempty"`
	TxErrors  uint64 `protobuf:"varint,5,opt,name=txErrors" json:"txErrors,omitempty"`
	RxErrors  uint64 `protobuf:"varint,6,opt,name=rxErrors" json:"rxErrors,omitempty"`
	TxDropped uint64 `protobuf:"varint,7,opt,name=txDropped" json:"txDropped,omitempty"`
	RxDropped uint64 `protobuf:"varint,8,opt,name=rxDropped" json:"rxDropped,omitempty"`
}

func (m *NetworkUsage) Reset()                    { *m = NetworkUsage{} }
func (m *NetworkUsage) String() string            { return proto.CompactTextString(m) }
func (*NetworkUsage) ProtoMessage()               {}
func (*NetworkUsage) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *NetworkUsage) GetTxBytes() uint64 {
	if m != nil {
		return m.TxBytes
	}
	return 0
}

func (m *NetworkUsage) GetRxBytes() uint64 {
	if m != nil {
		return m.RxBytes
	}
	return 0
}

func (m *NetworkUsage) GetTxPackets() uint64 {
	if m != nil {
		return m.TxPackets
	}
	return 0
}

func (m *NetworkUsage) GetRxPackets() uint64 {
	if m != nil {
		return m.RxPackets
	}
	return 0
}

func (m *NetworkUsage) GetTxErrors() uint64 {
	if m != nil {
		return m.TxErrors
	}
	return 0
}

func (m *NetworkUsage) GetRxErrors() uint64 {
	if m != nil {
		return m.RxErrors
	}
	return 0
}

func (m *NetworkUsage) GetTxDropped() uint64 {
	if m != nil {
		return m.TxDropped
	}
	return 0
}

func (m *NetworkUsage) GetRxDropped() uint64 {
	if m != nil {
		return m.RxDropped
	}
	return 0
}

type ResourceUsage struct {
	Cpu     *CPUUsage                `protobuf:"bytes,1,opt,name=cpu" json:"cpu,omitempty"`
	Memory  *MemoryUsage             `protobuf:"bytes,2,opt,name=memory" json:"memory,omitempty"`
	Network map[string]*NetworkUsage `protobuf:"bytes,3,rep,name=network" json:"network,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ResourceUsage) Reset()                    { *m = ResourceUsage{} }
func (m *ResourceUsage) String() string            { return proto.CompactTextString(m) }
func (*ResourceUsage) ProtoMessage()               {}
func (*ResourceUsage) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *ResourceUsage) GetCpu() *CPUUsage {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *ResourceUsage) GetMemory() *MemoryUsage {
	if m != nil {
		return m.Memory
	}
	return nil
}

func (m *ResourceUsage) GetNetwork() map[string]*NetworkUsage {
	if m != nil {
		return m.Network
	}
	return nil
}

type InfoReply struct {
	Usage        map[string]*ResourceUsage `protobuf:"bytes,1,rep,name=usage" json:"usage,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Name         string                    `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Capabilities *Capabilities             `protobuf:"bytes,3,opt,name=capabilities" json:"capabilities,omitempty"`
}

func (m *InfoReply) Reset()                    { *m = InfoReply{} }
func (m *InfoReply) String() string            { return proto.CompactTextString(m) }
func (*InfoReply) ProtoMessage()               {}
func (*InfoReply) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *InfoReply) GetUsage() map[string]*ResourceUsage {
	if m != nil {
		return m.Usage
	}
	return nil
}

func (m *InfoReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InfoReply) GetCapabilities() *Capabilities {
	if m != nil {
		return m.Capabilities
	}
	return nil
}

type StopTaskRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *StopTaskRequest) Reset()                    { *m = StopTaskRequest{} }
func (m *StopTaskRequest) String() string            { return proto.CompactTextString(m) }
func (*StopTaskRequest) ProtoMessage()               {}
func (*StopTaskRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

func (m *StopTaskRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type StopTaskReply struct {
}

func (m *StopTaskReply) Reset()                    { *m = StopTaskReply{} }
func (m *StopTaskReply) String() string            { return proto.CompactTextString(m) }
func (*StopTaskReply) ProtoMessage()               {}
func (*StopTaskReply) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

type TaskStatusRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *TaskStatusRequest) Reset()                    { *m = TaskStatusRequest{} }
func (m *TaskStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*TaskStatusRequest) ProtoMessage()               {}
func (*TaskStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{9} }

func (m *TaskStatusRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type TaskStatusReply struct {
	Status    TaskStatusReply_Status `protobuf:"varint,1,opt,name=status,enum=sonm.TaskStatusReply_Status" json:"status,omitempty"`
	ImageName string                 `protobuf:"bytes,2,opt,name=imageName" json:"imageName,omitempty"`
	Ports     string                 `protobuf:"bytes,3,opt,name=ports" json:"ports,omitempty"`
	Uptime    uint64                 `protobuf:"varint,4,opt,name=uptime" json:"uptime,omitempty"`
	Usage     *ResourceUsage         `protobuf:"bytes,5,opt,name=usage" json:"usage,omitempty"`
	MinerID   string                 `protobuf:"bytes,6,opt,name=minerID" json:"minerID,omitempty"`
}

func (m *TaskStatusReply) Reset()                    { *m = TaskStatusReply{} }
func (m *TaskStatusReply) String() string            { return proto.CompactTextString(m) }
func (*TaskStatusReply) ProtoMessage()               {}
func (*TaskStatusReply) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{10} }

func (m *TaskStatusReply) GetStatus() TaskStatusReply_Status {
	if m != nil {
		return m.Status
	}
	return TaskStatusReply_UNKNOWN
}

func (m *TaskStatusReply) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
}

func (m *TaskStatusReply) GetPorts() string {
	if m != nil {
		return m.Ports
	}
	return ""
}

func (m *TaskStatusReply) GetUptime() uint64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *TaskStatusReply) GetUsage() *ResourceUsage {
	if m != nil {
		return m.Usage
	}
	return nil
}

func (m *TaskStatusReply) GetMinerID() string {
	if m != nil {
		return m.MinerID
	}
	return ""
}

type StatusMapReply struct {
	Statuses map[string]*TaskStatusReply `protobuf:"bytes,1,rep,name=statuses" json:"statuses,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *StatusMapReply) Reset()                    { *m = StatusMapReply{} }
func (m *StatusMapReply) String() string            { return proto.CompactTextString(m) }
func (*StatusMapReply) ProtoMessage()               {}
func (*StatusMapReply) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{11} }

func (m *StatusMapReply) GetStatuses() map[string]*TaskStatusReply {
	if m != nil {
		return m.Statuses
	}
	return nil
}

type ContainerRestartPolicy struct {
	Name              string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	MaximumRetryCount uint32 `protobuf:"varint,2,opt,name=maximumRetryCount" json:"maximumRetryCount,omitempty"`
}

func (m *ContainerRestartPolicy) Reset()                    { *m = ContainerRestartPolicy{} }
func (m *ContainerRestartPolicy) String() string            { return proto.CompactTextString(m) }
func (*ContainerRestartPolicy) ProtoMessage()               {}
func (*ContainerRestartPolicy) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{12} }

func (m *ContainerRestartPolicy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContainerRestartPolicy) GetMaximumRetryCount() uint32 {
	if m != nil {
		return m.MaximumRetryCount
	}
	return 0
}

type TaskLogsRequest struct {
	Type          TaskLogsRequest_Type `protobuf:"varint,1,opt,name=type,enum=sonm.TaskLogsRequest_Type" json:"type,omitempty"`
	Id            string               `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Since         string               `protobuf:"bytes,3,opt,name=since" json:"since,omitempty"`
	AddTimestamps bool                 `protobuf:"varint,4,opt,name=addTimestamps" json:"addTimestamps,omitempty"`
	Follow        bool                 `protobuf:"varint,5,opt,name=Follow" json:"Follow,omitempty"`
	Tail          string               `protobuf:"bytes,6,opt,name=Tail" json:"Tail,omitempty"`
	Details       bool                 `protobuf:"varint,7,opt,name=Details" json:"Details,omitempty"`
}

func (m *TaskLogsRequest) Reset()                    { *m = TaskLogsRequest{} }
func (m *TaskLogsRequest) String() string            { return proto.CompactTextString(m) }
func (*TaskLogsRequest) ProtoMessage()               {}
func (*TaskLogsRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{13} }

func (m *TaskLogsRequest) GetType() TaskLogsRequest_Type {
	if m != nil {
		return m.Type
	}
	return TaskLogsRequest_STDOUT
}

func (m *TaskLogsRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TaskLogsRequest) GetSince() string {
	if m != nil {
		return m.Since
	}
	return ""
}

func (m *TaskLogsRequest) GetAddTimestamps() bool {
	if m != nil {
		return m.AddTimestamps
	}
	return false
}

func (m *TaskLogsRequest) GetFollow() bool {
	if m != nil {
		return m.Follow
	}
	return false
}

func (m *TaskLogsRequest) GetTail() string {
	if m != nil {
		return m.Tail
	}
	return ""
}

func (m *TaskLogsRequest) GetDetails() bool {
	if m != nil {
		return m.Details
	}
	return false
}

type TaskLogsChunk struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *TaskLogsChunk) Reset()                    { *m = TaskLogsChunk{} }
func (m *TaskLogsChunk) String() string            { return proto.CompactTextString(m) }
func (*TaskLogsChunk) ProtoMessage()               {}
func (*TaskLogsChunk) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{14} }

func (m *TaskLogsChunk) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type TaskResourceRequirements struct {
	// Number of CPU cores.
	CPUCores uint64 `protobuf:"varint,1,opt,name=CPUCores" json:"CPUCores,omitempty"`
	// Memory in bytes required.
	MaxMemory int64 `protobuf:"varint,2,opt,name=maxMemory" json:"maxMemory,omitempty"`
	// Describes whether a task requires GPU supoort.
	GPUSupport bool  `protobuf:"varint,3,opt,name=GPUSupport" json:"GPUSupport,omitempty"`
	NanoCPUs   int64 `protobuf:"varint,4,opt,name=nanoCPUs" json:"nanoCPUs,omitempty"`
}

func (m *TaskResourceRequirements) Reset()                    { *m = TaskResourceRequirements{} }
func (m *TaskResourceRequirements) String() string            { return proto.CompactTextString(m) }
func (*TaskResourceRequirements) ProtoMessage()               {}
func (*TaskResourceRequirements) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{15} }

func (m *TaskResourceRequirements) GetCPUCores() uint64 {
	if m != nil {
		return m.CPUCores
	}
	return 0
}

func (m *TaskResourceRequirements) GetMaxMemory() int64 {
	if m != nil {
		return m.MaxMemory
	}
	return 0
}

func (m *TaskResourceRequirements) GetGPUSupport() bool {
	if m != nil {
		return m.GPUSupport
	}
	return false
}

func (m *TaskResourceRequirements) GetNanoCPUs() int64 {
	if m != nil {
		return m.NanoCPUs
	}
	return 0
}

type Timestamp struct {
	// Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z.
	// Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive.
	Seconds int64 `protobuf:"varint,1,opt,name=seconds" json:"seconds,omitempty"`
	// Non-negative fractions of a second at nanosecond resolution. Negative
	// second values with fractions must still have non-negative nanos values
	// that count forward in time. Must be from 0 to 999'999'999
	// inclusive.
	Nanos int32 `protobuf:"varint,2,opt,name=nanos" json:"nanos,omitempty"`
}

func (m *Timestamp) Reset()                    { *m = Timestamp{} }
func (m *Timestamp) String() string            { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()               {}
func (*Timestamp) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{16} }

func (m *Timestamp) GetSeconds() int64 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *Timestamp) GetNanos() int32 {
	if m != nil {
		return m.Nanos
	}
	return 0
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "sonm.PingRequest")
	proto.RegisterType((*PingReply)(nil), "sonm.PingReply")
	proto.RegisterType((*CPUUsage)(nil), "sonm.CPUUsage")
	proto.RegisterType((*MemoryUsage)(nil), "sonm.MemoryUsage")
	proto.RegisterType((*NetworkUsage)(nil), "sonm.NetworkUsage")
	proto.RegisterType((*ResourceUsage)(nil), "sonm.ResourceUsage")
	proto.RegisterType((*InfoReply)(nil), "sonm.InfoReply")
	proto.RegisterType((*StopTaskRequest)(nil), "sonm.StopTaskRequest")
	proto.RegisterType((*StopTaskReply)(nil), "sonm.StopTaskReply")
	proto.RegisterType((*TaskStatusRequest)(nil), "sonm.TaskStatusRequest")
	proto.RegisterType((*TaskStatusReply)(nil), "sonm.TaskStatusReply")
	proto.RegisterType((*StatusMapReply)(nil), "sonm.StatusMapReply")
	proto.RegisterType((*ContainerRestartPolicy)(nil), "sonm.ContainerRestartPolicy")
	proto.RegisterType((*TaskLogsRequest)(nil), "sonm.TaskLogsRequest")
	proto.RegisterType((*TaskLogsChunk)(nil), "sonm.TaskLogsChunk")
	proto.RegisterType((*TaskResourceRequirements)(nil), "sonm.TaskResourceRequirements")
	proto.RegisterType((*Timestamp)(nil), "sonm.Timestamp")
	proto.RegisterEnum("sonm.TaskStatusReply_Status", TaskStatusReply_Status_name, TaskStatusReply_Status_value)
	proto.RegisterEnum("sonm.TaskLogsRequest_Type", TaskLogsRequest_Type_name, TaskLogsRequest_Type_value)
}

func init() { proto.RegisterFile("insonmnia.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 953 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0xdd, 0x8e, 0xe3, 0x34,
	0x14, 0xde, 0xf4, 0x6f, 0xd2, 0xd3, 0xe9, 0x4c, 0xc7, 0xc0, 0xaa, 0xaa, 0x56, 0xa8, 0x64, 0xb8,
	0x98, 0x11, 0xa8, 0x42, 0x03, 0x42, 0x68, 0x91, 0x90, 0x98, 0xb6, 0xbb, 0x5b, 0xed, 0x4e, 0x5a,
	0xb9, 0xad, 0x16, 0x71, 0xe7, 0x6d, 0xcd, 0x10, 0x35, 0x89, 0x83, 0xe3, 0x30, 0xcd, 0x5b, 0xf0,
	0x10, 0x3c, 0x01, 0x4f, 0xc3, 0x23, 0x70, 0xcd, 0x3d, 0x12, 0xb2, 0x8f, 0x93, 0xa6, 0x30, 0x70,
	0xe7, 0xcf, 0xdf, 0xe7, 0xf8, 0xfc, 0x7c, 0x27, 0x86, 0xf3, 0x20, 0x4e, 0x45, 0x1c, 0xc5, 0x01,
	0x1b, 0x25, 0x52, 0x28, 0x41, 0x1a, 0x1a, 0x0e, 0xc8, 0x86, 0x25, 0xec, 0x5d, 0x10, 0x06, 0x2a,
	0xe0, 0x29, 0x32, 0x5e, 0x17, 0x3a, 0x8b, 0x20, 0xbe, 0xa7, 0xfc, 0xa7, 0x8c, 0xa7, 0xca, 0xbb,
	0x84, 0x36, 0xc2, 0x24, 0xcc, 0xc9, 0x53, 0x68, 0xa5, 0x8a, 0xa9, 0x2c, 0xed, 0x3b, 0x43, 0xe7,
	0xaa, 0x4d, 0x2d, 0xf2, 0x86, 0xe0, 0x8e, 0x17, 0xeb, 0x75, 0xca, 0xee, 0x39, 0x79, 0x1f, 0x9a,
	0x4a, 0x28, 0x16, 0x1a, 0x49, 0x83, 0x22, 0xf0, 0xae, 0xa1, 0x73, 0xc7, 0x23, 0x21, 0x73, 0x14,
	0x0d, 0xc0, 0x8d, 0xd8, 0xde, 0xac, 0xad, 0xae, 0xc4, 0xde, 0x9f, 0x0e, 0x9c, 0xfa, 0x5c, 0x3d,
	0x08, 0xb9, 0x43, 0x71, 0x1f, 0x4e, 0xd4, 0xfe, 0x36, 0x57, 0x3c, 0xb5, 0xda, 0x02, 0x6a, 0x46,
	0x5a, 0xa6, 0x86, 0x8c, 0x85, 0xe4, 0x19, 0xb4, 0xd5, 0x7e, 0xc1, 0x36, 0x3b, 0xae, 0xd2, 0x7e,
	0xdd, 0x70, 0x87, 0x0d, 0xcd, 0xca, 0x92, 0x6d, 0x20, 0x5b, 0x6e, 0xe8, 0xe0, 0xd4, 0x7e, 0x2a,
	0xa5, 0x90, 0x69, 0xbf, 0x89, 0xc1, 0x15, 0x58, 0x73, 0xb2, 0xe0, 0x5a, 0xc8, 0x15, 0x18, 0xef,
	0x9c, 0x48, 0x91, 0x24, 0x7c, 0xdb, 0x3f, 0x29, 0xee, 0xb4, 0x1b, 0x78, 0x67, 0xc1, 0xba, 0xc5,
	0x9d, 0x76, 0xc3, 0xfb, 0xc3, 0x81, 0x2e, 0xe5, 0xa9, 0xc8, 0xe4, 0x86, 0x63, 0xd6, 0x43, 0xa8,
	0x6f, 0x92, 0xcc, 0x64, 0xdc, 0xb9, 0x39, 0x1b, 0xe9, 0x7e, 0x8d, 0x8a, 0x22, 0x53, 0x4d, 0x91,
	0x6b, 0x68, 0x45, 0xa6, 0xa6, 0x26, 0xf9, 0xce, 0xcd, 0x05, 0x8a, 0x2a, 0x75, 0xa6, 0x56, 0x40,
	0x9e, 0xc3, 0x49, 0x8c, 0x25, 0xed, 0xd7, 0x87, 0xf5, 0xab, 0xce, 0xcd, 0x10, 0xb5, 0x47, 0x57,
	0x8e, 0x6c, 0xd5, 0xa7, 0xb1, 0x92, 0x39, 0x2d, 0x0e, 0x0c, 0xfc, 0xb2, 0x1d, 0x86, 0x20, 0x3d,
	0xa8, 0xef, 0x78, 0x6e, 0x1d, 0xa0, 0x97, 0xe4, 0x0a, 0x9a, 0x3f, 0xb3, 0x30, 0xe3, 0x36, 0x0e,
	0x82, 0xdf, 0xae, 0xf6, 0x90, 0xa2, 0xe0, 0x79, 0xed, 0x2b, 0xc7, 0xfb, 0xdd, 0x81, 0xf6, 0x2c,
	0xfe, 0x41, 0xa0, 0xa5, 0x3e, 0x83, 0x66, 0x66, 0x6d, 0xa0, 0xe3, 0x1a, 0xe0, 0xd9, 0x92, 0x1f,
	0x99, 0xe3, 0x18, 0x11, 0x0a, 0x09, 0x81, 0x46, 0xcc, 0x22, 0xbc, 0xac, 0x4d, 0xcd, 0x9a, 0x7c,
	0x09, 0xa7, 0x55, 0x2b, 0x9b, 0x8e, 0x97, 0x81, 0x8c, 0x2b, 0x0c, 0x3d, 0xd2, 0x0d, 0xee, 0x00,
	0x0e, 0x17, 0x3c, 0x92, 0xd9, 0xf5, 0x71, 0x66, 0xef, 0x3d, 0x52, 0xb5, 0x6a, 0x6a, 0x1f, 0xc1,
	0xf9, 0x52, 0x89, 0x64, 0xc5, 0xd2, 0x9d, 0x9d, 0x1f, 0x72, 0x06, 0xb5, 0x60, 0x6b, 0x3f, 0x59,
	0x0b, 0xb6, 0xde, 0x39, 0x74, 0x0f, 0x92, 0x24, 0xcc, 0xbd, 0x4b, 0xb8, 0xd0, 0x60, 0x69, 0x26,
	0xe9, 0xbf, 0x4e, 0xfd, 0x56, 0x83, 0xf3, 0xaa, 0x4a, 0x57, 0xee, 0x8b, 0xa3, 0x61, 0x3c, 0xbb,
	0x79, 0x86, 0xc1, 0xfd, 0x43, 0x36, 0xb2, 0x6b, 0xab, 0xd5, 0x36, 0x0c, 0x22, 0x76, 0xcf, 0xfd,
	0x43, 0x09, 0x0f, 0x1b, 0x7a, 0x78, 0x13, 0x21, 0xed, 0xc8, 0xb4, 0x29, 0x02, 0x3d, 0xf6, 0x59,
	0xa2, 0x82, 0x88, 0xdb, 0x59, 0xb1, 0x48, 0x57, 0x07, 0x7b, 0xd7, 0xfc, 0x9f, 0xea, 0x64, 0xc5,
	0x0c, 0x47, 0x41, 0xcc, 0xe5, 0x6c, 0x62, 0xc6, 0xa6, 0x4d, 0x0b, 0xe8, 0x7d, 0x07, 0x2d, 0x0c,
	0x91, 0x74, 0xe0, 0x64, 0xed, 0xbf, 0xf6, 0xe7, 0x6f, 0xfd, 0xde, 0x13, 0x72, 0x0a, 0xee, 0x72,
	0x31, 0x9f, 0xbf, 0x99, 0xf9, 0x2f, 0x7b, 0x0e, 0xa2, 0x6f, 0xdf, 0xfa, 0x1a, 0xd5, 0xb4, 0x90,
	0xae, 0x7d, 0x03, 0xea, 0x9a, 0x7a, 0x31, 0xf3, 0x67, 0xcb, 0x57, 0xd3, 0x49, 0xaf, 0x41, 0x00,
	0x5a, 0xb7, 0x74, 0xfe, 0x7a, 0xea, 0xf7, 0x9a, 0xde, 0xaf, 0x0e, 0x9c, 0xe1, 0xa7, 0xef, 0x58,
	0x82, 0x35, 0xfb, 0x06, 0x5c, 0xac, 0x83, 0xf9, 0x97, 0x68, 0xc3, 0x79, 0x18, 0xf4, 0xb1, 0xce,
	0x42, 0x9e, 0xa2, 0xf1, 0xca, 0x33, 0x03, 0xaa, 0xbb, 0x57, 0xa1, 0x1e, 0xb1, 0xcc, 0x27, 0xc7,
	0x96, 0xf9, 0xe0, 0xd1, 0xae, 0x54, 0x4d, 0xf3, 0x3d, 0x3c, 0x1d, 0x8b, 0x58, 0x31, 0x5d, 0x0f,
	0xca, 0x53, 0xc5, 0xa4, 0x5a, 0x88, 0x30, 0xd8, 0xe4, 0xa5, 0xd3, 0x9d, 0x8a, 0xd3, 0x3f, 0x85,
	0x8b, 0x88, 0xed, 0x83, 0x28, 0x8b, 0x28, 0x57, 0x32, 0x1f, 0x8b, 0x2c, 0x56, 0xe6, 0xaa, 0x2e,
	0xfd, 0x37, 0xe1, 0xfd, 0xe5, 0xa0, 0x6f, 0xde, 0x88, 0xfb, 0xd2, 0x5b, 0x23, 0x68, 0xa8, 0x3c,
	0xe1, 0xd6, 0x35, 0x83, 0x43, 0x7c, 0x15, 0xd1, 0x68, 0x95, 0x27, 0x9c, 0x1a, 0x9d, 0xf5, 0x62,
	0xad, 0xf0, 0xa2, 0xf6, 0x48, 0x1a, 0xc4, 0x1b, 0x5e, 0x78, 0xc4, 0x00, 0xf2, 0x31, 0x74, 0xd9,
	0x76, 0xbb, 0x0a, 0x22, 0x9d, 0x41, 0x94, 0xe0, 0x6f, 0xd5, 0xa5, 0xc7, 0x9b, 0xda, 0x49, 0x2f,
	0x44, 0x18, 0x8a, 0x07, 0x63, 0x19, 0x97, 0x5a, 0xa4, 0x33, 0x5d, 0xb1, 0x20, 0xb4, 0xde, 0x30,
	0x6b, 0x6d, 0x99, 0x09, 0x57, 0x2c, 0x08, 0x53, 0xf3, 0x33, 0x75, 0x69, 0x01, 0xbd, 0x2b, 0x68,
	0xe8, 0xf8, 0x74, 0xb3, 0x97, 0xab, 0xc9, 0x7c, 0xbd, 0xea, 0x3d, 0xb1, 0xeb, 0x29, 0xa5, 0x3d,
	0x87, 0xb8, 0xd0, 0xb8, 0x9d, 0xaf, 0x5e, 0xf5, 0x6a, 0xde, 0x25, 0x74, 0x8b, 0xcc, 0xc6, 0x3f,
	0x66, 0xf1, 0x4e, 0x5f, 0xb4, 0x65, 0x8a, 0x99, 0xe4, 0x4f, 0xa9, 0x59, 0x7b, 0xbf, 0x38, 0xd0,
	0xc7, 0x79, 0x44, 0xe3, 0xea, 0x1a, 0x04, 0x92, 0x47, 0x3c, 0xc6, 0xc7, 0x60, 0xbc, 0x58, 0x8f,
	0x85, 0x2c, 0x5f, 0x9f, 0x12, 0xeb, 0x59, 0x8a, 0xd8, 0xfe, 0xee, 0xf0, 0x0f, 0xae, 0xd3, 0xc3,
	0x06, 0xf9, 0x10, 0xe0, 0xe5, 0x62, 0xbd, 0xcc, 0x12, 0x3d, 0x44, 0xa6, 0x58, 0x2e, 0xad, 0xec,
	0xe8, 0x2f, 0xc7, 0x2c, 0x16, 0xe3, 0xc5, 0x1a, 0x8b, 0x55, 0xa7, 0x25, 0xf6, 0xbe, 0x86, 0x76,
	0x59, 0x35, 0x5d, 0x88, 0x94, 0x6f, 0x44, 0xbc, 0xc5, 0x08, 0xea, 0xb4, 0x80, 0xba, 0x15, 0xfa,
	0x08, 0xbe, 0x7e, 0x4d, 0x8a, 0xe0, 0x5d, 0xcb, 0x3c, 0xe4, 0x9f, 0xff, 0x1d, 0x00, 0x00, 0xff,
	0xff, 0x5b, 0x44, 0x9c, 0x01, 0xf5, 0x07, 0x00, 0x00,
}
