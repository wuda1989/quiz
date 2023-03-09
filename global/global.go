package global

import (
	"gorm.io/gorm"
	"sync"
)

var (
	Global_DB     *gorm.DB
	Global_DBList map[string]*gorm.DB
	lock          sync.RWMutex
)
