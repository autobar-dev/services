package types

import "github.com/autobar-dev/services/module/repositories"

const (
	RequestReportCommandName repositories.CommandName = "request-report"
)

type RequestReportCommandArgs struct {
	Queue string `json:"queue"`
}
