module github.com/ozonmp/bss-workplace-api

go 1.16

require (
	github.com/Masterminds/squirrel v1.5.1
	github.com/Shopify/sarama v1.30.0
	github.com/golang/protobuf v1.5.2
	github.com/jackc/pgx/v4 v4.13.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/lib/pq v1.10.3
	github.com/ozonmp/bss-workplace-facade/pkg/bss-workplace-facade v0.0.0-20211120092448-783cfaf20bcb
	github.com/rs/zerolog v1.24.0
	golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace github.com/ozonmp/bss-workplace-facade/pkg/bss-workplace-facade => ./pkg/bss-workplace-facade
