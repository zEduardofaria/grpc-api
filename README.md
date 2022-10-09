# Exemplo de implementação de gRPC

## [**Link para o curso completo**](https://fullcycle.com.br/)

Essa API GRPC faz parte do treinamento do professor Wesley Willians.

1. Suba o container rodando:
```
docker-compose up -d
```

2. Acesse o container:
```
docker-compose exec app sh
```

3. Caso queira gerar novamente as stubs:
```
protoc --proto_path=proto/ proto/*.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc --go-grpc_out=. --go_out=.
```