package userdata

const (
	Photo = 0
	Video = 1
	Audio = 2
	Document = 3
	Wall = 4
	Market = 5
	Poll = 5
)

type User struct {
	Id int64
	IsUser bool
	Chat *Chat
}

type Chat struct {
	Id int64
}

type Attachment struct {
	Type int
	Data string
}


func encodeAttachments(attachments []*Attachment) string {
	att := ""
	for _, attachment := range attachments {
		att += attachment.Data + ","
	}

	return att[:len(att) - 1]
}