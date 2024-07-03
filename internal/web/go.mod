module web

go 1.22

require github.com/a-h/templ v0.2.731

require layouts v1.0.0

require components v1.0.0

require (
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/olahol/melody v1.2.1 // indirect
)

replace layouts => ./layouts

replace components => ./components
