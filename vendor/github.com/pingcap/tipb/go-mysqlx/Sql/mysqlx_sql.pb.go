// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/pingcap/tipb/go-mysqlx/Sql/mysqlx_sql.proto

/*
	Package Mysqlx_Sql is a generated protocol buffer package.

	Messages of the MySQL Package

	It is generated from these files:
		github.com/pingcap/tipb/go-mysqlx/Sql/mysqlx_sql.proto

	It has these top-level messages:
		StmtExecute
		StmtExecuteOk
*/
package Mysqlx_Sql

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	Mysqlx_Datatypes "github.com/pingcap/tipb/go-mysqlx/Datatypes"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// execute a statement in the given namespace
//
// .. uml::
//
//   client -> server: StmtExecute
//   ... zero or more Resultsets ...
//   server --> client: StmtExecuteOk
//
// Notices:
//   This message may generate a notice containing WARNINGs generated by its execution.
//   This message may generate a notice containing INFO messages generated by its execution.
//
// :param namespace: namespace of the statement to be executed
// :param stmt: statement that shall be executed.
// :param args: values for wildcard replacements
// :param compact_metadata: send only type information for :protobuf:msg:`Mysqlx.Resultset::ColumnMetadata`, skipping names and others
// :returns:
//    * zero or one :protobuf:msg:`Mysqlx.Resultset::` followed by :protobuf:msg:`Mysqlx.Sql::StmtExecuteOk`
type StmtExecute struct {
	Namespace        *string                 `protobuf:"bytes,3,opt,name=namespace,def=sql" json:"namespace,omitempty"`
	Stmt             []byte                  `protobuf:"bytes,1,req,name=stmt" json:"stmt,omitempty"`
	Args             []*Mysqlx_Datatypes.Any `protobuf:"bytes,2,rep,name=args" json:"args,omitempty"`
	CompactMetadata  *bool                   `protobuf:"varint,4,opt,name=compact_metadata,json=compactMetadata,def=0" json:"compact_metadata,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *StmtExecute) Reset()                    { *m = StmtExecute{} }
func (m *StmtExecute) String() string            { return proto.CompactTextString(m) }
func (*StmtExecute) ProtoMessage()               {}
func (*StmtExecute) Descriptor() ([]byte, []int) { return fileDescriptorMysqlxSql, []int{0} }

const Default_StmtExecute_Namespace string = "sql"
const Default_StmtExecute_CompactMetadata bool = false

func (m *StmtExecute) GetNamespace() string {
	if m != nil && m.Namespace != nil {
		return *m.Namespace
	}
	return Default_StmtExecute_Namespace
}

func (m *StmtExecute) GetStmt() []byte {
	if m != nil {
		return m.Stmt
	}
	return nil
}

func (m *StmtExecute) GetArgs() []*Mysqlx_Datatypes.Any {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *StmtExecute) GetCompactMetadata() bool {
	if m != nil && m.CompactMetadata != nil {
		return *m.CompactMetadata
	}
	return Default_StmtExecute_CompactMetadata
}

// statement executed successful
type StmtExecuteOk struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *StmtExecuteOk) Reset()                    { *m = StmtExecuteOk{} }
func (m *StmtExecuteOk) String() string            { return proto.CompactTextString(m) }
func (*StmtExecuteOk) ProtoMessage()               {}
func (*StmtExecuteOk) Descriptor() ([]byte, []int) { return fileDescriptorMysqlxSql, []int{1} }

func init() {
	proto.RegisterType((*StmtExecute)(nil), "Mysqlx.Sql.StmtExecute")
	proto.RegisterType((*StmtExecuteOk)(nil), "Mysqlx.Sql.StmtExecuteOk")
}
func (m *StmtExecute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StmtExecute) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Stmt == nil {
		return 0, new(proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMysqlxSql(dAtA, i, uint64(len(m.Stmt)))
		i += copy(dAtA[i:], m.Stmt)
	}
	if len(m.Args) > 0 {
		for _, msg := range m.Args {
			dAtA[i] = 0x12
			i++
			i = encodeVarintMysqlxSql(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Namespace != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMysqlxSql(dAtA, i, uint64(len(*m.Namespace)))
		i += copy(dAtA[i:], *m.Namespace)
	}
	if m.CompactMetadata != nil {
		dAtA[i] = 0x20
		i++
		if *m.CompactMetadata {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *StmtExecuteOk) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StmtExecuteOk) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintMysqlxSql(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *StmtExecute) Size() (n int) {
	var l int
	_ = l
	if m.Stmt != nil {
		l = len(m.Stmt)
		n += 1 + l + sovMysqlxSql(uint64(l))
	}
	if len(m.Args) > 0 {
		for _, e := range m.Args {
			l = e.Size()
			n += 1 + l + sovMysqlxSql(uint64(l))
		}
	}
	if m.Namespace != nil {
		l = len(*m.Namespace)
		n += 1 + l + sovMysqlxSql(uint64(l))
	}
	if m.CompactMetadata != nil {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StmtExecuteOk) Size() (n int) {
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMysqlxSql(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMysqlxSql(x uint64) (n int) {
	return sovMysqlxSql(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StmtExecute) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMysqlxSql
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StmtExecute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StmtExecute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stmt", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlxSql
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMysqlxSql
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stmt = append(m.Stmt[:0], dAtA[iNdEx:postIndex]...)
			if m.Stmt == nil {
				m.Stmt = []byte{}
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Args", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlxSql
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMysqlxSql
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Args = append(m.Args, &Mysqlx_Datatypes.Any{})
			if err := m.Args[len(m.Args)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlxSql
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMysqlxSql
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Namespace = &s
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompactMetadata", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlxSql
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.CompactMetadata = &b
		default:
			iNdEx = preIndex
			skippy, err := skipMysqlxSql(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMysqlxSql
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StmtExecuteOk) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMysqlxSql
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StmtExecuteOk: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StmtExecuteOk: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMysqlxSql(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMysqlxSql
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMysqlxSql(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMysqlxSql
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMysqlxSql
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMysqlxSql
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMysqlxSql
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMysqlxSql
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMysqlxSql(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMysqlxSql = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMysqlxSql   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/pingcap/tipb/go-mysqlx/Sql/mysqlx_sql.proto", fileDescriptorMysqlxSql)
}

var fileDescriptorMysqlxSql = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xf6, 0xad, 0x2c, 0x2e,
	0xcc, 0xa9, 0xd0, 0x0f, 0x2e, 0xcc, 0xd1, 0xcf, 0x05, 0x33, 0xe3, 0x8b, 0x0b, 0x73, 0xf4, 0x0a,
	0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xb8, 0x20, 0x92, 0x7a, 0xc1, 0x85, 0x39, 0x52, 0xea, 0x50, 0x85,
	0x2e, 0x89, 0x25, 0x89, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x30, 0xe5, 0x29, 0x30, 0x01, 0x88, 0x26,
	0xa5, 0x85, 0x8c, 0x5c, 0xdc, 0xc1, 0x25, 0xb9, 0x25, 0xae, 0x15, 0xa9, 0xc9, 0xa5, 0x25, 0xa9,
	0x42, 0x42, 0x5c, 0x2c, 0xc5, 0x25, 0xb9, 0x25, 0x12, 0x8c, 0x0a, 0x4c, 0x1a, 0x3c, 0x41, 0x60,
	0xb6, 0x90, 0x26, 0x17, 0x4b, 0x62, 0x51, 0x7a, 0xb1, 0x04, 0x93, 0x02, 0xb3, 0x06, 0xb7, 0x91,
	0xa8, 0x1e, 0xd4, 0x1e, 0xb8, 0xd9, 0x7a, 0x8e, 0x79, 0x95, 0x41, 0x60, 0x25, 0x42, 0x8a, 0x5c,
	0x9c, 0x79, 0x89, 0xb9, 0xa9, 0xc5, 0x05, 0x89, 0xc9, 0xa9, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0x9c,
	0x56, 0xcc, 0xc5, 0x85, 0x39, 0x41, 0x08, 0x51, 0x21, 0x03, 0x2e, 0x81, 0xe4, 0xfc, 0xdc, 0x82,
	0xc4, 0xe4, 0x92, 0xf8, 0xdc, 0xd4, 0x92, 0x44, 0x90, 0x83, 0x24, 0x58, 0x14, 0x18, 0x35, 0x38,
	0xac, 0x58, 0xd3, 0x12, 0x73, 0x8a, 0x53, 0x83, 0xf8, 0xa1, 0xd2, 0xbe, 0x50, 0x59, 0x25, 0x7e,
	0x2e, 0x5e, 0x24, 0x27, 0xfa, 0x67, 0x3b, 0xe9, 0x9d, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c,
	0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x70, 0xc9, 0x24, 0xe7, 0xe7, 0xea, 0x81,
	0xbd, 0xa8, 0x97, 0x9c, 0x05, 0x61, 0x54, 0x40, 0x7c, 0x98, 0x54, 0x9a, 0x06, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x69, 0x0b, 0x10, 0xa9, 0x37, 0x01, 0x00, 0x00,
}
