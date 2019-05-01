package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/spf13/viper"
)

func main() {
	actWrite()
	expWrite()
}

func actWrite() {
	viper.Set("version", "3")
	viper.SetConfigName("config-act")
	viper.AddConfigPath(".")

	viper.Set("kafka-server.image", "wurstmeister/kafka")
	viper.Set("kafka-server.environment.KAFKA_LISTENERS", "INSIDE://:9092,OUTSIDE://:29092")
	viper.WriteConfig()
}

func expWrite() {
	configName := "config-exp"
	viper := viper.New()
	viper.Set("version", "3")
	viper.SetConfigName(configName)
	viper.AddConfigPath(".")

	viper.Set("kafka-server.image", "wurstmeister/kafka")
	viper.Set("kafka-server.environment.KAFKA_LISTENERS", "INSIDE://:9092,OUTSIDE://:29092")

	ymlStr, err := yamlStringSettings(viper)
	if err != nil {
		fmt.Printf("write to %v.yml error:%v\n", configName, err)
		return
	}
	ymlStr = strings.Replace(ymlStr, "kafka_listeners", "KAFKA_LISTENERS", -1)

	if writeFile(configName+".yml", ymlStr); err != nil {
		fmt.Printf("write to %v.yml error:%v", configName, err)
		return
	}
}

func yamlStringSettings(vip *viper.Viper) (ymlString string, err error) {
	c := vip.AllSettings()
	bs, err := yaml.Marshal(c)
	if err != nil {
		return
	}
	ymlString = string(bs)
	return
}

func writeFile(fileName, content string) (err error) {
	out, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = out.WriteString(content)
	return
}
