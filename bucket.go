
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

// cacheDecorator is a decorator that caches results of the given function with additional functionalities.
func cacheDecorator(f myFunc, cacheDuration time.Duration, maxSize int) myFunc {
	cache := make(map[int]CacheEntry)
	var mu sync.Mutex
	var hits, misses int

	return func(n int) int {
		mu.Lock()
		defer mu.Unlock()

		// Concurrent Execution Handling: Ensure only one execution per unique input
		resultCh := make(chan int)
		go func() {
			if entry, found := cache[n]; found {
				// Check if the entry has expired
				if time.Since(entry.timestamp) < cacheDuration {
					hits++
					resultCh <- entry.value
					return
				}
				// Entry expired
				delete(cache, n)
			}
			misses++
			// Perform computation if not found or expired
			result := f(n)
			// Overwrite cache entry
			cache[n] = CacheEntry{result, time.Now()}
			// Maintain cache size
			if len(cache) > maxSize {
				for key := range cache {
					delete(cache, key)
					break
				}
			}
			resultCh <- result
		}()
		return <-resultCh
	}
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
