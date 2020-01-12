package conf

import (
	"github.com/youzeliang/gopractice/base"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
	"runtime"
)

//支持数组，工具

type TConf struct {
	LogLevel string `yaml:"loglevel"`
	Redis    struct {
		Addr string `yaml:"addr"`
	}
	Mysql struct {
		Addr         string `yaml:"addr"`
		User         string `yaml:"user"`
		PassWord     string `yaml:"password"`
		DataBase     string `yaml:"database"`
		MaxIdleConns int    `yaml:"maxidleconns"`
		MaxOpenConns int    `yaml:"maxopenconns"`
	}
	Port string
	Nsq  struct {
		Addr       string `yaml:"addr"`
		LookupAddr string `yaml:"lookupaddr"`
		Topic      string `yaml:"topic"`
		InFlight   int    `yaml:"inflight"`
	}
	Http struct {
		Timeout             int `yaml:"timeout"`
		MaxIdleConns        int `yaml:"maxidleconns"`
		MaxIdleConnsPerHost int `yaml:"maxidleconnsperhost"`
	}

	Warehouse struct {
		DispatcherTimeOut int    `yaml:"dispatcherTimeOut"`
		LoopAssignOrder   string `yaml:"loopAssignOrder"`
		LoopCallContainer string `yaml:"loopCallContainer"`
	}

	Elastic struct {
		Addr string
		User string
		Pass string
	}
}

var Conf TConf

func init() {
	yamlFile, err := ioutil.ReadFile(getCurrentPath() + "/conf.yaml")
	if err != nil {
		base.PanicfLogger(nil, "yamlfile get error: %v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		base.PanicfLogger(nil, "yaml unmarshal error: %v", err)
	}
}

func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
