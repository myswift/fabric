// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/configuration.proto

package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// AnchorPeers simply represents list of anchor peers which is used in ConfigurationItem
type AnchorPeers struct {
	AnchorPeers []*AnchorPeer `protobuf:"bytes,1,rep,name=anchor_peers,json=anchorPeers" json:"anchor_peers,omitempty"`
}

func (m *AnchorPeers) Reset()                    { *m = AnchorPeers{} }
func (m *AnchorPeers) String() string            { return proto.CompactTextString(m) }
func (*AnchorPeers) ProtoMessage()               {}
func (*AnchorPeers) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *AnchorPeers) GetAnchorPeers() []*AnchorPeer {
	if m != nil {
		return m.AnchorPeers
	}
	return nil
}

// AnchorPeer message structure which provides information about anchor peer, it includes host name,
// port number and peer certificate.
type AnchorPeer struct {
	// DNS host name of the anchor peer
	Host string `protobuf:"bytes,1,opt,name=host" json:"host,omitempty"`
	// The port number
	Port int32 `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
}

func (m *AnchorPeer) Reset()                    { *m = AnchorPeer{} }
func (m *AnchorPeer) String() string            { return proto.CompactTextString(m) }
func (*AnchorPeer) ProtoMessage()               {}
func (*AnchorPeer) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *AnchorPeer) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *AnchorPeer) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*AnchorPeers)(nil), "protos.AnchorPeers")
	proto.RegisterType((*AnchorPeer)(nil), "protos.AnchorPeer")
}

func init() { proto.RegisterFile("peer/configuration.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x8f, 0x3d, 0x8f, 0xc3, 0x20,
	0x0c, 0x86, 0xc5, 0x7d, 0x49, 0x47, 0x6e, 0x62, 0x62, 0x8c, 0x32, 0xe5, 0x16, 0x90, 0xee, 0xe3,
	0x07, 0xb4, 0xea, 0xde, 0x2a, 0x63, 0x97, 0x8a, 0x50, 0x02, 0x48, 0x6d, 0x8c, 0x0c, 0x19, 0xfa,
	0xef, 0x2b, 0x40, 0x55, 0x3a, 0xf1, 0x62, 0x3f, 0x8f, 0x6c, 0x53, 0x1e, 0x8c, 0x41, 0xa9, 0x61,
	0x9e, 0xbc, 0x5d, 0x50, 0x25, 0x0f, 0xb3, 0x08, 0x08, 0x09, 0xd8, 0x47, 0x79, 0x62, 0xb7, 0xa3,
	0xcd, 0x66, 0xd6, 0x0e, 0xf0, 0x60, 0x0c, 0x46, 0xf6, 0x4f, 0xbf, 0x54, 0xf9, 0x9e, 0xb2, 0x19,
	0x39, 0x69, 0x5f, 0xfb, 0xe6, 0x87, 0x55, 0x29, 0x8a, 0x15, 0x1d, 0x1a, 0xb5, 0x6a, 0xdd, 0x1f,
	0xa5, 0x6b, 0x8b, 0x31, 0xfa, 0xe6, 0x20, 0x26, 0x4e, 0x5a, 0xd2, 0x7f, 0x0e, 0x25, 0xe7, 0x5a,
	0x00, 0x4c, 0xfc, 0xa5, 0x25, 0xfd, 0xfb, 0x50, 0xf2, 0x76, 0x4f, 0x3b, 0x40, 0x2b, 0xdc, 0x2d,
	0x18, 0xbc, 0x98, 0xb3, 0x35, 0x28, 0x26, 0x35, 0xa2, 0xd7, 0x8f, 0x71, 0x79, 0x87, 0xe3, 0xb7,
	0xf5, 0xc9, 0x2d, 0xa3, 0xd0, 0x70, 0x95, 0x4f, 0xa8, 0xac, 0xa8, 0xac, 0xa8, 0xcc, 0xe8, 0x58,
	0x8f, 0xfa, 0xbd, 0x07, 0x00, 0x00, 0xff, 0xff, 0x6d, 0xf2, 0x7b, 0x6f, 0xf7, 0x00, 0x00, 0x00,
}
