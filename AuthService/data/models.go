package data

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User is the structure which holds one user from the database.
type User struct {
	ID        int       `gorm:"primaryKey" json:"id" autoIncrement:"true"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Password  string    `json:"password,omitempty"`
	Active    int       `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetAll returns a slice of all users, sorted by last name
func (u *User) GetAll(db *gorm.DB) ([]User, error) {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil

}

// GetByEmail returns one user by email
func (u *User) GetByEmail(email string, db *gorm.DB) (*User, error) {
	var user User
	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetOne returns one user by id
func (u *User) GetOne(id int, db *gorm.DB) (*User, error) {
	var user User
	result := db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// Update updates one user in the database, using the information
// stored in the receiver u
func (u *User) Update(user User, db *gorm.DB) error {
	var err = db.Model(&user).Updates(User{Email: user.Email, FirstName: user.FirstName, LastName: user.LastName, Active: user.Active, UpdatedAt: time.Now()})
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// Delete deletes one user from the database, by User.ID
func (u *User) Delete(db *gorm.DB) error {
	var err = db.Delete(&u)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// DeleteByID deletes one user from the database, by ID
func (u *User) DeleteByID(id int, db *gorm.DB) error {
	var err = db.Delete(&User{}, id)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// Insert inserts a new user into the database, and returns the ID of the newly inserted row
func (u *User) Insert(user User, db *gorm.DB) (int, error) {
	var err = db.Create(&user)
	if err.Error != nil {
		return 0, err.Error
	}
	return user.ID, nil
}

// ResetPassword is the method we will use to change a user's password.
func (u *User) ResetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// PasswordMatches uses Go's bcrypt package to compare a user supplied password
func (u *User) PasswordMatches(plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))
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
