package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	VERSION = "0.0.1"
)

var (
	configFile string
	version    bool
	testMode   bool
)

func init() {
	flag.StringVar(&configFile, "c", "config.json", "the config file")
	flag.BoolVar(&version, "V", false, "show version")
	flag.BoolVar(&testMode, "t", false, "test config")
}

func getVersion() string {
	return VERSION
}

func showVersion() {
	fmt.Println(getVersion())
	flag.Usage()
}
func main() {

	flag.Parse()

	if version {
		showVersion()
		return
	}

	if testMode {
		fmt.Println("config test ok")
		return
	}

	config, err := loadConfig(configFile)

	if err != nil {
		fmt.Errorf("load config error : %s", err.Error())
		os.Exit(2)
	}
	syncMgr := NewSyncManager()

	for _, plan := range config.Plans {
		syncMgr.AddPlan(plan)
	}
	syncMgr.Work()

}
