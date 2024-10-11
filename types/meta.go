package types

type Meta struct {
	Id      string `form:"id,omitempty"    json:"id,omitempty"`         // 请求id
	Size    int64  `form:"size,omitempty"  json:"size,omitempty"`       // 每页大小
	Page    int64  `form:"page,omitempty"  json:"page,omitempty"`       // 总页数
	Current int64  `form:"current,omitempty"  json:"current,omitempty"` // 当前页数
	Order   string `form:"order,omitempty" json:"order,omitempty"`      // 排序方式
	Total   int64  `form:"total,omitempty" json:"total,omitempty"`      // 总数
	Value   string `form:"value,omitempty" json:"value,omitempty"`      // 查询值
}

// func (err API) Error() string {
// 	return err.Message
// }