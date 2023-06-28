package types

import "go.a5r.dev/services/module/repositories"

const (
	RequestReportCommandName repositories.CommandName = "request-report"
)

type RequestReportCommandArgs struct {
	Queue string `json:"queue"`
}
