package hmacsig

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/google/uuid"
)

func CalculatePassportVerificationSignature(key []byte, nullifier, country, anonymousID string) (string, error) {
	bNull, err := hex.DecodeString(nullifier[2:])
	if err != nil {
		return "", fmt.Errorf("nullifier is not hex: %w", err)
	}

	bAID, err := hex.DecodeString(anonymousID)
	if err != nil {
		return "", fmt.Errorf("anonymousID is not hex: %w", err)
	}

	h := hmac.New(sha256.New, key)
	msg := append(bNull, []byte(country)...)
	msg = append(msg, bAID...)
	h.Write(msg)

	return hex.EncodeToString(h.Sum(nil)), nil
}

func CalculateQREventSignature(key []byte, nullifier, eventID, qrCode string) (string, error) {
	bNull, err := hex.DecodeString(nullifier[2:])
	if err != nil {
		return "", fmt.Errorf("nullifier is not hex: %w", err)
	}

	bID, err := uuid.Parse(eventID)
	if err != nil {
		return "", fmt.Errorf("eventID is not uuid: %w", err)
	}

	bQR, err := base64.StdEncoding.DecodeString(qrCode)
	if err != nil {
		return "", fmt.Errorf("qrCode is not base64: %w", err)
	}

	h := hmac.New(sha256.New, key)
	msg := append(bNull, bID[:]...)
	msg = append(msg, bQR...)
	h.Write(msg)

	return hex.EncodeToString(h.Sum(nil)), nil
}
