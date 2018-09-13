//adapter 适配器

package adapter

//Target 适配的目标接口
type Target interface {
	Request() string
}

//Adaptee 被适配的目标接口
type Adaptee interface {
	SpecificRequest() string
}

type adapteeImpl struct{}

func (a *adapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

type adapter struct {
	Adaptee
}

func (a *adapter) Request() string {
	return a.SpecificRequest()
}

func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}

func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}
