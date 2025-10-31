package builder

import "example.com/creational/internal/domain"

type BouquetBuilder struct {
    name    string
    flowers []string
}

func NewBouquetBuilder() *BouquetBuilder {
    return &BouquetBuilder{}
}

func (b *BouquetBuilder) SetName(name string) *BouquetBuilder {
    b.name = name
    return b
}

func (b *BouquetBuilder) AddFlower(f string) *BouquetBuilder {
    b.flowers = append(b.flowers, f)
    return b
}

func (b *BouquetBuilder) Build() domain.Bouquet {
    return domain.Bouquet{Name: b.name, Flowers: b.flowers}
}
