// Package scraperutil provides shared concurrent crawling and HTML directory parsing helpers for providers.
package scraperutil

import (
	"context"
	"fmt"
	"sync"
)

// CrawlConfig configures a concurrent breadth-first URL crawl.
type CrawlConfig struct {
	PanicErr         error
	StartURL         string
	ProviderName     string
	WorkChanCapacity int
	ConcurrencyLimit int
}

type crawlState struct {
	workers      *sync.WaitGroup
	onError      func(error)
	enqueue      func(string)
	tokens       chan struct{}
	panicErr     error
	providerName string
}

// PageFunc processes a crawled URL and returns discovered child URLs to enqueue.
type PageFunc func(ctx context.Context, pageURL string) ([]string, error)

// Crawl performs a concurrent crawl starting from cfg.StartURL.
// processPage is invoked for each unique URL and returns discovered child URLs to enqueue.
func Crawl(ctx context.Context, cfg CrawlConfig, processPage PageFunc) error {
	crawlCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	var (
		visitedMu     sync.Mutex
		visited       = make(map[string]bool)
		workChan      = make(chan string, cfg.WorkChanCapacity)
		activeWorkers sync.WaitGroup
		errOnce       sync.Once
		firstErr      error
	)

	recordErr := func(err error) {
		errOnce.Do(func() {
			firstErr = err

			cancel()
		})
	}

	enqueue := func(targetURL string) {
		if err := crawlCtx.Err(); err != nil {
			recordErr(fmt.Errorf("crawl context canceled: %w", err))

			return
		}

		visitedMu.Lock()
		if visited[targetURL] {
			visitedMu.Unlock()

			return
		}

		visited[targetURL] = true
		visitedMu.Unlock()

		activeWorkers.Add(1)

		go func() {
			select {
			case workChan <- targetURL:
			case <-crawlCtx.Done():
				activeWorkers.Done()
			}
		}()
	}

	state := &crawlState{
		workers:      &activeWorkers,
		onError:      recordErr,
		enqueue:      enqueue,
		tokens:       make(chan struct{}, cfg.ConcurrencyLimit),
		providerName: cfg.ProviderName,
		panicErr:     cfg.PanicErr,
	}

	enqueue(cfg.StartURL)

	var dispatcherWG sync.WaitGroup

	dispatcherWG.Go(func() {
		defer func() {
			if recovered := recover(); recovered != nil {
				state.onError(fmt.Errorf("%w: panic in %s dispatcher: %v", state.panicErr, state.providerName, recovered))
			}
		}()

		runDispatcher(crawlCtx, workChan, state, processPage)
	})

	activeWorkers.Wait()
	close(workChan)
	dispatcherWG.Wait()

	if err := ctx.Err(); err != nil && firstErr == nil {
		firstErr = fmt.Errorf("crawl context canceled: %w", err)
	}

	return firstErr
}

func runDispatcher(ctx context.Context, work <-chan string, state *crawlState, processPage PageFunc) {
	for target := range work {
		if ctx.Err() != nil {
			state.workers.Done()

			continue
		}

		select {
		case <-ctx.Done():
			state.onError(fmt.Errorf("crawl context canceled: %w", ctx.Err()))
			state.workers.Done()

			continue

		case state.tokens <- struct{}{}:
		}

		go runWorker(ctx, target, state, processPage)
	}
}

func runWorker(ctx context.Context, pageURL string, state *crawlState, processPage PageFunc) {
	defer func() {
		if recovered := recover(); recovered != nil {
			state.onError(fmt.Errorf("%w: panic in %s worker for %s: %v", state.panicErr, state.providerName, pageURL, recovered))
		}

		<-state.tokens
		state.workers.Done()
	}()

	subPages, err := processPage(ctx, pageURL)
	if err != nil {
		state.onError(err)

		return
	}

	for _, page := range subPages {
		state.enqueue(page)
	}
}
