package handlers

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"io"
	"net/http"
	"offering/internal/manager"
	"offering/internal/models"
)

type Controller struct {
	manager *manager.Manager
	Logger  *zap.Logger
}

func NewController(man *manager.Manager, logger *zap.Logger) *Controller {
	return &Controller{
		manager: man,
		Logger:  logger,
	}
}

func (c *Controller) ParseOffer(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		offId := r.URL.Query().Get("id")

		payload, ok := c.manager.JwtPayloadFromRequest(offId, c.manager.Cfg.JWT)
		if !ok {
			c.Logger.Fatal("Bad token")
		}

		orderJson, ok := payload["order"].(string)
		if !ok {
			c.Logger.Fatal("incorrect data in token")
		}

		w.Write([]byte(orderJson))
	}
}

func (c *Controller) CreateOffer(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		bytesBody, err := io.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("Плохое тело запроса"))
			c.Logger.Fatal(err.Error())
		}

		var respStruct models.Offer
		err = json.Unmarshal(bytesBody, &respStruct)
		if err != nil {
			c.Logger.Fatal(err.Error())
		}

		respStruct.Price.Amount = manager.GeneratePrice(respStruct.FROM, respStruct.TO)
		respStruct.Price.Currency = "$"

		bytes, err := json.Marshal(respStruct)
		if err != nil {
			c.Logger.Fatal(err.Error())
		}

		payload := jwt.MapClaims{
			"order": string(bytes),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

		jwtKey := c.manager.Cfg.JWT
		t, err := token.SignedString([]byte(jwtKey))
		if err != nil {
			c.Logger.Fatal(err.Error())
		}

		w.Write([]byte(t))
	}
}
