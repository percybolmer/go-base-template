package modbus

import (
	"errors"
	"net"
)

type Client struct {
	IP           net.IP
	BatteryLevel int
}

type Data struct {
	BatteryLevel int
}

// Take a look at https://github.com/goburrow/modbus btw
func Connnect(ip net.IP) (*Client, error) {
	// Connect to Modbus
	if ip == nil {
		return nil, errors.New("the modbusIP is not a valid IP addy")
	}
	return &Client{IP: ip, BatteryLevel: 100}, nil
}

func (c *Client) FetchData() (Data, error) {
	currentData := Data{
		BatteryLevel: c.BatteryLevel,
	}
	// reduce battery level until next time
	c.BatteryLevel--
	return currentData, nil
}
