package hmacsig

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestCalculateSignature use it to ensure matching signature on re-implementation
func TestCalculateSignature(t *testing.T) {
	key := "ab6b3f7796728e0df9696c4a3eb600b49b51db9d230e94e9c67fef756d695b63"
	nullifier := "0x973c253a93e8d2e6022721c6a8bd0205940b50cb478d485ca2cbc3354fae95ec"
	country := "UKR"
	anonymousID := "adeef82557bc0f95c8ffe38eca25e4d1d9da79ea14215ec52b4f21370dd60dbc"

	bKey, err := hex.DecodeString(key)
	require.NoError(t, err)
	sig, err := CalculatePassportVerificationSignature(bKey, nullifier, country, anonymousID)
	require.NoError(t, err)
	t.Log("Passport sig:", sig)

	eventID := "18593155-b6a3-4166-80f1-6bf4c5aeedf1"
	qrCode := "iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAABaElEQVR4AWP4//8/AyUYw000"
	sig, err = CalculateQREventSignature(bKey, nullifier, eventID, qrCode)
	require.NoError(t, err)
	t.Log("QR Event sig:", sig)
}
