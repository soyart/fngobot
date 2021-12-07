package fetch

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	// NotFound indicates the ticker symbol is not found in JSON data
	NotFound error = errors.New("Ticker not found in JSON")
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
