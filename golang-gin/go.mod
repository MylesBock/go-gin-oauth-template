module golang-gin

go 1.13

require (
	app v0.0.0-00010101000000-000000000000
	auth v0.0.0-00010101000000-000000000000 // indirect
	callback v0.0.0-00010101000000-000000000000
	github.com/codegangsta/negroni v1.0.0
	github.com/coreos/go-oidc v2.1.0+incompatible
	github.com/gorilla/mux v1.7.3
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	home v0.0.0-00010101000000-000000000000
)

replace app => ./app

replace auth => ./auth

replace home => ./routes/home

replace callback => ./routes/callback

replace templates => ./routes/templates
