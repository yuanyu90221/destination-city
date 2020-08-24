# destination-city

## 題目解讀：

### 題目來源:
[destination-city](https://leetcode.com/problems/destination-city/)

### 原文:
You are given the array paths, where paths[i] = [cityAi, cityBi] means there exists a direct path going from cityAi to cityBi. Return the destination city, that is, the city without any path outgoing to another city.

It is guaranteed that the graph of paths forms a line without any loop, therefore, there will be exactly one destination city.

### 解讀：

給定一個路徑陣列 paths 每個 元素 path[i] = [cityAi, cityBi]

代表 cityAi到 cityBi有一條直達的路

其中保證所給的路徑一定是一條直路 只有一個終點

要求出最後的終點是哪個城市


## 初步解法:
### 初步觀察:

如果假設 cityZ是 最後的終點 

cityZ 一定會出現在 每個元素的第二個位置

並且部會出現在 其他元素的第一個位置

因此把建立一個 map 把第一個元素的收集起來

每次loop檢查第二個元素是否有出現在第一個元素

如果有代表不是最後終點

### 初步設計:
Given an path array paths

Step 0: let startMap = make(map[string]bool), idx = 0

Step 1: if idx >= len(paths) go to step 4

Step 2: startMap[path[idx][0]] = true

Step 3: idx = idx + 1 go to step 1

Step 4: idx = 0

Step 5: if idx >= len(paths) go to step 8

Step 6: if startMap[path[idx][1]] == false return path[idx][1]

Step 7: idx = idx + 1

Step 8: return "no found"

## 遇到的困難
### 題目上理解的問題
因為英文不是筆者母語

所以在題意解讀上 容易被英文用詞解讀給搞模糊

### pseudo code撰寫

一開始不習慣把pseudo code寫下來

因此 不太容易把自己的code做解析

### golang table driven test不熟
對於table driven test還不太熟析

所以對於寫test還是耗費不少時間
## 我的github source code
```golang
package dest_city

func destCity(paths [][]string) string {
	result := ""
	startMap := make(map[string]bool)
	for _, path := range paths {
		startMap[path[0]] = true
	}
	for _, path := range paths {
		if ok, _ := startMap[path[1]]; ok == false {
			return path[1]
		}
	}
	return result
}

```
## 測資的撰寫
```golang
package dest_city

import "testing"

func Test_destCity(t *testing.T) {
	type args struct {
		paths [][]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example1",
			args: args{
				paths: [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}},
			},
			want: "Sao Paulo",
		},
		{
			name: "Example2",
			args: args{
				paths: [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}},
			},
			want: "A",
		},
		{
			name: "Example3",
			args: args{
				paths: [][]string{{"A", "Z"}},
			},
			want: "Z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := destCity(tt.args.paths); got != tt.want {
				t.Errorf("destCity() = %v, want %v", got, tt.want)
			}
		})
	}
}

```

## my self record
[golang leetcode 30 day 25th day](https://hackmd.io/j4XGFNv-S4OyxkIFKXpJSQ?view)
## 參考文章

[golang test](https://ithelp.ithome.com.tw/articles/10204692)