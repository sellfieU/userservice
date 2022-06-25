package main

import (
	"fmt"
	"github.com/sellfie/usermanagerservice/src/logrotate"
	"github.com/sellfie/usermanagerservice/src/server"
	"io"
	"os"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var (
	setupLog = ctrl.Log.WithName("setup")
	port     = "3550"
)

func main() {
	// Set log rotation
	logFile, err := logrotate.LogFile()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer func() {
		_ = logFile.Close()
	}()
	logWriter := io.MultiWriter(logFile, os.Stdout)
	ctrl.SetLogger(zap.New(zap.UseDevMode(true), zap.WriteTo(logWriter)))
	if err := logrotate.StartRotate("0 0 1 * * ?"); err != nil {
		setupLog.Error(err, "")
		os.Exit(1)
	}
	// Set ports
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	srv := server.New()
	srv.Start(port)
}
