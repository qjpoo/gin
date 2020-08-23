package session

import (
	"fmt"
	uuid2 "github.com/satori/go.uuid"
	"sync"
)

// session服务

// sessionData表示一个具体的用户session数据
type SessionData struct {
	ID     string
	Data   map[string]interface{}
	rwLock sync.RWMutex // 读写锁, 锁的是data
	// Expire 过期时间
}

// 构造函数
func NewSessionData(id string) *SessionData {
	return &SessionData{
		ID: id,
		Data: make(map[string]interface{}, 8),
	}

}


// mgr是一个全局的session管理
type Mgr struct {

	Session map[string]SessionData
	rwLock sync.RWMutex
}

// 对session的操作

//SessionData支持的操作




// 根据传进来的的session id,找到对就在的session记录
func (m *Mgr) GetSessionData(sessionID string) (sd SessionData, err error) {
	// 取之前加锁
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()
	// 根据传进来的sessionID来获取sessiondata
	sd, ok := m.Session[sessionID]
	if !ok {
		err = fmt.Errorf("invalid session id")
		return
	}
	return
}

// 创建一条session id的记录
func (m *Mgr)createSession() (sd *SessionData, err error)  {
	id, err :=  uuid2.NewV4()
	if err != nil {
		return
	}
	// 造一个和它对就在的sessiondata
	sd = NewSessionData(id.String())
	// 返回sessiondata
	return

}