package main

import (
	"log"
	"time"

	"github.com/prysmaticlabs/go-bls"
)

func main() {
	m := "super special message"
	var aggSig *bls.Sign
	var aggPub *bls.PublicKey
	startTime := time.Now()
	for i := 0; i < 1000; i++ {
		var sec bls.SecretKey
		sec.SetByCSPRNG()
		if i == 0 {
			aggSig = sec.Sign(m)
			aggPub = sec.GetPublicKey()
		} else {
			aggSig.Add(sec.Sign(m))
			aggPub.Add(sec.GetPublicKey())
		}
	}
	endTime := time.Now()
	log.Printf("Time required to sign 1000 messages and aggregate 1000 pub keys and signatures: %f seconds", endTime.Sub(startTime).Seconds())
	log.Printf("Aggregate Signature: 0x%x", aggSig.HexString())
	log.Printf("Aggregate Public Key: 0x%x", aggPub.HexString())

	startTime = time.Now()
	if !aggSig.Verify(aggPub, m) {
		log.Fatal("Aggregate Signature Does Not Verify")
	}
	log.Printf("Aggregate Signature Verifies Correctly!")
	endTime = time.Now()
	log.Printf("Time required to verify aggregate sig: %f seconds", endTime.Sub(startTime).Seconds())
}
