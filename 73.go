// +build ignore

package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type Crawler struct {
	urls map[string]bool
	sync.Mutex
	sync.WaitGroup
}

func (c *Crawler) crawlImpl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}
	c.Lock()
	_, visited := c.urls[url]
	if !visited {
		c.urls[url] = true
		c.Unlock()
		body, urls, err := fetcher.Fetch(url)
		if err == nil {
			fmt.Printf("found: %s %q\n", url, body)
			for _, u := range urls {
				c.Add(1)
				go c.crawlImpl(u, depth-1, fetcher)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		c.Unlock()
	}
	c.Done()
	return
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	c := &Crawler{urls: map[string]bool{}}
	c.Add(1)
	c.crawlImpl(url, depth, fetcher)
	c.Wait()
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
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

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
