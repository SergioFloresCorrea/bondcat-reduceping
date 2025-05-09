package bondcat

import (
	"context"
	"fmt"
	"log"
	"net"
)

/*
	This does not dial the bonded connection, this instead is a wrapper
	to make a dialer for our bonded connection libary.
*/

type bondCatDailer struct {
	dest   string
	dialer net.Dialer
}

func (b *bondCatDailer) DialContext(ctx context.Context) (net.Conn, error) {
	log.Printf("New TCP connection to %s", b.dest)
	return b.dialer.DialContext(ctx, "tcp", b.dest)
}

func (b *bondCatDailer) Label() string {
	return fmt.Sprintf("BCD - %s", b.dest)
}
