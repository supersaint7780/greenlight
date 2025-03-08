package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`                // unique integer id for the movie
	CreatedAt time.Time `json:"-"`                 // timestamp for when movie was added to our database
	Title     string    `json:"title"`             // Movie title
	Year      int32     `json:"year,omitempty"`    // Movie release year
	Runtime   int32     `json:"runtime,omitempty"` // Movie runtime (in minutes)
	Genres    []string  `json:"genres,omitempty"`  // Slice of genres for the movie (e.g comedy, sci-fi)
	Version   int32     `json:"version"`           // Starts at 1 incremented each time the movie info is updated
}
