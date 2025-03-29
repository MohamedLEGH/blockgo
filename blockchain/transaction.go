package blockchain

import (
	"fmt"
	"strings"
)

type Transaction struct {
	Sender    string
	Receiver  string
	Value     int
	Signature string
}

func (t *Transaction) Sign(hexKey string) {
	var buf strings.Builder
	fmt.Fprintf(&buf, "%s|%s|%d",
		t.Sender,
		t.Receiver,
		t.Value)
	msg := buf.String()
	t.Signature = SignMessage(hexKey, msg)
}

func (t *Transaction) Verify() {
	var buf strings.Builder
	fmt.Fprintf(&buf, "%s|%s|%d",
		t.Sender,
		t.Receiver,
		t.Value)
	msg := buf.String()
	VerifySignature(t.Sender, msg, t.Signature)
}
