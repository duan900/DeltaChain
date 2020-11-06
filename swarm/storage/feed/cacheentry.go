// Copyright 2018 The go-deltachaineum Authors
// This file is part of the go-deltachaineum library.
//
// The go-deltachaineum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-deltachaineum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-deltachaineum library. If not, see <http://www.gnu.org/licenses/>.

package feed

import (
	"bytes"
	"context"
	"time"

	"github.com/deltachaineum/go-deltachaineum/swarm/storage"
)

const (
	hasherCount            = 8
	feedsHashAlgorithm     = storage.SHA3Hash
	defaultRetrieveTimeout = 100 * time.Millisecond
)

// cacheEntry caches the last known update of a specific Swarm feed.
type cacheEntry struct {
	Update
	*bytes.Reader
	lastKey storage.Address
}

// implements storage.LazySectionReader
func (r *cacheEntry) Size(ctx context.Context, _ chan bool) (int64, error) {
	return int64(len(r.Update.data)), nil
}

//returns the feed's topic
func (r *cacheEntry) Topic() Topic {
	return r.Feed.Topic
}