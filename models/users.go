package models

type Users struct {
	ID			uint32		`json:"id" gorm:"primary_key:auto_increment"`
	Name		string		`json:"name" gorm:"size:255;not null"`
	Email		string		`json:"email" gorm:"size:255;not null;unique"`
	Password	string		`json:"password" gorm:"size:255;not null"`
	Address		string		`json:"address" gorm:"size:255"`
	Skill		string		`json:"skill" gorm:"size:255"`
	Phone		string		`json:"phone" gorm:"size:255"`
	Age			int			`json:"age" gorm:"size:255"`
	IsAvailable	bool		`json:"isAvailable" gorm:"size:255;default:true"`
	AvatarURL	string		`json:"avatarUrl" gorm:"size:255;default:https://avatars.dicebear.com/api/personas/default.svg"`
}