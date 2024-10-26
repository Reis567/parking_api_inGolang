package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger *zap.Logger
)

func init() {

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // Nível em maiúsculas e colorido
		EncodeTime:     zapcore.ISO8601TimeEncoder,       // Formato de tempo ISO8601
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,       // Arquivo e linha de onde veio o log
	}

	logLevel := getLogLevel()
	logOutput := getLogOutput()

	var syncer zapcore.WriteSyncer
	if logOutput == "stdout" {
		syncer = zapcore.AddSync(os.Stdout)
	} else if logOutput == "stderr" {
		syncer = zapcore.AddSync(os.Stderr)
	} else {
		file, err := os.OpenFile(logOutput, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic("Falha ao abrir o arquivo de log: " + err.Error())
		}
		syncer = zapcore.AddSync(file)
	}

	// Configurar o core do zap
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig), // Encoder para o console
		zapcore.NewMultiWriteSyncer(syncer),      // Output para console ou arquivo
		logLevel,                                 // Nível de log definido
	)

	// Criar o logger com as configurações especificadas
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	defer Logger.Sync() // Sincronizar antes de sair
}

// Info registra uma mensagem de nível informativo
func Info(message string, fields ...zap.Field) {
	Logger.Info(message, fields...)
	Logger.Sync() // Garante que a mensagem seja enviada imediatamente
}

// Error registra uma mensagem de nível de erro
func Error(message string, fields ...zap.Field) {
	Logger.Error(message, fields...)
	Logger.Sync() // Garante que a mensagem seja enviada imediatamente
}

// GetLogger retorna a instância do logger
func GetLogger() *zap.Logger {
	return Logger
}

// getLogOutput determina o destino de saída do log a partir da variável de ambiente LOG_OUTPUT
func getLogOutput() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv("LOG_OUTPUT")))
	if output == "" {
		return "stdout"
	}
	return output
}

// getLogLevel determina o nível de log a partir da variável de ambiente LOG_LEVEL
func getLogLevel() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv("LOG_LEVEL"))) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn", "warning":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel // Nível padrão
	}
}
