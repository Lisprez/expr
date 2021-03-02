// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internal/ast/ast.proto

package ast

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Node_Token int32

const (
	Node_EXIT Node_Token = 0
	Node_CALL Node_Token = 1
	Node_STR  Node_Token = 2
	Node_INT  Node_Token = 3
	Node_BOOL Node_Token = 4
	Node_ARR  Node_Token = 5
)

var Node_Token_name = map[int32]string{
	0: "EXIT",
	1: "CALL",
	2: "STR",
	3: "INT",
	4: "BOOL",
	5: "ARR",
}

var Node_Token_value = map[string]int32{
	"EXIT": 0,
	"CALL": 1,
	"STR":  2,
	"INT":  3,
	"BOOL": 4,
	"ARR":  5,
}

func (x Node_Token) String() string {
	return proto.EnumName(Node_Token_name, int32(x))
}

func (Node_Token) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8ad3fb669b35941e, []int{0, 0}
}

type Node struct {
	Token Node_Token `protobuf:"varint,1,opt,name=token,proto3,enum=ast.Node_Token" json:"token,omitempty"`
	// Types that are valid to be assigned to Data:
	//	*Node_S
	//	*Node_I
	//	*Node_B
	Data   isNode_Data `protobuf_oneof:"data"`
	Nested []*Node     `protobuf:"bytes,5,rep,name=nested,proto3" json:"nested,omitempty"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ad3fb669b35941e, []int{0}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Node.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return m.Size()
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

type isNode_Data interface {
	isNode_Data()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Node_S struct {
	S string `protobuf:"bytes,2,opt,name=s,proto3,oneof" json:"s,omitempty"`
}
type Node_I struct {
	I int64 `protobuf:"varint,3,opt,name=i,proto3,oneof" json:"i,omitempty"`
}
type Node_B struct {
	B bool `protobuf:"varint,4,opt,name=b,proto3,oneof" json:"b,omitempty"`
}

func (*Node_S) isNode_Data() {}
func (*Node_I) isNode_Data() {}
func (*Node_B) isNode_Data() {}

func (m *Node) GetData() isNode_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Node) GetToken() Node_Token {
	if m != nil {
		return m.Token
	}
	return Node_EXIT
}

func (m *Node) GetS() string {
	if x, ok := m.GetData().(*Node_S); ok {
		return x.S
	}
	return ""
}

func (m *Node) GetI() int64 {
	if x, ok := m.GetData().(*Node_I); ok {
		return x.I
	}
	return 0
}

func (m *Node) GetB() bool {
	if x, ok := m.GetData().(*Node_B); ok {
		return x.B
	}
	return false
}

func (m *Node) GetNested() []*Node {
	if m != nil {
		return m.Nested
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Node) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Node_S)(nil),
		(*Node_I)(nil),
		(*Node_B)(nil),
	}
}

func init() {
	proto.RegisterEnum("ast.Node_Token", Node_Token_name, Node_Token_value)
	proto.RegisterType((*Node)(nil), "ast.Node")
}

func init() { proto.RegisterFile("internal/ast/ast.proto", fileDescriptor_8ad3fb669b35941e) }

var fileDescriptor_8ad3fb669b35941e = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xf3, 0x96, 0xb4, 0x6e, 0x11, 0x66, 0xc8, 0x41, 0x72, 0x0a, 0x71, 0x20, 0xe4, 0x54,
	0x61, 0x7e, 0x01, 0x57, 0x11, 0x1c, 0x94, 0x0d, 0x62, 0x0f, 0x5e, 0x53, 0x9a, 0x43, 0x51, 0x5a,
	0x59, 0xf2, 0x41, 0xfc, 0x58, 0xe2, 0x69, 0x47, 0x8f, 0xd2, 0x7e, 0x11, 0x79, 0x53, 0x3c, 0x04,
	0xfe, 0xbf, 0xff, 0x2f, 0xbc, 0xc7, 0xe3, 0x97, 0x5d, 0x9f, 0xc2, 0xa1, 0xf7, 0xaf, 0x37, 0x3e,
	0x26, 0x7c, 0xc5, 0xdb, 0x61, 0x48, 0x83, 0xa4, 0x3e, 0xa6, 0xd5, 0x27, 0x70, 0xb6, 0x1b, 0xda,
	0x20, 0xaf, 0x79, 0x96, 0x86, 0x97, 0xd0, 0x2b, 0x30, 0x60, 0x97, 0xeb, 0x8b, 0x02, 0x3f, 0xa2,
	0x29, 0x6a, 0xac, 0xdd, 0xaf, 0x95, 0x4b, 0x0e, 0x51, 0xcd, 0x0c, 0xd8, 0xc5, 0x23, 0x71, 0x10,
	0x91, 0x3b, 0x45, 0x0d, 0x58, 0x8a, 0xdc, 0x21, 0x37, 0x8a, 0x19, 0xb0, 0x73, 0xe4, 0x46, 0x5e,
	0xf1, 0xbc, 0x0f, 0x31, 0x85, 0x56, 0x65, 0x86, 0xda, 0xf3, 0xf5, 0xe2, 0x7f, 0xae, 0xfb, 0x13,
	0xab, 0x3b, 0x9e, 0x9d, 0x56, 0xc8, 0x39, 0x67, 0x0f, 0xcf, 0xdb, 0x5a, 0x10, 0x4c, 0xf7, 0x9b,
	0xaa, 0x12, 0x20, 0xcf, 0x38, 0x7d, 0xaa, 0x9d, 0x98, 0x61, 0xd8, 0xee, 0x6a, 0x41, 0xd1, 0x95,
	0xfb, 0x7d, 0x25, 0x18, 0x56, 0x1b, 0xe7, 0x44, 0x56, 0xe6, 0x9c, 0xb5, 0x3e, 0xf9, 0x52, 0x7d,
	0x8c, 0x1a, 0x8e, 0xa3, 0x86, 0xef, 0x51, 0xc3, 0xfb, 0xa4, 0xc9, 0x71, 0xd2, 0xe4, 0x6b, 0xd2,
	0xa4, 0xc9, 0x4f, 0x27, 0xdf, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x14, 0x7a, 0xda, 0x0c,
	0x01, 0x00, 0x00,
}

func (m *Node) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Node) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Node) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Nested) > 0 {
		for iNdEx := len(m.Nested) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Nested[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAst(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Data != nil {
		{
			size := m.Data.Size()
			i -= size
			if _, err := m.Data.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.Token != 0 {
		i = encodeVarintAst(dAtA, i, uint64(m.Token))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Node_S) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Node_S) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.S)
	copy(dAtA[i:], m.S)
	i = encodeVarintAst(dAtA, i, uint64(len(m.S)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func (m *Node_I) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Node_I) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = encodeVarintAst(dAtA, i, uint64(m.I))
	i--
	dAtA[i] = 0x18
	return len(dAtA) - i, nil
}
func (m *Node_B) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Node_B) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i--
	if m.B {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x20
	return len(dAtA) - i, nil
}
func encodeVarintAst(dAtA []byte, offset int, v uint64) int {
	offset -= sovAst(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Node) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Token != 0 {
		n += 1 + sovAst(uint64(m.Token))
	}
	if m.Data != nil {
		n += m.Data.Size()
	}
	if len(m.Nested) > 0 {
		for _, e := range m.Nested {
			l = e.Size()
			n += 1 + l + sovAst(uint64(l))
		}
	}
	return n
}

func (m *Node_S) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.S)
	n += 1 + l + sovAst(uint64(l))
	return n
}
func (m *Node_I) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovAst(uint64(m.I))
	return n
}
func (m *Node_B) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 2
	return n
}

func sovAst(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAst(x uint64) (n int) {
	return sovAst(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Node) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAst
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Node: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Node: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			m.Token = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAst
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Token |= Node_Token(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field S", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAst
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAst
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAst
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = &Node_S{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field I", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAst
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Data = &Node_I{v}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAst
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Data = &Node_B{b}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nested", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAst
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAst
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAst
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nested = append(m.Nested, &Node{})
			if err := m.Nested[len(m.Nested)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAst(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAst
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAst(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAst
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
					return 0, ErrIntOverflowAst
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAst
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
			if length < 0 {
				return 0, ErrInvalidLengthAst
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAst
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAst
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAst        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAst          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAst = fmt.Errorf("proto: unexpected end of group")
)
