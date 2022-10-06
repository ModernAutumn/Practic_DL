package main

import (
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

var one = big.NewInt(1)

func pow(num, p int64) *big.Int {
	
  var result big.Int
	
  return result.Exp(big.NewInt(num), big.NewInt(p), nil)
}

func generateRandomKey(n int64) *big.Int {
	random := big.NewInt(0)
	for i := int64(0); i < n; i++ {
		
    randBit := rand.Intn(2)
		
    if randBit > 0 {
			
      random.Add(random, pow(int64(2), i))
		
    }

	}
	return random
}

func bruteForceKey(target *big.Int, n int64) (time.Duration, error) {
	
  numOfKeys := pow(2, n)
	
startTime := time.Now()

  for i := big.NewInt(0); i.Cmp(numOfKeys) < 0; i.Add(i, one) {
		
    if i.Cmp(target) == 0 {
      return time.Since(startTime), nil
    }
	
  }

	return 0, errors.New("The match not found")
}

func main() {
	
  const size = 10
	rand.Seed(time.Now().UnixNano())

	bits := [size]int64{8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096}
	
  var randomKeys [size]*big.Int

	fmt.Println(" Number of every n-bit keys")
	for i := 0; i < len(bits); i++ {
		
    numOfKeys := pow(2, bits[i])
		fmt.Printf("Number of keys for n=%d: %s\n", bits[i], numOfKeys)
	}

	fmt.Println("\n Generate a random key from n-bits keys space")
	
  for i := 0; i < len(bits); i++ {
		
    randomKey := generateRandomKey(bits[i])
		randomKeys[i] = randomKey
		
    fmt.Printf("Random key from %d-bits keys space: 0x%x\n", bits[i], randomKey)
	}

	fmt.Println("\n Bruteforce keys until the random key of second point is found")
	for i := 0; i < len(bits); i++ {
		
    timePassed, err := bruteForceKey(randomKeys[i], bits[i])
		
      if err != nil {
			fmt.Printf("Match for n=%d not found\n", bits[i])
		  }   else      {
			fmt.Printf("Brute force for n=%d found a key in %d ms\n",
				bits[i], timePassed.Milliseconds())
		  }
	}
}
