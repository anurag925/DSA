package dsa

type dataTime struct {
	data string
	ts   int
}

type TimeMap struct {
	store map[string][]dataTime
}

func Constructor() TimeMap {
	return TimeMap{
		store: make(map[string][]dataTime),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.store[key]; !ok {
		this.store[key] = make([]dataTime, 0)
	}
	this.store[key] = append(this.store[key], dataTime{value, timestamp})
}

func (this *TimeMap) getHelper(arr []dataTime, t int) dataTime {
	res := dataTime{ts: -1}
	start, end := 0, len(arr)-1
	for start <= end {
		mid := (start + end) / 2
		// fmt.Println(mid)
		i := arr[mid]
		if i.ts == t {
			return i
		} else if i.ts < t {
			res = i
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return res
}

func (this *TimeMap) Get(key string, timestamp int) string {
	// fmt.Println(this.store)
	res := this.getHelper(this.store[key], timestamp)
	return res.data
}
