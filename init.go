package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	// "e.coding.net/xverse-git/public/go_common/netutil"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"gopkg.in/yaml.v2"
)

const (
	defaultLogFile = "./logs/app.log"
)

var (
	confFile = "configs/config.yml"
	sugar    = &zap.SugaredLogger{} // skip 1
	sugar0   = &zap.SugaredLogger{} // skip 0
	logTag   = &LogTag{}
)

type LogTag struct {
	TaskID      string
	MachineAddr string
}

type LogConfig struct {
	LogLevel string `json:"logLevel"`
	LogFile  string `json:"logFile"`
}

func Sugar() *zap.SugaredLogger {
	return sugar
}
func init() {
	Init()
}

func main() {
	Init()
	return
}

// func getMachineAddr() string {
// 	localIP, err := netutil.GetLocalIP()
// 	if err != nil {
// 		err = fmt.Errorf("fatal while init log: %v", err)
// 		panic(err)
// 	}
// 	if sshdPort := os.Getenv("sshd_port"); sshdPort != "" {
// 		return localIP + ":" + sshdPort
// 	}
// 	return localIP
// }

func Init() {
	// init LogTag
	// logTag = &LogTag{
	// 	MachineAddr: getMachineAddr(),
	// 	TaskID:      os.Getenv("RLPLATFORM_TASK_ID"),
	// }
	if cf := os.Getenv("LoggerConfFile"); cf != "" {
		fmt.Printf("change log conf file to : %v\n", cf)
		confFile = cf
	}

	// init log
	file, err := ioutil.ReadFile(confFile)
	if err != nil {
		//panic(err)  update by carbinlin 2021.04.12
	}
	yamlString := string(file)

	lcfg := &LogConfig{}
	err = yaml.Unmarshal([]byte(yamlString), lcfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if lcfg.LogFile == "" {
		lcfg.LogFile = defaultLogFile
	}
	if lcfg.LogFile == "pidlog" {
		lcfg.LogFile = fmt.Sprintf("logs/app_%v.log", os.Getpid())
	}
	if lcfg.LogLevel == "" {
		lcfg.LogLevel = zap.DebugLevel.String()
	}
	var l zapcore.Level
	err = l.UnmarshalText([]byte(lcfg.LogLevel))
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   lcfg.LogFile,
		MaxSize:    20, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})

	zconf := zapcore.EncoderConfig{
		TimeKey:       "ts",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		//		EncodeTime:     LoggerTimeEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zconf),
		w,
		l,
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	logger0 := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(0))

	sugar = logger.Sugar()
	sugar0 = logger0.Sugar()
}
