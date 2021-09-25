package cmd

import (
	"fiber/internal/router"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)
 
var serveCmd = &cobra.Command{ 
	Use:   "serve", 
	Short: "Run Booking service",
	Run: func(cmd *cobra.Command, args []string) { 
		StartService() 
	}, 
}

func StartService() {
	server := router.New()

	log.Info("Service is running")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != err {
		log.Error("ENV PORT failed")
	}

	server.Start(port)
} 
