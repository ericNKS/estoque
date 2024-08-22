FROM golang:latest
# Defina o diretório de trabalho dentro do container
WORKDIR /app
# Copie os arquivos go.mod e go.sum e baixe as dependências
COPY go.mod go.sum ./
RUN go mod download
# Copie o restante do código
COPY . .
# Compile o binário para a pasta cmd/api
RUN go build -o /binario cmd/api/main.go
# Etapa 2: Construir a imagem final
FROM alpine:latest
# Defina o diretório de trabalho dentro do container
WORKDIR /root/
# Copie o binário da etapa anterior para a imagem final
COPY --from=builder /binario .
# Comando para rodar o binário quando o container iniciar
CMD ["./binario"]