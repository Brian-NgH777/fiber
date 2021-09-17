package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"

	"github.com/spf13/cobra"
)

func initService() {
	log.Info("Init Service!")
} 
 
func init() { 
	cobra.OnInitialize(initService) 
} 
 
// RootCmd . 
var RootCmd = &cobra.Command{ 
	Short: "Compile many function of Booking",
} 
 
// Execute ... 
func Execute() { 
	RootCmd.AddCommand(versionCmd) 
	RootCmd.AddCommand(serveCmd) 
	if err := RootCmd.Execute(); err != nil {
		log.Error(fmt.Sprintf("failed to start program, err: %v", err) )
		os.Exit(-1) 
	} 
} 
