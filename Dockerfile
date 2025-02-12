# 1️⃣ Fase de build: compila o Go
FROM golang:1.23.6 AS builder

# Defina o diretório de trabalho correto
WORKDIR /app

# Copie o go.mod e go.sum
COPY go.mod go.sum ./

# Baixe as dependências
RUN go mod tidy

# Copie todo o código-fonte
COPY . .

# Crie o diretório de saída dentro do builder
RUN mkdir -p /app/internal/output

# Compile o binário
RUN go build -o /app/blacklist ./cmd/main.go


# 2️⃣ Fase final: cria a imagem final baseada no Ubuntu
FROM ubuntu:22.04

# Defina o diretório de trabalho
WORKDIR /root/

# Instalar dependências do Ubuntu
RUN apt-get update && \
    apt-get install -y \
    libc6 \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Copiar o binário compilado da fase builder
COPY --from=builder /app/blacklist .

# Expor a porta, caso o app precise
EXPOSE 8080

# Rodar a aplicação
CMD ["./blacklist"]
