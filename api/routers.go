package api

import (
	"github.com/danesparza/iot-wifi-setup/ui"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"net/http"
)

func NewRouter(apiService Service) http.Handler {
	//	Create a router and set up our REST endpoints...
	r := chi.NewRouter()

	//	Add middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5))
	r.Use(ApiVersionMiddleware)

	//	... including CORS middleware
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Route("/v1", func(r chi.Router) {
		//	Network status
		r.Route("/status", func(r chi.Router) {
			r.Get("/", apiService.GetNetworkStatus) // Get network status
		})

		//	Wifi APs
		r.Route("/aps", func(r chi.Router) {
			r.Get("/", apiService.ListAccessPoints) // Get all nearby access points
		})

		r.Route("/configure", func(r chi.Router) {
			r.Put("/apmode", apiService.StartAPMode)   // Turn on AP mode
			r.Put("/client", apiService.SetClientWifi) // Set client network information
		})
	})

	//	SWAGGER
	r.Mount("/swagger", httpSwagger.WrapHandler)

	//	Route for embedded UI
	r.Handle("/*", http.FileServer(http.FS(ui.EmbeddedFS)))

	return r
}
