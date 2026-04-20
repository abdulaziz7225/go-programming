package visibility

var answerToEverything = 42

func Answer() int {
	return calculateAnswer()
}

func calculateAnswer() int {
	return answerToEverything + AlternativeAnswer
}
