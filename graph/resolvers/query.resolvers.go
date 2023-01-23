package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"

	"github.com/dsogari/barber-shop/graph/generated"
)

// ListShop is the resolver for the listShop field.
func (r *queryResolver) ListShop(ctx context.Context) ([]*generated.Shop, error) {
	panic(fmt.Errorf("not implemented: ListShop - listShop"))
}

// GetShop is the resolver for the getShop field.
func (r *queryResolver) GetShop(ctx context.Context, id int) (*generated.Shop, error) {
	panic(fmt.Errorf("not implemented: GetShop - getShop"))
}

// ListService is the resolver for the listService field.
func (r *queryResolver) ListService(ctx context.Context) ([]*generated.Service, error) {
	panic(fmt.Errorf("not implemented: ListService - listService"))
}

// GetService is the resolver for the getService field.
func (r *queryResolver) GetService(ctx context.Context, id int) (*generated.Service, error) {
	panic(fmt.Errorf("not implemented: GetService - getService"))
}

// ListClient is the resolver for the listClient field.
func (r *queryResolver) ListClient(ctx context.Context) ([]*generated.Client, error) {
	panic(fmt.Errorf("not implemented: ListClient - listClient"))
}

// GetClient is the resolver for the getClient field.
func (r *queryResolver) GetClient(ctx context.Context, id int) (*generated.Client, error) {
	panic(fmt.Errorf("not implemented: GetClient - getClient"))
}

// ListBarber is the resolver for the listBarber field.
func (r *queryResolver) ListBarber(ctx context.Context) ([]*generated.Barber, error) {
	panic(fmt.Errorf("not implemented: ListBarber - listBarber"))
}

// GetBarber is the resolver for the getBarber field.
func (r *queryResolver) GetBarber(ctx context.Context, id int) (*generated.Barber, error) {
	panic(fmt.Errorf("not implemented: GetBarber - getBarber"))
}

// ListAttendance is the resolver for the listAttendance field.
func (r *queryResolver) ListAttendance(ctx context.Context) ([]*generated.Attendance, error) {
	panic(fmt.Errorf("not implemented: ListAttendance - listAttendance"))
}

// GetAttendance is the resolver for the getAttendance field.
func (r *queryResolver) GetAttendance(ctx context.Context, id int) (*generated.Attendance, error) {
	panic(fmt.Errorf("not implemented: GetAttendance - getAttendance"))
}

// SearchAttendance is the resolver for the searchAttendance field.
func (r *queryResolver) SearchAttendance(ctx context.Context, input generated.AttendanceSearchInput) ([]*generated.Attendance, error) {
	panic(fmt.Errorf("not implemented: SearchAttendance - searchAttendance"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
