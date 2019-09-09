package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
)

const origin = "hogehoge"

/* PHP
<?php
$data = 'hogehoge'
for ($i = 0; $i < 100; $i++) {
    $data = hash_hmac('sha256', $data, 'test-key');
}
echo $data;
*/
func main() {
	data := []byte(origin)
	key := []byte("test-key")

	// hashToString
	for i := 0; i < 100; i++ {
		data = []byte(hashToString(data, key))
	}
	fmt.Println(string(data))
	data = []byte(origin)

	// hashToByte
	for i := 0; i < 100; i++ {
		data = hashToByte(data, key)
	}
	fmt.Println(string(data))
	data = []byte(origin)

	h := hmac.New(sha256.New, key)
	// hashToStringWithHash
	for i := 0; i < 100; i++ {
		data = []byte(hashToStringWithHash(h, data, key))
	}
	fmt.Println(string(data))
	data = []byte(origin)

	// hashToByteWithHash
	for i := 0; i < 100; i++ {
		data = hashToByteWithHash(h, data, key)
	}
	fmt.Println(string(data))
	data = []byte(origin)
}

func hashToString(data, key []byte) string {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func hashToByte(data, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	data = h.Sum(nil)
	dst := make([]byte, hex.EncodedLen(len(data)))
	hex.Encode(dst, data)
	return dst
}

func hashToStringWithHash(h hash.Hash, data, key []byte) string {
	defer h.Reset()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func hashToByteWithHash(h hash.Hash, data, key []byte) []byte {
	defer h.Reset()
	h.Write(data)
	data = h.Sum(nil)
	dst := make([]byte, hex.EncodedLen(len(data)))
	hex.Encode(dst, data)
	return dst
}
