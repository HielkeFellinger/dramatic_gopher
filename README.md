# Dramatic Gopher - TTRPG Engine - V0

- Old version: [go-dnd](https://github.com/HielkeFellinger/go-dnd)

This will, most likely, be nothing more than a simple gui+engine for a ""D&D""/TTRPG oneshot,
Build for entertainment (On LAN), with friends and getting more familiar with Go(TH) and (alternative) patterns like a 
`Entity Component System`

Is it possible to play a game currently? No, this iteration is ~ 11% on the way there

### Development

GoTH - stack
- [(Go)](https://go.dev/), [(T)empl](https://github.com/a-h/templ) + [(H)tmx](https://htmx.org)
- Extra: [Gin Web Framework](https://gin-gonic.com/) + [Gorilla](https://pkg.go.dev/github.com/gorilla/websocket@v1.5.3) (Main Framework + Websockets)

Recommendations:
- [Templ](https://templ.guide/): install/update: ```go install github.com/a-h/templ/cmd/templ@latest``` run: ```templ generate --watch```
- [Air](https://github.com/air-verse/air): install/update: ```go install github.com/air-verse/air@latest``` run ```air```