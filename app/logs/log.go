package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var SugarLogger *zap.SugaredLogger

func InitLogger() {
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	file, _ := os.OpenFile("./logs/log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writeSyncer := zapcore.AddSync(file)
	core := zapcore.NewCore(encoder, writeSyncer, zap.DebugLevel)

	logger := zap.New(core)
	SugarLogger = logger.Sugar()
}

func LogError(err error) {
	SugarLogger.Errorf("Error : %s", err)
}

func LogSuccess(message string) {
	SugarLogger.Infof("Message : %s", message)
}
