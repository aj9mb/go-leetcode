package leetcode

type TimeMap struct {
	store map[string][]Value
}

func Constructor() TimeMap {
	return TimeMap{
		store: make(map[string][]Value),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	v := Value{
		valueStr:  value,
		timestamp: timestamp,
	}
	if values, ok := this.store[key]; ok {
		values = append(values, v)
		this.store[key] = values
	} else {
		values = []Value{v}
		this.store[key] = values
	}
}

func (this *TimeMap) Get(key string, timestamp int) string {
	res, v := "", this.store[key]
	if v != nil {
		l := 0
		r := len(v) - 1
		for l <= r {
			mid := (l + r) / 2
			if v[mid].timestamp <= timestamp {
				res = v[mid].valueStr
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return res
}

type Value struct {
	valueStr  string
	timestamp int
}
