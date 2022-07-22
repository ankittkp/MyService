package models

import "github.com/google/uuid"

type Movie struct {
	ID         uuid.UUID `json:"id"`
	Title      string    `json:"title"`
	Year       int       `json:"year"`
	ImdbRating float64   `json:"imdb_rating"`
	Director   string    `json:"director"`
	Released   bool      `json:"released"`
	Runtime    int       `json:"runtime"`
	Genre      string    `json:"genre"`
	Plot       string    `json:"plot"`
	Country    string    `json:"country"`
}

var MapOfMovies []map[string]Movie

func init() {
	MapOfMovies = make([]map[string]Movie, 0)
	//MapOfMovies = []map[string]Movie{
	//	{
	//		uuid.New().String(): Movie{
	//			Title:      "The Shawshank Redemption",
	//			Year:       1994,
	//			ImdbRating: 9.3,
	//			Director:   "Frank Darabont",
	//			Released:   true,
	//			Runtime:    142,
	//			Genre:      "Drama",
	//			Plot: "Andy Dufresne (Tim Robbins) is sentenced to two consecutive life terms in prison for the murders of his wife and her lover and is sentenced to a tough prison. However," +
	//				" only Andy knows he didn't commit the crimes. While there, he forms a friendship with Red (Morgan Freeman), experiences brutality of prison life, adapts, helps the warden, etc, " +
	//				"all in 19 years. When the warden terminates him, Andy finds himself in the last cell and is forced to live with a new set of problems: his wife is dead and his daughter is missing. He forms a friendship with the warden, who is a prison guard at the state prison",
	//			Country: "USA",
	//		},
	//		uuid.New().String(): Movie{
	//			Title:      "The Godfather",
	//			Year:       1972,
	//			ImdbRating: 9.2,
	//			Director:   "Francis Ford Coppola",
	//			Released:   true,
	//			Runtime:    175,
	//			Genre:      "Drama",
	//			Plot: "Widely regarded as one of the greatest films of all time, this mob drama, based on Mario Puzo's novel of the same name, focuses on the powerful Italian-American crime family of Don Vito Corleone (Marlon Brando). When the don's youngest son, Michael (Al Pacino)," +
	//				" reluctantly joins the Mafia, he becomes involved in the inevitable cycle of violence and betrayal. Although Michael tries to maintain a normal relationship with his wife, Kay (Diane Keaton), he is drawn deeper into the family business. Michael's mother, however, remains an outcast, and he is forced to take matters into his own hands." +
	//				" Michael is given a role in a famous mob drama, but his real profession is a mob boss. He is able to make his way in the world as a godfather, but he must stay out of trouble.",
	//			Country: "USA",
	//		},
	//		uuid.New().String(): Movie{
	//			Title:      "The Godfather: Part II",
	//			Year:       1974,
	//			ImdbRating: 9.0,
	//			Director:   "Francis Ford Coppola",
	//			Released:   true,
	//			Runtime:    202,
	//			Genre:      "Drama",
	//			Plot:       "The continuing saga of the Corleone crime family tells the story of a young Vito Corleone growing up in Sicily and in 1910s New York; and follows Michael Corleone in the 1950s as he attempts to expand the family business into Las Vegas, Hollywood, and Cuba.",
	//			Country:    "USA",
	//		},
	//		uuid.New().String(): Movie{
	//			Title:      "The Dark Knight",
	//			Year:       2008,
	//			ImdbRating: 9.0,
	//			Director:   "Christopher Nolan",
	//			Released:   true,
	//			Runtime:    152,
	//			Genre:      "Action",
	//			Plot:       "When the menace known as the Joker wreaks havoc and chaos on the people of Gotham, the caped crusader must come to terms with one of the greatest psychological tests of his ability to fight injustice.",
	//			Country:    "USA",
	//		},
	//		uuid.New().String(): Movie{
	//			Title:      "Pulp Fiction",
	//			Year:       1994,
	//			ImdbRating: 8.9,
	//			Director:   "Quentin Tarantino",
	//			Released:   true,
	//			Runtime:    154,
	//			Genre:      "Crime",
	//			Plot:       "Vincent Vega (John Travolta) and Jules Winnfield (Samuel L. Jackson) are hitmen with a penchant for philosophical discussions. In this ultra-hip, multi-strand crime movie, their storyline is interwoven with those of their boss, gangster Marsellus Wallace (Ving Rhames) ; his actress wife, Mia (Uma Thurman) ; struggling boxer Butch Coolidge (Bruce Willis) ; master fixer Winston Wolfe (Harvey Keitel) and a nervous pair of armed robbers, Pumpkin (Tim Roth) and Honey Bunny (Amanda Plummer)",
	//			Country:    "USA",
	//		},
	//		uuid.New().String(): Movie{
	//			Title:      "The Matrix",
	//			Year:       1999,
	//			ImdbRating: 8.7,
	//			Director:   "Lana Wachowski",
	//			Released:   true,
	//			Runtime:    136,
	//			Genre:      "Action",
	//			Plot:       "A computer hacker learns from mysterious rebels about the true nature of his reality and his role in the war against its controllers.",
	//			Country:    "USA",
	//		},
	//	},
}
