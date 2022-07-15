package main

import (
	"database/sql"
	"math/rand"
	"net/http"
	"os"

	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/golang-module/carbon/v2"
)

type Server struct {
	Token string
	db    *sql.DB
	e     *echo.Echo
}

func (s *Server) Run(port string) {
	db, _ := sql.Open("sqlite3", "words.db")
	s.db = db
	s.e = echo.New()
	defer s.db.Close()

	s.e.GET("/word", s.getWord)
	s.e.GET("/check", s.checkWord)

	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())

	s.e.Start(port)
}

func main() {
	s := Server{Token: os.Getenv("WORDLE_TOKEN")}
	fmt.Println("token:", s.Token)
	s.Run(":" + os.Getenv("WORDLE_PORT"))
}

func (s *Server) getWord(c echo.Context) error {
	// token := c.QueryParam("token")
	// if token != s.Token {
	// 	return c.NoContent(http.StatusUnauthorized)
	// }

	lang := c.QueryParam("lang")
	word := s.syncAndFetch(lang)

	return c.JSON(http.StatusOK, Payload{
		Word: word.Text,
		Lang: word.Lang,
		Next: word.Next,
	})
}

func (s *Server) checkWord(c echo.Context) error {
	// token := c.QueryParam("token")
	// if token != s.Token {
	// 	return c.NoContent(http.StatusUnauthorized)
	// }

	testWord := c.QueryParam("word")
	trueWord := s.syncAndFetch(c.QueryParam("lang"))

	if trueWord.Text == testWord {
		return c.NoContent(http.StatusOK)
	}

	return c.NoContent(http.StatusBadRequest)
}

func (s *Server) syncAndFetch(lang string) WordOfDay {
	var word, err = s.getWordOfDay(lang)

	now := carbon.Now().Timestamp()

	if (err != nil) || (now >= word.Next) {
		var text string

		if lang == "en" {
			text = s.pickWordOfDay("English")
		} else if lang == "ru" {
			text = s.pickWordOfDay("Russian")
		}

		word = WordOfDay{
			Text: text,
			Lang: lang,
			Next: carbon.Tomorrow().Timestamp(),
		}
		fmt.Println("after pick:", word)

		s.db.Exec(`
			insert into
				History(Word, Lang, Next)
			values
				($1, $2, $3)
		`, word.Text, word.Lang, word.Next)
	}

	return word
}

func (s *Server) getWordOfDay(lang string) (WordOfDay, error) {
	var word WordOfDay
	err := s.db.QueryRow(`
		select
			*
		from
			History
		where
			Lang = $1
		order by Id desc
		limit 1
	`, lang).Scan(
		&word.Id, &word.Text, &word.Lang, &word.Next)

	if err != nil {
		return WordOfDay{}, err
	}

	return word, nil
}

func (s *Server) pickWordOfDay(tablePrefix string) string {
	tableName := tablePrefix + "Words"

	var count int
	s.db.QueryRow("select count(*) from " + tableName).Scan(&count)

	random := rand.Intn(count)

	var word string
	s.db.QueryRow("select Text from "+tableName+" where Id = $1", random).
		Scan(&word)

	return word
}
