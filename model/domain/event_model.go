package domain

type Event struct {
	Id                    int
	UserId                int
	CategoryId            int
	Name                  string
	Tagline               string
	Description           string
	Organizer             string
	StartDate             string
	EndDate               string
	RegistrationStartDate string
	RegistrationEndDate   string
	Location              string
	Capacity              int
	BannerUrl             string
	Status                string
}
