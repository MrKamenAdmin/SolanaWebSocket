package webserver

import (
	"GorillaWebSocket/internal/delivery"
	"GorillaWebSocket/internal/delivery/singleton"
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var m = delivery.Response{}

func StartServer() {
	http.HandleFunc("/", echo)
	http.ListenAndServe(":8081", nil) // Уводим http сервер в горутину
}

func echo(w http.ResponseWriter, r *http.Request) {
	connection, _ := upgrader.Upgrade(w, r, nil)

	go sendMessage(connection)
}

func sendMessage(connection *websocket.Conn) {
	defer connection.Close()

	cache := singleton.GetInstance()

	for {
		marshal, err := json.Marshal(cache.Get())
		if err != nil {
			break
		}
		err = connection.WriteMessage(websocket.TextMessage, marshal)
		//err = connection.WriteJSON(marshal)
		if err != nil {
			break
		}

		time.Sleep(5 * time.Second)
	}

}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Пропускаем любой запрос
	},
}
