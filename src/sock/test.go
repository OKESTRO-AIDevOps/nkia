package sock

import (
	"fmt"

	"github.com/gorilla/websocket"
)

func DetachedServerCommunicator_Test(address string) error {
	c, _, err := websocket.DefaultDialer.Dial(address, nil)
	if err != nil {
		return fmt.Errorf("comm failed: %s", err.Error())
	}
	defer c.Close()

	return nil
}
