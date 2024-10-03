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

func (v Verb) String() string {
	return [...]string{"fears", "is weak against", "looks like", "has"}[v]
}

type Object int

const (
	ObjectFire Object = iota
	ObjectPointyThings
	ObjectWater
	ObjectThunder
)

func (o Object) String() string {
	return [...]string{"fire", "pointy things", "water", "thunder"}[o]
}

func newAnswer(verb Verb, object Object) Answer {
	return Answer{"The Monster", verb, object}
}
