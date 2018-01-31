package mail

import (
	"io/ioutil"
	"testing"
)

func TestNewAddress(t *testing.T) {

	addr, err := NewAddress("<hoop>")
	if err == nil {
		t.Error("there should be an error:", addr)
	}

	addr, err = NewAddress(`Gogh Fir <tesst@test.com>`)
	if err != nil {
		t.Error("there should be no error:", addr.Host, err)
	}
}
func TestEnvelope(t *testing.T) {
	e := NewEnvelope("127.0.0.1", 22)

	e.QueuedId = "abc123"
	e.Helo = "helo.example.com"
	e.MailFrom = Address{User: "test", Host: "example.com"}
	e.TLS = true
	e.RemoteIP = "222.111.233.121"
	to := Address{User: "test", Host: "example.com"}
	e.PushRcpt(to)
	if to.String() != "test@example.com" {
		t.Error("to does not equal test@example.com, it was:", to.String())
	}
	e.Data.WriteString("Subject: Test\n\nThis is a test nbnb nbnb hgghgh nnnbnb nbnbnb nbnbn.")

	addHead := "Delivered-To: " + to.String() + "\n"
	addHead += "Received: from " + e.Helo + " (" + e.Helo + "  [" + e.RemoteIP + "])\n"
	e.DeliveryHeader = addHead

	r := e.NewReader()

	data, _ := ioutil.ReadAll(r)
	if len(data) != e.Len() {
		t.Error("e.Len() is inccorrect, it shown ", e.Len(), " but we wanted ", len(data))
	}
	e.ParseHeaders()
	if e.Subject != "Test" {
		t.Error("Subject expecting: Test, got:", e.Subject)
	}

}
