<br>

# Assorted-Golang

Example code, Codewars katas and a compiler for the Monkey programming language.

<br>

## Installation

Ben Davis has a nice video walk through on how to [install Go and set up the path variable](https://youtu.be/Q7uh85_i0-M)

<br>

## Syntax primer

"Learn X in Y minutes" has a [good introduction to Go](https://learnxinyminutes.com/docs/go/)

<br>

## To run a single main.go file

```sh
go run main.go
```

<br>

## To create a go.mod file for a project

```sh
go mod init name-of-project
```

<br>

## To compile a project that has a go.mod file

```sh
go build
```

<br>

## To run a project

```
./name-of-project
```

<br>

## To compile and run a project

```sh
go build && ./name-of-project
```

<br>

## Video course

The FreeCodeCamp YouTube channel has a [beginners course](https://youtu.be/un6ZyFkqFKo) to get you started.

<br>

## AI Code Generation

Large Language Models (LLMs) are only effective for me because I don't trust their output.

I always test the code that they generate!

However, given that Go is a mature language, backed by Google, there is a suprising amount of training data available.

<br>

### Function Generation With LLMs

I have had good results with prompts using the following template

```
Write a [name] function in Golang that takes
[name the parameters and their types] and returns
a [type] such that [describe what the function does].
Then show me the code.
```

Another strategy is to provide the LLM with a list of function signatures for a particular utility or feature and ask the LLM to implement the function bodies, one by one, asking for your feedback at each step.

<br>

### Example of Program Generation with Chat-GPT4

I asked the LLM to generate a TUI (Text User Interface) todo app using the gorm ORM and sqlite. 

In this [lengthy conversation with Chat-GPT](https://chat.openai.com/share/b7e57624-f1cd-4d2c-b0dc-fbd2daa7fc9f) I was given the basic functionality I needed.

<br>

### Hacking the context window limit

Although feeding errors back to Chat-GPT4 can yield good results, sometimes a conversation will drag on so long that the LLM "forgets" what you are talking about and the chatbots answers rapidly become useless.

This is because of a fundamental limit built into the LLM. Namely the maximum size of its context window. The size of the context grows as the conversation progresses.

To avoid running up against this limit when you are working on a code generation task, simply delete the chat you are using, start a new one, attach the source code that needs fixing to the chat (the "Advanced Data Analytics" plugin should be enabled), tell the chatbot that there are syntax errors or unexpected behaviour and ask for its help.

This process can be repeated as often as necessary!

<br>
