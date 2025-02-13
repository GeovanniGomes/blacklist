# Blacklist Project

## Description
Blacklist system to manage who can and cannot buy tickets at specific events or globally, ensuring flexibility, performance and scalability.

## Tecnologias Utilizadas
- **Go**
- **Gin** Web framework
- **RaabbitMQ** Is a Message Broker (i.e. a message intermediary) that allows different parts of a system to communicate in an asynchronous and decoupled way
- **PostgreSQL**  Open-source relational database
- **Redis** Open-source, high-performance, in-memory key-value data warehouse. It is widely used for use cases that require very high performance and very low latency,
- **Dig** Is a Dependency Injection (DI) library for Go, developed by Uber. Dig helps you manage dependencies declaratively, allowing you to register types and their dependencies, and the framework itself takes care of building the necessary objects.
- **MinIO** repository to store files.
- **Prometheus** monitoring and alerting focused on system and application metrics.
- **Grafana** open-source platform for visualizing and analyzing metrics, allowing you to create interactive dashboards from different data sources.

## Features
- Add blaacklist
- Chek blacklist
- Remove blacklist
- generate reports 

## Installation
1. Clone reposiitory:
   ```sh
   git clone https://github.com/GeovanniGomes/blacklist.git
   cd blacklist
   ```

2. Rum appllication:
   ```sh
   docker-compose --build
   ```

## Setup Tools 

### Step configuration Grafana
Go to Granafa to configure metrics views. url ```http://localhost:3000```

Log in with the default credentials (admin/admin).

You’ll be prompted to change the password; you can set a new one or skip.

Access side menu >> Connections >> data sources.

Click button Add new data source.
Select option prometheus.  In connections input value ```http://prometheus:9090```. And click in Save & Test. 

#### Create dashboards
Access side menu >>  Dashboards >> click New >> New Dashboard >> Add visualization

Select data source prometheus previously configured.

Input ```sum(blacklist_requests_total)``` in the PromQL query field. Save with title **Total All Requests**.


Repeat process created new Dashboard to query ```sum by (path)(blacklist_requests_total)``` in the PromQL query field . Save with title **Total Request Endpoint**


Repeat process created new Dashboard to query ```sum by (path)(blacklist_requests_errors_total)``` in the PromQL query field . Save with title **Total Request Error Endpoint**

## Uso
- To add an item to the blacklist via API:
  ```sh
  curl -X POST http://localhost:8000/api/v1/blacklist -H "Content-Type: application/json" -d ' {"user_identifier": 11,"event_id": "8a0b0a37-d739-4234-a356-8948e2cf2a57","scope": "global","document": "test","reason": "Fraude detectada"}
  ```

- To check if an item is blacklisted:
  ```sh
  curl -X GET http://localhost:8000/api/v1/blacklist/check?user_identifier=11&event_id=8a0b0a37-d739-4234-a356-8948e2cf2a57
  ```

- To remove a user from the blacklist:
  ```sh
  curl -X DELETE http://localhost:8000/api/v1/blacklist/remove -H "Content-Type: application/json" -d '{"user_identifier": 11,"event_id": "8a0b0a37-d739-4234-a356-8948e2cf2a57"}'
  ```

### Report generated

To access report generated go application  url ```play.min.io```
Log in with the default credentials (minioadmin/minioadmin).

After access go side menu >> Object Browser >> search for ```file-blacklist```  >> click no bucket >> select report generated >> Download

## Licença
Este projeto está licenciado sob a [MIT License](LICENSE).

