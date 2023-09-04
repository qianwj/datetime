package zone

var shortIDs = map[string]*Id{}

var (
	CTT           = &Id{short: "CTT", val: "Asia/Shanghai"}
	EST           = &Id{short: "EST", val: "-05:00"}
	defaultZoneId = &Id{short: "", val: ""}
)

func init() {
	shortIDs[CTT.short] = CTT
	shortIDs[EST.short] = EST
	shortIDs[defaultZoneId.short] = defaultZoneId
}

type Id struct {
	short string
	val   string
}

func DefaultId() *Id {
	return defaultZoneId
}
