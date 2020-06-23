package main

import (
	"log"
	"fmt"
	"time"
	"net/http"
	"errors"
	"crypto/subtle"
	"strings"
	//"io/ioutil"
	"strconv"
	"crypto/rand"
	"encoding/base64"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/argon2"
)

// DB CONST
var USER_DB 		= "root"
var PASS_DB 		= "root"
var IP_DB 			= "127.0.0.1:3306"
var TABLE_DB 		= "mailing"

// JWT CONST
var MY_SIGNING_KEY 	= []byte("mysupersecret")
var SERVER_URL 		= "http://localhost:9000/"

var (
    ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
    ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)

type params struct {
    memory      uint32
    iterations  uint32
    parallelism uint8
    saltLength  uint32
    keyLength   uint32
}

type User struct {
    ID   int    `json:"id"`
    Login string `json:"login"`
    Password string `json:"password"`
    Token string `json:"token"`
}

var user = User {}

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	insertUserToken(validToken, user.ID)

	/*client := &http.Client{}
	req, _ := http.NewRequest("GET", SERVER_URL, nil)
	req.Header.Set("Token", validToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}*/

	fmt.Fprintf(w, validToken)
	//fmt.Fprintf(w, string(body))
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Wyllis Monteiro"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(MY_SIGNING_KEY)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil

}

func handleRequests() {
	http.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":9001", nil))
}

func connectToBDD() (*sql.DB, error){
	db, err := sql.Open("mysql", USER_DB + ":" + PASS_DB + "@tcp(" + IP_DB + ")/" + TABLE_DB)

    if err != nil {
		return nil, err
	} else {
		fmt.Println("Logged to database")
		return db, nil
	}
}

func getOneUser(login string) error {

	db, err := connectToBDD()
	
	defer db.Close()

	if err != nil {
		return err
	}

	err = db.QueryRow("SELECT id, login, password FROM user WHERE login = ?", login).Scan(&user.ID, &user.Login, &user.Password)
	
	if err != nil {
		return err
	}

	return nil 
}

func insertUserToken(token string, user_id int) {
	db, err := connectToBDD()

	defer db.Close()

	// perform a db.Query insert
    insert, err := db.Query("UPDATE `user` SET `token` = ? WHERE `user`.`id` = ?", token, user_id)

    // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }
    // be careful deferring Queries if you are using transactions
    defer insert.Close()
}

func generateFromPassword(password string, p *params) (encodedHash string, err error) {
    salt, err := generateRandomBytes(p.saltLength)
    if err != nil {
        return "", err
    }

    hash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

    // Base64 encode the salt and hashed password.
    b64Salt := base64.RawStdEncoding.EncodeToString(salt)
    b64Hash := base64.RawStdEncoding.EncodeToString(hash)

    // Return a string using the standard encoded hash representation.
    encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.memory, p.iterations, p.parallelism, b64Salt, b64Hash)

    return encodedHash, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
    b := make([]byte, n)
    _, err := rand.Read(b)
    if err != nil {
        return nil, err
    }

    return b, nil
}

func comparePasswordAndHash(password, encodedHash string) (match bool, err error) {
    // Extract the parameters, salt and derived key from the encoded password
    // hash.
    p, salt, hash, err := decodeHash(encodedHash)
    if err != nil {
        return false, err
    }

    // Derive the key from the other password using the same parameters.
    otherHash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

    // Check that the contents of the hashed passwords are identical. Note
    // that we are using the subtle.ConstantTimeCompare() function for this
    // to help prevent timing attacks.
    if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
        return true, nil
    }
    return false, nil
}

func decodeHash(encodedHash string) (p *params, salt, hash []byte, err error) {
    vals := strings.Split(encodedHash, "$")
    if len(vals) != 6 {
        return nil, nil, nil, ErrInvalidHash
    }

    var version int
    _, err = fmt.Sscanf(vals[2], "v=%d", &version)
    if err != nil {
        return nil, nil, nil, err
    }
    if version != argon2.Version {
        return nil, nil, nil, ErrIncompatibleVersion
    }

    p = &params{}
    _, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
    if err != nil {
        return nil, nil, nil, err
    }

    salt, err = base64.RawStdEncoding.DecodeString(vals[4])
    if err != nil {
        return nil, nil, nil, err
    }
    p.saltLength = uint32(len(salt))

    hash, err = base64.RawStdEncoding.DecodeString(vals[5])
    if err != nil {
        return nil, nil, nil, err
    }
    p.keyLength = uint32(len(hash))

    return p, salt, hash, nil
}

func main() {
	fmt.Println("My Simple Client")

	err := getOneUser("wyllis")

	if err != nil {
		panic(err.Error())
		return
	}

	log.Printf(strconv.Itoa(user.ID))
	log.Printf(user.Login)
	log.Printf(user.Password)

	/*p := &params{
        memory:      64 * 1024,
        iterations:  3,
        parallelism: 2,
        saltLength:  16,
        keyLength:   32,
    }

    // Pass the plaintext password and parameters to our generateFromPassword
    // helper function.
    encodedHash, err := generateFromPassword("w", p)
    if err != nil {
        log.Fatal(err)
	}*/


	
	match, err := comparePasswordAndHash("w", user.Password)
    if err != nil {
        log.Fatal(err)
    }

	fmt.Printf("Match: %v\n", match)

	handleRequests()
}