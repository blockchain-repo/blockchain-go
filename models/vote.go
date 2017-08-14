package models

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
