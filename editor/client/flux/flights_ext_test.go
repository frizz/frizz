package flux_test

import (
	"testing"

	"fmt"

	"kego.io/editor/client/flux"
	"kego.io/kerr/assert"
)

type Flights struct {
	Dispatcher *flux.Dispatcher
	Country    *CountryStore
	City       *CityStore
	Price      *PriceStore
}

func TestFlights(t *testing.T) {

	a := &Flights{}
	a.Country = &CountryStore{app: a, Store: &flux.Store{}}
	a.City = &CityStore{app: a, Store: &flux.Store{}}
	a.Price = &PriceStore{app: a, Store: &flux.Store{}}
	a.Dispatcher = flux.NewDispatcher(a.Country, a.City, a.Price)

	a.Dispatcher.Dispatch(&UpdateCountry{Country: "France"})

	assert.Equal(t, "France", a.Country.GetCountry())
	assert.Equal(t, "Paris", a.City.GetCity())
	assert.Equal(t, "€100", a.Price.GetPrice())
}

type UpdateCity struct {
	City string
}

type CityStore struct {
	*flux.Store
	app  *Flights
	city string
}

func (m *CityStore) Handle(payload *flux.Payload) (finished bool) {
	switch action := payload.Action.(type) {
	case *UpdateCity:
		m.city = action.City
	case *UpdateCountry:
		payload.WaitFor(m.app.Country)
		m.city = getCapital(m.app.Country.GetCountry())
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

type UpdateCountry struct {
	Country string
}

type CountryStore struct {
	*flux.Store
	app     *Flights
	country string
}

func (m *CountryStore) Handle(payload *flux.Payload) (finished bool) {
	switch action := payload.Action.(type) {
	case *UpdateCountry:
		m.country = action.Country
	}
	return true
}

func (m *CountryStore) GetCountry() string {
	return m.country
}

type PriceStore struct {
	*flux.Store
	app   *Flights
	price string
}

func (m *PriceStore) Handle(payload *flux.Payload) (finished bool) {
	switch payload.Action.(type) {
	case *UpdateCountry, *UpdateCity:
		payload.WaitFor(m.app.City)
		m.price = calculatePrice(m.app.Country.GetCountry(), m.app.City.GetCity())
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
