package main

import (
	"net/http"
)

type GitHubToken struct {
	tr    http.RoundTripper
	token string
}

func (gh *GitHubToken) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+gh.token)
	return gh.tr.RoundTrip(req)
}
func NewGitHubToken(tr http.RoundTripper, token string) *GitHubToken {
	return &GitHubToken{tr: tr, token: token}
}
