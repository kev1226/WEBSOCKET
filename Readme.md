# WebSocket Hola Mundo en Go

Este proyecto es un ejemplo básico de un servidor WebSocket utilizando Go y la biblioteca Gorilla WebSocket.

## Pasos para ejecutar el proyecto

1. **Clona el proyecto desde GitHub**:
   ```bash
   git clone <URL-DE-TU-REPOSITORIO>
   ```

2. **Inicializa un módulo de Go** (si no lo has hecho):
   ```bash
   go mod init websocket-go-hello-world
   ```

3. **Instala la dependencia** `github.com/gorilla/websocket`:
   ```bash
   go get github.com/gorilla/websocket
   ```

4. **Ejecuta el servidor**:
   ```bash
   go run main.go
   ```

5. **Inicia `ngrok` para exponer el servidor de forma pública**:
   ```bash
   ./ngrok.exe http 8080
   ```

6. **Obtén la URL pública de `ngrok`**:
   ## Uso de `ngrok`

    1. **Configura tu token de autenticación**:
   ```bash
   ./ngrok.exe config add-authtoken <tu-token-de-autenticacion>
   ```

    2. **Expon el puerto 8080**:
   ```bash
   ./ngrok.exe http 8080
   ```

    3. **Usa la URL generada** para acceder a tu API desde cualquier lugar.

   - `ngrok` generará una URL como `https://9322-201-46-114-33.ngrok-free.app`. Usa esta URL para probar la conexión WebSocket.

7. **Configura la interfaz web (`ngrok.html`)**:
   - Abre `ngrok.html` y reemplaza la URL de la conexión WebSocket con la URL pública de `ngrok`:
     ```javascript
     const ws = new WebSocket('wss://9322-201-46-114-33.ngrok-free.app/ws');
     ```

8. **Prueba la conexión**:
   - Abre `index.html` en tu navegador.
   - Abre `ngrok.html` en tu navegador para probar ngrok.
   - Escribe un mensaje en la interfaz y haz clic en "Enviar".