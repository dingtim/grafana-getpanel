// Get a panel as png from grafana
// in grafana go to the panel and click on "share image" / "get direct link rendered image" to get the proper url
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cavaliercoder/grab"
	"github.com/pkg/errors"
)

func main() {
	// example url
	url := `http://localhost:3000/render/d-solo/6cE6Zmuiz/hitachi-copy?orgId=1&from=1549413887954&to=1549457721344&panelId=15&width=1000&height=500&tz=Europe%2FKiev`
	// exmaple key
	api_key := "eyJrIjoidjBnNG5OMEw2QnBpTW1KRjMzdE00NzdKUjFuMDZjSTAiLCJuIjoia2V5MSIsImlkIjoxfQ=="
	fname := "image.png"
	if err := GetImage(fname, url, api_key); err != nil {
		log.Fatalf("Error getting image : %v", err)
	}
	fmt.Println("Download saved to", fname)
}

func GetImage(fname, url, api_key string) error {
	os.Remove(fname)
	client := grab.NewClient()
	req, err := grab.NewRequest(fname, url)
	if err != nil {
		return errors.Wrap(err, "Error creating request")
	}
	req.NoResume = true
	req.HTTPRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", api_key))
	resp := client.Do(req)
	if err := resp.Err(); err != nil {
		return errors.Wrap(err, "Error executing request")
	}
	return nil
}
