package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sync"

	"github.com/atotto/clipboard"
	"golang.org/x/net/html"
)

// SafeVisitor to be used to determine if url has been visited
type SafeVisitor struct {
	v   map[string]bool
	mux *sync.Mutex
}

var sv SafeVisitor = SafeVisitor{v: make(map[string]bool), mux: &sync.Mutex{}}

func (s SafeVisitor) checkvisited(url string) bool {
	s.mux.Lock()
	defer s.mux.Unlock()
	_, ok := s.v[url]
	if ok == false {
		s.v[url] = true
	}
	return ok
}

func main() {
	argsWithoutProg := os.Args[1:]
	initialUrl := "https://play.google.com/store/apps/collection/cluster?clp=igM6ChkKEzgyMDQ2OTkzNjYyNDAwMTk3MDQQCBgDEhsKFWNvbS52encuaHNzLm15dmVyaXpvbhABGAMYAQ%3D%3D:S:ANO1ljK_y6A&gsr=Cj2KAzoKGQoTODIwNDY5OTM2NjI0MDAxOTcwNBAIGAMSGwoVY29tLnZ6dy5oc3MubXl2ZXJpem9uEAEYAxgB:S:ANO1ljLQ2zk&gl=US"
	if len(argsWithoutProg) > 0 {
		initialUrl = argsWithoutProg[0]
	}
	if appinfo, err := initialFetch(initialUrl); err == nil {
		if csvString, err := formatCsv(appinfo); err == nil {
			if err := clipboard.WriteAll(string(csvString)); err == nil {
				fmt.Println("\n\nCSV copied to your copy/paste buffer")
			} else {
				fmt.Println("Could not copy csv to your copy/paste buffer, outputting here:")
				fmt.Println(csvString)
			}
		}
	}
	fmt.Println("Hit enter to exit")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}

func formatCsv(appInf []appInfo) (string, error) {
	buffer := new(bytes.Buffer)
	csvW := csv.NewWriter(buffer)
	csvW.Comma = '\t'
	if err := csvW.Write([]string{"App Name", "App Total", "App Rating", "App URL"}); err != nil {
		return "", err
	}
	for i := range appInf {
		record := []string{appInf[i].appName, appInf[i].appTotal, appInf[i].appRating, appInf[i].appUrl}
		if err := csvW.Write(record); err != nil {
			return "", err
		}
	}
	csvW.Flush()
	return buffer.String(), nil
}

func findAttrValue(match string, atts []html.Attribute) string {
	for attr := range atts {
		if atts[attr].Key == match {
			return atts[attr].Val
		}
	}
	return ""
}

func findUrlsForClass(className string, n *html.Node) []string {
	var returnUrls []string
	if n.Type == html.ElementNode && n.Data == "a" && findAttrValue("class", n.Attr) == className {
		returnUrls = []string{findAttrValue("href", n.Attr)}
		// fmt.Printf("link: %s\n", returnUrls[0])
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		returnUrls = append(returnUrls, findUrlsForClass(className, c)...)
	}
	return returnUrls
}

func findNodeForDataAndAttrNameValue(findData string, findAttr string, findValue string, n *html.Node) *html.Node {
	if n.Type == html.ElementNode && n.Data == findData && findAttrValue(findAttr, n.Attr) == findValue {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if returnNode := findNodeForDataAndAttrNameValue(findData, findAttr, findValue, c); returnNode != nil {
			return returnNode
		}
	}
	return nil
}

type appInfo struct {
	appName   string
	appTotal  string
	appRating string
	appUrl    string
}

func fetchAppInfo(newUrlString string) (appInfo, error) {
	if doc, err := fetchDocFromUrl(newUrlString); err == nil {
		if titleNode := findNodeForDataAndAttrNameValue("h1", "class", "AHFaub", doc); titleNode != nil {
			// fmt.Printf("title: %s\n", titleNode.FirstChild.FirstChild.Data)
			if totalNode := findNodeForDataAndAttrNameValue("span", "class", "EymY4b", doc); totalNode != nil {
				// fmt.Printf("total: %s\n", totalNode.FirstChild.NextSibling.FirstChild.Data)
				if ratingNode := findNodeForDataAndAttrNameValue("div", "class", "BHMmbe", doc); ratingNode != nil {
					//fmt.Printf("rating: %s\n", ratingNode.FirstChild.Data)
					fetchedAppInfo := appInfo{appName: titleNode.FirstChild.FirstChild.Data,
						appRating: ratingNode.FirstChild.Data,
						appTotal:  totalNode.FirstChild.NextSibling.FirstChild.Data,
						appUrl:    newUrlString}
					return fetchedAppInfo, nil
				}
			}

		}
	} else {
		return appInfo{}, err
	}
	return appInfo{}, fmt.Errorf("nodes not found for %s", newUrlString)
}

func initialFetch(urlString string) ([]appInfo, error) {
	fetchedAppInfo := []appInfo{}
	url, err := url.Parse(urlString)
	if err != nil {
		return fetchedAppInfo, err
	}
	if doc, err := fetchDocFromUrl(urlString); err == nil {
		appUrls := findUrlsForClass("poRVub", doc)
		// fmt.Println(appUrls)
		for appUrl := range appUrls {
			newUrlString := url.Scheme + "://" + url.Host + appUrls[appUrl]
			if returnedAppInfo, err := fetchAppInfo(newUrlString); err == nil {
				fetchedAppInfo = append(fetchedAppInfo, returnedAppInfo)
			}
		}
		return fetchedAppInfo, nil
	} else {
		// fmt.Println("html parse failure")
		return fetchedAppInfo, err
	}
}

func fetchDocFromUrl(url string) (*html.Node, error) {
	fmt.Printf("fetching app info from: %s\n", url)
	if res, err := http.Get(url); err == nil {
		defer res.Body.Close()
		if doc, err := html.Parse(res.Body); err == nil {
			return doc, nil
		} else {
			fmt.Println("html parse failure")
			return nil, err
		}
	} else {
		fmt.Println("http.Get failure")
		return nil, err
	}
}
