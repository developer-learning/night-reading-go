package userservice

import (
	"ginexamples"
	"ginexamples/mock"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

var r = mock.UserRepository{
	StoreFnInvoked: false,
	StoreFn: func(user *ginexamples.User) error {
		if user.Email == "fail@mail.com" {
			return errors.New("error storing user")
		}
		return nil
	},
	UpdateFnInvoked: false,
	UpdateFn: func(user *ginexamples.User) error {
		return nil
	},
	FindFnInvoked: false,
	FindFn: func(id string) (*ginexamples.User, error) {
		if id == "0" {
			return &ginexamples.User{Model: gorm.Model{ID: 0}}, nil
		}
		return nil, errors.New("not found")
	},
	FindByEmailFnInvoked: false,
	FindByEmailFn: func(email string) (*ginexamples.User, error) {
		if email == "existing@mail.com" {
			return &ginexamples.User{Model: gorm.Model{ID: 0}, PasswordHash: "password"}, nil
		}
		return nil, errors.New("not found")
	},
	FindBySessionIDFnInvoked: false,
	FindBySessionIDFn: func(sessionID string) (*ginexamples.User, error) {
		if sessionID == "sessionID-0-0-0" {
			return &ginexamples.User{Model: gorm.Model{ID: 0}}, nil
		}
		return nil, errors.New("not found")
	},
}
var a = mock.AuthenticatorMock{
	HashFnInvoked: false,
	HashFn: func(password string) (string, error) {
		if password == "longpasswordbuthashfailed" {
			return "", errors.New("Error hashing password")
		}
		return password, nil
	},
	CompareHashFnInvoked: false,
	CompareHashFn: func(hashedPassword string, plainPassword string) error {
		if hashedPassword != plainPassword {
			return errors.New("incorrect password")
		}
		return nil
	},
	SessionIDFnInvoked: false,
	SessionIDFn: func() string {
		return "sessionID-0-0-0"
	},
}

var us = &UserService{&r, &a}

func TestNew(t *testing.T) {
	t.Parallel()
	u := New(&r)
	assert.Equal(t, &r, u.r, "repository does not match")
	assert.NotNil(t, u.a, "no authenticator")
}
func TestUserService_Login(t *testing.T) {
	var testCases = []struct {
		name     string
		email    string
		password string
		error    bool
	}{
		{"new user", "new@mail.com", "password", true},
		{"incorrect password", "existing@mail.com", "pass2", true},
		{"user exists", "existing@mail.com", "password", false},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			defer func() { r.FindByEmailFnInvoked = false }()
			defer func() { r.UpdateFnInvoked = false }()
			defer func() { a.CompareHashFnInvoked = false }()
			defer func() { a.SessionIDFnInvoked = false }()

			user, err := us.Login(v.email, v.password)
			if v.error {
				assert.NotNil(t, err, "did not fail to log in user")
				assert.Empty(t, user, "did not return empty user")
				return
			}
			assert.Nil(t, err, "should not fail to login userService")
			assert.True(t, r.FindByEmailFnInvoked, "FindByEmail not invoked")
			assert.True(t, r.UpdateFnInvoked, "Update not invoked")
			assert.True(t, a.CompareHashFnInvoked, "CompareHash not invoked")
			assert.True(t, a.SessionIDFnInvoked, "SessionID not invoked")
			assert.Equal(t, "sessionID-0-0-0", user.SessionID, "did not update sessionID")
		})
	}
}
