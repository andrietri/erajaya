package schema

type Product struct {
	Base
	Name        string `gorm:"index:idx_name;type:varchar(255);notNull"`
	Description string `gorm:"type:text;Null"`
	Price       int    `gorm:"type:int(11);notNull"`
	Quantity    int    `gorm:"type:int(11);notNull"`
}

func (Product) TableName() string {
	return "products"
}
