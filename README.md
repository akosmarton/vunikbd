# Virtual Unicode Keyboard for Linux in Go

## Installing
```bash
go get github.com/akosmarton/vunikbd
```

## Usage
```go
var err error

k, err := vunikbd.NewKeyboard("Virtual Unicode Keyboard", time.Microsecond*250)
if err != nil {
  panic(err)
}

k.TypeString("A ◉ B ⌘ C ◀\n")
```

## Requirements

- Write permission is required on /dev/uinput
