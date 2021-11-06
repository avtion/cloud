package main

import (
	"net"
	"net/http"
	"time"

	"github.com/andygrunwald/go-trending"
)

func newGithubTrending() *githubTrending {
	trend := trending.NewTrendingWithClient(&http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	})
	return &githubTrending{trend: trend}
}

// githubTrending github 热门项目
type githubTrending struct {
	trend *trending.Trending
}

// getProjects 获取热门项目
func (g *githubTrending) getProjects() ([]trending.Project, error) {
	return g.trend.GetProjects(trending.TimeToday, "go")
}
