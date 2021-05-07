# Interface

## What's interface

![](01.01.png)

![](01.02.png)

## Bot project

![](01.03.png)

![](01.04.png)

![](01.05.png)

```go
package main

type englishBot struct {
}

type spanishBot struct {
}

func main() {

}

func (eb englishBot) getGreeting() string {
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}
```

```go
type bot interface {
	getGreeting() string
}
```

```go
func printGreeting(b bot) {
    ...
}
```

![](01.06.png)

## Rule of Interface

![](01.07.png)

![](01.08.png)

## Concreate vs Interface

![](01.09.png)

**Concreate** Can create a value by direct to type

**Interface** Can create a value by indirect to type 

## Note for Interface

![](01.10.png)