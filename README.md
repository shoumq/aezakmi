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

Запуск:
> docker-compose up --build
