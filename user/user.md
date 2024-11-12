### 1. hello

1. route definition

- Url: /hello/:name
- Method: GET
- Request: `Request`
- Response: `Response`

2. request definition



```golang
type Request struct {
	Name string `path:"name,options=you|me"`
}
```


3. response definition



```golang
type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
}
```

### 2. add user

1. route definition

- Url: /user/add
- Method: POST
- Request: `User`
- Response: `Response`

2. request definition



```golang
type User struct {
	Name string `json:"name"`
	Email string `json:"email,optional"`
	Age int `json:"age"`
	Birthday string `json:"birthday,optional"`
}
```


3. response definition



```golang
type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
}
```

### 3. delete user

1. route definition

- Url: /user/delete/:id
- Method: DELETE
- Request: `DelUser`
- Response: `Response`

2. request definition



```golang
type DelUser struct {
	Id int `path:"id"`
}
```


3. response definition



```golang
type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
}
```

### 4. list user

1. route definition

- Url: /user/list
- Method: POST
- Request: `ListUser`
- Response: `ListUserResponse`

2. request definition



```golang
type ListUser struct {
	Name string `json:"name,optional"`
	Email string `json:"email,optional"`
}
```


3. response definition



```golang
type ListUserResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data []*UpdateUser `json:"data"`
}
```

### 5. update user

1. route definition

- Url: /user/update/:id
- Method: PUT
- Request: `UpdateUser`
- Response: `Response`

2. request definition



```golang
type UpdateUser struct {
	Id int `path:"id"`
	Name string `json:"name"`
	Email string `json:"email,optional"`
	Age int `json:"age"`
	Birthday string `json:"birthday,optional"`
}
```


3. response definition



```golang
type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
}
```

