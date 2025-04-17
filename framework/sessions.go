package framework

import (
	ginSession "github.com/gin-contrib/sessions"
	hertzSession "github.com/hertz-contrib/sessions"
)

type Session struct {
	ginSession   ginSession.Session   // Gin
	hertzSession hertzSession.Session // Hertz
}

// GetSession 获取 Session
func (c *Context) GetSession() Session {
	if c.IsGin() {
		session := ginSession.Default(c.ginCtx)
		return Session{
			ginSession: session,
		}
	}
	if c.IsHertz() {
		session := hertzSession.Default(c.hertzCtx)
		return Session{
			hertzSession: session,
		}
	}
	// 如果没有 Gin 或 Hertz 上下文返回一个空的 Session
	return Session{}
}

func (s Session) ID() string {
	if s.ginSession != nil {
		return s.ginSession.ID()
	}
	if s.hertzSession != nil {
		return s.hertzSession.ID()
	}
	return ""
}

func (s Session) Get(key any) any {
	if s.ginSession != nil {
		return s.ginSession.Get(key)
	}
	if s.hertzSession != nil {
		return s.hertzSession.Get(key)
	}
	return nil
}

func (s Session) Set(key any, val any) {
	if s.ginSession != nil {
		s.ginSession.Set(key, val)
	}
	if s.hertzSession != nil {
		s.hertzSession.Set(key, val)
	}
}

func (s Session) Delete(key any) {
	if s.ginSession != nil {
		s.ginSession.Delete(key)
	}
	if s.hertzSession != nil {
		s.hertzSession.Delete(key)
	}
}

func (s Session) Clear() {
	if s.ginSession != nil {
		s.ginSession.Clear()
	}
	if s.hertzSession != nil {
		s.hertzSession.Clear()
	}
}

func (s Session) Save() error {
	if s.ginSession != nil {
		return s.ginSession.Save()
	}
	if s.hertzSession != nil {
		return s.hertzSession.Save()
	}
	return nil
}
