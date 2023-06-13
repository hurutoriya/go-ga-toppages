package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/adrg/frontmatter"
	ga "google.golang.org/api/analyticsdata/v1beta"
)

type PageData struct {
	ViewCount int64
	PagePath  string
	PageTitle string
}

func main() {
	propertyID := flag.String("property_id", "", "Google Analytics(GA4) property ID")
	siteContentPath := flag.String("site_content_path", "", "Path to site content in static site generator (e.g. Hugo)")
	pagesRootURL := flag.String("pages_root_url", "", "Pages root URL (e.g. https://shunyaueta.com)")
	formatOption := flag.String("format_option", "markdown", "markdown or html")
	topN := flag.Int("top_n", 10, "Number of top pages to retrieve")
	flag.Parse()

	if *propertyID == "" || *siteContentPath == "" || *pagesRootURL == "" || *topN == 0 {
		log.Fatal("Missing required arguments")
	}

	ctx := context.Background()
	client, err := ga.NewService(ctx)

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	oneYearAgo := time.Now().AddDate(-1, 0, 0).Format("2006-01-02")

	request := &ga.RunReportRequest{
		DateRanges: []*ga.DateRange{
			{
				StartDate: oneYearAgo,
				EndDate:   "today",
			},
		},
		Dimensions: []*ga.Dimension{
			{Name: "pagePath"},
		},
		Metrics: []*ga.Metric{
			{Name: "totalUsers"},
		},
	}

	response, err := client.Properties.RunReport(fmt.Sprintf("properties/%s", *propertyID), request).Do()
	if err != nil {
		log.Fatalf("Failed to run GA report: %v", err)
	}

	if *formatOption == "markdown" {
		fmt.Printf("## 直近一年間の人気記事 Top%d\n\n", *topN)
	} else {
		// HTML option
		fmt.Printf("<h2> 直近一年間の人気記事 Top%d</h2> \n", *topN)
		fmt.Printf("<ol>\n")
	}

	for i, row := range response.Rows {
		if i >= *topN+1 {
			break
		}
		pagePath := row.DimensionValues[0].Value
		if pagePath == "/" {
			continue
		}
		viewCount := row.MetricValues[0].Value
		pageTitle := getPageTitle(*siteContentPath, pagePath)
		pageURL := *pagesRootURL + pagePath
		if *formatOption == "markdown" {
			fmt.Printf("1. `%s` unique user access: [%s](%s)\n", viewCount, pageTitle, pageURL)
		} else {
			// HTML option
			fmt.Printf("<li><code>%s</code> unique user access: <a href=\"%s\">%s</a></li> \n", viewCount, pagePath, pageTitle)
		}
	}
	if *formatOption == "html" {
		fmt.Println("</ol>")
	}
}

func getPageTitle(siteContentPath, pagePath string) string {

	markdownFilePath := siteContentPath + pagePath + "index.md"
	content, err := ioutil.ReadFile(markdownFilePath)
	if err != nil {
		log.Fatalf("Failed to read markdown file: %v", err)
	}
	markdownContent := string(content)

	var matter struct {
		Title string
	}
	// Read markdown file title name from front matter.
	_, err = frontmatter.Parse(strings.NewReader(markdownContent), &matter)
	if err != nil {
		log.Fatalf("Failed to parse markdown file: %v", err)
	}

	return matter.Title
}
