package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var version = "v0.0.1" 
 
// SetRevision inject version from git 
func SetRevision(r string) { 
	if len(r) > 0 { 
		version = fmt.Sprintf("%v-%v", version, r) 
	} 
	os.Setenv("version", version) 
} 
 
var versionCmd = &cobra.Command{ 
	Use:   "version", 
	Short: "Print the version number of service", 
	Run: func(cmd *cobra.Command, args []string) { 
		log.Info("Version -- %v\n", version)
	}, 
} 
