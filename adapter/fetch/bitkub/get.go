package bitkub

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

import (
	"fmt"
	"log"
	"net/url"

	"github.com/pkg/errors"

	"github.com/artnoi43/fngobot/adapter/fetch/common"
	"github.com/artnoi43/fngobot/usecase"
)

// Get fetches data from Bitkub JSON API,
// and parses the fetched JSON into quote struct
func (f *fetcher) Get(tick string) (usecase.Quoter, error) {

	/* Documentation for Bitkub:
	 * https://github.com/bitkub/bitkub-official common-docs */

	// @NOTE: Query string for this endpoint does not seem to work
	u, err := url.Parse(BaseURL)
	if err != nil {
		return nil, errors.Wrap(err, "parse url failed")
	}
	queryString := url.QueryEscape(fmt.Sprintf("sym=%s", tick))
	u.RawQuery = queryString
	data, err := common.FetchMapStrInf(u.String())
	if err != nil {
		return nil, errors.Wrap(err, "common failed()")
	}

	var found bool
	var q quote
	for key0, val0 := range data {
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
							q.last, ok = v.(float64)
							if !ok {
								return nil, errors.Wrap(err, "last")
							}
						case "highestBid":
							q.bid, ok = v.(float64)
							if !ok {
								return nil, errors.Wrap(err, "bid")
							}
						case "lowestAsk":
							q.ask, ok = v.(float64)
							if !ok {
								return nil, errors.Wrap(err, "ask")
							}
						case "high24hr":
							q.high, ok = v.(float64)
							if !ok {
								return nil, errors.Wrap(err, "high")
							}
						case "low24hr":
							q.low, ok = v.(float64)
							if !ok {
								return nil, errors.Wrap(err, "low")
							}
						case "percentageChange":
							q.change, ok = v.(float64)
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
		log.Printf("%s not found in Bitkub JSON\n", tick)
		return nil, common.ErrNotFound
	}
	return &q, nil
}
