package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type ConfigInfo struct {
	Plans []*Plan
}

func loadConfig(configFile string) (*ConfigInfo, error) {
	var c *ConfigInfo
	path := configFile
	fi, err := os.Open(path)
	defer fi.Close()
	if nil != err {
		return nil, err
	}

	fd, err := ioutil.ReadAll(fi)
	err = json.Unmarshal([]byte(fd), &c)
	if nil != err {
		return nil, err
	}

	return c, nil
}
