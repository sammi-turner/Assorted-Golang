<br>

# Key Value Pair Thing (KVPT)

KVPT is a lightweight library that provides a convenient way to store, update, and query key-value pairs using SQLite tables. It simplifies the process of working with key-value data by abstracting away the underlying SQLite operations.

<br>

## Features

- Create SQLite tables to store key-value pairs
- Insert key-value pairs into tables
- Retrieve all key-value pairs from a table
- Retrieve all keys from a table
- Retrieve all values from a table
- Delete a key-value pair from a table
- Check if a table contains a specific key
- Retrieve a row based on a key
- Retrieve a value based on a key

<br>

## Requirements

KVPT has no compile-time dependencies, but it requires SQLite3 to be installed as a runtime dependency for the functions to create and interact with key-value stores.

To check if you have SQLite3 installed on your machine, run:

```
sqlite3 --version
```

<br>

## Usage

Below are some examples drawn from the main.go file.

```go
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
```

<br>

## Functions

KVPT provides the following functions:

- `db_exists: string -> bool`: Checks if a SQLite database file exists.
- `shell: string -> string`: Executes a shell command and returns the output as a string.
- `is_sqlite_installed: unit -> bool`: Checks if SQLite3 is installed on the system.
- `create_key_value_table: string -> string`: Creates a SQLite table with the specified name to store key-value pairs.
- `insert_key_value_pair: string -> string -> string -> string`: Inserts a key-value pair into a specified table.
- `select_all_pairs: string -> string`: Retrieves all key-value pairs from a specified table.
- `select_all_keys: string -> string`: Retrieves all keys from a specified table.
- `select_all_values: string -> string`: Retrieves all values from a specified table.
- `delete_key_value_pair: string -> string -> string`: Deletes a key-value pair from a specified table.
- `select_row_from_key: string -> string -> string`: Retrieves a row based on a specified key from a table.
- `select_value_from_key: string -> string -> string`: Retrieves a value based on a specified key from a table.
- `count_string_instances: string list -> string -> int`: Counts the occurrences of a string in a list of strings.
- `table_contains_key: string -> string -> bool`: Checks if a table contains a specific key.

<br>

## Warning: User Input

SQL injection is a serious security risk when user input strings are directly plugged into larger strings that execute code. Always remember to sanitize user input to prevent SQL injection attacks.

<img src="bobby_tables.png" width="100%" height="auto" />

<br>
