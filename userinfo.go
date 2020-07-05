package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Address  string
	AuthID   string
	AuthSign string
	Client   *http.Client
}

type UserInfo struct {
	conf Config
}

type User struct {
	ID                 uint64    `json:"id"`
	Login              string    `json:"login"`
	FirstName          string    `json:"first_name"`
	MiddleName         string    `json:"middle_name"`
	LastName           string    `json:"last_name"`
	Email              string    `json:"email"`
	IsEmailValid       bool      `json:"is_email_valid"`
	IsActive           bool      `json:"is_active"`
	RegistrationDate   time.Time `json:"registration_date"`
	HasPhoto           bool      `json:"has_photo"`
	Phone              uint64    `json:"phone"`
	PhoneIsChecked     bool      `json:"phone_is_checked"`
	Birthday           time.Time `json:"birthday"`
	BirthdayVisibility int       `json:"birthday_visibility"`
	Gender             string    `json:"gender"`
	EmployeeID         uint64    `json:"employee_id"`
	Country            string    `json:"country"`
	Is2FAuth           bool      `json:"is_2f_auth"`
	QuickRegistration  bool      `json:"quick_registration"`
	Language           string    `json:"language"`
}

func (u *UserInfo) GetUser(uid uint64) (*User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/users/%d", u.conf.Address, uid), nil)
	if err != nil {
		return nil, fmt.Errorf("error while sending request %v", err)
	}

	req.Header.Set("x-auth-id", u.conf.AuthID)
	req.Header.Set("x-auth-sign", u.conf.AuthSign)

	resp, err := u.conf.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while client.Do %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	var user User

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("error while Unmarshal json %v", err)
	}

	return &user, nil
}

func NewUserInfo(conf Config) *UserInfo {
	return &UserInfo{conf: conf}
}
