package factory
import "testing"
func TestNewRestaurant(t *testing.T) {
	NewRestaurant("k").GetFood()
	NewRestaurant("m").GetFood()
}
