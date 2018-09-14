//proxy 代理模式

package proxy

type Subject interface {
	Do() string
}

type RealSubject struct{}

func (r *RealSubject) Do() string {
	return "real"
}

type Proxy struct {
	real RealSubject
}

func (p Proxy) Do() string {
	var res string

	//在调用之前
	res += "pre:"

	res += p.real.Do()

	//在调用之后

	res += ":after"

	return res
}
