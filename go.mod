module chess

go 1.22

require github.com/notnil/chess v1.9.0

require internal/web v1.0.0

require (
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/olahol/melody v1.2.1 // indirect
)

replace internal/web => ./internal/web
