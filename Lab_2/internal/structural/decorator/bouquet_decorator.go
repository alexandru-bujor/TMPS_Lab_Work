package decorator

import (
	"TMPS_Lab_Work/Lab_2/internal/domain"
)

type WithRibbon struct{ Wrapped domain.Bouquet }

func (d *WithRibbon) Description() string { return d.Wrapped.Description() + ", with ribbon" }
func (d *WithRibbon) Price() float64      { return d.Wrapped.Price() + 3.5 }

type WithCard struct{ Wrapped domain.Bouquet }

func (d *WithCard) Description() string { return d.Wrapped.Description() + ", with card" }
func (d *WithCard) Price() float64      { return d.Wrapped.Price() + 5 }

type WithVase struct{ Wrapped domain.Bouquet }

func (d *WithVase) Description() string { return d.Wrapped.Description() + ", with vase" }
func (d *WithVase) Price() float64      { return d.Wrapped.Price() + 25 }
