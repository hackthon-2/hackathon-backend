## todo模块接口文档 ##  

#### 创建待办表 /api/user/todoCreation

**请求方式:POST**  
**请求数据格式：application/json**

|请求参数|是否必选|说明|
|-------|-------|---|
|header|是|问题|
|todoItems|是 |一个对象数组，包含所有的待办项|
|time|是|日期|  

***todoItems对象结构示例:***

```json
{
  "id": "faskjfkaf",
  "item": "图书馆",
  "isComplete": false
}
```  

**请求成功示例**

```json
{
  "status": 1,
  "code": 200,
  "message": "请求成功",
  "data": []
}
```  

#### 更新待办表 /api/user/todoUpdate

**请求方式：POST**  
**请求数据格式：application/json**

|请求参数|是否必选|说明|
|-------|-------|----|
| id | 是 | 待办的id |
| header | 是 | 问题 |
| todoItems | 是 | 一个对象数组，包含所有的待办项,示例如上 |
| time | 是 | 日期 |  

**请求成功示例**

```json
{
  "status": 1,
  "code": 200,
  "message": "请求成功",
  "data": []
}
```  

#### 获取待办表 /api/user/todoList

**请求方式：GET**

|请求参数|是否必选|说明|
|-------|-------|----|
| time | 是 | 日期，格式YYYY-MM-DD |  

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

#### 删除待办表 /api/user/todoDeletion

**请求方式：DELETE**

| 请求参数 | 是否必选 | 说明 |
| ------- |------- | ---- |
| todoID | 是 | 待办的id |

**请求成功示例**

```json
{
  "status": 1,
  "code": 200,
  "message": "请求成功",
  "data": []
}
```