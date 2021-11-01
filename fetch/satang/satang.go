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

	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/artnoi43/fngobot/fetch"
)

/* Enum for parse() */
const (
	bid = iota
	ask
)

type Quote struct {
	Bid float64
	Ask float64
}

func parse(q *Quote, val interface{}, bidAsk int) error {
	for k, v := range val.(map[string]interface{}) {
		switch k {
		case "price":
			price, err := strconv.ParseFloat(v.(string), 64)
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

func Get(tick string) (*Quote, error) {

	/* Documentation for Satang:
	 * https://docs.satangcorp.com/apis/public/orders */

	resp, err := http.Get("https://satangcorp.com/api/orderbook-tickers/")
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

	var q Quote
	var found bool

	/* Range over the map[string]interface{}
	 * so that we can destructure our data into our struct */

	for key, val := range f.(map[string]interface{}) {
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
	if found {
		return &q, nil
	} else {
		log.Printf("%s not found in Satang JSON", tick)
		return nil, fetch.NotFound
	}
}
