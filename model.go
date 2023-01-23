package main

import (
	"gorm.io/gorm"
)

type Shop struct {
	gorm.Model
	Address     string
	PhoneNumber *string
	Notes       *string
	Attendances []Attendance `json:"-"`
}

type Barber struct {
	gorm.Model
	Name        string
	PhoneNumber *string
	Notes       *string
	Attendances []Attendance `json:"-"`
}

type Service struct {
	gorm.Model
	Name        string
	Cost        int `gorm:"default:30;"`
	Notes       *string
	Attendances []*Attendance `json:"-" gorm:"many2many:attendance_services;"`
}

type Client struct {
	gorm.Model
	Name        string
	PhoneNumber *string
	Notes       *string
	Attendances []Attendance `json:"-"`
}

type Attendance struct {
	gorm.Model
	ShopID     uint
	BarberID   uint
	ClientID   uint
	AttendedAt int `gorm:"autoCreateTime"`
	Notes      *string
	Services   []*Service `json:"-" gorm:"many2many:attendance_services;"`
}
