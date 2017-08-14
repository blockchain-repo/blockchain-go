package models

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