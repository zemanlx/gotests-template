# gotests-template

Opinionated template for generation of Go tests using
[gotests](https://github.com/cweill/gotests) utility which is natively
integrated to [VS Code Go
extension](https://marketplace.visualstudio.com/items?itemName=golang.go) and
Goland.

## Features

- Uses [go-cmp](github.com/google/go-cmp/cmp) instead of `reflect.DeepEqual`
  for non-basic types
- `wantErr`is type of `error` which is enabling use of `errors.Is` in assert
  part of tests
- Compatible with all `golangci-lint` linters (whitespace, variable names, etc.)

## Usage

Clone this repository and refer it in VS Code settings:

```json
  "go.generateTestsFlags": [
    "-template_dir=/path/to/this/repo/gotests-template/cmp-linted"
  ]
```

For Goland you may need to use *external tool* settings to pass `-template_dir`
option as kindly described in [gotests issue
161](https://github.com/cweill/gotests/issues/161#issuecomment-848329955).
