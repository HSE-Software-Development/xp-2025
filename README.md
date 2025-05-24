# xp-2025

## Запуск центрального сервера
Поднимает docker контейнер на 9092 порту
```bash
cd server
docker-compose up -d
```

## Запуск клиента:
```
cd client/frontend/xp-chat
make run

cd ../../backend
make run

```
Приложение будет доступно на 5173 порту локального host

## Документы
- [Техническое задание](docs/KR.txt)
- [Архитектура приложения](docs/architecture.md)
- [Архитектурная схема](docs/image.png)

## Видео работы приложения
[Тут](videos)