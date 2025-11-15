package domain

type Bouquet interface {
    Description() string
    Price() float64
}

type BasicBouquet struct {
    Name      string
    BasePrice float64
}

func (b *BasicBouquet) Description() string { return b.Name }
func (b *BasicBouquet) Price() float64      { return b.BasePrice }
