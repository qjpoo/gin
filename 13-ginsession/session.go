package ginsession

// 自己实现的gin框架session中间件

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid2 "github.com/satori/go.uuid"
	"sync"
)

const (
	SessionCookieName  = "session_id" // session id 存储在cookie里面的key
	SessionContextName = "session"    // session data在Context上下文中的key
)

var (
	MgrObj *Mgr // 定主一个全局的session管理对象
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
		ID:   id,
		Data: make(map[string]interface{}, 8),
	}

}

// mgr是一个全局的session管理
type Mgr struct {
	Session map[string]*SessionData
	rwLock  sync.RWMutex
}

func InitMgr() {
	MgrObj = &Mgr{
		Session: make(map[string]*SessionData, 1024), // 先初始化1024个红色的小框, 用来存储用户的session数据
	}
}

// 对session的操作

//SessionData支持的操作

// 根据传进来的的session id,找到对就在的session记录
func (m *Mgr) GetSessionData(sessionID string) (sd *SessionData, err error) {
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
func (m *Mgr) createSession() (sd *SessionData) {
	id, err := uuid2.NewV4()
	if err != nil {
		return
	}
	// 造一个和它对就在的sessiondata
	sd = NewSessionData(id.String())
	m.Session[sd.ID] = sd  // 保存到大仓库中
	// 返回sessiondata
	return
}

//  实现一个gin框架的中间件
/*
	r.Use(sessionMiddleware)
	中间件是一个 HandlerFunc类型 : func (engine *Engine) Use(middleware ...HandlerFunc) IRoutes
*/

// 所有流经我这个中间件的请求, 它的上下文肯定会有一个session -->session data
func SessionMiddleware(mgrObj *Mgr) gin.HandlerFunc {
	if mgrObj == nil {
		panic("mgrobj没有初始化")
	}
	return func(c *gin.Context) {
		// 1. 从请求的Cookie中获取session_id
		var sd *SessionData  // 申明了一个sd  sessiondata指针对象
		var sessionID string // 申明全局的sessionID
		sessionID, err := c.Cookie(SessionCookieName)
		if err != nil {
			// 1.1 取不到session id, 给这个用户创建一个新的session id,同时分配一个session_id
			sd = mgrObj.createSession()
			sessionID = sd.ID
		} else {
			// 1.2 取到cookie里面的session_id
			// 2. 根据session id去session大仓库里面查找,对应的session data数据
			sd, err = mgrObj.GetSessionData(sessionID)
			if err != nil {
				// 2.1 根据用户传过来的seesion id在大仓库中查找, 查找不到就创建session data(有可能这个sessionid是人造的假的,或者是过期的)
				sd = mgrObj.createSession()
				// 2.2 更新用户cookie中保存的那个session id
				sessionID = sd.ID
			}
		}

		// 3. 如何实现后续所有的处理请求都能拿到session data
		// 3.1 利用gin中的c.Set("session", "session data")
		c.Set(SessionContextName, sd)
		// 在gin框架中,要回写Cookie必须要在处理请求的函数之前
		c.SetCookie(SessionCookieName, sessionID, 20, "/", "localhost", false, true)
		c.Next() // 执行后续的handler
	}
}
