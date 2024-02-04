package meander

import "strings"

type Facade interface {
	Public() any
}

func Public(o any) any {
	if p, ok := o.(Facade); ok {
		return p.Public()
	}
	return o
}

var _ Facade = (*journey)(nil)

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

func (j *journey) Public() any {
	return map[string]any{
		"name":    j.Name,
		"journey": strings.Join(j.PlaceTypes, "|"),
	}
}
