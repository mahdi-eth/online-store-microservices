kubectl apply -f kubernetes/pvcs/product-pvc.yaml
kubectl apply -f kubernetes/pvcs/order-pvc.yaml
kubectl apply -f kubernetes/pvcs/redis-pvc.yaml

kubectl apply -f kubernetes/deployments/product-db-deployment.yaml
kubectl apply -f kubernetes/services/product-db-service.yaml

kubectl apply -f kubernetes/deployments/order-db-deployment.yaml
kubectl apply -f kubernetes/services/order-db-service.yaml

kubectl apply -f kubernetes/deployments/inventory-redis-deployment.yaml
kubectl apply -f kubernetes/services/inventory-redis-service.yaml

kubectl apply -f kubernetes/deployments/nats-deployment.yaml
kubectl apply -f kubernetes/services/nats-service.yaml

kubectl apply -f kubernetes/deployments/product-deployment.yaml
kubectl apply -f kubernetes/services/product-service.yaml

kubectl apply -f kubernetes/deployments/order-deployment.yaml
kubectl apply -f kubernetes/services/order-service.yaml

kubectl apply -f kubernetes/deployments/inventory-deployment.yaml
kubectl apply -f kubernetes/services/inventory-service.yaml
