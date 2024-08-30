module chessGame

go 1.22.5

replace internal/web => ./internal/web

replace components => ./internal/web/components

replace layouts => ./internal/web/layouts

replace engine => ./internal/engine

replace db => ./internal/db

require (
	github.com/a-h/templ v0.2.771 // indirect
	github.com/webview/webview_go v0.0.0-20240220051247-56f456ca3a43 // indirect
)
