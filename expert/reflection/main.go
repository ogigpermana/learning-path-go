// File: expert/reflection/main.go
// Level: Expert
// Topik: Reflection (reflect package)
//
// Reflection adalah kemampuan untuk memeriksa dan memanipulasi
// struktur program saat runtime. Berguna untuk:
// - Serialization/deserialization (JSON, XML, dll)
// - ORM (mapping struct ke database)
// - Validation library
// - Dependency injection
//
// Peringatan: hindari reflection jika ada alternatif lain (slow, complex)

package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"min=0"`
	City string `json:"city,omitempty"`
}

func (p Person) SayHello() string {
	return fmt.Sprintf("Hi, I'm %s", p.Name)
}

func (p Person) Greet(greeting string) string {
	return fmt.Sprintf("%s, I'm %s", greeting, p.Name)
}

func main() {
	// 1. TYPE & KIND
	fmt.Println("=== Type & Kind ===")
	p := Person{Name: "Anggi", Age: 20}
	t := reflect.TypeOf(p)

	fmt.Println("Type:", t)
	fmt.Println("Kind:", t.Kind())
	fmt.Println("Package:", t.PkgPath())
	fmt.Println()

	// 2. STRUCT FIELDS
	fmt.Println("=== Struct Fields ===")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field %d: name=%s, type=%s, tag=`%s`\n",
			i, field.Name, field.Type, field.Tag)
	}
	fmt.Println()

	// 3. GET & SET VALUES
	fmt.Println("=== Get & Set Values ===")
	v := reflect.ValueOf(&p).Elem() // pointer untuk bisa di-set

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fmt.Printf("Field %d: %v (canSet=%v)\n", i, field.Interface(), field.CanSet())
	}

	// Set value
	v.FieldByName("Name").SetString("Budi")
	v.FieldByName("Age").SetInt(25)
	fmt.Println("After set:", p)
	fmt.Println()

	// 4. TAGS
	fmt.Println("=== Tags ===")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("%s: json=%s, validate=%s\n",
			field.Name,
			field.Tag.Get("json"),
			field.Tag.Get("validate"))
	}
	fmt.Println()

	// 5. METHODS
	fmt.Println("=== Methods ===")
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("Method %d: name=%s, type=%s\n",
			i, method.Name, method.Type)
	}

	// Call method via reflection
	m := v.MethodByName("SayHello")
	result := m.Call([]reflect.Value{})
	fmt.Println("Called SayHello:", result[0].Interface())

	// Call method with parameter
	m2 := v.MethodByName("Greet")
	result2 := m2.Call([]reflect.Value{reflect.ValueOf("Hello")})
	fmt.Println("Called Greet:", result2[0].Interface())
	fmt.Println()

	// 6. DYNAMIC FUNCTION CALL
	fmt.Println("=== Dynamic Function Call ===")
	fn := func(a, b int) int { return a + b }
	fnValue := reflect.ValueOf(fn)
	fnType := reflect.TypeOf(fn)

	fmt.Println("Function type:", fnType)
	args := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}
	fnResult := fnValue.Call(args)
	fmt.Println("10 + 20 =", fnResult[0].Interface())
	fmt.Println()

	// 7. COMPARE INTERFACES
	fmt.Println("=== DeepEqual ===")
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	c := []int{3, 2, 1}

	fmt.Println("a == b:", reflect.DeepEqual(a, b))
	fmt.Println("a == c:", reflect.DeepEqual(a, c))
	fmt.Println()

	// 8. CREATE TYPE AT RUNTIME
	fmt.Println("=== Create Type at Runtime ===")
	typeMap := map[string]interface{}{
		"name": "Anggi",
		"age":  20,
	}

	mapType := reflect.TypeOf(typeMap)
	fmt.Println("Map type:", mapType)
	fmt.Println("Key type:", mapType.Key())
	fmt.Println("Value type:", mapType.Elem())
}

/*
Reflection use cases:
- JSON/XML/YAML serialization
- GORM/Ent (ORM)
- Validator libraries
- go-sqlmock (mocking)
- Dependency injection frameworks (Wire, dig, fx)
- Testing tools (testify)
*/