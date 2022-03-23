# rentateam-test
Тестовое задание для собеседования в Renta Team

## Как запускать
Миграции
```zsh
go run main.go migrate
```

Сервер
```zsh
go run main.go run
```

## API эндпоинты
### HTTP
1. Создать новый пост: [POST] `http://localhost:3010/api/v1/posts`
<br><em>Пример тела запроса</em>:
```json
{
  "header": "First post in my blog",
  "body": "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet",
  "tags": ["interviews", "golang"]
}
```

2. Получить список постов: [GET] `http://localhost:3010/api/v1/posts`

### gRPC
1. Создать новый пост: `api.ApiService.CreatePost`
1. Получить список постов: `api.ApiService.ListPosts`

## Модули системы
| Название | Описание | Путь |
|---|---|---|
|Config|Модуль с Viper конфигом|modules/config|
|Logger|Модуль с логгером, основан на Logrus|modules/logger|
|Database|Модуль для работы с PostgreSQL|modules/database|
|GRPC|Модуль с gRPC сервером|modules/server/grpc|
|HTTP|Модуль с HTTP сервером|modules/server/http|
|Controllers|Контроллеры для GRPC сервера|modules/server/grpc/controllers|

## Использованные технологии
### Proto
1. [protoc-gen-gorm](https://github.com/TheSDTM/protoc-gen-gorm) - генерация gORM структур на основании прото-структур;
2. [protoc-gen-validate](https://github.com/envoyproxy/protoc-gen-validate) - генерация валидаторов на основании параметров полей, указанных в прото;
3. [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) - генерация gRPC сервера и клиента на основании прото структур, реализация http -> grpc прокси;
4. [uber/prototool](https://github.com/Minish144/prototool-arm64-support) - удобная обертка на protoc для генерации из прото файликов с конфигами в yaml формате (по ссылке мой форк с маленькими изменениями для поддержки arm64 архитектуры процессоров для работы на Маках с m1)

### Go
1. [uber-go/fx](https://github.com/uber-go/fx) - удобная библиотека для реализации dependency injection;
2. [logrus](github.com/sirupsen/logrus) - логирование;
3. [viper](github.com/spf13/viper) - конфиг системы (сам конфиг в `config.yaml`);
4. [gORM](https://github.com/go-gorm/gorm) - ORM для SQL БД;
5. [cobra](https://github.com/spf13/cobra) - создание cli

## Что можно было бы улучшить
1. Сильно хотелось показать, как можно быстро написать одновременно и grpc и http api, однако в этом способе есть минус - своего рода проксирование. В GRPC-Gateway создается http сервер, под капотом которого grpc клиент, то есть получается такой локальный запрос на grpc сервер. Отсюда следует бóльшая задержка и меньшая производительность;
2. Второй минус, вытекающий отсюда, статус коды после проксирования. Правильно на успешный post запрос возвращать **201 Created**, но по дефолту grpc-gateway мапит все grpc.OK на 200;
2. По коду - в идеале все покрыть тестами, в т.ч. мокать базу данных и тестировать методы обращения к ней;
3. Можно было бы подумать над уровнем изоляции в транзакции при создании постов с тегами, возможно, есть смысл настроить что-то помимо дефолтного `Read committed`;