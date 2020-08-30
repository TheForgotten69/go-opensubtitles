package opensubtitles

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestInfoService_Formats(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       InfoService
		args    args
		want    *Formats
		want1   *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.s.Formats(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Formats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Formats() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Formats() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestInfoService_Languages(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       InfoService
		args    args
		want    *Languages
		want1   *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.s.Languages(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Languages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Languages() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Languages() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestInfoService_User(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       InfoService
		args    args
		want    *User
		want1   *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.s.User(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("User() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("User() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
