# Как запустить

1. docker build -t server .
2. docker run -p 5000:5000 server

Пока не добавил поддержку сваггера, имеется 3 эндпоинта:
* localhost/5000/header
* localhost/5000/footer
* localhost/5000/card

Пока каждый из них просто кидает эхо пост запроса.
Структуры можно посмотреть в папке schemas
