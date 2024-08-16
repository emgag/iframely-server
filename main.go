package main

import (
	"context"
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	apiKey := os.Getenv("SERVER_API_KEY")
	iframelyKey := os.Getenv("IFRAMELY_API_KEY")

	if apiKey == "" || iframelyKey == "" {
		fmt.Println("Environment variables SERVER_API_KEY and IFRAMELY_API_KEY must be set")
		return
	}

	app := fiber.New()

	app.Get("/info", func(c *fiber.Ctx) error {
		target := c.Query("url")
		key := c.Query("api_key")

		if subtle.ConstantTimeCompare([]byte(key), []byte(apiKey)) == 0 {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid API key")
		}

		u, err := url.Parse("https://iframe.ly/api/oembed")
		query := u.Query()
		query.Add("api_key", iframelyKey)
		query.Add("iframe", "0")
		query.Add("url", target)
		u.RawQuery = query.Encode()

		ctx, cancel := context.WithTimeout(c.Context(), time.Duration(time.Second))
		defer cancel()

		req, err := http.NewRequest(http.MethodGet, u.String(), nil)
		req = req.WithContext(ctx)

		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("failed to fetch data: %s", resp.Status))
		}

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		var iframelyResp IframelyResponse
		err = json.Unmarshal(body, &iframelyResp)

		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(Response{
			Provider: strings.ToLower(iframelyResp.ProviderName),
			URL:      iframelyResp.URL,
			HTML:     iframelyResp.HTML,
			Image:    iframelyResp.ThumbnailURL,
			Error:    iframelyResp.Error,
		})
	})

	err := app.Listen(":3000")
	log.Fatal(err)
}
