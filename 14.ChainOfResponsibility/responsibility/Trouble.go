package responsibility

import "strconv"

type Trouble struct {
	number int
}

func InitTrouble(n int) Trouble {
	return Trouble{n}
}

func (t Trouble) GetNumber() int {
	return t.number
}

func (t Trouble) ToString() string {
	return "[Trouble " + strconv.Itoa(t.number) + " ]"
}
