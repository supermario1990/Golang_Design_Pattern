// 对象池模型

package pool

// Object define
type Object struct {
	Desc string
}

// Pool defile
type Pool chan *Object

// New Object pool
func New(totle int) *Pool {
	p := make(Pool, totle)
	for i := 0; i < totle; i++ {
		p <- new(Object)
	}
	return &p
}
