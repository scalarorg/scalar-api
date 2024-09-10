package types

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type CovenantParamsType struct {
	CovenantPubkeys []string
	Qorum           int
	Tag             string
	Version         int
}

var CovenantParams CovenantParamsType

func InitCovenantParams() {
	covenantPubkeys := strings.Split(os.Getenv("COVENANT_PUBLIC_KEYS"), ",")
	for i := 0; i < len(covenantPubkeys); i++ {
		if len(covenantPubkeys[i]) != 66 {
			log.Fatalf("Invalid covenant pubkey (string length must be 66): %s", covenantPubkeys[i])
		}
	}

	qorumStr := os.Getenv("COVENANT_QUORUM")
	qorum, err := strconv.Atoi(qorumStr)
	if err != nil {
		log.Fatalf("Invalid qorum: %s", qorumStr)
	}
	if qorum < 1 {
		log.Fatalf("Invalid qorum (must bigger than 0): %d", qorum)
	}
	if qorum > len(covenantPubkeys) {
		log.Fatalf("Invalid qorum (must smaller than or equal to the number of covenant pubkeys): %d", qorum)
	}

	tagStr := os.Getenv("TAG")
	if len(tagStr) != 8 {
		log.Fatalf("Invalid tag (string length must be 8): %s", tagStr)
	}

	versionStr := os.Getenv("VERSION")
	version, err := strconv.Atoi(versionStr)
	if err != nil {
		log.Fatalf("Invalid version: %s", versionStr)
	}

	// Initialize the CovenantParams variable
	CovenantParams = CovenantParamsType{
		CovenantPubkeys: covenantPubkeys,
		Qorum:           qorum,
		Tag:             tagStr,
		Version:         version,
	}
}
