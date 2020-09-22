package main

import (
	"fmt"
	"math"
	"strings"

	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"github.com/afdolriski/golang-docker/database"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func echo(args []string) error {
	if len(args) < 2 {
		return errors.New("no message to echo")
	}
	_, err := fmt.Println(strings.Join(args[1:], " "))
	return err
}

func main() {
	// if err := echo(os.Args); err != nil {
	// 	fmt.Fprintf(os.Stderr, "%+v\n", err)
	// 	os.Exit(1)
	// }

	makeAPICall()
}

func makeAPICall() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/hmac", func(c *gin.Context) {
		val, err := generateSHA256MAC("stuffinmysecertemate")
		if err != nil {
			c.String(404, err.Error())
			return
		}
		c.JSON(200, gin.H{
			"result": val,
		})
	})

	r.GET("/sqrt", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": calculateSquareRoot(),
		})
	})

	r.GET("/animal/:name", func(c *gin.Context) {
		animal, err := database.GetAnimal(c.Param("name"))
		if err != nil {
			c.String(404, err.Error())
			return
		}
		c.JSON(200, animal)
	})

	r.Run(":3000")
}

// just generate something expensive
func generateSHA256MAC(secret string) (string, error) {
	if secret == "" {
		return "", errors.New("no secret passed in")
	}
	data := "data"
	fmt.Printf("Secret: %s Data: %s\n", secret, data)

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	fmt.Println("Result: " + sha)

	return sha, nil
}

func calculateSquareRoot() string {
	// 	$x = 0.0001;
	//   for ($i = 0; $i <= 1000000; $i++) {
	//     $x += sqrt($x);
	//   }
	//   echo "OK!";

	x := 0.0001
	for i := 0; i <= 1000000; i++ {
		x += math.Sqrt(x)
	}
	return "OK!!"
}
