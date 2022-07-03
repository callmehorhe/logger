# logger

## Installation
```
go get -u github.com/callmehorhe/logger
```

## Quick start
```sh
l := logger.New()
l.WithFields(logger.Data{
	"url": "hw.com",
}).Error("oshibka(((")
```
Output
```sh
time="2022-07-04T01:59:20+03:00" level=error msg=oshibka((( url=hw.com
```
-----
```sh
lvl, _ := logger.ParseLevel("debug")
logger.SetLevel(lvl)
logger.Info("msg")
logger.Debug("dbg")
logger.Error("err")
logger.Warn("warn")
```
Output
```sh
time="2022-07-04T02:42:39+03:00" level=info msg=msg
time="2022-07-04T02:42:40+03:00" level=debug msg=dbg
time="2022-07-04T02:42:40+03:00" level=error msg=err
time="2022-07-04T02:42:40+03:00" level=warning msg=warn
```

## Latency
```sh
logger.WithLatency(3600000000, time.Second).Info("msg")
```
Output
```sh
time="2022-07-04T02:24:47+03:00" level=info msg=msg latency=3.6
```
