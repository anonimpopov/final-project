package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"io"
	"net/http"
	"offering/internal/manager"
	"offering/internal/models"
)

type Controller struct {
	manager *manager.Manager
}

func NewController(man *manager.Manager) *Controller {
	return &Controller{manager: man}
}

func (c *Controller) ParseOffer(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		offId := r.URL.Query().Get("id")

		fmt.Println(offId)

		payload, ok := c.manager.JwtPayloadFromRequest(offId, c.manager.Cfg.JWT)
		if !ok {
			fmt.Println("Bad token")
		}

		orderJson, ok := payload["order"].(string)
		if !ok {
			fmt.Println("incorrect data in token")
		}

		w.Write([]byte(orderJson))
	}
}

func (c *Controller) CreateOffer(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		bytesBody, err := io.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("Плохое тело запроса"))
			fmt.Println(err)
		}

		var respStruct models.Offer
		err = json.Unmarshal(bytesBody, &respStruct)
		if err != nil {
			fmt.Println(err)
		}

		respStruct.Price.Amount = manager.GeneratePrice(respStruct.FROM, respStruct.TO)
		respStruct.Price.Currency = "$"

		fmt.Println(respStruct)

		bytes, err := json.Marshal(respStruct)
		if err != nil {
			fmt.Println("a")
		}

		payload := jwt.MapClaims{
			"order": string(bytes),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

		jwtKey := c.manager.Cfg.JWT
		t, err := token.SignedString([]byte(jwtKey))
		if err != nil {
			fmt.Println("Bad signing")
		}

		w.Write([]byte(t))
	}
}
