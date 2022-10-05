package key

import (
	"crypto/rand"
	"math/big"
	"sync"
	"time"
)

func KeySpace(bitsCount int64) *big.Int {
	bits := big.NewInt(bitsCount)
	two := big.NewInt(2)

	return two.Exp(two, bits, nil)
}

func GenerateKey(bitsCount int64) (*big.Int, error) {
	max := KeySpace(bitsCount)
	max = max.Sub(max, big.NewInt(1))

	n, err := rand.Int(rand.Reader, max)

	if err != nil {
		return nil, err
	}

	return n, nil
}

func searchInInterval(target *big.Int, lowerBound, upperBound big.Int) bool {
	one := big.NewInt(1)
	for i := lowerBound; i.Cmp(&upperBound) == -1; i.Add(&i, one) {
		if i.Cmp(target) == 0 {
			return true
		}
	}
	return false
}

func bruteforceSingle(target *big.Int, bitsCount int64) bool {
	maxNum := KeySpace(bitsCount)
	lowerBound := big.NewInt(0)
	return searchInInterval(target, *lowerBound, *maxNum)
}

func bruteforce(target *big.Int, bitsCount, threadCount int64) bool {
	var wg sync.WaitGroup
	maxNum := KeySpace(bitsCount)
	chunkSize := big.NewInt(0)
	chunkSize = chunkSize.Div(maxNum, big.NewInt(threadCount))
	lowerBound := big.NewInt(0)
	upperBound := big.NewInt(0)
	upperBound = upperBound.Add(lowerBound, chunkSize)

	success := make(chan bool, 1)

	var i int64
	for i = 0; i < threadCount; i++ {
		lowerBoundIt := big.NewInt(0)
		upperBoundIt := big.NewInt(0)
		if i == threadCount-1 {
			upperBound = maxNum
		}
		wg.Add(1)
		go func(target *big.Int, lowerBound, upperBound big.Int, group *sync.WaitGroup) {
			defer wg.Done()
			if searchInInterval(target, lowerBound, upperBound) {
				success <- true
			}
		}(target, *lowerBoundIt.Add(lowerBound, lowerBoundIt), *upperBoundIt.Add(upperBound, upperBoundIt), &wg)
		lowerBound = lowerBound.Add(lowerBound, chunkSize)
		upperBound = upperBound.Add(upperBound, chunkSize)
	}

	wg.Wait()
	select {
	case <-success:
		return true
	default:
		return false
	}
}

func TimeCounter(target *big.Int, bitsCount, threadCount int64) (bool, time.Duration) {
	start := time.Now()
	result := bruteforce(target, bitsCount, threadCount)
	//result := bruteforceSingle(target, bitsCount)
	end := time.Now()
	return result, end.Sub(start)
}
