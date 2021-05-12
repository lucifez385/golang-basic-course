# waitGroup

## Promise.all in JavaScript

```javascript
const fetch = require('node-fetch')

async function runAsync() {
    const promise1 = fetch('https://qvault.io')
    const promise2 = fetch('https://github.com')
    const promise3 = fetch('https://gitlab.io')

    await Promise.all([promise1, promise2, promise3]).then(async (values) => {
        ...
    });

    nextProcess();
}

runAsync()
```

## WaitGroup implementation

```go
sync.WaitGroup
```

```go
wg.Add(n)
```

```go
wg.Done()
```

```go
wg.Wait()
```

