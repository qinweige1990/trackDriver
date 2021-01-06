package DB

import "trackDriver/domain"

const (
	JourneyStart = iota
	JourneyEnd
)

func CreateJourney(start, end, driver string) {
	db.Create(&domain.Journey{
		Start:  start,
		End:    end,
		Driver: driver,
		Status: JourneyStart,
	})
}

func UpdateJourney(id int) {
	db.Model(&domain.Journey{}).Where("id", id).Update("status", JourneyEnd)
}

func GetEndJourney() (journeys []domain.Journey) {
	db.Where("status", JourneyEnd).Find(&journeys)
	return
}