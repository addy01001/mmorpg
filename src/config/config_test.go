package config_test

import (
	"mmorpg-bot/src/config"
	"testing"
)

func ReadConfigTest(t *testing.T){
	err:=config.ReadConfig()
	if err!=nil{
		t.Fail()
	}
}