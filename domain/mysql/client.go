package mysql

import "errors"

// Client is a future database client
type Client struct {
}

// Connect will connect to the db and return a struct
func Connect() (Client, error) {
	return Client{}, nil
}

func (c Client) InsertData(batteryLevel int) error {

	if batteryLevel == 95 {
		return errors.New("failed to insert into database")
	}
	return nil
}
