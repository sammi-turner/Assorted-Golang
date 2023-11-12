package main

import "fmt"

func createTable() {
	if DbExists("kv") {
		fmt.Println("\nTable already exists.")
	} else {
		CreateKeyValueTable("kv")
		InsertKeyValuePair("kv", "John", "Plumber")
		InsertKeyValuePair("kv", "Bill", "Firefighter")
		InsertKeyValuePair("kv", "Susan", "Engineer")
	}
}

func queryTable() {
	v0 := SelectAllPairs("kv")
	v1 := SelectAllKeys("kv")
	v2 := SelectAllValues("kv")

	DeleteKeyValuePair("kv", "Bill")
	v3 := TableContainsKey("kv", "Bill")

	v4 := SelectAllKeys("kv")
	v5 := SelectAllValues("kv")

	v6 := SelectRowFromKey("kv", "John")
	v7 := SelectValueFromKey("kv", "Susan")

	fmt.Printf("\n%s\n\n%s\n\n%s\n\n%t\n\n%s\n\n%s\n\n%s\n\n%s\n\n", v0, v1, v2, v3, v4, v5, v6, v7)
}

func main() {
	v, _ := IsSQLiteInstalled()
	if v {
		createTable()
		queryTable()
	} else {
		fmt.Println("SQLite is not installed.")
	}
}
