package main

// Provider ...
type Provider string

const (
	// Airtel ...
	Airtel Provider = "Airtel"
	// Safaricom ...
	Safaricom = "Safaricom"
	// Faiba ...
	Faiba = "Faiba 4G"

	// Homeland ...
	Homeland = "Homelands Media"

	// MobilePay ...
	MobilePay = "Mobile Pay"

	// Sema ...
	Sema = "Sema Mobile"

	// Telcom ...
	Telcom = "Telkom Kenya"

	// Eferio ...
	Eferio = "Eferio"

	// Equitel ...
	Equitel = "Equitel"
)

// Mobile ...
type Mobile struct {
	Code     int
	Pattern  string
	Number   *string
	Provider Provider
}

func getAllNumbers() []Mobile {
	return []Mobile{
		{10, "^25410", nil, Airtel},
		{11, "^25411", nil, Safaricom},
		{701, "^254701", nil, Safaricom},
		{702, "^254702", nil, Safaricom},
		{703, "^254703", nil, Safaricom},
		{704, "^254704", nil, Safaricom},
		{705, "^254705", nil, Safaricom},
		{706, "^254706", nil, Safaricom},
		{707, "^254707", nil, Safaricom},
		{708, "^254708", nil, Safaricom},
		{709, "^254709", nil, Safaricom},
		{710, "^254710", nil, Safaricom},
		{711, "^254711", nil, Safaricom},
		{712, "^254712", nil, Safaricom},
		{713, "^254713", nil, Safaricom},
		{714, "^254714", nil, Safaricom},
		{715, "^254715", nil, Safaricom},
		{716, "^254716", nil, Safaricom},
		{717, "^254717", nil, Safaricom},
		{718, "^254718", nil, Safaricom},
		{719, "^254719", nil, Safaricom},
		{720, "^254720", nil, Safaricom},
		{721, "^254721", nil, Safaricom},
		{722, "^254722", nil, Safaricom},
		{723, "^254723", nil, Safaricom},
		{724, "^254724", nil, Safaricom},
		{725, "^254725", nil, Safaricom},
		{726, "^254726", nil, Safaricom},
		{727, "^254727", nil, Safaricom},
		{728, "^254728", nil, Safaricom},
		{729, "^254729", nil, Safaricom},
		{730, "^254730", nil, Airtel},
		{731, "^254731", nil, Airtel},
		{732, "^254732", nil, Airtel},
		{733, "^254733", nil, Airtel},
		{734, "^254734", nil, Airtel},
		{735, "^254735", nil, Airtel},
		{736, "^254736", nil, Airtel},
		{737, "^254737", nil, Airtel},
		{738, "^254738", nil, Airtel},
		{739, "^254739", nil, Airtel},
		{740, "^254740", nil, Safaricom},
		{741, "^254741", nil, Safaricom},
		{742, "^254742", nil, Safaricom},
		{743, "^254743", nil, Safaricom},
		{744, "^254744", nil, Homeland},
		{745, "^254745", nil, Safaricom},
		{746, "^254746", nil, Safaricom},
		{747, "^254747", nil, Faiba},
		{748, "^254748", nil, Safaricom},
		{750, "^254750", nil, Airtel},
		{751, "^254751", nil, Airtel},
		{752, "^254752", nil, Airtel},
		{753, "^254753", nil, Airtel},
		{754, "^254754", nil, Airtel},
		{755, "^254755", nil, Airtel},
		{756, "^254756", nil, Airtel},
		{757, "^254757", nil, Safaricom},
		{758, "^254758", nil, Safaricom},
		{759, "^254759", nil, Safaricom},
		{760, "^254760", nil, MobilePay},
		{761, "^254761", nil, Eferio},
		{762, "^254762", nil, Airtel},
		{763, "^254763", nil, Equitel},
		{764, "^254764", nil, Equitel},
		{765, "^254765", nil, Equitel},
		{766, "^254766", nil, Equitel},
		{767, "^254767", nil, Sema},
		{768, "^254768", nil, Safaricom},
		{769, "^254769", nil, Safaricom},
		{770, "^254770", nil, Telcom},
		{771, "^254771", nil, Telcom},
		{772, "^254772", nil, Telcom},
		{773, "^254773", nil, Telcom},
		{774, "^254774", nil, Telcom},
		{775, "^254775", nil, Telcom},
		{776, "^254776", nil, Telcom},
		{777, "^254777", nil, Telcom},
		{778, "^254778", nil, Telcom},
		{779, "^254779", nil, Telcom},
		{780, "^254780", nil, Airtel},
		{781, "^254781", nil, Airtel},
		{782, "^254782", nil, Airtel},
		{783, "^254783", nil, Airtel},
		{784, "^254784", nil, Airtel},
		{785, "^254785", nil, Airtel},
		{786, "^254786", nil, Airtel},
		{787, "^254787", nil, Airtel},
		{788, "^254788", nil, Airtel},
		{789, "^254789", nil, Airtel},
		{790, "^254790", nil, Safaricom},
		{791, "^254791", nil, Safaricom},
		{792, "^254792", nil, Safaricom},
		{793, "^254793", nil, Safaricom},
		{794, "^254794", nil, Safaricom},
		{795, "^254795", nil, Safaricom},
		{796, "^254796", nil, Safaricom},
		{797, "^254797", nil, Safaricom},
		{798, "^254798", nil, Safaricom},
		{799, "^254799", nil, Safaricom},
	}
}
