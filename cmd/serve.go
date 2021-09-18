package cmd

import (
	"fiber/internal/router"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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
	server.Start(5000)
} 
