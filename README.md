Затрачано: 3 часа 

Auth
> curl -X POST http://localhost:8080/api/auth/simple \
>   -H "X-User-Id: user_123"

RTC Token
> curl -X POST http://localhost:8080/api/tokens/rtc \
>   -H "X-User-Id: user_123" \
>   -H "Content-Type: application/json" \
>   -d '{"channel":"room_1","uid":"user_456","role":"host"}'

RTM Token
> curl -X POST http://localhost:8080/api/tokens/rtm \
>   -H "X-User-Id: user_123" \
>   -H "Content-Type: application/json" \
>   -d '{"uid":"user_456"}'

Create Room
> curl -X POST http://localhost:8080/api/rooms \
>   -H "X-User-Id: user_123" \
>   -H "Content-Type: application/json" \
>   -d '{"name":"Вебинар","is_private":true,"max_participants":50}'

Запуск через Docker:
> docker-compose up --build

Запуск через Kubernetes (Minikube):
> minikube start
> 
> eval $(minikube docker-env)
> 
> docker build -t agora-app:latest .
> 
> kubectl apply -f k8s/
> 
> minikube service agora-app -n agora (берем порт из tunnel и вставляем его в @host в request.http)

Для доступа к подключению к клиенту БД можно использовать эту команду
> kubectl port-forward svc/postgres 5432:5432 -n agora