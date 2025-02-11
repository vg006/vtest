package server

import (
	"fmt"
	"mime"
	"net/http"
	"path"

	"github.com/labstack/echo/v4"
	"github.com/vg006/vtest/internal/types"
	"github.com/vg006/vtest/internal/utils"
)

func HomeHandler(e echo.Context) error {
	return e.String(200, "Hello, World!")
}

func RoutesHandler(e echo.Context) error {
	var req types.ReqSingleUrl
	if err := e.Bind(&req); err != nil {
		fmt.Println("Error binding request:", err)
		return e.String(http.StatusBadRequest, "Invalid request")
	}

	res, err := utils.FetchRoutes(req)
	if err != nil {
		fmt.Println("Error fetching response:", err)
		return e.String(http.StatusInternalServerError, "Error fetching response")
	}
	fmt.Println("Response:", res)
	return e.JSON(http.StatusOK, res)
}

func DetailsHandler(c echo.Context) error {
	var req types.ReqSingleUrl
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request")
	}

	res, err := utils.GetDetails(req)
	if err != nil {
		fmt.Println("Error fetching response:", err)
		return c.String(http.StatusInternalServerError, "Error fetching response")
	}
	fmt.Println("Response:", res)
	return c.JSON(http.StatusOK, res)
}

func spaHandler(fs http.FileSystem) http.HandlerFunc {
	fileServer := http.FileServer(fs)
	return func(w http.ResponseWriter, r *http.Request) {
		f, err := fs.Open(r.URL.Path)
		if err != nil {
			r.URL.Path = "/200.html"
		} else {
			fi, err := f.Stat()
			f.Close()
			if err != nil || fi.IsDir() {
				r.URL.Path = "/200.html"
			}
		}

		ext := path.Ext(r.URL.Path)
		if ext != "" {
			if ct := mime.TypeByExtension(ext); ct != "" {
				w.Header().Set("Content-Type", ct)
			}
		}
		fileServer.ServeHTTP(w, r)
	}
}

func init() {
	// Ensure proper MIME types for JavaScript and CSS.
	// This is especially important if the system's MIME types are not set correctly.
	mime.AddExtensionType(".js", "application/javascript")
	mime.AddExtensionType(".css", "text/css")
}
