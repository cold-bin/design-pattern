// @author cold bin
// @date 2023/8/28

package factory

type Config interface {
	Parse(data []byte) error
	UnParse(src []byte, dst []byte) error
}

type Json struct {
}

func (j Json) Parse(data []byte) error {
	panic("implement me")
}

func (j Json) UnParse(src []byte, dst []byte) error {
	panic("implement me")
}

type Yaml struct {
}

func (y Yaml) Parse(data []byte) error {
	panic("implement me")
}

func (y Yaml) UnParse(src []byte, dst []byte) error {
	panic("implement me")
}

// NewConfig 这里直接使用简单工厂方法创建最终对象
//
//	如果创建对象不复杂，不涉及对象之间的组合，可以使用
func NewConfig(name, typ string) Config {
	switch typ {
	case "json":
		return &Json{}
	case "yaml":
		return &Yaml{}
	}
	return nil
}
