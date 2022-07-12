package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

var (
	// ErrNotFound indicates the ticker symbol is not found in JSON data
	ErrNotFound   error = errors.New("Ticker not found in JSON")
	ErrBadRequest error = errors.New("Bad HTTP request")
)

// Fetch is a generic function used to fetch a HTTP response.
// It simply returns the response body
func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrapf(err, "HTTP GET %s failed", url)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(err, "non-200 status: %d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if resp.StatusCode < 500 {
			return nil, ErrBadRequest
		}
		return nil, errors.Wrap(err, "failed to read response body")
	}
	return body, nil
}

// FetchMapStrInf calls Fetch and return the data as map[string]interface{}
func FetchMapStrInf(url string) (map[string]interface{}, error) {
	body, err := Fetch(url)
	if err != nil {
		return nil, err
	}
	/* Since the JSON key is arbitary,
	 * we first unmarshal it into empty interface f */
	var v map[string]interface{}
	err = json.Unmarshal(body, &v)
	if err != nil {
		log.Println("Error parsing JSON: ", err)
		return nil, errors.Wrap(err, "error parsing JSON response body")
	}
	return v, nil
}
