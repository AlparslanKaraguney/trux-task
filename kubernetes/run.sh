kubectl apply -f kubernetes/database/deployment.yml
kubectl apply -f kubernetes/database/storage.yml
kubectl apply -f kubernetes/database/service.yml
kubectl apply -f kubernetes/application/deployment.yml
kubectl apply -f kubernetes/application/grpc-service.yml
kubectl apply -f kubernetes/application/health-service.yml
