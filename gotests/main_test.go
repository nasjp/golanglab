package main

import "testing"

func TestHoge(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Hoge()

			if (err != nil) != tt.wantErr {
				t.Errorf("Hoge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("Hoge() = %v, want %v", got, tt.want)
			}
		})
	}
}
