
import (
	"fmt"
	"sync"
)

// myFunc is a type that represents the function we want to decorate.
type myFunc func(int) int

// cacheDecorator is a decorator that caches results of the given function.
func cacheDecorator(f myFunc) myFunc {
	cache := make(map[int]int)
	var mu sync.Mutex

	return func(n int) int {
		mu.Lock()
		defer mu.Unlock()
		if result, found := cache[n]; found {
			return result
		}
		result := f(n)
		cache[n] = result
		return result
	}
}

// someExpensiveComputation is an example function we want to cache.
func someExpensiveComputation(n int) int {
	// Example expensive computation
	fmt.Printf("Computing result for %d\n", n)
	return n * n // For example purposes, let's assume the expensive computation is squaring the number.
}

func main() {
	// Decorate our expensive computation function with caching.
	cachedComputation := cacheDecorator(someExpensiveComputation)

	// Call the decorated function multiple times with same input.
	fmt.Println(cachedComputation(5)) // Should compute and cache
	fmt.Println(cachedComputation(5)) // Should return cached result
	fmt.Println(cachedComputation(10)) // Should compute and cache
	fmt.Println(cachedComputation(10)) // Should return cached result
}
