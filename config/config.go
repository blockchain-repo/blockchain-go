package config

import (
	"os/user"

	"unichain-go/log"
	"os"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"bufio"
	"unichain-go/common"
)

type _Config struct {
	Keypair Keypair
	Keyring []string `json:"Keyring"`
	LocalIp string   `json:"LocalIp"`
}

type Keypair struct {
	PublicKey  string `json:"PublicKey"`
	PrivateKey string `json:"PrivateKey"`
}

var Config _Config

func FileToConfig() {
	_user, err := user.Current()
	if err != nil {
		log.Error(err.Error())
	}
	fileName := _user.HomeDir + "/.unichain-go"
	file, err := os.Open(fileName)
	if err != nil {
		log.Error(err.Error())
	}
	_byte, err := ioutil.ReadAll(file)
	if err != nil {
		log.Error(err.Error())
	}
	err = json.Unmarshal(_byte, &Config)
	if err != nil {
		log.Error(err.Error())
	}
}

func ConfigToFile() {
	_user, err := user.Current()
	if err != nil {
		log.Error(err.Error())
	}
	fileName := _user.HomeDir + "/.unichain-go"
	fileInfo, err := os.Stat(fileName)
	if err == nil { //文件存在
		fmt.Println("Config file already exist, do you want to override it?")
		fmt.Println("Please input y(es) or n(o) ")
		inputReader := bufio.NewReader(os.Stdin)
		p := make([]byte, 10)
		inputReader.Read(p)
		if p[0] != []byte("y")[0] {
			fmt.Println("Give Up to override it!", fileInfo)
			return
		}
	}
	configfile, err := os.Create(fileName)
	defer configfile.Close()
	if err != nil {
		log.Error(err.Error())
	}
	var newConfig _Config
	c := common.GetCrypto()
	pub, priv := c.GenerateKeypair()
	newConfig.Keypair.PublicKey = pub
	newConfig.Keypair.PrivateKey = priv
	newConfig.Keyring = []string{}
	str := common.Serialize(newConfig)
	n, err := configfile.Write([]byte(str))
	if err != nil {
		log.Error(err.Error())
	} else {
		fmt.Println("crate config file successful", n)
	}

}