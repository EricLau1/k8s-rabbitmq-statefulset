# K8s RabbitMQ Statefulset & Golang

- Fazer o deploy do RabbitMQ

```bash
cd k8s

kubectl apply -f statefulset.yaml

kubectl apply -f service.yaml
```

- Acessar a dashboard do RabbitMQ

```bash
minikube service rabbitmq --url

# exemplo de saida
http://172.17.0.2:31112
http://172.17.0.2:31111

# Colar a primeira gerada URL no navegador
```

- Fazer o build da imagem

```bash
cd app

sh build.sh
```

- Fazer o deploy da imagem no minikube

```bash
cd app/deploy

sh deploy.sh

# visualizar os logs na interface
minikube dashboard
```