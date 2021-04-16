+## todo模块接口文档 ##

#### 创建计划/api/user/todoCreation

**请求方式:POST**\
**请求数据格式：x-www-form-urlencoded**

|请求参数|是否必选|说明|
|-------|-------|---|
|question|是|问题|
|text|是 |文本|
|time|是|日期|

**请求成功示例**

```json
{
  "status": 1,
  "code": 200,
  "message": "请求成功",
  "data": []
}
```

#### 更新计划/api/user/todoUpdate

**请求方式：POST**\
**请求数据格式：x-www-form-urlencoded**

|请求参数|是否必选|说明|
|-------|-------|----|
|id|是|计划的ID号|
|question|是|问题|
|test|是|文本|
|time|是|日期|

**请求成功示例**

```json
{
  "status": 1,
    "code": 200,
    "message": "请求成功",
    "data": []
}
```

#### 计划表/api/user/todoList

**请求方式：GET**\
**请求数据格式：x-www-form-urlencoded**

|请求参数|是否必选|说明|
|-------|-------|----|
|id|是|计划的ID号|
|time|是|日期|

**请求成功示例**

```json
{
     "status": 1,
     "code": 200,
     "message": "获取成功",
     "data": [
         {
             "id": 2,
             "user_id": 1,
             "header": "今天干了尼玛？",
             "todoItems": [
                 {
                     "id": "faskjfkaf",
                     "item": "图书馆",
                     "isComplete": false
                 },
                 {
                     "id": "faskjfkafgdas",
                     "item": "图书",
                     "isComplete": true
                 },
                 {
                     "id": "ffsaf",
                     "item": "书",
                     "isComplete": false
                 }
             ],
             "time": "2021-04-15"
         },
         {
             "id": 3,
             "user_id": 1,
             "header": "今天干了饭？",
             "todoItems": [
                 {
                     "id": "faskjfkaf",
                     "item": "起飞了奥",
                     "isComplete": false
                 },
                 {
                     "id": "faskjfkafgdas",
                     "item": "图书",
                     "isComplete": true
                 },
                 {
                     "id": "ffsaf",
                     "item": "书",
                     "isComplete": false
                 }
             ],
             "time": "2021-04-15"
         },
         {
             "id": 4,
             "user_id": 1,
             "header": "飞天吗？",
             "todoItems": [
                 {
                     "id": "faskjfkaf",
                     "item": "起飞了奥",
                     "isComplete": false
                 },
                 {
                     "id": "faskjfkafgdas",
                     "item": "图书",
                     "isComplete": true
                 },
                 {
                     "id": "ffsaf",
                     "item": "书",
                     "isComplete": false
                 }
             ],
             "time": "2021-04-15"
         },
         {
             "id": 5,
             "user_id": 1,
             "header": "淦碎了吗？",
             "todoItems": [
                 {
                     "id": "faskjfkaf",
                     "item": "起飞了奥",
                     "isComplete": false
                 },
                 {
                     "id": "faskjfkafgdas",
                     "item": "图书",
                     "isComplete": true
                 },
                 {
                     "id": "ffsaf",
                     "item": "书",
                     "isComplete": false
                 }
             ],
             "time": "2021-04-15"
         },
         {
             "id": 6,
             "user_id": 1,
             "header": "他喵的个？",
             "todoItems": [
                 {
                     "id": "faskjfkaf",
                     "item": "起飞了奥",
                     "isComplete": false
                 },
                 {
                     "id": "faskjfkafgdas",
                     "item": "图书",
                     "isComplete": true
                 },
                 {
                     "id": "ffsaf",
                     "item": "书",
                     "isComplete": false
                 }
             ],
             "time": "2021-04-15"
         }
     ]
 }
```

#### 删除计划/api/user/todoDeletion

**请求方式：DELETE**\
**请求数据格式：x-www-form-urlencoded**

|请求参数|是否必选|说明|
|-------|-------|----|
|id|是|计划的ID号|

**请求成功示例**

```json
{
  "status": 1,
    "code": 200,
    "message": "请求成功",
    "data": []
}
```