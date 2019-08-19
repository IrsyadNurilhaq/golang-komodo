package response

type Tweet struct {
	Id    int    `json:"id"`
	Tweet string `json:"tweet"`
	Email string `json:"email"`
}
