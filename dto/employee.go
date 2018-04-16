package dto

type CreateEmployeeResponse struct {
	UID string `json:"Uid"`
}

type CreateEmployeeRequest struct {
	PersonUID      string `json:"PersonUid"`
	WorkDocumentID string `json:"WorkDocumentId"`
}

type UpdateEmployeeRequest struct {
	WorkDocumentID string `json:"WorkDocumentId"`
}

type GetEmployeeResponse struct {
	UID            string `json:"Uid"`
	PersonUID      string `json:"PersonUid"`
	PersonName     string
	PersonSurname  string
	WorkDocumentID string `json:"WorkDocumentId"`
}

type GetEmployeesResponse struct {
	Employees []*GetEmployeeResponse
}
