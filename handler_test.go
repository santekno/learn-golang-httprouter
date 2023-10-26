package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
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

			assert.Equal(t, tt.want, bodyString)
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

			assert.Equal(t, tt.want, bodyString)
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
			want: "Product 1",
		},
		{
			name: "test get params with product 2",
			args: args{
				params: "2",
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

			assert.Equal(t, tt.want, bodyString)
		})
	}
}

func TestNamedParameterHandler(t *testing.T) {
	type args struct {
		id     string
		itemId string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test get params with product 1",
			args: args{
				id:     "1",
				itemId: "2",
			},
			want: "Product 1 Item 2",
		},
		{
			name: "test get params with product 2",
			args: args{
				id:     "2",
				itemId: "3",
			},
			want: "Product 2 Item 3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost/product/%s/item/%s", tt.args.id, tt.args.itemId), nil)
			recorder := httptest.NewRecorder()
			NamedParameterHandler(recorder, request, httprouter.Params{
				{
					Key:   "id",
					Value: tt.args.id,
				},
				{
					Key:   "itemId",
					Value: tt.args.itemId,
				},
			})

			response := recorder.Result()
			body, _ := io.ReadAll(response.Body)
			bodyString := string(body)

			assert.Equal(t, tt.want, bodyString)
		})
	}
}

func TestCatchAllParameterHandler(t *testing.T) {
	type args struct {
		image string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "get image",
			args: args{
				image: "photo.jpg",
			},
			want: "Image photo.jpg",
		},
		{
			name: "get image with path",
			args: args{
				image: "small/photo.jpg",
			},
			want: "Image small/photo.jpg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost/images/%s", tt.args.image), nil)
			recorder := httptest.NewRecorder()
			CatchAllParameterHandler(recorder, request, httprouter.Params{
				{
					Key:   "image",
					Value: tt.args.image,
				},
			})

			response := recorder.Result()
			body, _ := io.ReadAll(response.Body)
			bodyString := string(body)

			assert.Equal(t, tt.want, bodyString)
		})
	}
}
