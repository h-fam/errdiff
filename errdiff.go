// Package errdiff makes it easy to compare errors(by error, by string or by grpc code) in tests
//
// Similar in intended usage to cmp.Diff but for errors,
// particularly in table-driven tests.
//
// Example usage:
//
// The function Check is used mainly for the existence of an error
// alternatively Text(exact string matches) or Code(gRPC status codes) also can be used:
//
//	tests := []struct {
//		...
//		wantErr error
//	}{
//		// Success
//		{...},
//		// Failures
//		{..., wantErr: errors.New("my expected error string")}, // An explicit full error(case-sensitive)
//		{..., wantErr: io.EOF}, // A contained/wrapped error (sentimental error)
//	}
//	for _, c := range testCases {
//		got, err := fn(...)
//		if diff := errdiff.Check(err, c.wantErr); diff != "" {
//			t.Errorf("fn() %v", diff)
//			continue
//		}
//		...
package errdiff

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Check returns a message describing the difference between got error and want error.
func Check(got error, want error) string {
	if want == nil && got == nil {
		return ""
	}
	if want == nil && got != nil {
		return fmt.Sprintf("got err=%s, want err=nil", got.Error())
	}
	if want != nil && got == nil {
		return fmt.Sprintf("got err=nil, want err=%s", want.Error())
	}
	if got.Error() == want.Error() || errors.Is(got, want) {
		return ""
	}
	return fmt.Sprintf("got err=%s, want err=%s", got.Error(), want.Error())
}

// Text returns a message describing the difference between got error string and want string
// and it performs case-sensitive check.
func Text(got error, want string) string {
	if want == "" && got == nil {
		return ""
	}
	if want == "" && got != nil {
		return fmt.Sprintf("got err=%s, want err=nil", got.Error())
	}
	if want != "" && got == nil {
		return fmt.Sprintf("got err=nil, want err=%s", want)
	}
	if got.Error() == want {
		return ""
	}
	return fmt.Sprintf("got err=%s, want err=%s", got.Error(), want)
}

// Code returns a message describing the difference between the error's code
// and the desired codes. want code=codes.OK indicates that no error is wanted.
func Code(got error, want codes.Code) string {
	if want == codes.OK && got == nil {
		return ""
	}
	if want == codes.OK && got != nil {
		return fmt.Sprintf("got err=%s, want code=%s", got.Error(), want.String())
	}
	if want != codes.OK && got == nil {
		return fmt.Sprintf("got err=nil, want code=%s", want.String())
	}
	if status.Code(got) == want {
		return ""
	}
	return fmt.Sprintf("got err=%v, want code=%v", got, want.String())
}
