package game

type Answer struct {
	Subject string
	Verb    Verb
	Object  Object
	Found   bool
}

type Verb int

const (
	VerbDefault Verb = iota
	VerbFears
	VerbWeakAgainst
	VerbLooksLike
	VerbHas
)

func (v Verb) String() string {
	return [...]string{"null", "fears", "is weak against", "looks like", "has"}[v]
}

type Object int

const (
	ObjectDefault Object = iota
	ObjectFire
	ObjectPointyThings
	ObjectWater
	ObjectThunder
)

func (o Object) String() string {
	return [...]string{"null", "fire", "pointy things", "water", "thunder"}[o]
}

func newAnswer(verb Verb, object Object, found bool) Answer {
	return Answer{"The Monster", verb, object, found}
}
