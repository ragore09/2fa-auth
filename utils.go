package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"strconv"
	"strings"
	"time"
)

func getTOTP(secretKey string, timestamp int64) string {
	// the secret key is in base32, so we may decode it first
	key, err := base32.StdEncoding.DecodeString(strings.ToUpper(secretKey))
	check(err)
	//in this point, we know 'key' is an array of 20 bytes
	binaryTimestamp := make([]byte, 8)
	binary.BigEndian.PutUint64(binaryTimestamp, uint64(timestamp))
	hashKey := hmac.New(sha1.New, key)
	hashKey.Write(binaryTimestamp)
	hash := hashKey.Sum(nil)
	offset := hash[19] & 15 // we take the last 4 bits of the last byte
	var header uint32
	reader := bytes.NewReader(hash[offset : offset+4])
	err = binary.Read(reader, binary.BigEndian, &header)
	check(err)
	h12 := (int(header) & 0x7fffffff) % 1000000 //Converts number as a string
	otp := strconv.Itoa(int(h12))
	return otp
}

// GetTOTP gets a 2FA token for the given secret
func GetTOTP(secretKey string) string {
	interval := time.Now().Unix() / 30 // the OTP will live only during 30 seconds
	return getTOTP(secretKey, interval)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
