// Fetch prints the content found at a URL.
package main

import (
    "fmt"
    // "io/ioutil"
    "net/http"
    "os"
	"io"
	"strings"
)

func main() {
    for _, url := range os.Args[1:] {
        resp, err := http.Get(check_url(url))
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }
        // b, err := io.ReadAll(resp.Body)
		io.Copy(os.Stdout, resp.Body) // practice 1.7
		
		status := resp.Status // practice 1.9
        resp.Body.Close()
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
            os.Exit(1)
        }
        // fmt.Printf("%s", b)
		fmt.Printf("HTTP Status Code: %s\n", status) // practice 1.9
    }
}

// practice 1.8
func check_url(url string) string { 
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	return url
}