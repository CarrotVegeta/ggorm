package ggorm

import "gorm.io/gorm"

type Option func(db *gorm.DB)

type CRUD struct {
	*gorm.DB
	Query string
	Args  []any
}

func NewCRUD(db *gorm.DB) *CRUD {
	return &CRUD{}
}
func (c *CRUD) Raw(query string, args ...any) *CRUD {
	c.DB.Raw(query, args...)
	return c
}
func (c *CRUD) Find(res any) *CRUD {
	c.DB.Find(&res)
	return c
}
func (c *CRUD) Error() error {
	return c.DB.Error
}

type FieldColumn struct {
	Field string
	Value any
}

func (c *CRUD) Update(col FieldColumn) *CRUD {
	c.DB.Update(col.Field, col.Value)
	return c
}
func (c *CRUD) Session() *CRUD {
	c.DB = c.DB.Session(&gorm.Session{})
	return c
}
