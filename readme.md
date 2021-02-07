# Golang e RabbitMQ no Minikube

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
http://172.17.0.2:31112 # Colar a primeira URL gerada no navegador
http://172.17.0.2:31111

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