package flux_test

import (
	"testing"

	"fmt"

	"kego.io/editor/client/flux"
	"kego.io/kerr/assert"
)

var flights struct {
	country *CountryStore
	city    *CityStore
	price   *PriceStore
}

/*
func TestLoop(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovering...")
			assert.Equal(t, "panic: *flux_test.CityStore and *flux_test.CountryStore are waiting for each other.", r)
		}
	}()

	d := flux.NewDispatcher()
	flights.country = &CountryStore{}
	flights.city = &CityStore{}

	d.Register(flights.country)
	d.Register(flights.city)

	//fmt.Println("Before loop...")
	//d.Dispatch(Loop())
	//fmt.Println("After loop...")

}*/

func TestFlights(t *testing.T) {
	d := flux.NewDispatcher()
	flights.country = &CountryStore{}
	flights.city = &CityStore{}
	flights.price = &PriceStore{}

	d.Register(flights.country)
	d.Register(flights.city)
	d.Register(flights.price)

	d.Dispatch(UpdateCountry("France"))

	assert.Equal(t, "France", flights.country.GetCountry())
	assert.Equal(t, "Paris", flights.city.GetCity())
	assert.Equal(t, "€100", flights.price.GetPrice())
}

type LoopAction struct {
	*flux.Action
}

func Loop() *LoopAction {
	return &LoopAction{
		Action: &flux.Action{},
	}
}

type UpdateCityAction struct {
	*flux.Action
	City string
}

func UpdateCity(city string) *UpdateCityAction {
	return &UpdateCityAction{
		Action: &flux.Action{},
		City:   city,
	}
}

type CityStore struct {
	city string
}

func (m *CityStore) Handle(payload *flux.Payload) (finished bool) {
	switch action := payload.Action.(type) {
	case *UpdateCityAction:
		m.city = action.City
	case *UpdateCountryAction:
		payload.WaitFor(flights.country)
		m.city = getCapital(flights.country.GetCountry())
	case *LoopAction:
		payload.WaitFor(flights.country)
	}
	return true
}

func getCapital(country string) string {
	switch country {
	case "UK":
		return "London"
	case "France":
		return "Paris"
	}
	return fmt.Sprint("Capital of ", country)
}

func (m *CityStore) GetCity() string {
	return m.city
}

type UpdateCountryAction struct {
	*flux.Action
	Country string
}

func UpdateCountry(country string) *UpdateCountryAction {
	return &UpdateCountryAction{
		Action:  &flux.Action{},
		Country: country,
	}
}

type CountryStore struct {
	country string
}

func (m *CountryStore) Handle(payload *flux.Payload) (finished bool) {
	switch action := payload.Action.(type) {
	case *UpdateCountryAction:
		m.country = action.Country
	case *LoopAction:
		payload.WaitFor(flights.city)
	}
	return true
}

func (m *CountryStore) GetCountry() string {
	return m.country
}

type PriceStore struct {
	price string
}

func (m *PriceStore) Handle(payload *flux.Payload) (finished bool) {
	switch payload.Action.(type) {
	case *UpdateCountryAction, *UpdateCityAction:
		payload.WaitFor(flights.city)
		m.price = calculatePrice(flights.country.GetCountry(), flights.city.GetCity())
	}
	return true
}

func (m *PriceStore) GetPrice() string {
	return m.price
}

func calculatePrice(country string, city string) string {
	switch country {
	case "UK":
		return "£100"
	case "France":
		return "€100"
	}
	return "$100"
}
