package CoreManage

import (
	"github.com/PandaSteve/framework-lib/src/Lib/HttpService"
	"github.com/PandaSteve/framework-lib/src/Lib/Message"
	"github.com/PandaSteve/framework-lib/src/Lib/MySql"
)

type Config struct {
	DataBase     MySqlLib.Config       `yaml:"data_base"`
	HttpService  HttpServiceLib.Config `yaml:"http_service"`
	Service      MessageLib.Config     `yaml:"service"`
	Source       MessageLib.Config     `yaml:"source"`
	Gateway      MessageLib.Config     `yaml:"gateway"`
}
