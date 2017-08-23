package models

import (
	"unichain-go/common"
	"unichain-go/config"
)

type BlockBody struct {
	Transactions  []Transaction
	NodePubkey    string
	Voters        []string
	Timestamp     string
}

type Block struct {
	Id         string    `json:"id"`   //
	BlockBody  BlockBody               //
	Signature  string                  //
}

func (b *Block) Sign() string {
	priv_key := config.Config.Keypair.PrivateKey
	msg := b.BodyToString()
	c := common.GetCrypto()
	sig := c.Sign(priv_key, msg)
	b.Signature = sig
	return sig
}

func (b *Block) GenerateId() string {
	c := common.GetCrypto()
	_id := c.Hash(b.BodyToString())
	b.Id = _id
	return _id
}

func (b *Block) BodyToString() string {
	return common.Serialize(b.BlockBody)
}

func (b *Block) ToString() string {
	return common.Serialize(b)
}

func (b *Block) ValidateBlock() error {
	var err error = nil
	//Check if the block was created by a federation node

	//Check that the signature is valid

	//Check that the block contains no duplicated transactions

	return err
}
