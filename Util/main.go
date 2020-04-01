package main

import "github.com/Altair-cyx/Util/loggerutil"

func main() {
	log := loggerutil.NewConsoleLog()
	log.All("All日志")
	log.Debug("Debug日志")
	log.Trace("Trace日志")
	log.Info("Info日志")
	log.Warn("Warn日志")
	log.Error("Error日志")
	log.Fatal("Fatal日志")
	log.Off("Off日志")
}
