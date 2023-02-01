package game

type Memento struct {
	money  int
	fruits []string
}

func InitMemento(money int) Memento {
	return Memento{
		money:  money,
		fruits: make([]string, 0),
	}
}

func (m *Memento) AddFruit(fruit string) {
	m.fruits = append(m.fruits, fruit)
}

func (m *Memento) GetFruit() []string {
	return m.fruits
}

func (m *Memento) GetMoney() int {
	return m.money
}
