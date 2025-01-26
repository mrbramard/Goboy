package cartridge

import "fmt"

func checkNintendoLogo(logo []byte) bool {
	expectedLogo := []byte{
		0xCE, 0xED, 0x66, 0x66, 0xCC, 0x0D, 0x00, 0x0B, 0x03, 0x73, 0x00, 0x83, 0x00, 0x0C, 0x00, 0x0D,
		0x00, 0x08, 0x11, 0x1F, 0x88, 0x89, 0x00, 0x0E, 0xDC, 0xCC, 0x6E, 0xE6, 0xDD, 0xDD, 0xD9, 0x99,
		0xBB, 0xBB, 0x67, 0x63, 0x6E, 0x0E, 0xEC, 0xCC, 0xDD, 0xDC, 0x99, 0x9F, 0xBB, 0xB9, 0x33, 0x3E,
	}

	if len(logo) != len(expectedLogo) {
		return false
	}

	for i := range logo {
		if logo[i] != expectedLogo[i] {
			return false
		}
	}

	return true
}

func isSGBCompatible(sgbFlag byte) bool {
	return sgbFlag == 0x03
}

var cartridgeTypes = map[byte]string{
	0x00: "ROM ONLY",
	0x01: "MBC1",
	0x02: "MBC1+RAM",
	0x03: "MBC1+RAM+BATTERY",
	0x05: "MBC2",
	0x06: "MBC2+BATTERY",
	0x08: "ROM+RAM",
	0x09: "ROM+RAM+BATTERY",
	0x0B: "MMM01",
	0x0C: "MMM01+RAM",
	0x0D: "MMM01+RAM+BATTERY",
	0x0F: "MBC3+TIMER+BATTERY",
	0x10: "MBC3+TIMER+RAM+BATTERY",
	0x11: "MBC3",
	0x12: "MBC3+RAM",
	0x13: "MBC3+RAM+BATTERY",
	0x19: "MBC5",
	0x1A: "MBC5+RAM",
	0x1B: "MBC5+RAM+BATTERY",
	0x1C: "MBC5+RUMBLE",
	0x1D: "MBC5+RUMBLE+RAM",
	0x1E: "MBC5+RUMBLE+RAM+BATTERY",
	0x20: "MBC6",
	0x22: "MBC7+SENSOR+RUMBLE+RAM+BATTERY",
	0xFC: "POCKET CAMERA",
	0xFD: "BANDAI TAMA5",
	0xFE: "HuC3",
	0xFF: "HuC1+RAM+BATTERY",
}

func getCartridgeType(cartridgeType byte) string {
	if name, exists := cartridgeTypes[cartridgeType]; exists {
		return name
	}
	return "Unknown"
}

var licenseeCodes = map[string]string{
	"00": "None",
	"01": "Nintendo Research & Development 1",
	"08": "Capcom",
	"13": "EA (Electronic Arts)",
	"18": "Hudson Soft",
	"19": "B-AI",
	"20": "KSS",
	"22": "Planning Office WADA",
	"24": "PCM Complete",
	"25": "San-X",
	"28": "Kemco",
	"29": "SETA Corporation",
	"30": "Viacom",
	"31": "Nintendo",
	"32": "Bandai",
	"33": "Ocean Software/Acclaim Entertainment",
	"34": "Konami",
	"35": "HectorSoft",
	"37": "Taito",
	"38": "Hudson Soft",
	"39": "Banpresto",
	"41": "Ubi Soft",
	"42": "Atlus",
	"44": "Malibu Interactive",
	"46": "Angel",
	"47": "Bullet-Proof Software",
	"49": "Irem",
	"50": "Absolute",
	"51": "Acclaim Entertainment",
	"52": "Activision",
	"53": "Sammy USA Corporation",
	"54": "Konami",
	"55": "Hi Tech Expressions",
	"56": "LJN",
	"57": "Matchbox",
	"58": "Mattel",
	"59": "Milton Bradley Company",
	"60": "Titus Interactive",
	"61": "Virgin Games Ltd.",
	"64": "Lucasfilm Games",
	"67": "Ocean Software",
	"69": "EA (Electronic Arts)",
	"70": "Infogrames",
	"71": "Interplay Entertainment",
	"72": "Broderbund",
	"73": "Sculptured Software",
	"75": "The Sales Curve Limited",
	"78": "THQ",
	"79": "Accolade",
	"80": "Misawa Entertainment",
	"83": "lozc",
	"86": "Tokuma Shoten",
	"87": "Tsukuda Original",
	"91": "Chunsoft Co.",
	"92": "Video System",
	"93": "Ocean Software/Acclaim Entertainment",
	"95": "Varie",
	"96": "Yonezawa/s’pal",
	"97": "Kaneko",
	"99": "Pack-In-Video",
	"9H": "Bottom Up",
	"A4": "Konami (Yu-Gi-Oh!)",
	"BL": "MTO",
	"DK": "Kodansha",
}

func getNewLicenseeCode(licenseeCode string) string {
	if name, exists := licenseeCodes[licenseeCode]; exists {
		return name
	}
	return "Unknown"
}

var oldLicenseeCode = map[byte]string{
	0x00: "None",
	0x01: "Nintendo",
	0x08: "Capcom",
	0x09: "HOT-B",
	0x0A: "Jaleco",
	0x0B: "Coconuts Japan",
	0x0C: "Elite Systems",
	0x13: "EA (Electronic Arts)",
	0x18: "Hudson Soft",
	0x19: "ITC Entertainment",
	0x1A: "Yanoman",
	0x1D: "Japan Clary",
	0x1F: "Virgin Games Ltd.3",
	0x24: "PCM Complete",
	0x25: "San-X",
	0x28: "Kemco",
	0x29: "SETA Corporation",
	0x30: "Infogrames5",
	0x31: "Nintendo",
	0x32: "Bandai",
	0x33: "Indicates that the New licensee code should be used instead.",
	0x34: "Konami",
	0x35: "HectorSoft",
	0x38: "Capcom",
	0x39: "Banpresto",
	0x3C: "Entertainment Interactive (stub)",
	0x3E: "Gremlin",
	0x41: "Ubi Soft1",
	0x42: "Atlus",
	0x44: "Malibu Interactive",
	0x46: "Angel",
	0x47: "Spectrum HoloByte",
	0x49: "Irem",
	0x4A: "Virgin Games Ltd.3",
	0x4D: "Malibu Interactive",
	0x4F: "U.S. Gold",
	0x50: "Absolute",
	0x51: "Acclaim Entertainment",
	0x52: "Activision",
	0x53: "Sammy USA Corporation",
	0x54: "GameTek",
	0x55: "Park Place13",
	0x56: "LJN",
	0x57: "Matchbox",
	0x59: "Milton Bradley Company",
	0x5A: "Mindscape",
	0x5B: "Romstar",
	0x5C: "Naxat Soft14",
	0x5D: "Tradewest",
	0x60: "Titus Interactive",
	0x61: "Virgin Games Ltd.3",
	0x67: "Ocean Software",
	0x69: "EA (Electronic Arts)",
	0x6E: "Elite Systems",
	0x6F: "Electro Brain",
	0x70: "Infogrames5",
	0x71: "Interplay Entertainment",
	0x72: "Broderbund",
	0x73: "Sculptured Software6",
	0x75: "The Sales Curve Limited7",
	0x78: "THQ",
	0x79: "Accolade15",
	0x7A: "Triffix Entertainment",
	0x7C: "MicroProse",
	0x7F: "Kemco",
	0x80: "Misawa Entertainment",
	0x83: "LOZC G.",
	0x86: "Tokuma Shoten",
	0x8B: "Bullet-Proof Software2",
	0x8C: "Vic Tokai Corp.16",
	0x8E: "Ape Inc.17",
	0x8F: "I’Max18",
	0x91: "Chunsoft Co.8",
	0x92: "Video System",
	0x93: "Tsubaraya Productions",
	0x95: "Varie",
	0x96: "Yonezawa19/S’Pal",
	0x97: "Kemco",
	0x99: "Arc",
	0x9A: "Nihon Bussan",
	0x9B: "Tecmo",
	0x9C: "Imagineer",
	0x9D: "Banpresto",
	0x9F: "Nova",
	0xA1: "Hori Electric",
	0xA2: "Bandai",
	0xA4: "Konami",
	0xA6: "Kawada",
	0xA7: "Takara",
	0xA9: "Technos Japan",
	0xAA: "Broderbund",
	0xAC: "Toei Animation",
	0xAD: "Toho",
	0xAF: "Namco",
	0xB0: "Acclaim Entertainment",
	0xB1: "ASCII Corporation or Nexsoft",
	0xB2: "Bandai",
	0xB4: "Square Enix",
	0xB6: "HAL Laboratory",
	0xB7: "SNK",
	0xB9: "Pony Canyon",
	0xBA: "Culture Brain",
	0xBB: "Sunsoft",
	0xBD: "Sony Imagesoft",
	0xBF: "Sammy Corporation",
	0xC0: "Taito",
	0xC2: "Kemco",
	0xC3: "Square",
	0xC4: "Tokuma Shoten",
	0xC5: "cartridgea East",
	0xC6: "Tonkin House",
	0xC8: "Koei",
	0xC9: "UFL",
	0xCA: "Ultra Games",
	0xCB: "VAP, Inc.",
	0xCC: "Use Corporation",
	0xCD: "Meldac",
	0xCE: "Pony Canyon",
	0xCF: "Angel",
	0xD0: "Taito",
	0xD1: "SOFEL (Software Engineering Lab",
	0xD2: "Quest",
	0xD3: "Sigma Enterprises",
	0xD4: "ASK Kodansha Co.",
	0xD6: "Naxat Soft14",
	0xD7: "Copya System",
	0xD9: "Banpresto",
	0xDA: "Tomy",
	0xDB: "LJN",
	0xDD: "Nippon Computer Systems",
	0xDE: "Human Ent.",
	0xDF: "Altron",
	0xE0: "Jaleco",
	0xE1: "Towa Chiki",
	0xE2: "Yutaka",
	0xE3: "Varie",
	0xE5: "Epoch",
	0xE7: "Athena",
	0xE8: "Asmik Ace Entertainment",
	0xE9: "Natsume",
	0xEA: "King Records",
	0xEB: "Atlus",
	0xEC: "Epic/Sony Records",
	0xEE: "IGS",
	0xF0: "A Wave",
	0xF3: "Extreme Entertainment",
	0xFF: "LJN",
}

func getOldLicenseeCode(licenseeCode byte) string {
	if name, exists := oldLicenseeCode[licenseeCode]; exists {
		return name
	}
	return "Unknown"
}

func PrintHeader(cartridge []byte) {
	entryPoint := cartridge[0x100:0x104]
	nintendoLogo := cartridge[0x104:0x134]
	title := string(cartridge[0x134:0x144])

	// Only used in newer cartridges
	manufacturerCode := cartridge[0x13f:0x143]
	cgbFlag := cartridge[0x143]

	newLicenseeCode := string(cartridge[0x144:0x146])
	sgbFlag := cartridge[0x146]
	cartrigdeType := cartridge[0x147]
	romSize := cartridge[0x148]
	ramSize := cartridge[0x149]
	destinationCode := cartridge[0x14a]
	oldLicenseeCode := cartridge[0x14b]
	maskRomVersion := cartridge[0x14c]
	headerChecksum := cartridge[0x14d]
	globalChecksum := cartridge[0x14e:0x150]

	fmt.Println(fmt.Sprintf("%X", entryPoint))
	fmt.Println(fmt.Sprintf("%X", nintendoLogo) + " " + fmt.Sprintf("%t", checkNintendoLogo(nintendoLogo)))
	fmt.Println(title)
	fmt.Println(fmt.Sprintf("%X", manufacturerCode))
	fmt.Println(fmt.Sprintf("%X", cgbFlag))
	fmt.Println(newLicenseeCode + " " + getNewLicenseeCode(newLicenseeCode))
	fmt.Println(fmt.Sprintf("%X", sgbFlag) + " " + fmt.Sprintf("%t", isSGBCompatible(sgbFlag)))
	fmt.Println(fmt.Sprintf("%X", cartrigdeType) + " " + getCartridgeType(cartrigdeType))
	fmt.Println(fmt.Sprintf("%X", romSize))
	fmt.Println(fmt.Sprintf("%X", ramSize))
	fmt.Println(fmt.Sprintf("%X", destinationCode))
	fmt.Println(fmt.Sprintf("%X", oldLicenseeCode) + " " + getOldLicenseeCode(oldLicenseeCode))
	fmt.Println(fmt.Sprintf("%X", maskRomVersion))
	fmt.Println(fmt.Sprintf("%X", headerChecksum))
	fmt.Println(fmt.Sprintf("%X", globalChecksum))
}
