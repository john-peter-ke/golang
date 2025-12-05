// https://refactoring.guru/design-patterns/singleton/go/example#example-0

package singleton

import "sync"

type ConfigManager struct {
	AppConfig map[string]string
}

var (
	once     sync.Once
	instance *ConfigManager
)

func GetInstance() *ConfigManager {
	once.Do(func() {
		instance = &ConfigManager{
			AppConfig: map[string]string{
				"env":     "dev",
				"port":    "8080",
				"timeout": "30s",
			},
		}
	})
	return instance
}
