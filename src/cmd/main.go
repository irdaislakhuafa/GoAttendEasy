package main

import (
	"fmt"

	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/config"
	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/connection"
	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/flags"
)

const (
	cfgPath     = "etc/cfg/"
	cfgFileName = "conf.json"
)

func main() {
	// get env from option
	env, err := flags.ParseFlags(cfgPath, cfgFileName)
	if err != nil {
		panic(err)
	}

	cfg, err := config.ReadConfigFromFile(fmt.Sprintf("%s/%s/%s", cfgPath, *env, cfgFileName))
	if err != nil {
		panic(err)
	}

	client, err := connection.OpenConnection(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("client: %v\n", client)
}
