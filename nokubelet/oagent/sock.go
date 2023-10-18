package sock

import (
	"fmt"

	"github.com/gorilla/websocket"
)

var SESSION_SYM_KEY = ""

func DetachedServerCommunicator(address string, email string, cluster_id string) error {

	c, _, err := websocket.DefaultDialer.Dial(address, nil)
	if err != nil {
		return fmt.Errorf("comm failed: %s", err.Error())
	}
	defer c.Close()

	err = ServerAuthChallenge(c, email, cluster_id)

	if err != nil {
		return fmt.Errorf("comm failed: %s", err.Error())
	}

	if err := SockCommunicationHandler_LinearInstruction_PrintOnly(c); err != nil {
		return fmt.Errorf("comm failed: %s", err.Error())
	}

	return nil
}

func DetachedServerCommunicatorWithUpdate(address string, email string, cluster_id string, token string) error {

	c, _, err := websocket.DefaultDialer.Dial(address, nil)
	if err != nil {
		return fmt.Errorf("commwup failed: %s", err.Error())
	}
	defer c.Close()

	err = ServerUpdateChallenge(c, email, cluster_id, token)

	if err != nil {
		return fmt.Errorf("commwup failed: %s", err.Error())
	}

	err = ServerAuthChallenge(c, email, cluster_id)

	if err != nil {
		return fmt.Errorf("commwup failed: %s", err.Error())
	}

	if err := SockCommunicationHandler_LinearInstruction_PrintOnly(c); err != nil {
		return fmt.Errorf("commwup failed: %s", err.Error())
	}

	return nil
}
