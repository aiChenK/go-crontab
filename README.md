# go-crontab

go语言编写的定时任务脚本，可用于容器镜像内

## 使用

```bash
./go-crontab
    -c  string run app use config with -c=xxx.conf (default "./crontab.conf")
    -d  run app as a daemon with -d=true

# example
# ./go-crontab -d=true -c=/tmp/crontab.conf
```

## 配置说明

crontab.conf文件内容按行读取，使用cron组件，支持#注释
具体用法可参考：<https://pkg.go.dev/github.com/robfig/cron>
示例：

```text
0 */10 * * * * test.sh
@every 10m test.sh
@hourly test.sh
```
