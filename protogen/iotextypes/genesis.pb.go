// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/types/genesis.proto

package iotextypes

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Genesis struct {
	Blockchain           *GenesisBlockchain `protobuf:"bytes,1,opt,name=blockchain,proto3" json:"blockchain,omitempty"`
	Account              *GenesisAccount    `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Poll                 *GenesisPoll       `protobuf:"bytes,3,opt,name=poll,proto3" json:"poll,omitempty"`
	Rewarding            *GenesisRewarding  `protobuf:"bytes,4,opt,name=rewarding,proto3" json:"rewarding,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Genesis) Reset()         { *m = Genesis{} }
func (m *Genesis) String() string { return proto.CompactTextString(m) }
func (*Genesis) ProtoMessage()    {}
func (*Genesis) Descriptor() ([]byte, []int) {
	return fileDescriptor_8090b9f9a91af920, []int{0}
}

func (m *Genesis) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Genesis.Unmarshal(m, b)
}
func (m *Genesis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Genesis.Marshal(b, m, deterministic)
}
func (m *Genesis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Genesis.Merge(m, src)
}
func (m *Genesis) XXX_Size() int {
	return xxx_messageInfo_Genesis.Size(m)
}
func (m *Genesis) XXX_DiscardUnknown() {
	xxx_messageInfo_Genesis.DiscardUnknown(m)
}

var xxx_messageInfo_Genesis proto.InternalMessageInfo

func (m *Genesis) GetBlockchain() *GenesisBlockchain {
	if m != nil {
		return m.Blockchain
	}
	return nil
}

func (m *Genesis) GetAccount() *GenesisAccount {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *Genesis) GetPoll() *GenesisPoll {
	if m != nil {
		return m.Poll
	}
	return nil
}

func (m *Genesis) GetRewarding() *GenesisRewarding {
	if m != nil {
		return m.Rewarding
	}
	return nil
}

type GenesisBlockchain struct {
	Timestamp             int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	BlockGasLimit         uint64   `protobuf:"varint,2,opt,name=blockGasLimit,proto3" json:"blockGasLimit,omitempty"`
	ActionGasLimit        uint64   `protobuf:"varint,3,opt,name=actionGasLimit,proto3" json:"actionGasLimit,omitempty"`
	BlockInterval         int64    `protobuf:"varint,4,opt,name=blockInterval,proto3" json:"blockInterval,omitempty"`
	NumSubEpochs          uint64   `protobuf:"varint,5,opt,name=numSubEpochs,proto3" json:"numSubEpochs,omitempty"`
	NumDelegates          uint64   `protobuf:"varint,6,opt,name=numDelegates,proto3" json:"numDelegates,omitempty"`
	NumCandidateDelegates uint64   `protobuf:"varint,7,opt,name=numCandidateDelegates,proto3" json:"numCandidateDelegates,omitempty"`
	TimeBasedRotation     bool     `protobuf:"varint,8,opt,name=timeBasedRotation,proto3" json:"timeBasedRotation,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *GenesisBlockchain) Reset()         { *m = GenesisBlockchain{} }
func (m *GenesisBlockchain) String() string { return proto.CompactTextString(m) }
func (*GenesisBlockchain) ProtoMessage()    {}
func (*GenesisBlockchain) Descriptor() ([]byte, []int) {
	return fileDescriptor_8090b9f9a91af920, []int{1}
}

func (m *GenesisBlockchain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenesisBlockchain.Unmarshal(m, b)
}
func (m *GenesisBlockchain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenesisBlockchain.Marshal(b, m, deterministic)
}
func (m *GenesisBlockchain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisBlockchain.Merge(m, src)
}
func (m *GenesisBlockchain) XXX_Size() int {
	return xxx_messageInfo_GenesisBlockchain.Size(m)
}
func (m *GenesisBlockchain) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisBlockchain.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisBlockchain proto.InternalMessageInfo

func (m *GenesisBlockchain) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *GenesisBlockchain) GetBlockGasLimit() uint64 {
	if m != nil {
		return m.BlockGasLimit
	}
	return 0
}

func (m *GenesisBlockchain) GetActionGasLimit() uint64 {
	if m != nil {
		return m.ActionGasLimit
	}
	return 0
}

func (m *GenesisBlockchain) GetBlockInterval() int64 {
	if m != nil {
		return m.BlockInterval
	}
	return 0
}

func (m *GenesisBlockchain) GetNumSubEpochs() uint64 {
	if m != nil {
		return m.NumSubEpochs
	}
	return 0
}

func (m *GenesisBlockchain) GetNumDelegates() uint64 {
	if m != nil {
		return m.NumDelegates
	}
	return 0
}

func (m *GenesisBlockchain) GetNumCandidateDelegates() uint64 {
	if m != nil {
		return m.NumCandidateDelegates
	}
	return 0
}

func (m *GenesisBlockchain) GetTimeBasedRotation() bool {
	if m != nil {
		return m.TimeBasedRotation
	}
	return false
}

type GenesisAccount struct {
	InitBalanceAddrs     []string `protobuf:"bytes,1,rep,name=initBalanceAddrs,proto3" json:"initBalanceAddrs,omitempty"`
	InitBalances         []string `protobuf:"bytes,2,rep,name=initBalances,proto3" json:"initBalances,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenesisAccount) Reset()         { *m = GenesisAccount{} }
func (m *GenesisAccount) String() string { return proto.CompactTextString(m) }
func (*GenesisAccount) ProtoMessage()    {}
func (*GenesisAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_8090b9f9a91af920, []int{2}
}

func (m *GenesisAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenesisAccount.Unmarshal(m, b)
}
func (m *GenesisAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenesisAccount.Marshal(b, m, deterministic)
}
func (m *GenesisAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisAccount.Merge(m, src)
}
func (m *GenesisAccount) XXX_Size() int {
	return xxx_messageInfo_GenesisAccount.Size(m)
}
func (m *GenesisAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisAccount.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisAccount proto.InternalMessageInfo

func (m *GenesisAccount) GetInitBalanceAddrs() []string {
	if m != nil {
		return m.InitBalanceAddrs
	}
	return nil
}

func (m *GenesisAccount) GetInitBalances() []string {
	if m != nil {
		return m.InitBalances
	}
	return nil
}

type GenesisPoll struct {
	EnableGravityChainVoting bool               `protobuf:"varint,1,opt,name=enableGravityChainVoting,proto3" json:"enableGravityChainVoting,omitempty"`
	GravityChainStartHeight  uint64             `protobuf:"varint,2,opt,name=gravityChainStartHeight,proto3" json:"gravityChainStartHeight,omitempty"`
	RegisterContractAddress  string             `protobuf:"bytes,3,opt,name=registerContractAddress,proto3" json:"registerContractAddress,omitempty"`
	StakingContractAddress   string             `protobuf:"bytes,4,opt,name=stakingContractAddress,proto3" json:"stakingContractAddress,omitempty"`
	VoteThreshold            string             `protobuf:"bytes,5,opt,name=voteThreshold,proto3" json:"voteThreshold,omitempty"`
	ScoreThreshold           string             `protobuf:"bytes,6,opt,name=scoreThreshold,proto3" json:"scoreThreshold,omitempty"`
	SelfStakingThreshold     string             `protobuf:"bytes,7,opt,name=selfStakingThreshold,proto3" json:"selfStakingThreshold,omitempty"`
	Delegates                []*GenesisDelegate `protobuf:"bytes,8,rep,name=delegates,proto3" json:"delegates,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}           `json:"-"`
	XXX_unrecognized         []byte             `json:"-"`
	XXX_sizecache            int32              `json:"-"`
}

func (m *GenesisPoll) Reset()         { *m = GenesisPoll{} }
func (m *GenesisPoll) String() string { return proto.CompactTextString(m) }
func (*GenesisPoll) ProtoMessage()    {}
func (*GenesisPoll) Descriptor() ([]byte, []int) {
	return fileDescriptor_8090b9f9a91af920, []int{3}
}

func (m *GenesisPoll) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenesisPoll.Unmarshal(m, b)
}
func (m *GenesisPoll) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenesisPoll.Marshal(b, m, deterministic)
}
func (m *GenesisPoll) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisPoll.Merge(m, src)
}
func (m *GenesisPoll) XXX_Size() int {
	return xxx_messageInfo_GenesisPoll.Size(m)
}
func (m *GenesisPoll) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisPoll.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisPoll proto.InternalMessageInfo

func (m *GenesisPoll) GetEnableGravityChainVoting() bool {
	if m != nil {
		return m.EnableGravityChainVoting
	}
	return false
}

func (m *GenesisPoll) GetGravityChainStartHeight() uint64 {
	if m != nil {
		return m.GravityChainStartHeight
	}
	return 0
}

func (m *GenesisPoll) GetRegisterContractAddress() string {
	if m != nil {
		return m.RegisterContractAddress
	}
	return ""
}

func (m *GenesisPoll) GetStakingContractAddress() string {
	if m != nil {
		return m.StakingContractAddress
	}
	return ""
}

func (m *GenesisPoll) GetVoteThreshold() string {
	if m != nil {
		return m.VoteThreshold
	}
	return ""
}

func (m *GenesisPoll) GetScoreThreshold() string {
	if m != nil {
		return m.ScoreThreshold
	}
	return ""
}

func (m *GenesisPoll) GetSelfStakingThreshold() string {
	if m != nil {
		return m.SelfStakingThreshold
	}
	return ""
}

func (m *GenesisPoll) GetDelegates() []*GenesisDelegate {
	if m != nil {
		return m.Delegates
	}
	return nil
}

type GenesisDelegate struct {
	OperatorAddr         string   `protobuf:"bytes,1,opt,name=operatorAddr,proto3" json:"operatorAddr,omitempty"`
	RewardAddr           string   `protobuf:"bytes,2,opt,name=rewardAddr,proto3" json:"rewardAddr,omitempty"`
	Votes                string   `protobuf:"bytes,3,opt,name=votes,proto3" json:"votes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenesisDelegate) Reset()         { *m = GenesisDelegate{} }
func (m *GenesisDelegate) String() string { return proto.CompactTextString(m) }
func (*GenesisDelegate) ProtoMessage()    {}
func (*GenesisDelegate) Descriptor() ([]byte, []int) {
	return fileDescriptor_8090b9f9a91af920, []int{4}
}

func (m *GenesisDelegate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenesisDelegate.Unmarshal(m, b)
}
func (m *GenesisDelegate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenesisDelegate.Marshal(b, m, deterministic)
}
func (m *GenesisDelegate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisDelegate.Merge(m, src)
}
func (m *GenesisDelegate) XXX_Size() int {
	return xxx_messageInfo_GenesisDelegate.Size(m)
}
func (m *GenesisDelegate) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisDelegate.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisDelegate proto.InternalMessageInfo

func (m *GenesisDelegate) GetOperatorAddr() string {
	if m != nil {
		return m.OperatorAddr
	}
	return ""
}

func (m *GenesisDelegate) GetRewardAddr() string {
	if m != nil {
		return m.RewardAddr
	}
	return ""
}

func (m *GenesisDelegate) GetVotes() string {
	if m != nil {
		return m.Votes
	}
	return ""
}

type GenesisRewarding struct {
	InitAdminAddr              string   `protobuf:"bytes,1,opt,name=initAdminAddr,proto3" json:"initAdminAddr,omitempty"`
	InitBalance                string   `protobuf:"bytes,2,opt,name=initBalance,proto3" json:"initBalance,omitempty"`
	BlockReward                string   `protobuf:"bytes,3,opt,name=blockReward,proto3" json:"blockReward,omitempty"`
	EpochReward                string   `protobuf:"bytes,4,opt,name=epochReward,proto3" json:"epochReward,omitempty"`
	NumDelegatesForEpochReward uint64   `protobuf:"varint,5,opt,name=numDelegatesForEpochReward,proto3" json:"numDelegatesForEpochReward,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *GenesisRewarding) Reset()         { *m = GenesisRewarding{} }
func (m *GenesisRewarding) String() string { return proto.CompactTextString(m) }
func (*GenesisRewarding) ProtoMessage()    {}
func (*GenesisRewarding) Descriptor() ([]byte, []int) {
	return fileDescriptor_8090b9f9a91af920, []int{5}
}

func (m *GenesisRewarding) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenesisRewarding.Unmarshal(m, b)
}
func (m *GenesisRewarding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenesisRewarding.Marshal(b, m, deterministic)
}
func (m *GenesisRewarding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisRewarding.Merge(m, src)
}
func (m *GenesisRewarding) XXX_Size() int {
	return xxx_messageInfo_GenesisRewarding.Size(m)
}
func (m *GenesisRewarding) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisRewarding.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisRewarding proto.InternalMessageInfo

func (m *GenesisRewarding) GetInitAdminAddr() string {
	if m != nil {
		return m.InitAdminAddr
	}
	return ""
}

func (m *GenesisRewarding) GetInitBalance() string {
	if m != nil {
		return m.InitBalance
	}
	return ""
}

func (m *GenesisRewarding) GetBlockReward() string {
	if m != nil {
		return m.BlockReward
	}
	return ""
}

func (m *GenesisRewarding) GetEpochReward() string {
	if m != nil {
		return m.EpochReward
	}
	return ""
}

func (m *GenesisRewarding) GetNumDelegatesForEpochReward() uint64 {
	if m != nil {
		return m.NumDelegatesForEpochReward
	}
	return 0
}

func init() {
	proto.RegisterType((*Genesis)(nil), "iotextypes.Genesis")
	proto.RegisterType((*GenesisBlockchain)(nil), "iotextypes.GenesisBlockchain")
	proto.RegisterType((*GenesisAccount)(nil), "iotextypes.GenesisAccount")
	proto.RegisterType((*GenesisPoll)(nil), "iotextypes.GenesisPoll")
	proto.RegisterType((*GenesisDelegate)(nil), "iotextypes.GenesisDelegate")
	proto.RegisterType((*GenesisRewarding)(nil), "iotextypes.GenesisRewarding")
}

func init() { proto.RegisterFile("proto/types/genesis.proto", fileDescriptor_8090b9f9a91af920) }

var fileDescriptor_8090b9f9a91af920 = []byte{
	// 657 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0x5d, 0x6f, 0xd3, 0x3e,
	0x14, 0xc6, 0x95, 0xa6, 0x5b, 0x97, 0xd3, 0xff, 0x7f, 0x6c, 0xd6, 0x60, 0x61, 0x0c, 0x54, 0x45,
	0x08, 0x55, 0xbc, 0xb4, 0xd2, 0x98, 0xa6, 0x31, 0x09, 0xa4, 0x75, 0x8c, 0x81, 0xc4, 0x05, 0xf2,
	0x10, 0x17, 0x5c, 0xe1, 0x26, 0x26, 0x35, 0x4b, 0xec, 0xc8, 0x76, 0x07, 0xfb, 0x5a, 0x7c, 0x13,
	0xbe, 0x00, 0xb7, 0x7c, 0x0d, 0x64, 0xa7, 0x6d, 0x9c, 0xbe, 0x70, 0x99, 0xe7, 0xfc, 0x1e, 0xbb,
	0x8f, 0xcf, 0x39, 0x85, 0xbb, 0x85, 0x14, 0x5a, 0xf4, 0xf5, 0x4d, 0x41, 0x55, 0x3f, 0xa5, 0x9c,
	0x2a, 0xa6, 0x7a, 0x56, 0x43, 0xc0, 0x84, 0xa6, 0x3f, 0x6c, 0x25, 0xfa, 0xe3, 0x41, 0xeb, 0xa2,
	0xac, 0xa2, 0x97, 0x00, 0xc3, 0x4c, 0xc4, 0x57, 0xf1, 0x88, 0x30, 0x1e, 0x7a, 0x1d, 0xaf, 0xdb,
	0x3e, 0xb8, 0xdf, 0xab, 0xe0, 0xde, 0x04, 0x1c, 0xcc, 0x20, 0xec, 0x18, 0xd0, 0x21, 0xb4, 0x48,
	0x1c, 0x8b, 0x31, 0xd7, 0x61, 0xc3, 0x7a, 0xf7, 0x96, 0x78, 0x4f, 0x4b, 0x02, 0x4f, 0x51, 0xf4,
	0x04, 0x9a, 0x85, 0xc8, 0xb2, 0xd0, 0xb7, 0x96, 0xdd, 0x25, 0x96, 0x0f, 0x22, 0xcb, 0xb0, 0x85,
	0xd0, 0x09, 0x04, 0x92, 0x7e, 0x27, 0x32, 0x61, 0x3c, 0x0d, 0x9b, 0xd6, 0xb1, 0xbf, 0xc4, 0x81,
	0xa7, 0x0c, 0xae, 0xf0, 0xe8, 0x57, 0x03, 0xb6, 0x17, 0x02, 0xa0, 0x7d, 0x08, 0x34, 0xcb, 0xa9,
	0xd2, 0x24, 0x2f, 0x6c, 0x64, 0x1f, 0x57, 0x02, 0x7a, 0x08, 0xff, 0xdb, 0x80, 0x17, 0x44, 0xbd,
	0x67, 0x39, 0x2b, 0x83, 0x35, 0x71, 0x5d, 0x44, 0x8f, 0x60, 0x93, 0xc4, 0x9a, 0x09, 0x3e, 0xc3,
	0x7c, 0x8b, 0xcd, 0xa9, 0xb3, 0xd3, 0xde, 0x71, 0x4d, 0xe5, 0x35, 0xc9, 0x6c, 0x02, 0x1f, 0xd7,
	0x45, 0x14, 0xc1, 0x7f, 0x7c, 0x9c, 0x5f, 0x8e, 0x87, 0xe7, 0x85, 0x88, 0x47, 0x2a, 0x5c, 0xb3,
	0x67, 0xd5, 0xb4, 0x09, 0xf3, 0x9a, 0x66, 0x34, 0x25, 0x9a, 0xaa, 0x70, 0x7d, 0xc6, 0xcc, 0x34,
	0x74, 0x08, 0xb7, 0xf9, 0x38, 0x3f, 0x23, 0x3c, 0x61, 0x09, 0xd1, 0xb4, 0x82, 0x5b, 0x16, 0x5e,
	0x5e, 0x44, 0x4f, 0x61, 0xdb, 0xc4, 0x1f, 0x10, 0x45, 0x13, 0x2c, 0x34, 0x31, 0x01, 0xc2, 0x8d,
	0x8e, 0xd7, 0xdd, 0xc0, 0x8b, 0x85, 0xe8, 0x0b, 0x6c, 0xd6, 0xfb, 0x8a, 0x1e, 0xc3, 0x16, 0xe3,
	0x4c, 0x0f, 0x48, 0x46, 0x78, 0x4c, 0x4f, 0x93, 0x44, 0xaa, 0xd0, 0xeb, 0xf8, 0xdd, 0x00, 0x2f,
	0xe8, 0x26, 0x85, 0xa3, 0xa9, 0xb0, 0x61, 0xb9, 0x9a, 0x16, 0xfd, 0xf4, 0xa1, 0xed, 0xcc, 0x01,
	0x3a, 0x81, 0x90, 0x72, 0x32, 0xcc, 0xe8, 0x85, 0x24, 0xd7, 0x4c, 0xdf, 0x9c, 0x99, 0x2e, 0x7e,
	0x12, 0xda, 0x0c, 0x84, 0x67, 0x7f, 0xe6, 0xca, 0x3a, 0x3a, 0x86, 0xdd, 0xd4, 0x51, 0x2f, 0x35,
	0x91, 0xfa, 0x2d, 0x65, 0xe9, 0x68, 0xda, 0xd7, 0x55, 0x65, 0xe3, 0x94, 0x34, 0x65, 0x4a, 0x53,
	0x79, 0x26, 0xb8, 0x96, 0x24, 0xd6, 0x26, 0x02, 0x55, 0xca, 0xb6, 0x3a, 0xc0, 0xab, 0xca, 0xe8,
	0x08, 0xee, 0x28, 0x4d, 0xae, 0x18, 0x4f, 0xe7, 0x8d, 0x4d, 0x6b, 0x5c, 0x51, 0x35, 0xb3, 0x72,
	0x2d, 0x34, 0xfd, 0x38, 0x92, 0x54, 0x8d, 0x44, 0x96, 0xd8, 0x31, 0x08, 0x70, 0x5d, 0x34, 0x93,
	0xa7, 0x62, 0x21, 0x1d, 0x6c, 0xdd, 0x62, 0x73, 0x2a, 0x3a, 0x80, 0x1d, 0x45, 0xb3, 0xaf, 0x97,
	0xe5, 0x5d, 0x15, 0xdd, 0xb2, 0xf4, 0xd2, 0x1a, 0x7a, 0x01, 0x41, 0x32, 0x9b, 0x99, 0x8d, 0x8e,
	0xdf, 0x6d, 0x1f, 0xdc, 0x5b, 0xb2, 0x6b, 0xd3, 0xd1, 0xc1, 0x15, 0x1d, 0x5d, 0xc1, 0xad, 0xb9,
	0xaa, 0xe9, 0xb5, 0x28, 0xa8, 0x24, 0x5a, 0x48, 0x13, 0xd1, 0xf6, 0x2a, 0xc0, 0x35, 0x0d, 0x3d,
	0x00, 0x28, 0xd7, 0xd5, 0x12, 0x0d, 0x4b, 0x38, 0x0a, 0xda, 0x81, 0x35, 0x13, 0x7f, 0xfa, 0xe6,
	0xe5, 0x47, 0xf4, 0xdb, 0x83, 0xad, 0xf9, 0xbd, 0x37, 0xcf, 0x67, 0xc6, 0xe8, 0x34, 0xc9, 0x19,
	0x77, 0xee, 0xab, 0x8b, 0xa8, 0x03, 0x6d, 0x67, 0xd8, 0x26, 0x37, 0xba, 0x92, 0x21, 0xec, 0x76,
	0x96, 0x27, 0x4f, 0x2e, 0x76, 0x25, 0x43, 0x50, 0xb3, 0x94, 0x13, 0xa2, 0xec, 0xaa, 0x2b, 0xa1,
	0x57, 0xb0, 0xe7, 0x2e, 0xe6, 0x1b, 0x21, 0xcf, 0x1d, 0x43, 0xb9, 0xde, 0xff, 0x20, 0x06, 0xc7,
	0x9f, 0x8f, 0x52, 0xa6, 0x47, 0xe3, 0x61, 0x2f, 0x16, 0x79, 0xdf, 0x76, 0xa0, 0x90, 0xe2, 0x1b,
	0x8d, 0x75, 0xf9, 0xf1, 0xcc, 0xf4, 0xba, 0x6f, 0xff, 0xda, 0x53, 0xca, 0xfb, 0x55, 0x8b, 0x86,
	0xeb, 0x56, 0x7c, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0x05, 0xd5, 0xce, 0x9f, 0x0c, 0x06, 0x00,
	0x00,
}