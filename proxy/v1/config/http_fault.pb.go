// Code generated by protoc-gen-go.
// source: proxy/v1/config/http_fault.proto
// DO NOT EDIT!

package istio_proxy_v1_config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// HTTPFaultInjection can be used to specify one or more faults to inject
// while forwarding http requests to the destination specified in the route
// rule.  Fault specification is part of a route rule. Faults include
// aborting the Http request from downstream service, and/or delaying
// proxying of requests. A fault rule MUST HAVE delay or abort or both.
//
// *Note:* Delay and abort faults are independent of one another, even if
// both are specified simultaneously.
type HTTPFaultInjection struct {
	// Delay requests before forwarding, emulating various failures such as
	// network issues, overloaded upstream service, etc.
	Delay *HTTPFaultInjection_Delay `protobuf:"bytes,1,opt,name=delay" json:"delay,omitempty"`
	// Abort Http request attempts and return error codes back to downstream
	// service, giving the impression that the upstream service is faulty.
	Abort *HTTPFaultInjection_Abort `protobuf:"bytes,2,opt,name=abort" json:"abort,omitempty"`
}

func (m *HTTPFaultInjection) Reset()                    { *m = HTTPFaultInjection{} }
func (m *HTTPFaultInjection) String() string            { return proto.CompactTextString(m) }
func (*HTTPFaultInjection) ProtoMessage()               {}
func (*HTTPFaultInjection) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *HTTPFaultInjection) GetDelay() *HTTPFaultInjection_Delay {
	if m != nil {
		return m.Delay
	}
	return nil
}

func (m *HTTPFaultInjection) GetAbort() *HTTPFaultInjection_Abort {
	if m != nil {
		return m.Abort
	}
	return nil
}

// Delay specification is used to inject latency into the request
// forwarding path. The following example will introduce a 5 second delay
// in 10% of the requests to the "v1" version of the "reviews"
// service.
//
//     destination: reviews.default.svc.cluster.local
//     route:
//     - tags:
//         version: v1
//     httpFault:
//       delay:
//         percent: 10
//         fixedDelay: 5s
//
// The _fixedDelay_ field is used to indicate the amount of delay in
// seconds. An optional _percent_ field, a value between 0 and 100, can
// be used to only delay a certain percentage of requests. If left
// unspecified, all request will be delayed.
type HTTPFaultInjection_Delay struct {
	// percentage of requests on which the delay will be injected (0-100)
	Percent float32 `protobuf:"fixed32,1,opt,name=percent" json:"percent,omitempty"`
	// Types that are valid to be assigned to HttpDelayType:
	//	*HTTPFaultInjection_Delay_FixedDelay
	//	*HTTPFaultInjection_Delay_ExponentialDelay
	HttpDelayType isHTTPFaultInjection_Delay_HttpDelayType `protobuf_oneof:"http_delay_type"`
	// (-- Specify delay duration as part of Http request.
	// TODO: The semantics and syntax of the headers is undefined. --)
	OverrideHeaderName string `protobuf:"bytes,4,opt,name=override_header_name,json=overrideHeaderName" json:"override_header_name,omitempty"`
}

func (m *HTTPFaultInjection_Delay) Reset()                    { *m = HTTPFaultInjection_Delay{} }
func (m *HTTPFaultInjection_Delay) String() string            { return proto.CompactTextString(m) }
func (*HTTPFaultInjection_Delay) ProtoMessage()               {}
func (*HTTPFaultInjection_Delay) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 0} }

type isHTTPFaultInjection_Delay_HttpDelayType interface {
	isHTTPFaultInjection_Delay_HttpDelayType()
}

type HTTPFaultInjection_Delay_FixedDelay struct {
	FixedDelay *google_protobuf1.Duration `protobuf:"bytes,2,opt,name=fixed_delay,json=fixedDelay,oneof"`
}
type HTTPFaultInjection_Delay_ExponentialDelay struct {
	ExponentialDelay *google_protobuf1.Duration `protobuf:"bytes,3,opt,name=exponential_delay,json=exponentialDelay,oneof"`
}

func (*HTTPFaultInjection_Delay_FixedDelay) isHTTPFaultInjection_Delay_HttpDelayType()       {}
func (*HTTPFaultInjection_Delay_ExponentialDelay) isHTTPFaultInjection_Delay_HttpDelayType() {}

func (m *HTTPFaultInjection_Delay) GetHttpDelayType() isHTTPFaultInjection_Delay_HttpDelayType {
	if m != nil {
		return m.HttpDelayType
	}
	return nil
}

func (m *HTTPFaultInjection_Delay) GetPercent() float32 {
	if m != nil {
		return m.Percent
	}
	return 0
}

func (m *HTTPFaultInjection_Delay) GetFixedDelay() *google_protobuf1.Duration {
	if x, ok := m.GetHttpDelayType().(*HTTPFaultInjection_Delay_FixedDelay); ok {
		return x.FixedDelay
	}
	return nil
}

func (m *HTTPFaultInjection_Delay) GetExponentialDelay() *google_protobuf1.Duration {
	if x, ok := m.GetHttpDelayType().(*HTTPFaultInjection_Delay_ExponentialDelay); ok {
		return x.ExponentialDelay
	}
	return nil
}

func (m *HTTPFaultInjection_Delay) GetOverrideHeaderName() string {
	if m != nil {
		return m.OverrideHeaderName
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*HTTPFaultInjection_Delay) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _HTTPFaultInjection_Delay_OneofMarshaler, _HTTPFaultInjection_Delay_OneofUnmarshaler, _HTTPFaultInjection_Delay_OneofSizer, []interface{}{
		(*HTTPFaultInjection_Delay_FixedDelay)(nil),
		(*HTTPFaultInjection_Delay_ExponentialDelay)(nil),
	}
}

func _HTTPFaultInjection_Delay_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*HTTPFaultInjection_Delay)
	// http_delay_type
	switch x := m.HttpDelayType.(type) {
	case *HTTPFaultInjection_Delay_FixedDelay:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FixedDelay); err != nil {
			return err
		}
	case *HTTPFaultInjection_Delay_ExponentialDelay:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ExponentialDelay); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("HTTPFaultInjection_Delay.HttpDelayType has unexpected type %T", x)
	}
	return nil
}

func _HTTPFaultInjection_Delay_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*HTTPFaultInjection_Delay)
	switch tag {
	case 2: // http_delay_type.fixed_delay
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf1.Duration)
		err := b.DecodeMessage(msg)
		m.HttpDelayType = &HTTPFaultInjection_Delay_FixedDelay{msg}
		return true, err
	case 3: // http_delay_type.exponential_delay
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf1.Duration)
		err := b.DecodeMessage(msg)
		m.HttpDelayType = &HTTPFaultInjection_Delay_ExponentialDelay{msg}
		return true, err
	default:
		return false, nil
	}
}

func _HTTPFaultInjection_Delay_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*HTTPFaultInjection_Delay)
	// http_delay_type
	switch x := m.HttpDelayType.(type) {
	case *HTTPFaultInjection_Delay_FixedDelay:
		s := proto.Size(x.FixedDelay)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *HTTPFaultInjection_Delay_ExponentialDelay:
		s := proto.Size(x.ExponentialDelay)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Abort specification is used to prematurely abort a request with a
// pre-specified error code. The following example will return an HTTP
// 400 error code for 10% of the requests to the "ratings" service "v1".
//
//     destination: ratings.default.svc.cluster.local
//     route:
//     - tags:
//         version: v1
//     httpFault:
//       abort:
//         percent: 10
//         httpStatus: 400
//
// The _httpStatus_ field is used to indicate the HTTP status code to
// return to the caller. The optional _percent_ field, a value between 0
// and 100, is used to only abort a certain percentage of requests. If
// not specified, all requests are aborted.
type HTTPFaultInjection_Abort struct {
	// percentage of requests to be aborted with the error code provided (0-100).
	Percent float32 `protobuf:"fixed32,1,opt,name=percent" json:"percent,omitempty"`
	// Types that are valid to be assigned to ErrorType:
	//	*HTTPFaultInjection_Abort_GrpcStatus
	//	*HTTPFaultInjection_Abort_Http2Error
	//	*HTTPFaultInjection_Abort_HttpStatus
	ErrorType isHTTPFaultInjection_Abort_ErrorType `protobuf_oneof:"error_type"`
	// (-- Specify abort code as part of Http request.
	// TODO: The semantics and syntax of the headers is undefined. --)
	OverrideHeaderName string `protobuf:"bytes,5,opt,name=override_header_name,json=overrideHeaderName" json:"override_header_name,omitempty"`
}

func (m *HTTPFaultInjection_Abort) Reset()                    { *m = HTTPFaultInjection_Abort{} }
func (m *HTTPFaultInjection_Abort) String() string            { return proto.CompactTextString(m) }
func (*HTTPFaultInjection_Abort) ProtoMessage()               {}
func (*HTTPFaultInjection_Abort) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 1} }

type isHTTPFaultInjection_Abort_ErrorType interface {
	isHTTPFaultInjection_Abort_ErrorType()
}

type HTTPFaultInjection_Abort_GrpcStatus struct {
	GrpcStatus string `protobuf:"bytes,2,opt,name=grpc_status,json=grpcStatus,oneof"`
}
type HTTPFaultInjection_Abort_Http2Error struct {
	Http2Error string `protobuf:"bytes,3,opt,name=http2_error,json=http2Error,oneof"`
}
type HTTPFaultInjection_Abort_HttpStatus struct {
	HttpStatus int32 `protobuf:"varint,4,opt,name=http_status,json=httpStatus,oneof"`
}

func (*HTTPFaultInjection_Abort_GrpcStatus) isHTTPFaultInjection_Abort_ErrorType() {}
func (*HTTPFaultInjection_Abort_Http2Error) isHTTPFaultInjection_Abort_ErrorType() {}
func (*HTTPFaultInjection_Abort_HttpStatus) isHTTPFaultInjection_Abort_ErrorType() {}

func (m *HTTPFaultInjection_Abort) GetErrorType() isHTTPFaultInjection_Abort_ErrorType {
	if m != nil {
		return m.ErrorType
	}
	return nil
}

func (m *HTTPFaultInjection_Abort) GetPercent() float32 {
	if m != nil {
		return m.Percent
	}
	return 0
}

func (m *HTTPFaultInjection_Abort) GetGrpcStatus() string {
	if x, ok := m.GetErrorType().(*HTTPFaultInjection_Abort_GrpcStatus); ok {
		return x.GrpcStatus
	}
	return ""
}

func (m *HTTPFaultInjection_Abort) GetHttp2Error() string {
	if x, ok := m.GetErrorType().(*HTTPFaultInjection_Abort_Http2Error); ok {
		return x.Http2Error
	}
	return ""
}

func (m *HTTPFaultInjection_Abort) GetHttpStatus() int32 {
	if x, ok := m.GetErrorType().(*HTTPFaultInjection_Abort_HttpStatus); ok {
		return x.HttpStatus
	}
	return 0
}

func (m *HTTPFaultInjection_Abort) GetOverrideHeaderName() string {
	if m != nil {
		return m.OverrideHeaderName
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*HTTPFaultInjection_Abort) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _HTTPFaultInjection_Abort_OneofMarshaler, _HTTPFaultInjection_Abort_OneofUnmarshaler, _HTTPFaultInjection_Abort_OneofSizer, []interface{}{
		(*HTTPFaultInjection_Abort_GrpcStatus)(nil),
		(*HTTPFaultInjection_Abort_Http2Error)(nil),
		(*HTTPFaultInjection_Abort_HttpStatus)(nil),
	}
}

func _HTTPFaultInjection_Abort_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*HTTPFaultInjection_Abort)
	// error_type
	switch x := m.ErrorType.(type) {
	case *HTTPFaultInjection_Abort_GrpcStatus:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.GrpcStatus)
	case *HTTPFaultInjection_Abort_Http2Error:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Http2Error)
	case *HTTPFaultInjection_Abort_HttpStatus:
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.HttpStatus))
	case nil:
	default:
		return fmt.Errorf("HTTPFaultInjection_Abort.ErrorType has unexpected type %T", x)
	}
	return nil
}

func _HTTPFaultInjection_Abort_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*HTTPFaultInjection_Abort)
	switch tag {
	case 2: // error_type.grpc_status
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.ErrorType = &HTTPFaultInjection_Abort_GrpcStatus{x}
		return true, err
	case 3: // error_type.http2_error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.ErrorType = &HTTPFaultInjection_Abort_Http2Error{x}
		return true, err
	case 4: // error_type.http_status
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ErrorType = &HTTPFaultInjection_Abort_HttpStatus{int32(x)}
		return true, err
	default:
		return false, nil
	}
}

func _HTTPFaultInjection_Abort_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*HTTPFaultInjection_Abort)
	// error_type
	switch x := m.ErrorType.(type) {
	case *HTTPFaultInjection_Abort_GrpcStatus:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.GrpcStatus)))
		n += len(x.GrpcStatus)
	case *HTTPFaultInjection_Abort_Http2Error:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Http2Error)))
		n += len(x.Http2Error)
	case *HTTPFaultInjection_Abort_HttpStatus:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.HttpStatus))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*HTTPFaultInjection)(nil), "istio.proxy.v1.config.HTTPFaultInjection")
	proto.RegisterType((*HTTPFaultInjection_Delay)(nil), "istio.proxy.v1.config.HTTPFaultInjection.Delay")
	proto.RegisterType((*HTTPFaultInjection_Abort)(nil), "istio.proxy.v1.config.HTTPFaultInjection.Abort")
}

func init() { proto.RegisterFile("proxy/v1/config/http_fault.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0x6f, 0xb9, 0xf4, 0xde, 0x70, 0x30, 0x51, 0x1a, 0x4d, 0x2a, 0x0b, 0x83, 0xae, 0x58,
	0x4d, 0x05, 0xb7, 0x6e, 0x34, 0x60, 0xea, 0xc6, 0x98, 0xca, 0xbe, 0x19, 0xda, 0xd3, 0x52, 0x53,
	0x3a, 0x93, 0x61, 0x4a, 0xe0, 0x15, 0x7d, 0x0a, 0x9f, 0xc0, 0x67, 0x30, 0x73, 0x86, 0x26, 0x26,
	0x8a, 0xc6, 0x65, 0xcf, 0xfc, 0xff, 0x37, 0xe7, 0x9b, 0xc2, 0x40, 0x2a, 0xb1, 0xd9, 0x06, 0xeb,
	0x51, 0x90, 0x88, 0x2a, 0x2b, 0xf2, 0x60, 0xa1, 0xb5, 0x8c, 0x33, 0x5e, 0x97, 0x9a, 0x49, 0x25,
	0xb4, 0xf0, 0x4e, 0x8a, 0x95, 0x2e, 0x04, 0xa3, 0x1c, 0x5b, 0x8f, 0x98, 0xcd, 0xf5, 0xcf, 0x72,
	0x21, 0xf2, 0x12, 0x03, 0x0a, 0xcd, 0xeb, 0x2c, 0x48, 0x6b, 0xc5, 0x75, 0x21, 0x2a, 0x5b, 0xbb,
	0x78, 0x6d, 0x83, 0x17, 0xce, 0x66, 0x8f, 0x77, 0x06, 0x75, 0x5f, 0x3d, 0x63, 0x62, 0x0e, 0xbd,
	0x29, 0xb8, 0x29, 0x96, 0x7c, 0xeb, 0x3b, 0x03, 0x67, 0xd8, 0x1d, 0x07, 0xec, 0x4b, 0x3a, 0xfb,
	0xdc, 0x64, 0x13, 0x53, 0x8b, 0x6c, 0xdb, 0x60, 0xf8, 0x5c, 0x28, 0xed, 0xb7, 0x7e, 0x8b, 0xb9,
	0x31, 0xb5, 0xc8, 0xb6, 0xfb, 0x6f, 0x0e, 0xb8, 0xc4, 0xf5, 0x7c, 0xf8, 0x2f, 0x51, 0x25, 0x58,
	0x69, 0xda, 0xac, 0x15, 0x35, 0x9f, 0xde, 0x35, 0x74, 0xb3, 0x62, 0x83, 0x69, 0x6c, 0xf7, 0xb6,
	0x17, 0x9e, 0x32, 0xab, 0xcf, 0x1a, 0x7d, 0x36, 0xd9, 0xe9, 0x87, 0x7f, 0x22, 0xa0, 0xbc, 0xe5,
	0x86, 0xd0, 0xc3, 0x8d, 0x14, 0x15, 0x56, 0xba, 0xe0, 0xe5, 0x8e, 0xf1, 0xf7, 0x67, 0xc6, 0xd1,
	0x87, 0x96, 0x25, 0x5d, 0xc2, 0xb1, 0x58, 0xa3, 0x52, 0x45, 0x8a, 0xf1, 0x02, 0x79, 0x8a, 0x2a,
	0xae, 0xf8, 0x12, 0xfd, 0xf6, 0xc0, 0x19, 0x76, 0x22, 0xaf, 0x39, 0x0b, 0xe9, 0xe8, 0x81, 0x2f,
	0xf1, 0xb6, 0x07, 0x87, 0xf4, 0x37, 0xe9, 0xd2, 0x58, 0x6f, 0x25, 0xf6, 0x5f, 0x1c, 0x70, 0xe9,
	0x05, 0xbe, 0x11, 0x3e, 0x87, 0x6e, 0xae, 0x64, 0x12, 0xaf, 0x34, 0xd7, 0xf5, 0x8a, 0x84, 0x3b,
	0xc6, 0xca, 0x0c, 0x9f, 0x68, 0x66, 0x22, 0x86, 0x3c, 0x8e, 0x51, 0x29, 0xa1, 0xc8, 0x87, 0x22,
	0x34, 0x9c, 0x9a, 0x59, 0x13, 0x69, 0x28, 0x66, 0x4b, 0xb7, 0x89, 0xec, 0x28, 0xfb, 0x8c, 0xdc,
	0xbd, 0x46, 0x07, 0x00, 0x74, 0x23, 0xc9, 0xcc, 0xff, 0xd1, 0xc3, 0x5d, 0xbd, 0x07, 0x00, 0x00,
	0xff, 0xff, 0xdb, 0x48, 0xde, 0x73, 0xc4, 0x02, 0x00, 0x00,
}
