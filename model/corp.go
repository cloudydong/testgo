package model

// Corp 회사종목코드와 회사이름을 갖는다
type Corp struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// TableName 정확한 문자열 테이블 이름인 corp을 지정한다. 관례에 의해 db.Find(&corp) 호출하면 corps 테이블을 찾는것을 방지한다
func (Corp) TableName() string {
	return "corp"
}
