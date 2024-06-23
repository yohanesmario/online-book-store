package session

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yohanesmario/online-book-store/conf"
)

type SessionData struct {
	UserID    int32     `json:"user_id"`
	Email     string    `json:"email"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (s SessionData) ToJWTString() (string, error) {
	jsonSession, err := json.Marshal(s)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.RegisteredClaims{
		Subject:   base64.StdEncoding.EncodeToString(jsonSession),
		ExpiresAt: &jwt.NumericDate{Time: s.ExpiredAt},
	})

	jwtString, err := token.SignedString([]byte(conf.Configuration.JWTSecret))
	if err != nil {
		return "", err
	}
	return jwtString, nil
}

func (s SessionData) IsExpired() bool {
	return time.Now().After(s.ExpiredAt)
}

func FromJWTToken(token *jwt.Token) (*SessionData, error) {
	encodedClaimSubject, err := token.Claims.GetSubject()
	if err != nil {
		return nil, err
	}
	decodedClaimSubject, err := base64.StdEncoding.DecodeString(encodedClaimSubject)
	if err != nil {
		return nil, err
	}
	var sessionData SessionData
	err = json.Unmarshal(decodedClaimSubject, &sessionData)
	if err != nil {
		return nil, err
	}
	return &sessionData, nil
}
