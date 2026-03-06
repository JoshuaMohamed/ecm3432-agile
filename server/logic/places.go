package logic

import (
	"fmt"
	"regexp"
	"strings"
)

type Place struct {
	Name     string `json:"name"`
	Postcode string `json:"postcode"`
	CoverPath string `json:"cover_path,omitempty"`
}

func CreatePlace(db Database, place Place) error {
	if !IsValidPostcode(place.Postcode) {
		return fmt.Errorf("Invalid postcode: %s", place.Postcode)
	}

	err := db.CreateRow("Places", []string{"name", "postcode", "cover"}, []interface{}{place.Name, place.Postcode, place.CoverPath})
	if err != nil {
		return err
	}

	return nil
}

func GetPlaces(db Database, postcode, filter string, limit, offset int) ([]Place, error) {
	postcode = strings.ToUpper(strings.TrimSpace(postcode))

	if !IsValidPostcode(postcode) {
		return nil, fmt.Errorf("Invalid postcode: %s", postcode)
	}

	prefix, err := GetSearchPrefix(postcode, filter)
	if err != nil {
		return nil, err
	}

	rows, err := db.GetPlaces(prefix, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	places := make([]Place, 0)
	for rows.Next() {
		var place Place
		if err := rows.Scan(&place.Name, &place.Postcode, &place.CoverPath); err != nil {
			return nil, err
		}
		places = append(places, place)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return places, nil
}

func IsValidPostcode(postcode string) bool {
	re := regexp.MustCompile(`^[A-Z0-9]{2,4} [A-Z0-9]{3}$`)
	return re.MatchString(postcode)
}

func GetSearchPrefix(postcode, filter string) (string, error) {
	parts := strings.Split(postcode, " ")
	if len(parts) != 2 {
		return "", fmt.Errorf("Invalid postcode: %s", postcode)
	}
	district := parts[0]
	incode := parts[1]
	area := regexp.MustCompile(`^[A-Z]+`).FindString(district)
	sector := district + " " + string(incode[0])

	switch filter {
	case "area":
		return area, nil
	case "district":
		return district, nil
	case "sector":
		return sector, nil
	}
	return "", fmt.Errorf("Invalid filter: %s", filter)
}
