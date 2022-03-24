// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

// Package avf contains generated bindings for API file avf.api.
//
// Contents:
//   4 messages
//
package avf

import (
	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	interface_types "github.com/projectcalico/vpp-dataplane/vpplink/binapi/vppapi/interface_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "avf"
	APIVersion = "1.0.0"
	VersionCrc = 0x9f5a6a20
)

// AvfCreate defines message 'avf_create'.
type AvfCreate struct {
	PciAddr    uint32 `binapi:"u32,name=pci_addr" json:"pci_addr,omitempty"`
	EnableElog int32  `binapi:"i32,name=enable_elog" json:"enable_elog,omitempty"`
	RxqNum     uint16 `binapi:"u16,name=rxq_num" json:"rxq_num,omitempty"`
	RxqSize    uint16 `binapi:"u16,name=rxq_size" json:"rxq_size,omitempty"`
	TxqSize    uint16 `binapi:"u16,name=txq_size" json:"txq_size,omitempty"`
}

func (m *AvfCreate) Reset()               { *m = AvfCreate{} }
func (*AvfCreate) GetMessageName() string { return "avf_create" }
func (*AvfCreate) GetCrcString() string   { return "daab8ae2" }
func (*AvfCreate) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (m *AvfCreate) GetRetVal() error {
	return nil
}

func (m *AvfCreate) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.PciAddr
	size += 4 // m.EnableElog
	size += 2 // m.RxqNum
	size += 2 // m.RxqSize
	size += 2 // m.TxqSize
	return size
}
func (m *AvfCreate) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.PciAddr)
	buf.EncodeInt32(m.EnableElog)
	buf.EncodeUint16(m.RxqNum)
	buf.EncodeUint16(m.RxqSize)
	buf.EncodeUint16(m.TxqSize)
	return buf.Bytes(), nil
}
func (m *AvfCreate) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.PciAddr = buf.DecodeUint32()
	m.EnableElog = buf.DecodeInt32()
	m.RxqNum = buf.DecodeUint16()
	m.RxqSize = buf.DecodeUint16()
	m.TxqSize = buf.DecodeUint16()
	return nil
}

// AvfCreateReply defines message 'avf_create_reply'.
type AvfCreateReply struct {
	Retval    int32                          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *AvfCreateReply) Reset()               { *m = AvfCreateReply{} }
func (*AvfCreateReply) GetMessageName() string { return "avf_create_reply" }
func (*AvfCreateReply) GetCrcString() string   { return "5383d31f" }
func (*AvfCreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (m *AvfCreateReply) GetRetVal() error {
	return api.RetvalToVPPApiError(m.Retval)
}

func (m *AvfCreateReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *AvfCreateReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *AvfCreateReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// AvfDelete defines message 'avf_delete'.
type AvfDelete struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *AvfDelete) Reset()               { *m = AvfDelete{} }
func (*AvfDelete) GetMessageName() string { return "avf_delete" }
func (*AvfDelete) GetCrcString() string   { return "f9e6675e" }
func (*AvfDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (m *AvfDelete) GetRetVal() error {
	return nil
}

func (m *AvfDelete) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *AvfDelete) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *AvfDelete) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// AvfDeleteReply defines message 'avf_delete_reply'.
type AvfDeleteReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *AvfDeleteReply) Reset()               { *m = AvfDeleteReply{} }
func (*AvfDeleteReply) GetMessageName() string { return "avf_delete_reply" }
func (*AvfDeleteReply) GetCrcString() string   { return "e8d4e804" }
func (*AvfDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (m *AvfDeleteReply) GetRetVal() error {
	return api.RetvalToVPPApiError(m.Retval)
}

func (m *AvfDeleteReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *AvfDeleteReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *AvfDeleteReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_avf_binapi_init() }
func file_avf_binapi_init() {
	api.RegisterMessage((*AvfCreate)(nil), "avf_create_daab8ae2")
	api.RegisterMessage((*AvfCreateReply)(nil), "avf_create_reply_5383d31f")
	api.RegisterMessage((*AvfDelete)(nil), "avf_delete_f9e6675e")
	api.RegisterMessage((*AvfDeleteReply)(nil), "avf_delete_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*AvfCreate)(nil),
		(*AvfCreateReply)(nil),
		(*AvfDelete)(nil),
		(*AvfDeleteReply)(nil),
	}
}
