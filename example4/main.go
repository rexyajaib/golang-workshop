package main

import "fmt"

// Logger is a tiny example interface.
type Logger interface {
	Log(msg string)
}

// ----------- Concrete implementations -----------

type StdLogger struct{}

func (StdLogger) Log(msg string) {
	fmt.Println("[STD]", msg)
}

type MockLogger struct{}

func (MockLogger) Log(msg string) {
	fmt.Println("[MOCK]", msg)
}

// ----------- 1) Struct implements the interface -----------
//
// FileLogger (StdLogger) implements Logger because it defines Log.
// We can pass it anywhere Logger is expected (polymorphism).
func useLogger(l Logger) {
	l.Log("called useLogger")
}

// ----------- 2) Embedding an interface into a struct (dependency) -----------
//
// OrderService embeds Logger as a dependency. It does NOT itself implement Logger
// just by embedding the interface. Method calls are forwarded to the concrete value stored.

type OrderService struct {
	Logger // embedded interface field: promotes Logger methods for ergonomic calls
}

func (s OrderService) CreateOrder() {
	fmt.Println("OrderService: create order")
	// This forwards to the concrete Logger value held in s.Logger:
	s.Log("order created (from OrderService)")
}

// ----------- 3) Embedding a concrete type -----------
//
// If we embed a concrete type that implements Logger, the outer struct *does*
// implement Logger because methods are promoted from the concrete type.

type OrderWithImpl struct {
	StdLogger // embedding concrete type promotes its methods; OrderWithImpl implements Logger
}

func main() {
	fmt.Println("=== 1) Implementing interface (polymorphism) ===")
	var l Logger = StdLogger{} // StdLogger implements Logger
	useLogger(l)

	fmt.Println("\n=== 2) Embedding interface (dependency injection) ===")
	svc := OrderService{
		Logger: MockLogger{}, // inject a mock implementation
	}
	svc.CreateOrder()

	// We can swap the dependency at runtime:
	svc.Logger = StdLogger{}
	svc.CreateOrder()

	var x Logger = svc // compile run:
	x.Log("from OrderService")

	fmt.Println("\n=== 3) Embedding concrete implementation (reuse + implements) ===")
	var y Logger = OrderWithImpl{} // OK: methods promoted, OrderWithImpl implements Logger
	y.Log("from OrderWithImpl")
}
