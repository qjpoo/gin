package session

import "fmt"

// 内存版 的session服务


// 根据key来获取value
func (s *SessionData) Get(key string) (value interface{}, err error) {
	// 获取读锁
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	value, ok := s.Data[key]
	if !ok {
		err = fmt.Errorf("Invalid Key ...")
		return
	}
	return
}

// set
func (s *SessionData) Set(key string, value interface{}) {
	// 获取写锁
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	s.Data[key] = value
}

// delete 删除key对应的值
func (s *SessionData) Del(key string)  {
	// 获取写锁
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	delete(s.Data, key)
}