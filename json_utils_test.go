package caliber

import (
	"testing"
)

func TestJsonToStructPtr(t *testing.T) {
	// Define a test struct type
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	tests := []struct {
		name      string
		jsonStr   string
		wantErr   bool
		expected  *Person
		expectNil bool
	}{
		{
			name:    "valid json",
			jsonStr: `{"name": "Alice", "age": 30}`,
			wantErr: false,
			expected: &Person{
				Name: "Alice",
				Age:  30,
			},
		},
		{
			name:      "invalid json",
			jsonStr:   `{"name": "Bob", "age": "thirty"}`, // age should be number
			wantErr:   true,
			expectNil: true,
		},
		{
			name:      "empty json",
			jsonStr:   `{}`,
			wantErr:   false,
			expected:  &Person{},
			expectNil: false,
		},
		{
			name:      "malformed json",
			jsonStr:   `{"name": "Charlie", "age": 40`, // missing closing brace
			wantErr:   true,
			expectNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JsonToStructPtr[Person](tt.jsonStr)

			if (err != nil) != tt.wantErr {
				t.Errorf("JsonToStructPtr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.expectNil {
				if got != nil {
					t.Errorf("JsonToStructPtr() = %v, expected nil", got)
				}
				return
			}

			if got == nil {
				t.Error("JsonToStructPtr() returned nil when expecting non-nil")
				return
			}

			if !tt.wantErr {
				if got.Name != tt.expected.Name || got.Age != tt.expected.Age {
					t.Errorf("JsonToStructPtr() = %+v, expected %+v", got, tt.expected)
				}
			}
		})
	}
}

func TestJsonToStructPtrWithDifferentType(t *testing.T) {
	type Product struct {
		ID    string  `json:"id"`
		Price float64 `json:"price"`
	}

	jsonStr := `{"id": "123", "price": 19.99}`
	got, err := JsonToStructPtr[Product](jsonStr)
	if err != nil {
		t.Errorf("JsonToStructPtr() unexpected error: %v", err)
		return
	}

	expected := &Product{
		ID:    "123",
		Price: 19.99,
	}

	if got.ID != expected.ID || got.Price != expected.Price {
		t.Errorf("JsonToStructPtr() = %+v, expected %+v", got, expected)
	}
}
