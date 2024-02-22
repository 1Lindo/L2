/*
Программа реализует утилиту wget с возможностью скачивать сайты целиком.

Для тестов:
go get golang.org/x/net/html - установите библиотеку
go run main.go -url "ссылка на сайт" -output output_directory
*/
package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Определение флагов
	urlFlag := flag.String("url", "", "URL of the website to download")
	outputFlag := flag.String("output", "", "Directory to save the downloaded files")
	flag.Parse()

	// Проверка обязательного флага URL
	if *urlFlag == "" {
		fmt.Println("Usage: ./wget -url <URL> [-output <directory>]")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Парсинг URL
	parsedURL, err := url.Parse(*urlFlag)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		os.Exit(1)
	}

	// Получение имени директории для сохранения файлов
	outputDir := *outputFlag
	if outputDir == "" {
		outputDir = getOutputDirectory(parsedURL)
	}

	// Создание директории, если её нет
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating output directory:", err)
		os.Exit(1)
	}

	// Загрузка веб-сайта
	err = downloadWebsite(parsedURL, outputDir)
	if err != nil {
		fmt.Println("Error downloading website:", err)
		os.Exit(1)
	}

	fmt.Println("Website downloaded successfully.")
}

// getOutputDirectory возвращает имя директории для сохранения файлов на основе хоста URL
func getOutputDirectory(parsedURL *url.URL) string {
	host := strings.TrimPrefix(parsedURL.Hostname(), "www.")
	return filepath.Join(".", host)
}

// downloadWebsite выполняет загрузку веб-сайта и сохранение всех файлов
func downloadWebsite(parsedURL *url.URL, outputDir string) error {
	resp, err := http.Get(parsedURL.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request failed with status code %d", resp.StatusCode)
	}

	return processHTML(resp.Body, parsedURL, outputDir)
}

// processHTML извлекает все ссылки из HTML и загружает соответствующие файлы
func processHTML(reader io.Reader, baseURL *url.URL, outputDir string) error {
	links, err := extractLinks(reader)
	if err != nil {
		return err
	}

	for _, link := range links {
		absoluteURL, err := baseURL.Parse(link)
		if err != nil {
			fmt.Println("Error parsing link:", err)
			continue
		}

		err = downloadFile(absoluteURL, outputDir)
		if err != nil {
			fmt.Println("Error downloading file:", err)
		}
	}

	return nil
}

// extractLinks извлекает все ссылки из HTML
func extractLinks(reader io.Reader) ([]string, error) {
	var links []string

	// Используется golang.org/x/net/html для более надежного парсинга HTML
	doc, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}

	var visitNode func(*html.Node)
	visitNode = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visitNode(c)
		}
	}

	visitNode(doc)
	return links, nil
}

// downloadFile загружает указанный файл и сохраняет его на диске
func downloadFile(fileURL *url.URL, outputDir string) error {
	filePath := filepath.Join(outputDir, fileURL.Path)

	resp, err := http.Get(fileURL.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request failed with status code %d", resp.StatusCode)
	}

	err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
