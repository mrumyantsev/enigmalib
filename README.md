# Шифровальные машины

**Шифровальные машины** - это веб-приложение, предназначенное для перевода текста в шифр, в точности реализующее механизм шифровальных машин фирмы "Enigma", выпускавшихся в Германии с начала 20-х годов XX века, а также поздних измененных моделей, предназначавшихся для работы госструктур в других странах после Второй мировой войны.

![Шифровальные машины](./cipher-machines-web-ui.png "Веб-интерфейс приложения \"Шифровальные машины\"")\
*Вид приложения в браузере*

## Пакеты для импортирования

* `pkg/machine` - вспомогательные типы для инициализации любых шифровальных машин и для работы веб-сервера данного проекта
* `pkg/machine/enigma` - конкретные моделями машин Enigma
* `pkg/machine/enigma/base` - базовые детали для построения новой машины с механизмом Enigma

## Системные требования

**Операционная система:**

- Windows / MacOS / Linux.

**Программное обеспечение:**

- средства разработки языка Go >=v1.20 (только для локальной сборки);
- платформа контейнеризации Docker;
- утилита make.

## Установка и запуск

Для **сборки** приложения в **Docker** выполните эту команду:

#### Для Linux:

```
docker compose build
```

#### Для MacOS:

```
docker-compose build
```

Для **запуска** приложения в контейнерах **Docker** выполните следующую команду (она также запустит сборку, если таковой не было произведено):

#### Для Linux:

```
docker compose up
```

#### Для MacOS:

```
docker-compose up
```

Чтобы воспользоваться **веб-интерфейсом** приложения откройте интернет-браузер и перейдите в нем по адресу:

```
http://localhost:8090/
```

Для того, чтобы **остановить** запущенные **Docker**-контейнеры и конфигурации текущего развертывания, выполните команду:

#### Для Linux:

```
docker compose down
```

#### Для MacOS:

```
docker-compose down
```

Для **сборки серверного компонента** приложения на **локальном** компьютере выполните команду:

```
make build
```

Чтобы **запустить серверный компонент** приложения на **локальном** компьютере выполните следующую команду:

```
make run
```

Чтобы **запустить демонстрацию**, выполните команду:

```
make demo-run
```

Или выполните ее короткую версию:

```
make
```

## Траблшутинг

Если при развертывании в Docker постоянно появляется ошибка *"This port already in use"* попробуйте поменять этот порт, о котором говорится в ошибке, с помощью того же файла с параметрами `.env`.

Если при открытии веб-интерфейса приложения в браузере появляется страница с ошибкой *"Could not resolve host: localhost"*, замените в поле адреса слово `localhost` на `127.0.0.1` и перейдите по данному адресу.
