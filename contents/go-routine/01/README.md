# Go routine / Channel

## Website uptime checker

![](01.01.png)

```go
links := []string{
    "https://google.com",
    "https://facebook.com",
    "https://stackoverflow.com",
    "https://golang.org",
    "https://amazon.com",
}
```

```go
func checkLink(link string) {
	...
}
```

## Synchronous process

![](01.02.png)

![](01.03.png)

## Parallel

![](01.04.png)

## Go routine?

![](01.05.png)

![](01.06.png)

![](01.07.png)

![](01.08.png)

## Go routine syntax

![](01.09.png)

## Theory of go routine

![](01.10.png)

![](01.11.png)

## Concurrency vs Parallelism

![](01.12.png)

![](01.13.png)

![](01.14.png)

## Run program with go

```
go run main.go
```

## Channel?

![](01.15.png)

![](01.16.png)

![](01.17.png)

## Channel implementation

![](01.18.png)

![](01.19.png)

```
https://golang.org is up :)
is up :)
```

## Blocking channels

![](01.20.png)

```
https://golang.org is up :)
is up :)
https://google.com is up :)
is up :)
```

![](01.21.png)

## Receiving message

```go
for i := 0; i < len(links); i++ {
    fmt.Println(<-c)
}
```

![](01.22.png)

## Repeting routine

```go
for {
    go checkLink(<-c, c)
}
```

```go
for l := range c {
    go checkLink(l, c)
}
```

## Sleeping routine

```go
time.Sleep(5 * time.Second)
```