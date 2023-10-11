package flags

import (
	"flag"
	"fmt"

	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/errors"
	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/errors/codes"
	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/files"
)

// default env is "local"
func ParseFlags(configPath, configFile string) (*string, error) {
	env := ""
	flag.StringVar(&env, "env", "local", fmt.Sprintf("mode from dir name in %v/", configPath))
	flag.Parse()

	if dir := fmt.Sprintf("%v/%v", configPath, env); !files.IsFileExist(dir) {
		return nil, errors.NewWithCode(codes.CodeFileNotExist, "env '%v' not found in dir '%v'", env, configPath)
	} else {
		if cfgFile := fmt.Sprintf("%v/%v", dir, configFile); !files.IsFileExist(cfgFile) {
			return nil, errors.NewWithCode(codes.CodeFileNotExist, "cfg file '%v' not found in '%v'", configFile, dir)
		}
	}

	return &env, nil
}
