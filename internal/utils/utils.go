package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/vg006/vtest/internal/types"
)

var securityHeaders = []types.SecurityHeaderInfo{
	{
		Name:       "Strict-Transport-Security",
		Importance: "Enforces HTTPS connections and prevents downgrade attacks",
		Level:      "Critical",
	},
	{
		Name:       "Content-Security-Policy",
		Importance: "Prevents XSS and other code injection attacks",
		Level:      "Critical",
	},
	{
		Name:       "X-Frame-Options",
		Importance: "Prevents clickjacking attacks by controlling frame embedding",
		Level:      "High",
	},
	{
		Name:       "X-Content-Type-Options",
		Importance: "Prevents MIME-type sniffing security exploits",
		Level:      "Medium",
	},
	{
		Name:       "X-XSS-Protection",
		Importance: "Legacy protection, modern browsers use CSP instead",
		Level:      "Low",
	},
	{
		Name:       "Referrer-Policy",
		Importance: "Controls how much referrer information is included with requests",
		Level:      "Medium",
	},
	{
		Name:       "Permissions-Policy",
		Importance: "Controls which browser features and APIs the site can use",
		Level:      "Medium",
	},
}

func Contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func GetDetails(req types.ReqSingleUrl) (types.ResSingleUrl, error) {
	var res types.ResSingleUrl
	url := "http://localhost:" + req.Port + req.Endpoint

	fmt.Println("URL:", url)
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	httpReq, err := http.NewRequest(req.Method, url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return res, err
	}

	resp, err := client.Do(httpReq)
	if err != nil {
		fmt.Println("Error executing request:", err)
		res.Headers = append(res.Headers, types.Header{
			Key:   "error",
			Value: err.Error(),
		})
		return res, err
	}
	defer resp.Body.Close()
	res.StatusCode = resp.StatusCode

	for _, header := range securityHeaders {
		if value := resp.Header.Get(header.Name); value != "" {
			res.Headers = append(res.Headers, types.Header{
				Key:      header.Name,
				Value:    value,
				Usage:    header.Importance,
				Level:    header.Level,
				Presence: true,
			})
		} else {
			res.Headers = append(res.Headers, types.Header{
				Key:      header.Name,
				Value:    "Not set",
				Usage:    header.Importance,
				Level:    header.Level,
				Presence: false,
			})
		}
	}
	fmt.Println("Response headers:", res.Headers)
	return res, nil
}

func FetchRoutes(req types.ReqSingleUrl) ([]types.Route, error) {
	var res []types.Route

	url := "http://localhost:" + req.Port + req.Endpoint
	fmt.Println("Fetching routes from:", url)

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	httpReq, err := http.NewRequest(req.Method, url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return res, err
	}

	resp, err := client.Do(httpReq)
	if err != nil {
		fmt.Println("Error executing request:", err)
		return res, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return res, err
	}

	if err := json.Unmarshal(body, &res); err != nil {
		fmt.Println("Error unmarshaling response JSON:", err)
		return res, err
	}

	for _, route := range res {
		fmt.Printf("Method: %s, Path: %s\n", route.Method, route.Path)
	}

	return res, nil
}
