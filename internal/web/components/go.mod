module components

go 1.22

require (
	github.com/a-h/templ v0.2.731
	github.com/notnil/chess v1.9.0
	github.com/olahol/melody v1.2.1
)

require github.com/gorilla/websocket v1.5.0 // indirect

require layouts v1.0.0
replace layouts => ../layouts
