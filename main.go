package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/examples/fs/ent/migrate"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nathanhack/entgotest2/ent"
	"github.com/nathanhack/entgotest2/ent2"
)

func main() {
	db, err := sql.Open("sqlite3", "file:./db.sqlite?_fk=1")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	client := ent.NewClient(ent.Driver(db))
	// Run the migrations.
	if err := client.Schema.Create(context.Background(),
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
	); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 100; i++ {
		client.Todo.Create().SetTitle(fmt.Sprintf("Note %v", i)).Save(context.Background())
	}

	// now with the new one
	client2 := ent2.NewClient(ent2.Driver(db))
	if err != nil {
		log.Fatal(err)
	}
	// Run the migrations.
	if err := client2.Schema.WriteTo(context.Background(), os.Stdout,
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
	); err != nil {
		log.Fatal(err)
	}
}
