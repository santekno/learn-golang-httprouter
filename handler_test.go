package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestSampleGetHandler(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "test get",
			want: "Hello Get",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "http://localhost/", nil)
			recorder := httptest.NewRecorder()
			SampleGetHandler(recorder, request, nil)

			response := recorder.Result()
			body, _ := io.ReadAll(response.Body)
			bodyString := string(body)

			if !reflect.DeepEqual(bodyString, tt.want) {
				t.Errorf("response = %v, want %v", bodyString, tt.want)
			}
		})
	}
}

func TestSamplePostHandler(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "test post",
			want: "Hello Post",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "http://localhost/", nil)
			recorder := httptest.NewRecorder()
			SamplePostHandler(recorder, request, nil)

			response := recorder.Result()
			body, _ := io.ReadAll(response.Body)
			bodyString := string(body)

			if !reflect.DeepEqual(bodyString, tt.want) {
				t.Errorf("response = %v, want %v", bodyString, tt.want)
			}
		})
	}
}

func TestGetUsedParamsHandler(t *testing.T) {
	type args struct {
		params string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test get params with product 1",
			args: args{
				params: "1",
			},
			want: "Product 2",
		},
		{
			name: "test get params with product 2",
			args: args{
				params: "1",
			},
			want: "Product 2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost/product/%s", tt.args.params), nil)
			recorder := httptest.NewRecorder()
			GetUsedParamsHandler(recorder, request, httprouter.Params{
				{
					Key:   "id",
					Value: tt.args.params,
				},
			})

			response := recorder.Result()
			body, _ := io.ReadAll(response.Body)
			bodyString := string(body)

			if !reflect.DeepEqual(bodyString, tt.want) {
				t.Errorf("response = %v, want %v", bodyString, tt.want)
			}
		})
	}
}
