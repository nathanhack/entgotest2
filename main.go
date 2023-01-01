package main

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	"entgo.io/ent/examples/fs/ent/migrate"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nathanhack/entgotest2/ent"
	"github.com/nathanhack/entgotest2/ent2"
)

func main() {
	client, err := ent.Open(dialect.SQLite, "./db.sqlite?_fk=1")
	if err != nil {
		log.Fatal(err)
	}
	// Run the migrations.
	if err := client.Schema.Create(context.Background(),
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
	); err != nil {
		log.Fatal(err)
	}

	for i:=0;i<100;i++{
		client.Todo.Create().SetTitle(fmt.Sprintf("Note %v",i)).Save(context.Background())
	}

	client.Close()

	//now with the new one
	client2, err := ent2.Open(dialect.SQLite, "./db.sqlite?_fk=1")
	if err != nil {
		log.Fatal(err)
	}
	// Run the migrations.
	if err := client2.Schema.Create(context.Background(),
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
	); err != nil {
		log.Fatal(err)
	}
}
