package plugin

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/vg006/vtest/internal/types"
	"github.com/vg006/vtest/internal/utils"
)

func Chi(r *chi.Mux, opts types.Options) {
	r.Get(opts.Path, func(w http.ResponseWriter, r *http.Request) {
		var routes []types.Route
		rc := chi.RouteContext(r.Context())
		err := chi.Walk(rc.Routes, func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
			if utils.Contains(opts.ExcludeMethods, method) {
				return nil
			}
			routes = append(routes, types.Route{Method: method, Endpoint: route})
			return nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(routes)
	})
}
