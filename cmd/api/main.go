package main

import (
	"flag"
	"log"
	"os"

	"github.com/go-logr/zapr"
	"github.com/gobuffalo/pop/v6"
	"github.com/pseudomuto/pseudocms/pkg/models"
	"github.com/pseudomuto/pseudocms/pkg/repo"
	"github.com/pseudomuto/pseudocms/pkg/server"
	"go.uber.org/zap"
)

func main() {
	addr := flag.String("addr", ":0", "the address to bind to")
	flag.Parse()

	zlog, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer zlog.Sync()

	env := "development"
	if v, ok := os.LookupEnv("ENV"); ok {
		env = v
	}

	db, err := pop.Connect(env)
	if err != nil {
		zlog.Fatal(err.Error())
	}

	_, done := server.ListenAndServe(
		*addr,
		server.WithLogger(zapr.NewLogger(zlog).WithName("api-server")),
		server.WithRepoFactory(&repoFactory{db: db}),
	)
	<-done
}

// repoFactory satisfies server.RepoFactory by returning repos that use the db connection.
type repoFactory struct {
	db *pop.Connection
}

func (r *repoFactory) Definitions() server.DefinitionsRepo {
	return repo.New[models.Definition](r.db)
}

func (r *repoFactory) Fields() server.FieldsRepo {
	return repo.New[models.Field](r.db)
}
