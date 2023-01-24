package model

import "github.com/dsogari/barber-shop/graph/generated"

type Shop struct {
	generated.Shop
	Attendances []*Attendance `json:"-"`
}
type Barber struct {
	generated.Barber
	Attendances []*Attendance `json:"-"`
}
type Client struct {
	generated.Client
	Attendances []*Attendance `json:"-"`
}
type Service struct {
	generated.Service
	Attendances []*Attendance `json:"-" gorm:"many2many:attendance_services;"`
}
type Attendance struct {
	generated.Attendance
	ShopID   int `json:"-"`
	BarberID int `json:"-"`
	ClientID int `json:"-"`
}

func (o *Shop) UpdateFrom(input generated.ShopInput) {
	if input.Address != nil {
		o.Address = *input.Address
	}
	if input.PhoneNumber != nil {
		o.PhoneNumber = *input.PhoneNumber
	}
	if input.Notes != nil {
		o.Notes = *input.Notes
	}
}

func (o *Service) UpdateFrom(input generated.ServiceInput) {
	if input.Name != nil {
		o.Name = *input.Name
	}
	if input.Cost != nil {
		o.Cost = *input.Cost
	}
	if input.Notes != nil {
		o.Notes = *input.Notes
	}
}

func (o *Client) UpdateFrom(input generated.ClientInput) {
	if input.Name != nil {
		o.Name = *input.Name
	}
	if input.PhoneNumber != nil {
		o.PhoneNumber = *input.PhoneNumber
	}
	if input.Notes != nil {
		o.Notes = *input.Notes
	}
}

func (o *Barber) UpdateFrom(input generated.BarberInput) {
	if input.Name != nil {
		o.Name = *input.Name
	}
	if input.PhoneNumber != nil {
		o.PhoneNumber = *input.PhoneNumber
	}
	if input.Notes != nil {
		o.Notes = *input.Notes
	}
}

func (o *Attendance) UpdateFrom(input generated.AttendanceInput) {
	if input.ShopID != nil {
		o.ShopID = *input.ShopID
	}
	if input.BarberID != nil {
		o.BarberID = *input.BarberID
	}
	if input.ClientID != nil {
		o.ClientID = *input.ClientID
	}
	if input.AttendedAt != nil {
		o.AttendedAt = *input.AttendedAt
	}
	if input.Notes != nil {
		o.Notes = *input.Notes
	}
}
