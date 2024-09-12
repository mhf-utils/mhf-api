package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"mhf-api/config"
	"mhf-api/server"
	"mhf-api/utils/ascii"
	"mhf-api/utils/logger"
	"mhf-api/utils/newrelic"
)

func main() {
	newRelicApp := newrelic.InitNewRelic()
	log         := logger.NewLogger(newRelicApp, logger.Config(config.GlobalConfig.Logger), "main")
	commit      := getCommit()

	fmt.Printf(ascii.ServerTitle, commit)
	server.Init(log, newRelicApp)
}

func getCommit() string {
	environment := os.Getenv("ENVIRONMENT")
	cmd         := exec.Command("git", "rev-parse", "--short", "HEAD")
	output, _   := cmd.Output()

	if environment == "dev" {
		environment = " " + environment
	}
	return strings.TrimSpace(string(output))
}
