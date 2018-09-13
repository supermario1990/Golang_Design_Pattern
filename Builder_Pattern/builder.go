package builder

//Builder 生成器接口
type Builder interface {
	Part1()
	Part2()
	Part3()
}

type Director struct {
	build Builder
}

// NewDirector new a Director
func NewDirector(build Builder) *Director {
	return &Director{
		build: build,
	}
}

//construct product
func (d *Director) Construct() {
	d.build.Part1()
	d.build.Part2()
	d.build.Part3()
}

type Builder1 struct {
	result string
}

func (b *Builder1) Part1() {
	b.result += "1"
}

func (b *Builder1) Part2() {
	b.result += "2"
}

func (b *Builder1) Part3() {
	b.result += "3"
}

func (b *Builder1) GetResult() string {
	return b.result
}
