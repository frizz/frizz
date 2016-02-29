package flux_test

import (
	"testing"

	"kego.io/editor/client/flux"
)

var flights struct {
	country *CountryStore
	city    *CityStore
	price   *PriceStore
}

func TestFlights(t *testing.T) {
	/*
		d := flux.NewDispatcher()
		store.message = NewMessageStore()
		store.topic = NewTopicStore()
		d.Register(store.message)
		d.Register(store.topic)
		d.Dispatch(AddMessage("a"))
		m := store.message.GetMessages()
		assert.Equal(t, []string{"a"}, m)
	*/
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
	}
	return true
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
		//payload.WaitFor(flights.price)
		m.country = action.Country
	}
	return true
}

func (m *CountryStore) GetCountry() string {
	return m.country
}

type UpdatePriceAction struct {
	*flux.Action
	Price string
}

func UpdatePrice(price string) *UpdatePriceAction {
	return &UpdatePriceAction{
		Action: &flux.Action{},
		Price:  price,
	}
}

type PriceStore struct {
	price string
}

func (m *PriceStore) Handle(payload *flux.Payload) (finished bool) {
	switch action := payload.Action.(type) {
	case *UpdatePriceAction:
		payload.WaitFor(flights.city, flights.country)
		m.price = action.Price
	}
	return true
}

func (m *PriceStore) GetPrice() string {
	return m.price
}
