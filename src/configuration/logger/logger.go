package logger

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "os"
)

var (
    Logger *zap.Logger
)

func init() {
    // Configuração básica do `zapcore.EncoderConfig`
    encoderConfig := zapcore.EncoderConfig{
        TimeKey:        "timestamp",
        LevelKey:       "level",
        NameKey:        "logger",
        CallerKey:      "caller",
        MessageKey:     "message",
        StacktraceKey:  "stacktrace",
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.CapitalColorLevelEncoder, // Capitalize o nível (INFO, ERROR)
        EncodeTime:     zapcore.ISO8601TimeEncoder,       // Formato de tempo ISO8601
        EncodeDuration: zapcore.StringDurationEncoder,
        EncodeCaller:   zapcore.ShortCallerEncoder,       // Mostra o arquivo e linha de onde veio o log
    }

    // Configurando o `zapcore.Core` para console e arquivo
    core := zapcore.NewCore(
        zapcore.NewConsoleEncoder(encoderConfig),                      // Encoder para console
        zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),       // Output para console
        zapcore.DebugLevel,                                            // Nível mínimo de log
    )

    // Criação do logger
    Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

    // Configuração opcional: Redirecionar `zap.SugaredLogger` para facilitar o uso de logs formatados
    defer Logger.Sync()
}

// GetLogger retorna a instância do logger
func GetLogger() *zap.Logger {
    return Logger
}
