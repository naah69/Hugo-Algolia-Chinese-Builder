package po

import "sync"

type ConcurrentMap struct {
	lock sync.RWMutex
	data map[string]interface{}
}

func NewConcurrentMap(data map[string]interface{}) *ConcurrentMap {
	return &ConcurrentMap{
		lock: sync.RWMutex{},
		data: data,
	}
}

func (m *ConcurrentMap) AddData(key string, value interface{}) {
	m.lock.Lock()
	m.data[key] = value
	m.lock.Unlock()
}

func (m *ConcurrentMap) GetValue(key string) interface{} {
	m.lock.RLock()
	var temp interface{} = m.data[key]
	m.lock.RUnlock()
	return temp
}

func (m *ConcurrentMap) GetData() map[string]interface{} {
	return m.data
}
