package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var DoNotGrab = []string{"EXAMPLE", "EXAMPLE2"}

var AuthyDigits int = 7
var AuthyInterval int = 20
var AuthyIssuer string = "synack"
var AuthySecret string = ReadConfig(ConfigFile).Authy_Secret
var ConfigFile string = "config/config.json"
var Channel string = ReadConfig(ConfigFile).Channel_ID
var Email string = ReadConfig(ConfigFile).Email_Address
var Password string = ReadConfig(ConfigFile).Password
var Token string = ReadConfig(ConfigFile).Discord_Token
var Delay uint = 20
var LoggedIn = false

type Configurations struct {
	Channel_ID    string `json:"CHANNEL_ID"`
	Discord_Token string `json:"DISCORD_TOKEN"`
	BotPrefix     string `json:"BOT_PREFIX"`
	Authy_Secret  string `json:"AUTHY_SECRET"`
	Email_Address string `json:"EMAIL_ADDRESS"`
	Password      string `json:"PASSWORD"`
}

func ReadConfig(ConfigFile string) *Configurations {
	body, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		fmt.Println("Error reading config from file, ", err)
	}
	var conf Configurations
	json.Unmarshal(body, &conf)
	return &conf
}
