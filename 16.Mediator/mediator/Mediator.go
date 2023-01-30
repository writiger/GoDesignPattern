package mediator

import "fmt"

type UnitedNations interface {
	ForwardMessage(message string, country Country)
}

type UnitedNationsSecurityCouncil struct {
	USA
	Iraq
}

func (unsc UnitedNationsSecurityCouncil) ForwardMessage(message string, country Country) {
	switch country.(type) {
	case USA:
		unsc.Iraq.GetMessage(message)
	case Iraq:
		unsc.USA.GetMessage(message)
	default:
		fmt.Println("The country is not a member of UNSC")
	}

}
