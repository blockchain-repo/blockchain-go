package models

import (
	"reflect"

	"unichain-go/common"
	"unichain-go/config"
	"unichain-go/log"
)

type PreOut struct {
	Tx    string
	Index string
}

type Input struct {
	OwnersBefore string  //
	Signature    string  //
	PreOut       *PreOut //
}

type Output struct {
	OwnersAfter string  //
	Amount      float64 //
}

type Transaction struct {
	Id        string                 `json:"id,omitempty"` //
	Inputs    []Input                //
	Outputs   []Output               //
	Operation string                 //???
	Asset     string                 //
	Chain     string                 //???
	Metadata  map[string]interface{} //
	Version   string                 //
}

func (t *Transaction) GenerateId() string {
	c := common.GetCrypto()
	_id := c.Hash(t.BodyToString())
	t.Id = _id
	return _id
}

func (t *Transaction) Sign() string {
	priv_key := config.Config.Keypair.PrivateKey
	msg := t.StringForSign()
	c := common.GetCrypto()
	sig := c.Sign(priv_key, msg)

	for i := 0; i < len(t.Inputs); i++ {
		t.Inputs[i].Signature = sig
	}
	return sig
}

func (t *Transaction) ToString() string {
	return common.Serialize(t)
}
func (t *Transaction) TxToMap() string {
	m, err := common.StructToMap(t)
	if err != nil {
		log.Error(err.Error())
	}
	delete(m, "id")
	s := ToSlice(m["Inputs"])
	for i, in := range s {
		delete(in.(map[string]interface{}), "Signature")
		s[i] = in
	}
	m["Inputs"] = s
	return common.Serialize(m)
}
func (t *Transaction) BodyToString() string {
	m, err := common.StructToMap(t)
	if err != nil {
		log.Error(err.Error())
	}
	delete(m, "id")
	return common.Serialize(m)
}

func (t *Transaction) StringForSign() string {
	m, err := common.StructToMap(t)
	if err != nil {
		log.Error(err.Error())
	}
	delete(m, "id")
	s := ToSlice(m["Inputs"])
	for i, in := range s {
		delete(in.(map[string]interface{}), "Signature")
		s[i] = in
	}
	m["Inputs"] = s
	return common.Serialize(m)
}

func ToSlice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		panic("arr not slice")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}
