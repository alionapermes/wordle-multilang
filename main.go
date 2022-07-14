package main

import (
	"database/sql"
	"net/http"
	"time"

	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	db *sql.DB
	e  *echo.Echo
}

func (s *Server) Run(port uint32) {
	db, _ := sql.Open("sqlite3", "words.db")
	s.db = db
	s.e = echo.New()
	defer s.db.Close()

	s.e.GET("/word", s.handler)

	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())

	s.e.Start(fmt.Sprintf(":%d", port))
}

func main() {
	s := Server{}
	s.Run(1234)
}

// func (s *Server) getWord(c echo.Context) error {
// 	row := s.db.QueryRow(
// 		"select Text from RussianWords where Id = $1",
// 		rand.Intn(3000),
// 	)

// 	var word string
// 	row.Scan(&word)
// 	return c.String(http.StatusOK, word)
// }

func (s *Server) handler(c echo.Context) error {
	lang := c.QueryParam("lang")
	word := s.sync(lang).getWordOfDay(lang)

	return c.JSON(http.StatusOK, word)
}

func (s *Server) sync(lang string) *Server {
	var wordOfDay = s.getWordOfDay(lang)

	t, _ := time.Parse("2006-01-02 00:00:00", wordOfDay.Day)
	wordDate := t.Unix()
	today := time.Today().Unix()

	// if wordDate

	fmt.Println("t:", t.Unix())
	fmt.Println(wordOfDay)

	return s
}

func (s *Server) getWordOfDay(lang string) WordOfDay {
	var wordOfDay WordOfDay
	s.db.QueryRow(`
		select
			*
		from
			WordOfDayHistory
		where
			Language = $1
		order by Id desc
		limit 1
	`, lang).Scan(
		&wordOfDay.Id, &wordOfDay.Text, &wordOfDay.Day, &wordOfDay.Lang)

	return wordOfDay
}

func (s *Server) pickWordsOfDay() WordOfDay {
	var count int
	s.db.QueryRow("select count(*) from RussianWords").Scan(&count)

	var word WordOfDay
	s.db.QueryRow("select * from RussianWords where Id = $1", count).Scan(
		&word.Id, &word.Text, &word.Lang, &word.Day)

	return word
}
