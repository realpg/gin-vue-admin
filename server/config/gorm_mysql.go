package config

import "strings"

type Mysql struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (m *Mysql) Dsn() string {
	if strings.HasPrefix(m.Path, "unix:") {
		//兼容某些newbie不注意 配置的unix:和socket path之间多了空格 或者尾部有多余空格
		return m.Username + ":" + m.Password + "@unix(" + strings.TrimSpace(strings.TrimPrefix(m.Path, "unix:")) + ")/" + m.Dbname + "?" + m.Config
	}
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

func (m *Mysql) GetLogMode() string {
	return m.LogMode
}
