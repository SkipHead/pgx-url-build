package pgxUrlBuild

import (
	"strings"
)

func (c *Config) options() string {
	var options []string
	if c.SslMode != "" {
		options = append(options, "sslmode="+c.SslMode)
	}
	if c.PoolMaxConns != "" {
		options = append(options, "pool_max_conns="+c.PoolMaxConns)
	}
	if c.StatementCacheMode != "" {
		options = append(options, "statement_cache_mode="+c.StatementCacheMode)
	}

	return strings.Join(options, "&")
}

func (c *Config) makeURL() string {
	var url string
	if c.Host != "" {
		url = "postgres://" +
			c.User + ":" +
			c.Password + "@" +
			c.Host + ":" +
			c.Port + "/" +
			c.DbName
	} else {
		url = "postgres://" +
			c.User + ":" +
			c.Password + "@" +
			c.MasterHost + ":" +
			c.MasterPort + "," +
			c.SecondHost + ":" +
			c.SecondPort + "/" +
			c.DbName
	}
	if c.options() != "" {
		url = url + "?" + c.options()
	}
	return url
}
