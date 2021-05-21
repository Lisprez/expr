// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package bytecode

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type OpPushVector struct {
	_tab flatbuffers.Table
}

func GetRootAsOpPushVector(buf []byte, offset flatbuffers.UOffsetT) *OpPushVector {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &OpPushVector{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *OpPushVector) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *OpPushVector) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *OpPushVector) Elems() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *OpPushVector) MutateElems(n uint16) bool {
	return rcv._tab.MutateUint16Slot(4, n)
}

func OpPushVectorStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func OpPushVectorAddElems(builder *flatbuffers.Builder, elems uint16) {
	builder.PrependUint16Slot(0, elems, 0)
}
func OpPushVectorEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}