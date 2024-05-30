package repository

import (
	"AuthService/data"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetAll(db *gorm.DB) ([]data.User, error) {
	var users []data.User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// GetByEmail returns one user by email
func GetByEmail(email string, db *gorm.DB) (*data.User, error) {
	var user data.User
	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetOne returns one user by id
func GetOne(id int, db *gorm.DB) (*data.User, error) {
	var user data.User
	result := db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// Update updates one user in the database, using the information
// stored in the receiver u
func Update(user data.User, db *gorm.DB) error {
	var err = db.Model(&user).Updates(data.User{Email: user.Email, FirstName: user.FirstName, LastName: user.LastName, Active: user.Active, UpdatedAt: time.Now()})
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// Delete deletes one user from the database, by User.ID
func Delete(u data.User, db *gorm.DB) error {
	var err = db.Delete(&u)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// DeleteByID deletes one user from the database, by ID
func DeleteByID(id int, db *gorm.DB) error {
	var err = db.Delete(&data.User{}, id)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// Insert inserts a new user into the database, and returns the ID of the newly inserted row
func Insert(user data.User, db *gorm.DB) (int, error) {
	var err = db.Create(&user)
	if err.Error != nil {
		return 0, err.Error
	}
	return user.ID, nil
}

// ResetPassword is the method we will use to change a user's password.
func ResetPassword(u *data.User, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// PasswordMatches uses Go's bcrypt package to compare a user supplied password
func PasswordMatches(password string, plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid password
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
