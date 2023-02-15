package notification

import (
	"fmt"
	"strings"
)

type NotificationErrorProps struct {
	Message []string
	Context string
}

type Notification struct {
	errors []*NotificationErrorProps
}

func NewNotification() *Notification {
	return &Notification{}
}

func (n *Notification) Errors() []*NotificationErrorProps {
	return n.errors
}

func (n *Notification) Messages(context string) string {
	for _, err := range n.errors {
		if err.Context == context {
			return fmt.Sprintf("%s: %s", context, strings.Join(err.Message, ","))
		}
	}

	return ""
}

func (n *Notification) AddError(message string, context string) {
	hasContext := false
	for _, e := range n.errors {
		if e.Context == context {
			e.Message = append(e.Message, message)
			hasContext = true
		}
	}

	if !hasContext {
		n.errors = append(n.errors, &NotificationErrorProps{
			Message: []string{message},
			Context: context,
		})
	}
}

func (n *Notification) HasErrors() bool {
	return len(n.errors) > 0
}

func (n *Notification) Clear() {
	n.errors = []*NotificationErrorProps{}
}
