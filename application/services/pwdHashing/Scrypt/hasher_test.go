package Scrypt

import (
	"github.com/Thenecromance/OurStories/application/services/pwdHashing"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want pwdHashing.PwdHasher
	}{
		{
			name: "TestNew",
			want: &scryptor{
				cfg: defaultConfig(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scryptor_Hash(t *testing.T) {
	type fields struct {
		cfg *Setting
	}
	type args struct {
		password string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantKey  string
		wantSalt string
	}{
		{
			name: "TestHash1",
			fields: fields{
				cfg: defaultConfig(),
			},
			args: args{
				password: "password",
			},
			wantKey:  "key",
			wantSalt: "salt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &scryptor{
				cfg: tt.fields.cfg,
			}
			gotKey, gotSalt := s.Hash(tt.args.password)
			if gotKey != tt.wantKey {
				t.Errorf("Hash() gotKey = %v, want %v", gotKey, tt.wantKey)
			}
			if gotSalt != tt.wantSalt {
				t.Errorf("Hash() gotSalt = %v, want %v", gotSalt, tt.wantSalt)
			}
		})
	}
}

func Test_scryptor_Verify(t *testing.T) {
	type fields struct {
		cfg *Setting
	}
	type args struct {
		password string
		hash     string
		salt     string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &scryptor{
				cfg: tt.fields.cfg,
			}
			if got := s.Verify(tt.args.password, tt.args.hash, tt.args.salt); got != tt.want {
				t.Errorf("Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scryptor_generateKeyWithSalt(t *testing.T) {
	type fields struct {
		cfg *Setting
	}
	type args struct {
		password []byte
		salt     []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantKey []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &scryptor{
				cfg: tt.fields.cfg,
			}
			if gotKey := s.generateKeyWithSalt(tt.args.password, tt.args.salt); !reflect.DeepEqual(gotKey, tt.wantKey) {
				t.Errorf("generateKeyWithSalt() = %v, want %v", gotKey, tt.wantKey)
			}
		})
	}
}

func Test_scryptor_randomSalt(t *testing.T) {
	type fields struct {
		cfg *Setting
	}
	tests := []struct {
		name     string
		fields   fields
		wantSalt []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &scryptor{
				cfg: tt.fields.cfg,
			}
			if gotSalt := s.randomSalt(); !reflect.DeepEqual(gotSalt, tt.wantSalt) {
				t.Errorf("randomSalt() = %v, want %v", gotSalt, tt.wantSalt)
			}
		})
	}
}
