package backend

import "fmt"

func ExampleConnection() {
	conn := GetConnection()
	//	conn.InitDatabase("unichain")

	int_res := conn.SetTransactionToBacklog(`{"id":"5556","back":"j22222ihhh"}`)
	fmt.Println(int_res)
	map_string := conn.GetTransactionFromBacklog("5556")
	fmt.Printf("tx:%s\n", map_string)

	// Output:
	//{0 1 0 0 0 0 0 0 0 0 0 0 0 0 []  [] []}1
	//tx:{"back":"j22222ihhh","id":"5555"}
}