package golang_quiz_app

type Quiz struct {
	Questions []Question
}

type Question interface {
	GetText() string
}

type TrueFalseQuestion struct {
	Text   string
	Answer bool
}

type SingleChoiceQuestion struct {
	Text            string
	PossibleAnswers []string
	Answer          string
}

type MultipleChoiceQuestion struct {
	Text            string
	PossibleAnswers []string
	Answers         []string
}

type FillInTheBlankQuestion struct {
	Text   string
	Answer string
}
