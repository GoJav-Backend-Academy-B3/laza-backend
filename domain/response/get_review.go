package response

import "time"

type Review struct {
	Id        uint      `json:"id" gorm:"primarykey"`
	Comment   string    `json:"comment"`
	Rating    float32   `json:"rating"`
	Namauser  string    `json:"nama_user"`
	CreatedAt time.Time `json:"created_at"`
}
