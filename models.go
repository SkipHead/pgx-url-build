package pgxUrlBuild

type Config struct {
	User               string `json:"user"`
	Password           string `json:"password"`
	Host               string `json:"host,omitempty"`
	Port               string `json:"port,omitempty"`
	DbName             string `json:"db_name"`
	SslMode            string `json:"sslmode,omitempty"`
	PoolMaxConns       string `json:"pool_max_conns,omitempty"`
	MasterHost         string `json:"master_host,omitempty"`
	MasterPort         string `json:"master_port,omitempty"`
	SecondHost         string `json:"second_host,omitempty"`
	SecondPort         string `json:"second_port,omitempty"`
	StatementCacheMode string `json:"statement_cache_mode,omitempty"`
	Schema             string `json:"schema,omitempty"`
}
