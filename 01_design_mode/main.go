package main

import (
	chainofresponsibility "desgin_model/01_chainofresponsibility"
	command "desgin_model/02_command"
	iterator "desgin_model/03_iterator"
	mediator "desgin_model/04_mediator"
	memento "desgin_model/05_memento"
	observer "desgin_model/06_observer"
	state "desgin_model/07_state"
	strategy "desgin_model/08_strategy"
	templatemethod "desgin_model/09_templatemethod"
	visitor "desgin_model/10_visitor"
	interpreter "desgin_model/11_interpreter"
	adapter "desgin_model/12_adapter"
	bridge "desgin_model/13_bridge"
	composite "desgin_model/14_composite"
	decorator "desgin_model/15_decorator"
	facade "desgin_model/16_facade"
	flyweight "desgin_model/17_flyweight"
	proxy "desgin_model/18_proxy"
	factory "desgin_model/19_factory"
	abstract_factory "desgin_model/20_abstractfactory"
	builder "desgin_model/21_builder"
	prototype "desgin_model/22_prototype"
	singleton "desgin_model/23_singleton"
	"fmt"
)

func main() {
	fmt.Println("=== 01 chainofresponsibility ===")
	chainofresponsibility.ExecChainOfResponsibility()

	fmt.Println("\n=== 02 command ===")
	command.ExecuteCookCommand()

	fmt.Println("\n=== 03 iterator ===")
	iterator.ExecIterator()

	fmt.Println("\n=== 04 mediator ===")
	mediator.ExecMediator()

	fmt.Println("\n=== 05 memento ===")
	memento.ExecMemento()

	fmt.Println("\n=== 06 observer ===")
	observer.ExecObserver()

	fmt.Println("\n=== 07 state ===")
	state.ExecState()

	fmt.Println("\n=== 08 strategy ===")
	strategy.ExecStrategy()

	fmt.Println("\n=== 09 templatemethod ===")
	templatemethod.ExecTemplateMethod()

	fmt.Println("\n=== 10 visitor ===")
	visitor.ExecVisitor()

	fmt.Println("\n=== 11 interpreter ===")
	interpreter.ExecInterpreter()

	fmt.Println("\n=== 12 adapter ===")
	adapter.ExecAdapter()

	fmt.Println("\n=== 13 bridge ===")
	bridge.ExecBridge()

	fmt.Println("\n=== 14 composite ===")
	composite.ExecComposite()

	fmt.Println("\n=== 15 decorator ===")
	decorator.ExecDecorator()

	fmt.Println("\n=== 16 facade ===")
	facade.ExecFacade()

	fmt.Println("\n=== 17 flyweight ===")
	flyweight.ExecFlyweight()

	fmt.Println("\n=== 18 proxy ===")
	proxy.ExecProxy()

	fmt.Println("\n=== 19 factory ===")
	factory.ExecFactoryMethod()

	fmt.Println("\n=== 20 abstract_factory ===")
	abstract_factory.ExecAbstractFactory()

	fmt.Println("\n=== 21 builder ===")
	builder.ExecBuilder()

	fmt.Println("\n=== 22 prototype ===")
	prototype.ExecPrototype()

	fmt.Println("\n=== 23 singleton ===")
	singleton.ExecSingleton()
}
