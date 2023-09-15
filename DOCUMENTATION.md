## hng stage 2

### setting up the project

The project setup can be found in the [README.md](README.md) file

### requests & responses

### GET /api/:id


```bash
GET http://localhost:8080/api/oluwajuwon
```

```json
{
    "id": 1,
    "name": "oluwajuwon",
}
```

### POST /api/:id

expected request body:

```json
{
    "name": "oluwajuwon",
}
```

expected response body:

```json
{
    "id": 1,
    "name": "oluwajuwon",
}
```

### PUT /api/:id

expected request body:

```json
{
    "name": "Tomiwa",
}
```

expected response body:

```json
{
    "message": "Person updated",
}
```

### DELETE /api/:id

expected response body:

```json
{
    "message": "Person deleted",
}
```


