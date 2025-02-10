# Use a imagem oficial do Go como base para compilar o código
FROM golang:1.23.6 AS builder

# Defina o diretório de trabalho
WORKDIR /app

# Copie o go.mod e go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Baixe as dependências
RUN go mod tidy

# Copie todo o código-fonte, incluindo os arquivos da pasta cmd e o restante da estrutura
COPY . .

# Compile o aplicativo Go
RUN go build -o blacklist ./cmd/main.go

# Crie a imagem final com Ubuntu 22.04
FROM ubuntu:22.04

# Defina o diretório de trabalho
WORKDIR /root/

# Instalar dependências do Ubuntu, como GLIBC e outras
RUN apt-get update && \
    apt-get install -y \
    libc6 \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Copie o binário compilado da etapa anterior
COPY --from=builder /app/blacklist .

# Exponha a porta, caso o seu app precise
EXPOSE 8080

# Defina o comando para rodar a aplicação
CMD ["./blacklist"]
