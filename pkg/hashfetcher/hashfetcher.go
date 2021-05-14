package hashfetcher

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/painhardcore/WebHasher/pkg/md5hasher"
)

const (
	// format for the ouput
	outputFormat = "%s %s\n"
)

func Run(parallel int, urls []string) {
	// set the correct parallel limit
	limit := parallel
	// no reason to make more than maximum
	if parallel > len(urls) {
		limit = len(urls)
	}

	var wg sync.WaitGroup
	barrier := make(chan struct{}, limit)
	hasher := md5hasher.New(http.DefaultClient)

	for i := range urls {
		// block if we hit the limit
		barrier <- struct{}{}
		wg.Add(1)

		go func(ii int) {
			defer wg.Done()

			fullURL := "http://" + urls[ii]
			md5, err := hasher.Hash(fullURL)
			if err != nil {
				fmt.Printf(outputFormat, fullURL, err)
			} else {
				fmt.Printf(outputFormat, fullURL, md5)
			}

			// release
			<-barrier
		}(i)

	}
	// Wait all requests to finish.
	wg.Wait()
}
