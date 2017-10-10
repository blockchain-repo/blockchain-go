package models

import (
	"reflect"

	"unichain-go/common"
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
	OwnersAfter string //
	Amount      int    //
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

func (t *Transaction) CheckId() bool {
	c := common.GetCrypto()
	_id := c.Hash(t.BodyToString())
	return t.Id == _id
}

func (t *Transaction) Sign(private []string) bool {
	c := common.GetCrypto()
	keyMap := make(map[string]string)
	for _, seed := range private {
		pub, pri := c.GenerateKeypair(seed)
		keyMap[pub] = pri
	}
	msg := t.StringForSign()
	for i := 0; i < len(t.Inputs); i++ {
		pub := t.Inputs[i].OwnersBefore
		priv, ok := keyMap[pub]
		if !ok {
			return false
		}
		sig := c.Sign(priv, msg)
		t.Inputs[i].Signature = sig
	}
	return true
}

func (t *Transaction) Verify() bool {
	msg := t.StringForSign()
	c := common.GetCrypto()
	for i := 0; i < len(t.Inputs); i++ {
		flag := c.Verify(t.Inputs[i].OwnersBefore, msg, t.Inputs[i].Signature)
		if flag == false {
			return false
		}
	}
	return true
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
