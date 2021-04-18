package service

import (
	"encoding/json"
	"errors"
	database "hackthon/database/mysql"
	database2 "hackthon/database/redis"
	"hackthon/model"
	"sort"
	"strings"
)

var GenerateStatisticsError = errors.New("generating statistics error")

func Statistics(userId uint, from string) ([]model.Statistics, error) {
	dat, err := database2.FindStatisticsCache(userId, from)
	if err != nil || dat == "" {
		headers, e := database.ListHeaderFromTime(userId, from)
		var data = make([]model.Statistics, len(headers))
		if e != nil {
			return nil, GenerateStatisticsError
		}
		//把这些header下面的待办项全部反序列化
		for in, v := range headers {
			val, er := database.ListItemFromTime(v, userId, from)
			if er != nil {
				return nil, GenerateStatisticsError
			}
			items := make([]model.ToDoItem, 0, 600)
			for _, value := range val {
				itemVal := strings.Split(value, "/")
				for _, va := range itemVal {
					var item model.ToDoItem
					_ = json.Unmarshal([]byte(va), &item)
					items = append(items, item)
				}
			}
			//把非完成项全部删掉
			for i := 0; i < len(items); {
				if !items[i].IsComplete {
					items = append(items[:i], items[i+1:]...)
				} else {
					i++
				}
			}
			if len(items) < 1 {
				data[in] = model.Statistics{
					Header: v,
				}
				continue
			}
			//根据item进行排序
			sort.Slice(items, func(i, j int) bool {
				return items[i].Item < items[j].Item
			})
			var d = make(map[string]int64)
			var temp = items[0].Item
			var sum int64 = 1
			for i := 0; i < len(items); i++ {
				//如果只有一个元素，那直接赋值
				if len(items) == 1 {
					d[temp] = 1
				}
				//到最后一项时，如果还是相等则不会赋值，不相等也会导致最后一个元素没有赋值
				if i+1 < len(items) {
					//如果下一个元素和当前元素相等，就+1
					if items[i+1].Item == temp {
						sum++
					} else {
						//不然的话就已经统计完所有的该元素了,赋值后重置
						d[temp] = sum
						sum = 1
						temp = items[i+1].Item
					}
				}
			}
			//所以在此处进行最后一次赋值
			d[temp] = sum
			var da = make([]model.Item, len(d))
			var s int64 = 0
			for k, va := range d {
				da[s].Item = k
				da[s].Times = uint(va)
				s++
			}
			data[in].Header = v
			data[in].Items = da
		}
		val, _ := json.Marshal(&data)
		err = database2.CreateStatisticsCache(userId, from, string(val))
		return data, err
	}
	var data []model.Statistics
	err = json.Unmarshal([]byte(dat), &data)
	return data, err
}

