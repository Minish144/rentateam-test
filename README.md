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

## Использованные технологии
### Proto
1. [protoc-gen-gorm](https://github.com/TheSDTM/protoc-gen-gorm) - генерация gORM структур на основании прото-структур;
2. [protoc-gen-validate](https://github.com/envoyproxy/protoc-gen-validate) - генерация валидаторов на основании параметров полей, указанных в прото;
3. [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) - генерация gRPC сервера и клиента на основании прото структур, реализация http -> grpc прокси;
4. [uber/prototool](https://github.com/uber/prototool) - удобная обертка на protoc для генерации из прото файликов с конфигами в yaml формате

### Go
1. [uber-go/fx](https://github.com/uber-go/fx) - удобная библиотека для реализации dependency injection;
2. [logrus](github.com/sirupsen/logrus) - логирование;
3. [viper](github.com/spf13/viper) - конфиг системы (сам конфиг в `config.yaml`);
4. [gORM](https://github.com/go-gorm/gorm) - ORM для SQL БД;
5. [cobra](https://github.com/spf13/cobra) - создание cli

## Что можно было бы улучшить
<!-- TODO -->