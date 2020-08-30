package opensubtitles

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestSearchService_Movie(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/v1/search/movie", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
    "data": [
        {
            "id": "646193",
            "type": "movie",
            "attributes": {
                "title": "The Matrix",
                "original_title": "The Matrix",
                "imdb_id": 133093,
                "tmdb_id": 603,
                "feature_id": "646193",
                "year": "1999",
                "subtitles_counts": {
                    "pl": 151,
                    "en": 117,
                    "tr": 68,
                    "ro": 56,
                    "cs": 54,
                    "es": 54,
                    "pt-BR": 49,
                    "sl": 41,
                    "pt-PT": 39,
                    "sr": 38,
                    "el": 37,
                    "bg": 29,
                    "he": 28,
                    "nl": 22,
                    "fi": 19,
                    "fr": 19,
                    "hu": 17,
                    "ar": 16,
                    "ru": 15,
                    "hr": 14,
                    "da": 13,
                    "et": 11,
                    "sv": 11,
                    "sq": 10,
                    "bs": 7,
                    "de": 7,
                    "it": 7,
                    "ko": 7,
                    "no": 7,
                    "fa": 7,
                    "sk": 7,
                    "mk": 6,
                    "zh-CN": 5,
                    "ms": 4,
                    "zh-TW": 4,
                    "bn": 3,
                    "id": 3,
                    "lt": 3,
                    "is": 2,
                    "ja": 2,
                    "th": 2,
                    "ca": 1,
                    "hi": 1,
                    "ml": 1,
                    "mn": 1,
                    "vi": 1
                },
                "url": "https://www.opensubtitles.com/en/movies/1999-the-matrix",
                "img_url": "https://s9.osdb.link/features/3/9/1/646193.jpg"
            }
        }
    ]
}`)
	})
	opt := SearchOptions{Query: "The Matrix"}
	shows, _, err := client.Search.Movie(context.Background(), &opt)
	if err != nil {
		t.Errorf("Search.Movie returned error: %v", err)
	}
	want := Shows{Data: []Show{{
		ID:         "646193",
		Type:       "movie",
		Attributes: ShowAttributes{
			Title:           "The Matrix",
			OriginalTitle:   "The Matrix",
			ImdbID:          133093,
			TmdbID:          603,
			FeatureID:       "646193",
			Year:            "1999",
			TitleAka:        nil,
			SubtitlesCounts: SubtitlesCounts{
				Pl:   151,
				En:   117,
				Tr:   68,
				Ro:   56,
				Cs:   54,
				Es:   54,
				PtBR: 49,
				Sl:   41,
				PtPT: 39,
				Sr:   38,
				El:   37,
				Bg:   29,
				He:   28,
				Nl:   22,
				Fi:   19,
				Fr:   19,
				Hu:   17,
				Ar:   16,
				Ru:   15,
				Hr: 14,
				Da: 13,
				Et: 11,
				Sv: 11,
				Sq: 10,
				Bs: 7,
				De: 7,
				It: 7,
				Ko: 7,
				No: 7,
				Fa: 7,
				Sk: 7,
				Mk: 6,
				ZhCN: 5,
				Ms: 4,
				ZhTW: 4,
				Bn: 3,
				ID: 3,
				Lt: 3,
				Is: 2,
				Ja: 2,
				Th: 2,
				Ca: 1,
				Hi: 1,
				Ml: 1,
				Mn: 1,
				Vi: 1,
			},
			URL:             "https://www.opensubtitles.com/en/movies/1999-the-matrix",
			ImgURL:          "https://s9.osdb.link/features/3/9/1/646193.jpg",
		},
	},
		
	}}
	if !reflect.DeepEqual(shows, &want) {
		t.Errorf("Search.Movie returned %+v, want %+v", shows, want)
	}
}

func TestSearchService_TV(t *testing.T) {

}

func TestSearchService_Title(t *testing.T) {

}
