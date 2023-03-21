package testdb

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/Pallinder/go-randomdata"
	"github.com/cockroachdb/cockroach-go/v2/testserver"
	"github.com/gobuffalo/pop/v6"
	"github.com/stretchr/testify/require"
	"github.com/tanimutomo/sqlfile"
)

var (
	// projectDir is the directory containing go.mod. This is the working directory
	// used by commands that are run here, allowing us to reference paths relative
	// from the project root.
	projectDir = ""

	pdSync sync.Once
)

// TestDB defines functionality for connecting to an ephemeral test database
// that has the schema preloaded.
type TestDB struct {
	name   string
	server testserver.TestServer
}

// Open opens a new connection to the database.
func (db *TestDB) Open() (*pop.Connection, error) {
	url := db.server.PGURL()
	url.Path = db.name

	conn, err := pop.NewConnection(&pop.ConnectionDetails{
		URL:    url.String(),
		Driver: "postgres",
	})
	if err != nil {
		return nil, err
	}

	return conn, conn.Open()
}

// Close ensures the database server is closed.
func (db *TestDB) Close() {
	db.server.Stop()
}

// New creates a new TestDB using a randomly generated database name. This new database
// has the schema preloaded and all tables are empty.
func New(t *testing.T) *TestDB {
	ts, err := testserver.NewTestServer()
	require.NoError(t, err)

	url := ts.PGURL()
	url.Path = strings.ToLower(randomdata.Noun())
	conn, err := sql.Open("postgres", url.String())
	require.NoError(t, err)
	defer conn.Close()

	// Ensure test db is created
	_, err = conn.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", url.Path))
	require.NoError(t, err)

	// load schema
	s := sqlfile.New()
	s.File(filepath.Join(moduleDir(), "migrations", "schema.sql"))
	_, err = s.Exec(conn)
	require.NoError(t, err)

	return &TestDB{
		name:   url.Path,
		server: ts,
	}
}

// moduleDir finds the directory containing go.mod.
func moduleDir() string {
	pdSync.Do(func() {
		projectDir, _ = os.Getwd()
		projectDir = filepath.Clean(projectDir)

		for {
			if fi, err := os.Stat(filepath.Join(projectDir, "go.mod")); err == nil && !fi.IsDir() {
				break
			}

			d := filepath.Dir(projectDir)
			if d == projectDir {
				break
			}
			projectDir = d
		}
	})

	return projectDir
}
