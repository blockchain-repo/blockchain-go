package models

type BlockBody struct {
	Transaction   []Transaction
	NodePubkey    string
	Voters        []string
	Timestamp     string
}

type Block struct {
	Id         string   `json:"id"`   //
	BlockBody  BlockBody              //
	Signature  string                 //
}