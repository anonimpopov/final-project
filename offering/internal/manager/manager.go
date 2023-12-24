package manager

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"math"
	"offering/internal/config"
	"offering/internal/models"
)

const hashMod int = 9859

type Manager struct {
	Cfg *config.Config
}

func NewManager(cfg *config.Config) *Manager {
	return &Manager{Cfg: cfg}
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

func GeneratePrice(from models.Location, to models.Location) int {
	x := int(math.Abs(from.Lat + from.Lng - to.Lng - to.Lat))
	return (x*31)%hashMod + 100
}

func (man *Manager) JwtPayloadFromRequest(tokenString string, secret string) (jwt.MapClaims, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		fmt.Println(err)
	}

	return claims, true
}
