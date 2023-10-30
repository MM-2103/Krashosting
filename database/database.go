package database

import (
	"fmt"
	"log"
	"os"

	// Bcrypt

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	dsn string
)

func LoadEnv() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn = fmt.Sprintf(
		"%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("db_USER"),
		os.Getenv("db_PASS"),
		os.Getenv("db_DATABASE"),
	)
}

func Init() {
	LoadEnv()

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

type (
	User struct {
		gorm.Model
		Username    string    `gorm:"unique,not null,type:varchar(100)"`
		PhoneNumber uint      `gorm:"unique,not null,type:varchar(100)"`
		Email       string    `gorm:"unique,not null,type:varchar(255)"`
		Password    string    `gorm:"not null,type:varchar(500)"`
		Articles    []Article `gorm:"foreignKey:Author"`
		Sessions    []Session
	}

	Session struct {
		gorm.Model
		SessionToken string `gorm:"not null,type:varchar(100)"`
		AccessToken  string `gorm:"not null,type:varchar(100)"`
		IP           string `gorm:"not null,type:varchar(60)"`
		UserID       uint   `gorm:"not null"`
		User         User   `gorm:"foreignKey:UserID;references:ID"`
		Valid        bool   `gorm:"not null,default:true"`
	}

	Article struct {
		gorm.Model
		Title  string `gorm:"not null,type:varchar(100)"`
		Body   string `gorm:"not null"`
		Author uint   `gorm:"not null"`
		User   User   `gorm:"foreignKey:Author;references:ID"`
	}

	Package struct {
		gorm.Model
		Type          string `gorm:"not null,type:varchar(50)"`
		Specefication string `gorm:"not null"`
	}
)

func (u *User) ValidatePassword(password string) bool {
	var err error
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func (u *User) HashPassword(password string) (string, error) {
	var err error
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func Migrate() {
	DB.AutoMigrate(&User{})
}
