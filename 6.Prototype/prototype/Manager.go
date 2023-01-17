package prototype

import _interface "6.Prototype/interface"

type Manager struct {
	Showcase map[string]_interface.Product
}

func (m *Manager) Register(name string, pro _interface.Product) {
	if _, ok := m.Showcase[name]; ok {
		return
	}
	m.Showcase[name] = pro
}

func (m *Manager) Create(name string) _interface.Product {
	if pro, ok := m.Showcase[name]; ok {
		return pro.Clone()
	}
	return nil
}
