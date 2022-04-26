/*
Copyright 2020 LINE Corporation

LINE Corporation licenses this file to you under the Apache License,
version 2.0 (the "License"); you may not use this file except in compliance
with the License. You may obtain a copy of the License at:

  https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
License for the specific language governing permissions and limitations
under the License
*/
package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

type APIConfig struct {
	LBDAPIEndpoint       string `json:"lbd-api-endpoint"`
	LINEAPIEndpoint      string `json:"line-api-endpoint"`
	LINEAccessEndpoint   string `json:"lineAccessEndpoint"`
	Endpoint             string `json:"endpoint"`
	WalletAddress        string `json:"walletAddress"`
	WalletSecret         string `json:"walletSecret"`
	APIKey               string `json:"apiKey"`
	APISecret            string `json:"apiSecret"`
	ChannelID            string `json:"channel-id"`
	ChannelSecret        string `json:"channelSecret"`
	ServiceContractID    string `json:"serviceContract-id"`
	ItemContractID       string `json:"itemContract-id"`
	FungibleTokenType    string `json:"fungibleTokenType"`
	NonFungibleTokenType string `json:"non-fungibleTokenType"`
	UserID               string `json:"user-id"`
	GuiUserId            string `json:"gui-user-id"`
	GuiUserPw            string `json:"gui-user-pw"`
	TestUserId            string `json:"test-user-id"`
	TestUserPw            string `json:"test-user-pw"`
}

const (
	Path = "CONFIG_PATH"
)

var (
	apiConfig = &APIConfig{}
)

func GetAPIConfig() *APIConfig {
	return apiConfig
}

func SetAPIConfig(config *APIConfig) {
	apiConfig = config
}

func LoadAPIConfig(path string) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if _, err := toml.Decode(string(dat), apiConfig); err != nil {
		fmt.Println(err.Error())
		return
	}
}

func InitAPIConfig() {
	// apiConfig.LBDAPIEndpoint = os.Getenv("LBDAPIEndpoint")
	// apiConfig.LINEAPIEndpoint = os.Getenv("LINEAPIEndpoint")
	// apiConfig.LINEAccessEndpoint = os.Getenv("LINEAccessEndpoint")
	apiConfig.Endpoint = os.Getenv("Endpoint")
// 	apiConfig.WalletAddress = os.Getenv("WalletAddress")
// 	apiConfig.WalletSecret = os.Getenv("WalletSecret")
// 	apiConfig.APIKey = os.Getenv("APIKey")
// 	apiConfig.APISecret = os.Getenv("APISecret")
// 	apiConfig.ChannelID = os.Getenv("ChannelID")
// 	apiConfig.ChannelSecret = os.Getenv("ChannelSecret")
// 	apiConfig.ServiceContractID = os.Getenv("ServiceContractID")
// 	apiConfig.ItemContractID = os.Getenv("ItemContractID")
// 	apiConfig.FungibleTokenType = os.Getenv("FungibleTokenType")
// 	apiConfig.NonFungibleTokenType = os.Getenv("NonFungibleTokenType")
// 	apiConfig.UserID = os.Getenv("UserID")
// 	apiConfig.GuiUserId = os.Getenv("GuiUserId")
// 	apiConfig.GuiUserPw = os.Getenv("GuiUserPw")
// 	apiConfig.TestUserId = os.Getenv("TestUserId")
// 	apiConfig.TestUserPw = os.Getenv("TestUserPw")
}
