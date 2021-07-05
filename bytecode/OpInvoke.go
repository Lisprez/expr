// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package bytecode

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type OpInvoke struct {
	_tab flatbuffers.Table
}

func GetRootAsOpInvoke(buf []byte, offset flatbuffers.UOffsetT) *OpInvoke {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &OpInvoke{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *OpInvoke) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *OpInvoke) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *OpInvoke) Args() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *OpInvoke) MutateArgs(n uint16) bool {
	return rcv._tab.MutateUint16Slot(4, n)
}

func (rcv *OpInvoke) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func OpInvokeStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func OpInvokeAddArgs(builder *flatbuffers.Builder, args uint16) {
	builder.PrependUint16Slot(0, args, 0)
}
func OpInvokeAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(name), 0)
}
func OpInvokeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}