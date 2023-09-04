package zone

var shortIDs = map[string]*Id{}

var (
	EST           = &Id{short: "EST", val: "-05:00"}
	HST           = &Id{short: "HST", val: "-10:00"}
	MST           = &Id{short: "MST", val: "-07:00"}
	ACT           = &Id{short: "ACT", val: "Australia/Darwin"}
	AET           = &Id{short: "AET", val: "Australia/Sydney"}
	AGT           = &Id{short: "AGT", val: "America/Argentina/Buenos_Aires"}
	ART           = &Id{short: "ART", val: "Africa/Cairo"}
	AST           = &Id{short: "AST", val: "America/Anchorage"}
	BET           = &Id{short: "BET", val: "America/Sao_Paulo"}
	BST           = &Id{short: "BST", val: "Asia/Dhaka"}
	CAT           = &Id{short: "CAT", val: "Africa/Harare"}
	CNT           = &Id{short: "CNT", val: "America/St_Johns"}
	CST           = &Id{short: "CST", val: "America/Chicago"}
	CTT           = &Id{short: "CTT", val: "Asia/Shanghai"}
	EAT           = &Id{short: "EAT", val: "Africa/Addis_Ababa"}
	ECT           = &Id{short: "ECT", val: "Europe/Paris"}
	IET           = &Id{short: "IET", val: "America/Indiana/Indianapolis"}
	IST           = &Id{short: "IST", val: "Asia/Kolkata"}
	JST           = &Id{short: "JST", val: "Asia/Tokyo"}
	MIT           = &Id{short: "MIT", val: "Pacific/Apia"}
	NET           = &Id{short: "NET", val: "Asia/Yerevan"}
	NST           = &Id{short: "NST", val: "Pacific/Auckland"}
	PLT           = &Id{short: "PLT", val: "Asia/Karachi"}
	PNT           = &Id{short: "PNT", val: "America/Phoenix"}
	PRT           = &Id{short: "PRT", val: "America/Puerto_Rico"}
	PST           = &Id{short: "PST", val: "America/Los_Angeles"}
	SST           = &Id{short: "SST", val: "Pacific/Guadalcanal"}
	VST           = &Id{short: "VST", val: "Asia/Ho_Chi_Minh"}
	defaultZoneId = &Id{short: "", val: ""}
)

func init() {
	shortIDs[EST.short] = EST
	shortIDs[HST.short] = HST
	shortIDs[MST.short] = MST
	shortIDs[ACT.short] = ACT
	shortIDs[AET.short] = AET
	shortIDs[AGT.short] = AGT
	shortIDs[ART.short] = ART
	shortIDs[AST.short] = AST
	shortIDs[BET.short] = BET
	shortIDs[BST.short] = BST
	shortIDs[CAT.short] = CAT
	shortIDs[CNT.short] = CNT
	shortIDs[CST.short] = CST
	shortIDs[CTT.short] = CTT
	shortIDs[EAT.short] = EAT
	shortIDs[ECT.short] = ECT
	shortIDs[IET.short] = IET
	shortIDs[IST.short] = IST
	shortIDs[JST.short] = JST
	shortIDs[MIT.short] = MIT
	shortIDs[NET.short] = NET
	shortIDs[NST.short] = NST
	shortIDs[PLT.short] = PLT
	shortIDs[PNT.short] = PNT
	shortIDs[PRT.short] = PRT
	shortIDs[PST.short] = PST
	shortIDs[SST.short] = SST
	shortIDs[VST.short] = VST

	shortIDs[defaultZoneId.short] = defaultZoneId
}

type Id struct {
	short string
	val   string
}

func DefaultId() *Id {
	return defaultZoneId
}
