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

import "net/http"

// Token is the GitHub authentication token
// It also implements the http.RoundTripper interface to automatically add
// the credentials to every request.
type Token string

// RoundTrip implements the http.RoundTripper interface to automatically add the
// token as a header to every HTTP request.
func (t Token) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("authorization", "token "+string(t))
	return http.DefaultTransport.RoundTrip(req)
}
