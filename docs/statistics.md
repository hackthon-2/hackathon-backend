## 统计接口文档

#### 查看统计 /api/user/statistics

**请求方式GET**

|请求参数|是否必须|说明|  
|-------|------|---|  
|time|否|根据这个日期，算当天到这一天的统计数据，一天只能算一次，之后就不再更新|  

**请求成功示例：**

```json
{
  "status": 1,
  "code": 200,
  "message": "获取成功",
  "data": {
    "daskljdakslf": {
      "图书": 5,
      "图书馆": 4
    }
  }
}
```