// @author cold bin
// @date 2023/8/30

package builder

import "errors"

const (
	DefaultMaxTotal = 10
	DefaultMaxIdle  = 9
	DefaultMinIdle  = 1
)

type Build struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

// Build 建造器
//
//	这里负责参数校验，通过后真正new对象
func (b *Build) Build() (*Pool, error) {
	if b.maxIdle <= 0 {
		b.maxIdle = DefaultMaxIdle
	}
	if b.minIdle <= 0 {
		b.minIdle = DefaultMinIdle
	}
	if b.maxTotal <= 0 {
		b.maxTotal = DefaultMaxTotal
	}
	if b.minIdle > b.maxIdle {
		return nil, errors.New("minIdle is more than maxIdle")
	}

	if b.maxIdle > b.maxTotal {
		return nil, errors.New("maxIdle is more than maxTotal")
	}

	return &Pool{
		name:     b.name,
		maxTotal: b.maxTotal,
		maxIdle:  b.maxIdle,
		minIdle:  b.minIdle,
	}, nil
}

func (b *Build) SetName(name string) {
	b.name = name
}

func (b *Build) SetMaxTotal(maxTotal int) {
	b.maxTotal = maxTotal
}

func (b *Build) SetMaxIdle(maxIdle int) {
	b.maxIdle = maxIdle
}

func (b *Build) SetMinIdle(minIdle int) {
	b.minIdle = minIdle
}

type Pool struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}
