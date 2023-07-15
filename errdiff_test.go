package errdiff

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type elist []string

func (e elist) Error() string { return strings.Join(e, ", ") }

func TestCheck(t *testing.T) {
	elist1 := elist{"error1a", "error1b"}
	elist2 := elist{"error2a", "error2b"}

	tests := []struct {
		name       string
		got        error
		want       error
		wantResult string
	}{
		{
			name: "empty",
		},
		{
			name: "exact same error",
			got:  io.EOF,
			want: io.EOF,
		},
		{
			name: "contained same error",
			got:  fmt.Errorf("something: %w", io.EOF),
			want: io.EOF,
		},
		{
			name: "different errors with same string",
			want: errors.New("an error"),
			got:  errors.New("an error"),
		},
		{
			name:       "nil want",
			got:        io.EOF,
			wantResult: "got err=EOF, want err=nil",
		},
		{
			name:       "nil got",
			want:       io.EOF,
			wantResult: "got err=nil, want err=EOF",
		},
		{
			name:       "different error",
			want:       errors.New("this error"),
			got:        errors.New("that error"),
			wantResult: "got err=that error, want err=this error",
		},
		{
			name:       "unexpected errlist",
			got:        elist1,
			wantResult: "got err=error1a, error1b, want err=nil",
		},
		{
			name:       "missing errlist",
			want:       elist1,
			wantResult: "got err=nil, want err=error1a, error1b",
		},
		{
			name: "correct errlist",
			got:  elist1,
			want: elist1,
		},
		{
			name:       "wrong errlist",
			got:        elist1,
			want:       elist2,
			wantResult: "got err=error1a, error1b, want err=error2a, error2b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Check(tt.got, tt.want); gotResult != tt.wantResult {
				t.Errorf("Check(%v, %v): gotResult=%q wantResult=%q", tt.got, tt.want, gotResult, tt.wantResult)
			}
		})
	}
}

func TestText(t *testing.T) {
	tests := []struct {
		name       string
		got        error
		want       string
		wantResult string
	}{
		{
			name: "empty error",
		},
		{
			name:       "match",
			got:        errors.New("abc"),
			want:       "abc",
			wantResult: "",
		},
		{
			name:       "message no match",
			got:        errors.New("ab"),
			want:       "abc",
			wantResult: "got err=ab, want err=abc"},
		{
			name:       "want nil",
			got:        errors.New("ab"),
			wantResult: "got err=ab, want err=nil",
		},
		{
			name:       "want nil got message",
			got:        errors.New(""),
			wantResult: "got err=, want err=nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Text(tt.got, tt.want); gotResult != tt.wantResult {
				t.Errorf("Text(%v, %v): gotResult=%q wantResult=%q", tt.got, tt.want, gotResult, tt.wantResult)
			}
		})
	}
}

func TestCode(t *testing.T) {
	tests := []struct {
		name       string
		got        error
		want       codes.Code
		wantResult string
	}{
		{
			name: "empty message",
		},
		{
			name:       "Unimplemented match",
			got:        status.Errorf(codes.Unimplemented, ""),
			want:       codes.Unimplemented,
			wantResult: "",
		},
		{
			name:       "code no match",
			got:        status.Errorf(codes.Unimplemented, ""),
			want:       codes.InvalidArgument,
			wantResult: "got err=rpc error: code = Unimplemented desc = , want code=InvalidArgument",
		},
		{
			name:       "nil match",
			got:        status.Errorf(codes.Unimplemented, ""),
			wantResult: "got err=rpc error: code = Unimplemented desc = , want code=OK",
		},
		{
			name:       "no code",
			got:        errors.New("other"),
			want:       codes.InvalidArgument,
			wantResult: "got err=other, want code=InvalidArgument",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Code(tt.got, tt.want); gotResult != tt.wantResult {
				t.Errorf("CodeCompare(): gotResult=%q wantResult=%q", gotResult, tt.wantResult)
			}
		})
	}
}
