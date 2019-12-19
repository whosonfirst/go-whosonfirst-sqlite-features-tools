package main

import (
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-github/organizations"
	_ "github.com/whosonfirst/go-whosonfirst-index-git"
	_ "github.com/whosonfirst/go-whosonfirst-index-sqlite"
	"github.com/whosonfirst/go-whosonfirst-sqlite-features-index"
	"github.com/whosonfirst/go-whosonfirst-sqlite-features/tables"
	"github.com/whosonfirst/go-whosonfirst-sqlite/database"
	"log"
	"time"
)

type Inspector struct {
	Owner string
}

func (i *Inspector) Inspect(repos ...string) {

	// how many of these can we do concurrently?
	
	for _, repo := range repos {

		err := i.InspectRepo(repo)

		if err != nil {
			log.Printf("[WARNING] '%s' failed indexing because: %v\n", repo, err)
		}
	}
}

func (i *Inspector) InspectRepo(repo string) error {

	t1 := time.Now()
	log.Printf("[INFO] Inspect %s at %v\n", repo, t1)
	
	defer func() {
		log.Printf("[INFO] Time to inspect %s, %v\n", repo, time.Since(t1))
	}()
	
	db, err := database.NewDBWithDriver("sqlite3", ":memory:")

	if err != nil {
		return err
	}
	
	defer db.Close()
	
	err = db.LiveHardDieFast()

	if err != nil {
		return err
	}
	
	db_tables, err := tables.CommonTablesWithDatabase(db)

	if err != nil {
		return err
	}
	
	opts := index.DefaultSQLiteFeaturesIndexerCallbackOptions()
	cb := index.SQLiteFeaturesIndexerCallback(opts)

	idx, err := index.NewSQLiteFeaturesIndexerWithCallback(db, db_tables, cb)

	if err != nil {
		return err
	}

	repo_uri := fmt.Sprintf("git@github.com:%s/%s.git", i.Owner, repo)
	to_index := []string{ repo_uri }

	return idx.IndexPaths("git://", to_index)
}

func main() {

	org := flag.String("org", "whosonfirst-data", "The name of the organization to clone repositories from")
	prefix := flag.String("prefix", "whosonfirst-data", "Limit repositories to only those with this prefix")
	exclude := flag.String("exclude", "", "Exclude repositories with this prefix")
	token := flag.String("token", "", "A valid GitHub API access token")

	flag.Parse()

	opts := organizations.NewDefaultListOptions()

	opts.Prefix = *prefix
	opts.Exclude = *exclude
	opts.NotForked = true
	opts.AccessToken = *token

	repos, err := organizations.ListRepos(*org, opts)

	if err != nil {
		log.Fatal(err)
	}

	i := &Inspector{
		Owner: *org,
	}
	
	i.Inspect(repos...)
}
