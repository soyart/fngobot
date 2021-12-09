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

	"github.com/artnoi43/fngobot/fetch"
	"github.com/pkg/errors"
)

// Quote struct for Bitkub
type Quote struct {
	Last   float64 `json:"last"`
	Bid    float64 `json:"bid"`
	Ask    float64 `json:"ask"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Change float64 `json:"change"`
}

// Get fetches data from Bitkub JSON API,
// and parses the fetched JSON into Quote struct
func Get(tick string) (*Quote, error) {

	/* Documentation for Bitkub:
	 * https://github.com/bitkub/bitkub-official fetch-docs */

	data, err := fetch.Fetch("https://bitkub.com/api/market/ticker/")
	if err != nil {
		return nil, err
	}

	var found bool
	var q Quote
	for key0, val0 := range data.(map[string]interface{}) {
		switch key0 {
		case "data":
			/* Inner keys and values */
			for key1, val1 := range val0.(map[string]interface{}) {
				/* Filter ticker */
				switch key1 {
				case "THB_" + tick:
					var ok bool
					var err error = errors.New("failed to parse float")
					for k, v := range val1.(map[string]interface{}) {
						switch k {
						case "last":
							q.Last, ok = v.(float64)
							if !ok {
								return nil, errors.Wrap(err, "last")
							}
						case "highestBid":
							q.Bid, ok = v.(float64)
							if !ok {
								return nil, errors.Wrap(err, "bid")
							}
						case "lowestAsk":
							q.Ask, ok = v.(float64)
							if !ok {
								return nil, errors.Wrap(err, "ask")
							}
						case "high24hr":
							q.High, ok = v.(float64)
							if !ok {
								return nil, errors.Wrap(err, "high")
							}
						case "low24hr":
							q.Low, ok = v.(float64)
							if !ok {
								return nil, errors.Wrap(err, "low")
							}
						case "percentageChange":
							q.Change, ok = v.(float64)
							if !ok {
								return nil, errors.Wrap(err, "change")
							}
						}
					}
					found = true
				}
			}
		}
	}
	if !found {
		log.Printf("%s not found in Bitkub JSON", tick)
		return nil, fetch.ErrNotFound
	}
	return &q, nil
}
