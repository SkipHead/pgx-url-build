# pgx-url-build
Build postgres url string from json


1. Example json format config standalone

    ```json
    {
        "user": "test",
        "password": "test",
        "host":"10.10.10.1",
        "port": "5454",
        "db_name":"test_db",
        "sslmode": "verify-ca",
        "pool_max_conns":"100",
        "statement_cache_mode":  "",
        "schema": "test_schema"
    }
    ```
output: postgres://test:test@10.10.10.1:5454/test_db?sslmode=verify-ca&pool_max_conns=100


2. Example json format config cluster

    ```json
    {
      "user": "test",
      "password": "test",
      "master_host":"1.1.1.1",
      "master_port": "5555",
      "second_host": "2.2.2.2",
      "second_port": "5555",
      "db_name":"test_db",
      "sslmode": "",
      "pool_max_conns":"100",
      "statement_cache_mode":  "",
      "schema": "test_schema"
    }
    ```
output: postgres://test:test@1.1.1.1:5555,2.2.2.2:5555/test_db?pool_max_conns=100