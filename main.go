package main

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	broadcast "github.com/Maki-Daisuke/go-broadcast-channel"
	"github.com/gorilla/websocket"
)

//go:embed index.html
var indexHTML []byte

//go:embed popper.html
var popperHTML []byte

//go:embed pop.mp3
var popSound []byte

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var imgBroadcast = broadcast.New[[]byte](10).WithTimeout(2 * time.Second)

func handleRecv(w http.ResponseWriter, r *http.Request) {
	// HTTP接続をWebSocketにアップグレード
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	fmt.Println("Client Connected")

	ch := make(chan []byte)
	imgBroadcast.Subscribe(ch)
	defer close(ch)
	for img := range ch {
		if err := conn.WriteMessage(websocket.TextMessage, img); err != nil {
			log.Println(err)
			return
		}
	}
}

func handleSend(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	img, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	if string(img[0:11]) != "data:image/" {
		log.Printf("Not an image: %s\n", string(img))
		http.Error(w, "Not an image", http.StatusForbidden)
		return
	}
	imgBroadcast.Chan() <- img
	log.Println("Image received")
	w.WriteHeader(http.StatusOK)
	w.Write(([]byte)("OK"))
}

func main() {
	port := "25252"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/recv", handleRecv)
	mux.HandleFunc("/send", handleSend)
	mux.HandleFunc("/popup", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write(popperHTML)
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write(indexHTML)
	})
	mux.HandleFunc("/pop.mp3", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "audio/mpeg")
		w.Write(popSound)
	})
	http.Handle("/", mux)
	log.Printf("Server started on :%s\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
