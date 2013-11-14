package mailbox

import (
	"errors"
	"io"
)

type Mailbox interface {
	MailReceived() <-chan *Email // new *Email sent when new email is received
	Mailboxes() []string // names of available mailboxes
	RenameMailbox(mboxName string, newName string) error // renames mailbox
	Search(searchStr string) []*Email // returns list of Email that matches seacrh string
}

type Email struct {
	Header map[string]string
	Body   string
}

type ParserState int

const (
	READ_SZ = 512

	HEADER_DEF_STATE = iota
	HEADER_CONTENTS_STATE
	HEADER_NEWLINE
)

type EmailParser struct {
	state ParserState
	keyword []byte
	tokenQ *queue
}

func NewEmailParser() *EmailParser {
	return &EmailParser{state: HEADER_DEF_STATE, tokenQ: newQueue()}
}

func (e *EmailParser) Parse(r io.Reader) error {
	buffer := make([]byte, READ_SZ)
	n, err := r.Read(buffer)
	for ; ;  {
		for i := 0; i < n; i++ {
			ch := buffer[i]

			switch(e.state) {
			case HEADER_DEF_STATE:
				switch(ch) {
				case ':':
					a := e.tokenQ.Clear()
					e.keyword = make([]byte, len(a))
					for i, v := range a {
						var found bool
						e.keyword[i], found = v.(byte)
						if found == false {
							return errors.New("error using parser queue")
						}
					}
					e.state = HEADER_CONTENTS_STATE
				case '\n':
					e.state = HEADER_NEWLINE
				default:
					e.tokenQ.Push(ch)
				}
			default:
				break
			}
		}
		if err != nil {
			break
		}
	}
	return nil
}