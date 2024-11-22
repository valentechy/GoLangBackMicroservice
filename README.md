# GoLangBackMicroservice
Microservicio de prueba utilizado para artículo en [www.valentinpalacios.com](https://valentinpalacios.com/post/plataforma-microservicios-herramientas-cncf-parte1-base/).

Este microservicio expone dos endpoints:
- `/primario`: Devuelve un JSON con el color "blue".
- `/secundario`: Devuelve un JSON con el color "green".

## Pasos para poner en funcionamiento el microservicio

1. Construir la imagen de Docker:
    ```sh
    docker build -t microservicio-backend .
    ```

2. Cargar la imagen de Docker en el clúster de kind:

Primero descargamos la imagen en formato .tar:

```sh
podman save frontend-microservice:1.0.0 -o frontend-microservice.tar
```

Y la subimos al registry de kind:

```sh
kind load docker-image microservicio-backend:1.0.0
```

3. Aplicar el archivo de despliegue de Kubernetes:

```sh
kubectl apply -f k8s/
```

4. Verificar que el servicio está corriendo:

```sh
kubectl get pods
kubectl get services
```

5. Acceder al microservicio:

A este pod no se podrá acceder desde fuera del cluster. Se utilizarán las DNS internas para acceder desde el [frontend](https://github.com/valentechy/GoLangFrontMicroservice/tree/1.0.0)

    - Para el endpoint `/primario`: `http://backend-service.default.svc.cluster.local/primario`
    - Para el endpoint `/secundario`: `http://backend-service.default.svc.cluster.local/secundario`

Nota: `<NodeIP>` y `<NodePort>` se pueden obtener ejecutando `kubectl get services` y buscando el servicio `microservicio-backend`.
