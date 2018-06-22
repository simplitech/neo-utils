package neoutils

import (
	"fmt"

	"github.com/o3labs/ont-mobile/ontmobile"
)

func OntologyTransfer(endpoint string, wif string, asset string, to string, amount float64) (string, error) {
	raw, err := ontmobile.Transfer(wif, asset, to, amount)
	if err != nil {
		return "", err
	}

	txid, err := ontmobile.SendRawTransaction(endpoint, fmt.Sprintf("%x", raw.Data))
	if err != nil {
		return "", err
	}

	return txid, nil
}
