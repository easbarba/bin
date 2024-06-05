/// 2>/dev/null ; gorun "$0" "$@" ; exit $?

package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

const opmURL = "https://gist.githubusercontent.com/funkyhippo/1d40bd5dae11e03a6af20e5a9a030d81/raw/?"

func main() {
	response := parseResponse()
	printInitialInfo(response)
	getPics(response.Chapters)
}

func printInitialInfo(response Opm) {
	fmt.Println("cover:", getDecodedLink(response.Cover))
	fmt.Println("title:", response.Title)
}

func destDir() string {
	if len(os.Args) == 2 {
		return os.Args[1]
	}

	home := home()
	result := filepath.Join(home, "Downloads", "opm")

	return result
}

func getPics(chapters map[string]Chapter) {
	for k, v := range chapters {
		chapterDir := strings.Replace(
			filepath.Join(destDir(), k),
			".", "_", -1)

		fmt.Println("\nChapter:", k)
		fmt.Printf("%-2sDirectory: %s\n", "", chapterDir)

		for _, link := range v.Groups["/r/OnePunchMan"] {
			finalUrl := getDecodedLink(link)
			picNameParsed, err := url.Parse(finalUrl)
			if err != nil {
				fmt.Println("Error parsing image url", err)
				os.Exit(1)
			}
			picName := picNameParsed.Path[1:]
			destinationPath := filepath.Join(chapterDir, picName)

			if _, err := os.Stat(destinationPath); err == nil {
				fmt.Printf("%-4sfound: %s\n", "", destinationPath)
				continue
			}

			time.Sleep(5 * time.Second)

			err = os.MkdirAll(filepath.Dir(destinationPath), 0755)
			if err != nil {
				fmt.Println("Error: unable to create parents dirs of:", destinationPath, err)
				os.Exit(1)
			}

			fmt.Printf("%-4surl: %s - filepath: %s\n", "", finalUrl, destinationPath)
			err = downloadImage(finalUrl, destinationPath)
			if err != nil {
				fmt.Println("Error: unable to download file:", destinationPath, err)
			}
		}
	}
}

func downloadImage(url, destinationPath string) error {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}

	// Create a new HTTP request with the specified URL
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// Set the User-Agent header
	userAgent := "Mozilla/5.0"
	req.Header.Set("User-Agent", userAgent)

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the destination file
	file, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the image data to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func home() string {
	result, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home", err)
		os.Exit(1)
	}

	return result
}

func getDecodedLink(link string) string {
	var result string
	uLink, err := url.Parse(link)
	if err != nil {
		fmt.Println("Error parsing URI", link, err)
		os.Exit(1)
	}

	result = path.Base(uLink.Path)
	decoded, err := base64.StdEncoding.DecodeString(result)
	if err != nil {
		fmt.Println("Error decoding base64", err)
		os.Exit(1)
	}

	result = string(decoded)
	return result
}

func getRequest() []byte {
	filepath := "/tmp/opm.json"
	if _, err := os.Stat(filepath); err == nil {
		content, err := os.ReadFile(filepath)
		if err != nil {
			fmt.Println("Error reading file: ", filepath, err)
			os.Exit(1)
		}

		fmt.Printf("INFO: File found at %s, imma use it!\n\n", filepath)
		return content
	}

	response, err := http.Get(opmURL)
	if err != nil {
		fmt.Println("Error: requesting main json file!", err)
		os.Exit(1)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error: reading response body!")
		os.Exit(1)
	}

	err = os.WriteFile(filepath, body, 0644)
	if err != nil {
		fmt.Println("Error saving response body to file: ", filepath, err)
		os.Exit(1)
	}

	return body
}

func parseResponse() Opm {
	var opm Opm
	err := json.Unmarshal(getRequest(), &opm)
	if err != nil {
		fmt.Println("Error at parsing json response!", err)
		os.Exit(1)
	}

	return opm
}

type Opm struct {
	Chapters    map[string]Chapter `json:"chapters"`
	Cover       string
	Description string
	Title       string
}

type Chapter struct {
	Groups      map[string][]string `json:"groups"`
	LastUpdated int64               `json:"last_updated"`
	Title       string              `json:"title"`
	Volume      string              `json:"volume"`
}
