// Code generated by protoc-gen-gogo.
// source: hello/hello.proto
// DO NOT EDIT!

/*
Package app_service_hello is a generated protocol buffer package.

It is generated from these files:
	hello/hello.proto

It has these top-level messages:
	Request
	Response
*/
package app_service_hello

import proto "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto/gogo.pb"

import io "io"
import fmt "fmt"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Request struct {
	Name             string `protobuf:"bytes,1,opt,name=name" json:"name"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Message          string `protobuf:"bytes,1,opt,name=message" json:"message"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
}
func (m *Request) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[index:postIndex])
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := github_com_gogo_protobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *Response) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(data[index:postIndex])
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := github_com_gogo_protobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *Request) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovHello(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Response) Size() (n int) {
	var l int
	_ = l
	l = len(m.Message)
	n += 1 + l + sovHello(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovHello(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHello(x uint64) (n int) {
	return sovHello(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Request) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintHello(data, i, uint64(len(m.Name)))
	i += copy(data[i:], m.Name)
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Response) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Response) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintHello(data, i, uint64(len(m.Message)))
	i += copy(data[i:], m.Message)
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Hello(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Hello(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintHello(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
