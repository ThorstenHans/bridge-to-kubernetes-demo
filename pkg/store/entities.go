package store

type Product struct {
	Id        int     `yaml:"id"`
	Name      string  `yaml:"name"`
	Price     float32 `yaml:"price"`
	IsInStock bool    `yaml:"isInStock"`
}
