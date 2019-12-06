package model

import "testing"

func TestProfile_String(t *testing.T) {
	type fields struct {
		Name      string
		Id        string
		Age       int
		Education string
		Local     string
		Marriage  string
		Height    string
		Wage      string
		WageMin   int
		WageMax   int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Profile{
				Name:      tt.fields.Name,
				Id:        tt.fields.Id,
				Age:       tt.fields.Age,
				Education: tt.fields.Education,
				Local:     tt.fields.Local,
				Marriage:  tt.fields.Marriage,
				Height:    tt.fields.Height,
				Wage:      tt.fields.Wage,
				WageMin:   tt.fields.WageMin,
				WageMax:   tt.fields.WageMax,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}