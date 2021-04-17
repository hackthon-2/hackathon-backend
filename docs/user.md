## 用户相关接口文档

#### 上传头像 /api/user/avatar

**请求方式:POST**  
**请求数据格式:FormData**

|请求参数|是否必须|说明|  
|-------|------|---|  
| document | 是 | 图片文件 |  

***请求成功示例***

```json
{
  "status": 1,
  "code": 200,
  "message": "请求成功",
  "data": []
}
```  

#### 获取个人资料 /api/user/profile

**请求方式:GET**  
***无请求参数***

***请求成功示例***

```json
{
  "status": 1,
  "code": 200,
  "message": "获取成功",
  "data": {
    "id": 1,
    "username": "admin",
    "password": "",
    "email": "123456@email.com",
    "avatar": "https://oss.onesnowwarrior.cn/avatars/55c212c0c4e67b0cc7c69db6414bc9ae1d9135c3aea12ed9d319d141c7fe7700.jpg"
  }
}
```  

#### 更新个人资料 /api/user/profileUpdate

**请求方式：POST**  
**请求数据格式:x-www-form-urlencoded**

|请求参数|是否必须|说明|  
|-------|------|---|  
| username | 是 | 用户名 |  
| email | 是 | 电子邮箱 |  

***请求成功示例***

```json
{
  "status": 1,
  "code": 200,
  "message": "请求成功",
  "data": []
}
```  






