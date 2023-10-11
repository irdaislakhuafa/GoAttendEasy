package config

import (
	"encoding/json"
	"os"

	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/errors"
	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/errors/codes"
	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/files"
)

func ReadConfigFromFile(filePath string) (AppConfig, error) {
	if !files.IsFileExist(filePath) {
		return AppConfig{}, errors.NewWithCode(codes.CodeFileNotExist, "file with path '%s' not exist", filePath)
	}

	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return AppConfig{}, errors.NewWithCode(codes.CodeCannotReadFile, "cannot read file, %s", err.Error())
	}

	cfg := AppConfig{}
	if err := json.Unmarshal(fileBytes, &cfg); err != nil {
		return AppConfig{}, errors.NewWithCode(codes.CodeJSONUnmarshalError, err.Error())
	}

	return cfg, nil

}
