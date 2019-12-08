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

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nacx/backlog-labeller/github"
)

// DefaultTimeout is the default timeout used in HitHub API calls
const DefaultTimeout = 30 * time.Second

func main() {
	token := flag.String("token", "", "GitHub API token")
	timeout := flag.Duration("default-timeout", DefaultTimeout, "Default timeout for the GitHub API calls")
	flag.Parse()

	if *token == "" {
		fmt.Println("missing API token")
		flag.Usage()
		os.Exit(1)
	}

	event, err := github.ReadProjectCardEvent(os.Getenv("GITHUB_EVENT_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	gh := github.New(*token, *timeout)

	i, err := gh.GetIssue(context.Background(), event.ProjectCard.GetContentURL())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("issue: #%d (%s)\n", i.GetID(), i.GetTitle())
	fmt.Printf("details: %+v", i)
}
