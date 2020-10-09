# First graphql API

Step 1: Create main directory or clone an empty github project then bootstrap it with `go run github.com/99designs/gqlgen init`. It will create some code.

Step 2: Modified `schema.graphqls` in the `graph` directory with you desired schema.

Step 3: Remove `schema.resolvers.go` to avoid error with the new schema updated.

Step 4: Update the existed code with the new schema typing: `go run github.com/99designs/gqlgen generate` or
        Copy this in `resolver.go` between packege and imports: `//go:generate go run github.com/99designs/gqlgen` and then type only `go generate`

Step 5: Include database module, in this case I configured a `Postgres` db with `pg`
