package pairing

import (
	"bytes"
	"encoding/json"
	"testing"
)

var keyOperations KeyOperations = NewEcdsaKeyOperationsWithRand(onesReader)

func TestPublicKeyRoundTrip(t *testing.T) {
	// Example JWK values taken from RFC 7517.
	key, err := keyOperations.importJwkPublicKey(`
	{
		"kty": "EC",
		"crv": "P-256",
		"x": "MKBCTNIcKUSDii11ySs3526iDZ8AiTo7Tu6KPAqv7D4",
		"y": "4Etl6SRW2YiLUrN5vfvVHuhp7x8PxltmWWlbbM4IFyM"
	}
	`)

	if err != nil {
		t.Error(err)
	}

	exportedKey := key.exportJwk()
	var parsedKey map[string]interface{}
	err = json.Unmarshal([]byte(exportedKey), &parsedKey)
	if err != nil {
		t.Error(err)
	}

	if parsedKey["kty"] != "EC" {
		t.Errorf("Round-trip with invalid kty: %v", parsedKey["kty"])
	}
	if parsedKey["crv"] != "P-256" {
		t.Errorf("Round-trip with invalid crv: %v", parsedKey["crv"])
	}
	if parsedKey["x"] != "MKBCTNIcKUSDii11ySs3526iDZ8AiTo7Tu6KPAqv7D4" {
		t.Errorf("Round-trip with invalid x: %v", parsedKey["x"])
	}
	if parsedKey["y"] != "4Etl6SRW2YiLUrN5vfvVHuhp7x8PxltmWWlbbM4IFyM" {
		t.Errorf("Round-trip with invalid y: %v", parsedKey["y"])
	}
}

func TestPublicKeyImportExtraValues(t *testing.T) {
	_, err := keyOperations.importJwkPublicKey(`
	{
		"kty": "EC",
		"crv": "P-256",
		"x": "MKBCTNIcKUSDii11ySs3526iDZ8AiTo7Tu6KPAqv7D4",
		"y": "4Etl6SRW2YiLUrN5vfvVHuhp7x8PxltmWWlbbM4IFyM",
		"kid": "1",
		"use": "sig"
	}
	`)

	// Extra JSON values should be ignored and not cause failure.
	if err != nil {
		t.Error(err)
	}
}

func TestPublicKeyImportInvalidJson(t *testing.T) {
	_, err := keyOperations.importJwkPublicKey(`blah`)
	if err == nil {
		t.Error("Failed to raise error on invalid JSON.")
	}
}

func TestPublicKeyImportInvalidAlgorithm(t *testing.T) {
	_, err := keyOperations.importJwkPublicKey(`
	{
		"kty": "RSA",
		"crv": "P-256",
		"x": "0",
		"y": "0"
	}
	`)
	if err == nil {
		t.Error("Failed to raise error on invalid algorithm.")
	}
}

func TestPublicKeyImportInvalidCurve(t *testing.T) {
	_, err := keyOperations.importJwkPublicKey(`
	{
		"kty": "EC",
		"crv": "P-384",
		"x": "0",
		"y": "0"
	}
	`)
	if err == nil {
		t.Error("Failed to raise error on invalid curve.")
	}
}

func TestPublicKeyImportInvalidParameters(t *testing.T) {
	_, err := keyOperations.importJwkPublicKey(`
	{
		"kty": "EC",
		"crv": "P-256",
		"x": "0",
		"y": "0"
	}
	`)
	if err == nil {
		t.Error("Failed to raise error on invalid parameters.")
	}
}

func TestGenerateExportPublicKey(t *testing.T) {
	keyPair, err := keyOperations.generateKeyPair()
	if err != nil {
		t.Error(err)
	}

	exportedKey := keyPair.PublicKey.exportJwk()
	var parsedKey map[string]interface{}
	err = json.Unmarshal([]byte(exportedKey), &parsedKey)
	if err != nil {
		t.Error(err)
	}

	if parsedKey["kty"] != "EC" {
		t.Errorf("Export with invalid kty: %v", parsedKey["kty"])
	}
	if parsedKey["crv"] != "P-256" {
		t.Errorf("Export with invalid crv: %v", parsedKey["crv"])
	}
	if parsedKey["x"] != "b_A7lJJBzh2t1DUZ5pYOCoW0GmmgXDKBA6orzhWUyhY" {
		t.Errorf("Export with invalid x: %v", parsedKey["x"])
	}
	if parsedKey["y"] != "PE91OlW_AdxT9sCwx-7ni0DG_30lqW4igrmJzvccFEo" {
		t.Errorf("Export with invalid y: %v", parsedKey["y"])
	}
}

func TestPrivateKeyRoundTrip(t *testing.T) {
	// Example JWK values taken from RFC 7517.
	key, err := keyOperations.importJwkPrivateKey(`
	{
		"kty": "EC",
		"crv": "P-256",
		"x": "MKBCTNIcKUSDii11ySs3526iDZ8AiTo7Tu6KPAqv7D4",
		"y": "4Etl6SRW2YiLUrN5vfvVHuhp7x8PxltmWWlbbM4IFyM",
		"d": "870MB6gfuTJ4HtUnUvYMyJpr5eUZNP4Bk43bVdj3eAE"
	}
	`)

	if err != nil {
		t.Error(err)
	}

	exportedKey := key.exportJwk()
	var parsedKey map[string]interface{}
	err = json.Unmarshal([]byte(exportedKey), &parsedKey)
	if err != nil {
		t.Error(err)
	}

	if parsedKey["kty"] != "EC" {
		t.Errorf("Round-trip with invalid kty: %v", parsedKey["kty"])
	}
	if parsedKey["crv"] != "P-256" {
		t.Errorf("Round-trip with invalid crv: %v", parsedKey["crv"])
	}
	if parsedKey["x"] != "MKBCTNIcKUSDii11ySs3526iDZ8AiTo7Tu6KPAqv7D4" {
		t.Errorf("Round-trip with invalid x: %v", parsedKey["x"])
	}
	if parsedKey["y"] != "4Etl6SRW2YiLUrN5vfvVHuhp7x8PxltmWWlbbM4IFyM" {
		t.Errorf("Round-trip with invalid y: %v", parsedKey["y"])
	}
	if parsedKey["d"] != "870MB6gfuTJ4HtUnUvYMyJpr5eUZNP4Bk43bVdj3eAE" {
		t.Errorf("Round-trip with invalid y: %v", parsedKey["y"])
	}
}

func TestGenerateExportPrivateKey(t *testing.T) {
	keyPair, err := keyOperations.generateKeyPair()
	if err != nil {
		t.Error(err)
	}

	exportedKey := keyPair.PrivateKey.exportJwk()
	var parsedKey map[string]interface{}
	err = json.Unmarshal([]byte(exportedKey), &parsedKey)
	if err != nil {
		t.Error(err)
	}

	if parsedKey["kty"] != "EC" {
		t.Errorf("Export with invalid kty: %v", parsedKey["kty"])
	}
	if parsedKey["crv"] != "P-256" {
		t.Errorf("Export with invalid crv: %v", parsedKey["crv"])
	}
	if parsedKey["x"] != "b_A7lJJBzh2t1DUZ5pYOCoW0GmmgXDKBA6orzhWUyhY" {
		t.Errorf("Export with invalid x: %v", parsedKey["x"])
	}
	if parsedKey["y"] != "PE91OlW_AdxT9sCwx-7ni0DG_30lqW4igrmJzvccFEo" {
		t.Errorf("Export with invalid y: %v", parsedKey["y"])
	}
	if parsedKey["d"] != "AQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQE" {
		t.Errorf("Export with invalid d: %v", parsedKey["d"])
	}
}

func TestGenerateFullRoundTrip(t *testing.T) {
	keyPair, err := keyOperations.generateKeyPair()
	if err != nil {
		t.Error(err)
	}

	// Export then import the public key.
	exportedPublicKey := keyPair.PublicKey.exportJwk()
	importedPublicKey, err := keyOperations.importJwkPublicKey(exportedPublicKey)
	if err != nil {
		t.Error(err)
	}

	// Export then import the private key.
	exportedPrivateKey := keyPair.PrivateKey.exportJwk()
	importedPrivateKey, err := keyOperations.importJwkPrivateKey(exportedPrivateKey)
	if err != nil {
		t.Error(err)
	}

	// Sign a message and then verify with the imported public key.
	signature, err := importedPrivateKey.signMessage("hello")
	if err != nil {
		t.Error(err)
	}

	verified := importedPublicKey.verifyMessage(signature, "hello")
	if !verified {
		t.Errorf("Failed to verify own signature: %v", signature)
	}
}

func TestGenerateNonce(t *testing.T) {
	nonce := GenerateNonceWithRand(onesReader)
	if nonce != "AQEBAQEBAQEBAQEBAQEBAQEB" {
		t.Errorf("Unexpected nonce value: %v", nonce)
	}
}

func TestSignatureEncodingPadded(t *testing.T) {
	// Value found by iterating until we found a value which produced a non-equal length r,s pair when signing "hello".
	constRand := constReader{31}
	keyOperations := NewEcdsaKeyOperationsWithRand(constRand)

	keyPair, err := keyOperations.generateKeyPair()
	if err != nil {
		t.Error(err)
		return
	}

	signature, err := keyPair.PrivateKey.signMessage("hello")
	if err != nil {
		t.Error(err)
		return
	}

	// Value of s is expected to be padded with a zero at index zero.
	rBytes := []byte{252, 192, 250, 114, 162, 172, 104, 96, 223, 157, 84, 160, 73, 142, 70, 223, 32, 57, 52, 26, 90, 10, 136, 159, 142, 225, 102, 17, 90, 246, 168, 110}
	sBytes := []byte{0, 35, 194, 80, 208, 112, 92, 52, 233, 190, 89, 156, 98, 16, 247, 234, 3, 95, 244, 81, 133, 117, 81, 159, 156, 168, 184, 50, 2, 189, 178, 188}
	expectedSignature := append(rBytes, sBytes...)

	if !bytes.Equal(signature, expectedSignature) {
		t.Errorf("Signature '%v' does not match expected padded signature '%v'", signature, expectedSignature)
	}

	verified := keyPair.PublicKey.verifyMessage(signature, "hello")
	if !verified {
		t.Errorf("Failed to verify own signature: %v", signature)
	}
}

func TestVerifyUnpaddedSignature(t *testing.T) {
	// Random source and message found by iterating until we found values which produced an r,s pair both of length 31 bytes.
	// If verify doesn't check r,s padding on the signature, it will verify these r,s values as they are a valid signature, despite not being in the correct padded format.
	constRand := constReader{1}
	message := "iNWPwVjsWF"
	keyOperations := NewEcdsaKeyOperationsWithRand(constRand)

	keyPair, err := keyOperations.generateKeyPair()
	if err != nil {
		t.Error(err)
		return
	}
	rBytes := []byte{153, 180, 197, 99, 63, 229, 132, 85, 154, 146, 112, 217, 156, 206, 59, 246, 85, 49, 80, 38, 154, 99, 229, 104, 111, 208, 94, 161, 219, 204, 177}
	sBytes := []byte{138, 235, 92, 132, 36, 80, 130, 27, 230, 90, 109, 204, 167, 170, 234, 35, 85, 235, 244, 17, 75, 190, 124, 29, 101, 192, 72, 126, 210, 224, 187}
	unpaddedSignature := append(rBytes, sBytes...)

	verified := keyPair.PublicKey.verifyMessage(unpaddedSignature, message)
	if verified {
		t.Errorf("Incorrectly verified unpadded signature!")
	}
}
