package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"io"
	"net/http"
	"offering/internal/manager"
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

		payload, ok := manager.JwtPayloadFromRequest(offId, c.manager.Cfg.JWT)
		if !ok {
			fmt.Println("Bad token")
		}

		fmt.Println(payload)

		ans := manager.Answer{}
		ans.ID = offId
		ans.TO = payload["to"].(string)
		ans.FROM = payload["from"].(string)
		ans.ClientID = int(payload["client_id"].(float64))
		ans.Price = payload["price"].(float64)

		resp, err := json.Marshal(ans)
		if err != nil {
			fmt.Println("Bad marshal")
		}

		w.Write(resp)
	}
}

func (c *Controller) CreateOffer(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		bytesBody, err := io.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("Плохое тело запроса"))
			fmt.Println(err)
		}

		var req manager.CreateRequest
		var respStruct manager.CreateResponse
		err = json.Unmarshal(bytesBody, &req)
		if err != nil {
			fmt.Println(err)
		}

		payload := jwt.MapClaims{
			"from":      req.FROM,
			"to":        req.TO,
			"client_id": req.ClientID,
			"price":     manager.GeneratePrice(req.FROM, req.TO),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

		jwtKey := c.manager.Cfg.JWT
		t, err := token.SignedString([]byte(jwtKey))
		if err != nil {
			fmt.Println("Bad signing")
		}

		respStruct.ID = t
		respStruct.FROM = req.FROM
		respStruct.TO = req.TO
		respStruct.ClientID = req.ClientID
		respStruct.Price = manager.GeneratePrice(req.FROM, req.TO)

		response, err := json.Marshal(respStruct)

		fmt.Println(response)

		if err != nil {
			fmt.Println(err)
		}

		w.Write(response)
	}
}
