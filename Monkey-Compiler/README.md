<br>

# Monkey compiler 🐒

A Monkey programming language compiler from  [Writing A Compiler In Go](https://compilerbook.com).

<br>

## Building

To build the binary, enter

```sh
go build
```

<br>

## REPL

To run the REPL, enter

```
$ ./monkey-compiler
This is the Monkey programming language!
Feel free to type in commands
>> puts("Hello, world!")
Hello, world!
>> 
```

<br>

## Scripts

This compiler also supports running a single Monkey script file. For example,

```sh
$ ./monkey-compiler scripts/hello.monkey
Hello, world!
```

<br>

## The Monkey Language

<br>

### Number types and variable bindings

Define and reassign to variables using the `=` operator. Variables are dynamically typed and can be assigned to objects of any type in Monkey. You can use `let` keyword when defining variables, but it's completely optional and there is no difference between with and without `let` keyword. 

Two number types are supported in this implementation: integers and floating-point numbers.

```sh
>> let a = 1;  # Assignment with `let` keyword
>> a
1
>> b = 2.5;  # Assignment without `let` keyword
>> b
2.5
>> b = "a";  # Reassignment to b
>> b
a
```

<br>

### Arithmetic and comparison expressions

You can do basic arithmetic and comparison operations for numbers, such as `+`, `-`, `*`, `/`, `<`, `>`, `<=`, `>=`, `==`, `!=`, `&&` and `||`.

```sh
>> let a = 10;
>> let b = a * 2;
>> (a + b) / 2 - 3;
12
>> let c = 2.25;
>> let d = -5.5;
>> b + c * d
7.625
>> a < b
true
>> c == d
false
```

<br>

### If expressions

You can use `if` and `else` keywords for conditional expressions. The last value in an executed block is returned from the expression.

```sh
>> let a = 10;
>> let b = a * 2;
>> let c = if (b > a) { 99 } else { 100 };
>> c
99
>> let d = if (b > a && c < b) { 199 } else { 200 };
>> d
200
```

<br>

### Functions and closures

You can define functions using `fn` keyword. All functions are closures in Monkey and you have to use `let` along with `fn` to bind a closure to a variable. Closures close over an environment where they are defined, and are evaluated in *the* environment when called. The last value in an executed function body is returned as a return value.

```sh
>> let multiply = fn(x, y) { x * y };
>> multiply(50 / 2, 1 * 2)
50
>> fn(x) { x + 10 }(10)
20
>> let newAdder = fn(x) { fn(y) { x + y }; };
>> let addTwo = newAdder(2);
>> addTwo(3);
5
>> let sub = fn(a, b) { a - b };
>> let applyFunc = fn(a, b, func) { func(a, b) };
>> applyFunc(10, 2, sub);
8
```

<br>

### Strings

You can build strings using a pair of double quotes `""`. Strings are immutable values just like numbers. You can concatenate strings with `+` operator.

```sh
>> let makeGreeter = fn(greeting) { fn(name) { greeting + " " + name + "!" } };
>> let hello = makeGreeter("Hello");
>> hello("John");
Hello John!
```

<br>

### Arrays

You can build arrays using square brackets `[]`. Array literal is `[value1, value2, ...]`. Arrays can contain values of any type, such as integers, strings, even arrays and functions (closures). To get an element at an index from an array, use `array[index]` syntax. To set a value at an index in an array to another value, use `array[index] = value` syntax.

```sh
>> let myArray = ["Thorsten", "Ball", 28, fn(x) { x * x }];
>> myArray[0]
Thorsten
>> myArray[4 - 2]
28
>> myArray[3](2);
4
>> myArray[2] = myArray[2] + 1
>> myArray[2]
29
```

<br>

### Hash maps

You can build hash maps using curly brackets `{}`. Hash literal is `{key1: value1, key2: value2, ...}`. You can use numbers, strings and booleans as keys, and objects of any type as values. To get a value under a key from a hash map, use `hash[key]` syntax. To set a value under a key in a hash map to another value, use `hash[key] = value` syntax.

```sh
>> let myHash = {"name": "Jimmy", "age": 72, true: "yes, a boolean", 99: "correct, an integer"};
>> myHash["name"]
Jimmy
>> myHash["age"]
72
>> myHash[true]
yes, a boolean
>> myHash[99]
correct, an integer
>> myHash[0] = "right, zero"
>> myHash[0]
right, zero
```

<br>

### Built-in functions

There are some built-in functions in Monkey.

<br>

#### `len`

`len` built-in function allows you to get the length of strings or arrays. Note that `len` returns the number of bytes instead of characters for strings.

```sh
>> len("hello");
5
>> len("∑");
3
>> let myArray = ["one", "two", "three"];
>> len(myArray)
3
```

<br>

#### `puts`

`puts` built-in function allows you to print out one or more objects to console (i.e. stdout).

```sh
>> puts("Hello, World")
Hello, World
```

<br>

#### `first`

`first` built-in function allows you to get the first element from an array. If the array is empty, `first` returns `nil`.

```sh
>> let myArray = ["one", "two", "three"];
>> first(myArray)
one
>> first([])
nil
```

<br>

#### `last`

`last` built-in function allows you to get the last element from an array. If the array is empty, `last` returns `nil`.

```sh
>> let myArray = ["one", "two", "three"];
>> last(myArray)
three
>> last([])
nil
```

<br>

#### `rest`

`rest` built-in function allows you to create a new array containing all elements of a given array except the first one. If the array is empty, `rest` returns `nil`.

```sh
>> let myArray = ["one", "two", "three"];
>> rest(myArray)
[two, three]
>> rest([])
nil
```

<br>

#### `push`

`push` built-in function allows you to add a new element to the end of an existing array. It allocates a new array instead of modifying the given one.

```sh
>> let myArray = ["one", "two", "three"];
>> push(myArray, "four")
[one, two, three, four]
```

<br>

#### `quote` / `unquote`

Special function, `quote`, returns an unevaluated code block (think it as an AST). Opposite function to `quote`, `unquote`, evaluates code inside `quote`.

```sh
>> quote(2 + 2)
Quote((2 + 2)) # Unevaluated code
>> quote(unquote(1 + 2))
Quote(3)
```

<br>

### Single-line comments

Comments begin with a hash mark (`#`) and continue to the end of the line. Thery are ignored by the compiler.

```sh
>> # This line is just a comment.
>> let a = 1;  # This is an integer.
1
```

<br>

### Macros

You can define macros using `macro` keyword. Note that macro definitions must return `Quote` objects generated from `quote` function.

```sh
# Define `unless` macro which does the opposite to `if`
>> let unless = macro(condition, consequence, alternative) {
     quote(
       if (!(unquote(condition))) {
         unquote(consequence);
       } else {
         unquote(alternative);
       }
     );
   };
>> unless(10 > 5, puts("not greater"), puts("greater"));
greater
nil
```

<br>