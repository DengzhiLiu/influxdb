// Code generated by protoc-gen-gogo.
// source: internal.proto
// DO NOT EDIT!

/*
Package internal is a generated protocol buffer package.

It is generated from these files:
	internal.proto

It has these top-level messages:
	Source
	Dashboard
	DashboardCell
	Color
	Axis
	Template
	TemplateValue
	TemplateQuery
	Server
	Layout
	Cell
	Query
	TimeShift
	Range
	AlertRule
	User
	Role
	Organization
*/
package internal

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Source struct {
	ID                 int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name               string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Type               string `protobuf:"bytes,3,opt,name=Type,proto3" json:"Type,omitempty"`
	Username           string `protobuf:"bytes,4,opt,name=Username,proto3" json:"Username,omitempty"`
	Password           string `protobuf:"bytes,5,opt,name=Password,proto3" json:"Password,omitempty"`
	URL                string `protobuf:"bytes,6,opt,name=URL,proto3" json:"URL,omitempty"`
	Default            bool   `protobuf:"varint,7,opt,name=Default,proto3" json:"Default,omitempty"`
	Telegraf           string `protobuf:"bytes,8,opt,name=Telegraf,proto3" json:"Telegraf,omitempty"`
	InsecureSkipVerify bool   `protobuf:"varint,9,opt,name=InsecureSkipVerify,proto3" json:"InsecureSkipVerify,omitempty"`
	MetaURL            string `protobuf:"bytes,10,opt,name=MetaURL,proto3" json:"MetaURL,omitempty"`
	SharedSecret       string `protobuf:"bytes,11,opt,name=SharedSecret,proto3" json:"SharedSecret,omitempty"`
	Organization       string `protobuf:"bytes,12,opt,name=Organization,proto3" json:"Organization,omitempty"`
	Role               string `protobuf:"bytes,13,opt,name=Role,proto3" json:"Role,omitempty"`
}

func (m *Source) Reset()                    { *m = Source{} }
func (m *Source) String() string            { return proto.CompactTextString(m) }
func (*Source) ProtoMessage()               {}
func (*Source) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{0} }

type Dashboard struct {
	ID           int64            `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name         string           `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Cells        []*DashboardCell `protobuf:"bytes,3,rep,name=cells" json:"cells,omitempty"`
	Templates    []*Template      `protobuf:"bytes,4,rep,name=templates" json:"templates,omitempty"`
	Organization string           `protobuf:"bytes,5,opt,name=Organization,proto3" json:"Organization,omitempty"`
}

func (m *Dashboard) Reset()                    { *m = Dashboard{} }
func (m *Dashboard) String() string            { return proto.CompactTextString(m) }
func (*Dashboard) ProtoMessage()               {}
func (*Dashboard) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{1} }

func (m *Dashboard) GetCells() []*DashboardCell {
	if m != nil {
		return m.Cells
	}
	return nil
}

func (m *Dashboard) GetTemplates() []*Template {
	if m != nil {
		return m.Templates
	}
	return nil
}

type DashboardCell struct {
	X       int32            `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y       int32            `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	W       int32            `protobuf:"varint,3,opt,name=w,proto3" json:"w,omitempty"`
	H       int32            `protobuf:"varint,4,opt,name=h,proto3" json:"h,omitempty"`
	Queries []*Query         `protobuf:"bytes,5,rep,name=queries" json:"queries,omitempty"`
	Name    string           `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Type    string           `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	ID      string           `protobuf:"bytes,8,opt,name=ID,proto3" json:"ID,omitempty"`
	Axes    map[string]*Axis `protobuf:"bytes,9,rep,name=axes" json:"axes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	Colors  []*Color         `protobuf:"bytes,10,rep,name=colors" json:"colors,omitempty"`
}

func (m *DashboardCell) Reset()                    { *m = DashboardCell{} }
func (m *DashboardCell) String() string            { return proto.CompactTextString(m) }
func (*DashboardCell) ProtoMessage()               {}
func (*DashboardCell) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{2} }

func (m *DashboardCell) GetQueries() []*Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

func (m *DashboardCell) GetAxes() map[string]*Axis {
	if m != nil {
		return m.Axes
	}
	return nil
}

func (m *DashboardCell) GetColors() []*Color {
	if m != nil {
		return m.Colors
	}
	return nil
}

type Color struct {
	ID    string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Type  string `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	Hex   string `protobuf:"bytes,3,opt,name=Hex,proto3" json:"Hex,omitempty"`
	Name  string `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	Value string `protobuf:"bytes,5,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (m *Color) Reset()                    { *m = Color{} }
func (m *Color) String() string            { return proto.CompactTextString(m) }
func (*Color) ProtoMessage()               {}
func (*Color) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{3} }

type Axis struct {
	LegacyBounds []int64  `protobuf:"varint,1,rep,name=legacyBounds" json:"legacyBounds,omitempty"`
	Bounds       []string `protobuf:"bytes,2,rep,name=bounds" json:"bounds,omitempty"`
	Label        string   `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	Prefix       string   `protobuf:"bytes,4,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Suffix       string   `protobuf:"bytes,5,opt,name=suffix,proto3" json:"suffix,omitempty"`
	Base         string   `protobuf:"bytes,6,opt,name=base,proto3" json:"base,omitempty"`
	Scale        string   `protobuf:"bytes,7,opt,name=scale,proto3" json:"scale,omitempty"`
}

func (m *Axis) Reset()                    { *m = Axis{} }
func (m *Axis) String() string            { return proto.CompactTextString(m) }
func (*Axis) ProtoMessage()               {}
func (*Axis) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{4} }

type Template struct {
	ID      string           `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TempVar string           `protobuf:"bytes,2,opt,name=temp_var,json=tempVar,proto3" json:"temp_var,omitempty"`
	Values  []*TemplateValue `protobuf:"bytes,3,rep,name=values" json:"values,omitempty"`
	Type    string           `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Label   string           `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
	Query   *TemplateQuery   `protobuf:"bytes,6,opt,name=query" json:"query,omitempty"`
}

func (m *Template) Reset()                    { *m = Template{} }
func (m *Template) String() string            { return proto.CompactTextString(m) }
func (*Template) ProtoMessage()               {}
func (*Template) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{5} }

func (m *Template) GetValues() []*TemplateValue {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *Template) GetQuery() *TemplateQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

type TemplateValue struct {
	Type     string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Value    string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Selected bool   `protobuf:"varint,3,opt,name=selected,proto3" json:"selected,omitempty"`
}

func (m *TemplateValue) Reset()                    { *m = TemplateValue{} }
func (m *TemplateValue) String() string            { return proto.CompactTextString(m) }
func (*TemplateValue) ProtoMessage()               {}
func (*TemplateValue) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{6} }

type TemplateQuery struct {
	Command     string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Db          string `protobuf:"bytes,2,opt,name=db,proto3" json:"db,omitempty"`
	Rp          string `protobuf:"bytes,3,opt,name=rp,proto3" json:"rp,omitempty"`
	Measurement string `protobuf:"bytes,4,opt,name=measurement,proto3" json:"measurement,omitempty"`
	TagKey      string `protobuf:"bytes,5,opt,name=tag_key,json=tagKey,proto3" json:"tag_key,omitempty"`
	FieldKey    string `protobuf:"bytes,6,opt,name=field_key,json=fieldKey,proto3" json:"field_key,omitempty"`
}

func (m *TemplateQuery) Reset()                    { *m = TemplateQuery{} }
func (m *TemplateQuery) String() string            { return proto.CompactTextString(m) }
func (*TemplateQuery) ProtoMessage()               {}
func (*TemplateQuery) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{7} }

type Server struct {
	ID           int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Username     string `protobuf:"bytes,3,opt,name=Username,proto3" json:"Username,omitempty"`
	Password     string `protobuf:"bytes,4,opt,name=Password,proto3" json:"Password,omitempty"`
	URL          string `protobuf:"bytes,5,opt,name=URL,proto3" json:"URL,omitempty"`
	SrcID        int64  `protobuf:"varint,6,opt,name=SrcID,proto3" json:"SrcID,omitempty"`
	Active       bool   `protobuf:"varint,7,opt,name=Active,proto3" json:"Active,omitempty"`
	Organization string `protobuf:"bytes,8,opt,name=Organization,proto3" json:"Organization,omitempty"`
}

func (m *Server) Reset()                    { *m = Server{} }
func (m *Server) String() string            { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()               {}
func (*Server) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{8} }

type Layout struct {
	ID          string  `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Application string  `protobuf:"bytes,2,opt,name=Application,proto3" json:"Application,omitempty"`
	Measurement string  `protobuf:"bytes,3,opt,name=Measurement,proto3" json:"Measurement,omitempty"`
	Cells       []*Cell `protobuf:"bytes,4,rep,name=Cells" json:"Cells,omitempty"`
	Autoflow    bool    `protobuf:"varint,5,opt,name=Autoflow,proto3" json:"Autoflow,omitempty"`
}

func (m *Layout) Reset()                    { *m = Layout{} }
func (m *Layout) String() string            { return proto.CompactTextString(m) }
func (*Layout) ProtoMessage()               {}
func (*Layout) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{9} }

func (m *Layout) GetCells() []*Cell {
	if m != nil {
		return m.Cells
	}
	return nil
}

type Cell struct {
	X       int32            `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y       int32            `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	W       int32            `protobuf:"varint,3,opt,name=w,proto3" json:"w,omitempty"`
	H       int32            `protobuf:"varint,4,opt,name=h,proto3" json:"h,omitempty"`
	Queries []*Query         `protobuf:"bytes,5,rep,name=queries" json:"queries,omitempty"`
	I       string           `protobuf:"bytes,6,opt,name=i,proto3" json:"i,omitempty"`
	Name    string           `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Yranges []int64          `protobuf:"varint,8,rep,name=yranges" json:"yranges,omitempty"`
	Ylabels []string         `protobuf:"bytes,9,rep,name=ylabels" json:"ylabels,omitempty"`
	Type    string           `protobuf:"bytes,10,opt,name=type,proto3" json:"type,omitempty"`
	Axes    map[string]*Axis `protobuf:"bytes,11,rep,name=axes" json:"axes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Cell) Reset()                    { *m = Cell{} }
func (m *Cell) String() string            { return proto.CompactTextString(m) }
func (*Cell) ProtoMessage()               {}
func (*Cell) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{10} }

func (m *Cell) GetQueries() []*Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

func (m *Cell) GetAxes() map[string]*Axis {
	if m != nil {
		return m.Axes
	}
	return nil
}

type Query struct {
	Command  string       `protobuf:"bytes,1,opt,name=Command,proto3" json:"Command,omitempty"`
	DB       string       `protobuf:"bytes,2,opt,name=DB,proto3" json:"DB,omitempty"`
	RP       string       `protobuf:"bytes,3,opt,name=RP,proto3" json:"RP,omitempty"`
	GroupBys []string     `protobuf:"bytes,4,rep,name=GroupBys" json:"GroupBys,omitempty"`
	Wheres   []string     `protobuf:"bytes,5,rep,name=Wheres" json:"Wheres,omitempty"`
	Label    string       `protobuf:"bytes,6,opt,name=Label,proto3" json:"Label,omitempty"`
	Range    *Range       `protobuf:"bytes,7,opt,name=Range" json:"Range,omitempty"`
	Source   string       `protobuf:"bytes,8,opt,name=Source,proto3" json:"Source,omitempty"`
	Shifts   []*TimeShift `protobuf:"bytes,9,rep,name=Shifts" json:"Shifts,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{11} }

func (m *Query) GetRange() *Range {
	if m != nil {
		return m.Range
	}
	return nil
}

func (m *Query) GetShifts() []*TimeShift {
	if m != nil {
		return m.Shifts
	}
	return nil
}

type TimeShift struct {
	Label    string `protobuf:"bytes,1,opt,name=Label,proto3" json:"Label,omitempty"`
	Unit     string `protobuf:"bytes,2,opt,name=Unit,proto3" json:"Unit,omitempty"`
	Quantity string `protobuf:"bytes,3,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
}

func (m *TimeShift) Reset()                    { *m = TimeShift{} }
func (m *TimeShift) String() string            { return proto.CompactTextString(m) }
func (*TimeShift) ProtoMessage()               {}
func (*TimeShift) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{12} }

type Range struct {
	Upper int64 `protobuf:"varint,1,opt,name=Upper,proto3" json:"Upper,omitempty"`
	Lower int64 `protobuf:"varint,2,opt,name=Lower,proto3" json:"Lower,omitempty"`
}

func (m *Range) Reset()                    { *m = Range{} }
func (m *Range) String() string            { return proto.CompactTextString(m) }
func (*Range) ProtoMessage()               {}
func (*Range) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{13} }

type AlertRule struct {
	ID     string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	JSON   string `protobuf:"bytes,2,opt,name=JSON,proto3" json:"JSON,omitempty"`
	SrcID  int64  `protobuf:"varint,3,opt,name=SrcID,proto3" json:"SrcID,omitempty"`
	KapaID int64  `protobuf:"varint,4,opt,name=KapaID,proto3" json:"KapaID,omitempty"`
}

func (m *AlertRule) Reset()                    { *m = AlertRule{} }
func (m *AlertRule) String() string            { return proto.CompactTextString(m) }
func (*AlertRule) ProtoMessage()               {}
func (*AlertRule) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{14} }

type User struct {
	ID         uint64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name       string  `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Provider   string  `protobuf:"bytes,3,opt,name=Provider,proto3" json:"Provider,omitempty"`
	Scheme     string  `protobuf:"bytes,4,opt,name=Scheme,proto3" json:"Scheme,omitempty"`
	Roles      []*Role `protobuf:"bytes,5,rep,name=Roles" json:"Roles,omitempty"`
	SuperAdmin bool    `protobuf:"varint,6,opt,name=SuperAdmin,proto3" json:"SuperAdmin,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{15} }

func (m *User) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

type Role struct {
	Organization string `protobuf:"bytes,1,opt,name=Organization,proto3" json:"Organization,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (m *Role) Reset()                    { *m = Role{} }
func (m *Role) String() string            { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()               {}
func (*Role) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{16} }

type Organization struct {
	ID          uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	DefaultRole string `protobuf:"bytes,3,opt,name=DefaultRole,proto3" json:"DefaultRole,omitempty"`
	Public      bool   `protobuf:"varint,4,opt,name=Public,proto3" json:"Public,omitempty"`
}

func (m *Organization) Reset()                    { *m = Organization{} }
func (m *Organization) String() string            { return proto.CompactTextString(m) }
func (*Organization) ProtoMessage()               {}
func (*Organization) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{17} }

func init() {
	proto.RegisterType((*Source)(nil), "internal.Source")
	proto.RegisterType((*Dashboard)(nil), "internal.Dashboard")
	proto.RegisterType((*DashboardCell)(nil), "internal.DashboardCell")
	proto.RegisterType((*Color)(nil), "internal.Color")
	proto.RegisterType((*Axis)(nil), "internal.Axis")
	proto.RegisterType((*Template)(nil), "internal.Template")
	proto.RegisterType((*TemplateValue)(nil), "internal.TemplateValue")
	proto.RegisterType((*TemplateQuery)(nil), "internal.TemplateQuery")
	proto.RegisterType((*Server)(nil), "internal.Server")
	proto.RegisterType((*Layout)(nil), "internal.Layout")
	proto.RegisterType((*Cell)(nil), "internal.Cell")
	proto.RegisterType((*Query)(nil), "internal.Query")
	proto.RegisterType((*TimeShift)(nil), "internal.TimeShift")
	proto.RegisterType((*Range)(nil), "internal.Range")
	proto.RegisterType((*AlertRule)(nil), "internal.AlertRule")
	proto.RegisterType((*User)(nil), "internal.User")
	proto.RegisterType((*Role)(nil), "internal.Role")
	proto.RegisterType((*Organization)(nil), "internal.Organization")
}

func init() { proto.RegisterFile("internal.proto", fileDescriptorInternal) }

var fileDescriptorInternal = []byte{
	// 1264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x57, 0xdf, 0x8e, 0xdb, 0xc4,
	0x17, 0x96, 0xe3, 0x38, 0xb1, 0x4f, 0xb6, 0xfd, 0x55, 0xf3, 0xab, 0xa8, 0x29, 0x12, 0x0a, 0x16,
	0x88, 0x45, 0xd0, 0x05, 0xb5, 0x42, 0x42, 0x5c, 0x20, 0x65, 0x37, 0xa8, 0x2c, 0xfd, 0xb7, 0x9d,
	0x74, 0xcb, 0x15, 0xaa, 0x26, 0xce, 0x49, 0x62, 0xd5, 0xb1, 0xcd, 0xd8, 0xde, 0x8d, 0x79, 0x18,
	0x24, 0x24, 0x9e, 0x00, 0x71, 0xcf, 0x2d, 0xe2, 0x96, 0x77, 0xe0, 0x15, 0xb8, 0x45, 0x67, 0x66,
	0xec, 0x38, 0x9b, 0x50, 0xf5, 0x02, 0x71, 0x37, 0xdf, 0x39, 0x93, 0x33, 0x67, 0xce, 0xf9, 0xce,
	0x37, 0x0e, 0x5c, 0x8f, 0x92, 0x02, 0x65, 0x22, 0xe2, 0xa3, 0x4c, 0xa6, 0x45, 0xca, 0xdc, 0x1a,
	0x07, 0x7f, 0x76, 0xa0, 0x37, 0x49, 0x4b, 0x19, 0x22, 0xbb, 0x0e, 0x9d, 0xd3, 0xb1, 0x6f, 0x0d,
	0xad, 0x43, 0x9b, 0x77, 0x4e, 0xc7, 0x8c, 0x41, 0xf7, 0xb1, 0x58, 0xa1, 0xdf, 0x19, 0x5a, 0x87,
	0x1e, 0x57, 0x6b, 0xb2, 0x3d, 0xab, 0x32, 0xf4, 0x6d, 0x6d, 0xa3, 0x35, 0xbb, 0x0d, 0xee, 0x79,
	0x4e, 0xd1, 0x56, 0xe8, 0x77, 0x95, 0xbd, 0xc1, 0xe4, 0x3b, 0x13, 0x79, 0x7e, 0x99, 0xca, 0x99,
	0xef, 0x68, 0x5f, 0x8d, 0xd9, 0x0d, 0xb0, 0xcf, 0xf9, 0x43, 0xbf, 0xa7, 0xcc, 0xb4, 0x64, 0x3e,
	0xf4, 0xc7, 0x38, 0x17, 0x65, 0x5c, 0xf8, 0xfd, 0xa1, 0x75, 0xe8, 0xf2, 0x1a, 0x52, 0x9c, 0x67,
	0x18, 0xe3, 0x42, 0x8a, 0xb9, 0xef, 0xea, 0x38, 0x35, 0x66, 0x47, 0xc0, 0x4e, 0x93, 0x1c, 0xc3,
	0x52, 0xe2, 0xe4, 0x65, 0x94, 0x3d, 0x47, 0x19, 0xcd, 0x2b, 0xdf, 0x53, 0x01, 0xf6, 0x78, 0xe8,
	0x94, 0x47, 0x58, 0x08, 0x3a, 0x1b, 0x54, 0xa8, 0x1a, 0xb2, 0x00, 0x0e, 0x26, 0x4b, 0x21, 0x71,
	0x36, 0xc1, 0x50, 0x62, 0xe1, 0x0f, 0x94, 0x7b, 0xcb, 0x46, 0x7b, 0x9e, 0xc8, 0x85, 0x48, 0xa2,
	0xef, 0x45, 0x11, 0xa5, 0x89, 0x7f, 0xa0, 0xf7, 0xb4, 0x6d, 0x54, 0x25, 0x9e, 0xc6, 0xe8, 0x5f,
	0xd3, 0x55, 0xa2, 0x75, 0xf0, 0x8b, 0x05, 0xde, 0x58, 0xe4, 0xcb, 0x69, 0x2a, 0xe4, 0xec, 0xb5,
	0x6a, 0x7d, 0x07, 0x9c, 0x10, 0xe3, 0x38, 0xf7, 0xed, 0xa1, 0x7d, 0x38, 0xb8, 0x7b, 0xeb, 0xa8,
	0x69, 0x62, 0x13, 0xe7, 0x04, 0xe3, 0x98, 0xeb, 0x5d, 0xec, 0x13, 0xf0, 0x0a, 0x5c, 0x65, 0xb1,
	0x28, 0x30, 0xf7, 0xbb, 0xea, 0x27, 0x6c, 0xf3, 0x93, 0x67, 0xc6, 0xc5, 0x37, 0x9b, 0x76, 0xae,
	0xe2, 0xec, 0x5e, 0x25, 0xf8, 0xa3, 0x03, 0xd7, 0xb6, 0x8e, 0x63, 0x07, 0x60, 0xad, 0x55, 0xe6,
	0x0e, 0xb7, 0xd6, 0x84, 0x2a, 0x95, 0xb5, 0xc3, 0xad, 0x8a, 0xd0, 0xa5, 0xe2, 0x86, 0xc3, 0xad,
	0x4b, 0x42, 0x4b, 0xc5, 0x08, 0x87, 0x5b, 0x4b, 0xf6, 0x01, 0xf4, 0xbf, 0x2b, 0x51, 0x46, 0x98,
	0xfb, 0x8e, 0xca, 0xee, 0x7f, 0x9b, 0xec, 0x9e, 0x96, 0x28, 0x2b, 0x5e, 0xfb, 0xa9, 0x1a, 0x8a,
	0x4d, 0x9a, 0x1a, 0x6a, 0x4d, 0xb6, 0x82, 0x98, 0xd7, 0xd7, 0x36, 0x5a, 0x9b, 0x2a, 0x6a, 0x3e,
	0x50, 0x15, 0x3f, 0x85, 0xae, 0x58, 0x63, 0xee, 0x7b, 0x2a, 0xfe, 0x3b, 0xff, 0x50, 0xb0, 0xa3,
	0xd1, 0x1a, 0xf3, 0x2f, 0x93, 0x42, 0x56, 0x5c, 0x6d, 0x67, 0xef, 0x43, 0x2f, 0x4c, 0xe3, 0x54,
	0xe6, 0x3e, 0x5c, 0x4d, 0xec, 0x84, 0xec, 0xdc, 0xb8, 0x6f, 0xdf, 0x07, 0xaf, 0xf9, 0x2d, 0xd1,
	0xf7, 0x25, 0x56, 0xaa, 0x12, 0x1e, 0xa7, 0x25, 0x7b, 0x17, 0x9c, 0x0b, 0x11, 0x97, 0xba, 0x8b,
	0x83, 0xbb, 0xd7, 0x37, 0x61, 0x46, 0xeb, 0x28, 0xe7, 0xda, 0xf9, 0x79, 0xe7, 0x33, 0x2b, 0x58,
	0x80, 0xa3, 0x22, 0xb7, 0x78, 0xe0, 0xd5, 0x3c, 0x50, 0xf3, 0xd5, 0x69, 0xcd, 0xd7, 0x0d, 0xb0,
	0xbf, 0xc2, 0xb5, 0x19, 0x39, 0x5a, 0x36, 0x6c, 0xe9, 0xb6, 0xd8, 0x72, 0x13, 0x9c, 0xe7, 0xea,
	0x70, 0xdd, 0x45, 0x0d, 0x82, 0x9f, 0x2d, 0xe8, 0xd2, 0xe1, 0xd4, 0xeb, 0x18, 0x17, 0x22, 0xac,
	0x8e, 0xd3, 0x32, 0x99, 0xe5, 0xbe, 0x35, 0xb4, 0x0f, 0x6d, 0xbe, 0x65, 0x63, 0x6f, 0x40, 0x6f,
	0xaa, 0xbd, 0x9d, 0xa1, 0x7d, 0xe8, 0x71, 0x83, 0x28, 0x74, 0x2c, 0xa6, 0x18, 0x9b, 0x14, 0x34,
	0xa0, 0xdd, 0x99, 0xc4, 0x79, 0xb4, 0x36, 0x69, 0x18, 0x44, 0xf6, 0xbc, 0x9c, 0x93, 0x5d, 0x67,
	0x62, 0x10, 0x25, 0x3d, 0x15, 0x79, 0xd3, 0x54, 0x5a, 0x53, 0xe4, 0x3c, 0x14, 0x71, 0xdd, 0x55,
	0x0d, 0x82, 0x5f, 0x2d, 0x9a, 0x76, 0xcd, 0xd2, 0x9d, 0x0a, 0xbd, 0x09, 0x2e, 0x31, 0xf8, 0xc5,
	0x85, 0x90, 0xa6, 0x4a, 0x7d, 0xc2, 0xcf, 0x85, 0x64, 0x1f, 0x43, 0x4f, 0x95, 0x78, 0xcf, 0xc4,
	0xd4, 0xe1, 0x54, 0x55, 0xb8, 0xd9, 0xd6, 0x70, 0xaa, 0xdb, 0xe2, 0x54, 0x73, 0x59, 0xa7, 0x7d,
	0xd9, 0x3b, 0xe0, 0x10, 0x39, 0x2b, 0x95, 0xfd, 0xde, 0xc8, 0x9a, 0xc2, 0x7a, 0x57, 0x70, 0x0e,
	0xd7, 0xb6, 0x4e, 0x6c, 0x4e, 0xb2, 0xb6, 0x4f, 0xda, 0xd0, 0xc5, 0x33, 0xf4, 0x20, 0xa5, 0xcb,
	0x31, 0xc6, 0xb0, 0xc0, 0x99, 0xaa, 0xb7, 0xcb, 0x1b, 0x1c, 0xfc, 0x68, 0x6d, 0xe2, 0xaa, 0xf3,
	0x48, 0xcb, 0xc2, 0x74, 0xb5, 0x12, 0xc9, 0xcc, 0x84, 0xae, 0x21, 0xd5, 0x6d, 0x36, 0x35, 0xa1,
	0x3b, 0xb3, 0x29, 0x61, 0x99, 0x99, 0x0e, 0x76, 0x64, 0xc6, 0x86, 0x30, 0x58, 0xa1, 0xc8, 0x4b,
	0x89, 0x2b, 0x4c, 0x0a, 0x53, 0x82, 0xb6, 0x89, 0xdd, 0x82, 0x7e, 0x21, 0x16, 0x2f, 0x88, 0xe4,
	0xa6, 0x93, 0x85, 0x58, 0x3c, 0xc0, 0x8a, 0xbd, 0x05, 0xde, 0x3c, 0xc2, 0x78, 0xa6, 0x5c, 0xba,
	0x9d, 0xae, 0x32, 0x3c, 0xc0, 0x2a, 0xf8, 0xcd, 0x82, 0xde, 0x04, 0xe5, 0x05, 0xca, 0xd7, 0x12,
	0xb9, 0xf6, 0xe3, 0x61, 0xbf, 0xe2, 0xf1, 0xe8, 0xee, 0x7f, 0x3c, 0x9c, 0xcd, 0xe3, 0x71, 0x13,
	0x9c, 0x89, 0x0c, 0x4f, 0xc7, 0x2a, 0x23, 0x9b, 0x6b, 0x40, 0x6c, 0x1c, 0x85, 0x45, 0x74, 0x81,
	0xe6, 0x45, 0x31, 0x68, 0x47, 0xfb, 0xdc, 0x3d, 0xda, 0xf7, 0x83, 0x05, 0xbd, 0x87, 0xa2, 0x4a,
	0xcb, 0x62, 0x87, 0x85, 0x43, 0x18, 0x8c, 0xb2, 0x2c, 0x8e, 0x42, 0xfd, 0x6b, 0x7d, 0xa3, 0xb6,
	0x89, 0x76, 0x3c, 0x6a, 0xd5, 0x57, 0xdf, 0xad, 0x6d, 0x22, 0xb9, 0x38, 0x51, 0xfa, 0xae, 0xc5,
	0xba, 0x25, 0x17, 0x5a, 0xd6, 0x95, 0x93, 0x8a, 0x30, 0x2a, 0x8b, 0x74, 0x1e, 0xa7, 0x97, 0xea,
	0xb6, 0x2e, 0x6f, 0x70, 0xf0, 0x7b, 0x07, 0xba, 0xff, 0x95, 0x26, 0x1f, 0x80, 0x15, 0x99, 0x66,
	0x5b, 0x51, 0xa3, 0xd0, 0xfd, 0x96, 0x42, 0xfb, 0xd0, 0xaf, 0xa4, 0x48, 0x16, 0x98, 0xfb, 0xae,
	0x52, 0x97, 0x1a, 0x2a, 0x8f, 0x9a, 0x23, 0x2d, 0xcd, 0x1e, 0xaf, 0x61, 0x33, 0x17, 0xd0, 0x9a,
	0x8b, 0x8f, 0x8c, 0x8a, 0x0f, 0x54, 0x46, 0xfe, 0x76, 0x59, 0xae, 0x8a, 0xf7, 0xbf, 0xa7, 0xc9,
	0x7f, 0x59, 0xe0, 0x34, 0x43, 0x75, 0xb2, 0x3d, 0x54, 0x27, 0x9b, 0xa1, 0x1a, 0x1f, 0xd7, 0x43,
	0x35, 0x3e, 0x26, 0xcc, 0xcf, 0xea, 0xa1, 0xe2, 0x67, 0xd4, 0xac, 0xfb, 0x32, 0x2d, 0xb3, 0xe3,
	0x4a, 0x77, 0xd5, 0xe3, 0x0d, 0x26, 0x26, 0x7e, 0xb3, 0x44, 0x69, 0x4a, 0xed, 0x71, 0x83, 0x88,
	0xb7, 0x0f, 0x95, 0xe0, 0xe8, 0xe2, 0x6a, 0xc0, 0xde, 0x03, 0x87, 0x53, 0xf1, 0x54, 0x85, 0xb7,
	0xfa, 0xa2, 0xcc, 0x5c, 0x7b, 0x29, 0xa8, 0xfe, 0x7a, 0x33, 0x04, 0xae, 0xbf, 0xe5, 0x3e, 0x84,
	0xde, 0x64, 0x19, 0xcd, 0x8b, 0xfa, 0x2d, 0xfc, 0x7f, 0x4b, 0xb0, 0xa2, 0x15, 0x2a, 0x1f, 0x37,
	0x5b, 0x82, 0xa7, 0xe0, 0x35, 0xc6, 0x4d, 0x3a, 0x56, 0x3b, 0x1d, 0x06, 0xdd, 0xf3, 0x24, 0x2a,
	0xea, 0xd1, 0xa5, 0x35, 0x5d, 0xf6, 0x69, 0x29, 0x92, 0x22, 0x2a, 0xaa, 0x7a, 0x74, 0x6b, 0x1c,
	0xdc, 0x33, 0xe9, 0x53, 0xb8, 0xf3, 0x2c, 0x43, 0x69, 0x64, 0x40, 0x03, 0x75, 0x48, 0x7a, 0x89,
	0x5a, 0xc1, 0x6d, 0xae, 0x41, 0xf0, 0x2d, 0x78, 0xa3, 0x18, 0x65, 0xc1, 0xcb, 0x18, 0xf7, 0xbd,
	0x8c, 0x5f, 0x4f, 0x9e, 0x3c, 0xae, 0x33, 0xa0, 0xf5, 0x66, 0xe4, 0xed, 0x2b, 0x23, 0xff, 0x40,
	0x64, 0xe2, 0x74, 0xac, 0x78, 0x6e, 0x73, 0x83, 0x82, 0x9f, 0x2c, 0xe8, 0x92, 0xb6, 0xb4, 0x42,
	0x77, 0x5f, 0xa5, 0x4b, 0x67, 0x32, 0xbd, 0x88, 0x66, 0x28, 0xeb, 0xcb, 0xd5, 0x58, 0x15, 0x3d,
	0x5c, 0x62, 0xf3, 0x00, 0x1b, 0x44, 0x5c, 0xa3, 0x4f, 0xbd, 0x7a, 0x96, 0x5a, 0x5c, 0x23, 0x33,
	0xd7, 0x4e, 0xf6, 0x36, 0xc0, 0xa4, 0xcc, 0x50, 0x8e, 0x66, 0xab, 0x28, 0x51, 0x4d, 0x77, 0x79,
	0xcb, 0x12, 0x7c, 0xa1, 0x3f, 0x1e, 0x77, 0x14, 0xca, 0xda, 0xff, 0xa1, 0x79, 0x35, 0xf3, 0x20,
	0xde, 0xfe, 0xdd, 0x6b, 0xdd, 0x76, 0x08, 0x03, 0xf3, 0xa5, 0xad, 0xbe, 0x5b, 0x8d, 0x58, 0xb5,
	0x4c, 0x74, 0xe7, 0xb3, 0x72, 0x1a, 0x47, 0xa1, 0xba, 0xb3, 0xcb, 0x0d, 0x9a, 0xf6, 0xd4, 0x1f,
	0x8a, 0x7b, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xc4, 0x7a, 0x3e, 0x62, 0x0c, 0x00, 0x00,
}
