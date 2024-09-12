package newrelic

import (
	"fmt"
	"os"
	"time"
	
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
	
	"mhf-api/config"
)

type logrusAdapter struct {
	*logrus.Entry
}

func (l *logrusAdapter) Error(msg string, context map[string]interface{}) {
	l.WithFields(logrus.Fields(context)).Error(msg)
}

func (l *logrusAdapter) Warn(msg string, context map[string]interface{}) {
	l.WithFields(logrus.Fields(context)).Warn(msg)
}

func (l *logrusAdapter) Info(msg string, context map[string]interface{}) {
	l.WithFields(logrus.Fields(context)).Info(msg)
}

func (l *logrusAdapter) Debug(msg string, context map[string]interface{}) {
	l.WithFields(logrus.Fields(context)).Debug(msg)
}

func (l *logrusAdapter) DebugEnabled() bool {
	return l.Logger.IsLevelEnabled(logrus.DebugLevel)
}

func InitNewRelic() *newrelic.Application {
	log := logrus.New()
	log.Out = os.Stdout
	log.SetLevel(logrus.InfoLevel)
	
	logEntry := logrus.NewEntry(log)
	logger := &logrusAdapter{logEntry}

	log.Info("Trying to init NewRelic Application")
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(config.GlobalConfig.NewRelic.AppName),
		newrelic.ConfigLicense(config.GlobalConfig.NewRelic.License),
		newrelic.ConfigLogger(logger),
	)
	if err != nil {
		log.Error("Error on 'InitNewRelic' New Relic:", err)
	}
	if app == nil {
		log.Error("New Relic application initialization failed")
	}

	err = app.WaitForConnection(10 * time.Second)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to connect to New Relic: %v", err))
	}
	log.Info("New Relic application connected")
	return app
}
