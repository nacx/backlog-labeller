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

package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/google/go-github/github"
)

// ReadProjectCardEvent reads a ProjectCardEvent from the given path.
// Events are stored in JSON format by GitHub actions in a given path so they can be read
// from the action code.
func ReadProjectCardEvent(path string) (github.ProjectCardEvent, error) {
	log.Printf("reading event from %q\n", path)

	var event github.ProjectCardEvent
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return event, fmt.Errorf("error reading event file %q: %w", path, err)
	}

	if err = json.Unmarshal(bytes, &event); err != nil {
		return event, fmt.Errorf("error unmarshalling event: %w", err)
	}

	return event, nil
}
