package glog

import (
	"github.com/heqiaomu/gtools/gos"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var l *zap.Logger

func Logger() {
	//l = CreateZapFactory(ZapLogHandler)
}

func (cfg *Config) CreateZapFactory(entry func(zapcore.Entry) error) *zap.Logger {

	// 获取程序所处的模式：  开发调试 、 生产
	//appDebug := config.GetAppDebug()
	appDebug := true
	if appDebug == true {
		if logger, err := zap.NewDevelopment(zap.Hooks(entry)); err == nil {
			return logger
		} else {
			log.Fatal("创建zap日志包失败，详情：" + err.Error())
		}
	}

	// 以下才是 非调试（生产）模式所需要的代码
	encoderConfig := zap.NewProductionEncoderConfig()

	var recordTimeFormat string
	switch cfg.TimePrecision {
	case "second":
		recordTimeFormat = "2006-01-02 15:04:05"
	case "millisecond":
		recordTimeFormat = "2006-01-02 15:04:05.000"
	default:
		recordTimeFormat = "2006-01-02 15:04:05"

	}
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(recordTimeFormat))
	}
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.TimeKey = "created_at" // 生成json格式日志的时间键字段，默认为 ts,修改以后方便日志导入到 ELK 服务器

	var encoder zapcore.Encoder
	switch cfg.TextFormat {
	case "console":
		encoder = zapcore.NewConsoleEncoder(encoderConfig) // 普通模式
	case "json":
		encoder = zapcore.NewJSONEncoder(encoderConfig) // json格式
	default:
		encoder = zapcore.NewConsoleEncoder(encoderConfig) // 普通模式
	}

	//写入器
	fileName := gos.HomeDir() + cfg.GoSkeletonLogName
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,       //日志文件的位置
		MaxSize:    cfg.MaxSize,    //在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: cfg.MaxBackups, //保留旧文件的最大个数
		MaxAge:     cfg.MaxAge,     //保留旧文件的最大天数
		Compress:   cfg.Compress,   //是否压缩/归档旧文件
	}
	writer := zapcore.AddSync(lumberJackLogger)
	// 开始初始化zap日志核心参数，
	//参数一：编码器
	//参数二：写入器
	//参数三：参数级别，debug级别支持后续调用的所有函数写日志，如果是 fatal 高级别，则级别>=fatal 才可以写日志
	zapCore := zapcore.NewCore(encoder, writer, zap.InfoLevel)
	return zap.New(zapCore, zap.AddCaller(), zap.Hooks(entry), zap.AddStacktrace(zap.WarnLevel))
}

type Level string

const (
	DebugLevel Level = "debug"
	InfoLevel  Level = "info"
	WarnLevel  Level = "warn"
	ErrorLevel Level = "error"
	PanicLevel Level = "panic"
	FatalLevel Level = "fatal"
)

func Log(format string, fields ...zap.Field) {
	l.Info(format, fields...)
}
func Debug(format string, fields ...zap.Field) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Debug(format, fields...)
	// TODO 发往日志系统
}
func Debugf(template string, args ...interface{}) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Debugf(template, args...)
	// TODO 发往日志系统
}
func Debugw(template string, fields ...zap.Field) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Debug(template, fields...)
	// TODO 发往日志系统
}

func Info(format string, fields ...zap.Field) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Info(format, fields...)
	// TODO 发往日志系统
}
func Infof(template string, args ...interface{}) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Infof(template, args...)
	// TODO 发往日志系统
}
func Infow(template string, keysAndValues ...interface{}) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Infow(template, keysAndValues...)
	// TODO 发往日志系统
}

func Warn(format string, fields ...zap.Field) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Warn(format, fields...)
	// TODO 发往日志系统
}
func Warnf(template string, args ...interface{}) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Warnf(template, args...)
	// TODO 发往日志系统
}
func Warnw(template string, keysAndValues ...interface{}) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Warnw(template, keysAndValues...)
	// TODO 发往日志系统
}

func Error(format string, fields ...zap.Field) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Error(format, fields...)
	// TODO 发往日志系统
}
func Errorf(template string, args ...interface{}) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Errorf(template, args...)
	// TODO 发往日志系统
}
func Errorw(template string, keysAndValues ...interface{}) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Errorw(template, keysAndValues...)
	// TODO 发往日志系统
}

func Panic(format string, fields ...zap.Field) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Panic(format, fields...)
	// TODO 发往日志系统
}
func Panicf(template string, args ...interface{}) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Panicf(template, args...)
	// TODO 发往日志系统
}
func Panicw(template string, keysAndValues ...interface{}) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Panicw(template, keysAndValues...)
	// TODO 发往日志系统
}

func Falal(format string, fields ...zap.Field) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Fatal(format, fields...)
	// TODO 发往日志系统
}
func Fatalf(template string, args ...interface{}) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Fatalf(template, args...)
	// TODO 发往日志系统
}
func Fatalw(template string, keysAndValues ...interface{}) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Fatalw(template, keysAndValues...)
	// TODO 发往日志系统
}

func stackTrace() (funcName, caller string) {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, false)]
	stacks := strings.Split(string(buf), "\n")

	funcName = stacks[5]
	funcName = funcName[strings.LastIndex(funcName, "/")+1 : strings.LastIndex(funcName, "(")]
	callers := strings.Split(stacks[6][1:], " ")

	return funcName, callers[0]
}

func constructFieldMap(level, funcName, caller string, fields []zap.Field) map[string]interface{} {
	m := make(map[string]interface{})
	m["level"] = level
	m["funcName"] = funcName
	m["caller"] = caller
	for _, f := range fields {
		m[f.Key] = f.String
	}
	return m
}

func constructArrayMap(level, funcName, caller string, fields []interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	m["level"] = level
	m["funcName"] = funcName
	m["caller"] = caller
	for i, f := range fields {
		m[strconv.Itoa(i)] = f
	}
	return m
}

// Sugar return zap SugaredLogger instance
func Sugar() *zap.SugaredLogger {
	return l.Sugar()
}

// Named return zap instance
func Named(s string) *zap.Logger {
	return l.Named(s)
}

// WithOptions log with option
func WithOptions(opts ...zap.Option) *zap.Logger {
	return l.WithOptions(opts...)
}

// With log with field
func With(fields ...zap.Field) *zap.Logger {
	return l.With(fields...)
}

// Check level check
func Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry {
	return l.Check(lvl, msg)
}

// Sync log sync
func Sync() error {
	return l.Sync()
}

// Core return log core
func Core() zapcore.Core {
	return l.Core()
}
