package noshorthandinit

type Person struct {
	Name string
	Age  int
}

type Empty struct{}

func nested() {
	_ = Person{Name: "Bob", Age: 25}
	_ = Person{"Alice", 30} // want `avoid short-hand struct initialization`
}

func stuff() {
	_ = interface{}(nil)
	_ = []byte{1, 2, 3}
	_ = map[string]int{"a": 1, "b": 2}
	_ = Empty{}
	_ = Person{Name: "Bob", Age: 25}
	_ = Person{"Alice", 30} // want `avoid short-hand struct initialization`
	f := func(Name string, Age int) {
		_ = Person{Name, Age} // want `avoid short-hand struct initialization`
		_ = Person{Name: "Alice", Age: 30}
	}
	f("Bob", 25)
}
