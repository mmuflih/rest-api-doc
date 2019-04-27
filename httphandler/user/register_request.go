package user

/*
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-04-27 08:00
 */

type registerRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (lr registerRequest) GetEmail() string {
	return lr.Email
}

func (lr registerRequest) GetName() string {
	return lr.Name
}

func (lr registerRequest) GetPhone() string {
	return lr.Phone
}

func (lr registerRequest) GetPassword() string {
	return lr.Password
}
