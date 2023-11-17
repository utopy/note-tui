package core

import (
	"os"
	"path"
)

type PathsInterface struct {
	ConfigDirName string
	ConfigPath    string
	Home          string
}

func StartPathInterface(configDirName string) *PathsInterface {

	homeDir, err := os.UserHomeDir()

	if err != nil {
		panic("UNABLE TO GET HOME DIR")
	}

	return &PathsInterface{
		Home:          homeDir,
		ConfigDirName: configDirName,
		ConfigPath:    path.Join(homeDir, configDirName),
	}

}

func (p *PathsInterface) Setup() {

	err := os.MkdirAll(p.ConfigPath, 0755)

	if err != nil {
		panic("ERROR CREATING FOLDER AT CONFIG DIR PATH")
	}

}
