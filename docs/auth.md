## Auth模块接口文档

#### 注册接口 /api/auth/registration

**请求方式:POST**  
**请求数据格式:x-www-form-urlencoded**

|请求参数 |是否必选 |说明 |
|  ----  | ----  |----|
| username | 是 | 用户名 |
| password  | 是 | 密码 |
| email | 是 | 邮箱 |
| sex | 是 | 性别 |  

**请求成功示例**

```json
{
  "status": 1,
  "code": 200,
  "message": "请求成功",
  "data": []
}
```

#### 登陆接口 /api/auth/login

**请求方式:POST**  
**请求数据格式:x-www-form-urlencoded**

|请求参数 |是否必选 |说明 |
|  ----  | ----  |----|
| username | 是 | 用户名 |
| password  | 是 | 密码 |  

**请求成功示例**

```json
{
  "status": 1,
  "code": 200,
  "message": "请求成功",
  "data": {
    "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6NCwiVXNlcm5hbWUiOiJhZG1pbjIzNCIsImV4cCI6MTYxODIwOTg3NSwiaWF0IjoxNjE4MjAyNjc1LCJpc3MiOiJ4emgiLCJuYmYiOjE2MTgyMDE2NzV9.CUkbMIt0tfuR1c1y5g-d5limys48lhoxKjGSveYbgSe2uajVOuefjjmZee2zxB2gdrSy7OUz7FMNPgDj6vQ0yXGqe-_-5Q5Z0WWf8YMmp7rHh0XcOBwFSYFPFjdohpEKr8m1ltgliLfeUw0xBVNfoA2NavXNjKCLXcSgPw1xz8c"
  }
}
```