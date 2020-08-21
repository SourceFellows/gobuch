package jwt

import (
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
)

var key = []byte("meinGeheimerSchlüssel")

func CreateToken(user string) ([]byte, error) {
	claims := jws.Claims{}
	claims.SetAudience("source-fellows.com")
	claims.SetIssuer(user)
	claims.SetExpiration(time.Now().Add(time.Minute * 10))
	token := jws.NewJWT(claims, crypto.SigningMethodHS256)
	return token.Serialize(key)
}

func ValidateToken(token []byte) error {
	newToken, err := jws.ParseJWT(token)
	if err != nil {
		return err
	}
	return newToken.Validate(key, crypto.SigningMethodHS256)
}
