package question

import (
)

//  A questionare is compose of an slice of questions
//  Question string
//  Options  []string
//  Answer   string
type Question struct {
    Question string
    Options  []string
    Answer   string
}

type Questionare []Question
