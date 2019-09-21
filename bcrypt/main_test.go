package main

import (
	"fmt"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func BenchmarkBcryptCost(b *testing.B) {
	// cost:4-15
	for cost := 4; cost <= 15; cost++ {
		b.Run(fmt.Sprintf("cost=%d", cost), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bcrypt.GenerateFromPassword([]byte("test"), cost)
			}
		})
	}
}
