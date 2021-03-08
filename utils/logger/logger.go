package logger

import (
	"github.com/mkruczek/user-store/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

type Logger interface {
	Errorf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Warnf(format string, v ...interface{})
}

type LOG struct {
	logger *zap.SugaredLogger
}

func New(cfg *config.Config) *LOG {
	logConf := zap.Config{
		Level:       zap.NewAtomicLevelAt(cfg.Logging.Level),
		Encoding:    cfg.Logging.Encoding,
		OutputPaths: cfg.Logging.OutputPaths,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   cfg.Logging.EncoderConfig.MessageKey,
			LevelKey:     cfg.Logging.EncoderConfig.LevelKey,
			TimeKey:      cfg.Logging.EncoderConfig.TimeKey,
			EncodeTime:   zapcore.EpochTimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	zapLog, err := logConf.Build()
	if err != nil {
		log.Fatal(err)
	}

	return &LOG{logger: zapLog.Sugar()}
}

func (l *LOG) Errorf(format string, v ...interface{}) {
	l.logger.Errorf(format, v...)
}

func (l *LOG) Warnf(format string, v ...interface{}) {
	l.logger.Warnf(format, v...)
}

func (l *LOG) Infof(format string, v ...interface{}) {
	l.logger.Infof(format, v...)
}

func (l *LOG) Debugf(format string, v ...interface{}) {
	l.logger.Debugf(format, v...)
}
