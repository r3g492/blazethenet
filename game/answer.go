package game

type Answer struct {
	Subject string
	Verb    Verb
	Object  Object
}

type Verb int

const (
	VerbFears Verb = iota
	VerbWeakAgainst
	VerbLooksLike
	VerbHas
)

type Object int

const (
	ObjectFire Object = iota
	ObjectPointyThings
	ObjectWater
	ObjectThunder
)

func newAnswer(verb Verb, object Object) Answer {
	return Answer{"The Monster", verb, object}
}
