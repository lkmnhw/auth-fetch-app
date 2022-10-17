# auth-app
A simple app for jwt authentication

## available endpoints

#### register

```http
  POST /register
```
| Parameter | Type     | Description  |
| :-------- | :------- | :----------- |
| `phone`   | `string` | **Required** |
| `name`    | `string` | **Required** |
| `role`    | `string` | **Required** |
| `password`| `string` | **Optional**. automatically generated if empty |

#### login

```http
  POST /login
```

| Parameter | Type     | Description  |
| :-------- | :------- | :----------- |
| `phone`   | `string` | **Required** |
| `password`| `string` | **Required** |

#### info

```http
  GET /info
```