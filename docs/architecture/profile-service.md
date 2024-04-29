## Profile service

### Сценарии

Отвечает за редактирование пользователем своего профиля (никнейм - уникальный, информация о себе)

- Изменить никнейм
- Отредактировать информацию пользователя о себе

### API

- Изменить никнейм
  `PUT /:uuid/nickname`

  ```
  {
      nickname: string
  }
  ```

  Изменяет никнейм пользователя и отправляет сообщение с новым никнейм в `user_updated` топик.

- Редактирование информации о себе
  `PUT /:uuid/info`

  ```
  {
      firstName?: string
      lastName?: string
      description?: string
  }
  ```

- Получение информации о пользователе
  `GET /:uuid/info`
  Ответ сервера:

  ```
  {
      uuid: string
      email: string
      nickname: string
      firstName?: string
      lastName?: string
      description?: string
      avatar?: url
  }
  ```

- Поиск пользователя по никнейм или email
  `GET /info?email=?&nickname=?`
  Возвращает список пользователей подходящих под данные критерии
  ```
  {
     uuid: string
     email: string
     nickname: string
     firstName?: string
     lastName?: string
     description?: string
     avatar?: url
  }
  ```

### Хранение данных

Для хранения данных сервис использует реляционную базу данных (PostgreSQL).
Пользователи хранятся в таблице `users`:

```
user_id uuid PK
email varchar not null unique constraint
nickname varchar not null unique constraint
```

Информация о пользователе хранится в таблице `user_info`:

```
id int PK
firstName varchar nullable
lastName varchar nullable
description text nullable
avatar varchar nullable

```

### Интеграция с другими сервисами

- Подписан на `user_created` топик и создает профиль пользователя после регистрации
- Подписан на `avatar_uploaded` топик, для того чтобы обновить ссылку на аватар, после его загрузки
- При изменении nickname пользователя отправляетс информацию в `user_updated` топик
