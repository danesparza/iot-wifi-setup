package cmd

import (
	"context"
	"fmt"
	"github.com/danesparza/iot-wifi-setup/api"
	_ "github.com/danesparza/iot-wifi-setup/docs" // swagger docs location
	"github.com/danesparza/iot-wifi-setup/internal/network"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the IoT wifi setup service",
	Long: `Starts an access point and configuration UI 
if there is no active Wifi connection configured.  

Otherwise this simply starts the utility API service to show 
wifi status and allow for resetting the wifi config`,
	Run: start,
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func start(cmd *cobra.Command, args []string) {

	//	If we have a config file, report it:
	if viper.ConfigFileUsed() != "" {
		log.Debug().Str("configFile", viper.ConfigFileUsed()).Msg("Using config file")
	} else {
		log.Debug().Msg("No config file found")
	}

	//	Create an api service object
	apiService := api.Service{
		StartTime: time.Now(),
		NM:        network.NewNetworkManagerService(),
	}

	//	Trap program exit appropriately
	ctx, cancel := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 2)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	go handleSignals(ctx, sigs, cancel)

	//	Set up the API routes
	r := api.NewRouter(apiService)

	//	Format the bound interface:
	formattedServerPort := fmt.Sprintf(":%v", viper.GetString("server.port"))

	//	Start the service and display how to access it
	log.Info().Str("server", formattedServerPort).Msg("Started REST service")
	log.Err(http.ListenAndServe(formattedServerPort, r)).Msg("HTTP API service error")
}

func handleSignals(ctx context.Context, sigs <-chan os.Signal, cancel context.CancelFunc) {
	select {
	case <-ctx.Done():
	case sig := <-sigs:
		switch sig {
		case os.Interrupt:
			log.Info().Msg("SIGINT")
		case syscall.SIGTERM:
			log.Info().Msg("SIGTERM")
		}

		log.Info().Msg("Shutting down ...")
		cancel()
		os.Exit(0)
	}
}
