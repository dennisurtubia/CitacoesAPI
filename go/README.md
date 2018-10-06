# CitacoesAPI
API criada para retornar citações famosas. No momento estão disponíveis 1035 citações.

## Requisitos
```
[Go](https://golang.org/dl/)
```

## Implementação
Para rodar o projeto, adicione a pasta do projeto na variavel GOPATH
```
$ export GOPATH=$../../CitacoesAPI/go
```
Inicie o servidor:
```
go run main
```

## Endpoints
Retorna uma citação aleatória
```
http://localhost:8080/random
```
