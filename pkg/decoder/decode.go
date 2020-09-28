/*
Copyright 2020, Author: Ashutosh Kumar (GithubID: @sonasingh46).

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package decoder

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Decode the request body to appropriate structure based on content
// type
func DecodeBody(req *http.Request, out interface{}) error {
	cType, err := GetContentType(req)
	if err != nil {
		return err
	}

	if cType != "application/json" {
		return fmt.Errorf("Invalid content type:%s", cType)
	}
	return DecodeJsonBody(req, out)
}

// Get the value of Content-Type that is set in http request header
func GetContentType(req *http.Request) (string, error) {

	if req.Header == nil {
		return "", fmt.Errorf("Request does not have any header")
	}

	return req.Header.Get("Content-Type"), nil
}

// decodeJsonBody is used to decode a JSON request body
func DecodeJsonBody(req *http.Request, out interface{}) error {
	dec := json.NewDecoder(req.Body)
	return dec.Decode(&out)
}
