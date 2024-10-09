package game

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Logic() bool {
	rl.ClearBackground(rl.Black)
	answers := getAnswers()
	drawAnswers(answers)

	var chosenVerb = VerbDefault
	var chosenObject = ObjectDefault
	rl.DrawText(chosenVerb.String()+" "+chosenObject.String(),
		400,
		10,
		50,
		rl.Brown,
	)

	answerCheck := false

	if answerCheck {
		for answerIdx := range answers {
			if answers[answerIdx].Verb == chosenVerb &&
				answers[answerIdx].Object == chosenObject {
				answers[answerIdx].Found = true
			}
		}
	}

	return true
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
