package monitor

import (
	"context"
	"github.com/danesparza/iot-wifi-setup/internal/network"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

// Service encapsulates API service operations
type Service struct {
	NM network.NetworkManagerService
}

func (service Service) NetworkStatus(ctx context.Context) {
	frequencystring, freqset := os.LookupEnv("NETWORK_MONITOR_FREQ")
	if freqset != true {
		frequencystring = "30s"
	}

	log.Info().Str("frequency", frequencystring).Msg("NetworkStatus check starting...")

	//	Parse the frequency string
	frequency, err := time.ParseDuration(frequencystring)
	if err != nil {
		log.Warn().Str("frequency", frequencystring).Msg("Problem parsing NETWORK_MONITOR_FREQ duration.  Using default of 30s.")
		frequency = 30 * time.Second
	}

	for {
		select {
		//	Execute it every so often
		case <-time.After(frequency):
			//	As we get a request on a channel ...
			log.Debug().Msg("Checking network status")
			//go CheckStatusAndHandleAPMode(ctx) // Launch the goroutine
		case <-ctx.Done():
			log.Info().Msg("NetworkStatus check stopping")
			return
		}
	}
}
