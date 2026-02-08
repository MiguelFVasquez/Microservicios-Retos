# Imagen oficial de Go liviana
FROM golang:1.25-alpine

# Crear directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar go.mod primero (mejor uso de cache)
COPY go.mod ./

# Descargar dependencias
RUN go mod download

# Copiar todo el proyecto
COPY . .

# Compilar el binario
RUN go build -o server ./cmd/server/main.go

# Exponer puerto del microservicio
EXPOSE 8080

# Comando para ejecutar el servidor
CMD ["./server"]
