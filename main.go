package main

import (
	"bytes"
	"errors"
	"fmt"
	"go-crontab/pkg/bootstrap"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	_ "github.com/icattlecoder/godaemon"
	"github.com/robfig/cron"
)

func main() {
	//获取配置地址
	var configFile = *bootstrap.ConfigFile

	//读取配置
	jobs := ReadFile(configFile)
	if len(jobs) == 0 {
		log.Println("[ERROR] Config File Read Failed")
		return
	}
	log.Println("[INFO] Config File Read Success")

	//设置任务
	c := cron.New()
	for _, line := range jobs {
		cronStr, cmdStr, err := ParseCron(line)
		if err != nil {
			// log.Println("[ERROR]", err, ":", line)
			continue
		}
		c.AddFunc(cronStr, func() {
			log.Println("[INFO] Run:", cmdStr)
			Command(cmdStr)
		})
		log.Println("[INFO] AddFunc [", cronStr, "]:", cmdStr)
	}
	c.Start()
	select {}
}

// 解析定时脚本
func ParseCron(str string) (cronStr, cmdStr string, err error) {
	re := regexp.MustCompile("^((([\\d*-,? \\/]+){6} )|(@every [\\d\\w]+ )|(@\\w+ ))")
	cronStr = re.FindString(str)
	if len(cronStr) == 0 {
		err = errors.New("Parse cron failed")
		return
	}
	cmdStr = strings.Trim(strings.Replace(str, cronStr, "", 1), " ")
	if len(cmdStr) == 0 {
		err = errors.New("Parse cmd faild")
		return
	}
	cronStr = strings.Trim(cronStr, " ")
	return
}

// 读取文件信息
func ReadFile(path string) []string {
	fileHanle, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return []string{}
	}

	defer fileHanle.Close()

	readBytes, err := ioutil.ReadAll(fileHanle)
	if err != nil {
		return []string{}
	}

	results := strings.Split(string(readBytes), "\n")
	return results
}

func Command(cmd string) error {
	c := exec.Command("bash", "-c", cmd)
	output, err := c.CombinedOutput()
	fmt.Println(string(output))
	return err
}

func Exec(command string) error {
	in := bytes.NewBuffer(nil)
	cmd := exec.Command("/bin/bash")
	cmd.Stdin = in
	in.WriteString(command)
	in.WriteString("exit\n")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
