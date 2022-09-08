package bootstrap

import "flag"

var ConfigFile *string

func init() {
	ConfigFile = flag.String("c", "./crontab.conf", "run app use config with -c=xxx.conf")
}
