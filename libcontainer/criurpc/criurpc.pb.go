// Code generated by protoc-gen-go.
// source: criurpc.proto
// DO NOT EDIT!

/*
Package criurpc is a generated protocol buffer package.

It is generated from these files:
	criurpc.proto

It has these top-level messages:
	CriuPageServerInfo
	CriuVethPair
	ExtMountMap
	InheritFd
	CgroupRoot
	UnixSk
	CriuOpts
	CriuDumpResp
	CriuRestoreResp
	CriuNotify
	CriuFeatures
	CriuReq
	CriuResp
*/
package criurpc

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type CriuCgMode int32

const (
	CriuCgMode_IGNORE  CriuCgMode = 0
	CriuCgMode_NONE    CriuCgMode = 1
	CriuCgMode_PROPS   CriuCgMode = 2
	CriuCgMode_SOFT    CriuCgMode = 3
	CriuCgMode_FULL    CriuCgMode = 4
	CriuCgMode_STRICT  CriuCgMode = 5
	CriuCgMode_DEFAULT CriuCgMode = 6
)

var CriuCgMode_name = map[int32]string{
	0: "IGNORE",
	1: "NONE",
	2: "PROPS",
	3: "SOFT",
	4: "FULL",
	5: "STRICT",
	6: "DEFAULT",
}
var CriuCgMode_value = map[string]int32{
	"IGNORE":  0,
	"NONE":    1,
	"PROPS":   2,
	"SOFT":    3,
	"FULL":    4,
	"STRICT":  5,
	"DEFAULT": 6,
}

func (x CriuCgMode) Enum() *CriuCgMode {
	p := new(CriuCgMode)
	*p = x
	return p
}
func (x CriuCgMode) String() string {
	return proto.EnumName(CriuCgMode_name, int32(x))
}
func (x *CriuCgMode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CriuCgMode_value, data, "CriuCgMode")
	if err != nil {
		return err
	}
	*x = CriuCgMode(value)
	return nil
}

type CriuReqType int32

const (
	CriuReqType_EMPTY         CriuReqType = 0
	CriuReqType_DUMP          CriuReqType = 1
	CriuReqType_RESTORE       CriuReqType = 2
	CriuReqType_CHECK         CriuReqType = 3
	CriuReqType_PRE_DUMP      CriuReqType = 4
	CriuReqType_PAGE_SERVER   CriuReqType = 5
	CriuReqType_NOTIFY        CriuReqType = 6
	CriuReqType_CPUINFO_DUMP  CriuReqType = 7
	CriuReqType_CPUINFO_CHECK CriuReqType = 8
	CriuReqType_FEATURE_CHECK CriuReqType = 9
)

var CriuReqType_name = map[int32]string{
	0: "EMPTY",
	1: "DUMP",
	2: "RESTORE",
	3: "CHECK",
	4: "PRE_DUMP",
	5: "PAGE_SERVER",
	6: "NOTIFY",
	7: "CPUINFO_DUMP",
	8: "CPUINFO_CHECK",
	9: "FEATURE_CHECK",
}
var CriuReqType_value = map[string]int32{
	"EMPTY":         0,
	"DUMP":          1,
	"RESTORE":       2,
	"CHECK":         3,
	"PRE_DUMP":      4,
	"PAGE_SERVER":   5,
	"NOTIFY":        6,
	"CPUINFO_DUMP":  7,
	"CPUINFO_CHECK": 8,
	"FEATURE_CHECK": 9,
}

func (x CriuReqType) Enum() *CriuReqType {
	p := new(CriuReqType)
	*p = x
	return p
}
func (x CriuReqType) String() string {
	return proto.EnumName(CriuReqType_name, int32(x))
}
func (x *CriuReqType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CriuReqType_value, data, "CriuReqType")
	if err != nil {
		return err
	}
	*x = CriuReqType(value)
	return nil
}

type CriuPageServerInfo struct {
	Address          *string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Port             *int32  `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	Pid              *int32  `protobuf:"varint,3,opt,name=pid" json:"pid,omitempty"`
	Fd               *int32  `protobuf:"varint,4,opt,name=fd" json:"fd,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CriuPageServerInfo) Reset()         { *m = CriuPageServerInfo{} }
func (m *CriuPageServerInfo) String() string { return proto.CompactTextString(m) }
func (*CriuPageServerInfo) ProtoMessage()    {}

func (m *CriuPageServerInfo) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

func (m *CriuPageServerInfo) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

func (m *CriuPageServerInfo) GetPid() int32 {
	if m != nil && m.Pid != nil {
		return *m.Pid
	}
	return 0
}

func (m *CriuPageServerInfo) GetFd() int32 {
	if m != nil && m.Fd != nil {
		return *m.Fd
	}
	return 0
}

type CriuVethPair struct {
	IfIn             *string `protobuf:"bytes,1,req,name=if_in" json:"if_in,omitempty"`
	IfOut            *string `protobuf:"bytes,2,req,name=if_out" json:"if_out,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CriuVethPair) Reset()         { *m = CriuVethPair{} }
func (m *CriuVethPair) String() string { return proto.CompactTextString(m) }
func (*CriuVethPair) ProtoMessage()    {}

func (m *CriuVethPair) GetIfIn() string {
	if m != nil && m.IfIn != nil {
		return *m.IfIn
	}
	return ""
}

func (m *CriuVethPair) GetIfOut() string {
	if m != nil && m.IfOut != nil {
		return *m.IfOut
	}
	return ""
}

type ExtMountMap struct {
	Key              *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Val              *string `protobuf:"bytes,2,req,name=val" json:"val,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ExtMountMap) Reset()         { *m = ExtMountMap{} }
func (m *ExtMountMap) String() string { return proto.CompactTextString(m) }
func (*ExtMountMap) ProtoMessage()    {}

func (m *ExtMountMap) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *ExtMountMap) GetVal() string {
	if m != nil && m.Val != nil {
		return *m.Val
	}
	return ""
}

type InheritFd struct {
	Key              *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Fd               *int32  `protobuf:"varint,2,req,name=fd" json:"fd,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *InheritFd) Reset()         { *m = InheritFd{} }
func (m *InheritFd) String() string { return proto.CompactTextString(m) }
func (*InheritFd) ProtoMessage()    {}

func (m *InheritFd) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *InheritFd) GetFd() int32 {
	if m != nil && m.Fd != nil {
		return *m.Fd
	}
	return 0
}

type CgroupRoot struct {
	Ctrl             *string `protobuf:"bytes,1,opt,name=ctrl" json:"ctrl,omitempty"`
	Path             *string `protobuf:"bytes,2,req,name=path" json:"path,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CgroupRoot) Reset()         { *m = CgroupRoot{} }
func (m *CgroupRoot) String() string { return proto.CompactTextString(m) }
func (*CgroupRoot) ProtoMessage()    {}

func (m *CgroupRoot) GetCtrl() string {
	if m != nil && m.Ctrl != nil {
		return *m.Ctrl
	}
	return ""
}

func (m *CgroupRoot) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

type UnixSk struct {
	Inode            *uint32 `protobuf:"varint,1,req,name=inode" json:"inode,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UnixSk) Reset()         { *m = UnixSk{} }
func (m *UnixSk) String() string { return proto.CompactTextString(m) }
func (*UnixSk) ProtoMessage()    {}

func (m *UnixSk) GetInode() uint32 {
	if m != nil && m.Inode != nil {
		return *m.Inode
	}
	return 0
}

type CriuOpts struct {
	ImagesDirFd       *int32              `protobuf:"varint,1,req,name=images_dir_fd" json:"images_dir_fd,omitempty"`
	Pid               *int32              `protobuf:"varint,2,opt,name=pid" json:"pid,omitempty"`
	LeaveRunning      *bool               `protobuf:"varint,3,opt,name=leave_running" json:"leave_running,omitempty"`
	ExtUnixSk         *bool               `protobuf:"varint,4,opt,name=ext_unix_sk" json:"ext_unix_sk,omitempty"`
	TcpEstablished    *bool               `protobuf:"varint,5,opt,name=tcp_established" json:"tcp_established,omitempty"`
	EvasiveDevices    *bool               `protobuf:"varint,6,opt,name=evasive_devices" json:"evasive_devices,omitempty"`
	ShellJob          *bool               `protobuf:"varint,7,opt,name=shell_job" json:"shell_job,omitempty"`
	FileLocks         *bool               `protobuf:"varint,8,opt,name=file_locks" json:"file_locks,omitempty"`
	LogLevel          *int32              `protobuf:"varint,9,opt,name=log_level,def=2" json:"log_level,omitempty"`
	LogFile           *string             `protobuf:"bytes,10,opt,name=log_file" json:"log_file,omitempty"`
	Ps                *CriuPageServerInfo `protobuf:"bytes,11,opt,name=ps" json:"ps,omitempty"`
	NotifyScripts     *bool               `protobuf:"varint,12,opt,name=notify_scripts" json:"notify_scripts,omitempty"`
	Root              *string             `protobuf:"bytes,13,opt,name=root" json:"root,omitempty"`
	ParentImg         *string             `protobuf:"bytes,14,opt,name=parent_img" json:"parent_img,omitempty"`
	TrackMem          *bool               `protobuf:"varint,15,opt,name=track_mem" json:"track_mem,omitempty"`
	AutoDedup         *bool               `protobuf:"varint,16,opt,name=auto_dedup" json:"auto_dedup,omitempty"`
	WorkDirFd         *int32              `protobuf:"varint,17,opt,name=work_dir_fd" json:"work_dir_fd,omitempty"`
	LinkRemap         *bool               `protobuf:"varint,18,opt,name=link_remap" json:"link_remap,omitempty"`
	Veths             []*CriuVethPair     `protobuf:"bytes,19,rep,name=veths" json:"veths,omitempty"`
	CpuCap            *uint32             `protobuf:"varint,20,opt,name=cpu_cap,def=4294967295" json:"cpu_cap,omitempty"`
	ForceIrmap        *bool               `protobuf:"varint,21,opt,name=force_irmap" json:"force_irmap,omitempty"`
	ExecCmd           []string            `protobuf:"bytes,22,rep,name=exec_cmd" json:"exec_cmd,omitempty"`
	ExtMnt            []*ExtMountMap      `protobuf:"bytes,23,rep,name=ext_mnt" json:"ext_mnt,omitempty"`
	ManageCgroups     *bool               `protobuf:"varint,24,opt,name=manage_cgroups" json:"manage_cgroups,omitempty"`
	CgRoot            []*CgroupRoot       `protobuf:"bytes,25,rep,name=cg_root" json:"cg_root,omitempty"`
	RstSibling        *bool               `protobuf:"varint,26,opt,name=rst_sibling" json:"rst_sibling,omitempty"`
	InheritFd         []*InheritFd        `protobuf:"bytes,27,rep,name=inherit_fd" json:"inherit_fd,omitempty"`
	AutoExtMnt        *bool               `protobuf:"varint,28,opt,name=auto_ext_mnt" json:"auto_ext_mnt,omitempty"`
	ExtSharing        *bool               `protobuf:"varint,29,opt,name=ext_sharing" json:"ext_sharing,omitempty"`
	ExtMasters        *bool               `protobuf:"varint,30,opt,name=ext_masters" json:"ext_masters,omitempty"`
	SkipMnt           []string            `protobuf:"bytes,31,rep,name=skip_mnt" json:"skip_mnt,omitempty"`
	EnableFs          []string            `protobuf:"bytes,32,rep,name=enable_fs" json:"enable_fs,omitempty"`
	UnixSkIno         []*UnixSk           `protobuf:"bytes,33,rep,name=unix_sk_ino" json:"unix_sk_ino,omitempty"`
	ManageCgroupsMode *CriuCgMode         `protobuf:"varint,34,opt,name=manage_cgroups_mode,enum=CriuCgMode" json:"manage_cgroups_mode,omitempty"`
	GhostLimit        *uint32             `protobuf:"varint,35,opt,name=ghost_limit,def=1048576" json:"ghost_limit,omitempty"`
	IrmapScanPaths    []string            `protobuf:"bytes,36,rep,name=irmap_scan_paths" json:"irmap_scan_paths,omitempty"`
	External          []string            `protobuf:"bytes,37,rep,name=external" json:"external,omitempty"`
	EmptyNs           *uint32             `protobuf:"varint,38,opt,name=empty_ns" json:"empty_ns,omitempty"`
	NoSeccomp         *bool               `protobuf:"varint,39,opt,name=no_seccomp" json:"no_seccomp,omitempty"`
	LazyPages         *bool               `protobuf:"varint,48,opt,name=lazy_pages" json:"lazy_pages,omitempty"`
	XXX_unrecognized  []byte              `json:"-"`
}

func (m *CriuOpts) Reset()         { *m = CriuOpts{} }
func (m *CriuOpts) String() string { return proto.CompactTextString(m) }
func (*CriuOpts) ProtoMessage()    {}

const Default_CriuOpts_LogLevel int32 = 2
const Default_CriuOpts_CpuCap uint32 = 4294967295
const Default_CriuOpts_GhostLimit uint32 = 1048576

func (m *CriuOpts) GetImagesDirFd() int32 {
	if m != nil && m.ImagesDirFd != nil {
		return *m.ImagesDirFd
	}
	return 0
}

func (m *CriuOpts) GetPid() int32 {
	if m != nil && m.Pid != nil {
		return *m.Pid
	}
	return 0
}

func (m *CriuOpts) GetLeaveRunning() bool {
	if m != nil && m.LeaveRunning != nil {
		return *m.LeaveRunning
	}
	return false
}

func (m *CriuOpts) GetExtUnixSk() bool {
	if m != nil && m.ExtUnixSk != nil {
		return *m.ExtUnixSk
	}
	return false
}

func (m *CriuOpts) GetTcpEstablished() bool {
	if m != nil && m.TcpEstablished != nil {
		return *m.TcpEstablished
	}
	return false
}

func (m *CriuOpts) GetLazyPages() bool {
	if m != nil && m.LazyPages != nil {
		return *m.LazyPages
	}
	return false
}

func (m *CriuOpts) GetEvasiveDevices() bool {
	if m != nil && m.EvasiveDevices != nil {
		return *m.EvasiveDevices
	}
	return false
}

func (m *CriuOpts) GetShellJob() bool {
	if m != nil && m.ShellJob != nil {
		return *m.ShellJob
	}
	return false
}

func (m *CriuOpts) GetFileLocks() bool {
	if m != nil && m.FileLocks != nil {
		return *m.FileLocks
	}
	return false
}

func (m *CriuOpts) GetLogLevel() int32 {
	if m != nil && m.LogLevel != nil {
		return *m.LogLevel
	}
	return Default_CriuOpts_LogLevel
}

func (m *CriuOpts) GetLogFile() string {
	if m != nil && m.LogFile != nil {
		return *m.LogFile
	}
	return ""
}

func (m *CriuOpts) GetPs() *CriuPageServerInfo {
	if m != nil {
		return m.Ps
	}
	return nil
}

func (m *CriuOpts) GetNotifyScripts() bool {
	if m != nil && m.NotifyScripts != nil {
		return *m.NotifyScripts
	}
	return false
}

func (m *CriuOpts) GetRoot() string {
	if m != nil && m.Root != nil {
		return *m.Root
	}
	return ""
}

func (m *CriuOpts) GetParentImg() string {
	if m != nil && m.ParentImg != nil {
		return *m.ParentImg
	}
	return ""
}

func (m *CriuOpts) GetTrackMem() bool {
	if m != nil && m.TrackMem != nil {
		return *m.TrackMem
	}
	return false
}

func (m *CriuOpts) GetAutoDedup() bool {
	if m != nil && m.AutoDedup != nil {
		return *m.AutoDedup
	}
	return false
}

func (m *CriuOpts) GetWorkDirFd() int32 {
	if m != nil && m.WorkDirFd != nil {
		return *m.WorkDirFd
	}
	return 0
}

func (m *CriuOpts) GetLinkRemap() bool {
	if m != nil && m.LinkRemap != nil {
		return *m.LinkRemap
	}
	return false
}

func (m *CriuOpts) GetVeths() []*CriuVethPair {
	if m != nil {
		return m.Veths
	}
	return nil
}

func (m *CriuOpts) GetCpuCap() uint32 {
	if m != nil && m.CpuCap != nil {
		return *m.CpuCap
	}
	return Default_CriuOpts_CpuCap
}

func (m *CriuOpts) GetForceIrmap() bool {
	if m != nil && m.ForceIrmap != nil {
		return *m.ForceIrmap
	}
	return false
}

func (m *CriuOpts) GetExecCmd() []string {
	if m != nil {
		return m.ExecCmd
	}
	return nil
}

func (m *CriuOpts) GetExtMnt() []*ExtMountMap {
	if m != nil {
		return m.ExtMnt
	}
	return nil
}

func (m *CriuOpts) GetManageCgroups() bool {
	if m != nil && m.ManageCgroups != nil {
		return *m.ManageCgroups
	}
	return false
}

func (m *CriuOpts) GetCgRoot() []*CgroupRoot {
	if m != nil {
		return m.CgRoot
	}
	return nil
}

func (m *CriuOpts) GetRstSibling() bool {
	if m != nil && m.RstSibling != nil {
		return *m.RstSibling
	}
	return false
}

func (m *CriuOpts) GetInheritFd() []*InheritFd {
	if m != nil {
		return m.InheritFd
	}
	return nil
}

func (m *CriuOpts) GetAutoExtMnt() bool {
	if m != nil && m.AutoExtMnt != nil {
		return *m.AutoExtMnt
	}
	return false
}

func (m *CriuOpts) GetExtSharing() bool {
	if m != nil && m.ExtSharing != nil {
		return *m.ExtSharing
	}
	return false
}

func (m *CriuOpts) GetExtMasters() bool {
	if m != nil && m.ExtMasters != nil {
		return *m.ExtMasters
	}
	return false
}

func (m *CriuOpts) GetSkipMnt() []string {
	if m != nil {
		return m.SkipMnt
	}
	return nil
}

func (m *CriuOpts) GetEnableFs() []string {
	if m != nil {
		return m.EnableFs
	}
	return nil
}

func (m *CriuOpts) GetUnixSkIno() []*UnixSk {
	if m != nil {
		return m.UnixSkIno
	}
	return nil
}

func (m *CriuOpts) GetManageCgroupsMode() CriuCgMode {
	if m != nil && m.ManageCgroupsMode != nil {
		return *m.ManageCgroupsMode
	}
	return CriuCgMode_IGNORE
}

func (m *CriuOpts) GetGhostLimit() uint32 {
	if m != nil && m.GhostLimit != nil {
		return *m.GhostLimit
	}
	return Default_CriuOpts_GhostLimit
}

func (m *CriuOpts) GetIrmapScanPaths() []string {
	if m != nil {
		return m.IrmapScanPaths
	}
	return nil
}

func (m *CriuOpts) GetExternal() []string {
	if m != nil {
		return m.External
	}
	return nil
}

func (m *CriuOpts) GetEmptyNs() uint32 {
	if m != nil && m.EmptyNs != nil {
		return *m.EmptyNs
	}
	return 0
}

func (m *CriuOpts) GetNoSeccomp() bool {
	if m != nil && m.NoSeccomp != nil {
		return *m.NoSeccomp
	}
	return false
}

type CriuDumpResp struct {
	Restored         *bool  `protobuf:"varint,1,opt,name=restored" json:"restored,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CriuDumpResp) Reset()         { *m = CriuDumpResp{} }
func (m *CriuDumpResp) String() string { return proto.CompactTextString(m) }
func (*CriuDumpResp) ProtoMessage()    {}

func (m *CriuDumpResp) GetRestored() bool {
	if m != nil && m.Restored != nil {
		return *m.Restored
	}
	return false
}

type CriuRestoreResp struct {
	Pid              *int32 `protobuf:"varint,1,req,name=pid" json:"pid,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CriuRestoreResp) Reset()         { *m = CriuRestoreResp{} }
func (m *CriuRestoreResp) String() string { return proto.CompactTextString(m) }
func (*CriuRestoreResp) ProtoMessage()    {}

func (m *CriuRestoreResp) GetPid() int32 {
	if m != nil && m.Pid != nil {
		return *m.Pid
	}
	return 0
}

type CriuNotify struct {
	Script           *string `protobuf:"bytes,1,opt,name=script" json:"script,omitempty"`
	Pid              *int32  `protobuf:"varint,2,opt,name=pid" json:"pid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CriuNotify) Reset()         { *m = CriuNotify{} }
func (m *CriuNotify) String() string { return proto.CompactTextString(m) }
func (*CriuNotify) ProtoMessage()    {}

func (m *CriuNotify) GetScript() string {
	if m != nil && m.Script != nil {
		return *m.Script
	}
	return ""
}

func (m *CriuNotify) GetPid() int32 {
	if m != nil && m.Pid != nil {
		return *m.Pid
	}
	return 0
}

//
// List of features which can queried via
// CRIU_REQ_TYPE__FEATURE_CHECK
type CriuFeatures struct {
	MemTrack         *bool  `protobuf:"varint,1,opt,name=mem_track" json:"mem_track,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CriuFeatures) Reset()         { *m = CriuFeatures{} }
func (m *CriuFeatures) String() string { return proto.CompactTextString(m) }
func (*CriuFeatures) ProtoMessage()    {}

func (m *CriuFeatures) GetMemTrack() bool {
	if m != nil && m.MemTrack != nil {
		return *m.MemTrack
	}
	return false
}

type CriuReq struct {
	Type          *CriuReqType `protobuf:"varint,1,req,name=type,enum=CriuReqType" json:"type,omitempty"`
	Opts          *CriuOpts    `protobuf:"bytes,2,opt,name=opts" json:"opts,omitempty"`
	NotifySuccess *bool        `protobuf:"varint,3,opt,name=notify_success" json:"notify_success,omitempty"`
	//
	// When set service won't close the connection but
	// will wait for more req-s to appear. Works not
	// for all request types.
	KeepOpen *bool `protobuf:"varint,4,opt,name=keep_open" json:"keep_open,omitempty"`
	//
	// 'features' can be used to query which features
	// are supported by the installed criu/kernel
	// via RPC.
	Features         *CriuFeatures `protobuf:"bytes,5,opt,name=features" json:"features,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *CriuReq) Reset()         { *m = CriuReq{} }
func (m *CriuReq) String() string { return proto.CompactTextString(m) }
func (*CriuReq) ProtoMessage()    {}

func (m *CriuReq) GetType() CriuReqType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return CriuReqType_EMPTY
}

func (m *CriuReq) GetOpts() *CriuOpts {
	if m != nil {
		return m.Opts
	}
	return nil
}

func (m *CriuReq) GetNotifySuccess() bool {
	if m != nil && m.NotifySuccess != nil {
		return *m.NotifySuccess
	}
	return false
}

func (m *CriuReq) GetKeepOpen() bool {
	if m != nil && m.KeepOpen != nil {
		return *m.KeepOpen
	}
	return false
}

func (m *CriuReq) GetFeatures() *CriuFeatures {
	if m != nil {
		return m.Features
	}
	return nil
}

type CriuResp struct {
	Type             *CriuReqType        `protobuf:"varint,1,req,name=type,enum=CriuReqType" json:"type,omitempty"`
	Success          *bool               `protobuf:"varint,2,req,name=success" json:"success,omitempty"`
	Dump             *CriuDumpResp       `protobuf:"bytes,3,opt,name=dump" json:"dump,omitempty"`
	Restore          *CriuRestoreResp    `protobuf:"bytes,4,opt,name=restore" json:"restore,omitempty"`
	Notify           *CriuNotify         `protobuf:"bytes,5,opt,name=notify" json:"notify,omitempty"`
	Ps               *CriuPageServerInfo `protobuf:"bytes,6,opt,name=ps" json:"ps,omitempty"`
	CrErrno          *int32              `protobuf:"varint,7,opt,name=cr_errno" json:"cr_errno,omitempty"`
	Features         *CriuFeatures       `protobuf:"bytes,8,opt,name=features" json:"features,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *CriuResp) Reset()         { *m = CriuResp{} }
func (m *CriuResp) String() string { return proto.CompactTextString(m) }
func (*CriuResp) ProtoMessage()    {}

func (m *CriuResp) GetType() CriuReqType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return CriuReqType_EMPTY
}

func (m *CriuResp) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

func (m *CriuResp) GetDump() *CriuDumpResp {
	if m != nil {
		return m.Dump
	}
	return nil
}

func (m *CriuResp) GetRestore() *CriuRestoreResp {
	if m != nil {
		return m.Restore
	}
	return nil
}

func (m *CriuResp) GetNotify() *CriuNotify {
	if m != nil {
		return m.Notify
	}
	return nil
}

func (m *CriuResp) GetPs() *CriuPageServerInfo {
	if m != nil {
		return m.Ps
	}
	return nil
}

func (m *CriuResp) GetCrErrno() int32 {
	if m != nil && m.CrErrno != nil {
		return *m.CrErrno
	}
	return 0
}

func (m *CriuResp) GetFeatures() *CriuFeatures {
	if m != nil {
		return m.Features
	}
	return nil
}

func init() {
	proto.RegisterEnum("CriuCgMode", CriuCgMode_name, CriuCgMode_value)
	proto.RegisterEnum("CriuReqType", CriuReqType_name, CriuReqType_value)
}
