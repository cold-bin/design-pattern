// @author cold bin
// @date 2023/8/29

package factory

// ConfigFactory Config 的工厂方法抽象
type ConfigFactory interface {
	Create() Config
}

type JsonFactory struct{}

func (j JsonFactory) Create() Config {
	return Json{}
}

type YamlFactory struct{}

func (y YamlFactory) Create() Config {
	return Json{}
}

// NewConfigFactory 这里使用简单工厂的方式创建工厂
//
//	与简单工厂不同的是，这里拿的还是工厂，而不是直接的对象
func NewConfigFactory(typ string) ConfigFactory {
	switch typ {
	case "json":
		return JsonFactory_
	case "yaml":
		return YamlFactory_
	}
	return nil
}

// 工厂复用
var (
	JsonFactory_ = JsonFactory{}
	YamlFactory_ = YamlFactory{}
)
