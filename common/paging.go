package common

type Paging struct {
	Page  int `form:"page" json:"page"`
	Size  int `form:"size" json:"size"`
	Total int `form:"total" json:"total"`
}

func (p *Paging) Process() {
	if p.Size < 1 || p.Size > 100 {
		p.Size = 10
	}

	if p.Page < 1 {
		p.Page = 1
	}
}
