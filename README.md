# Ozon Marketplace Workplace API

##Задание 6
Имплементировать для сервиса {domain-kw}-{subdomain}-facade чтение данных из kafka топика в stdout
```zsh
сделано
Убрал вывод в stdout, данные пишутся в БД
```

Имплементировать сохранение прочтенных данных в базу 
```zsh
сделано
```

Добавить поддержку работы с множеством брокеров 
```zsh
сделано
```


##Задание 4

Реализовать методы для интерфейса Repo - 
```zsh
сделано
```

Написать миграции для создания таблиц и создания индексов - 
```zsh
сделано, настройки в пакете migrations, файл 20211107174407_init_db.sql
```

Реализовать методы для интерфейса RepoEvent (сообщения в proto) - 
```zsh
сделано
```

Подготовить dataset для таблиц subdomains и subdomains_events gem - 
```zsh
сделано. Пакет migrations, файл 20211108152300_fill_db.tmp
```

Реализовать поддержку вариаций типов событий на обновление сущности subdomain gem - в процессе

Обеспечить защиту от sql-инъекции gem - 
```zsh
сделано, все запросы работают через параметры
```

Настроить партиципирование таблицы на N частей gem - 
```zsh
сделано, 4 партиции по hash. Настройки в пакете migrations, файл 20211107174407_init_db.sql
```

Написать тесты gem - 
```zsh
сделано, пакет internal/repo
```

## Задание 3. Описание реализации:

Имплементация ручек согласно п.6:
- Вызовы логируются
- Возвращаются пустые ответы

###Базовая версия задания:
Из корня bss-workplace-api запустить docker-compose: 

```zsh
$ docker-compose -f docker-compose-mini.yaml up
```

После запуска проверить работу gRPC можно скриптами: scripts/grpc_server_tests (используется grpcurl)

###Задание 9(со звездочкой):
Из корня bss-workplace-api запустить docker-compose: 

```zsh
$ docker-compose -f docker-compose-traefik.yaml up
```

После запуска проверить работу gateway можно скриптами: scripts/gw_server_tests (используется curl)
Также в этом варианте доступен swagger: Доступ из браузера по URL: http://0.0.0.0:8080/swagger

---

## Build project

### Local

For local assembly you need to perform

```zsh
$ make deps # Installation of dependencies
$ make build # Build project
```
## Running

### For local development

```zsh
$ docker-compose up -d
```

---

## Services

### Swagger UI

The Swagger UI is an open source project to visually render documentation for an API defined with the OpenAPI (Swagger) Specification

- http://localhost:8081

### Grafana:

- http://localhost:3000
- - login `admin`
- - password `MYPASSWORT`

### gRPC:

- http://localhost:8082

```sh
[I] ➜ grpc_cli call localhost:8082 DescribeWorkplaceV1 "id: 1"
connecting to localhost:8082
Rpc failed with status code 5, error message: workplace not found
```

### Gateway:

It reads protobuf service definitions and generates a reverse-proxy server which translates a RESTful HTTP API into gRPC

- http://localhost:8080

```sh
[I] ➜ curl -s -X 'POST' \
  'http://localhost:8080/v1/workplaces' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "id": "1"
}' | jq .
{
  "code": 5,
  "message": "workplace not found",
  "details": []
}
```

### Metrics:

Metrics GRPC Server

- http://localhost:9100/metrics

### Status:

Service condition and its information

- http://localhost:8000
- - `/live`- Layed whether the server is running
- - `/ready` - Is it ready to accept requests
- - `/version` - Version and assembly information

### Prometheus:

Prometheus is an open-source systems monitoring and alerting toolkit

- http://localhost:9090

### Kafka

Apache Kafka is an open-source distributed event streaming platform used by thousands of companies for high-performance data pipelines, streaming analytics, data integration, and mission-critical applications.

- http://localhost:9094

### Kafka UI

UI for Apache Kafka is a simple tool that makes your data flows observable, helps find and troubleshoot issues faster and deliver optimal performance. Its lightweight dashboard makes it easy to track key metrics of your Kafka clusters - Brokers, Topics, Partitions, Production, and Consumption.

- http://localhost:9001

### Jaeger UI

Monitor and troubleshoot transactions in complex distributed systems.

- http://localhost:16686

### Graylog

Graylog is a leading centralized log management solution for capturing, storing, and enabling real-time analysis of terabytes of machine data.

- http://localhost:9000
- - login `admin`
- - password `admin`

### PostgreSQL

For the convenience of working with the database, you can use the [pgcli](https://github.com/dbcli/pgcli) utility. Migrations are rolled out when the service starts. migrations are located in the **./migrations** directory and are created using the [goose](https://github.com/pressly/goose) tool.

```sh
$ pgcli "postgresql://docker:docker@localhost:5432/bss_workplace_api"
```

### Python client

```shell
$ python -m venv .venv
$ . .venv/bin/activate
$ make deps
$ make generate
$ cd pypkg/bss-workplace-api
$ python setup.py install
$ cd ../..
$ docker-compose up -d
$ python scripts/grpc_client.py
```


### Thanks

- [Evald Smalyakov](https://github.com/evald24)
- [Michael Morgoev](https://github.com/zerospiel)
