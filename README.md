# Config

A very simple util to inject variable into env variable, convenient when developing in local environment. You can define variable in yml file, it'll inject it into env variable.

**warning: still in development, use it carefully**

## Getting Started

```golang
func main() {
  ok := config.Load("./config/env.yml") // your config file path, return true if successfully load file, return false if can not find file.

  if os.Getenv("ENV") == "development" {
    fmt.Println("my debug info, don't show it in production.")
  }

  // or use config.Get
  if val, ok := config.Get("ENV"); !ok {
    fmt.Println("ENV doesn't exist in")
  }

  // or use config.MustGet
  // panics if can not find env
  if val := config.MustGet("ENV") {

  }
}
```

## API

### config.Load(filename string) (ok bool)

Load a file and inject it into env, only .yml format is supported

TODO: error handling if wrong format

### config.Get(key string) (val string, exist bool)

Get global env from os

### config.MustGet(key string) (val string)

Get global env from os, panics if env doesn't exist.

## TODO

- [ ] bettor error handling

## License

MIT
