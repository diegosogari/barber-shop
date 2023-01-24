package generated

func (o *Shop) UpdateFrom(input ShopInput) {
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

func (o *Service) UpdateFrom(input ServiceInput) {
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

func (o *Client) UpdateFrom(input ClientInput) {
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

func (o *Barber) UpdateFrom(input BarberInput) {
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

func (o *Attendance) UpdateFrom(input AttendanceInput) {
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
