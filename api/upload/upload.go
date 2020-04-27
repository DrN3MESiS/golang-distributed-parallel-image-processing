package upload

import (
	"errors"
	"fmt"
	"golang-distributed-parallel-image-processing/api/helpers"
	"io"
	"net/http"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func UploadResponse(c echo.Context) error {
	fmt.Println("[ACCESS] New connection to:\t/upload")
	user := c.Get("user").(*jwt.Token)
	token := user.Raw
	valid := helpers.IsTokenActive(token)
	if !valid {
		return helpers.ReturnJSON(c, http.StatusConflict, "Token is invalid or revoked")
	}

	image, err := c.FormFile("data")
	if err != nil {
		return errors.New("[] " + err.Error())
	}

	src, err := image.Open()
	if err != nil {
		return errors.New("[] " + err.Error())
	}
	defer src.Close()

	dst, err := os.Create(primitive.NewObjectID().Hex() + "_" + image.Filename)
	if err != nil {
		return errors.New("[] " + err.Error())
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return errors.New("[] " + err.Error())
	}

	fmt.Println("\t[OPERATION] Uploaded Image")
	return helpers.ReturnJSONMap(c, http.StatusOK, map[string]string{
		"message":  "An image has been successfully uploaded",
		"filename": image.Filename,
		"size":     strconv.Itoa(int(image.Size)),
	})
}
