package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/karashiiro/bingode"
	"github.com/xivapi/godestone/v2"
)

type characterResult struct {
	Bio string `json:"bio"`
}

func main() {
	s := godestone.NewScraper(bingode.New(), godestone.EN)

	r := gin.Default()

	// Character endpoint
	r.GET("/character/:id", func(c *gin.Context) {
		cIdStr := c.Param("id")

		cId, err := strconv.ParseUint(cIdStr, 10, 32)
		if err != nil {
			c.AbortWithError(400, err)
			return
		}

		character, err := s.FetchCharacter(uint32(cId))
		if err != nil {
			c.AbortWithError(404, err)
			return
		}

		res := characterResult{
			Bio: character.Bio,
		}

		c.JSON(200, res)
	})

	// Character search endpoint
	r.GET("/search/character/:world/:first/:last", func(c *gin.Context) {
		worldName := strings.ToLower(c.Param("world"))
		if worldName == "" {
			c.AbortWithError(400, errors.New("world name not provided"))
			return
		}

		characterName := strings.ToLower(fmt.Sprintf("%s %s", c.Param("first"), c.Param("last")))
		if characterName == "" {
			c.AbortWithError(400, errors.New("character name not provided"))
			return
		}

		for res := range s.SearchCharacters(godestone.CharacterOptions{
			Name:  characterName,
			World: strings.ToUpper(string(worldName[0])) + worldName[1:], // World name must be captialized
		}) {
			if res.Error != nil {
				c.AbortWithError(500, res.Error)
				return
			}

			if strings.ToLower(res.Name) == characterName && strings.ToLower(res.World) == worldName {
				c.JSON(200, res)
				return
			}
		}

		c.AbortWithError(404, errors.New("no character matching those parameters was found"))
	})

	r.Run(":3999")
}
