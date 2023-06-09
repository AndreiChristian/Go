package handlers

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)

// upgrades the connection
var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// ws josn reponse defines the reponse sent back from the websocket
type wsJsonReponse struct {
	Action      string `json:"action"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {

	ws, err := upgradeConnection.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
	}

	log.Println("client connected to endpoint")

	var response wsJsonReponse
	response.Message = `</em><small>Connected to service</small></em>`

	err = ws.WriteJSON(response)

	if err != nil {
		log.Println(err)
	}

}

func Home(w http.ResponseWriter, r *http.Request) {

	err := renderPage(w, "home.jet", nil)

	if err != nil {
		log.Println(err)
	}

}

func renderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {

	view, err := views.GetTemplate(tmpl)

	if err != nil {
		log.Println(err)
		return err
	}

	err = view.Execute(w, data, nil)

	if err != nil {
		log.Println(err)
		return err
	}

	return err
}
