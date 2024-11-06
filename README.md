# GoLangBackMicroservice
Microservicio de prueba utilizado para artículo en www.valentinpalacios.com

```
docker build -t microservicio-backend .
```

```
docker run -p 8080:8080 microservicio-backend
```

```
kind load docker-image microservicio-backend:latest
```

```
kubectl apply -f deployment.yaml
```