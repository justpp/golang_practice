package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Emoji string `yaml:"emoji"`
}

func main() {
	start := time.Now()
	// 读取YAML配置文件
	data, err := ioutil.ReadFile("test.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// 将YAML配置文件解析为map[string]interface{}
	var configMap map[string]interface{}
	err = yaml.Unmarshal(data, &configMap)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// 将map[string]interface{}转换为Config结构体
	var config Config
	configBytes, err := yaml.Marshal(configMap)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// 修改emoji表情
	config.Emoji = "👍"

	// 将Config结构体转换为map[string]interface{}
	configMap = make(map[string]interface{})
	configBytes, err = yaml.Marshal(&config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = yaml.Unmarshal(configBytes, &configMap)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// 将map[string]interface{}中的Unicode转义序列替换为实际的emoji字符
	data, err = yaml.Marshal(&configMap)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	re := regexp.MustCompile(`\\u[0-9a-fA-F]{4}`)
	data = []byte(re.ReplaceAllStringFunc(string(data), func(match string) string {
		bytes := []byte(match)[2:]
		var codepoint uint32
		for _, b := range bytes {
			codepoint *= 16
			switch {
			case b >= '0' && b <= '9':
				codepoint += uint32(b - '0')
			case b >= 'a' && b <= 'f':
				codepoint += uint32(b - 'a' + 10)
			case b >= 'A' && b <= 'F':
				codepoint += uint32(b - 'A' + 10)
			}
		}
		return strings.ToUpper(string(rune(codepoint)))
	}))

	// 将字符串中的Unicode转义序列替换为实际的emoji字符
	re = regexp.MustCompile(`\\u[0-9a-fA-F]{4}`)
	data = []byte(re.ReplaceAllStringFunc(string(data), func(match string) string {
		bytes := []byte(match)[2:]
		var codepoint uint32
		for _, b := range bytes {
			codepoint *= 16
			switch {
			case b >= '0' && b <= '9':
				codepoint += uint32(b - '0')
			case b >= 'a' && b <= 'f':
				codepoint += uint32(b - 'a' + 10)
			case b >= 'A' && b <= 'F':
				codepoint += uint32(b - 'A' + 10)
			}
		}
		return strings.ToUpper(string(rune(codepoint)))
	}))

	// 将字符串中的双引号去掉，以便正确地解析emoji字符
	data = []byte(strings.ReplaceAll(string(data), "\"", ""))

	// 将map[string]interface{}写回到YAML配置文件中
	err = ioutil.WriteFile("new_config.yaml", data, 0644)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println("Done!")
	fmt.Println(time.Now().Sub(start).Seconds())
}
