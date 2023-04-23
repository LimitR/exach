package threads_controllers

import "gochan/internal/database/threads"

func GetAllThreads(h *threads.Thread, limit int32) []map[string]interface{} {
	return h.GetThreads(limit)
}
