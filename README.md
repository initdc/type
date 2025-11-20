# type

Bring the Rust [Result Option] type to Golang

## Installation

  `go get github.com/initdc/type`

## Usage

```go
import (
  . "github.com/initdc/type/option"
  . "github.com/initdc/type/result"
)

// Option
s := Some(1)

var s Option[int]
s.None()

n := None[int]()

// Result
var r Result[int, string]
r.Ok(1)

var e Result[int, string]
e.Err("error")


r := Ok[int, string](1)
e := Err[int, string]("error")
```

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/initdc/type.
