// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lxd/migration/migrate.proto

package migration

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MigrationFSType int32

const (
	MigrationFSType_RSYNC           MigrationFSType = 0
	MigrationFSType_BTRFS           MigrationFSType = 1
	MigrationFSType_ZFS             MigrationFSType = 2
	MigrationFSType_RBD             MigrationFSType = 3
	MigrationFSType_BLOCK_AND_RSYNC MigrationFSType = 4
)

var MigrationFSType_name = map[int32]string{
	0: "RSYNC",
	1: "BTRFS",
	2: "ZFS",
	3: "RBD",
	4: "BLOCK_AND_RSYNC",
}

var MigrationFSType_value = map[string]int32{
	"RSYNC":           0,
	"BTRFS":           1,
	"ZFS":             2,
	"RBD":             3,
	"BLOCK_AND_RSYNC": 4,
}

func (x MigrationFSType) Enum() *MigrationFSType {
	p := new(MigrationFSType)
	*p = x
	return p
}

func (x MigrationFSType) String() string {
	return proto.EnumName(MigrationFSType_name, int32(x))
}

func (x *MigrationFSType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MigrationFSType_value, data, "MigrationFSType")
	if err != nil {
		return err
	}
	*x = MigrationFSType(value)
	return nil
}

func (MigrationFSType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fe8772548dc4b615, []int{0}
}

type CRIUType int32

const (
	CRIUType_CRIU_RSYNC CRIUType = 0
	CRIUType_PHAUL      CRIUType = 1
	CRIUType_NONE       CRIUType = 2
)

var CRIUType_name = map[int32]string{
	0: "CRIU_RSYNC",
	1: "PHAUL",
	2: "NONE",
}

var CRIUType_value = map[string]int32{
	"CRIU_RSYNC": 0,
	"PHAUL":      1,
	"NONE":       2,
}

func (x CRIUType) Enum() *CRIUType {
	p := new(CRIUType)
	*p = x
	return p
}

func (x CRIUType) String() string {
	return proto.EnumName(CRIUType_name, int32(x))
}

func (x *CRIUType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CRIUType_value, data, "CRIUType")
	if err != nil {
		return err
	}
	*x = CRIUType(value)
	return nil
}

func (CRIUType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fe8772548dc4b615, []int{1}
}

type IDMapType struct {
	Isuid                *bool    `protobuf:"varint,1,req,name=isuid" json:"isuid,omitempty"`
	Isgid                *bool    `protobuf:"varint,2,req,name=isgid" json:"isgid,omitempty"`
	Hostid               *int32   `protobuf:"varint,3,req,name=hostid" json:"hostid,omitempty"`
	Nsid                 *int32   `protobuf:"varint,4,req,name=nsid" json:"nsid,omitempty"`
	Maprange             *int32   `protobuf:"varint,5,req,name=maprange" json:"maprange,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IDMapType) Reset()         { *m = IDMapType{} }
func (m *IDMapType) String() string { return proto.CompactTextString(m) }
func (*IDMapType) ProtoMessage()    {}
func (*IDMapType) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe8772548dc4b615, []int{0}
}

func (m *IDMapType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IDMapType.Unmarshal(m, b)
}
func (m *IDMapType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IDMapType.Marshal(b, m, deterministic)
}
func (m *IDMapType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDMapType.Merge(m, src)
}
func (m *IDMapType) XXX_Size() int {
	return xxx_messageInfo_IDMapType.Size(m)
}
func (m *IDMapType) XXX_DiscardUnknown() {
	xxx_messageInfo_IDMapType.DiscardUnknown(m)
}

var xxx_messageInfo_IDMapType proto.InternalMessageInfo

func (m *IDMapType) GetIsuid() bool {
	if m != nil && m.Isuid != nil {
		return *m.Isuid
	}
	return false
}

func (m *IDMapType) GetIsgid() bool {
	if m != nil && m.Isgid != nil {
		return *m.Isgid
	}
	return false
}

func (m *IDMapType) GetHostid() int32 {
	if m != nil && m.Hostid != nil {
		return *m.Hostid
	}
	return 0
}

func (m *IDMapType) GetNsid() int32 {
	if m != nil && m.Nsid != nil {
		return *m.Nsid
	}
	return 0
}

func (m *IDMapType) GetMaprange() int32 {
	if m != nil && m.Maprange != nil {
		return *m.Maprange
	}
	return 0
}

type Config struct {
	Key                  *string  `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value                *string  `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe8772548dc4b615, []int{1}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Config) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type Device struct {
	Name                 *string   `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Config               []*Config `protobuf:"bytes,2,rep,name=config" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe8772548dc4b615, []int{2}
}

func (m *Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Device.Unmarshal(m, b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Device.Marshal(b, m, deterministic)
}
func (m *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(m, src)
}
func (m *Device) XXX_Size() int {
	return xxx_messageInfo_Device.Size(m)
}
func (m *Device) XXX_DiscardUnknown() {
	xxx_messageInfo_Device.DiscardUnknown(m)
}

var xxx_messageInfo_Device proto.InternalMessageInfo

func (m *Device) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Device) GetConfig() []*Config {
	if m != nil {
		return m.Config
	}
	return nil
}

type Snapshot struct {
	Name                 *string   `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	LocalConfig          []*Config `protobuf:"bytes,2,rep,name=localConfig" json:"localConfig,omitempty"`
	Profiles             []string  `protobuf:"bytes,3,rep,name=profiles" json:"profiles,omitempty"`
	Ephemeral            *bool     `protobuf:"varint,4,req,name=ephemeral" json:"ephemeral,omitempty"`
	LocalDevices         []*Device `protobuf:"bytes,5,rep,name=localDevices" json:"localDevices,omitempty"`
	Architecture         *int32    `protobuf:"varint,6,req,name=architecture" json:"architecture,omitempty"`
	Stateful             *bool     `protobuf:"varint,7,req,name=stateful" json:"stateful,omitempty"`
	CreationDate         *int64    `protobuf:"varint,8,opt,name=creation_date,json=creationDate" json:"creation_date,omitempty"`
	LastUsedDate         *int64    `protobuf:"varint,9,opt,name=last_used_date,json=lastUsedDate" json:"last_used_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Snapshot) Reset()         { *m = Snapshot{} }
func (m *Snapshot) String() string { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()    {}
func (*Snapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe8772548dc4b615, []int{3}
}

func (m *Snapshot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Snapshot.Unmarshal(m, b)
}
func (m *Snapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Snapshot.Marshal(b, m, deterministic)
}
func (m *Snapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Snapshot.Merge(m, src)
}
func (m *Snapshot) XXX_Size() int {
	return xxx_messageInfo_Snapshot.Size(m)
}
func (m *Snapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_Snapshot.DiscardUnknown(m)
}

var xxx_messageInfo_Snapshot proto.InternalMessageInfo

func (m *Snapshot) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Snapshot) GetLocalConfig() []*Config {
	if m != nil {
		return m.LocalConfig
	}
	return nil
}

func (m *Snapshot) GetProfiles() []string {
	if m != nil {
		return m.Profiles
	}
	return nil
}

func (m *Snapshot) GetEphemeral() bool {
	if m != nil && m.Ephemeral != nil {
		return *m.Ephemeral
	}
	return false
}

func (m *Snapshot) GetLocalDevices() []*Device {
	if m != nil {
		return m.LocalDevices
	}
	return nil
}

func (m *Snapshot) GetArchitecture() int32 {
	if m != nil && m.Architecture != nil {
		return *m.Architecture
	}
	return 0
}

func (m *Snapshot) GetStateful() bool {
	if m != nil && m.Stateful != nil {
		return *m.Stateful
	}
	return false
}

func (m *Snapshot) GetCreationDate() int64 {
	if m != nil && m.CreationDate != nil {
		return *m.CreationDate
	}
	return 0
}

func (m *Snapshot) GetLastUsedDate() int64 {
	if m != nil && m.LastUsedDate != nil {
		return *m.LastUsedDate
	}
	return 0
}

type RsyncFeatures struct {
	Xattrs               *bool    `protobuf:"varint,1,opt,name=xattrs" json:"xattrs,omitempty"`
	Delete               *bool    `protobuf:"varint,2,opt,name=delete" json:"delete,omitempty"`
	Compress             *bool    `protobuf:"varint,3,opt,name=compress" json:"compress,omitempty"`
	Bidirectional        *bool    `protobuf:"varint,4,opt,name=bidirectional" json:"bidirectional,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsyncFeatures) Reset()         { *m = RsyncFeatures{} }
func (m *RsyncFeatures) String() string { return proto.CompactTextString(m) }
func (*RsyncFeatures) ProtoMessage()    {}
func (*RsyncFeatures) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe8772548dc4b615, []int{4}
}

func (m *RsyncFeatures) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsyncFeatures.Unmarshal(m, b)
}
func (m *RsyncFeatures) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsyncFeatures.Marshal(b, m, deterministic)
}
func (m *RsyncFeatures) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsyncFeatures.Merge(m, src)
}
func (m *RsyncFeatures) XXX_Size() int {
	return xxx_messageInfo_RsyncFeatures.Size(m)
}
func (m *RsyncFeatures) XXX_DiscardUnknown() {
	xxx_messageInfo_RsyncFeatures.DiscardUnknown(m)
}

var xxx_messageInfo_RsyncFeatures proto.InternalMessageInfo

func (m *RsyncFeatures) GetXattrs() bool {
	if m != nil && m.Xattrs != nil {
		return *m.Xattrs
	}
	return false
}

func (m *RsyncFeatures) GetDelete() bool {
	if m != nil && m.Delete != nil {
		return *m.Delete
	}
	return false
}

func (m *RsyncFeatures) GetCompress() bool {
	if m != nil && m.Compress != nil {
		return *m.Compress
	}
	return false
}

func (m *RsyncFeatures) GetBidirectional() bool {
	if m != nil && m.Bidirectional != nil {
		return *m.Bidirectional
	}
	return false
}

type ZfsFeatures struct {
	Compress             *bool    `protobuf:"varint,1,opt,name=compress" json:"compress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ZfsFeatures) Reset()         { *m = ZfsFeatures{} }
func (m *ZfsFeatures) String() string { return proto.CompactTextString(m) }
func (*ZfsFeatures) ProtoMessage()    {}
func (*ZfsFeatures) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe8772548dc4b615, []int{5}
}

func (m *ZfsFeatures) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZfsFeatures.Unmarshal(m, b)
}
func (m *ZfsFeatures) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZfsFeatures.Marshal(b, m, deterministic)
}
func (m *ZfsFeatures) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZfsFeatures.Merge(m, src)
}
func (m *ZfsFeatures) XXX_Size() int {
	return xxx_messageInfo_ZfsFeatures.Size(m)
}
func (m *ZfsFeatures) XXX_DiscardUnknown() {
	xxx_messageInfo_ZfsFeatures.DiscardUnknown(m)
}

var xxx_messageInfo_ZfsFeatures proto.InternalMessageInfo

func (m *ZfsFeatures) GetCompress() bool {
	if m != nil && m.Compress != nil {
		return *m.Compress
	}
	return false
}

type MigrationHeader struct {
	Fs                   *MigrationFSType `protobuf:"varint,1,req,name=fs,enum=migration.MigrationFSType" json:"fs,omitempty"`
	Criu                 *CRIUType        `protobuf:"varint,2,opt,name=criu,enum=migration.CRIUType" json:"criu,omitempty"`
	Idmap                []*IDMapType     `protobuf:"bytes,3,rep,name=idmap" json:"idmap,omitempty"`
	SnapshotNames        []string         `protobuf:"bytes,4,rep,name=snapshotNames" json:"snapshotNames,omitempty"`
	Snapshots            []*Snapshot      `protobuf:"bytes,5,rep,name=snapshots" json:"snapshots,omitempty"`
	Predump              *bool            `protobuf:"varint,7,opt,name=predump" json:"predump,omitempty"`
	RsyncFeatures        *RsyncFeatures   `protobuf:"bytes,8,opt,name=rsyncFeatures" json:"rsyncFeatures,omitempty"`
	Refresh              *bool            `protobuf:"varint,9,opt,name=refresh" json:"refresh,omitempty"`
	ZfsFeatures          *ZfsFeatures     `protobuf:"bytes,10,opt,name=zfsFeatures" json:"zfsFeatures,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *MigrationHeader) Reset()         { *m = MigrationHeader{} }
func (m *MigrationHeader) String() string { return proto.CompactTextString(m) }
func (*MigrationHeader) ProtoMessage()    {}
func (*MigrationHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe8772548dc4b615, []int{6}
}

func (m *MigrationHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MigrationHeader.Unmarshal(m, b)
}
func (m *MigrationHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MigrationHeader.Marshal(b, m, deterministic)
}
func (m *MigrationHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MigrationHeader.Merge(m, src)
}
func (m *MigrationHeader) XXX_Size() int {
	return xxx_messageInfo_MigrationHeader.Size(m)
}
func (m *MigrationHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_MigrationHeader.DiscardUnknown(m)
}

var xxx_messageInfo_MigrationHeader proto.InternalMessageInfo

func (m *MigrationHeader) GetFs() MigrationFSType {
	if m != nil && m.Fs != nil {
		return *m.Fs
	}
	return MigrationFSType_RSYNC
}

func (m *MigrationHeader) GetCriu() CRIUType {
	if m != nil && m.Criu != nil {
		return *m.Criu
	}
	return CRIUType_CRIU_RSYNC
}

func (m *MigrationHeader) GetIdmap() []*IDMapType {
	if m != nil {
		return m.Idmap
	}
	return nil
}

func (m *MigrationHeader) GetSnapshotNames() []string {
	if m != nil {
		return m.SnapshotNames
	}
	return nil
}

func (m *MigrationHeader) GetSnapshots() []*Snapshot {
	if m != nil {
		return m.Snapshots
	}
	return nil
}

func (m *MigrationHeader) GetPredump() bool {
	if m != nil && m.Predump != nil {
		return *m.Predump
	}
	return false
}

func (m *MigrationHeader) GetRsyncFeatures() *RsyncFeatures {
	if m != nil {
		return m.RsyncFeatures
	}
	return nil
}

func (m *MigrationHeader) GetRefresh() bool {
	if m != nil && m.Refresh != nil {
		return *m.Refresh
	}
	return false
}

func (m *MigrationHeader) GetZfsFeatures() *ZfsFeatures {
	if m != nil {
		return m.ZfsFeatures
	}
	return nil
}

type MigrationControl struct {
	Success *bool `protobuf:"varint,1,req,name=success" json:"success,omitempty"`
	// optional failure message if sending a failure
	Message              *string  `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MigrationControl) Reset()         { *m = MigrationControl{} }
func (m *MigrationControl) String() string { return proto.CompactTextString(m) }
func (*MigrationControl) ProtoMessage()    {}
func (*MigrationControl) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe8772548dc4b615, []int{7}
}

func (m *MigrationControl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MigrationControl.Unmarshal(m, b)
}
func (m *MigrationControl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MigrationControl.Marshal(b, m, deterministic)
}
func (m *MigrationControl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MigrationControl.Merge(m, src)
}
func (m *MigrationControl) XXX_Size() int {
	return xxx_messageInfo_MigrationControl.Size(m)
}
func (m *MigrationControl) XXX_DiscardUnknown() {
	xxx_messageInfo_MigrationControl.DiscardUnknown(m)
}

var xxx_messageInfo_MigrationControl proto.InternalMessageInfo

func (m *MigrationControl) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

func (m *MigrationControl) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type MigrationSync struct {
	FinalPreDump         *bool    `protobuf:"varint,1,req,name=finalPreDump" json:"finalPreDump,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MigrationSync) Reset()         { *m = MigrationSync{} }
func (m *MigrationSync) String() string { return proto.CompactTextString(m) }
func (*MigrationSync) ProtoMessage()    {}
func (*MigrationSync) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe8772548dc4b615, []int{8}
}

func (m *MigrationSync) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MigrationSync.Unmarshal(m, b)
}
func (m *MigrationSync) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MigrationSync.Marshal(b, m, deterministic)
}
func (m *MigrationSync) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MigrationSync.Merge(m, src)
}
func (m *MigrationSync) XXX_Size() int {
	return xxx_messageInfo_MigrationSync.Size(m)
}
func (m *MigrationSync) XXX_DiscardUnknown() {
	xxx_messageInfo_MigrationSync.DiscardUnknown(m)
}

var xxx_messageInfo_MigrationSync proto.InternalMessageInfo

func (m *MigrationSync) GetFinalPreDump() bool {
	if m != nil && m.FinalPreDump != nil {
		return *m.FinalPreDump
	}
	return false
}

// This one contains statistics about dump/restore process
type DumpStatsEntry struct {
	FreezingTime         *uint32  `protobuf:"varint,1,req,name=freezing_time,json=freezingTime" json:"freezing_time,omitempty"`
	FrozenTime           *uint32  `protobuf:"varint,2,req,name=frozen_time,json=frozenTime" json:"frozen_time,omitempty"`
	MemdumpTime          *uint32  `protobuf:"varint,3,req,name=memdump_time,json=memdumpTime" json:"memdump_time,omitempty"`
	MemwriteTime         *uint32  `protobuf:"varint,4,req,name=memwrite_time,json=memwriteTime" json:"memwrite_time,omitempty"`
	PagesScanned         *uint64  `protobuf:"varint,5,req,name=pages_scanned,json=pagesScanned" json:"pages_scanned,omitempty"`
	PagesSkippedParent   *uint64  `protobuf:"varint,6,req,name=pages_skipped_parent,json=pagesSkippedParent" json:"pages_skipped_parent,omitempty"`
	PagesWritten         *uint64  `protobuf:"varint,7,req,name=pages_written,json=pagesWritten" json:"pages_written,omitempty"`
	IrmapResolve         *uint32  `protobuf:"varint,8,opt,name=irmap_resolve,json=irmapResolve" json:"irmap_resolve,omitempty"`
	PagesLazy            *uint64  `protobuf:"varint,9,req,name=pages_lazy,json=pagesLazy" json:"pages_lazy,omitempty"`
	PagePipes            *uint64  `protobuf:"varint,10,opt,name=page_pipes,json=pagePipes" json:"page_pipes,omitempty"`
	PagePipeBufs         *uint64  `protobuf:"varint,11,opt,name=page_pipe_bufs,json=pagePipeBufs" json:"page_pipe_bufs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DumpStatsEntry) Reset()         { *m = DumpStatsEntry{} }
func (m *DumpStatsEntry) String() string { return proto.CompactTextString(m) }
func (*DumpStatsEntry) ProtoMessage()    {}
func (*DumpStatsEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe8772548dc4b615, []int{9}
}

func (m *DumpStatsEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DumpStatsEntry.Unmarshal(m, b)
}
func (m *DumpStatsEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DumpStatsEntry.Marshal(b, m, deterministic)
}
func (m *DumpStatsEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DumpStatsEntry.Merge(m, src)
}
func (m *DumpStatsEntry) XXX_Size() int {
	return xxx_messageInfo_DumpStatsEntry.Size(m)
}
func (m *DumpStatsEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_DumpStatsEntry.DiscardUnknown(m)
}

var xxx_messageInfo_DumpStatsEntry proto.InternalMessageInfo

func (m *DumpStatsEntry) GetFreezingTime() uint32 {
	if m != nil && m.FreezingTime != nil {
		return *m.FreezingTime
	}
	return 0
}

func (m *DumpStatsEntry) GetFrozenTime() uint32 {
	if m != nil && m.FrozenTime != nil {
		return *m.FrozenTime
	}
	return 0
}

func (m *DumpStatsEntry) GetMemdumpTime() uint32 {
	if m != nil && m.MemdumpTime != nil {
		return *m.MemdumpTime
	}
	return 0
}

func (m *DumpStatsEntry) GetMemwriteTime() uint32 {
	if m != nil && m.MemwriteTime != nil {
		return *m.MemwriteTime
	}
	return 0
}

func (m *DumpStatsEntry) GetPagesScanned() uint64 {
	if m != nil && m.PagesScanned != nil {
		return *m.PagesScanned
	}
	return 0
}

func (m *DumpStatsEntry) GetPagesSkippedParent() uint64 {
	if m != nil && m.PagesSkippedParent != nil {
		return *m.PagesSkippedParent
	}
	return 0
}

func (m *DumpStatsEntry) GetPagesWritten() uint64 {
	if m != nil && m.PagesWritten != nil {
		return *m.PagesWritten
	}
	return 0
}

func (m *DumpStatsEntry) GetIrmapResolve() uint32 {
	if m != nil && m.IrmapResolve != nil {
		return *m.IrmapResolve
	}
	return 0
}

func (m *DumpStatsEntry) GetPagesLazy() uint64 {
	if m != nil && m.PagesLazy != nil {
		return *m.PagesLazy
	}
	return 0
}

func (m *DumpStatsEntry) GetPagePipes() uint64 {
	if m != nil && m.PagePipes != nil {
		return *m.PagePipes
	}
	return 0
}

func (m *DumpStatsEntry) GetPagePipeBufs() uint64 {
	if m != nil && m.PagePipeBufs != nil {
		return *m.PagePipeBufs
	}
	return 0
}

type RestoreStatsEntry struct {
	PagesCompared        *uint64  `protobuf:"varint,1,req,name=pages_compared,json=pagesCompared" json:"pages_compared,omitempty"`
	PagesSkippedCow      *uint64  `protobuf:"varint,2,req,name=pages_skipped_cow,json=pagesSkippedCow" json:"pages_skipped_cow,omitempty"`
	ForkingTime          *uint32  `protobuf:"varint,3,req,name=forking_time,json=forkingTime" json:"forking_time,omitempty"`
	RestoreTime          *uint32  `protobuf:"varint,4,req,name=restore_time,json=restoreTime" json:"restore_time,omitempty"`
	PagesRestored        *uint64  `protobuf:"varint,5,opt,name=pages_restored,json=pagesRestored" json:"pages_restored,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RestoreStatsEntry) Reset()         { *m = RestoreStatsEntry{} }
func (m *RestoreStatsEntry) String() string { return proto.CompactTextString(m) }
func (*RestoreStatsEntry) ProtoMessage()    {}
func (*RestoreStatsEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe8772548dc4b615, []int{10}
}

func (m *RestoreStatsEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestoreStatsEntry.Unmarshal(m, b)
}
func (m *RestoreStatsEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestoreStatsEntry.Marshal(b, m, deterministic)
}
func (m *RestoreStatsEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestoreStatsEntry.Merge(m, src)
}
func (m *RestoreStatsEntry) XXX_Size() int {
	return xxx_messageInfo_RestoreStatsEntry.Size(m)
}
func (m *RestoreStatsEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_RestoreStatsEntry.DiscardUnknown(m)
}

var xxx_messageInfo_RestoreStatsEntry proto.InternalMessageInfo

func (m *RestoreStatsEntry) GetPagesCompared() uint64 {
	if m != nil && m.PagesCompared != nil {
		return *m.PagesCompared
	}
	return 0
}

func (m *RestoreStatsEntry) GetPagesSkippedCow() uint64 {
	if m != nil && m.PagesSkippedCow != nil {
		return *m.PagesSkippedCow
	}
	return 0
}

func (m *RestoreStatsEntry) GetForkingTime() uint32 {
	if m != nil && m.ForkingTime != nil {
		return *m.ForkingTime
	}
	return 0
}

func (m *RestoreStatsEntry) GetRestoreTime() uint32 {
	if m != nil && m.RestoreTime != nil {
		return *m.RestoreTime
	}
	return 0
}

func (m *RestoreStatsEntry) GetPagesRestored() uint64 {
	if m != nil && m.PagesRestored != nil {
		return *m.PagesRestored
	}
	return 0
}

type StatsEntry struct {
	Dump                 *DumpStatsEntry    `protobuf:"bytes,1,opt,name=dump" json:"dump,omitempty"`
	Restore              *RestoreStatsEntry `protobuf:"bytes,2,opt,name=restore" json:"restore,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *StatsEntry) Reset()         { *m = StatsEntry{} }
func (m *StatsEntry) String() string { return proto.CompactTextString(m) }
func (*StatsEntry) ProtoMessage()    {}
func (*StatsEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe8772548dc4b615, []int{11}
}

func (m *StatsEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsEntry.Unmarshal(m, b)
}
func (m *StatsEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsEntry.Marshal(b, m, deterministic)
}
func (m *StatsEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsEntry.Merge(m, src)
}
func (m *StatsEntry) XXX_Size() int {
	return xxx_messageInfo_StatsEntry.Size(m)
}
func (m *StatsEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsEntry.DiscardUnknown(m)
}

var xxx_messageInfo_StatsEntry proto.InternalMessageInfo

func (m *StatsEntry) GetDump() *DumpStatsEntry {
	if m != nil {
		return m.Dump
	}
	return nil
}

func (m *StatsEntry) GetRestore() *RestoreStatsEntry {
	if m != nil {
		return m.Restore
	}
	return nil
}

func init() {
	proto.RegisterEnum("migration.MigrationFSType", MigrationFSType_name, MigrationFSType_value)
	proto.RegisterEnum("migration.CRIUType", CRIUType_name, CRIUType_value)
	proto.RegisterType((*IDMapType)(nil), "migration.IDMapType")
	proto.RegisterType((*Config)(nil), "migration.Config")
	proto.RegisterType((*Device)(nil), "migration.Device")
	proto.RegisterType((*Snapshot)(nil), "migration.Snapshot")
	proto.RegisterType((*RsyncFeatures)(nil), "migration.rsyncFeatures")
	proto.RegisterType((*ZfsFeatures)(nil), "migration.zfsFeatures")
	proto.RegisterType((*MigrationHeader)(nil), "migration.MigrationHeader")
	proto.RegisterType((*MigrationControl)(nil), "migration.MigrationControl")
	proto.RegisterType((*MigrationSync)(nil), "migration.MigrationSync")
	proto.RegisterType((*DumpStatsEntry)(nil), "migration.dump_stats_entry")
	proto.RegisterType((*RestoreStatsEntry)(nil), "migration.restore_stats_entry")
	proto.RegisterType((*StatsEntry)(nil), "migration.stats_entry")
}

func init() { proto.RegisterFile("lxd/migration/migrate.proto", fileDescriptor_fe8772548dc4b615) }

var fileDescriptor_fe8772548dc4b615 = []byte{
	// 1063 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0x9e, 0x6d, 0x39, 0xb1, 0x8e, 0xec, 0xc4, 0x65, 0x82, 0x42, 0x68, 0xf7, 0xe3, 0xa9, 0x1d,
	0xe6, 0xe6, 0x22, 0xe9, 0x5c, 0x0c, 0xe8, 0xd5, 0x80, 0xc6, 0x5e, 0xd6, 0x62, 0xa9, 0x1b, 0xd0,
	0x09, 0x86, 0xed, 0x46, 0x60, 0xa5, 0x23, 0x87, 0x88, 0xfe, 0x40, 0xca, 0x49, 0x9c, 0x9b, 0x61,
	0x8f, 0xb1, 0x07, 0xd8, 0xf3, 0xec, 0x6a, 0xef, 0x33, 0x90, 0x94, 0x14, 0x39, 0x1d, 0xb0, 0x3b,
	0x9e, 0xef, 0x7c, 0xfa, 0x0e, 0x79, 0xfe, 0x04, 0x4f, 0xe3, 0xdb, 0xf0, 0x28, 0xe1, 0x4b, 0xc1,
	0x0a, 0x9e, 0xa5, 0xe5, 0x09, 0x0f, 0x73, 0x91, 0x15, 0x19, 0xb1, 0x6b, 0x87, 0xf7, 0x3b, 0xd8,
	0xef, 0x66, 0xef, 0x59, 0x7e, 0xbe, 0xce, 0x91, 0xec, 0x43, 0x97, 0xcb, 0x15, 0x0f, 0xdd, 0xd6,
	0xa8, 0x3d, 0xee, 0x51, 0x63, 0x18, 0x74, 0xc9, 0x43, 0xb7, 0x5d, 0xa1, 0x4b, 0x1e, 0x92, 0xc7,
	0xb0, 0x75, 0x99, 0xc9, 0x82, 0x87, 0x6e, 0x67, 0xd4, 0x1e, 0x77, 0x69, 0x69, 0x11, 0x02, 0x56,
	0x2a, 0x79, 0xe8, 0x5a, 0x1a, 0xd5, 0x67, 0xf2, 0x04, 0x7a, 0x09, 0xcb, 0x05, 0x4b, 0x97, 0xe8,
	0x76, 0x35, 0x5e, 0xdb, 0xde, 0x4b, 0xd8, 0x9a, 0x66, 0x69, 0xc4, 0x97, 0x64, 0x08, 0x9d, 0x2b,
	0x5c, 0xeb, 0xd8, 0x36, 0x55, 0x47, 0x15, 0xf9, 0x9a, 0xc5, 0x2b, 0xd4, 0x91, 0x6d, 0x6a, 0x0c,
	0xef, 0x27, 0xd8, 0x9a, 0xe1, 0x35, 0x0f, 0x50, 0xc7, 0x62, 0x09, 0x96, 0x9f, 0xe8, 0x33, 0x79,
	0x01, 0x5b, 0x81, 0xd6, 0x73, 0xdb, 0xa3, 0xce, 0xd8, 0x99, 0x3c, 0x3a, 0xac, 0x1f, 0x7b, 0x68,
	0x02, 0xd1, 0x92, 0xe0, 0xfd, 0xdd, 0x86, 0xde, 0x22, 0x65, 0xb9, 0xbc, 0xcc, 0x8a, 0xff, 0xd4,
	0x7a, 0x05, 0x4e, 0x9c, 0x05, 0x2c, 0x9e, 0xfe, 0x8f, 0x60, 0x93, 0xa5, 0x1e, 0x9b, 0x8b, 0x2c,
	0xe2, 0x31, 0x4a, 0xb7, 0x33, 0xea, 0x8c, 0x6d, 0x5a, 0xdb, 0xe4, 0x73, 0xb0, 0x31, 0xbf, 0xc4,
	0x04, 0x05, 0x8b, 0x75, 0x86, 0x7a, 0xf4, 0x1e, 0x20, 0xdf, 0x43, 0x5f, 0x0b, 0x99, 0xd7, 0x49,
	0xb7, 0xfb, 0x49, 0x3c, 0xe3, 0xa1, 0x1b, 0x34, 0xe2, 0x41, 0x9f, 0x89, 0xe0, 0x92, 0x17, 0x18,
	0x14, 0x2b, 0x81, 0xee, 0x96, 0xce, 0xf0, 0x06, 0xa6, 0x2e, 0x25, 0x0b, 0x56, 0x60, 0xb4, 0x8a,
	0xdd, 0x6d, 0x1d, 0xb7, 0xb6, 0xc9, 0x33, 0x18, 0x04, 0x02, 0x75, 0x00, 0x3f, 0x64, 0x05, 0xba,
	0xbd, 0x51, 0x6b, 0xdc, 0xa1, 0xfd, 0x0a, 0x9c, 0xb1, 0x02, 0xc9, 0x73, 0xd8, 0x89, 0x99, 0x2c,
	0xfc, 0x95, 0xc4, 0xd0, 0xb0, 0x6c, 0xc3, 0x52, 0xe8, 0x85, 0xc4, 0x50, 0xb1, 0xbc, 0x3f, 0x5a,
	0x30, 0x10, 0x72, 0x9d, 0x06, 0x27, 0xc8, 0x54, 0x5c, 0xa9, 0xda, 0xe4, 0x96, 0x15, 0x85, 0x90,
	0x6e, 0x6b, 0xd4, 0x1a, 0xf7, 0x68, 0x69, 0x29, 0x3c, 0xc4, 0x18, 0x0b, 0x55, 0x5b, 0x8d, 0x1b,
	0x4b, 0x5d, 0x34, 0xc8, 0x92, 0x5c, 0xa0, 0x54, 0xd9, 0x53, 0x9e, 0xda, 0x26, 0xcf, 0x61, 0xf0,
	0x91, 0x87, 0x5c, 0x60, 0xa0, 0xae, 0xa5, 0x33, 0xa8, 0x08, 0x9b, 0xa0, 0xf7, 0x02, 0x9c, 0xbb,
	0x48, 0xd6, 0x17, 0x68, 0x0a, 0xb6, 0x36, 0x05, 0xbd, 0x3f, 0x3b, 0xb0, 0xfb, 0xbe, 0x4a, 0xee,
	0x5b, 0x64, 0x21, 0x0a, 0x72, 0x00, 0xed, 0x48, 0xea, 0x2e, 0xd8, 0x99, 0x3c, 0x69, 0xa4, 0xbe,
	0xe6, 0x9d, 0x2c, 0xd4, 0xac, 0xd0, 0x76, 0x24, 0xc9, 0xb7, 0x60, 0x05, 0x82, 0xaf, 0xf4, 0x13,
	0x76, 0x26, 0x7b, 0xcd, 0xc6, 0xa0, 0xef, 0x2e, 0x34, 0x4d, 0x13, 0xc8, 0x01, 0x74, 0x79, 0x98,
	0xb0, 0x5c, 0x37, 0x84, 0x33, 0xd9, 0x6f, 0x30, 0xeb, 0xe9, 0xa3, 0x86, 0xa2, 0x5e, 0x29, 0xcb,
	0xa6, 0x9c, 0xb3, 0x04, 0xa5, 0x6b, 0xe9, 0x26, 0xda, 0x04, 0xc9, 0x77, 0x60, 0x57, 0x40, 0xd5,
	0x28, 0xcd, 0xf8, 0x55, 0x5b, 0xd3, 0x7b, 0x16, 0x71, 0x61, 0x3b, 0x17, 0x18, 0xae, 0x92, 0xdc,
	0xdd, 0xd6, 0x89, 0xa8, 0x4c, 0xf2, 0xc3, 0x83, 0xaa, 0xe9, 0x0e, 0x70, 0x26, 0x6e, 0x43, 0x70,
	0xc3, 0x4f, 0x1f, 0x14, 0xd9, 0x85, 0x6d, 0x81, 0x91, 0x40, 0x79, 0xa9, 0xbb, 0xa2, 0x47, 0x2b,
	0x93, 0xbc, 0xde, 0x28, 0x86, 0x0b, 0x5a, 0xf7, 0x71, 0x43, 0xb7, 0xe1, 0xa5, 0x4d, 0xaa, 0x77,
	0x02, 0xc3, 0x3a, 0xe5, 0xd3, 0x2c, 0x2d, 0x44, 0x16, 0xab, 0x38, 0x72, 0x15, 0x04, 0xa6, 0x94,
	0xaa, 0x89, 0x2b, 0x53, 0x79, 0x12, 0x94, 0x92, 0x2d, 0x4d, 0x3f, 0xd9, 0xb4, 0x32, 0xbd, 0x57,
	0x30, 0xa8, 0x75, 0x16, 0xeb, 0x34, 0x50, 0xe3, 0x12, 0xf1, 0x94, 0xc5, 0x67, 0x02, 0x67, 0x2a,
	0x17, 0x46, 0x69, 0x03, 0xf3, 0xfe, 0xea, 0xc0, 0x50, 0x65, 0xc6, 0x57, 0x43, 0x22, 0x7d, 0x4c,
	0x0b, 0xb1, 0x56, 0x73, 0x12, 0x09, 0xc4, 0x3b, 0x9e, 0x2e, 0xfd, 0x82, 0x97, 0xab, 0x62, 0x40,
	0xfb, 0x15, 0x78, 0xce, 0x13, 0x24, 0x5f, 0x81, 0x13, 0x89, 0xec, 0x0e, 0x53, 0x43, 0x69, 0x6b,
	0x0a, 0x18, 0x48, 0x13, 0xbe, 0x86, 0x7e, 0x82, 0x89, 0x16, 0xd7, 0x8c, 0x8e, 0x66, 0x38, 0x25,
	0xa6, 0x29, 0xcf, 0x60, 0x90, 0x60, 0x72, 0x23, 0x78, 0x81, 0x86, 0x63, 0x99, 0x40, 0x15, 0x58,
	0x91, 0x72, 0xb6, 0x44, 0xe9, 0xcb, 0x80, 0xa5, 0x29, 0x86, 0x7a, 0xb1, 0x5a, 0xb4, 0xaf, 0xc1,
	0x85, 0xc1, 0xc8, 0x4b, 0xd8, 0x2f, 0x49, 0x57, 0x3c, 0xcf, 0x31, 0xf4, 0x73, 0x26, 0x30, 0x2d,
	0xf4, 0x8a, 0xb0, 0x28, 0x31, 0x5c, 0xe3, 0x3a, 0xd3, 0x9e, 0x7b, 0x59, 0x15, 0xa9, 0xc0, 0x54,
	0x6f, 0x8b, 0x4a, 0xf6, 0x17, 0x83, 0x29, 0x12, 0x17, 0x09, 0xcb, 0x7d, 0x81, 0x32, 0x8b, 0xaf,
	0xcd, 0xc6, 0x18, 0xd0, 0xbe, 0x06, 0xa9, 0xc1, 0xc8, 0x17, 0x00, 0x46, 0x29, 0x66, 0x77, 0x6b,
	0xd7, 0xd6, 0x32, 0xb6, 0x46, 0x4e, 0xd9, 0xdd, 0xba, 0x72, 0xfb, 0x39, 0xcf, 0xcb, 0xc6, 0x28,
	0xdd, 0x67, 0x0a, 0x50, 0xfb, 0xa6, 0x76, 0xfb, 0x1f, 0x57, 0x91, 0x74, 0x1d, 0x4d, 0xe9, 0x57,
	0x94, 0xe3, 0x55, 0x24, 0xbd, 0x7f, 0x5a, 0xb0, 0x27, 0x50, 0x16, 0x99, 0xc0, 0x8d, 0x52, 0x7d,
	0x63, 0xbe, 0x96, 0xbe, 0x1a, 0x75, 0x26, 0xd0, 0xfc, 0xd1, 0x2c, 0x6a, 0xde, 0x36, 0x2d, 0x41,
	0x72, 0x00, 0x8f, 0x36, 0xd3, 0x13, 0x64, 0x37, 0xba, 0x64, 0x16, 0xdd, 0x6d, 0xe6, 0x66, 0x9a,
	0xdd, 0xa8, 0xba, 0x45, 0x99, 0xb8, 0xaa, 0x8b, 0x5f, 0xd6, 0xad, 0xc4, 0xaa, 0xd2, 0x56, 0x97,
	0x69, 0x94, 0xcd, 0x29, 0x31, 0x4d, 0xa9, 0x2f, 0x56, 0x82, 0xaa, 0x6c, 0xad, 0xfa, 0x62, 0xb4,
	0x04, 0xbd, 0x5b, 0x70, 0x9a, 0xcf, 0x39, 0x02, 0x2b, 0x34, 0xad, 0xaa, 0xc6, 0xe7, 0x69, 0x63,
	0x7c, 0x1e, 0x36, 0x29, 0xd5, 0x44, 0xf2, 0x5a, 0x0d, 0xa4, 0xd6, 0xd2, 0xe3, 0xe0, 0x4c, 0xbe,
	0x6c, 0x8e, 0xf2, 0xa7, 0x09, 0xa3, 0x15, 0xfd, 0x60, 0xde, 0xd8, 0x88, 0x66, 0xd3, 0x11, 0x1b,
	0xba, 0x74, 0xf1, 0xeb, 0x7c, 0x3a, 0xfc, 0x4c, 0x1d, 0x8f, 0xcf, 0xe9, 0xc9, 0x62, 0xd8, 0x22,
	0xdb, 0xd0, 0xf9, 0xed, 0x64, 0x31, 0x6c, 0xab, 0x03, 0x3d, 0x9e, 0x0d, 0x3b, 0x64, 0x0f, 0x76,
	0x8f, 0x4f, 0x3f, 0x4c, 0x7f, 0xf6, 0xdf, 0xcc, 0x67, 0xbe, 0xf9, 0xc2, 0x3a, 0x38, 0x82, 0x5e,
	0xb5, 0x0b, 0xc9, 0x0e, 0x80, 0x3a, 0xfb, 0x0d, 0xb5, 0xb3, 0xb7, 0x6f, 0x2e, 0x4e, 0x87, 0x2d,
	0xd2, 0x03, 0x6b, 0xfe, 0x61, 0xfe, 0xe3, 0xb0, 0xfd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc9,
	0x83, 0x44, 0xf9, 0xb9, 0x08, 0x00, 0x00,
}
