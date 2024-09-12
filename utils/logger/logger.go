package logger

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/nrlogrus"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Format   string
	FilePath string
}

type Logger struct {
	*logrus.Logger
}

func NewLogger(newRelicApp *newrelic.Application, config Config, scope string, contexts ...string) *Logger {
	environment := os.Getenv("ENVIRONMENT")
	var context string
	if len(contexts) > 0 {
		context = contexts[0]
	}

	message := "mhf-api:" + scope
	if context != "" {
		message += ":" + context
	}

	logger := createLogger(newRelicApp, config, message, environment)

	return &logger
}

func createLogger(newRelicApp *newrelic.Application, config Config, message string, environement string) Logger {
	logrusLogger := logrus.New()
	var formatter logrus.Formatter

	if environement == "dev" {
		formatter = &logrus.TextFormatter{
			ForceColors:     true,
			FullTimestamp:   true,
			TimestampFormat: convertToGoTimeFormat(config.Format),
		}
	}
	if environement == "prod" {
		formatter = &logrus.JSONFormatter{
			TimestampFormat: convertToGoTimeFormat(config.Format),
		}
	}

	nrlogrusFormatter := nrlogrus.NewFormatter(newRelicApp, formatter)
	logrusLogger.SetFormatter(nrlogrusFormatter)

	output := getOutput(
		message,
		config.FilePath != "",
		config.FilePath,
	)

	logrusLogger.SetOutput(output)

	return Logger{logrusLogger}
}

func convertToGoTimeFormat(format string) string {
	replacements := map[string]string{
		"YYYY": "2006", "MM": "01", "DD": "02",
		"hh": "15", "mm": "04", "ss": "05", "ms": "000",
	}
	for old, new := range replacements {
		format = strings.ReplaceAll(format, old, new)
	}
	return format
}

func getOutput(message string, hasPath bool, path string) io.Writer {
	fmt.Printf("getOutput: %s", message)

	fileName := "merged_logs"

	if !hasPath {
		return os.Stdout
	}

	directoryPath := path + "/" + time.Now().Format(convertToGoTimeFormat("YYYY-DD-MM"))
	filePath := fmt.Sprintf("%s/%s.log", directoryPath, fileName)

	if _, err := os.Stat(filePath); err == nil {
		file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			fmt.Printf("cannot open log file: %v\n", err)
			return nil
		}
		return io.MultiWriter(os.Stdout, file)
	}

	err := os.MkdirAll(directoryPath, os.ModePerm)
	if err != nil {
		fmt.Printf("cannot create log directory %s: %v\n", directoryPath, err)
		return nil
	}

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("cannot create log file: %v\n", err)
		return nil
	}
	return io.MultiWriter(os.Stdout, file)
}

func (l *Logger) Float32(label string, values []float32) []logrus.Fields {
	formattedValues := make([]string, len(values))
	for i, v := range values {
		formattedValues[i] = fmt.Sprintf("%.2f", v)
	}
	return []logrus.Fields{{label: strings.Join(formattedValues, ", ")}}
}

func (l *Logger) BytesToString(key string, payload []byte) logrus.Fields {
	return logrus.Fields{key: payload}
}

func (l *Logger) ReadBytes(data []byte) logrus.Fields {
	isPrintableASCII := true
	for _, b := range data {
		if !unicode.IsPrint(rune(b)) {
			isPrintableASCII = false
			break
		}
	}
	if isPrintableASCII {
		return logrus.Fields{"string": string(data)}
	}

	hexStr := hex.EncodeToString(data)
	if len(hexStr)%2 == 0 {
		_, err := hex.DecodeString(hexStr)
		if err == nil {
			return logrus.Fields{"hex": hexStr}
		}
	}

	intVal, err := strconv.Atoi(string(data))
	if err == nil {
		return logrus.Fields{"int": intVal}
	}

	return logrus.Fields{"unknown": hexStr}
}

func (l *Logger) WithField(key string, value interface{}) *logrus.Entry {
	return logrus.WithField(key, value)
}

func (l *Logger) WithContext(context string, fields ...string) *logrus.Entry {
	commonFields := make(logrus.Fields)
	commonFields["context"] = context

	for _, field := range fields {
		parts := strings.SplitN(field, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			commonFields[key] = value
		} else {
			logrus.Warnf("Invalid field format: %s. Skipping.", field)
		}
	}

	return logrus.WithFields(commonFields)
}
