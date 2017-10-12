package models

import (
	"unichain-go/common"
)

type BlockBody struct {
	Transactions  []Transaction
	NodePubkey    string
	Voters        []string
	Timestamp     string
}

type Block struct {
	Id         string    `json:"id,omitempty"`   //
	BlockBody  BlockBody               //
	Signature  string                  //
}

func (b *Block) Sign(private string) string {
	msg := b.BodyToString()
	c := common.GetCrypto()
	sig := c.Sign(private, msg)
	b.Signature = sig
	return sig
}

func (b *Block) Verify() bool {
	msg := b.BodyToString()
	c := common.GetCrypto()
	flag := c.Verify(b.BlockBody.NodePubkey, msg, b.Signature)
	if flag == false {
		return false
	}
	return true
}

func (b *Block) GenerateId() string {
	c := common.GetCrypto()
	_id := c.Hash(b.BodyToString())
	b.Id = _id
	return _id
}

func (b *Block) CheckId() bool {
	c := common.GetCrypto()
	_id := c.Hash(b.BodyToString())
	return b.Id == _id
}

func (b *Block) BodyToString() string {
	return common.Serialize(b.BlockBody)
}

func (b *Block) ToString() string {
	return common.Serialize(b)
}
