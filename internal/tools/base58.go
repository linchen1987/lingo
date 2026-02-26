package tools

import "github.com/btcsuite/btcd/btcutil/base58"

func Base58Encode(input string) string {
	return base58.Encode([]byte(input))
}

func Base58Decode(input string) (string, error) {
	decoded, _, err := base58.CheckDecode(input)
	if err == nil {
		return string(decoded), nil
	}

	decodedRaw := base58.Decode(input)
	if len(decodedRaw) == 0 && len(input) > 0 {
		return "", err
	}

	return string(decodedRaw), nil
}
