package meander

type journey struct {
	Name       string
	PlaceTypes []string
}

var Journeys = []any{
	&journey{Name: "ロマンティック", PlaceTypes: []string{"park", "bar", "movie_theater", "restaurant", "florist", "taxi_stand"}},
	&journey{Name: "ショッピング", PlaceTypes: []string{"department_store", "cafe", "clothing_store", "jewelry_store", "shoe_store"}},
	&journey{Name: "ナイトライフ", PlaceTypes: []string{"bar", "casino", "food", "bar", "night_club", "bar", "bar", "hospital"}},
	&journey{Name: "カルチャー", PlaceTypes: []string{"museum", "cafe", "cemetery", "library", "art_gallery"}},
	&journey{Name: "リラックス", PlaceTypes: []string{"hair_care", "beauty_salon", "cafe", "spa"}},
}
