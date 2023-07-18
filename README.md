# Errdiff

[![Version](https://img.shields.io/github/tag/mrwormhole/errdiff.svg)](https://github.com/mrwormhole/errdiff/tags)
[![CI Build](https://github.com/mrwormhole/errdiff/actions/workflows/test.yaml/badge.svg)](https://github.com/mrwormhole/errdiff/actions/workflows/test.yaml)
[![GoDoc](https://godoc.org/github.com/mrwormhole/errdiff?status.svg)](https://godoc.org/github.com/mrwormhole/errdiff)
[![Report Card](https://goreportcard.com/badge/github.com/mrwormhole/errdiff)](https://goreportcard.com/report/github.com/mrwormhole/errdiff)
[![License](https://img.shields.io/github/license/mrwormhole/errdiff)](https://github.com/mrwormhole/errdiff/blob/master/LICENSE)

This is a fork of h-fam/errdiff, this is created in order to achieve type-safety and better Check() method that can understand wrapped/contained errors.
In the process of doing so, I have removed Substring() method due to issues that cause and interface{} argument that is passed to Check() method. 
Also cleaned up deprecated grpc status code things.

# Usage

### errdiff.Check

The most common option that is used to compare against error

```go
tests := []struct {
  ...
  wantErr error
}{
  // Success
  {...},
  // Failures
  {..., wantErr: errors.New("something failed: EOF")}, // an explicit full error
  {..., wantErr: io.EOF}, // a contained/wrapped error
}
for _, tt := range tests {
  t.Run(tt.name, func(t *testing.T) {
    got, err := fn(...)
    if diff := errdiff.Check(err, tt.wantErr); diff != "" {
      t.Errorf("fn() %s", diff)
    }
  })
}
```

### errdiff.Text

It is used for exact strings to compare against error

```go
tests := []struct {
  ...
  wantErr string
}{
  // Success
  {...},
  // Failures
  {..., wantErr: "something failed: EOF"}, // full text case-sensitive
}
for _, tt := range tests {
  t.Run(tt.name, func(t *testing.T) {
    got, err := fn(...)
    if diff := errdiff.Text(err, tt.wantErr); diff != "" {
      t.Errorf("fn() %s", diff)
    }
  })
}
```

### errdiff.Code

It is used for grpc status codes to compare against error

```go
tests := []struct {
  ...
  wantCode codes.Code
}{
  // Success
  {...},
  // Failures
  {..., wantCode: codes.InvalidArgument}, // grpc status code
}
for _, tt := range tests {
  t.Run(tt.name, func(t *testing.T) {
    got, err := fn(...)
    if diff := errdiff.Code(err, tt.wantCode); diff != "" {
      t.Errorf("fn() %s", diff)
    }
  })
}
```
