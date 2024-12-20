// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/masterchef/incentive.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// Incentive Info
type LegacyIncentiveInfo struct {
	// reward amount in eden for 1 year
	EdenAmountPerYear cosmossdk_io_math.Int `protobuf:"bytes,1,opt,name=eden_amount_per_year,json=edenAmountPerYear,proto3,customtype=cosmossdk.io/math.Int" json:"eden_amount_per_year"`
	// starting block height of the distribution
	DistributionStartBlock cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=distribution_start_block,json=distributionStartBlock,proto3,customtype=cosmossdk.io/math.Int" json:"distribution_start_block"`
	// distribution duration - block number per year
	TotalBlocksPerYear cosmossdk_io_math.Int `protobuf:"bytes,3,opt,name=total_blocks_per_year,json=totalBlocksPerYear,proto3,customtype=cosmossdk.io/math.Int" json:"total_blocks_per_year"`
	// blocks distributed
	BlocksDistributed cosmossdk_io_math.Int `protobuf:"bytes,4,opt,name=blocks_distributed,json=blocksDistributed,proto3,customtype=cosmossdk.io/math.Int" json:"blocks_distributed"`
}

func (m *LegacyIncentiveInfo) Reset()         { *m = LegacyIncentiveInfo{} }
func (m *LegacyIncentiveInfo) String() string { return proto.CompactTextString(m) }
func (*LegacyIncentiveInfo) ProtoMessage()    {}
func (*LegacyIncentiveInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbfbd92cca5ccf94, []int{0}
}
func (m *LegacyIncentiveInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LegacyIncentiveInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LegacyIncentiveInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LegacyIncentiveInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LegacyIncentiveInfo.Merge(m, src)
}
func (m *LegacyIncentiveInfo) XXX_Size() int {
	return m.Size()
}
func (m *LegacyIncentiveInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LegacyIncentiveInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LegacyIncentiveInfo proto.InternalMessageInfo

type IncentiveInfo struct {
	// reward amount in eden for 1 year
	EdenAmountPerYear cosmossdk_io_math.Int `protobuf:"bytes,1,opt,name=eden_amount_per_year,json=edenAmountPerYear,proto3,customtype=cosmossdk.io/math.Int" json:"eden_amount_per_year"`
	// blocks distributed
	BlocksDistributed int64 `protobuf:"varint,2,opt,name=blocks_distributed,json=blocksDistributed,proto3" json:"blocks_distributed,omitempty"`
}

func (m *IncentiveInfo) Reset()         { *m = IncentiveInfo{} }
func (m *IncentiveInfo) String() string { return proto.CompactTextString(m) }
func (*IncentiveInfo) ProtoMessage()    {}
func (*IncentiveInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbfbd92cca5ccf94, []int{1}
}
func (m *IncentiveInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IncentiveInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IncentiveInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IncentiveInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncentiveInfo.Merge(m, src)
}
func (m *IncentiveInfo) XXX_Size() int {
	return m.Size()
}
func (m *IncentiveInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_IncentiveInfo.DiscardUnknown(m)
}

var xxx_messageInfo_IncentiveInfo proto.InternalMessageInfo

func (m *IncentiveInfo) GetBlocksDistributed() int64 {
	if m != nil {
		return m.BlocksDistributed
	}
	return 0
}

func init() {
	proto.RegisterType((*LegacyIncentiveInfo)(nil), "elys.masterchef.LegacyIncentiveInfo")
	proto.RegisterType((*IncentiveInfo)(nil), "elys.masterchef.IncentiveInfo")
}

func init() { proto.RegisterFile("elys/masterchef/incentive.proto", fileDescriptor_fbfbd92cca5ccf94) }

var fileDescriptor_fbfbd92cca5ccf94 = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xb1, 0x8e, 0xd3, 0x40,
	0x10, 0x86, 0x6d, 0x82, 0x90, 0x58, 0x09, 0x21, 0x4c, 0x82, 0x4c, 0x0a, 0x1b, 0xa5, 0x42, 0x42,
	0xf1, 0x16, 0x3c, 0x01, 0x11, 0x8d, 0x25, 0x0a, 0x04, 0x15, 0x11, 0xc2, 0x5a, 0xdb, 0x13, 0x67,
	0x15, 0x7b, 0xc7, 0xda, 0x1d, 0x03, 0x7e, 0x0b, 0x2a, 0xaa, 0x7b, 0x8c, 0x7b, 0x88, 0x94, 0xd1,
	0x55, 0xa7, 0x2b, 0xa2, 0x53, 0xf2, 0x22, 0x27, 0xdb, 0xc9, 0x25, 0x45, 0x9a, 0x5c, 0x71, 0x9d,
	0x47, 0xf3, 0xfb, 0xfb, 0x67, 0x76, 0x7e, 0xe6, 0x43, 0x5e, 0x1b, 0x5e, 0x08, 0x43, 0xa0, 0x93,
	0x39, 0xcc, 0xb8, 0x54, 0x09, 0x28, 0x92, 0xbf, 0x21, 0x28, 0x35, 0x12, 0x3a, 0x2f, 0x1b, 0x41,
	0x70, 0x10, 0x0c, 0xfb, 0x19, 0x66, 0xd8, 0xf6, 0x78, 0xf3, 0xd5, 0xc9, 0x86, 0x7e, 0x86, 0x98,
	0xe5, 0xc0, 0xdb, 0x2a, 0xae, 0x66, 0x9c, 0x64, 0x01, 0x86, 0x44, 0x51, 0xee, 0x04, 0x6f, 0x13,
	0x34, 0x05, 0x9a, 0xa8, 0xfb, 0xb3, 0x2b, 0xba, 0xd6, 0xe8, 0x7f, 0x8f, 0xbd, 0xfe, 0x02, 0x99,
	0x48, 0xea, 0x70, 0x6f, 0x1e, 0xaa, 0x19, 0x3a, 0x3f, 0x59, 0x1f, 0x52, 0x50, 0x91, 0x28, 0xb0,
	0x52, 0x14, 0x95, 0xa0, 0xa3, 0x1a, 0x84, 0x76, 0xed, 0x77, 0xf6, 0xfb, 0xe7, 0x93, 0x0f, 0xcb,
	0xb5, 0x6f, 0xdd, 0xac, 0xfd, 0x41, 0xc7, 0x32, 0xe9, 0x22, 0x90, 0xc8, 0x0b, 0x41, 0xf3, 0x20,
	0x54, 0x74, 0x75, 0x39, 0x66, 0x3b, 0x93, 0x50, 0xd1, 0xb7, 0x57, 0x0d, 0xe8, 0x53, 0xcb, 0xf9,
	0x0a, 0xfa, 0x07, 0x08, 0xed, 0x00, 0x73, 0x53, 0x69, 0x48, 0xcb, 0xb8, 0x22, 0x89, 0x2a, 0x32,
	0x24, 0x34, 0x45, 0x71, 0x8e, 0xc9, 0xc2, 0x7d, 0x72, 0xbe, 0xc3, 0x9b, 0x63, 0xd8, 0xf7, 0x86,
	0x35, 0x69, 0x50, 0xce, 0x2f, 0x36, 0x20, 0x24, 0x91, 0x77, 0x64, 0x73, 0xd8, 0xa2, 0x77, 0xbe,
	0x87, 0xd3, 0x92, 0x5a, 0xae, 0xd9, 0xaf, 0x31, 0x65, 0xce, 0x8e, 0x7c, 0x3f, 0x00, 0xa4, 0xee,
	0xd3, 0x07, 0x3c, 0x51, 0x87, 0xf9, 0x7c, 0xa0, 0x8c, 0x2e, 0x6c, 0xf6, 0xe2, 0x31, 0x4f, 0x32,
	0x3e, 0xb9, 0x4b, 0x73, 0x8c, 0xde, 0x89, 0xf1, 0x26, 0xe1, 0x72, 0xe3, 0xd9, 0xab, 0x8d, 0x67,
	0xdf, 0x6e, 0x3c, 0xfb, 0xdf, 0xd6, 0xb3, 0x56, 0x5b, 0xcf, 0xba, 0xde, 0x7a, 0xd6, 0x94, 0x67,
	0x92, 0xe6, 0x55, 0x1c, 0x24, 0x58, 0xf0, 0x26, 0xbf, 0x63, 0x05, 0xf4, 0x07, 0xf5, 0xa2, 0x2d,
	0xf8, 0xdf, 0xe3, 0xbc, 0x53, 0x5d, 0x82, 0x89, 0x9f, 0xb5, 0x49, 0xfc, 0x78, 0x17, 0x00, 0x00,
	0xff, 0xff, 0xa8, 0xff, 0x3f, 0x7e, 0x0f, 0x03, 0x00, 0x00,
}

func (m *LegacyIncentiveInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LegacyIncentiveInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LegacyIncentiveInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.BlocksDistributed.Size()
		i -= size
		if _, err := m.BlocksDistributed.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.TotalBlocksPerYear.Size()
		i -= size
		if _, err := m.TotalBlocksPerYear.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.DistributionStartBlock.Size()
		i -= size
		if _, err := m.DistributionStartBlock.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.EdenAmountPerYear.Size()
		i -= size
		if _, err := m.EdenAmountPerYear.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *IncentiveInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IncentiveInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IncentiveInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlocksDistributed != 0 {
		i = encodeVarintIncentive(dAtA, i, uint64(m.BlocksDistributed))
		i--
		dAtA[i] = 0x10
	}
	{
		size := m.EdenAmountPerYear.Size()
		i -= size
		if _, err := m.EdenAmountPerYear.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintIncentive(dAtA []byte, offset int, v uint64) int {
	offset -= sovIncentive(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LegacyIncentiveInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.EdenAmountPerYear.Size()
	n += 1 + l + sovIncentive(uint64(l))
	l = m.DistributionStartBlock.Size()
	n += 1 + l + sovIncentive(uint64(l))
	l = m.TotalBlocksPerYear.Size()
	n += 1 + l + sovIncentive(uint64(l))
	l = m.BlocksDistributed.Size()
	n += 1 + l + sovIncentive(uint64(l))
	return n
}

func (m *IncentiveInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.EdenAmountPerYear.Size()
	n += 1 + l + sovIncentive(uint64(l))
	if m.BlocksDistributed != 0 {
		n += 1 + sovIncentive(uint64(m.BlocksDistributed))
	}
	return n
}

func sovIncentive(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIncentive(x uint64) (n int) {
	return sovIncentive(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LegacyIncentiveInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentive
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
			return fmt.Errorf("proto: LegacyIncentiveInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LegacyIncentiveInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EdenAmountPerYear", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EdenAmountPerYear.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionStartBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DistributionStartBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalBlocksPerYear", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalBlocksPerYear.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlocksDistributed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BlocksDistributed.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIncentive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentive
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
func (m *IncentiveInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentive
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
			return fmt.Errorf("proto: IncentiveInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IncentiveInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EdenAmountPerYear", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EdenAmountPerYear.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlocksDistributed", wireType)
			}
			m.BlocksDistributed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlocksDistributed |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIncentive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentive
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
func skipIncentive(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIncentive
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
					return 0, ErrIntOverflowIncentive
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
					return 0, ErrIntOverflowIncentive
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
				return 0, ErrInvalidLengthIncentive
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIncentive
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIncentive
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIncentive        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIncentive          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIncentive = fmt.Errorf("proto: unexpected end of group")
)
