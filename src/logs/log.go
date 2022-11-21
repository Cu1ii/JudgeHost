package logs

import (
	"bytes"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"path"
	"time"
)

type mineFomatter struct{}

func (m *mineFomatter) Format(entry *logrus.Entry) ([]byte, error) {
	return nil, nil
}

func init() {
	gin.DefaultWriter = os.Stdout
	initRuntimeLog()
	initWebLog()
}

var logger *logrus.Logger

func initRuntimeLog() {
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&nested.Formatter{
		TimestampFormat: time.RFC3339,
	})
	writer1 := &bytes.Buffer{}
	writer2 := os.Stdout
	writer3, err := os.OpenFile("log/runtime.logs", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("create file log/runtime.logs failed: %v", err)
	}
	logrus.SetOutput(io.MultiWriter(writer1, writer2, writer3))
	logrus.Info("\truntime log start")
}

func initWebLog() {
	// 获得实例
	logger = logrus.New()
	logger.SetReportCaller(true)
	var (
		logFilePath = "log/" //文件存储路径
		logFileName = "system.logs"
	)
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	// 写入文件
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		log.Fatal("打开/写入文件失败", err)
	}
	// 日志级别
	logger.SetLevel(logrus.DebugLevel)
	// 设置输出
	logger.Out = file
	// 设置 rotatelogs,实现文件分割
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.logs",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour), //以hour为单位的整数
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(1*time.Hour),
	)
	// hook机制的设置
	writerMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	//给logrus添加hook
	// hook 自己的 formatter 不需要再使用, 让 logger 调用自己的 formatter 处理日志即可
	logger.AddHook(lfshook.NewHook(writerMap, &mineFomatter{}))

	logger.SetFormatter(&nested.Formatter{
		// HideKeys:        true,
		FieldsOrder:     []string{"status_code", "client_ip", "req_method", "req_uri"},
		TimestampFormat: time.RFC3339,
		NoColors:        false,
	})
}

func LogMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		//请求方式
		method := c.Request.Method
		//请求路由
		reqUrl := c.Request.RequestURI
		//状态码
		statusCode := c.Writer.Status()
		//请求ip
		clientIP := c.ClientIP()
		// 打印日志
		logger.WithFields(logrus.Fields{
			"status_code": statusCode,
			"client_ip":   clientIP,
			"req_method":  method,
			"req_uri":     reqUrl,
		}).Info()
	}
}
