package types

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type CovenantParamsType struct {
	CovenantPubkeys []string
	Quorum          int
	Tag             string
	Version         int
}

var CovenantParams CovenantParamsType

func initCovenantParams() {
	covenantPubkeys := strings.Split(os.Getenv("COVENANT_PUBLIC_KEYS"), ",")
	for i := 0; i < len(covenantPubkeys); i++ {
		if len(covenantPubkeys[i]) != 66 {
			log.Fatalf("Invalid covenant pubkey (string length must be 66): %s", covenantPubkeys[i])
		}
	}

	quorumStr := os.Getenv("COVENANT_QUORUM")
	quorum, err := strconv.Atoi(quorumStr)
	if err != nil {
		log.Fatalf("Invalid quorum: %s", quorumStr)
	}
	if quorum < 1 {
		log.Fatalf("Invalid quorum (must bigger than 0): %d", quorum)
	}
	if quorum > len(covenantPubkeys) {
		log.Fatalf("Invalid quorum (must smaller than or equal to the number of covenant pubkeys): %d", quorum)
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
		Quorum:          quorum,
		Tag:             tagStr,
		Version:         version,
	}
}

func isZeroCovenantParams() bool {
	return CovenantParams.CovenantPubkeys == nil &&
		CovenantParams.Quorum == 0 &&
		CovenantParams.Tag == "" &&
		CovenantParams.Version == 0
}

func GetCovenantParamsVar() CovenantParamsType {
	if isZeroCovenantParams() {
		initCovenantParams()
	}
	return CovenantParams
}
