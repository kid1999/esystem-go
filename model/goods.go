/**
@auther: kid1999
@date: 2021/7/13 21:47
@desc: goods
**/

package model


type Goods struct {
	Model
	GoodsName 		string		`gorm:"type:varchar(50);" json:"GoodsName,omitempty"`
	Description		string		`gorm:"type:varchar(250);not null;default:''" json:"Description,omitempty"`
	OwnerStudentID	string		`gorm:"type:varchar(50);not null;" json:"OwnerStudentID,omitempty"`
	Username		string		`gorm:"type:varchar(30);not null;default:''" json:"Username,omitempty"`
	Phone			string		`gorm:"type:varchar(20);not null;default:''" json:"Phone,omitempty"`
	Wx				string		`gorm:"type:varchar(50);default:''" json:"Wx,omitempty"`
	QQ				string		`gorm:"type:varchar(50);default:''" json:"QQ,omitempty"`
	Price			float64		`gorm:"type:decimal;default:0" json:"Price,omitempty"`
	State			string		`gorm:"type:varchar(20);default:''" json:"State,omitempty"`
	ImgUrl			string		`gorm:"type:varchar(250);default:'http://kid1999.top:9000/images/favor.png'" json:"ImgUrl,omitempty"`
}
