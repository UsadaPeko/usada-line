package domain

type MessageBroadcastRetryManager interface {
	RetryBackgroundTask(roomID string, user *User, message string)
}
