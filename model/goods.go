/**
@auther: kid1999
@date: 2021/7/13 21:47
@desc: goods
**/

package model


type Goods struct {
	Model
	GoodsName 		string		`gorm:"type:varchar(50);unique_index" json:"goods_name,omitempty"`
	Description		string		`gorm:"type:varchar(250);not null;default:''" json:"password,omitempty"`
	OwnerStudentID	string		`gorm:"type:varchar(50);not null;default:''" json:"owner_student_id,omitempty"`
	UserName		string		`gorm:"type:varchar(30);not null;default:''" json:"username,omitempty"`
	Phone			string		`gorm:"type:varchar(20);not null;default:''" json:"phone,omitempty"`
	Wx				string		`gorm:"type:varchar(50);default:''" json:"wx,omitempty"`
	QQ				string		`gorm:"type:varchar(50);default:''" json:"qq,omitempty"`
	Price			float64		`gorm:"type:decimal;;default:0" json:"price,omitempty"`
	State			string		`gorm:"type:varchar(20);default:'applying'" json:"state,omitempty"`
}
