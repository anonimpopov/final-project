package manager

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"math"
	"offering/internal/config"
)

const hashConstant int = 17
const hashMod int = 9859

type Manager struct {
	Cfg *config.Config
}

func NewManager(cfg *config.Config) *Manager {
	return &Manager{Cfg: cfg}
}

//type offeringID struct {
//	OfferingID string `json:"offering_id"`
//}

type Answer struct {
	ID       string  `json:"id"`
	FROM     string  `json:"from"`
	TO       string  `json:"to"`
	ClientID int     `json:"client_id"`
	Price    float64 `json:"price"`
}

type CreateRequest struct {
	FROM     string `json:"from"`
	TO       string `json:"to"`
	ClientID int    `json:"client_id"`
}

type CreateResponse struct {
	ID       string  `json:"id"`
	FROM     string  `json:"from"`
	TO       string  `json:"to"`
	ClientID int     `json:"client_id"`
	Price    float64 `json:"price"`
}

func hashString(s string) float64 {
	var hash int = 0
	var power int = 1
	for i := 0; i < len(s); i++ {
		hash += int(s[i]) * power % hashMod
		power *= hashConstant % hashMod
	}

	return float64(hash)
}

func GeneratePrice(from string, to string) float64 {
	return math.Abs(hashString(from) - hashString(to))
}

func JwtPayloadFromRequest(tokenString string, secret string) (jwt.MapClaims, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		fmt.Println(err)
	}

	return claims, true
}
