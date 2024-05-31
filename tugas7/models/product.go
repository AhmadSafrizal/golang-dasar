package model

type Product struct {
	ID          int    `json:"id" gorm:"column:id"`
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
	Price       int    `json:"price" gorm:"column:price"`
	UserID      uint64 `json:"user_id" gorm:"column:user_id"`
	UserDetail  *User  `json:"user_detail" gorm:"foreignKey:UserID"`
}
