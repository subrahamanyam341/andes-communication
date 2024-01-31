package connection

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/subrahamanyam341/andes-communication/websocket/data"
	logger "github.com/subrahamanyam341/andes-logger-123"
)

var log = logger.GetOrCreate("connection")

type wsConnClient struct {
	mut      sync.RWMutex
	conn     *websocket.Conn
	clientID string
}

// NewWSConnClient creates a new wrapper over a websocket connection
func NewWSConnClient() *wsConnClient {
	return &wsConnClient{}
}

// NewWSConnClientWithConn creates a new wrapper over a provided websocket connection
func NewWSConnClientWithConn(conn *websocket.Conn) *wsConnClient {
	wsc := &wsConnClient{
		conn: conn,
	}
	wsc.clientID = fmt.Sprintf("%p", wsc)

	return wsc
}

// OpenConnection will open a new client with a background context
func (wsc *wsConnClient) OpenConnection(url string) error {
	wsc.mut.Lock()
	defer wsc.mut.Unlock()

	if wsc.conn != nil {
		return data.ErrConnectionAlreadyOpen
	}

	var err error
	wsc.conn, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return err
	}

	return nil
}

// ReadMessage calls the underlying reading message ws connection func
func (wsc *wsConnClient) ReadMessage() (messageType int, p []byte, err error) {
	conn, err := wsc.getConn()
	if err != nil {
		return 0, nil, err
	}

	return conn.ReadMessage()
}

// WriteMessage calls the underlying write message ws connection func
func (wsc *wsConnClient) WriteMessage(messageType int, payload []byte) error {
	wsc.mut.Lock()
	defer wsc.mut.Unlock()

	if wsc.conn == nil {
		return data.ErrConnectionNotOpen
	}

	return wsc.conn.WriteMessage(messageType, payload)
}

// IsOpen will return true if the connection is open, false otherwise
func (wsc *wsConnClient) IsOpen() bool {
	wsc.mut.RLock()
	defer wsc.mut.RUnlock()

	return wsc.conn != nil
}

func (wsc *wsConnClient) getConn() (*websocket.Conn, error) {
	wsc.mut.RLock()
	defer wsc.mut.RUnlock()

	if wsc.conn == nil {
		return nil, data.ErrConnectionNotOpen
	}

	conn := wsc.conn

	return conn, nil
}

// GetID will return the unique id of the client
func (wsc *wsConnClient) GetID() string {
	return wsc.clientID
}

// Close will try to cleanly close the connection, if possible
func (wsc *wsConnClient) Close() error {
	// critical section
	wsc.mut.Lock()
	defer wsc.mut.Unlock()

	if wsc.conn == nil {
		return data.ErrConnectionNotOpen
	}

	log.Debug("closing ws connection...")

	//Cleanly close the connection by sending a close message and then
	//waiting (with timeout) for the server to close the connection.
	err := wsc.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Trace("cannot send close message", "error", err)
	}

	wsc.conn.CloseHandler()

	err = wsc.conn.Close()
	if err != nil {
		return err
	}

	wsc.conn = nil
	return nil
}

// IsInterfaceNil -
func (wsc *wsConnClient) IsInterfaceNil() bool {
	return wsc == nil
}
