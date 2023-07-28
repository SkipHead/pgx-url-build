package pgxUrlBuild

import (
	"testing"
)

var standalone = Config{
	User:               "test",
	Password:           "test",
	Host:               "10.10.10.1",
	Port:               "5454",
	DbName:             "test_db",
	SslMode:            "verify-ca",
	PoolMaxConns:       "100",
	StatementCacheMode: "",
	Schema:             "test_schema",
}

func TestURLStandalone(t *testing.T) {
	got := standalone.makeURL()
	want := "postgres://test:test@10.10.10.1:5454/test_db?sslmode=verify-ca&pool_max_conns=100"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

var cluster = Config{
	User:               "test",
	Password:           "test",
	MasterHost:         "1.1.1.1",
	MasterPort:         "5555",
	SecondHost:         "2.2.2.2",
	SecondPort:         "5555",
	DbName:             "test_db",
	SslMode:            "verify-ca",
	PoolMaxConns:       "100",
	StatementCacheMode: "",
	Schema:             "test_schema",
}

func TestURLCluster(t *testing.T) {
	got := cluster.makeURL()
	want := "postgres://test:test@1.1.1.1:5555,2.2.2.2:5555/test_db?sslmode=verify-ca&pool_max_conns=100"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
