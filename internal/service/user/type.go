package user

type UserDTO struct {
	Id           uint64 `gorm:"primary_key" json:"id"`
	Pid          uint64 `json:"pid"`
	PhoneNumber  string `json:"phone_number"`
	Name         string `json:"name"`
	Balance      int64  `json:"balance"`
	HeadImg      string `json:"head_img"`
	Addr         string `json:"addr"`
	IdCardName   string `json:"id_card_name"`
	IdCardNumber string `json:"id_card_number"`
	IdCardInfo   string `json:"id_card_info"`
	CreateTime   string `json:"create_time"`
}
