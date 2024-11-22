# Usar una imagen base de Go
FROM golang:1.23-alpine

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar el código fuente
COPY . .

# Compilar la aplicación
RUN go build -o /backend-microservice .

# Exponer el puerto
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["/backend-microservice"]
