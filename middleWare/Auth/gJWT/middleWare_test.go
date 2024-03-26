package gJWT

import (
	"testing"
	"time"
)

func TestAuthToken(t *testing.T) {
	New(
		WithExpireTime(2*time.Hour+30*time.Minute), // signed token will expired in 2 hours and 30 minutes
		WithKey("123"),
	)
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestAuthToken",
			args: args{
				"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySW5mbyI6eyJ1c2VybmFtZSI6ImRlbW9cdTAwMDAiLCJwYXNzd29yZCI6ImRlbW9cdTAwMDAiLCJlbWFpbCI6IjEyM0BxcS5jb20ifSwiZXhwIjoxNzExNDQ2MzE5LCJpYXQiOjE3MTE0MzczMTl9.mhqsYpBpO_zm9UCYOKF5Q9KtF8qpJKNU9cP8qEqdDiY",
			},
			wantErr: false,
		},
		{
			name: "WrongToken Input",
			args: args{
				"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySW5mbyI6eyJ1c2VybmFtZSI6ImRlbW9cdTAwMDMiLCJwYXNzd29yZCI6ImRlbW9cdTAwMDMiLCJlbWFpbCI6IjEyM0BxcS5jb20ifSwiZXhwIjoxNzExNDQ2NTIyLCJpYXQiOjE3MTE0Mzc1MjJ9.J4Yob9ZKgiqOrfQHV-POm-Zc4Rsjcg7ICuHLHiCIR3Q",
			},
			wantErr: true,
		},
		{
			name: "EmptyToken",
			args: args{
				"",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AuthToken(tt.args.tokenString); (err != nil) != tt.wantErr {
				t.Errorf("AuthToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSignedToken(t *testing.T) {
	type args struct {
		arg interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "TestSignedToken",
			args: args{
				arg: DemoUser{
					Username: "demo",
					Password: "demo",
					Email:    "123@qq.com",
				},
			},
			wantErr: false,
		},
		{
			name: "TestSignedEmptyInfoToToken",
			args: args{
				arg: DemoUser{},
			},
			wantErr: true, // should not be empty, if signed a empty struct, it will return error
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SignedToken(tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignedToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) == 0 {
				t.Errorf("SignedToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
