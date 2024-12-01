package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        // Permitir cualquier origen (desactivado en producción por razones de seguridad)
        return true
    },
}

func handler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Error al actualizar la conexión:", err)
        return
    }
    defer conn.Close()

    for {
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            fmt.Println("Error al leer el mensaje:", err)
            return
        }
        fmt.Printf("Mensaje recibido: %s\n", p)

        if err := conn.WriteMessage(messageType, p); err != nil {
            fmt.Println("Error al escribir el mensaje:", err)
            return
        }
    }
}

func main() {
    http.HandleFunc("/ws", handler)
    fmt.Println("Servidor WebSocket corriendo en ws://localhost:8080/ws")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error al iniciar el servidor:", err)
    }
}
