package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/vg006/vtest/internal/types"
)

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

	securityHeaders := []string{
		"Strict-Transport-Security", // HSTS
		"Content-Security-Policy",   // CSP
		"X-Frame-Options",           // Prevent clickjacking
		"X-Content-Type-Options",    // Prevent MIME sniffing
		"X-XSS-Protection",          // Legacy XSS filter
		"Referrer-Policy",           // Control referrer info
		"Permissions-Policy",        // Control browser features
	}

	for _, headerName := range securityHeaders {
		if value := resp.Header.Get(headerName); value != "" {
			res.Headers = append(res.Headers, types.Header{
				Key:   headerName,
				Value: value,
			})
		} else {
			res.Headers = append(res.Headers, types.Header{
				Key:   headerName,
				Value: "Not set",
			})
		}
	}
	fmt.Println("Response headers:", res.Headers)
	return res, nil
}
