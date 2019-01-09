package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//! Default fuction API
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/assasin", GetdataAssasin)
	e.GET("/mage", GetdataMage)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

//Get Assasin data
func GetdataAssasin(c echo.Context) error {
	//! Openfile
	fileHandle, err := os.Open("assassin_db.txt")
	if err != nil {
		log.Fatalf("Failed open file message.txt: %s", err)
	}
	defer fileHandle.Close()

	//! Scan text in File
	byteValue, _ := ioutil.ReadAll(fileHandle)
	result := HeroList2{}
	//making data to correct json.format
	json.Unmarshal([]byte(byteValue), &result)
	fmt.Print(string(byteValue))

	//! Return this if it's work!!
	return c.JSON(http.StatusOK, result)
}

//Get Mage data
func GetdataMage(c echo.Context) error {
	//! Openfile
	fileHandle, err := os.Open("mage_db.txt")
	if err != nil {
		log.Fatalf("Failed open file message.txt: %s", err)
	}
	defer fileHandle.Close()

	//! Scan text in File
	byteValue, _ := ioutil.ReadAll(fileHandle)
	result := HeroList2{}
	//! "[" and "]" are making data to correct json.format
	json.Unmarshal([]byte(byteValue), &result)
	fmt.Print(string(byteValue))

	//! Return this if it's work!!
	return c.JSON(http.StatusOK, result)
}

//json struct
type HeroList2 struct {
	Imgtitleurl         string `json:"img_title_url"`
	Imgstatus           string `json:"img_status"`
	Name                string `json:"name"`
	Class               []string `json:"class"`
	Role                string `json:"role"`
	Hp                  int    `json:"hp"`
	AttackDamage        int    `json:"attackdamage"`
	MagicDamage         int    `json:"magicdamage"`
	Armor               int    `json:"armor"`
	MagicDefense        int    `json:"magicdefense"`
	Mana                int    `json:"mana"`
	SpeedMovement       int    `json:"speedmovement"`
	IgnoreAttackdefense int    `json:"ignoreattackdefense"`
	IgnoreMagicdefense  int    `json:"ignoremagicdefense"`
	SpeedAttack         int    `json:"speedattack"`
	RateCritical        int    `json:"ratecritical"`
	DamageCritical      string `json:"damagecritical"`
	LifeSteal           int    `json:"lifesteal"`
	MagicLifeSteal      int    `json:"magiclifesteal"`
	CoolDownSpeed       int    `json:"cooldownspeed"`
	AttackRange         string `json:"attackrange"`
	Resistance          int    `json:"resistance"`
	HPper5s             int    `json:"hpper5s"`
	MPper5s             int    `json:"mpper5s"`
}
