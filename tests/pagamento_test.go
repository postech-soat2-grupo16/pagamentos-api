package tests

import (
	"fmt"
	"github.com/cucumber/godog"
	"github.com/joaocampari/postech-soat2-grupo16/api"
	"github.com/joaocampari/postech-soat2-grupo16/tests/tutils"
	"net/http"
	"os"
	"testing"
)

var baseURL string

func TestFeatures(t *testing.T) {
	server := setup()
	defer server.Close()

	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func setup() *http.Server {
	os.Setenv("DATABASE_URL", "host=localhost user=postgres password=postgres dbname=payments_db sslmode=disable TimeZone=UTC")
	db := api.SetupDB()
	sqs := api.SetupQueue()
	r := api.SetupRouter(db, sqs)

	server := http.Server{
		Handler: r,
	}
	serverAddress := tutils.StartNewTestServer(&server)
	baseURL = fmt.Sprintf("http://%s", serverAddress)

	return &server
}
