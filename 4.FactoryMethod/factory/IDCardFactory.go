package factory

import "fmt"

// IDCardFactory 子类
type IDCardFactory struct {
	IDCardFactoryTemplate
}

// IDCardFactoryTemplate 父类
type IDCardFactoryTemplate struct {
	Owners []string
}

func (idcft *IDCardFactoryTemplate) CreateProduct(name string) IDCard {
	fmt.Println("制作" + name + "的IDCard")
	return IDCard{Owner: name}
}

func (idcft *IDCardFactoryTemplate) RegisterProduct(name string) {
	idcft.Owners = append(idcft.Owners, name)
}

// Create 具体处理
func (idcf *IDCardFactory) Create(name string) IDCard {
	idcf.RegisterProduct(name)
	return idcf.CreateProduct(name)
}

func InitIDCardFactory() IDCardFactory {
	return IDCardFactory{InitFactory("IDCard").(IDCardFactoryTemplate)}
}
