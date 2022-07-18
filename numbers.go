package main

import "regexp"

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
	Regex    string
	Provider Provider
}

// MatchMobile ...
func MatchMobile(mobile string) *Provider {
	for _, m := range AllNumbers() {
		if m.Regex != "" && regexp.MustCompile(m.Regex).MatchString(mobile) {
			return &m.Provider
		}
	}

	return nil
}

// AllNumbers ...
func AllNumbers() []Mobile {
	return []Mobile{
		{10, "^25410", Airtel},
		{11, "^25411", Safaricom},
		{701, "^254701", Safaricom},
		{702, "^254702", Safaricom},
		{703, "^254703", Safaricom},
		{704, "^254704", Safaricom},
		{705, "^254705", Safaricom},
		{706, "^254706", Safaricom},
		{707, "^254707", Safaricom},
		{708, "^254708", Safaricom},
		{709, "^254709", Safaricom},
		{710, "^254710", Safaricom},
		{711, "^254711", Safaricom},
		{712, "^254712", Safaricom},
		{713, "^254713", Safaricom},
		{714, "^254714", Safaricom},
		{715, "^254715", Safaricom},
		{716, "^254716", Safaricom},
		{717, "^254717", Safaricom},
		{718, "^254718", Safaricom},
		{719, "^254719", Safaricom},
		{720, "^254720", Safaricom},
		{721, "^254721", Safaricom},
		{722, "^254722", Safaricom},
		{723, "^254723", Safaricom},
		{724, "^254724", Safaricom},
		{725, "^254725", Safaricom},
		{726, "^254726", Safaricom},
		{727, "^254727", Safaricom},
		{728, "^254728", Safaricom},
		{729, "^254729", Safaricom},
		{730, "^254730", Airtel},
		{731, "^254731", Airtel},
		{732, "^254732", Airtel},
		{733, "^254733", Airtel},
		{734, "^254734", Airtel},
		{735, "^254735", Airtel},
		{736, "^254736", Airtel},
		{737, "^254737", Airtel},
		{738, "^254738", Airtel},
		{739, "^254739", Airtel},
		{740, "^254740", Safaricom},
		{741, "^254741", Safaricom},
		{742, "^254742", Safaricom},
		{743, "^254743", Safaricom},
		{744, "^254744", Homeland},
		{745, "^254745", Safaricom},
		{746, "^254746", Safaricom},
		{747, "^254747", Faiba},
		{748, "^254748", Safaricom},
		{750, "^254750", Airtel},
		{751, "^254751", Airtel},
		{752, "^254752", Airtel},
		{753, "^254753", Airtel},
		{754, "^254754", Airtel},
		{755, "^254755", Airtel},
		{756, "^254756", Airtel},
		{757, "^254757", Safaricom},
		{758, "^254758", Safaricom},
		{759, "^254759", Safaricom},
		{760, "^254760", MobilePay},
		{761, "^254761", Eferio},
		{762, "^254762", Airtel},
		{763, "^254763", Equitel},
		{764, "^254764", Equitel},
		{765, "^254765", Equitel},
		{766, "^254766", Equitel},
		{767, "^254767", Sema},
		{768, "^254768", Safaricom},
		{769, "^254769", Safaricom},
		{770, "^254770", Telcom},
		{771, "^254771", Telcom},
		{772, "^254772", Telcom},
		{773, "^254773", Telcom},
		{774, "^254774", Telcom},
		{775, "^254775", Telcom},
		{776, "^254776", Telcom},
		{777, "^254777", Telcom},
		{778, "^254778", Telcom},
		{779, "^254779", Telcom},
		{780, "^254780", Airtel},
		{781, "^254781", Airtel},
		{782, "^254782", Airtel},
		{783, "^254783", Airtel},
		{784, "^254784", Airtel},
		{785, "^254785", Airtel},
		{786, "^254786", Airtel},
		{787, "^254787", Airtel},
		{788, "^254788", Airtel},
		{789, "^254789", Airtel},
		{790, "^254790", Safaricom},
		{791, "^254791", Safaricom},
		{792, "^254792", Safaricom},
		{793, "^254793", Safaricom},
		{794, "^254794", Safaricom},
		{795, "^254795", Safaricom},
		{796, "^254796", Safaricom},
		{797, "^254797", Safaricom},
		{798, "^254798", Safaricom},
		{799, "^254799", Safaricom},
	}
}
