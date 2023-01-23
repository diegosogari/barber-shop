package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"

	"github.com/dsogari/barber-shop/graph/generated"
)

// CreateShop is the resolver for the createShop field.
func (r *mutationResolver) CreateShop(ctx context.Context, input generated.ShopInput) (*generated.Shop, error) {
	panic(fmt.Errorf("not implemented: CreateShop - createShop"))
}

// UpdateShop is the resolver for the updateShop field.
func (r *mutationResolver) UpdateShop(ctx context.Context, id int, input generated.ShopInput) (*generated.Shop, error) {
	panic(fmt.Errorf("not implemented: UpdateShop - updateShop"))
}

// DeleteShop is the resolver for the deleteShop field.
func (r *mutationResolver) DeleteShop(ctx context.Context, id int) (*generated.Shop, error) {
	panic(fmt.Errorf("not implemented: DeleteShop - deleteShop"))
}

// CreateService is the resolver for the createService field.
func (r *mutationResolver) CreateService(ctx context.Context, input generated.ServiceInput) (*generated.Service, error) {
	panic(fmt.Errorf("not implemented: CreateService - createService"))
}

// UpdateService is the resolver for the updateService field.
func (r *mutationResolver) UpdateService(ctx context.Context, id int, input generated.ServiceInput) (*generated.Service, error) {
	panic(fmt.Errorf("not implemented: UpdateService - updateService"))
}

// DeleteService is the resolver for the deleteService field.
func (r *mutationResolver) DeleteService(ctx context.Context, id int) (*generated.Service, error) {
	panic(fmt.Errorf("not implemented: DeleteService - deleteService"))
}

// CreateClient is the resolver for the createClient field.
func (r *mutationResolver) CreateClient(ctx context.Context, input generated.ClientInput) (*generated.Client, error) {
	panic(fmt.Errorf("not implemented: CreateClient - createClient"))
}

// UpdateClient is the resolver for the updateClient field.
func (r *mutationResolver) UpdateClient(ctx context.Context, id int, input generated.ClientInput) (*generated.Client, error) {
	panic(fmt.Errorf("not implemented: UpdateClient - updateClient"))
}

// DeleteClient is the resolver for the deleteClient field.
func (r *mutationResolver) DeleteClient(ctx context.Context, id int) (*generated.Client, error) {
	panic(fmt.Errorf("not implemented: DeleteClient - deleteClient"))
}

// CreateBarber is the resolver for the createBarber field.
func (r *mutationResolver) CreateBarber(ctx context.Context, input generated.BarberInput) (*generated.Barber, error) {
	panic(fmt.Errorf("not implemented: CreateBarber - createBarber"))
}

// UpdateBarber is the resolver for the updateBarber field.
func (r *mutationResolver) UpdateBarber(ctx context.Context, id int, input generated.BarberInput) (*generated.Barber, error) {
	panic(fmt.Errorf("not implemented: UpdateBarber - updateBarber"))
}

// DeleteBarber is the resolver for the deleteBarber field.
func (r *mutationResolver) DeleteBarber(ctx context.Context, id int) (*generated.Barber, error) {
	panic(fmt.Errorf("not implemented: DeleteBarber - deleteBarber"))
}

// CreateAttendance is the resolver for the createAttendance field.
func (r *mutationResolver) CreateAttendance(ctx context.Context, input generated.AttendanceInput) (*generated.Attendance, error) {
	panic(fmt.Errorf("not implemented: CreateAttendance - createAttendance"))
}

// UpdateAttendance is the resolver for the updateAttendance field.
func (r *mutationResolver) UpdateAttendance(ctx context.Context, id int, input generated.AttendanceInput) (*generated.Attendance, error) {
	panic(fmt.Errorf("not implemented: UpdateAttendance - updateAttendance"))
}

// DeleteAttendance is the resolver for the deleteAttendance field.
func (r *mutationResolver) DeleteAttendance(ctx context.Context, id int) (*generated.Attendance, error) {
	panic(fmt.Errorf("not implemented: DeleteAttendance - deleteAttendance"))
}

// AddAttendanceServices is the resolver for the addAttendanceServices field.
func (r *mutationResolver) AddAttendanceServices(ctx context.Context, id int, serviceIDs []int) (*generated.Attendance, error) {
	panic(fmt.Errorf("not implemented: AddAttendanceServices - addAttendanceServices"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
