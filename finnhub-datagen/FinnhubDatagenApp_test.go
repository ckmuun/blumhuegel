package main

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"reflect"
	"testing"
)

func Test_getFinnhubAuth(t *testing.T) {
	tests := []struct {
		name  string
		want  error
		want1 context.Context
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getFinnhubAuth()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFinnhubAuth() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("getFinnhubAuth() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getProducer(t *testing.T) {
	tests := []struct {
		name string
		want pulsar.Producer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getProducer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getProducer() = %v, want %v", got, tt.want)
			}
		})
	}
}
