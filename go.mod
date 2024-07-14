module chessGame

go 1.22

require internal/web v1.0.0

replace internal/web => ./internal/web

require components v1.0.0 // indirect

replace components => ./internal/web/components

require layouts v1.0.0 // indirect

replace layouts => ./internal/web/layouts

require engine v1.0.0 // indirect

require (
	github.com/a-h/templ v0.2.747 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/notnil/chess v1.9.0 // indirect
	github.com/olahol/melody v1.2.1 // indirect
)

replace engine => ./internal/engine
