package main

import (
	crypto "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
)


func main() {

	// This approach is usually used in tests.
	// The reason for doing so is for the
	// reproducibility of the sequence.
	sec1 := rand.New( rand.NewSource(10) )
	sec2 := rand.New( rand.NewSource(10) )
	for i := 0; i < 5; i++ {
		rnd1 := sec1.Int()
		rnd2 := sec2.Int()
		if rnd1 != rnd2 {
			fmt.Println("Rand generated non-equal sequence")
			break
		} else {
			fmt.Printf("Math/Rand1: %d , Math/Rand2: %d\n", rnd1, rnd2)
			// Math/Rand1: 5221277731205826435 , Math/Rand2: 5221277731205826435
			// ..
		}
	}


	// The API uses the Reader to provide the instance of
	// a cryptographically strong pseudo-random generator.
	for i := 0; i < 5; i++ {
		safeNum := NewCryptoRand()
		safeNum2 := NewCryptoRand()
		if safeNum == safeNum2 {
			fmt.Println("Crypto generated equal numbers")
			break
		} else {
			fmt.Printf("Crypto/Rand1: %d , Crypto/Rand2: %d\n",
				safeNum, safeNum2)
			// Crypto/Rand1: 50285 , Crypto/Rand2: 46056
			// ..
		}
	}
}


// ..

func NewCryptoRand() int64 {
	safeNum, err := crypto.Int(crypto.Reader, big.NewInt(100234))
	if err != nil {
		panic(err)
	}

	return safeNum.Int64()
}
