module github.com/blocky/basm-go-sdk/example

go 1.22.6

require (
	github.com/blocky/basm-go-sdk v0.0.0-20250228222656-2d98e5763856
	github.com/mailru/easyjson v0.9.0 // indirect
)

require github.com/josharian/intern v1.0.0 // indirect

replace github.com/blocky/basm-go-sdk => ../ // use the local version of the SDK
