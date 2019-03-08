// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type RSC_DataF struct {
	_tab flatbuffers.Table
}

func GetRootAsRSC_DataF(buf []byte, offset flatbuffers.UOffsetT) *RSC_DataF {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &RSC_DataF{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *RSC_DataF) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *RSC_DataF) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *RSC_DataF) Tracts(obj *RSC_TractF, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 20
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *RSC_DataF) TractsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func RSC_DataFStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func RSC_DataFAddTracts(builder *flatbuffers.Builder, tracts flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(tracts), 0)
}
func RSC_DataFStartTractsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(20, numElems, 4)
}
func RSC_DataFEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
