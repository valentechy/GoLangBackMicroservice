# GoLangBackMicroservice
Microservicio de prueba utilizado para artículo en www.valentinpalacios.com

Este microservicio expone dos endpoints:
- `/primario`: Devuelve un JSON con el color "blue".
- `/secundario`: Devuelve un JSON con el color "green".

## Pasos para poner en funcionamiento el microservicio

1. Construir la imagen de Docker:
    ```sh
    docker build -t microservicio-backend .
    ```

2. Ejecutar el contenedor localmente (opcional):
    ```sh
    docker run -p 8080:8080 microservicio-backend
    ```

3. Cargar la imagen de Docker en el clúster de kind:
    ```sh
    kind load docker-image microservicio-backend:latest
    ```

4. Aplicar el archivo de despliegue de Kubernetes:
    ```sh
    kubectl apply -f K8s/Deployment.yaml
    ```

5. Verificar que el servicio está corriendo:
    ```sh
    kubectl get pods
    kubectl get services
    ```

6. Acceder al microservicio:
    - Para el endpoint `/primario`: `http://<NodeIP>:<NodePort>/primario`
    - Para el endpoint `/secundario`: `http://<NodeIP>:<NodePort>/secundario`

Nota: `<NodeIP>` y `<NodePort>` se pueden obtener ejecutando `kubectl get services` y buscando el servicio `microservicio-backend`.