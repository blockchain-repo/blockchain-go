package models

import (
	"unichain-go/common"
	"unichain-go/config"
)

type VoteBody struct {
	IsValid       bool   //合约、合约交易投票结果，如true,false
	InvalidReason string //合约、合约交易投无效票原因
	PreviousBlock string //
	VoteBlock     string //
	Timestamp     string //节点投票时间戳
}

type Vote struct {
	Id         string   `json:"id"`    //投票唯一标识ID，UUID
	NodePubkey string                  //投票节点的公钥
	VoteBody   VoteBody                //投票信息
	Signature  string                  //投票节点签名
}

func (v *Vote) VerifySig() bool {
	signature := v.Signature
	pub := v.NodePubkey
	msg := v.BodyToString()
	c := common.GetCrypto()
	return c.Verify(pub, msg, signature)
}

func (v *Vote) Sign() string {
	priv_key := config.Config.Keypair.PrivateKey
	msg := v.BodyToString()
	c := common.GetCrypto()
	sig := c.Sign(priv_key, msg)
	v.Signature = sig
	return sig
}

func (v *Vote) GenerateId() string {
	_id := common.GenerateUUID()
	v.Id = _id
	return _id
}

func (v *Vote) BodyToString() string {
	return common.Serialize(v.VoteBody)
}

func (v *Vote) ToString() string {
	return common.Serialize(v)
}