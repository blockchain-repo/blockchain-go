package config

import (
	"os/user"

	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"unichain-go/common"
	"unichain-go/log"
)

type _Config struct {
	Keypair Keypair
	Keyring []string `json:"Keyring"`
	LocalIp string   `json:"LocalIp"`
	Log     Log      `json:"Log"`
}

type Keypair struct {
	PublicKey  string `json:"PublicKey"`
	PrivateKey string `json:"PrivateKey"`
}

type Log struct {
	LogName      string
	LogSaveLevel int
	LogMaxDays   int
	LogMaxLines  int
	LogMaxSize   int
	LogRotate    bool
	LogDaily     bool
	LogSeparate  []string
	//#error = 3; warning = 4; info = 6; debug = 7
	LogLevel         int
	LogEnableConsole bool
}

var Config _Config

func init() {
	_user, err := user.Current()
	if err != nil {
		log.Error(err.Error())
	}
	fileName := _user.HomeDir + "/.unichain-go"
	_, err = os.Open(fileName)
	if err != nil {
		log.Info(err.Error())
		return
	}
	FileToConfig()
}

func FileToConfig() {
	_user, err := user.Current()
	if err != nil {
		log.Error(err.Error())
	}
	fileName := _user.HomeDir + "/.unichain-go"
	file, err := os.Open(fileName)
	if err != nil {
		log.Error(err.Error())
		log.Error("please create default config by 'unichain-go configure' or 'go run main.go configure'")
		os.Exit(1)
	}
	_byte, err := ioutil.ReadAll(file)
	if err != nil {
		log.Error(err.Error())
		log.Error("please checkout your config file OR remove it", fileName)
		os.Exit(1)
	}
	err = json.Unmarshal(_byte, &Config)
	if err != nil {
		log.Error(err.Error())
		log.Error("please checkout your config file OR remove it", fileName)
		os.Exit(1)
	}
}

func createNewConfig() _Config {
	var newConfig _Config
	c := common.GetCrypto()
	//keypair
	pub, priv := c.GenerateKeypair()
	newConfig.Keypair.PublicKey = pub
	newConfig.Keypair.PrivateKey = priv
	//keyring
	newConfig.Keyring = []string{}
	//LocalIp
	newConfig.LocalIp = "localhost"
	//log
	newConfig.Log.LogName = "/tmp/unichain-go-logs/unichain-go.log"
	newConfig.Log.LogSaveLevel = 7
	newConfig.Log.LogMaxDays = 10
	newConfig.Log.LogMaxLines = 0
	newConfig.Log.LogMaxSize = 0
	newConfig.Log.LogRotate = true
	newConfig.Log.LogDaily = true
	newConfig.Log.LogSeparate = []string{"error", "warning", "info", "debug"}
	newConfig.Log.LogLevel = 7
	newConfig.Log.LogEnableConsole = true
	return newConfig
}

func ConfigToFile() {
	_user, err := user.Current()
	if err != nil {
		log.Error(err.Error())
	}
	fileName := _user.HomeDir + "/.unichain-go"
	_, err = os.Stat(fileName)
	if err == nil { //文件存在
		fmt.Println("Config file already exist, do you want to override it?")
		fmt.Println("Please input y(es) or n(o) ")
		inputReader := bufio.NewReader(os.Stdin)
		p := make([]byte, 10)
		inputReader.Read(p)
		if p[0] != []byte("y")[0] {
			fmt.Println("Give up to override it!")
			return
		}
	}
	configfile, err := os.Create(fileName)
	defer configfile.Close()
	if err != nil {
		log.Error(err.Error())
	}

	newConfig := createNewConfig()
	str := common.SerializePretty(newConfig)
	_, err = configfile.Write([]byte(str + "\n"))
	if err != nil {
		log.Error(err.Error())
	} else {
		fmt.Println("crate config file successful!\n", str)
	}

}
