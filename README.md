# How to run

```go
go run main.go -commit="some hash" -pipeline="some url" -env="dev"
```

# How to test

go test -v .

# Hot to build

go build -o bin/service

# Task description

1) сделать форк проекта
2) добавить в проект докер файл(ы)
3) настроить ci/cd для проекта

## Dockerfile

В результате должен собираться образ, который должен содержать только бинарный файл приложения (`bin/service`) запущенный с переданными параметрами.
Во всех случаях в commit записать хеш текущего коммита и в pipeline нужно записать url на pipeline

## Pipeline

Pipeline содержит 3 stages:

1) [test](#How-to-test) - тесты сервиса
2) [build](#Hot-to-build) - компиляция сервиса
3) push

### Prod pipeline

Тригеррится при коммите в master ветку. Состоит из всех 3 stages. В env параметр нужно передать `prod`.

### Dev pipeline

Тригеррится при коммите в dev ветку. Состоит из всех 3 stages. В env параметр нужно передать `dev`.

### Feature pipeline

Тригеррится при коммите в feature ветку (ветку с префиксом feature). Состоит из stages test и build