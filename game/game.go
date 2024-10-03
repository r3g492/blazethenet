package game

import rl "github.com/gen2brain/raylib-go/raylib"

func Logic() bool {
	rl.ClearBackground(rl.Black)

	answers := []Answer{
		newAnswer(VerbFears, ObjectFire),
		newAnswer(VerbWeakAgainst, ObjectWater),
		newAnswer(VerbLooksLike, ObjectPointyThings),
	}

	var y int32 = 100
	for _, answer := range answers {
		rl.DrawText(answer.Subject+" "+answer.Verb.String()+" "+answer.Object.String(),
			10,
			y,
			100,
			rl.Blue,
		)
		y += 100
	}

	return true
}
