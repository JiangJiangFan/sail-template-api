package utils

import (
	"sail-chat/global"
	"sail-chat/types"

	"gorm.io/gorm"
)

type Regular[T any] struct {
	Data T
}

// Paginate 分页
func Paginate(meta *types.Meta) func(c *gorm.DB) *gorm.DB {
	return func(c *gorm.DB) *gorm.DB {
		if meta.Current == 0 {
			meta.Current = 1
		}
		switch {
		case meta.Size > 100:
			meta.Size = 100
		case meta.Size <= 0:
			meta.Size = 10
		}
		meta.Page = meta.Total / meta.Size
		if meta.Total%meta.Size > 0 {
			meta.Page++
		}
		p := meta.Current
		if p > meta.Page {
			p = meta.Page
		}
		size := meta.Size
		offset := size * (p - 1)
		return c.Offset(int(offset)).Limit(int(size))
	}
}

func (r *Regular[T]) SelectOne(wrapper map[string]interface{}) (t T, e error) {
	var model T
	err := global.App.DB.Model(&model).Where(wrapper).Limit(1).Find(&model).Error
	return model, err
}
