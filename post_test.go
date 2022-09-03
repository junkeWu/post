package main

import (
	"encoding/json"
	"testing"
)

func Test_getCsrfToken(t *testing.T) {
	type args struct {
		url   string
		param map[string]int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success case1",
			args: args{
				url:   GetCsrfTokenUrl,
				param: map[string]int{"portal_entrance": 1},
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "failed case1",
			args: args{
				url:   GetCsrfTokenErrUrl,
				param: map[string]int{"portal_entrance": 1},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCsrfToken(tt.args.url, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCsrfToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				return
			}
			if got.Code != tt.want {
				t.Errorf("getCsrfToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_post(t *testing.T) {
	type args struct {
		url   string
		token string
		body  GetPostDataRequest
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "success case1",
			args: args{
				url:   GetPostUrl,
				token: "",
				body: GetPostDataRequest{
					Limit:             1,
					Offset:            0,
					PortalType:        6,
					JobFunctionIdList: nil,
					PortalEntrance:    1,
				},
			},
		},
	}
	token, err := getCsrfToken(GetCsrfTokenUrl, map[string]int{})
	if err != nil {
		t.Errorf("getCsrfToken() error = %v", err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := post(tt.args.url, token.Data.Token, tt.args.body)
			if err != nil {
				t.Errorf("post() error = %v", err)
			}
			var postResp GetPostDataResp
			err = json.Unmarshal([]byte(got), &postResp)
			if err != nil {
				t.Errorf("post() error = %v", err)
			}
			if postResp.Code != 0 {
				t.Errorf("post() got = %v", got)
			}
		})
	}
}
