package response

type DashboardList struct {
	DataType  string `json:"dataType"`
	DataName  string `json:"dataName"`
	DataCount int64  `json:"dataCount"`
	Icon      string `json:"icon"`
	Path      string `json:"path"`
}
