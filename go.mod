module github.com/filmlaunchus/gqlserver

go 1.14

require (
	github.com/buger/jsonparser v1.0.0 // indirect
	github.com/filmlaunchus/gqlserver/server v0.0.0-20200819022312-ef929bff12dd
	github.com/gorilla/mux v1.7.4 // indirect
	github.com/graphql-go/graphql v0.7.9 // indirect
	github.com/rs/xid v1.2.1 // indirect
)

replace (
	github.com/filmlaunchus/gqlserver/datastores => ./datastores
	github.com/filmlaunchus/gqlserver/gql => ./gql
	github.com/filmlaunchus/gqlserver/server => ./server
	github.com/filmlaunchus/gqlserver/utils => ./utils
)
