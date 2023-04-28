package authcontroller

import (
	"encoding/json"
	"jwtMux/config"
	"jwtMux/helper"
	"jwtMux/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// mengambil inputan json
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		// log.Fatal("gagal mendecode json")
		resp := map[string]string{"message": err.Error()}
		helper.ResponJSON(w, http.StatusBadRequest, resp)
		return
	}
	defer r.Body.Close()

	// ambil data user berdasarkan username
	var user models.User
	if err := models.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "username atau password salah"}
			helper.ResponJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": "gagal login"}
			helper.ResponJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	// cek apakah password valid
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"message": "username atau password salah"}
		helper.ResponJSON(w, http.StatusUnauthorized, response)
		return
	}

	// proses pembuatan token jwt
	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "golang-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// mendeklarasikan algoritma yang akan digunakan untuk signing
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// signed token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponJSON(w, http.StatusInternalServerError, response)
		return
	}

	// set token yang ke cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	response := map[string]string{"message": "login berhasil"}
	helper.ResponJSON(w, http.StatusOK, response)
	return
}

func Register(w http.ResponseWriter, r *http.Request) {
	// mengambil inputan json
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		// log.Fatal("gagal mendecode json")
		resp := map[string]string{"message": err.Error()}
		helper.ResponJSON(w, http.StatusBadRequest, resp)
		return
	}
	defer r.Body.Close()

	// hash password menggunakan bycrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	// insert ke database
	if err := models.DB.Create(&userInput).Error; err != nil {
		// log.Fatal("gagal menyimpan data")
		resp := map[string]string{"message": err.Error()}
		helper.ResponJSON(w, http.StatusInternalServerError, resp)
		return
	}

	// resp, _ := json.Marshal(map[string]string{"message": "register success"})
	// w.Header().Add("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(resp)

	resp := map[string]string{"message": "register success"}
	helper.ResponJSON(w, http.StatusOK, resp)
}

func Logout(w http.ResponseWriter, r *http.Request) {

}
