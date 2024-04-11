<br>

# S-Expression REPL

This REPL supports the following arithmetic operations

- Addition: (add x y)
- Subtraction: (sub x y)
- Multiplication: (mul x y)
- Division: (div x y)
- Modulo: (mod x y)
- Exponentiation: (pow x y)

Here, x and y are numbers (integers or floating-point values).

<br>

## Evaluation

To evaluate an expression, simply enter it at the prompt and press Enter. 

The REPL will parse the expression, evaluate it, and print the result.

Here are a few examples

```
> (add 3 4)
7

> (sub 10 5)
5

> (mul 2 6)
12

> (div 10 2)
5

> (mod 7 3)
1

> (pow 2 3)
8
```

<br>

## Nested Operations

You can enter expressions with nested operations as well. For example

```
> (add (mul 2 3) (div 10 2))
11
```

<br>

## Exit

To exit the REPL, simply press Enter without any expression.

<br>


## To build

```
make
```

## To run the binary

```
./bin/main
```

## To build and run the binary

```
make && ./bin/main
```
<br>
