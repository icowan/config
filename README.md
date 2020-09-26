# config

## install

```bash
$ go get github.com/icowan/config
```

## use

```go
cfg, err = config.NewConfig(configPath)
if err != nil {
    panic(err)
}

cfg.GetString("mysql", "host")
```