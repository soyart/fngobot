/*
 * Copyright 2021 Prem Phansuriyanon
 * Redistribution and use in source and binary forms,
 * with or without modification, are permitted provided
 * that the following conditions are met:
 *
 * 1. Redistributions of source code must retain
 * the above copyright notice, this list of condition
 * and the following disclaimer.
 *
 * 2. Redistributions in binary form must reproduce
 * the above copyright notice, this list of conditions
 * and the following disclaimer in the documentation
 * and/or other materials provided with
 * the distribution.
 *
 * 3. Neither the name of the copyright holder nor the
 * names of its contributors may be used to endorse or
 * promote products derived from this software without
 * specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS
 * AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED
 * WARRANTIES, INCLUDING, BUT NOT LIMITED TO,
 * THE IMPLIED WARRANTIES OF MERCHANTABILITY AND
 * FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
 *
 * IN NO EVENT SHALL THE COPYRIGHT HOLDER OR
 * CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
 * INCIDENTAL, SPECIAL, EXEMPLARY, OR
 * CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED
 * TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 * LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
 * INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
 * LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
 * OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING
 * IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
 * ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

package bitkub

import (
	"log"

	"encoding/json"
	"io/ioutil"
	"net/http"

	fetch "github.com/artnoi43/fngobot/fetch"
)

type Quote struct {
	Last   float64
	Bid    float64
	Ask    float64
	High   float64
	Low    float64
	Change float64
}

func Get(tick string) (*Quote, error) {

	/* Documentation for Bitkub:
	 * https://github.com/bitkub/bitkub-official fetch-docs */

	resp, err := http.Get("https://bitkub.com/api/market/ticker/")
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

	var found bool
	var q Quote

	/* Range over the map[string]interface{}
	 * so that we can destructure data into our struct */

	for key0, val0 := range f.(map[string]interface{}) {
		switch key0 {
		case "data":
			/* Inner keys and values */
			for key1, val1 := range val0.(map[string]interface{}) {
				/* Filter ticker */
				switch key1 {
				case "THB_" + tick:
					for k, v := range val1.(map[string]interface{}) {
						switch k {
						case "last":
							q.Last = v.(float64)
						case "highestBid":
							q.Bid = v.(float64)
						case "lowestAsk":
							q.Ask = v.(float64)
						case "high24hr":
							q.High = v.(float64)
						case "low24hr":
							q.Low = v.(float64)
						case "percentageChange":
							q.Change = v.(float64)
						}
					}
					found = true
				}
			}
		}
	}
	if found {
		return &q, nil
	} else {
		log.Printf("%s not found in Bitkub JSON", tick)
		return nil, fetch.NotFound
	}
}
