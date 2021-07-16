/**
@auther: kid1999
@date: 2021/7/14 14:25
@desc: log
**/

package model


type Log struct {
	Model
	StudentID		string		`gorm:"type:varchar(50);not null;default:''" json:"StudentID,omitempty"`
	UserName 		string		`gorm:"type:varchar(50);" json:"UserName,omitempty"`
	Type			string		`gorm:"type:varchar(50)" json:"Type,omitempty"`
	AimID			uint		`gorm:"type:integer" json:"AimID,omitempty"`
}


