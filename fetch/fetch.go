package fetch

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

type FetchFunc func(string) (interface{}, error)

// Quoter is returned by all Get functions
type Quoter interface {
	Last() (float64, error)
	Bid() (float64, error)
	Ask() (float64, error)
}

var (
	// ErrNotFound indicates the ticker symbol is not found in JSON data
	ErrNotFound error = errors.New("Ticker not found in JSON")
)

// Fetch is a generic function used to fetch a HTTP response.
// It simply returns the response body
func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.Wrap(err, "non-200 status")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
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
	var f interface{}
	err = json.Unmarshal(body, &f)
	if err != nil {
		log.Println("Error parsing JSON: ", err)
		return nil, err
	}
	return f.(map[string]interface{}), nil
}
