package txt

var SpecialWords = map[string]string{
	"av":                 "AV",
	"xxs":                "XXS",
	"xs":                 "XS",
	"s":                  "S",
	"sm":                 "SM",
	"r":                  "R",
	"m":                  "M",
	"md":                 "MD",
	"mm":                 "mm",
	"cm":                 "cm",
	"km":                 "km",
	"l":                  "L",
	"ii":                 "II",
	"iii":                "III",
	"ti":                 "Ti",
	"xl":                 "XL",
	"xxl":                "XXL",
	"xxxl":               "XXXL",
	"xxxxl":              "XXXXL",
	"xtc":                "XTC",
	"rpg":                "RPG",
	"rtx":                "RTX",
	"rty":                "RTY",
	"nq":                 "NQ",
	"nqq":                "NQQ",
	"ym":                 "YM",
	"cme":                "CME",
	"esx":                "ESX",
	"nwog":               "NWOG",
	"ndog":               "NDOG",
	"ict":                "ICT",
	"smt":                "SMT",
	"rth":                "RTH",
	"tradingview":        "TradingView",
	"ftse":               "FTSE",
	"usd":                "USD",
	"usdt":               "USDT",
	"gbp":                "GBP",
	"chf":                "CHF",
	"btc":                "BTC",
	"eth":                "ETH",
	"bnb":                "BNB",
	"xd":                 "XD",
	"dx":                 "DX",
	"ds":                 "DS",
	"hd":                 "HD",
	"uk":                 "UK",
	"nyc":                "NYC",
	"ny":                 "NY",
	"uae":                "UAE",
	"usa":                "USA",
	"amd":                "AMD",
	"tiff":               "TIFF",
	"tga":                "TGA",
	"eos":                "EOS",
	"slr":                "SLR",
	"dslr":               "DSLR",
	"dslrs":              "DSLRs",
	"ibm":                "IBM",
	"ios":                "iOS",
	"bvg":                "BVG",
	"bsr":                "BSR",
	"ceo":                "CEO",
	"cto":                "CTO",
	"cfo":                "CFO",
	"cia":                "CIA ",
	"fbi":                "FBI",
	"bnd":                "BND",
	"fsb":                "FSB",
	"nsa":                "NSA",
	"nas":                "NAS",
	"ssd":                "SSD",
	"raid":               "RAID",
	"nasa":               "NASA",
	"esa":                "ESA",
	"csa":                "CSA",
	"hst":                "HST",
	"stsci":              "STScI",
	"jwst":               "JWST",
	"nircam":             "NIRCam",
	"lax":                "LAX",
	"sxf":                "SXF",
	"ber":                "BER",
	"sfo":                "SFO",
	"lh":                 "LH",
	"lhr":                "LHR",
	"afl":                "AFL",
	"nrl":                "NRL",
	"nsw":                "NSW",
	"qld":                "QLD",
	"vic":                "VIC",
	"aws":                "AWS",
	"cpu":                "CPU",
	"ram":                "RAM",
	"rgb":                "RGB",
	"srgb":               "SRGB",
	"hdtv":               "HDTV",
	"dci":                "DCI",
	"xmp":                "XMP",
	"xml":                "XML",
	"svg":                "SVG",
	"ai":                 "AI",
	"psd":                "PSD",
	"pdf":                "PDF",
	"pcx":                "PCX",
	"sct":                "SCT",
	"bpm":                "BPM",
	"ps":                 "PS",
	"postscript":         "PostScript",
	"proraw":             "ProRAW",
	"gopro":              "GoPro",
	"eps":                "EPS",
	"mov":                "MOV",
	"avc":                "AVC",
	"hvc":                "HVC",
	"hevc":               "HEVC",
	"heif":               "HEIF",
	"heic":               "HEIC",
	"avif":               "AVIF",
	"bmp":                "BMP",
	"gif":                "GIF",
	"dng":                "DNG",
	"wto":                "WTO",
	"fifa":               "FIFA",
	"yaml":               "YAML",
	"json":               "JSON",
	"uhd":                "UHD",
	"uav":                "UAV",
	"cmy":                "CMY",
	"cmyk":               "CMYK",
	"cdu":                "CDU",
	"spd":                "SPD",
	"fdp":                "FDP",
	"afd":                "AFD",
	"hsl":                "HSL",
	"hsv":                "HSV",
	"lcd":                "LCD",
	"oled":               "OLED",
	"fullhd":             "FullHD",
	"ultrahd":            "UltraHD",
	"uwqhd":              "UWQHD",
	"unrwa":              "UNRWA",
	"nvidia":             "NVIDIA",
	"geforce":            "GeForce",
	"iphone":             "iPhone",
	"imac":               "iMac",
	"ipad":               "iPad",
	"ipod":               "iPod",
	"macbook":            "MacBook",
	"airplay":            "AirPlay",
	"airpods":            "AirPods",
	"youtube":            "YouTube",
	"github":             "GitHub",
	"tensorflow":         "TensorFlow",
	"digitalocean":       "DigitalOcean",
	"photosync":          "PhotoSync",
	"photoprism":         "PhotoPrism",
	"macgyver":           "MacGyver",
	"o'brien":            "O'Brien",
	"mcgregor":           "McGregor",
	"mcdonald":           "McDonald",
	"mcdonalds":          "McDonald's",
	"mcdonald's":         "McDonald's",
	"macalister":         "MacAlister",
	"mcalister":          "McAlister",
	"mcallister":         "McAllister",
	"macauley":           "MacAuley",
	"mccauley":           "McCauley",
	"mcawley":            "McAwley",
	"macauliffe":         "MacAuliffe",
	"macbride":           "MacBride",
	"mcbride":            "McBride",
	"maccabe":            "MacCabe",
	"mccabe":             "McCabe",
	"maccann":            "MacCann",
	"mccann":             "McCann",
	"maccarthy":          "MacCarthy",
	"mccarthy":           "McCarthy",
	"maccormack":         "MacCormack",
	"mccormick":          "McCormick",
	"maccullagh":         "MacCullagh",
	"macnully":           "MacNully",
	"mackenna":           "MacKenna",
	"macnamara":          "MacNamara",
	"mcnamara":           "McNamara",
	"gelaende":           "Gelände",
	"schwaebisch":        "Schwäbisch",
	"schwaebische":       "Schwäbische",
	"aegypten":           "Ägypten",
	"muenchen":           "München",
	"wuerttemberg":       "Württemberg",
	"baden-wuerttemberg": "Baden-Württemberg",
	"nuernberg":          "Nürnberg",
	"wuerzburg":          "Würzburg",
	"tubingen":           "Tübingen",
	"tuebingen":          "Tübingen",
	"koeln":              "Köln",
	"oesterreich":        "Österreich",
	"woerthersee":        "Wörthersee",
	"oeland":             "Öland",
	"schoenefeld":        "Schönefeld",
	"duesseldorf":        "Düsseldorf",
	"dusseldorf":         "Düsseldorf",
	"saarbrucken":        "Saarbrücken",
	"saarbruecken":       "Saarbrücken",
	"zuerich":            "Zürich",
	"bverfge":            "BVerfGE",
	"bverwg":             "BVerwG",
	"bhg":                "BGH",
	"olg":                "OLG",
	"lg":                 "LG",
	"ag":                 "AG",
	"ug":                 "UG",
	"gmbh":               "GmbH",
	"eugh":               "EuGH",
	"ic":                 "IC",
	"icc":                "ICC",
	"iaf":                "IAF",
	"idf":                "IDF",
	"smd":                "SMD",
	"cmos":               "CMOS",
	"mosfet":             "MOSFET",
	"finfet":             "FinFET",
	"igh":                "IGH",
	"istgh":              "IStGH",
	"wpo":                "WPO",
	"wg":                 "WG",
	"sso":                "SSO",
	"sdk":                "SDK",
	"sos":                "SOS",
	"api":                "API",
	"oidc":               "OIDC",
	"openid":             "OpenID",
	"oauth":              "OAuth",
	"webdav":             "WebDAV",
	"webp":               "WebP",
	"webm":               "WebM",
	"whatsapp":           "WhatsApp",
	"wechat":             "WeChat",
	"chatgpt":            "ChatGPT",
	"goland":             "GoLand",
	"phpstorm":           "PhpStorm",
	"vuejs":              "VueJS",
	"nodejs":             "NodeJS",
	"d'honneur":          "d'Honneur",
	"île-de-france":      "Île-de-France",
	"ile-de-france":      "Île-de-France",
}
