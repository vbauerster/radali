package tracker

import (
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

const (
	maxRedirects = 10
)

var ErrMaxRedirect = fmt.Errorf("tracker: max redirects (%d) followed", maxRedirects)

type ExtractTarget func(tracker *url.URL) (*url.URL, error)

var trackers = make(map[string]ExtractTarget)

type CleanUpRule struct {
	Params       []string
	InvertParams bool
	EmptyParams  bool
	EmptyPath    bool
	Scheme       string
}

var knownShops = map[string]CleanUpRule{
	"tmall.aliexpress.com": CleanUpRule{
		Params:       []string{"SearchText"},
		InvertParams: true,
	},
	"ru.aliexpress.com": CleanUpRule{
		Params:       []string{"SearchText"},
		InvertParams: true,
	},
	"www.gearbest.com": CleanUpRule{
		EmptyParams: true,
	},
	"www.coolicool.com": CleanUpRule{
		EmptyParams: true,
	},
	"www.tinydeal.com": CleanUpRule{
		EmptyParams: true,
	},
	"www.banggood.com": CleanUpRule{
		EmptyParams: true,
	},
}

// RegisterTracker ...
func RegisterTracker(host string, fn ExtractTarget) ExtractTarget {
	prevFn := trackers[host]
	trackers[host] = fn
	return prevFn
}

func follow(rawurl string) (*url.URL, error) {
	trackURL, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	// number of redirects followed
	var redirectsFollowed int
	client := &http.Client{
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	for {
		if f, ok := trackers[trackURL.Host]; ok {
			return f(trackURL)
		}
		resp, err := client.Get(trackURL.String())
		if err != nil {
			return nil, err
		}
		resp.Body.Close()

		if isRedirect(resp.StatusCode) {
			loc, err := resp.Location()
			if err != nil {
				return nil, err
			}

			// fmt.Println("get:", trackURL)
			// fmt.Println("loc:", loc)
			// fmt.Println()

			redirectsFollowed++
			if redirectsFollowed > maxRedirects {
				return nil, ErrMaxRedirect
			}
			trackURL = loc
		} else {
			return trackURL, nil
		}
	}
}

func Untrack(rawurl string) (string, error) {
	if !strings.Contains(rawurl, "://") && !strings.HasPrefix(rawurl, "//") {
		rawurl = "//" + rawurl
	}
	targetURL, err := follow(rawurl)
	if err != nil {
		return "", err
	}
	if rule, ok := knownShops[targetURL.Host]; ok {
		// fmt.Printf("%s = %+values\n", targetURL.Host, rule)
		if rule.EmptyParams {
			targetURL.RawQuery = ""
		} else if len(rule.Params) != 0 {
			values := targetURL.Query()
			toDelKeys := make(map[string]bool, len(values))
			for k := range values {
				toDelKeys[k] = rule.InvertParams
			}
			for _, k := range rule.Params {
				toDelKeys[k] = !rule.InvertParams
			}
			for k, toDel := range toDelKeys {
				if toDel {
					values.Del(k)
				}
			}
			targetURL.RawQuery = values.Encode()
		}

		if rule.EmptyPath {
			targetURL.Path = ""
		}
		if rule.Scheme != "" {
			targetURL.Scheme = rule.Scheme
		}
	}
	return targetURL.String(), nil
}

// KnownTrackers ...
func KnownTrackers() []string {
	list := make([]string, 0, len(trackers))
	for k := range trackers {
		list = append(list, k)
	}
	sort.Strings(list)
	return list
}

func isRedirect(status int) bool {
	return status > 299 && status < 400
}
