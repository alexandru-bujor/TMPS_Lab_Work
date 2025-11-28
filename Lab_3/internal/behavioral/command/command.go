package command

import "TMPS_Lab_Work/Lab_3/internal/domain"

type Command interface {
	Execute()
	Undo()
}

type AddFlower struct {
	Bouquet *domain.Bouquet
	Flower  string
}

func (c *AddFlower) Execute() { c.Bouquet.Flowers = append(c.Bouquet.Flowers, c.Flower) }
func (c *AddFlower) Undo()    { c.Bouquet.Flowers = c.Bouquet.Flowers[:len(c.Bouquet.Flowers)-1] }

type Manager struct{ History []Command }

func (m *Manager) Run(cmd Command) { cmd.Execute(); m.History = append(m.History, cmd) }

func (m *Manager) Undo() {
	if len(m.History) == 0 {
		return
	}
	last := m.History[len(m.History)-1]
	last.Undo()
	m.History = m.History[:len(m.History)-1]
}
