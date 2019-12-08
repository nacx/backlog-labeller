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
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	token := flag.String("token", "", "GitHub API token")
	flag.Parse()

	if *token == "" {
		fmt.Println("missing API token")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("repo: %s\n", os.Getenv("GITHUB_REPOSITORY"))
	fmt.Printf("user: %s\n", os.Getenv("GITHUB_ACTOR"))

	bytes, err := ioutil.ReadFile(os.Getenv("GITHUB_EVENT_PATH"))
	if err != nil {
		fmt.Printf("error reading file: %v", err)
		os.Exit(1)
	}

	var event github.ProjectCardEvent
	if err = json.Unmarshal(bytes, &event); err != nil {
		fmt.Printf("error unmarshalling event: %v", err)
		os.Exit(1)
	}
	iss := event.ProjectCard.GetContentURL()

	fmt.Printf("issue: %s\n", iss)

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: *token})
	oc := oauth2.NewClient(context.Background(), ts)
	_ = github.NewClient(oc)
}
