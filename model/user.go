/**
@auther: kid1999
@date: 2021/7/13 21:44
@desc: user
**/

package model


type User struct {
	Model
	StudentID 		string		`gorm:"type:varchar(50);unique_index" json:"StudentID,omitempty"`
	Username 		string		`gorm:"type:varchar(50)" json:"Username,omitempty"`
	Password		string		`gorm:"type:varchar(64);not null;default:''" json:"Password,omitempty"`
	Role			string		`gorm:"type:varchar(50);not null;default:'user'" json:"Role,omitempty"`
	Phone			string		`gorm:"type:varchar(20);not null;default:''" json:"Phone,omitempty"`
	Wx				string		`gorm:"type:varchar(50);default:''" json:"Wx,omitempty"`
	QQ				string		`gorm:"type:varchar(50);default:''" json:"QQ,omitempty"`
	Email			string		`gorm:"type:varchar(50);default:''" json:"Email,omitempty"`
	State			string		`gorm:"type:varchar(20);default:''" json:"State,omitempty"`
}

