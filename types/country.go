package types

type Country struct {
	PID  uint   `form:"pid,omitempty" json:"pid"`
	Name string `form:"name,omitempty" json:"name"`
}
