package controller

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	graphql "github.com/graph-gophers/graphql-go"
	resolver "github.com/syads321/starter/resolver"
	schemas "github.com/syads321/starter/schema"
	"net/http"
	"os"
)

var (
	// We can pass an option to the schema so we don’t need to
	// write a method to access each type’s field:
	opts = []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	// Schema get schemas

)

type TokenClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type ClientQuery struct {
	OpName    string
	Query     string
	Variables map[string]interface{}
}

// ExecuteQuery for grapgql
func ExecuteQuery(query string, request *http.Request) *graphql.Response {
	ctx := context.Background()
	q1 := ClientQuery{
		Query:     query,
		Variables: nil,
	}
	tokenClaim := TokenClaim{}
	headertoken := request.Header.Get("Token-Key")

	if headertoken != "" {
		token, _ := jwt.ParseWithClaims(headertoken, &TokenClaim{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SIGNING_KEY")), nil
		})
		if token != nil && token.Valid {
			claims := token.Claims.(*TokenClaim)
			tokenClaim.Email = claims.Email
		}

	}

	Schema := graphql.MustParseSchema(schemas.Schema, &resolver.RootResolver{
		Session: tokenClaim.Email,
	}, opts...)
	resp1 := Schema.Exec(ctx, q1.Query, "", q1.Variables)

	return resp1
}
