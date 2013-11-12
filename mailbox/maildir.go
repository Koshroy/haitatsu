package mailbox

import (
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

