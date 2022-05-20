package main

import (
	"reflect"
	"testing"
)

func Test_getResponse(t *testing.T) {
	type args struct {
		commit      string
		pipelineURL string
		env         string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				commit:      "some hash",
				pipelineURL: "some url",
				env:         "dev",
			},
			want:    []byte(`{"commit":"some hash","pipeline_url":"some url","env":"dev"}`),
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				commit:      "",
				pipelineURL: "",
				env:         "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getResponse(tt.args.commit, tt.args.pipelineURL, tt.args.env)
			if (err != nil) != tt.wantErr {
				t.Errorf("getResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
