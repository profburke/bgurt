// Copyright (c) 2020 BlueDino Software (http://bluedino.net)
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation and/or
//    other materials provided with the distribution.
// 3. Neither the name of the copyright holder nor the names of its contributors may be
//    used to endorse or promote products derived from this software without specific prior
//    written permission.
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY
// EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL
// THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT
// OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
// HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR
// TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// Package bggclient encapsulates the setup of a TLS connection to https://boardgamegeek.com
// and is reponsible for including the user's login credentials as cookies with each HTTP
// request.
//
package bggclient

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"path/filepath"

	"golang.org/x/net/publicsuffix"
)

var tlsConfig *tls.Config // TODO: does this (and transport) need to be global?
var client *http.Client
var transport *http.Transport
var bggURL *url.URL

type Credentials struct {
	Username string `toml:"username"`
	PassHash string `toml:"passhash"`
}

func (c Credentials) IsSet() bool {
	return len(c.Username) > 0 && len(c.PassHash) > 0
}

// TODO: func (c Credentials) NewFromFile(filename string) ....

// TODO: need to decide how to handle URLs. We could almost completely get away
// with just using strings since the methods on http.Client require strings, not URLs,
// with one exception: the cookiejar.SetCookies call :(

func init() {
	tlsConfig = &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	transport = &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		log.Fatalf("bggclient: Error creating cookie jar: %v", err)
	}

	client = &http.Client{
		Transport: transport,
		Jar:       jar,
	}

	bggURL, err = url.Parse("https://boardgamegeek.com")
	if err != nil {
		log.Fatalf("bggclient: Error parsing URL: %v", err)
	}
}

// Download handls an HTTP response that includes a file download.
//
func Download(relativeURL *url.URL) (data []byte, err error) {
	u := bggURL.ResolveReference(relativeURL)
	res, err := client.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)

	_, err = io.Copy(writer, res.Body)
	if err == nil {
		data = b.Bytes()
	}

	return
}

// Get sends an HTTP GET request to the given URL.
//
func Get(relativeURL *url.URL) (page string, err error) {
	u := bggURL.ResolveReference(relativeURL)
	res, err := client.Get(u.String())
	defer res.Body.Close()

	if err != nil {
		return "", err
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	page = string(bytes)

	return
}

// Post sends an HTTP POST request to the given URL with the given form data.
//
func Post(relativeURL *url.URL, data url.Values) (res *http.Response, err error) {
	u := bggURL.ResolveReference(relativeURL)
	res, err = client.PostForm(u.String(), data)

	return
}

// SetCredentials creates cookies containing the BGG username and password hash
// for use by the client.
//
func SetCredentials(c Credentials) {
	usernameCookie := http.Cookie{Name: "bggusername", Value: c.Username}
	passHashCookie := http.Cookie{Name: "bggpassword", Value: c.PassHash}
	client.Jar.SetCookies(bggURL, []*http.Cookie{&usernameCookie, &passHashCookie})
}

// Upload sends an HTTP POST request with a file upload.
//
func Upload(relativeURL *url.URL, data []byte, filename, fileFieldName string, fields map[string]string) (err error) {
	u := bggURL.ResolveReference(relativeURL)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(fileFieldName, filepath.Base(filename))
	if err != nil {
		return err
	}

	r := bytes.NewReader(data)
	_, err = io.Copy(part, r)

	for key, val := range fields {
		_ = writer.WriteField(key, val)
	}

	err = writer.Close()
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := client.Do(request)
	if err != nil {
		return err
	}

	res.Body.Close()

	return nil
}

// Local Variables:
// compile-command: "go build"
// End:
