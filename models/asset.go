package models

import (
	"unichain-go/common"
)

type Asset struct {
	Id   string `json:"id"` //
	Data map[string]interface{}
}

func (a *Asset) GenerateId() string {
	_id := common.GenerateUUID()
	a.Id = _id
	return _id
}

func (a *Asset) ToString() string {
	return common.Serialize(a)
}
