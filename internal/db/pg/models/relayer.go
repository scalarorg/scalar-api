package models

type RelayDataStatus string

const (
	Pending   RelayDataStatus = "pending"
	Approved  RelayDataStatus = "approved"
	Success   RelayDataStatus = "success"
	Failed    RelayDataStatus = "failed"
	Undefined RelayDataStatus = "undefined"
)

func ToReadableStatus(status int) RelayDataStatus {
	switch status {
	case 0:
		return Pending
	case 1:
		return Approved
	case 2:
		return Success
	case 3:
		return Failed
	default:
		return Undefined
	}
}
