package address

import "testing"

func TestTypeOfAddress(t *testing.T) {
	type Address struct {
		Address string
		Type    string
	}
	tests := []Address{
		{"Rua abc", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dale", "Rodovia"},
		{"Praça da ce", "Invalid Type of Address"},
		{"RUA TESTE", "Rua"},
		{"", "Invalid Type of Address"},
	}
	for _, stub := range tests {
		receiveType := TypeOfAddress(stub.Address)
		if receiveType != stub.Type {
			t.Errorf("The type receive %s is differ the expected %s", receiveType, stub.Address)
		}

	}
}
