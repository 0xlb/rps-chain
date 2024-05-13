package types

import (
	"crypto/sha256"
	"encoding/hex"
	"regexp"
)

// isValidHash is a helper function to check if
// the provided string is a valid SHA256 hash
func isValidHash(s string) bool {
	pattern := "^[a-fA-F0-9]{64}$"
	match, _ := regexp.MatchString(pattern, s)
	return match
}

// isMoveRevelead is a helper function to compare the commitment
// and the hash of the revealed move and the salt
func isMoveRevelead(commitment string, move string, salt string) bool {
	hash := CalculateHash(move, salt)
	return hash == commitment
}

// CalculateHash is a helper function to calculate the
// sha256 hash to validate the submitted move commitment
func CalculateHash(move string, salt string) string {
	// Concatenate input and salt
	data := []byte(move + salt)

	// Calculate SHA256 hash
	hash := sha256.Sum256(data)

	// Convert hash to hexadecimal string
	hashString := hex.EncodeToString(hash[:])

	return hashString
}