package rest

import (
	"context"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	dblayer "github.com/hariboGCS/Back/src/dbconn"
	"github.com/hariboGCS/Back/src/model"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

type handlerInterface interface {
	GetScore(c echo.Context)
	ReceiveScore(c echo.Context)
	Login(c echo.Context) error
	Signup(c echo.Context) error
	updateUser(c echo.Context) error
	SignOut(c echo.Context)
	deleteUser(c echo.Context) error
	GetNotice(c echo.Context)
	GetComplaints(c echo.Context)
}
type Handler struct {
	Handler handlerInterface
}

func NewHandler() *Handler {
	return new(Handler)
}
func GetMainPage(c echo.Context) (err error) {
	return c.String(200, "main page")
}
func Signup(c echo.Context) (err error) {
	// Bind
	u := &model.User{ID: bson.NewObjectId().Hex()}
	if err = c.Bind(u); err != nil {
		return err
	}
	// Validate
	if u.Email == "" || u.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	}
	collection, err := dblayer.GetDBCollection()
	collection.InsertOne(context.TODO(), u)
	if err != nil {
		return err
	}
	defer collection.Database().Client().Disconnect(context.TODO())

	return c.JSON(http.StatusCreated, u)
}

func Signin(c echo.Context) (err error) {
	// Bind
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return
	}
	filter := bson.M{"token": u.Token}
	collection, err := dblayer.GetDBCollection()
	if err != nil {
		return err
		// return &echo.HTTPError{Code: http.StatusUnauthorized,Message:"invalid email or password"}
	}
	err = collection.FindOne(context.TODO(), filter).Decode(&u)
	_, err = collection.UpdateOne(context.TODO(), filter, &u)
	defer collection.Database().Client().Disconnect(context.TODO())
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response
	u.Token, err = token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	u.Password = "" // Don't send password
	return c.JSON(http.StatusOK, u)
}
