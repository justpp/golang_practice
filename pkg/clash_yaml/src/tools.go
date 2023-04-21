package src

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"giao/pkg/util"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type V2Vmess struct {
	V    string `json:"v"`
	Ps   []byte `json:"ps"`
	Add  string `json:"add"`
	Port string `json:"port"`
	Id   string `json:"id"`
	Aid  string `json:"aid"`
	Scy  string `json:"scy"`
	Net  string `json:"net"`
	Type string `json:"type"`
	Host string `json:"host"`
	Path string `json:"path"`
	Tls  bool   `json:"tls"`
	Sni  string `json:"sni"`
	Alpn string `json:"alpn"`
}

// type ClashVmess struct {
// 	Name           string `json:"name" yaml:"name,flow"`
// 	Type           string `json:"type" yaml:"type,flow"`
// 	Server         string `json:"server" yaml:"server,flow"`
// 	Port           string `json:"port" yaml:"port,flow"`
// 	Cipher         string `json:"cipher" yaml:"cipher,flow"`
// 	Uuid           string `json:"uuid" yaml:"uuid,flow"`
// 	AlterId        int    `json:"alterId" yaml:"alterId,flow"`
// 	Tls            bool   `json:"tls" yaml:"tls,flow"`
// 	SkipCertVerify bool   `json:"skip-cert-verify" yaml:"skip-cert-verify,flow"`
// 	Network        string `json:"network" yaml:"network,flow"`
// }

type ClashVmess struct {
	Name           []byte `json:"name" yaml:"name"`
	Type           string `json:"type" yaml:"type"`
	Server         string `json:"server" yaml:"server"`
	Port           string `json:"port" yaml:"port"`
	Cipher         string `json:"cipher" yaml:"cipher"`
	Uuid           string `json:"uuid" yaml:"uuid"`
	AlterId        int    `json:"alterId" yaml:"alterId"`
	Tls            bool   `json:"tls" yaml:"tls"`
	SkipCertVerify bool   `json:"skip-cert-verify" yaml:"skip-cert-verify"`
	Network        string `json:"network" yaml:"network"`
}

func Vmess2Clash(vmessStr string, tName string) {
	split := strings.Split(strings.TrimSpace(vmessStr), "\n")
	var vmessNameS [][]byte
	var vssS []*ClashVmess
	for _, s := range split {
		clashT := v2EncodeClashT(s)
		vmessNameS = append(vmessNameS, clashT.Name)
		vssS = append(vssS, clashT)
	}

	// var proxies []map[string][][]byte

	config := GetTempConfig("")
	// err := config.UnmarshalKey("proxy-groups", &proxies)
	// util.CheckErr(err)
	//
	// for i, proxy := range proxies {
	// 	proxiesItem := proxy["proxies"]
	//
	// 	i2 := append(proxiesItem, vmessNameS...)
	//
	// 	proxies[i]["proxies"] = i2
	//
	// }
	//
	// config.Set("proxies", vssS)
	// config.Set("proxy-groups", proxies)
	fmt.Println(config.AllSettings())
	marshal, err := yaml.Marshal(config.AllSettings())

	if err != nil {
		util.CheckErr(err)
		return
	}
	err = os.WriteFile("own.yaml", marshal, 0644)
	// err = config.WriteConfigAs("own.yaml")
	util.CheckErr(err)

}

func v2EncodeClashT(str string) *ClashVmess {
	str = strings.Replace(str, "vmess://", "", 1)
	decodeString, err := base64.StdEncoding.DecodeString(str)
	util.CheckErr(err)
	var v2 V2Vmess
	err = json.Unmarshal(decodeString, &v2)
	util.CheckErr(err)
	return &ClashVmess{
		Name:           v2.Ps,
		Type:           "vmess",
		Server:         v2.Add,
		Port:           v2.Port,
		Cipher:         "auto",
		Uuid:           v2.Id,
		Tls:            v2.Tls,
		SkipCertVerify: true,
		Network:        v2.Net,
	}
}

func GetTempConfig(name string) *viper.Viper {
	if name == "" {
		name = "default"
	}
	name = "t_" + name + ".yaml"

	return util.NewEnv(
		util.SetEnvName(name),
		util.SetEnvPath("./template"))
}
