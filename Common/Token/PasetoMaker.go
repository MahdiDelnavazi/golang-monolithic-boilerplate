package token

import (
	"log"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

// PasetoMaker is a PASETO Token maker
type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

var maker Maker
var err error

func NewPasetoMaker(symmetricKey string) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		log.Fatal("cannot create token: %W", "invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker = &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

}

// CreateToken creates a new Token for a specific username and duration
func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", nil
	}

	return maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
}

// VerifyToken checks if the Token is valid or not
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
