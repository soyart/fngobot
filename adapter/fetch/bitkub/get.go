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
	"encoding/json"
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
	data, err := common.Fetch(u.String())
	if err != nil {
		return nil, errors.Wrap(err, "common failed()")
	}

	var resp bitkubResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal bitkub response")
	}

	var found bool
	var q common.ApiQuote
	for tokenName, tokenInfo := range resp.DataPart {
		if tokenName == TickerPrefix+tick {
			found = true
			q.Bid = tokenInfo.Bid
			q.Ask = tokenInfo.Ask
			q.Last = tokenInfo.Last
			break
		}
	}

	if !found {
		log.Printf("%s not found in Bitkub JSON\n", tick)
		return nil, common.ErrNotFound
	}
	return &q, nil
}
