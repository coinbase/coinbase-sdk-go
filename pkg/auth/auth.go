package auth

import (
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math"
	"math/big"
	"strings"
	"time"

	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

// APIKeyClaims holds public claim values for a JWT, as well as a URI.
type APIKeyClaims struct {
	*jwt.Claims
	URI string `json:"uri"`
}

// BuildJWT constructs and returns the JWT as a string along with an error, if any.
func BuildJWT(apiKey APIKey, service, uri string) (string, error) {
	privateKey := apiKey.PrivateKey

	// Replace escaped newline sequence
	privateKey = strings.ReplaceAll(privateKey, "\\n", "\n")

	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return "", fmt.Errorf("could not decode private key")
	}

	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("could not parse private key: %w", err)
	}

	sig, err := jose.NewSigner(
		jose.SigningKey{Algorithm: jose.ES256, Key: key},
		(&jose.SignerOptions{NonceSource: nonceSource{}}).WithType("JWT").WithHeader("kid", apiKey.Name),
	)
	if err != nil {
		return "", fmt.Errorf("could not create signer: %w", err)
	}

	cl := &APIKeyClaims{
		Claims: &jwt.Claims{
			Subject:   apiKey.Name,
			Issuer:    "cdp",
			NotBefore: jwt.NewNumericDate(time.Now()),
			Expiry:    jwt.NewNumericDate(time.Now().Add(1 * time.Minute)),
			Audience:  jwt.Audience{service},
		},
		URI: uri,
	}

	jwtString, err := jwt.Signed(sig).Claims(cl).CompactSerialize()
	if err != nil {
		return "", fmt.Errorf("could not sign: %w", err)
	}

	return jwtString, nil
}

// nonceSource implements the jose.NonceSource interface. It is used for building
// the JWT.
type nonceSource struct{}

// Nonce calculates and returns nonce as a string and an error, if any.
func (n nonceSource) Nonce() (string, error) {
	r, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		return "", fmt.Errorf("could not generate nonce: %w", err)
	}

	return r.String(), nil
}
