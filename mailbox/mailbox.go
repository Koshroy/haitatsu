package mailbox

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




