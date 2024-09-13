package service

type ServiceStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	STATUS_OK                               = ServiceStatus{0, "OK"}
	STATUS_DB_ERROR                         = ServiceStatus{10001, "mySQL database query error"}
	STATUS_NO_SCHEME_RECORD                 = ServiceStatus{10002, "no scheme record found in mySQL database"}
	STATUS_NO_APPLICANT_RECORD              = ServiceStatus{10003, "no applicant record found in mySQL database"}
	STATUS_NO_APPLICATION_RECORD            = ServiceStatus{10004, "no application record found in mySQL database"}
	STATUS_INVALID_NRIC                     = ServiceStatus{10005, "applicant's nric is invalid"}
	STATUS_PARAMS_EMPTY                     = ServiceStatus{10006, "all parameters need to be filled up"}
	STATUS_UPDATE_APPLICATION_STATUS_FAILED = ServiceStatus{10007, "update application status failed"}
	STATUS_APPLICATION_ALREADY_APPROVED     = ServiceStatus{10008, "application status already updated to approved"}
	STATUS_APPLICATION_ALREADY_REJECTED     = ServiceStatus{10009, "application status already updated to rejected"}
	STATUS_APPLICATION_ALREADY_WITHDRAWN    = ServiceStatus{10010, "application status already updated to withdrawn"}
	STATUS_APPLICATION_STILL_PENDING        = ServiceStatus{10011, "application status still pending"}
	STATUS_APPLICANT_ID_NOT_INTEGER         = ServiceStatus{10012, "applicant id should be an integer"}
	STATUS_INVALID_NRIC_HOUSEHOLD           = ServiceStatus{10013, "applicant's houusehold nric is invalid"}
	STATUS_NO_ADMIN_RECORD                  = ServiceStatus{10014, "no admin record found in mySQL database"}
	STATUS_PARAMS_ERROR                     = ServiceStatus{10015, "parameters error"}
	STATUS_PASSWORD_ERROR                   = ServiceStatus{10016, "password error"}
	STATUS_HASHING_ERROR                    = ServiceStatus{10017, "hashing of password error"}
	STATUS_GENERATE_TOKEN_ERROR             = ServiceStatus{10018, "generate jwt token error"}
	STATUS_DEL_REDIS_TOKEN_ERROR            = ServiceStatus{10019, "delete redis token error"}
	STATUS_INVALID_TOKEN                    = ServiceStatus{10020, "invalid token"}
	STATUS_NO_TOKEN                         = ServiceStatus{10021, "no token provided"}
	STATUS_INVALID_DOB                      = ServiceStatus{10022, "applicant's date of birth is invalid"}
	STATUS_INVALID_DOB_HOUSEHOLD            = ServiceStatus{10023, "applicant's household date of birth is invalid"}
)

const (
	APPLICATION_PENDING = iota
	APPLICATION_APPROVED
	APPLICATION_REJECTED
	APPLICATION_WITHDRAWN
)
