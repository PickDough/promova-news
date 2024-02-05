## Запуск api в Docker:
```
    docker compose up
```

## Makefile команди:
```
    $ make build
    $ make run
    $ make clean
    $ make test
```

### [OpenApi](./src/docs/swagger.json)

---
Для продакшн коду, варто було би покрити тестами всі ділянки коду, але оскільки в цьому api бізнес логіка не є комплексною, то я не став дублювати тести. З їх прикладами можна ознайомитись у наступних файлах:

* [repository_test.go](./src/persistance/postsRepository/repository_test.go)
* [handler_test.go](./src/app/query/getPost/handler_test.go)
* [getPost_test.go](./src/network/postsApi/getPostApi/getPost_test.go)