package main

import (
	"github.com/hieuphq/ask-me/model"
	"github.com/k0kubun/pp"
)

func main() {
	var q model.Question

	q = &model.YesNoQuestion{
		Content: "Cau hoi gi?",
		Answers: []model.Answer{
			model.Answer("Yes"),
			model.Answer("No"),
		},
	}

	q.UpVote(model.User{ID: 1, Name: "Hieu"})
	q.UpVote(model.User{ID: 2, Name: "Hieu NM"})

	rs, err := q.GetRating()
	if err != nil {
		print(err)
	}

	pp.Print(rs)
}
