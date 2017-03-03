/*
Copyright 2017 Aylien, Inc. All Rights Reserved.

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

package newsapi

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type APIClient struct {
}

func (c *APIClient) SelectHeaderContentType(contentTypes []string) string {

	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

func (c *APIClient) SelectHeaderAccept(accepts []string) string {

	if len(accepts) == 0 {
		return ""
	}
	if contains(accepts, "application/json") {
		return "application/json"
	}
	return strings.Join(accepts, ",")
}

func contains(source []string, containvalue string) bool {
	for _, a := range source {
		if strings.ToLower(a) == strings.ToLower(containvalue) {
			return true
		}
	}
	return false
}

func (c *APIClient) CallAPI(path string, method string,
	headerParams map[string]string,
	queryParams []QueryParams,
	formParams []FormParams) (*http.Response, error) {

	client := &http.Client{}
	request := prepareRequest(path, strings.ToUpper(method), headerParams, queryParams, formParams)

	response, err := client.Do(request)

	return response, err
}

func (c *APIClient) ParameterToString(obj interface{}) string {
	objType := reflect.TypeOf(obj).String()

	switch objType {
	case "int32":
		return strconv.FormatInt(int64(obj.(int32)), 10)
	case "int":
		return strconv.FormatInt(int64(obj.(int)), 10)
	case "int64":
		return strconv.FormatInt(int64(obj.(int64)), 10)
	case "bool":
		return strconv.FormatBool(obj.(bool))
	case "time.Time":
		return obj.(time.Time).UTC().Format("2006-01-02T15:04:05Z")
	default:
		return obj.(string)
	}
}

func prepareRequest(path string,
	method string,
	headerParams map[string]string,
	queryParams []QueryParams,
	formParams []FormParams) *http.Request {

	var body io.Reader
	var request *http.Request

	// add form parameter, if any
	if len(formParams) > 0 {
		data := url.Values{}

		for _, element := range formParams {
			data.Set(element.Item1, element.Item2)
		}

		body = bytes.NewBufferString(data.Encode())
	}

	if method == "GET" {
		request, _ = http.NewRequest(method, path, nil)
	} else {
		request, _ = http.NewRequest(method, path, body)
	}

	//set Client User Agent header
	configuration := NewConfiguration()
	request.Header.Add("User-Agent", configuration.UserAgent)

	// add header parameter, if any
	if len(headerParams) > 0 {
		for k, v := range headerParams {
			request.Header.Add(k, v)
		}
	}

	// add query parameter, if any
	if len(queryParams) > 0 {
		q := request.URL.Query()

		for _, element := range queryParams {
			q.Add(element.Item1, element.Item2)
		}

		request.URL.RawQuery = q.Encode()
	}

	return request
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Func, reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.Array:
		z := true
		for i := 0; i < v.Len(); i++ {
			z = z && isZero(v.Index(i))
		}
		return z
	case reflect.Struct:
		z := true
		for i := 0; i < v.NumField(); i++ {
			z = z && isZero(v.Field(i))
		}
		return z
	}
	// Compare other types directly:
	z := reflect.Zero(v.Type())
	return v.Interface() == z.Interface()
}
