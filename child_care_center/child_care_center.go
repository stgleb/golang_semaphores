package child_care_center

import "log"

type ChildCareCenter struct {
	ChildIn  chan Child
	ChildOut chan Child

	AdultIn  chan Adult
	AdultOut chan Adult

	childCount int
	adultCount int
}

func NewChildCenter() ChildCareCenter {
	return ChildCareCenter{
		ChildIn:    nil,
		ChildOut:   make(chan Child),
		AdultIn:    make(chan Adult),
		AdultOut:   make(chan Adult),
		childCount: 0,
		adultCount: 0,
	}
}

func (center *ChildCareCenter) Run() {
	var (
		cIn  chan Child
		aOut chan Adult
	)
	for {
		select {
		case child := <-center.ChildIn:
			log.Printf("Enter: %s", child)
			center.childCount++
			// Lock children channel
			if center.childCount == center.adultCount*3 {
				cIn = center.ChildIn
				center.ChildIn = nil
			}

			// Unlock adult out channel.
			if center.adultCount == center.childCount/3-1 {
				center.AdultOut = aOut
			}
		case child := <-center.ChildOut:
			log.Printf("Exit: %s", child)
			// Unlock children in channel
			if center.childCount == center.adultCount*3 {
				center.ChildIn = cIn
			}

			center.childCount--

			if center.adultCount == center.childCount/3 {
				aOut = center.AdultOut
				center.AdultOut = nil
			}
		case adult := <-center.AdultIn:
			log.Printf("Enter: %s", adult)
			// Unlock children channel since count of adults become bigger
			if center.childCount == center.adultCount*3 {
				center.ChildIn = cIn
			}

			if center.AdultOut == nil {
				center.AdultOut = aOut
			}

			center.adultCount++
		case adult := <-center.AdultOut:
			log.Printf("Exit: %s", adult)
			center.adultCount--

			if center.adultCount == center.childCount/3 {
				aOut = center.AdultOut
				center.AdultOut = nil
			}
		}
	}
}
