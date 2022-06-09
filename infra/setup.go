package infra

import (
	"context"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/kcmvp/idm/ent"
	"github.com/kcmvp/idm/ent/migrate"
	"github.com/kcmvp/profile"
	"log"
)

var Client *ent.Client

func SetupClient() error {
	var err error
	ds, err := profile.GetDatasource()
	url, err := ds.DsName()
	Client, err = ent.Open(ds.Driver, url)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer Client.Close()
	// Run the auto migration tool.
	if err = Client.Schema.Create(context.Background(),
		migrate.WithGlobalUniqueID(true),
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
		schema.WithAtlas(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return err
}
