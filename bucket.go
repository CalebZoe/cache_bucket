
import (
	"fmt"
	"sync"
	"time"
)

// myFunc is a type that represents the function we want to decorate.
type myFunc func(int) int

// CacheEntry represents a cache entry
type CacheEntry struct {
	value     int
	timestamp time.Time
}

func cacheDecorator(f myFunc, cacheDuration time.Duration, maxSize int) (myFunc, func() (int, int, int)) {
	cache := make(map[int]CacheEntry)
	var mu sync.Mutex
	var hits, misses int

	getStats := func() (int, int, int) {
		mu.Lock()
		defer mu.Unlock()
		return hits, misses, len(cache)
	}

	// Schedule periodic cache cleanup
	go func() {
		ticker := time.NewTicker(cacheDuration)
		for range ticker.C {
			mu.Lock()
			for key, entry := range cache {
				if time.Since(entry.timestamp) >= cacheDuration {
					delete(cache, key)
				}
			}
			mu.Unlock()
		}
	}()

	return func(n int) int {
		mu.Lock()
		defer mu.Unlock()

		if entry, found := cache[n]; found {
			if time.Since(entry.timestamp) < cacheDuration {
				hits++
				return entry.value
			}
			delete(cache, n)
		}
		misses++
		result := f(n)

		// Only add if not exceeding max size
		if len(cache) >= maxSize {
			var oldest int
			oldestTime := time.Now()
			for key, entry := range cache {
				if entry.timestamp.Before(oldestTime) {
					oldest = key
					oldestTime = entry.timestamp
				}
			}
			delete(cache, oldest)
		}
		
		cache[n] = CacheEntry{result, time.Now()}
		return result
	}, getStats
}

// someExpensiveComputation is an example function we want to cache.
func someExpensiveComputation(n int) int {
	fmt.Printf("Computing result for %d\n", n)
	time.Sleep(2 * time.Second) // Simulating expensive computation
	return n * n
}

func main() {
	// Decorate our expensive computation function with caching.
	cacheDuration := 10 * time.Second // Cache expiry duration
	maxCacheSize := 5                 // Max cache size entries
	cachedComputation := cacheDecorator(someExpensiveComputation, cacheDuration, maxCacheSize)

	// Call the decorated function multiple times with same input.
	fmt.Println(cachedComputation(5))
	fmt.Println(cachedComputation(5))
	fmt.Println(cachedComputation(10))
	fmt.Println(cachedComputation(10))
	fmt.Println(cachedComputation(7))
	fmt.Println(cachedComputation(7))
	fmt.Println(cachedComputation(7))
	fmt.Println(cachedComputation(20))
	fmt.Println(cachedComputation(15))
	fmt.Println(cachedComputation(20))
	fmt.Println(cachedComputation(25))
}
