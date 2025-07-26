
# escolhi a tecnologia, versao dela e o SO
FROM golang:1.24.2-alpine

# colei ela dentro uma pasta, ou seja, um diretorio que se chama app
COPY . /app

# apartir daqui irei rodar todos comandos dentro desse diretorio app
WORKDIR /app

# o comando que roda o golang
CMD go run main.go