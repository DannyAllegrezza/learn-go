package router

import (
  "github.com/go-chi/chi/v5"
  "github.com/go-chi/cors"
  "net/http"
  "HelloWorld/api"
  "HelloWorld/impl"
)

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler() http.Handler {
    serviceImpl := api.HelloWorldDelegate{
            ApiDelegate: &impl.ApiImpl{},
            RootDelegate: &impl.RootImpl{},
    }

  return RouterHandler(serviceImpl, chi.NewRouter())
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func RouterHandler(serviceImpl api.HelloWorldDelegate, r chi.Router) http.Handler {
    // This is here for local dev server development where the client consuming the API is NOT the same IP
    // and/or PORT as the server is running on. This allows local development from say, a web UI that runs
    // on a different port, accessing the API, so a reverse proxy is not needed.
    cors := cors.New(cors.Options{
        AllowedOrigins:     []string{"*"},
        AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
        AllowedHeaders:     []string{"Accept", "Authorization", "Content-Length", "Cache-Control", "Accept-Encoding", "Content-Type", "X-CSRF-Token"},
        AllowCredentials:   true,
        MaxAge:             300, // Maximum value not ignored by any of major browsers
        OptionsPassthrough: false,
        Debug:              true,
    })
    r.Use(cors.Handler)

    apiWrapper := api.ApiWrapper {
      ApiDelegate: serviceImpl.ApiDelegate,
    }

    r.Group(func(r chi.Router) {
      r.Get("/api/volunteers", apiWrapper.GetVolunteers)
    })

    rootWrapper := api.RootWrapper {
      RootDelegate: serviceImpl.RootDelegate,
    }

    r.Group(func(r chi.Router) {
      r.Get("/", rootWrapper.Get)
    })

    return r
}
