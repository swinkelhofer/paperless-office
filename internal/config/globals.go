package config

import "sync"

var (
	once     sync.Once
	instance Configuration
)
