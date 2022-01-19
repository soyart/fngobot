package fetch

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

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

// Fetch is a generic function used to fetch HTTP response
func Fetch(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
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
	return f, nil
}
