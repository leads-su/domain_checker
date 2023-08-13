# Многопоточная проверка доменов/сайтов

Доступные проверки:
* наличие одного `ipv4` в `A` записи из предложенного списка
* наличие действительного сертификата у сайта

> Проверки собираются в список и для одного домена идут последовательно, если текущая проверка закончилась провалом то следующая не начнется.

## Сборка

```bash
$ make
```

## Использование

```bash
$ domain_checker [OPTIONS] checker_list path_of_file
```

Опции:
* `-a` - список `IP` адресов через запятую, которые разрешено указывать в `A` записи домена, например `"127.0.0.1,127.0.0.2"`
* `-n` - количество потоков для проверки `[1, 100]`, по умолчанию `5`
* `-dry-run` - сухой запуск, валидация параметров запуска без выполнения многопоточной проверки

Обязательные аргументы:
* `checker_list` - список чекеров для проверки разделенных запятой, доступны:
  * `ip` - проверка наличия разрешенного `IP` в `A` записи домена
  * `cert` - проверка валидности сертификата (`443` порт по умолчанию)
* `path_of_file` - путь до файла с доменами (без протокола, путей и портов) разделенных новой строкой

Пример запуска:
```bash
$ domain_checker -a "127.0.0.1,127.0.0.2" -n 10 "ip,cert" domains.txt
```

## Лицензия

MIT
