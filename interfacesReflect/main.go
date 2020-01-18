package main

import (
	"fmt"
	"math"
	"reflect"
)

/*
	An interface is a type that holds a set of methods that can be called for a variety of types
*/

// Animal ... This interface 'Animal' has the method speak. Thus, an 'animal' is that which can speak
type Animal interface {
	Speak() string
}

// Life implements animal in addition to others
type Life interface {
	Animal
	Grow()
}

// Dog ...
type Dog struct {
}

// Speak ...
func (d Dog) Speak() string {
	return "Woof!"
}

// Cat ...
type Cat struct {
}

// Speak ... note that this is a pointer to Cat
func (c *Cat) Speak() string {
	return "Meow!"
}

// Llama ...
type Llama struct {
}

// Speak ...
func (l Llama) Speak() string {
	return "?????"
}

func (l *Llama) Grow() {
	fmt.Println("I am growing all the time")
	return
}

// JavaProgrammer ...
type JavaProgrammer struct {
	name string
}

func (j *JavaProgrammer) Drink() {
	j.name = "Super" + j.name
	fmt.Printf("This is from the Drink function: %s\n", j.name)
	return
}

func (j *JavaProgrammer) Grow() {
	fmt.Println("I am growing all the time!")
	return
}

// Speak ...
func (j JavaProgrammer) Speak() string {
	return "Design Patterns!"
}

// DoSomething ...
func DoSomething(v interface{}) {
	// This function will accept any parameter - because all types implement 0 methods
	// ! this is fugly code - the better approach is to create a receiver function that takes an interface that is implemented by all expected types
	// ! only here for demo purposes
	// Use type checking to see whether the input is iterable
	// if v, ok := v.([]string); ok {
	// 	for _, value := range v {
	// 		fmt.Println(value)
	// 	}
	// 	return
	// }
	// Or use switch with a default handler
	fmt.Println(math.Pow(2, 3))
	switch x := v.(type) {
	case string:
		fmt.Println(x)
		return
	case []string:
		for _, value := range x {
			fmt.Println(value)
		}
		return
	case int:
		fmt.Println(x)
	default:
		// Pass silently
		return
	}
	// fmt.Println(v)
	// return
}

// This will demo the reflect package
func doSomethingElse(v interface{}) {
	fmt.Printf("Type is: %s\n", reflect.TypeOf(v))
}

func doEvenMoreStuff(v interface{}) {
	i := reflect.ValueOf(v)
	switch i.Kind() {
	case reflect.Float64:
		fmt.Println("We have determined that this is a Float64")
		break
	case reflect.Float32:
		fmt.Println("We have determined that this is a Float32")
		break
	case reflect.Int:
		fmt.Println("We have determined that this is an Int")
		break
	case reflect.Int8:
		fmt.Println("We have determined that this is an Int8")
		break
	default:
		fmt.Println("We have determined that this is a .... ")
		break
	}
	return
}

// ! This is the correct approach to creating a function that accepts any type.
// At runtime, any type that implements the interface can be passed here and its method will be called
func speak(a Animal) {
	fmt.Printf("%s says: %s\n", reflect.TypeOf(a).Name(), a.Speak())
}

func main() {
	animals := []Animal{new(Dog), new(Cat), Llama{}, JavaProgrammer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

	// example of an interface that implements another interfaces
	// Just did some different syntax here to loop through the slice for grins
	lives := []Life{new(JavaProgrammer), new(Llama)}
	for j := 0; j < len(lives); j++ {
		lives[j].Grow()
	}

	newProgrammer := JavaProgrammer{"Blaine"}
	newProgrammer2 := JavaProgrammer{"James"}

	fmt.Println()
	fmt.Println("==== USE THE CORRECT APPROACH ====")
	fmt.Println()
	speak(newProgrammer)

	newProgrammer.Drink()
	// newProgrammer2.Drink()

	// JavaProgrammer.Drink(newProgrammer2)

	fmt.Println(newProgrammer.name)
	fmt.Println(newProgrammer2.name)

	fmt.Println()
	fmt.Println("==== REFLECT BASICS ====")
	fmt.Println()

	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x).String())
	fmt.Println("kind:", reflect.Kind(x).String())

	fmt.Println()
	fmt.Println("==== REFLECT VALUEOF ====")
	fmt.Println()

	var y float64 = 5.6654
	v := reflect.ValueOf(y)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64) // this is a useful comparitor. reflect has these types
	fmt.Println("value:", v.Float())                             // 5.6654

	fmt.Println()
	fmt.Println("==== DO SOMETHING ====")
	fmt.Println()

	DoSomething(12)       // Call DoSomething on an integer
	DoSomething("Blaine") // Call DoSomething on a string
	DoSomething([]string{"List Elem 1", "List Elem 2", "List Elem 3"})

	fmt.Println()
	fmt.Println("==== DO SOMETHING ELSE ====")
	fmt.Println()

	doSomethingElse([]string{"Blaine"})

	fmt.Println()
	fmt.Println("==== HERE BE DRAGONS ====")
	fmt.Println()

	var h uint8 = 'h'                                         // remember that this is a rune - hence why it is of the uint8 sort
	i := reflect.ValueOf(h)                                   // align
	fmt.Println("type:", i.Type())                            // uint8.
	fmt.Println("kind:", i.Kind())                            // uint8.
	fmt.Println("kind is uint8: ", i.Kind() == reflect.Uint8) // true.

	h = uint8(i.Uint()) // v.Uint returns a uint64.
	fmt.Println(h)      // 104 - creates a new variable

	fmt.Println()
	fmt.Println("==== DO EVEN MORE STUFF ====")
	fmt.Println()

	var j int8 = 1
	doEvenMoreStuff(j)

	fmt.Println()
	fmt.Println("==== TYPING ====")
	fmt.Println()
	type MyInt int
	var k MyInt = 7
	l := reflect.ValueOf(k)
	fmt.Println("type:", l.Type())
	fmt.Println("kind:", l.Kind()) // Note that the kind gives the underlying type of 'l'; despite that l's static type is MyInt

	fmt.Println()
	fmt.Println("==== REFLECTION: GENERATING THE INVERSE ====")
	fmt.Println()
	// m := v.Interface().(float64) // y will have type float64.
	// ! this is taking the valueOf - reflect back "interface values to reflection objects and back again"
	testVar := v.Interface().(float64)

	fmt.Printf("The type of testVar is %T\n", testVar)

	fmt.Println(v.Interface().(float64))
	fmt.Printf("value is %1.2e\n", v.Interface())

	fmt.Println()
	fmt.Println("==== REFLECT AND SET ====")
	fmt.Println()
	var m float64 = 3.4
	n := reflect.ValueOf(m)
	// n.SetFloat(7.1)
	// ! Error: will panic. "panic: reflect.Value.SetFloat using unaddressable value"
	// ! This is because "m" here is not settable `Settability is a property of a reflection Value, and not all reflection Values have it.`

	// Use CanSet method to determine whether the value is settable
	fmt.Println("Settability of 'n' is:", n.CanSet())

	// ! Settability ... is the property that a reflection object can modify the actual storage that was used to create the reflection object. ...
	// ! determined by whether the reflection object holds the original item.
	// ! In the above example, a copy of the variable m was sent to `reflect.ValueOf`; not the actual variable.

	var o float64 = 55.6
	p := reflect.ValueOf(&o)
	fmt.Println("Type of 'p' is:", p.Type())
	fmt.Println("Settability of 'p' is:", p.CanSet())
	// ! Note that the .Elem() method here determines that to which 'p' points
	q := p.Elem()
	fmt.Println("However, settability of 'p.Elem()' is:", q.CanSet())
	q.SetFloat(5.6)
	fmt.Println("The output of 'p.Elem().SetFloat(5.6)' - in other words, the originating variable `o`'s new value - is:", o)

	fmt.Println()
	fmt.Println("==== USING THESE METHODS TO ALTER STRUCTS' VALUES ====")
	fmt.Println()
	type T struct {
		A int
		B string
	}
	t := T{23, "skidoo"}
	// ! note that we are using a pointer in the `reflect.ValueOf` method, because we will need it later
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type() // ? This = `main.T`. The kind (`.Kind()`) = `struct`.
	// ? The `reflect.NumField` method can only be called on reflect.ValueOf objects of the kind `struct`
	// ? only Exported fields of a struct are settable
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
		switch f.Kind() {
		case reflect.Int:
			s.Field(i).SetInt(545)

		case reflect.String:
			s.Field(i).SetString("blainesworld")

		default:
			// empty return
		}
	}
	// ? `Field` returns the i'th field of the struct s. It panics if s's Kind is not Struct or i is out of range.
	// s.Field(0).SetInt(77)
	// s.Field(1).SetString("Blainesworld")
	// ? If we modified the program so that 's' was created from 't', not '&t',
	// ? the calls to SetInt and SetString would fail as the fields of 't' would not be settable.
	fmt.Println("t is now", t)

	// ! "Once you understand these laws reflection in Go becomes much easier to use,
	// ! although it remains subtle.
	// ! It's a powerful tool that should be used with care and avoided unless strictly necessary.
}
