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

package satang

import (
	"log"
	"strconv"

	"github.com/artnoi43/fngobot/fetch"
)

// Enum for parse() */
const (
	bid = iota
	ask
)

// Quote for Satang only has Bid and Ask fields
type Quote struct {
	Bid float64
	Ask float64
}

func parse(q *Quote, val interface{}, bidAsk int) error {
	for k, v := range val.(map[string]string) {
		switch k {
		case "price":
			price, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return err
			}
			switch bidAsk {
			case bid:
				q.Bid = price
			case ask:
				q.Ask = price
			}
		}
	}
	return nil
}

// Get fetches data from Satang JSON API,
// and parses the fetched JSON into Quote struct
func Get(tick string) (*Quote, error) {

	/* Documentation for Satang:
	 * https://docs.satangcorp.com/apis/public/orders */

	data, err := fetch.Fetch("https://satangcorp.com/api/orderbook-tickers/")
	if err != nil {
		return nil, err
	}

	var q Quote
	var found bool
	for key, val := range data.(map[string]interface{}) {
		/* Filter ticker */
		switch key {
		case tick + "_THB":
			/* Inner keys and values */
			for k, v := range val.(map[string]interface{}) {
				switch k {
				case "bid":
					err := parse(&q, v, bid)
					if err != nil {
						return nil, err
					}
				case "ask":
					err := parse(&q, v, ask)
					if err != nil {
						return nil, err
					}
				}
			}
			found = true
		}
	}
	if !found {
		log.Printf("%s not found in Satang JSON", tick)
		return nil, fetch.ErrNotFound
	}
	return &q, nil
}
