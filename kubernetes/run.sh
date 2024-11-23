kubectl apply -f kubernetes/database/postgresql.yml
kubectl apply -f kubernetes/application/deployment.yml
kubectl apply -f kubernetes/application/grpc-service.yml
kubectl apply -f kubernetes/application/health-service.yml
