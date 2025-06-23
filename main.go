package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Alonza0314/free-ran-ue/gnb"
	"github.com/Alonza0314/free-ran-ue/logger"
	"github.com/Alonza0314/free-ran-ue/model"
	"github.com/Alonza0314/free-ran-ue/util"
)

func main() {
	gnbConfig := model.GnbConfig{}
	if err := util.LoadFromYaml("config/gnb.yaml", &gnbConfig); err != nil {
		panic(err)
	}

	logger, err := logger.NewLogger(gnbConfig.Logger.Level)
	if err != nil {
		panic(err)
	}

	gnb := gnb.NewGnb(&gnbConfig, logger)
	gnb.Start()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	gnb.Stop()
}
