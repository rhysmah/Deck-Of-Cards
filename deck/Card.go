package deck

type Card struct {
	Suit  string
	Value int8
}

const (
	Spades  string = "Spade"
	Diamond string = "Diamond"
	Club    string = "Club"
	Hearts  string = "Heart"
)

func New() []Card {

}
