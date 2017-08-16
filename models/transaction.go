package models

import (
	"unichain-go/common"
	"unichain-go/log"
)

type PreOut struct {
	Tx            string
	Index         string
}

type Input struct {
	OwnersBefore  string     //
	Signature     string     //
	PreOut        PreOut     //
}

type Output struct {
	OwnersAfter   string   //
	Amount        string   //
}

type Transaction struct {
	Id         string   `json:"id"`      //
	Inputs     []Input                   //
	Outputs    []Output                  //
	Operation  string                    //???
	Asset      string                    //
	Chain      string                    //???
	Metadata   map[string]interface{}    //
	Version    string                    //
}

func (t *Transaction) GenerateId() string {
	c := common.GetCrypto()
	_id := c.Hash(t.BodyToString())
	t.Id = _id
	return _id
}

func (t *Transaction) ToString() string {
	return common.Serialize(t)
}

func (t *Transaction) BodyToString() string {
	m,err := common.StructToMap(t)
	if err != nil {
		log.Error(err.Error())
	}
	delete(m, "Id")
	return common.Serialize(m)
}

//tx := models.Transaction{}
//fmt.Println(common.Serialize(tx))
//m,err := common.StructToMap(tx)
//if err != nil {
//log.Error(err.Error())
//}
//delete(m, "id")
//fmt.Println(common.Serialize(m))