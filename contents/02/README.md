# Introduction
##  How to we run code?

![](02.01.png)

```
go run main.go
```

```
go build main.go
./main
```

![](02.07.png)

## `package main`?

Package == Project == Workspace

### Package?

![](02.02.png)

### Type of package

![](02.03.png)

### Executable

![](02.04.png)

### Reuseable

![](02.05.png)

### Comparision

![](02.06.png)

## import "fmt"

![](02.08.png)

![](02.09.png)

### Standard lib

https://golang.org/pkg

## func?

![](02.10.png)

## How is the main.go file organinzed?

![](02.11.png)

# Quiz

1. What is the purpose of a package in Go?

```
A. To specify the load order of files
B. To group together code with a similar purpose
C. To customize your program for running on MacOS or Windows
```
<!-- B. -->

2. What is the special name we use for a package to tell Go that we want it to be turned into a file that can be executed?

```
A. fmt
B. main
C. pkg
```
<!-- B. -->

3. The one requirement of packages named "main" is that we...

```
A. Export all functions declared within it
B. Import at least one system package
C. Define a function named "main", which is ran automatically when the program runs
```
<!-- C. -->

4. Why do we use "import" statements?
```
A. To give our package access to code written in another package
B. To optimize garbage collection
C. To declare functions we that our package should optimize
```
<!-- A. -->



