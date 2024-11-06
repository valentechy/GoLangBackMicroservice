# Usar una imagen base de Go
FROM golang:1.17-alpine

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar el archivo go.mod y go.sum
COPY go.mod ./
COPY go.sum ./

# Descargar las dependencias
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar la aplicación
RUN go build -o /microservicio-backend

# Exponer el puerto
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["/microservicio-backend"]