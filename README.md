# Вычислитель отличий на Go

[![hexlet-check](https://github.com/kimchi-lover/go-from-scratch-project-244/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/kimchi-lover/go-from-scratch-project-244/actions)
[![CI](https://github.com/kimchi-lover/go-from-scratch-project-244/actions/workflows/ci.yml/badge.svg)](https://github.com/kimchi-lover/go-from-scratch-project-244/actions/workflows/ci.yml)

Консольная утилита для сравнения вложенных структур (JSON, YAML)

Учебный проект Хекслета: https://ru.hexlet.io/programs/go-from-scratch


## Стек

- Go

## Установка

<!-- Опишите установку: клонирование, зависимости, переменные окружения -->

```bash
git clone https://github.com/kimchi-lover/go-from-scratch-project-244.git
cd go-from-scratch-project-244
```

## Использование

Поддерживаемые форматы: JSON (`.json`) и YAML (`.yml`, `.yaml`). Формат определяется по расширению файла.

```bash
go run ./cmd/gendiff testdata/fixture/file1.json testdata/fixture/file2.json
go run ./cmd/gendiff testdata/fixture/file1.yml testdata/fixture/file2.yaml
```

Запись сравнения двух JSON-файлов:

[![Пример работы gendiff](https://asciinema.org/a/pIKCYanzhh8F431w.svg)](https://asciinema.org/a/pIKCYanzhh8F431w)

## Разработка

```bash
make lint           # линтер golangci-lint
make test           # тесты
make test-coverage  # тесты с проверкой порога покрытия (COVERAGE_MIN в Makefile)
```

---

<details>
<summary>Автоматические тесты Хекслета</summary>

Тесты запускаются на каждый коммит. За запуск отвечает файл `.github/workflows/hexlet-check.yml` — не удаляйте и не переименовывайте ни его, ни репозиторий.

</details>

## О Хекслете

[Хекслет](https://ru.hexlet.io/) — школа программирования: авторские программы обучения с практикой, поддержкой наставников и реальными проектами, которые остаются в резюме. Этот репозиторий — один из таких проектов.
