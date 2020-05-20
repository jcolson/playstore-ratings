package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sync"

	"golang.org/x/net/html"
)

// Fetcher interface
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

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

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	defer wg.Done()
	if depth <= 0 || sv.checkvisited(url) {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher, wg)
	}
	return
}

func main() {
	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithoutProg)
	initialUrl := "https://play.google.com/store/apps/collection/cluster?clp=igM6ChkKEzgyMDQ2OTkzNjYyNDAwMTk3MDQQCBgDEhsKFWNvbS52encuaHNzLm15dmVyaXpvbhABGAMYAQ%3D%3D:S:ANO1ljK_y6A&gsr=Cj2KAzoKGQoTODIwNDY5OTM2NjI0MDAxOTcwNBAIGAMSGwoVY29tLnZ6dy5oc3MubXl2ZXJpem9uEAEYAxgB:S:ANO1ljLQ2zk&gl=US"
	if len(argsWithoutProg) > 0 {
		initialUrl = argsWithoutProg[0]
	}
	// var wg sync.WaitGroup
	// wg.Add(1)
	fetcher.Fetch(initialUrl)
	// Crawl(initialUrl, 4, fetcherx, &wg)
	// Wait for all Crawls to complete
	// wg.Wait()
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

// an http fetcher
type httpFetcher struct {
}

func (f httpFetcher) Fetch(urlString string) ([]appInfo, error) {
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
			if doc, err := fetchDocFromUrl(newUrlString); err == nil {
				if titleNode := findNodeForDataAndAttrNameValue("h1", "class", "AHFaub", doc); titleNode != nil {
					// fmt.Printf("title: %s\n", titleNode.FirstChild.FirstChild.Data)
					if totalNode := findNodeForDataAndAttrNameValue("span", "class", "EymY4b", doc); totalNode != nil {
						// fmt.Printf("total: %s\n", totalNode.FirstChild.NextSibling.FirstChild.Data)
						if ratingNode := findNodeForDataAndAttrNameValue("div", "class", "BHMmbe", doc); ratingNode != nil {
							// fmt.Printf("rating: %s\n", ratingNode.FirstChild.Data)
							fetchedAppInfo = append(fetchedAppInfo, appInfo{
								appName:   titleNode.FirstChild.FirstChild.Data,
								appRating: ratingNode.FirstChild.Data,
								appTotal:  totalNode.FirstChild.NextSibling.FirstChild.Data,
								appUrl:    newUrlString})
						}
					}

				}
			} else {
				return fetchedAppInfo, err
			}
		}
		return fetchedAppInfo, nil
	} else {
		// fmt.Println("html parse failure")
		return fetchedAppInfo, err
	}
}

func fetchDocFromUrl(url string) (*html.Node, error) {
	// fmt.Printf("url fetching: %s\n", url)
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

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = httpFetcher{}

// fetcher is a populated fakeFetcher.
var fetcherx = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
