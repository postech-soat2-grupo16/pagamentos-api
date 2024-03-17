package cliente

type Cliente struct {
	ID        uint32 `json:"id"`
	CPF       string `json:"cpf"`
	Email     string `json:"email"`
	Nome      string `json:"nome"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
