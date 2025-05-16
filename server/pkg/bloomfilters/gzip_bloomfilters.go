package bloomfilters

import (
	"bufio"
	"compress/gzip"
	"os"
	"sync"

	"github.com/bits-and-blooms/bloom/v3"
)

type IBloomFilter interface {
	Load(filename []string) error
	ElmentExistsInWhichFiles(coupon string) []string
	IsLoaded() bool
}

type FileBloomFilter struct {
	fileName string
	filter   bloom.BloomFilter
}

// GzipBloomFilter uses Bloom filter to check data.
// It works only with .gz (GZIP) files.
// It reads each line from the file and saves into the filter.
type GzipBloomFilter struct {
	store    []FileBloomFilter
	mu       sync.Mutex
	isLoaded bool
}

// Load loads all .gz files into GzipBloomFilter.
func (filter *GzipBloomFilter) Load(files []string) error {
	var wg sync.WaitGroup
	// buffer to hold all possible errors
	errCh := make(chan error, len(files))

	for _, filename := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			err := filter.loadSingleFile(file)
			if err != nil {
				errCh <- err
				return
			}
		}(filename)
	}

	// Wait for all goroutines
	wg.Wait()

	// Mark loaded
	filter.mu.Lock()
	filter.isLoaded = true
	filter.mu.Unlock()

	// Close channel since all goroutines are done
	close(errCh)

	// Return first error if any
	for err := range errCh {
		if err != nil {
			return err
		}
	}
	return nil
}

// couponExistsInrReturns list of files coupon existed in
func (filter *GzipBloomFilter) ElmentExistsInWhichFiles(coupon string) []string {
	filter.mu.Lock()
	loaded := filter.isLoaded
	filter.mu.Unlock()

	if !loaded {
		return []string{}
	}
	var foundInFiles []string
	for _, couponStore := range filter.store {
		if couponStore.filter.Test([]byte(coupon)) {
			foundInFiles = append(foundInFiles, couponStore.fileName)
		}
	}
	return foundInFiles
}

func (filter *GzipBloomFilter) IsLoaded() bool {
	filter.mu.Lock()
	loaded := filter.isLoaded
	filter.mu.Unlock()
	return loaded
}

// loadSingleFile loads one .gz file into a Bloom filter and saves it.
// It adds the result to the internal store safely using mutex.
func (filter *GzipBloomFilter) loadSingleFile(file string) error {
	bloomfilter, err := fileToBloomFilter(file)
	if err != nil {
		return err
	}

	filter.mu.Lock()
	filter.store = append(filter.store, FileBloomFilter{
		fileName: file,
		filter:   *bloomfilter,
	})
	filter.mu.Unlock()
	return nil
}

// fileToBloomFilter loads a single file into a Bloom filter, and returns bloom filter.
func fileToBloomFilter(filename string) (*bloom.BloomFilter, error) {
	totalWords, err := countLines(filename)
	if err != nil {
		return nil, err
	}
	bloomFilter := bloom.NewWithEstimates(uint(totalWords+100000), 0.01)
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	gzReader, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}
	defer gzReader.Close()

	scanner := bufio.NewScanner(gzReader)
	for scanner.Scan() {
		bloomFilter.Add([]byte(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return bloomFilter, nil
}

func countLines(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	gzReader, err := gzip.NewReader(file)
	if err != nil {
		return 0, err
	}
	defer gzReader.Close()

	scanner := bufio.NewScanner(gzReader)
	count := 0

	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return count, nil
}

func NewGzipBloomFilter() IBloomFilter {
	return &GzipBloomFilter{store: []FileBloomFilter{}}
}
