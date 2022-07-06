package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"

	"github.com/CavernaTechnologies/pgext"
	"github.com/CavernaTechnologies/pgext/example/database"
	"github.com/jackc/pgx/v4"
)

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/example")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	queries := database.New(conn)

	queries.InsertNum(ctx, pgext.Puint(100000000))
	queries.InsertNum(ctx, pgext.Puint(1123412340000))

	for i := 0; i < 10; i++ {
		r, err := queries.InsertNum(ctx, pgext.Puint(rand.Uint64()))
		if err != nil {
			panic(err)
		}
		fmt.Println(r)
	}

	r, err := queries.GetNums(ctx)
	if err != nil {
		panic(err)
	}

	for _, num := range r {
		fmt.Println(num)
	}
}
