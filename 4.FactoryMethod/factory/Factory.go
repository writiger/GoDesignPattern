package factory

func InitFactory(factoryName string) interface{} {
	switch factoryName {
	case "IDCard":
		return IDCardFactoryTemplate{}
	default:
		return nil
	}
}
