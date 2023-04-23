package threads_controllers

import "gochan/internal/database/threads"

func GetThreadAndPosts(h *threads.Thread, threadId string, limit int32) []map[string]interface{} {
	return h.GetThreadAndPosts(threadId, limit)
}
