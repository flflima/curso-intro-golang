# Comandos aprendidos:

## Para executar a aplicação
```go
package main

func main() // deve estar no pacote main
```
```go
go run main.go
```

## Para criar um módulo

```go 
go mod init modulo 
```

## Para buildar um módulo
```go
go build
```
```console
./modulo
```

## Para builadr mas enviar a pasta <dir_instalacao>/bin
```go
go install
```

## Para adicionar um pacote externo
```go
go get github.com/badoux/checkmail
```

## Para remover um pacote externo
```go
go mod tidy
```