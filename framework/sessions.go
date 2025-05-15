package framework

import (
	"fmt"

	ginSession "github.com/gin-contrib/sessions"
	hertzSession "github.com/hertz-contrib/sessions"
)

type Session struct {
	c            *Context
	ginSession   ginSession.Session   // Gin
	hertzSession hertzSession.Session // Hertz
}

// GetSession 获取 Session
func (c *Context) GetSession() (*Session, error) {
	if c.IsGin() {
		session := ginSession.Default(c.GetGinContext())
		if session == nil {
			return nil, fmt.Errorf("gin session is nil, maybe middleware not registered")
		}
		return &Session{
			ginSession: session,
		}, nil
	}
	if c.IsHertz() {
		session := hertzSession.Default(c.GetHertzContext())
		if session == nil {
			return nil, fmt.Errorf("hertz session is nil, maybe middleware not registered")
		}
		return &Session{
			hertzSession: session,
		}, nil
	}
	if c.IsEcho() {
	}
	return nil, fmt.Errorf("unsupported framework")
}

// 会话的ID
func (s *Session) ID() string {
	if s.ginSession != nil {
		return s.ginSession.ID()
	}
	if s.hertzSession != nil {
		return s.hertzSession.ID()
	}
	return ""
}

func (s *Session) Get(key any) any {
	if s.ginSession != nil {
		return s.ginSession.Get(key)
	}
	if s.hertzSession != nil {
		return s.hertzSession.Get(key)
	}
	return nil
}

func (s *Session) Set(key any, val any) {
	if s.ginSession != nil {
		s.ginSession.Set(key, val)
	}
	if s.hertzSession != nil {
		s.hertzSession.Set(key, val)
	}
	if s.c.IsEcho() {
	}
}

func (s *Session) Delete(key any) {
	if s.ginSession != nil {
		s.ginSession.Delete(key)
	}
	if s.hertzSession != nil {
		s.hertzSession.Delete(key)
	}
	if s.c.IsEcho() {
	}
}

func (s *Session) Clear() {
	if s.ginSession != nil {
		s.ginSession.Clear()
	}
	if s.hertzSession != nil {
		s.hertzSession.Clear()
	}
	if s.c.IsEcho() {
	}
}

func (s *Session) Save() error {
	if s.ginSession != nil {
		return s.ginSession.Save()
	}
	if s.hertzSession != nil {
		return s.hertzSession.Save()
	}
	return nil
}
