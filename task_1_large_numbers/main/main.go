package main

import (
	"fmt"
	"large-numbers/key"
)

func main() {
	var (
		bits              = []int64{8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096}
		threadCount int64 = 8
	)
	for _, bitNum := range bits {
		keySpace := key.KeySpace(bitNum)
		fmt.Printf("bits: %v\nkeyspace: %v\n", bitNum, keySpace.String())
		target, _ := key.GenerateKey(bitNum)
		fmt.Printf("key generated: %v\n", target.String())
		result, time := key.TimeCounter(target, bitNum, threadCount)
		fmt.Printf("bruteforce successful: %v\ntime elapsed to bruteforce the key: %vms\n\n", result, time.Milliseconds())
	}
}
