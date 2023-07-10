// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/amm/pool.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type Pool struct {
	PoolId            uint64                                 `protobuf:"varint,1,opt,name=poolId,proto3" json:"poolId,omitempty"`
	Address           string                                 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	PoolParams        PoolParams                             `protobuf:"bytes,3,opt,name=poolParams,proto3" json:"poolParams"`
	TotalShares       types.Coin                             `protobuf:"bytes,4,opt,name=totalShares,proto3" json:"totalShares"`
	PoolAssets        []PoolAsset                            `protobuf:"bytes,5,rep,name=poolAssets,proto3" json:"poolAssets"`
	TotalWeight       github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=totalWeight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"totalWeight"`
	RebalanceTreasury string                                 `protobuf:"bytes,7,opt,name=rebalanceTreasury,proto3" json:"rebalanceTreasury,omitempty"`
}

func (m *Pool) Reset()         { *m = Pool{} }
func (m *Pool) String() string { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()    {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ac3be9a215271f9, []int{0}
}
func (m *Pool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return m.Size()
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

func (m *Pool) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *Pool) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Pool) GetPoolParams() PoolParams {
	if m != nil {
		return m.PoolParams
	}
	return PoolParams{}
}

func (m *Pool) GetTotalShares() types.Coin {
	if m != nil {
		return m.TotalShares
	}
	return types.Coin{}
}

func (m *Pool) GetPoolAssets() []PoolAsset {
	if m != nil {
		return m.PoolAssets
	}
	return nil
}

func (m *Pool) GetRebalanceTreasury() string {
	if m != nil {
		return m.RebalanceTreasury
	}
	return ""
}

func init() {
	proto.RegisterType((*Pool)(nil), "elysnetwork.elys.amm.Pool")
}

func init() { proto.RegisterFile("elys/amm/pool.proto", fileDescriptor_3ac3be9a215271f9) }

var fileDescriptor_3ac3be9a215271f9 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xb1, 0xce, 0xd3, 0x30,
	0x18, 0x4c, 0x68, 0xe8, 0x2f, 0xdc, 0x09, 0x53, 0x21, 0xb7, 0x43, 0x1a, 0x31, 0xa0, 0x0c, 0xd4,
	0x56, 0xcb, 0x13, 0x34, 0x08, 0xa4, 0x6e, 0x51, 0x40, 0x42, 0x62, 0xa9, 0x9c, 0xc4, 0x4a, 0xa3,
	0x26, 0x71, 0x64, 0xbb, 0x40, 0x5f, 0x81, 0x89, 0xc7, 0xea, 0xd8, 0x11, 0x31, 0x54, 0xa8, 0x7d,
	0x11, 0x64, 0xc7, 0x45, 0xa9, 0x80, 0x29, 0xdf, 0xe7, 0xbb, 0xfb, 0x7c, 0xf7, 0xc5, 0xe0, 0x19,
	0xab, 0x0e, 0x92, 0xd0, 0xba, 0x26, 0x2d, 0xe7, 0x15, 0x6e, 0x05, 0x57, 0x1c, 0x8e, 0xf5, 0x61,
	0xc3, 0xd4, 0x17, 0x2e, 0x76, 0x58, 0xd7, 0x98, 0xd6, 0xf5, 0x74, 0x7a, 0x47, 0xdd, 0xb4, 0x54,
	0xd0, 0x5a, 0x76, 0x8a, 0xe9, 0xe4, 0x1e, 0xa3, 0x52, 0x32, 0x65, 0xa1, 0x71, 0xc1, 0x0b, 0x6e,
	0x4a, 0xa2, 0x2b, 0x7b, 0xea, 0x67, 0x5c, 0xd6, 0x5c, 0x92, 0x94, 0x4a, 0x46, 0x3e, 0x2f, 0x52,
	0xa6, 0xe8, 0x82, 0x64, 0xbc, 0x6c, 0x6e, 0x03, 0x3b, 0x7c, 0xd3, 0x09, 0xbb, 0xa6, 0x83, 0x5e,
	0x7c, 0x1b, 0x00, 0x2f, 0xe6, 0xbc, 0x82, 0xcf, 0xc1, 0x50, 0xdf, 0xb6, 0xce, 0x91, 0x1b, 0xb8,
	0xa1, 0x97, 0xd8, 0x0e, 0x22, 0xf0, 0x40, 0xf3, 0x5c, 0x30, 0x29, 0xd1, 0xa3, 0xc0, 0x0d, 0x9f,
	0x24, 0xb7, 0x16, 0xbe, 0x03, 0x40, 0x73, 0x62, 0x63, 0x1d, 0x0d, 0x02, 0x37, 0x1c, 0x2d, 0x03,
	0xfc, 0xaf, 0xb4, 0x38, 0xfe, 0xc3, 0x8b, 0xbc, 0xe3, 0x79, 0xe6, 0x24, 0x3d, 0x25, 0x5c, 0x81,
	0x91, 0xe2, 0x8a, 0x56, 0xef, 0xb7, 0x54, 0x30, 0x89, 0x3c, 0x33, 0x68, 0x82, 0xad, 0x4d, 0x9d,
	0x09, 0xdb, 0x4c, 0xf8, 0x0d, 0x2f, 0x1b, 0x3b, 0xa1, 0xaf, 0x81, 0x6f, 0x3b, 0x2b, 0x2b, 0xbd,
	0x29, 0x89, 0x1e, 0x07, 0x83, 0x70, 0xb4, 0x9c, 0xfd, 0xdf, 0x8a, 0xe1, 0xf5, 0x9d, 0x74, 0x42,
	0x18, 0x5b, 0x27, 0x1f, 0x59, 0x59, 0x6c, 0x15, 0x1a, 0xea, 0xbc, 0x11, 0xd6, 0xb4, 0x9f, 0xe7,
	0xd9, 0xcb, 0xa2, 0x54, 0xdb, 0x7d, 0x8a, 0x33, 0x5e, 0xdb, 0x15, 0xda, 0xcf, 0x5c, 0xe6, 0x3b,
	0xa2, 0x0e, 0x2d, 0x93, 0x78, 0xdd, 0xa8, 0xa4, 0x3f, 0x02, 0xbe, 0x02, 0x4f, 0x05, 0x4b, 0x69,
	0x45, 0x9b, 0x8c, 0x7d, 0x10, 0x8c, 0xca, 0xbd, 0x38, 0xa0, 0x07, 0xb3, 0xc7, 0xbf, 0x81, 0x28,
	0x3a, 0x5e, 0x7c, 0xf7, 0x74, 0xf1, 0xdd, 0x5f, 0x17, 0xdf, 0xfd, 0x7e, 0xf5, 0x9d, 0xd3, 0xd5,
	0x77, 0x7e, 0x5c, 0x7d, 0xe7, 0x53, 0xd8, 0xbb, 0x5c, 0x47, 0x99, 0xdb, 0x5c, 0xa6, 0x21, 0x5f,
	0xcd, 0x63, 0x31, 0x16, 0xd2, 0xa1, 0xf9, 0xaf, 0xaf, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x64,
	0x85, 0x5f, 0xc3, 0x8c, 0x02, 0x00, 0x00,
}

func (m *Pool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RebalanceTreasury) > 0 {
		i -= len(m.RebalanceTreasury)
		copy(dAtA[i:], m.RebalanceTreasury)
		i = encodeVarintPool(dAtA, i, uint64(len(m.RebalanceTreasury)))
		i--
		dAtA[i] = 0x3a
	}
	{
		size := m.TotalWeight.Size()
		i -= size
		if _, err := m.TotalWeight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.PoolAssets) > 0 {
		for iNdEx := len(m.PoolAssets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolAssets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPool(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	{
		size, err := m.TotalShares.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.PoolParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintPool(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.PoolId != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovPool(uint64(m.PoolId))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = m.PoolParams.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.TotalShares.Size()
	n += 1 + l + sovPool(uint64(l))
	if len(m.PoolAssets) > 0 {
		for _, e := range m.PoolAssets {
			l = e.Size()
			n += 1 + l + sovPool(uint64(l))
		}
	}
	l = m.TotalWeight.Size()
	n += 1 + l + sovPool(uint64(l))
	l = len(m.RebalanceTreasury)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	return n
}

func sovPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPool(x uint64) (n int) {
	return sovPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: Pool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PoolParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalShares", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalShares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolAssets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolAssets = append(m.PoolAssets, PoolAsset{})
			if err := m.PoolAssets[len(m.PoolAssets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalWeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalWeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RebalanceTreasury", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RebalanceTreasury = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPool
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
func skipPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
				return 0, ErrInvalidLengthPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPool = fmt.Errorf("proto: unexpected end of group")
)