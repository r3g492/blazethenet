package game

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var chosenVerb = VerbDefault
var chosenObject = ObjectDefault
var answers = getAnswers()

func Logic() bool {
	rl.ClearBackground(rl.Black)
	drawAnswers(answers)

	setAnswerFromUserInput()

	rl.DrawText(chosenVerb.String()+" "+chosenObject.String(),
		400,
		10,
		50,
		rl.Brown,
	)

	buttonBounds := rl.Rectangle{X: 50, Y: 800, Width: 200, Height: 50}
	rl.DrawRectangleRec(buttonBounds, rl.Gray)
	rl.DrawText("answerCheck", int32(buttonBounds.X+10), int32(buttonBounds.Y+10), 20, rl.Black)
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), buttonBounds) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		checkAnswer(answers)
	}

	return true
}

func checkAnswer(answers []Answer) {
	for answerIdx := range answers {
		if answers[answerIdx].Verb == chosenVerb &&
			answers[answerIdx].Object == chosenObject {
			answers[answerIdx].Found = true
		}
	}
}

func setAnswerFromUserInput() {
	buttonVerbBounds := rl.Rectangle{X: 50, Y: 600, Width: 200, Height: 50}
	buttonObjectBounds := rl.Rectangle{X: 300, Y: 600, Width: 200, Height: 50}

	rl.DrawRectangleRec(buttonVerbBounds, rl.Gray)
	rl.DrawText("Change Verb", int32(buttonVerbBounds.X+10), int32(buttonVerbBounds.Y+10), 20, rl.Black)

	rl.DrawRectangleRec(buttonObjectBounds, rl.Gray)
	rl.DrawText("Change Object", int32(buttonObjectBounds.X+10), int32(buttonObjectBounds.Y+10), 20, rl.Black)

	if rl.CheckCollisionPointRec(rl.GetMousePosition(), buttonVerbBounds) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		if chosenVerb >= GetLastVerb() {
			chosenVerb = 0
		} else {
			chosenVerb = chosenVerb + 1
		}
	}

	if rl.CheckCollisionPointRec(rl.GetMousePosition(), buttonObjectBounds) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		if chosenObject >= GetLastObject() {
			chosenObject = 0
		} else {
			chosenObject = chosenObject + 1
		}
	}
}

func drawAnswers(answers []Answer) {
	var y int32 = 200
	rl.DrawText(fmt.Sprintf("answer count: %d", len(answers)), 10, 50, 50, rl.Red)
	for _, answer := range answers {
		if !answer.Found {
			continue
		}
		rl.DrawText(answer.Subject+" "+answer.Verb.String()+" "+answer.Object.String(),
			10,
			y,
			50,
			rl.Blue,
		)
		y += 100
	}
}

func getAnswers() []Answer {
	return []Answer{
		newAnswer(
			VerbFears,
			ObjectFire,
			true,
		),
		newAnswer(
			VerbWeakAgainst,
			ObjectWater,
			false,
		),
		newAnswer(
			VerbLooksLike,
			ObjectPointyThings,
			false,
		),
	}
}

func ScenarioEndCondition() bool {
	for answer := range answers {
		if answers[answer].Found == false {
			return false
		}
	}
	return true
}
