package models

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func (u *User) UpdateEmail(newEmail string) {
    u.Email = newEmail
}