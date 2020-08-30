package opensubtitles

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestFindService_Feature(t *testing.T) {
	type args struct {
		ctx context.Context
		opt *FeatureOptions
	}
	tests := []struct {
		name        string
		s           FindService
		args        args
		wantFeature *Feature
		wantResp    *http.Response
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFeature, gotResp, err := tt.s.Feature(tt.args.ctx, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("Feature() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFeature, tt.wantFeature) {
				t.Errorf("Feature() gotFeature = %v, want %v", gotFeature, tt.wantFeature)
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Feature() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestFindService_Find(t *testing.T) {
	type args struct {
		ctx context.Context
		opt *FindOptions
	}
	tests := []struct {
		name        string
		s           FindService
		args        args
		wantFeature *Feature
		wantResp    *http.Response
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFeature, gotResp, err := tt.s.Find(tt.args.ctx, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFeature, tt.wantFeature) {
				t.Errorf("Find() gotFeature = %v, want %v", gotFeature, tt.wantFeature)
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Find() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
