## Friend service

### Сценарии

- Добавить пользователя в друзья
- Удалить пользователя из друзей
- Подтвердить или отклонить запрос на дружбу
- Просмотр списка своих друзей и запросов на дружбу

### API

- Поиск пользователя, для добавления в друзья (возвращает список совпадающих с search term пользователей)
  `GET /users?nickname=?&fname=?&lname=?`

- Добавить пользователя в друзья
  `POST /:uuid/friends/:uuid`
  сохранить запрос на добавление в друзья и отправить уведомление в notification topik которое
  должен обработать notification service

- Удалить пользователя из друзей
  `DELETE /:uuid/friends/:uuid`
  удалить пользователя из друзей и отправить уведомление в notification topik которое
  должен обработать notification service

- Подтвердить запрос на дружбу и отправить уведомление в notification topik которое
  должен обработать notification service
  `POST /:uuid/approve/:uuid`

- Отклонить запрос на дружбу и отправить уведомление в notification topik которое
  должен обработать notification service
  `POST /:uuid/decline/:uuid`

- Вывести список своих друзей и запросов на дружбу
  `GET /:uuuid/friends`

### Хранение данных

Для MVP можно хранить данные в MongoDB и допустить избыточность (мы будем сохранять связь в коллекции для обоих пользователей
которые находятся в отношении). В дальнейшем, когда сценарии будут более понятны, можно было бы подумать об использовании специальной графовой базы данных.

Пример коллекции:

```
{
    userId: "bd845fda-96d0-43c6-bbc5-85735b2bb2d1",
    friendship: [
       {
            userId: "5423e587-4376-444c-8ea3-bd37678d75af",
            name: "Bob",
            avatar: "Bob.jpg",
            isConfirmed: true,
       },
       {
            userId: "21d4c118-88fb-49f2-afbc-8eac1d0a7738",
            name: "Alice",
            avatar: "Alice.jpg",
            isConfirmed: false,
       }
    ]
},
{
    userId: "5423e587-4376-444c-8ea3-bd37678d75af",
    friendship: [
       {
            userId: "bd845fda-96d0-43c6-bbc5-85735b2bb2d1",
            name: "John",
            avatar: "John.jpg",
            isConfirmed: true,
       }
    ]
}
```

### Интеграция с другими сервисами

- Может слушать `user_created` и `user_updated` события для того что бы создать коллекцию для пользователя и, возможно, дублировать для него какую-то информацию для имени и аватара.
- отправляет сообщения в `notification` топик для уведомления о добавлении в друзья/удалении из друзей