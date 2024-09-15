package wget

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/net/html"
)

var outputDir string

var wgetCmd = &cobra.Command{
	Use:   "go run main.go [url]",
	Short: "Download a webpage and its resources",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		if err := downloadPage(url, outputDir); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
	},
}

func init() {
	wgetCmd.Flags().StringVarP(&outputDir, "output", "o", ".", "Directory to save the downloaded files")
}

func sanitizeFileName(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "unknown"
	}
	fileName := path.Base(u.Path)
	fileName = strings.ReplaceAll(fileName, "?", "_")
	fileName = strings.ReplaceAll(fileName, "&", "_")
	fileName = strings.ReplaceAll(fileName, "=", "_")
	if fileName == "" || fileName == "/" {
		fileName = "index.html"
	}
	return fileName
}

func downloadPage(url, dir string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch page: %s", resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	htmlFile := path.Join(dir, "index.html")
	file, err := os.Create(htmlFile)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := html.Render(file, doc); err != nil {
		return err
	}

	var resources []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key == "href" || attr.Key == "src" {
					resources = append(resources, attr.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	for _, res := range resources {
		if strings.HasPrefix(res, "http") {
			if err := downloadResource(res, dir); err != nil {
				fmt.Fprintf(os.Stderr, "Error downloading resource %s: %v\n", res, err)
			}
		}
	}

	return nil
}

func downloadResource(url, dir string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch resource: %s", resp.Status)
	}

	fileName := sanitizeFileName(url)
	filePath := path.Join(dir, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

func Execute() error {
	return wgetCmd.Execute()
}
