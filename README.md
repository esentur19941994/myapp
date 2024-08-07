Этот проект представляет собой сервис для управления пользователями с использованием MongoDB в качестве базы данных. Он предоставляет возможность регистрации новых пользователей, генерации JWT токенов для аутентификации и авторизации, а также доступ к информации о зарегистрированных пользователях через API endpoint.

Функциональность
Регистрация пользователя

Пользователь может зарегистрироваться, предоставив свое имя пользователя и дату регистрации.
Данные сохраняются в MongoDB.
Выдача JWT токена

После успешной регистрации пользователю выдается JWT токен.
Токен используется для аутентификации при последующих запросах к API.
Endpoint для получения списка пользователей

API предоставляет endpoint /users, который возвращает список всех зарегистрированных пользователей.
Аутентификация с использованием JWT

При запросе к API пользователь должен предоставить JWT токен.
Сервис проверяет валидность токена и разрешает доступ к данным пользователей только авторизованным пользователям.
