package chatwork

import (
	"fmt"
	"testing"
)

func rejectFmt(t *testing.T, a interface{}, b interface{}) {
	va := fmt.Sprintf("%+v", a)
	vb := fmt.Sprintf("%+v", b)

	if va == vb {
		t.Errorf("Struct has unexpected shape: %s", vb)
	}
}

func TestMeEndpoint(t *testing.T) {
	c := NewClient(apiKey)
	if me, err := c.Me(); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Me: %+v", me)
		rejectFmt(t, me, &Me{})
	}
}

func TestMyEndpoint(t *testing.T) {
	c := NewClient(apiKey)
	if s, err := c.MyStatus(); err != nil {
		t.Fatal("/my/tasks", err)
	} else {
		refute(t, s, nil)
	}

	if tasks, err := c.MyTasks(&Params{"status": "done"}); err != nil {
		t.Fatal("/my/tasks", err)
	} else {
		refute(t, tasks, nil)
	}
}

func TestContactsEndpoint(t *testing.T) {
	c := NewClient(apiKey)
	if contacts, err := c.Contacts(); err != nil {
		t.Fatal("/contacts", err)
	} else {
		t.Logf("Contacts: %+v", contacts)
		refute(t, contacts, nil)
	}
}
