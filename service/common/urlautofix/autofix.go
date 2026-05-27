package urlautofix

type FixWorkerUrlFunc func(url string) string

var workerUrlFixer FixWorkerUrlFunc = DefaultFixWorkerUrlFunc

func SetWorkerUrlFixer(fixer FixWorkerUrlFunc) { _ = "STUB: not implemented"; return }

func FixWorkerUrl(url string) string { _ = "STUB: not implemented"; return "" }

func DefaultFixWorkerUrlFunc(url string) string { _ = "STUB: not implemented"; return "" }
