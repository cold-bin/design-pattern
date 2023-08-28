// @author cold bin
// @date 2023/8/28

package singleton

import "sync"

/*
	单例模式的懒汉式：支持延迟加载，实现范式会导致频繁加锁和释放锁，而且并发度也低。双重检测版本
*/

var (
	log__ *log
	once  sync.Once
)

func GetLazyInstance() Logger {
	if log__ == nil {
		once.Do(func() {
			log__ = &log{}
		})
	}
	return log__
}
