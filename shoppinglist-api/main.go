package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	// Test credentials for demo purposes
	// ideally it should utilize credential service to authenticate
	testUsername = "test"
	testPassword = "test"
)

// BasicAuthMiddleware validates basic auth credentials
func BasicAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//check request method
		if c.Request().Method == "OPTIONS" {
			return next(c)
		}

		// Get Authorization header
		auth := c.Request().Header.Get("Authorization")
		if auth == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Please provide basic auth credentials")
		}

		// Parse Basic Auth credentials
		const prefix = "Basic "
		if !strings.HasPrefix(auth, prefix) {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authentication method")
		}

		payload, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid basic auth format")
		}

		basicauth_pair := strings.SplitN(string(payload), ":", 2)
		if len(basicauth_pair) != 2 {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid basic auth format")
		}

		username := basicauth_pair[0]
		password := basicauth_pair[1]

		// Validate credentials
		if username != testUsername || password != testPassword {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
		}

		return next(c)
	}
}

// ProxyHandler forwards requests to the target API
func ProxyHandler(c echo.Context) error {
	// Target API base URL - replace with your target API
	targetBaseURL := os.Getenv("SHOPPING_LIST_SERVICE_URI")

	// Get the original request method and path
	method := c.Request().Method
	path := c.Request().URL.Path

	// Create new request to target API
	targetURL := targetBaseURL + path
	req, err := http.NewRequest(method, targetURL, c.Request().Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error creating proxy request")
	}

	// Copy original headers
	for key, values := range c.Request().Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	// Copy query parameters
	req.URL.RawQuery = c.Request().URL.RawQuery

	// Send request to target API
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, "Error forwarding request")
	}
	defer resp.Body.Close()

	// Copy response headers
	for key, values := range resp.Header {
		for _, value := range values {
			c.Response().Header().Add(key, value)
		}
	}

	// Copy response status
	c.Response().WriteHeader(resp.StatusCode)

	// Copy response body
	_, err = io.Copy(c.Response().Writer, resp.Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error copying response")
	}

	return nil
}

func main() {
	// Create new Echo instance
	e := echo.New()

	//static file server
	//e.Static("/", "public")
	e.File("/", "public/index.html")



	// Disable CORS for development purposes
	e.Use(middleware.CORS())

	corsConfig := middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}

	e.Use(middleware.CORSWithConfig(corsConfig))

	// Add logging middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Add basic auth middleware to all routes
	//e.Use(BasicAuthMiddleware)

	// Define routes - this will match any path and method
	e.Any("/*", ProxyHandler)

	// Start server
	port := ":8080"
	fmt.Printf("Starting proxy server on %s\n", port)
	e.Logger.Fatal(e.Start(port))
}
