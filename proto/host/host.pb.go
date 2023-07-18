// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: host/host.proto

package host

import (
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	network "github.com/maxlandon/aims/proto/network"
	nmap "github.com/maxlandon/aims/proto/scan/nmap"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Host - Represents a remote computer host in DB
type Host struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: display:"ID" readonly:"true" strict:"yes"
	// int64 Id = 1 [(gorm.field).tag = {primary_key: true}];
	// string Id = 1 [(gorm.field).tag = {type:"uuid" primary_key: true }];
	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty" display:"ID" readonly:"true" strict:"yes"`
	// @gotags: display:"Created at" readonly:"true" xml:"-"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" display:"Created at" readonly:"true" xml:"-"`
	// @gotags: display:"Updated at" readonly:"true" xml:"-"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty" display:"Updated at" readonly:"true" xml:"-"`
	// General ------------------------------------
	// @gotags: display:"MAC Address"
	MAC string `protobuf:"bytes,10,opt,name=MAC,proto3" json:"MAC,omitempty" display:"MAC Address"`
	// @gotags: display:"Comm"
	Comm string `protobuf:"bytes,11,opt,name=Comm,proto3" json:"Comm,omitempty" display:"Comm"`
	// @gotags: display:"OS Name"
	OSName string `protobuf:"bytes,12,opt,name=OSName,proto3" json:"OSName,omitempty" display:"OS Name"`
	// @gotags: display:"OS Flavor"
	OSFlavor string `protobuf:"bytes,13,opt,name=OSFlavor,proto3" json:"OSFlavor,omitempty" display:"OS Flavor"`
	// @gotags: display:"OS Service Pack"
	OSSp string `protobuf:"bytes,14,opt,name=OSSp,proto3" json:"OSSp,omitempty" display:"OS Service Pack"`
	// @gotags: display:"OS Language"
	OSLang string `protobuf:"bytes,15,opt,name=OSLang,proto3" json:"OSLang,omitempty" display:"OS Language"`
	// @gotags: display:"OS Family"
	OSFamily string `protobuf:"bytes,16,opt,name=OSFamily,proto3" json:"OSFamily,omitempty" display:"OS Family"`
	// @gotags: display:"CPU Arch"
	Arch string `protobuf:"bytes,17,opt,name=Arch,proto3" json:"Arch,omitempty" display:"CPU Arch"`
	// @gotags: display:"Purpose"
	Purpose string `protobuf:"bytes,18,opt,name=Purpose,proto3" json:"Purpose,omitempty" display:"Purpose"`
	// @gotags: display:"Info"
	Info string `protobuf:"bytes,19,opt,name=Info,proto3" json:"Info,omitempty" display:"Info"`
	// @gotags: display:"Scope"
	Scope string `protobuf:"bytes,20,opt,name=Scope,proto3" json:"Scope,omitempty" display:"Scope"`
	// @gotags: display:"Virtual Host"
	VirtualHost string `protobuf:"bytes,21,opt,name=VirtualHost,proto3" json:"VirtualHost,omitempty" display:"Virtual Host"`
	// @gotags: display:"Users"
	Users []*User `protobuf:"bytes,22,rep,name=Users,proto3" json:"Users,omitempty" display:"Users"`
	// @gotags: display:"Processes"
	Processes []*Process `protobuf:"bytes,23,rep,name=Processes,proto3" json:"Processes,omitempty" display:"Processes"`
	// @gotags: display:"Filesystem"
	FS *FileSystem `protobuf:"bytes,24,opt,name=FS,proto3" json:"FS,omitempty" display:"Filesystem"`
	// @gotags: xml:"address"
	Addresses []*network.Address `protobuf:"bytes,25,rep,name=Addresses,proto3" json:"Addresses,omitempty" xml:"address"`
	// Nmap ---------------------------------------
	// @gotags: xml:"os"
	OS *OS `protobuf:"bytes,30,opt,name=OS,proto3" json:"OS,omitempty" xml:"os"`
	// @gotags: xml:"status"
	Status *Status `protobuf:"bytes,31,opt,name=Status,proto3" json:"Status,omitempty" xml:"status"`
	// @gotags: xml:"distance"
	Distance *network.Distance `protobuf:"bytes,32,opt,name=Distance,proto3" json:"Distance,omitempty" xml:"distance"`
	// @gotags: xml:"starttime,attr,omitempty"
	StartTime int64 `protobuf:"varint,33,opt,name=StartTime,proto3" json:"StartTime,omitempty" xml:"starttime,attr,omitempty"`
	// @gotags: xml:"endtime,attr,omitempty"
	EndTime int64 `protobuf:"varint,34,opt,name=EndTime,proto3" json:"EndTime,omitempty" xml:"endtime,attr,omitempty"` // Might have issues here with XML unmarshalling
	// @gotags: xml:"ipidsequence"
	IPIDSequence *network.IPIDSequence `protobuf:"bytes,35,opt,name=IPIDSequence,proto3" json:"IPIDSequence,omitempty" xml:"ipidsequence"`
	// @gotags: xml:"tcpsequence"
	TCPSequence *network.TCPSequence `protobuf:"bytes,36,opt,name=TCPSequence,proto3" json:"TCPSequence,omitempty" xml:"tcpsequence"`
	// @gotags: xml:"tcptssequence"
	TCPTSSequence *network.TCPTSSequence `protobuf:"bytes,37,opt,name=TCPTSSequence,proto3" json:"TCPTSSequence,omitempty" xml:"tcptssequence"`
	ICMPResponse  *network.ICMPResponse  `protobuf:"bytes,38,opt,name=ICMPResponse,proto3" json:"ICMPResponse,omitempty"`
	// @gotags: xml:"times"
	Testing *network.Times `protobuf:"bytes,39,opt,name=Testing,proto3" json:"Testing,omitempty" xml:"times"`
	// @gotags: xml:"trace"
	Trace *network.Trace `protobuf:"bytes,40,opt,name=Trace,proto3" json:"Trace,omitempty" xml:"trace"`
	// @gotags: xml:"uptime"
	Uptime *Uptime `protobuf:"bytes,41,opt,name=Uptime,proto3" json:"Uptime,omitempty" xml:"uptime"`
	// @gotags: xml:"comment,attr"
	Comment string `protobuf:"bytes,42,opt,name=Comment,proto3" json:"Comment,omitempty" xml:"comment,attr"`
	// @gotags: xml:"hostscript>script"
	HostScripts []*nmap.Script `protobuf:"bytes,43,rep,name=HostScripts,proto3" json:"HostScripts,omitempty" xml:"hostscript>script"`
	// @gotags: xml:"smurf"
	Smurfs []*nmap.Smurf `protobuf:"bytes,44,rep,name=Smurfs,proto3" json:"Smurfs,omitempty" xml:"smurf"`
	// @gotags: xml:"hostnames>hostname"
	Hostnames []*Hostname `protobuf:"bytes,45,rep,name=Hostnames,proto3" json:"Hostnames,omitempty" xml:"hostnames>hostname"`
	// @gotags: xml:"ports>port"
	Ports []*Port `protobuf:"bytes,46,rep,name=Ports,proto3" json:"Ports,omitempty" xml:"ports>port"`
	// @gotags: xml:"ports>extraports"
	ExtraPorts []*ExtraPort `protobuf:"bytes,47,rep,name=ExtraPorts,proto3" json:"ExtraPorts,omitempty" xml:"ports>extraports"`
}

func (x *Host) Reset() {
	*x = Host{}
	if protoimpl.UnsafeEnabled {
		mi := &file_host_host_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Host) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Host) ProtoMessage() {}

func (x *Host) ProtoReflect() protoreflect.Message {
	mi := &file_host_host_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Host.ProtoReflect.Descriptor instead.
func (*Host) Descriptor() ([]byte, []int) {
	return file_host_host_proto_rawDescGZIP(), []int{0}
}

func (x *Host) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Host) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Host) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Host) GetMAC() string {
	if x != nil {
		return x.MAC
	}
	return ""
}

func (x *Host) GetComm() string {
	if x != nil {
		return x.Comm
	}
	return ""
}

func (x *Host) GetOSName() string {
	if x != nil {
		return x.OSName
	}
	return ""
}

func (x *Host) GetOSFlavor() string {
	if x != nil {
		return x.OSFlavor
	}
	return ""
}

func (x *Host) GetOSSp() string {
	if x != nil {
		return x.OSSp
	}
	return ""
}

func (x *Host) GetOSLang() string {
	if x != nil {
		return x.OSLang
	}
	return ""
}

func (x *Host) GetOSFamily() string {
	if x != nil {
		return x.OSFamily
	}
	return ""
}

func (x *Host) GetArch() string {
	if x != nil {
		return x.Arch
	}
	return ""
}

func (x *Host) GetPurpose() string {
	if x != nil {
		return x.Purpose
	}
	return ""
}

func (x *Host) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *Host) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *Host) GetVirtualHost() string {
	if x != nil {
		return x.VirtualHost
	}
	return ""
}

func (x *Host) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *Host) GetProcesses() []*Process {
	if x != nil {
		return x.Processes
	}
	return nil
}

func (x *Host) GetFS() *FileSystem {
	if x != nil {
		return x.FS
	}
	return nil
}

func (x *Host) GetAddresses() []*network.Address {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *Host) GetOS() *OS {
	if x != nil {
		return x.OS
	}
	return nil
}

func (x *Host) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *Host) GetDistance() *network.Distance {
	if x != nil {
		return x.Distance
	}
	return nil
}

func (x *Host) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *Host) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *Host) GetIPIDSequence() *network.IPIDSequence {
	if x != nil {
		return x.IPIDSequence
	}
	return nil
}

func (x *Host) GetTCPSequence() *network.TCPSequence {
	if x != nil {
		return x.TCPSequence
	}
	return nil
}

func (x *Host) GetTCPTSSequence() *network.TCPTSSequence {
	if x != nil {
		return x.TCPTSSequence
	}
	return nil
}

func (x *Host) GetICMPResponse() *network.ICMPResponse {
	if x != nil {
		return x.ICMPResponse
	}
	return nil
}

func (x *Host) GetTesting() *network.Times {
	if x != nil {
		return x.Testing
	}
	return nil
}

func (x *Host) GetTrace() *network.Trace {
	if x != nil {
		return x.Trace
	}
	return nil
}

func (x *Host) GetUptime() *Uptime {
	if x != nil {
		return x.Uptime
	}
	return nil
}

func (x *Host) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *Host) GetHostScripts() []*nmap.Script {
	if x != nil {
		return x.HostScripts
	}
	return nil
}

func (x *Host) GetSmurfs() []*nmap.Smurf {
	if x != nil {
		return x.Smurfs
	}
	return nil
}

func (x *Host) GetHostnames() []*Hostname {
	if x != nil {
		return x.Hostnames
	}
	return nil
}

func (x *Host) GetPorts() []*Port {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *Host) GetExtraPorts() []*ExtraPort {
	if x != nil {
		return x.ExtraPorts
	}
	return nil
}

// A hostname for a host.
type Hostname struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: display:"ID" readonly:"true" strict:"yes"
	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty" display:"ID" readonly:"true" strict:"yes"`
	// @gotags: display:"Created at" readonly:"true" xml:"-"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" display:"Created at" readonly:"true" xml:"-"`
	// @gotags: display:"Updated at" readonly:"true" xml:"-"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty" display:"Updated at" readonly:"true" xml:"-"`
	// @gotags: display:"Name" xml:"name,attr"
	Name string `protobuf:"bytes,10,opt,name=Name,proto3" json:"Name,omitempty" display:"Name" xml:"name,attr"`
	// @gotags: display:"Type" xml:"type,attr"
	Type string `protobuf:"bytes,11,opt,name=Type,proto3" json:"Type,omitempty" display:"Type" xml:"type,attr"`
}

func (x *Hostname) Reset() {
	*x = Hostname{}
	if protoimpl.UnsafeEnabled {
		mi := &file_host_host_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hostname) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hostname) ProtoMessage() {}

func (x *Hostname) ProtoReflect() protoreflect.Message {
	mi := &file_host_host_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hostname.ProtoReflect.Descriptor instead.
func (*Hostname) Descriptor() ([]byte, []int) {
	return file_host_host_proto_rawDescGZIP(), []int{1}
}

func (x *Hostname) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Hostname) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Hostname) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Hostname) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Hostname) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

// Uptime - The amount of time the host has been up
type Uptime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,100,opt,name=Id,proto3" json:"Id,omitempty"`
	// @gotags: xml:"seconds,attr"
	Seconds int32 `protobuf:"varint,1,opt,name=Seconds,proto3" json:"Seconds,omitempty" xml:"seconds,attr"`
	// @gotags: xml:"lastboot,attr"
	LastBoot string `protobuf:"bytes,2,opt,name=LastBoot,proto3" json:"LastBoot,omitempty" xml:"lastboot,attr"`
}

func (x *Uptime) Reset() {
	*x = Uptime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_host_host_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Uptime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Uptime) ProtoMessage() {}

func (x *Uptime) ProtoReflect() protoreflect.Message {
	mi := &file_host_host_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Uptime.ProtoReflect.Descriptor instead.
func (*Uptime) Descriptor() ([]byte, []int) {
	return file_host_host_proto_rawDescGZIP(), []int{2}
}

func (x *Uptime) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Uptime) GetSeconds() int32 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

func (x *Uptime) GetLastBoot() string {
	if x != nil {
		return x.LastBoot
	}
	return ""
}

// Status represents a host's status.
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// @gotags: xml:"state,attr"
	State string `protobuf:"bytes,2,opt,name=State,proto3" json:"State,omitempty" xml:"state,attr"`
	// @gotags: xml:"reason,attr"
	Reason string `protobuf:"bytes,3,opt,name=Reason,proto3" json:"Reason,omitempty" xml:"reason,attr"`
	// @gotags: xml:"reason_ttl,attr"
	ReasonTTL string `protobuf:"bytes,4,opt,name=ReasonTTL,proto3" json:"ReasonTTL,omitempty" xml:"reason_ttl,attr"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_host_host_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_host_host_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_host_host_proto_rawDescGZIP(), []int{3}
}

func (x *Status) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Status) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Status) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *Status) GetReasonTTL() string {
	if x != nil {
		return x.ReasonTTL
	}
	return ""
}

var File_host_host_proto protoreflect.FileDescriptor

var file_host_host_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x68, 0x6f, 0x73, 0x74, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x68, 0x6f,
	0x73, 0x74, 0x2f, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x68, 0x6f, 0x73,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0f, 0x68, 0x6f, 0x73, 0x74, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0f, 0x68, 0x6f, 0x73, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x15, 0x68, 0x6f, 0x73, 0x74, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x63, 0x61, 0x6e, 0x2f, 0x6e, 0x6d, 0x61, 0x70,
	0x2f, 0x6e, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x0b, 0x0a, 0x04,
	0x48, 0x6f, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0e, 0xba, 0xb9, 0x19, 0x0a, 0x0a, 0x08, 0x12, 0x04, 0x75, 0x75, 0x69, 0x64, 0x28, 0x01,
	0x52, 0x02, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38,
	0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x41, 0x43, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x41, 0x43, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f,
	0x6d, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x6d, 0x6d, 0x12, 0x16,
	0x0a, 0x06, 0x4f, 0x53, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x4f, 0x53, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x53, 0x46, 0x6c, 0x61, 0x76,
	0x6f, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4f, 0x53, 0x46, 0x6c, 0x61, 0x76,
	0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x4f, 0x53, 0x53, 0x70, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4f, 0x53, 0x53, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x53, 0x4c, 0x61, 0x6e, 0x67,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4f, 0x53, 0x4c, 0x61, 0x6e, 0x67, 0x12, 0x1a,
	0x0a, 0x08, 0x4f, 0x53, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x4f, 0x53, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x72,
	0x63, 0x68, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x41, 0x72, 0x63, 0x68, 0x12, 0x18,
	0x0a, 0x07, 0x50, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x50, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05,
	0x53, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x63, 0x6f,
	0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x48, 0x6f, 0x73,
	0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c,
	0x48, 0x6f, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x18, 0x16, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x2b, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x18, 0x17, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x68, 0x6f, 0x73, 0x74,
	0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x02, 0x46, 0x53, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x42, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x22, 0x00, 0x52, 0x02, 0x46, 0x53, 0x12, 0x36, 0x0a,
	0x09, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x19, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x42, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x32, 0x00, 0x52, 0x09, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x02, 0x4f, 0x53, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x4f, 0x53, 0x42, 0x06, 0xba, 0xb9, 0x19,
	0x02, 0x22, 0x00, 0x52, 0x02, 0x4f, 0x53, 0x12, 0x24, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x35, 0x0a,
	0x08, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x20, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x42, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x22, 0x00, 0x52, 0x08, 0x44, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x21, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x22, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0c,
	0x49, 0x50, 0x49, 0x44, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x23, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x49, 0x50, 0x49,
	0x44, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x22,
	0x00, 0x52, 0x0c, 0x49, 0x50, 0x49, 0x44, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x3e, 0x0a, 0x0b, 0x54, 0x43, 0x50, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x24,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x54,
	0x43, 0x50, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x06, 0xba, 0xb9, 0x19, 0x02,
	0x22, 0x00, 0x52, 0x0b, 0x54, 0x43, 0x50, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x44, 0x0a, 0x0d, 0x54, 0x43, 0x50, 0x54, 0x53, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x18, 0x25, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2e, 0x54, 0x43, 0x50, 0x54, 0x53, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x06,
	0xba, 0xb9, 0x19, 0x02, 0x22, 0x00, 0x52, 0x0d, 0x54, 0x43, 0x50, 0x54, 0x53, 0x53, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x49, 0x43, 0x4d, 0x50, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x26, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x49, 0x43, 0x4d, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x22, 0x00, 0x52, 0x0c, 0x49, 0x43, 0x4d, 0x50,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x54, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x27, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x42, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x22,
	0x00, 0x52, 0x07, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x2c, 0x0a, 0x05, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x42, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x22,
	0x00, 0x52, 0x05, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x55, 0x70, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x29, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e,
	0x55, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x22, 0x00, 0x52, 0x06,
	0x55, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x3b, 0x0a, 0x0b, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x18,
	0x2b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x6e, 0x6d, 0x61,
	0x70, 0x2e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x42, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x32, 0x00,
	0x52, 0x0b, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x12, 0x30, 0x0a,
	0x06, 0x53, 0x6d, 0x75, 0x72, 0x66, 0x73, 0x18, 0x2c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x73, 0x63, 0x61, 0x6e, 0x2e, 0x6e, 0x6d, 0x61, 0x70, 0x2e, 0x53, 0x6d, 0x75, 0x72, 0x66, 0x42,
	0x06, 0xba, 0xb9, 0x19, 0x02, 0x32, 0x00, 0x52, 0x06, 0x53, 0x6d, 0x75, 0x72, 0x66, 0x73, 0x12,
	0x2c, 0x0a, 0x09, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x2d, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x52, 0x09, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x28, 0x0a,
	0x05, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x2e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x68,
	0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x42, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x22, 0x00,
	0x52, 0x05, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x37, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x50, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x2f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x68, 0x6f,
	0x73, 0x74, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x50, 0x6f, 0x72, 0x74, 0x42, 0x06, 0xba, 0xb9,
	0x19, 0x02, 0x22, 0x00, 0x52, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61, 0x50, 0x6f, 0x72, 0x74, 0x73,
	0x3a, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x22, 0xce, 0x01, 0x0a, 0x08, 0x48, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0e, 0xba, 0xb9, 0x19, 0x0a, 0x0a, 0x08, 0x12, 0x04, 0x75, 0x75, 0x69, 0x64, 0x28,
	0x01, 0x52, 0x02, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x38, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x3a, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x22, 0x66, 0x0a, 0x06, 0x55, 0x70, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0e, 0xba, 0xb9, 0x19, 0x0a, 0x0a, 0x08, 0x12, 0x04, 0x75, 0x75, 0x69, 0x64, 0x28, 0x01, 0x52,
	0x02, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x4c, 0x61, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x4c, 0x61, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x74, 0x3a, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x08,
	0x01, 0x22, 0x7c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xba, 0xb9, 0x19, 0x0a, 0x0a, 0x08, 0x12,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x28, 0x01, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x54, 0x54, 0x4c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x54, 0x54, 0x4c, 0x3a, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x42,
	0x6b, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x42, 0x09, 0x48, 0x6f, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x78, 0x6c, 0x61, 0x6e, 0x64, 0x6f, 0x6e, 0x2f, 0x61,
	0x69, 0x6d, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0xa2, 0x02,
	0x03, 0x48, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x48, 0x6f, 0x73, 0x74, 0xca, 0x02, 0x04, 0x48, 0x6f,
	0x73, 0x74, 0xe2, 0x02, 0x10, 0x48, 0x6f, 0x73, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_host_host_proto_rawDescOnce sync.Once
	file_host_host_proto_rawDescData = file_host_host_proto_rawDesc
)

func file_host_host_proto_rawDescGZIP() []byte {
	file_host_host_proto_rawDescOnce.Do(func() {
		file_host_host_proto_rawDescData = protoimpl.X.CompressGZIP(file_host_host_proto_rawDescData)
	})
	return file_host_host_proto_rawDescData
}

var file_host_host_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_host_host_proto_goTypes = []interface{}{
	(*Host)(nil),                  // 0: host.Host
	(*Hostname)(nil),              // 1: host.Hostname
	(*Uptime)(nil),                // 2: host.Uptime
	(*Status)(nil),                // 3: host.Status
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*User)(nil),                  // 5: host.User
	(*Process)(nil),               // 6: host.Process
	(*FileSystem)(nil),            // 7: host.FileSystem
	(*network.Address)(nil),       // 8: network.Address
	(*OS)(nil),                    // 9: host.OS
	(*network.Distance)(nil),      // 10: network.Distance
	(*network.IPIDSequence)(nil),  // 11: network.IPIDSequence
	(*network.TCPSequence)(nil),   // 12: network.TCPSequence
	(*network.TCPTSSequence)(nil), // 13: network.TCPTSSequence
	(*network.ICMPResponse)(nil),  // 14: network.ICMPResponse
	(*network.Times)(nil),         // 15: network.Times
	(*network.Trace)(nil),         // 16: network.Trace
	(*nmap.Script)(nil),           // 17: scan.nmap.Script
	(*nmap.Smurf)(nil),            // 18: scan.nmap.Smurf
	(*Port)(nil),                  // 19: host.Port
	(*ExtraPort)(nil),             // 20: host.ExtraPort
}
var file_host_host_proto_depIdxs = []int32{
	4,  // 0: host.Host.CreatedAt:type_name -> google.protobuf.Timestamp
	4,  // 1: host.Host.UpdatedAt:type_name -> google.protobuf.Timestamp
	5,  // 2: host.Host.Users:type_name -> host.User
	6,  // 3: host.Host.Processes:type_name -> host.Process
	7,  // 4: host.Host.FS:type_name -> host.FileSystem
	8,  // 5: host.Host.Addresses:type_name -> network.Address
	9,  // 6: host.Host.OS:type_name -> host.OS
	3,  // 7: host.Host.Status:type_name -> host.Status
	10, // 8: host.Host.Distance:type_name -> network.Distance
	11, // 9: host.Host.IPIDSequence:type_name -> network.IPIDSequence
	12, // 10: host.Host.TCPSequence:type_name -> network.TCPSequence
	13, // 11: host.Host.TCPTSSequence:type_name -> network.TCPTSSequence
	14, // 12: host.Host.ICMPResponse:type_name -> network.ICMPResponse
	15, // 13: host.Host.Testing:type_name -> network.Times
	16, // 14: host.Host.Trace:type_name -> network.Trace
	2,  // 15: host.Host.Uptime:type_name -> host.Uptime
	17, // 16: host.Host.HostScripts:type_name -> scan.nmap.Script
	18, // 17: host.Host.Smurfs:type_name -> scan.nmap.Smurf
	1,  // 18: host.Host.Hostnames:type_name -> host.Hostname
	19, // 19: host.Host.Ports:type_name -> host.Port
	20, // 20: host.Host.ExtraPorts:type_name -> host.ExtraPort
	4,  // 21: host.Hostname.CreatedAt:type_name -> google.protobuf.Timestamp
	4,  // 22: host.Hostname.UpdatedAt:type_name -> google.protobuf.Timestamp
	23, // [23:23] is the sub-list for method output_type
	23, // [23:23] is the sub-list for method input_type
	23, // [23:23] is the sub-list for extension type_name
	23, // [23:23] is the sub-list for extension extendee
	0,  // [0:23] is the sub-list for field type_name
}

func init() { file_host_host_proto_init() }
func file_host_host_proto_init() {
	if File_host_host_proto != nil {
		return
	}
	file_host_os_proto_init()
	file_host_process_proto_init()
	file_host_port_proto_init()
	file_host_user_proto_init()
	file_host_filesystem_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_host_host_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Host); i {
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
		file_host_host_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hostname); i {
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
		file_host_host_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Uptime); i {
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
		file_host_host_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_host_host_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_host_host_proto_goTypes,
		DependencyIndexes: file_host_host_proto_depIdxs,
		MessageInfos:      file_host_host_proto_msgTypes,
	}.Build()
	File_host_host_proto = out.File
	file_host_host_proto_rawDesc = nil
	file_host_host_proto_goTypes = nil
	file_host_host_proto_depIdxs = nil
}
