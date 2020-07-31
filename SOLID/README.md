# 模式

SOLID 是基本指导原则，用于帮助设计稳健的软件。

- [OCP - Open Closed Principle](#ocp---open-closed-principle)
- [SRP - Single Responsibility Principle](#srp---single-responsibility-principle)
- [LSP - Liskov Substitution Principle](#lsp---liskov-substitution-principle)
- [LoD - Law of Demeter](#lod---law-of-demeter)
- [ISP - Interface Segregation Principle](#isp---interface-segregation-principle)
- [DIP - Dependence Inversion Principle](#dip---dependence-inversion-principle)

## OCP - Open Closed Principle

开闭原则：一个软件实体如类、模块和函数应该对扩展开放，对修改关闭。

> Software entities like classes,modules and functions should be open for extension but closed for modifications.

通过扩展而不是修改来实现变化。

```go
package main

import "fmt"

type calculate struct {
	a, b int
}

func newCalculate(a, b int) *calculate {
	return &calculate{
		a: a,
		b: b,
	}
}

func (c *calculate) do() int {
	return c.a + c.b
}

func main() {
	c := newCalculate(2, 3)
	fmt.Println(c.do())
}
```

在上面的代码中, `calculate.do()` 会执行加法运算。如果想要其执行乘法运算。就必须要修改成

```go
func (c *calculate) do() int {
	return c.a * c.b
}
```

这样就违反了开闭原则。因为需要修改元代码，才能实现加法变乘法。

如果运用[策略方法](../DesignPatterns/strategy)，就可以不违背开闭原则。

```go
package main

import "fmt"

type calculate struct {
	a, b int
	executer
}

type executer interface {
	execute(int, int) int
}

func newCalculate(a, b int, e executer) *calculate {
	return &calculate{
		a:        a,
		b:        b,
		executer: e,
	}
}

func (c *calculate) do() int {
	return c.execute(c.a, c.b)
}

type add struct {
}

func (a add) execute(i, j int) int {
	return i + j
}

type multiply struct {
}

func (m multiply) execute(a, b int) int {
	return a * b
}

func main() {
	c1 := newCalculate(2, 3, add{})
	fmt.Println("calculate result is ", c1.do())

	c2 := newCalculate(2, 3, multiply{})
	fmt.Println("calculate result is ", c2.do())
}

```

通过注入不同的计算方式，来修改计算方法。以后还需要修改成除法或者减法的时候。也只需要添加相关的计算方式即可。

## SRP - Single Responsibility Principle

单一职责原则: 每个类被修改的理由，只有一个。

> There should never be more than one reason for a class to change.

- 类的职责越多，被修改的可能性就越大，这个类就越不稳健。
- 单一职责也同样适用于接口和方法。
- 越是抽象，越要遵守单一职责原则。所以，接口 > 类 > 方法。

## LSP - Liskov Substitution Principle

里氏替换原则:使用父类的程序，不用修改，就可以使用子类替代父类。

> If for each object o1 of type S there is an object o2 of type T such that for all programs P defined in terms of T,the behavior of P is unchanged when o1 is substituted for o2 then S is a subtype of T.

在使用父类的程序看来: 子类与父类

- 拥有相同的公开属性和公开方法
- 相同方法输入和输出类型一致。

所以，程序根本无法区分子类与父类。所以，可以直接进行替换。

总之，子类必须是父类的超集。

## LoD - Law of Demeter

迪米特法则：一个对象应该对其他对象有最少的了解。

- 只与朋友说话。 朋友是指出现在成员变量、方法的输入输出参数中的类。出现在方法体内部的类不属于朋友类。所以在自己的方法内，只调用朋友的方法，不要引入非朋友类，形成依赖。
- 跟朋友也不要说太多话。如果调用了朋友的三四个方法组成了一个流程，那么说明自己的这个方法的抽象级别不对，应该把此流程下放到朋友的新方法中。就只有调用朋友的那一个新方法了。
- 自己的事自己做。对于可以放在本类或者朋友类的方法，如果这个方法放在本类中，既不增加类间关系，也对本类不产生负面影响，那就放置在本类中。

## ISP - Interface Segregation Principle

接口隔离原则：客户端不应该依赖它不需要的接口，类间的依赖关系应该建立在最小的接口上。

> Clients should not be forced to depend upon interfaces that they don't use, The dependency of one class to another one should depend on the smallest possible interface.

接口要小，尽可能的小。

## DIP - Dependence Inversion Principle

依赖倒置原则:高层模块不应该依赖低层模块，两者都应该依赖其抽象；抽象不应该依赖细节；细节应该依赖抽象

> High level modules should not depend upon low level modules.Both should depend upon abstractions.Abstractions should not depend upon details.Details should depend upon abstractions.

高层模块依赖于底层模块

```go
type Tesla strut{}

func (t *Tesla) Run() {}

type Driver strut{
	Car *Tesla
}

func newDriver(t *Tesla) *Driver {
	return &Driver{
		Car: t,
	}
}

// 高层模块 Driver 依赖了低层模块 Tesla
// Driver 只能驾驶 Tesla 汽车
func (d *Driver) Drive() {
	d.Car.Run()
}
```

高层模块和依赖低层模块，都依赖其抽象

```go
// Car 接口是抽象的
type Car interface{
	Run()
}

type Tesla strut{}

// Tesla 实现了 Car 接口
// 所以，Tesla 依赖于 Car 接口
func (t *Tesla) Run() {}

// Driver 包含了 Car 接口
// 所以，Driver 依赖于 Car 接口
type Driver strut{
	Car Car
}

// Driver 依赖于 Car 接口
func newDriver(c Car) *Driver {
	return &Driver{
		Car: c,
	}
}

// Driver 可以驾驶所有实现了 Car 接口的汽车
func (d *Driver) Drive() {
	d.Car.Run()
}
```

- 越高层的模块越稳定
- 越抽象的接口越稳定
- 依赖倒置想要不稳定的依赖稳定的
