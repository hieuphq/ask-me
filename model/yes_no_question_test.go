package model

import "testing"

func TestYesNoQuestion_GetQuestion(t *testing.T) {
	type fields struct {
		Content string
		Answers []Answer
		Users   []User
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &YesNoQuestion{
				Content: tt.fields.Content,
				Answers: tt.fields.Answers,
				Users:   tt.fields.Users,
			}
			got, err := q.GetQuestion()
			if (err != nil) != tt.wantErr {
				t.Errorf("YesNoQuestion.GetQuestion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("YesNoQuestion.GetQuestion() = %v, want %v", got, tt.want)
			}
		})
	}
}
