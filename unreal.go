package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	// UnrealSiteURL ...
	UnrealSiteURL = "https://issues.unrealengine.com/issue/search?"
)

// CrawlingData ...
type CrawlingData struct {
	URL   string
	Title string
	Desc  string
}

// GetPageData ...
func GetPageData(v string) ([]CrawlingData, error) {
	ret := []CrawlingData{}

	query := fmt.Sprintf("q=%s&resolution=fixed&component=&sort=updated", url.QueryEscape(v))

	doc, err := goquery.NewDocument(UnrealSiteURL + query)
	if err != nil {
		return ret, err
	}

	doc.Find(".result").Each(func(i int, s *goquery.Selection) {

		cd := CrawlingData{}

		h4 := s.Find("h4 a").First()
		url, exist := h4.Attr("href")
		if exist {
			cd.URL = "https://issues.unrealengine.com" + url
			cd.Title = h4.Text()
		}

		p := s.Find("p").First()
		desc := p.Text()
		if len(desc) > 0 {
			cd.Desc = strings.Replace(desc, `"`, `*`, -1)
			cd.Desc = strings.Replace(cd.Desc, `\t`, `    `, -1)
		}

		if len(cd.URL) > 0 && len(cd.Title) > 0 {
			ret = append(ret, cd)
		}
	})

	return ret, nil
}

func loadUnrealVersion(path string) string {
	loadedVersion := "fix:4.24.2"

	data, err := ioutil.ReadFile(path)
	if err != nil {
		ioutil.WriteFile(path, []byte(loadedVersion), os.ModePerm)
	}
	if data == nil {
		log.Println("load default version", loadedVersion)
	} else {
		log.Println("version loaded ", string(data))
		loadedVersion = string(data)
	}
	return loadedVersion
}
