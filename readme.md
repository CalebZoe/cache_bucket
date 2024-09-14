# Caching Decorator Example

This repository contains a simple example of how to implement a caching decorator in Go to optimize function calls that are expensive to compute multiple times.

## Table of Contents

1. [Overview](#overview)
2. [Files](#files)
3. [How It Works](#how-it-works)
4. [Installation](#installation)
5. [Usage](#usage)
6. [Example Output](#example-output)

## Overview

The Go program in this repository demonstrates how to use a decorator pattern to cache results of a function, specifically an expensive computation. The program includes:

- A decorator function (`cacheDecorator`) that wraps around another function to cache its results.
- An example function (`someExpensiveComputation`) that performs a costly computation, which in this case is squaring an integer.
  
The caching system ensures that repeated calls with the same input return a cached result, improving performance by avoiding redundant computations.

## Files

- `bucket.go`: Contains the main code for the caching decorator 
## How It Works

1. **Function Decoration**:
   - The `cacheDecorator` function takes a function of type `myFunc` (which is `func(int) int`) and returns a new function of the same type.
   - A mutex (`sync.Mutex`) is used to ensure that caching is thread-safe in a concurrent environment.
   - Results of the decorated function are stored in a map (`cache`), where the key is the function input, and the value is the computed result.

2. **Example Function**:
   - The `someExpensiveComputation` function simulates an expensive operation by computing the square of an integer.

3. **Caching Process**:
   - When the decorated function (`cachedComputation`) is called, the cache is checked. If the result for the input exists in the cache, it is returned. Otherwise, the original function is called, and the result is stored in the cache.

## Installation

To run this code, you'll need to have Go installed on your machine. If you don't have Go installed, you can follow the [official Go installation guide](https://golang.org/doc/install).

### Steps:

1. Clone the repository or download the `bucket.go` file.
2. Navigate to the folder containing `bucket.go`.
3. Ensure you have Go installed by running `go version`.

## Usage

1. Open your terminal and navigate to the directory where `bucket.go` is located.
2. Run the program using the Go command:

   ```bash
   go run bucket.go
   ```

The program will execute and print results of the cached computations. The second time an input is used, the cached result will be retrieved instead of recomputing it.

## Example Output

```bash
Computing result for 5
25
25
Computing result for 10
100
100
```

In this example, you can see that when the function is called with `5` and `10` for the first time, the expensive computation is performed. For subsequent calls with the same inputs, the cached results are returned without performing the computation again.

## Notes

- This example uses a `sync.Mutex` to ensure thread safety. This means it can be safely used in concurrent Go programs.
- You can replace the `someExpensiveComputation` function with any other function that is computationally expensive, and the caching decorator will optimize it similarly.

Feel free to extend and modify the code to suit your use case!
