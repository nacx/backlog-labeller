// Copyright 2019 Ignasi Barrera
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package github provides a wrapper on the GitHub API to make it easier to
// make HTTP calls when the entire URI of the resources is known in advance.
package github

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/go-github/github"
)

// GitHub is a small wrapper on top of the GitHub API to make it easier to run
// API calls when we already have the full URI of the resources to retrieve.
type GitHub struct {
	token   string
	timeout time.Duration
	client  *github.Client
}

// New creates a new GitHub wrapper that authenticates to the API using the given token and performs
// all requests using the configured timeout.
func New(token string, timeout time.Duration) GitHub {
	return GitHub{
		token:   token,
		timeout: timeout,
		client: github.NewClient(&http.Client{
			Timeout:   timeout,
			Transport: Token(token),
		}),
	}
}

// GetIssue gets a GitHub issue given its URL
func (g GitHub) GetIssue(ctx context.Context, url string) (*github.Issue, error) {
	log.Printf("getting issue %q\n", url)

	req, err := g.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating get issue request: %w", err)
	}

	issue := new(github.Issue)
	if _, err = g.client.Do(ctx, req, issue); err != nil {
		return nil, fmt.Errorf("error getting issue %q: %w", url, err)
	}

	return issue, nil
}
