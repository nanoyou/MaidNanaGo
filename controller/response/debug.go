package response

type DebugInfo struct {
	Version   string
	GoVersion string
	QQ        struct {
		Account int64
		Online  bool `json:"online"`
	}
}
