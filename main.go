package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Song struct{
	ID int64
	Title string
	Artist string
}

func main() {
	cfg := mysql.Config{
		User: os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "music",
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")

	songs, err := songsByArtist("Between August and December")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Songs found: %v\n", songs)

	// Hard-code ID 2 here to test the query
	sng, err := songByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Song found: %v\n", sng)

	sngID, err := addSong(Song{
		Title: "Man Of The Year",
		Artist: "Juice WRLD",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added song: %v\n", sngID)
}

// songsByArtist queries for albums that have the specified artist name.
func songsByArtist(name string) ([]Song, error) {
	var songs []Song
	
	rows, err := db.Query("SELECT * FROM music.song  WHERE artist= ?", name)
	if err != nil {
		return nil, fmt.Errorf("songsByArtist %q: %v", name, err)
	}
	defer rows.Close()

	for rows.Next() {
		var sng Song
		if err := rows.Scan(&sng.ID, &sng.Title, &sng.Artist); err != nil {
			return nil, fmt.Errorf("songsByArtist %q: %v", name, err)
		}
		songs = append(songs, sng)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("songsByArtist %q: %v", name, err)
	}
	return songs, nil
}

// songByID queries for the song with the specified ID.
func songByID(id int64) (Song, error) {
	var sng Song
	row := db.QueryRow("SELECT * FROM music.song  WHERE id = ?", id)
	if err := row.Scan(&sng.ID, &sng.Title, &sng.Artist); err != nil {
		if err == sql.ErrNoRows {
			return sng, fmt.Errorf("songByID %d: no such song", id)
		}
		return sng, fmt.Errorf("songByID %d: %v", id, err)
	}
	return sng, nil
}

// addSong adds the specified song to the database,
// returning the album ID of the new entry
func addSong(sng Song) (int64, error) {
	result, err := db.Exec("INSERT INTO music.song (title, artist) VALUES (?, ?)", sng.Title, sng.Artist)
	if err != nil {
		return 0, fmt.Errorf("addSong: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addSong: %v", err)
	}
	return id, nil
}