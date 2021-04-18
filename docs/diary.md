## diary模块接口文档

#### 创建日记 /api/user/diaryCreation

**请求方式POST**  
**请求数据格式:x-www-form-urlencoded**

|请求参数 |是否必选 |说明 |
|  ----  | ----  |----|
| question | 是 | 问题 |
| text  | 是 | 文本 |
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

#### 更新日记 /api/user/diaryUpdate

**请求方式POST**  
**请求数据格式:x-www-form-urlencoded**

|请求参数 |是否必选 |说明 |
|  ----  | ----  |----|
| diary_id | 是 | 日记的id号 |
| question | 是 | 问题 |
| text  | 是 | 文本 |
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

#### 获取日记 /api/user/diaryList

**请求方式GET**

|请求参数 |是否必选 |说明 |
|  ----  | ----  |----|
| time | 否 | 如果不输入特定值，默认为当天的数据 |  

**请求成功示例**

```json
{
  "status": 1,
  "code": 200,
  "message": "获取成功",
  "data": [
    {
      "id": 5,
      "user_id": 1,
      "question": "今天干了些啥？",
      "text": "睡觉睡觉睡觉写代码",
      "time": "2021-04-16"
    }
  ]
}
```  

#### 删除日记 /api/user/diaryDeletion

**请求方式DELETE**

|请求参数 |是否必选 |说明 |
|  ----  | ----  |----|
| diaryID | 是 | 日记的id |  

**请求成功示例**

```json
{
  "status": 1,
  "code": 200,
  "message": "请求成功",
  "data": []
}
```  

