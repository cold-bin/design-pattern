// @author cold bin
// @date 2023/8/30

package option

const (
	DefaultMaxTotal = 10
	DefaultMaxIdle  = 9
	DefaultMinIdle  = 1
)

type Pool struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

func NewPool(name string, opts ...Option) *Pool {
	// 填必须参数
	p := &Pool{name: name, maxTotal: DefaultMaxTotal, maxIdle: DefaultMaxIdle, minIdle: DefaultMinIdle}
	// 填入选项
	for _, opt := range opts {
		opt(p)
	}
	return p
}

type Option func(*Pool)

func WithMaxTotal(maxTotal int) Option {
	return func(pool *Pool) {
		pool.maxTotal = maxTotal
	}
}

func WithMaxIdle(maxIdle int) Option {
	return func(pool *Pool) {
		pool.maxIdle = maxIdle
	}
}

func WithMinIdle(minIdle int) Option {
	return func(pool *Pool) {
		pool.minIdle = minIdle
	}
}
