package mailbox

import (
	"os"
	"errors"
	"io/ioutil"
)

type Maildirbox struct {
	folder string
	newMailChan <-chan *Email
}

func NewMaildirbox(folder string) *Maildirbox {
	return &Maildirbox{folder: folder, newMailChan: make(chan *Email)}
}

func (m *Maildirbox) MailReceived() <-chan *Email {
	return m.newMailChan
}

func (m *Maildirbox) Mailboxes() []string {
	arr, err := ioutil.ReadDir(m.folder)
	outArr := make([]string, 0)
	if err != nil {
		return outArr
	}
	for _, f := range arr {
		if f.IsDir() {
			n := f.Name()
			if n[0] == '.' {
				outArr = append(outArr, n[1:])
			} else {
				if n == "new" {
					outArr = append(outArr, n)
				}
			}
		}
	}
	return outArr
}

func (m *Maildirbox) RenameMailbox(oldName string, newName string) error {
	fi, err := os.Stat("."+oldName)
	if err != nil {
		return err
	}
	if !fi.IsDir() {
		return errors.New(oldName + " mailbox does not exist")
	}

	err = os.Rename("." + oldName, "." + newName)
	return err
}

func (m *Maildirbox) Search(searchStr string) []*Email {
	outArr := make([]*Email, 0)
	fi, err := os.Stat(".new")
	if err != nil {
		return outArr
	}
	if !fi.IsDir() {
		return outArr
	}
	return outArr // meh
}

