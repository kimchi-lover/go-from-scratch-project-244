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

Поддерживаемые форматы: JSON (`.json`) и YAML (`.yml`, `.yaml`). Формат определяется по расширению файла, файлы могут быть в разных форматах. Сравниваются структуры любой вложенности.

```bash
make build
./bin/gendiff testdata/fixture/file1.json testdata/fixture/file2.json
./bin/gendiff --format stylish testdata/fixture/file1.yml testdata/fixture/file2.yaml
```

Флаг `--format` (`-f`) задаёт формат вывода: `stylish` (по умолчанию), `plain` или `json`.

### stylish

Дерево изменений: `+` — ключ добавлен, `-` — удалён, без знака — не изменился. Изменённое значение выводится двумя строками: старое с `-`, новое с `+`.

[![Пример работы gendiff в формате stylish](https://asciinema.org/a/G91DWL7bSUMHKV3Y.svg)](https://asciinema.org/a/G91DWL7bSUMHKV3Y)

### plain

Список изменений в виде текста, по строке на каждое изменённое свойство. Для вложенных свойств выводится полный путь от корня, составные значения заменяются на `[complex value]`, строки выводятся в одинарных кавычках. Неизменившиеся свойства не выводятся.

```bash
./bin/gendiff --format plain testdata/fixture/file1.json testdata/fixture/file2.json
```

```
Property 'common.follow' was added with value: false
Property 'common.setting2' was removed
Property 'common.setting3' was updated. From true to null
Property 'common.setting4' was added with value: 'blah blah'
Property 'common.setting5' was added with value: [complex value]
Property 'common.setting6.doge.wow' was updated. From '' to 'so much'
Property 'common.setting6.ops' was added with value: 'vops'
Property 'group1.baz' was updated. From 'bas' to 'bars'
Property 'group1.nest' was updated. From [complex value] to 'str'
Property 'group2' was removed
Property 'group3' was added with value: [complex value]
```

<!-- asciinema: plain -->

### json

Дерево изменений в формате JSON — для обработки другими программами. Корень — объект с полем `diff`: массивом узлов, отсортированных по ключу. У каждого узла есть поля `key`, `type`, `oldValue` и `newValue`, у вложенных — ещё `children`. Значимые поля определяются типом узла:

| `type`      | Значимые поля          | Описание                                         |
|-------------|------------------------|--------------------------------------------------|
| `added`     | `newValue`             | ключ добавлен                                    |
| `removed`   | `oldValue`             | ключ удалён                                      |
| `unchanged` | `oldValue`, `newValue` | значение не изменилось                           |
| `changed`   | `oldValue`, `newValue` | значение изменилось                              |
| `nested`    | `children`             | оба значения — объекты, `children` — их сравнение |

```bash
./bin/gendiff --format json testdata/fixture/file1.json testdata/fixture/file2.json
```

Фрагмент вывода — узел `group1`:

```json
{
  "key": "group1",
  "type": "nested",
  "oldValue": null,
  "newValue": null,
  "children": [
    {
      "key": "baz",
      "type": "changed",
      "oldValue": "bas",
      "newValue": "bars"
    },
    {
      "key": "foo",
      "type": "unchanged",
      "oldValue": "bar",
      "newValue": "bar"
    },
    {
      "key": "nest",
      "type": "changed",
      "oldValue": {
        "key": "value"
      },
      "newValue": "str"
    }
  ]
}
```

<!-- asciinema: json -->

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
