package message_domain

import "errors"

var ErrMessageNotFound = errors.New("message not found")
var ErrMessageAlreadyBound = errors.New("message already bound to a chat")
