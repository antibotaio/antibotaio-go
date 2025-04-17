package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/antibotaio/antibotaio-go"
)

func main() {
	session := antibotaio.NewSession("your_api_key")

	session.WithClient(http.DefaultClient)

	body := base64.StdEncoding.EncodeToString([]byte("example body"))

	input := &antibotaio.SyncInput{
		Href:      "https://example.com",
		Website:   "example.com",
		Body:      body,
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		Language:  "en-US",
		Timezone:  "America/New_York",
	}

	response, err := session.SolveSync(input)
	if err != nil {
		log.Fatalf("failed to solve sync: %v", err)
	}

	fmt.Printf("sync response: %v", response)

	differentBody := base64.StdEncoding.EncodeToString([]byte("different body"))

	widgetInput := &antibotaio.NuDataWidgetTask{
		Body: differentBody,
	}

	widgetResponse, err := session.SolveWidget(widgetInput, response.ID)
	if err != nil {
		log.Fatalf("failed to solve widget: %v", err)
	}

	fmt.Printf("widget response: %v", widgetResponse)
}
