package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"testing"
)

var data = []byte("hogehoge")
var key = []byte("test-key")

func BenchmarkEncodeToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var v string
		for i := 0; i < 100; i++ {
			v = hashToString(data, key)
		}
		_ = v
	}
}

func BenchmarkEncodeToByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var v []byte
		for i := 0; i < 100; i++ {
			v = hashToByte(data, key)
		}
		_ = string(v)
	}
}

func BenchmarkEncodeToStringWithHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := hmac.New(sha256.New, key)
		var v string
		for i := 0; i < 100; i++ {
			v = hashToStringWithHash(h, data, key)
		}
		_ = v
	}
}

func BenchmarkEncodeToByteWithHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := hmac.New(sha256.New, key)
		var v []byte
		for i := 0; i < 100; i++ {
			v = hashToByteWithHash(h, data, key)
		}
		_ = string(v)
	}
}
