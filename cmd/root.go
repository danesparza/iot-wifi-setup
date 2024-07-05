package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "iot-wifi-setup",
	Short: "Simple wifi configuration service for an IoT app",
	Long:  `Present a wifi AP to get connected to local wifi network and then hand-off to your app`,
}

var (
	cfgFile               string
	problemWithConfigFile bool
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/iotwifisetup.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(home)      // adding home directory as first search path
		viper.AddConfigPath(".")       // also look in the working directory
		viper.SetConfigName("fxaudio") // name the config file (without extension)
	}

	viper.AutomaticEnv() // read in environment variables that match

	//	Set our defaults
	viper.SetDefault("datastore.system", filepath.Join(home, "fxaudio", "db", "fxaudio.db"))
	viper.SetDefault("datastore.migrationsource", "./scripts/sqlite/migrations")
	viper.SetDefault("upload.path", path.Join(home, "fxaudio", "uploads"))
	viper.SetDefault("upload.bytelimit", 15*1024*1024) // 15MB
	viper.SetDefault("server.port", 3070)
	viper.SetDefault("server.allowed-origins", "*")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		problemWithConfigFile = true
	}
}
