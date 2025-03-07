package behaviorconstraint

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type user struct {
	name  string
	email string
}

func (u user) String() string {
	return fmt.Sprintf("{type: \"user\", name: %q, email: %q}", u.name, u.email)
}

type customer struct {
	name  string
	email string
}

func (u customer) String() string {
	return fmt.Sprintf("{type: \"customer\", name: %q, email: %q}", u.name, u.email)
}

func stringifyUsers(users []user) []string {
	ret := make([]string, 0, len(users))
	for _, user := range users {
		ret = append(ret, user.String())
	}
	return ret
}

func stringifyCustomers(customers []customer) []string {
	ret := make([]string, 0, len(customers))
	for _, customer := range customers {
		ret = append(ret, customer.String())
	}
	return ret
}

func stringifyAssert(v interface{}) []string {
	switch list := v.(type) {
	case []user:
		ret := make([]string, 0, len(list))
		for _, value := range list {
			ret = append(ret, value.String())
		}
		return ret

	case []customer:
		ret := make([]string, 0, len(list))
		for _, value := range list {
			ret = append(ret, value.String())
		}
		return ret
	}

	return nil
}

func stringifyReflect(v interface{}) []string {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Slice {
		return nil
	}

	ret := make([]string, 0, val.Len())

	for i := 0; i < val.Len(); i++ {
		m := val.Index(i).MethodByName("String")
		if !m.IsValid() {
			return nil
		}

		data := m.Call(nil)
		ret = append(ret, data[0].String())
	}

	return ret
}

func stringify[T fmt.Stringer](slice []T) []string {
	ret := make([]string, 0, len(slice))
	for _, value := range slice {
		ret = append(ret, value.String())
	}
	return ret
}

func Example() {
	users := []user{
		{
			name:  "Bill",
			email: "bill@ardanlabs.com",
		},
		{
			name:  "Ale",
			email: "ale@whatever.com",
		},
	}

	u1 := stringifyUsers(users)
	fmt.Println("users:", u1)

	u2 := stringifyAssert(users)
	fmt.Println("users:", u2)

	u3 := stringifyReflect(users)
	fmt.Println("users:", u3)

	u4 := stringify(users)
	fmt.Println("users:", u4)

	customers := []customer{
		{
			name:  "Google",
			email: "you@google.com",
		},
		{
			name:  "MSFT",
			email: "you@msft.com",
		},
	}

	c1 := stringifyCustomers(customers)
	fmt.Println("customers:", c1)

	c2 := stringifyCustomers(customers)
	fmt.Println("customers:", c2)

	c3 := stringifyReflect(customers)
	fmt.Println("customers:", c3)

	c4 := stringify(customers)
	fmt.Println("customers:", c4)

	user := User{
		Name: "Alice",
		Age:  30,
	}

	jsonData, err := marshal(user)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}

func marshal[T json.Marshaler](value T) ([]byte, error) {
	return value.MarshalJSON()
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(&struct {
		Alias
	}{
		Alias: Alias(u),
	})
}
