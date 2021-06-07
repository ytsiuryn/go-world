package world

import "strings"

// https://www.loc.gov/standards/iso639-2/php/code_list.php
// https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes

// https://en.wikipedia.org/wiki/List_of_official_languages_by_country_and_territory

// TODO: всегда синхронизировать с БД!

// Languages - перечень языков мира.
var Languages = map[string]string{
	"aar": "afar",
	"abk": "abkhazian",
	"ace": "achinese",
	"ach": "acoli",
	"ada": "adangme",
	"ady": "adygei",
	"afa": "afro-asiatic languages",
	"afh": "afrihili",
	"afr": "afrikaans",
	"ain": "ainu",
	"aka": "akan",
	"akk": "akkadian",
	"ale": "aleut",
	"alg": "algonquian languages",
	"alt": "southern altai",
	"amh": "amharic",
	"ang": "english: old ca.450-1100",
	"anp": "angika",
	"apa": "apache languages",
	"ara": "arabic",
	"arc": "official aramaic 700-300 bce",
	"arg": "aragonese",
	"arn": "mapudungun: mapuche",
	"arp": "arapaho",
	"art": "artificial languages",
	"arw": "arawak",
	"asm": "assamese",
	"ast": "asturian",
	"ath": "athapascan languages",
	"aus": "australian languages",
	"ava": "avaric",
	"ave": "avestan",
	"awa": "awadhi",
	"aym": "aymara",
	"aze": "azerbaijani",
	"bad": "banda languages",
	"bai": "bamileke languages",
	"bak": "bashkir",
	"bal": "baluchi",
	"bam": "bambara",
	"ban": "balinese",
	"bas": "basa",
	"bat": "baltic languages",
	"bej": "beja: bedawiyet",
	"bel": "belarusian",
	"bem": "bemba",
	"ben": "bengali",
	"ber": "berber languages",
	"bho": "bhojpuri",
	"bih": "bihari languages",
	"bik": "bikol",
	"bin": "bini: edo",
	"bis": "bislama",
	"bla": "siksika",
	"bnt": "bantu languages",
	"bos": "bosnian",
	"bra": "braj",
	"bre": "breton",
	"btk": "batak languages",
	"bua": "buriat",
	"bug": "buginese",
	"bul": "bulgarian",
	"byn": "blin: bilin",
	"cad": "caddo",
	"cai": "central american indian languages",
	"car": "galibi carib",
	"cat": "catalan: valencian",
	"cau": "caucasian languages",
	"ceb": "cebuano",
	"cel": "celtic languages",
	"cha": "chamorro",
	"chb": "chibcha",
	"che": "chechen",
	"chg": "chagatai",
	"chk": "chuukese",
	"chm": "mari",
	"chn": "chinook jargon",
	"cho": "choctaw",
	"chp": "chipewyan: dene suline",
	"chr": "cherokee",
	"chu": "church slavic: old slavonic",
	"chv": "chuvash",
	"chy": "cheyenne",
	"cmc": "chamic languages",
	"cnr": "montenegrin",
	"cop": "coptic",
	"cor": "cornish",
	"cos": "corsican",
	"cpe": "creoles and pidgins: english based",
	"cpf": "creoles and pidgins: french-based",
	"cpp": "creoles and pidgins: portuguese-based",
	"cre": "cree",
	"crh": "crimean tatar",
	"crp": "creoles and pidgins",
	"csb": "kashubian",
	"cus": "cushitic languages",
	"dak": "dakota",
	"dan": "danish",
	"dar": "dargwa",
	"day": "land dayak languages",
	"del": "delaware",
	"den": "slave athapascan",
	"dgr": "dogrib",
	"din": "dinka",
	"div": "dhivehi: maldivian",
	"doi": "dogri",
	"dra": "dravidian languages",
	"dsb": "lower sorbian",
	"dua": "duala",
	"dum": "dutch: middle ca.1050-1350",
	"dyu": "dyula",
	"dzo": "dzongkha",
	"efi": "efik",
	"egy": "egyptian ancient",
	"eka": "ekajuk",
	"elx": "elamite",
	"eng": "english",
	"enm": "english: middle 1100-1500",
	"epo": "esperanto",
	"est": "estonian",
	"ewe": "ewe",
	"ewo": "ewondo",
	"fan": "fang",
	"fao": "faroese",
	"fat": "fanti",
	"fij": "fijian",
	"fil": "filipino: pilipino",
	"fin": "finnish",
	"fiu": "finno-ugrian languages",
	"fon": "fon",
	"frm": "french: middle ca.1400-1600",
	"fro": "french: old 842-ca.1400",
	"frr": "northern frisian",
	"frs": "eastern frisian",
	"fry": "western frisian",
	"ful": "fulah",
	"fur": "friulian",
	"gaa": "ga",
	"gay": "gayo",
	"gba": "gbaya",
	"gem": "germanic languages",
	"gez": "geez",
	"gil": "gilbertese",
	"gla": "gaelic: scottish gaelic",
	"gle": "irish",
	"glg": "galician",
	"glv": "manx",
	"gmh": "german: middle high ca.1050-1500",
	"goh": "german: old high ca.750-1050",
	"gon": "gondi",
	"gor": "gorontalo",
	"got": "gothic",
	"grb": "grebo",
	"grc": "greek: ancient to 1453",
	"grn": "guarani",
	"gsw": "swiss german: alemannic: alsatian",
	"guj": "gujarati",
	"gwi": "gwich in",
	"hai": "haida",
	"hat": "haitian: haitian creole",
	"hau": "hausa",
	"haw": "hawaiian",
	"heb": "hebrew",
	"her": "herero",
	"hil": "hiligaynon",
	"him": "himachali languages",
	"hin": "hindi",
	"hit": "hittite",
	"hmn": "hmong: mong",
	"hmo": "hiri motu",
	"hrv": "croatian",
	"hsb": "upper sorbian",
	"hun": "hungarian",
	"hup": "hupa",
	"iba": "iban",
	"ibo": "igbo",
	"ido": "ido",
	"iii": "sichuan yi: nuosu",
	"ijo": "ijo languages",
	"iku": "inuktitut",
	"ile": "interlingue: occidental",
	"ilo": "iloko",
	"ina": "interlingua international auxiliary language association",
	"inc": "indic languages",
	"ind": "indonesian",
	"ine": "indo-european languages",
	"inh": "ingush",
	"ipk": "inupiaq",
	"ira": "iranian languages",
	"iro": "iroquoian languages",
	"ita": "italian",
	"jav": "javanese",
	"jbo": "lojban",
	"jpn": "japanese",
	"jpr": "judeo-persian",
	"jrb": "judeo-arabic",
	"kaa": "kara-kalpak",
	"kab": "kabyle",
	"kac": "kachin: jingpho",
	"kal": "kalaallisut: greenlandic",
	"kam": "kamba",
	"kan": "kannada",
	"kar": "karen languages",
	"kas": "kashmiri",
	"kau": "kanuri",
	"kaw": "kawi",
	"kaz": "kazakh",
	"kbd": "kabardian",
	"kha": "khasi",
	"khi": "khoisan languages",
	"khm": "central khmer",
	"kho": "khotanese: sakan",
	"kik": "kikuyu: gikuyu",
	"kin": "kinyarwanda",
	"kir": "kirghiz",
	"kmb": "kimbundu",
	"kok": "konkani",
	"kom": "komi",
	"kon": "kongo",
	"kor": "korean",
	"kos": "kosraean",
	"kpe": "kpelle",
	"krc": "karachay-balkar",
	"krl": "karelian",
	"kro": "kru languages",
	"kru": "kurukh",
	"kua": "kuanyama: kwanyama",
	"kum": "kumyk",
	"kur": "kurdish",
	"kut": "kutenai",
	"lad": "ladino",
	"lah": "lahnda",
	"lam": "lamba",
	"lao": "lao",
	"lat": "latin",
	"lav": "latvian",
	"lez": "lezghian",
	"lim": "limburgan",
	"lin": "lingala",
	"lit": "lithuanian",
	"lol": "mongo",
	"loz": "lozi",
	"ltz": "luxembourgish",
	"lua": "luba-lulua",
	"lub": "luba-katanga",
	"lug": "ganda",
	"lui": "luiseno",
	"lun": "lunda",
	"luo": "luo kenya and tanzania",
	"lus": "lushai",
	"mad": "madurese",
	"mag": "magahi",
	"mah": "marshallese",
	"mai": "maithili",
	"mak": "makasar",
	"mal": "malayalam",
	"man": "mandingo",
	"map": "austronesian languages",
	"mar": "marathi",
	"mas": "masai",
	"mdf": "moksha",
	"mdr": "mandar",
	"men": "mende",
	"mga": "irish: middle 900-1200",
	"mic": "mi kmaq: micmac",
	"min": "minangkabau",
	"mis": "uncoded languages",
	"mkh": "mon-khmer languages",
	"mlg": "malagasy",
	"mlt": "maltese",
	"mnc": "manchu",
	"mni": "manipuri",
	"mno": "manobo languages",
	"moh": "mohawk",
	"mon": "mongolian",
	"mos": "mossi",
	"mul": "multiple languages",
	"mun": "munda languages",
	"mus": "creek",
	"mwl": "mirandese",
	"mwr": "marwari",
	"myn": "mayan languages",
	"myv": "erzya",
	"nah": "nahuatl languages",
	"nai": "north american indian languages",
	"nap": "neapolitan",
	"nau": "nauru",
	"nav": "navajo",
	"nbl": "ndebele: south: south ndebele",
	"nde": "ndebele: north: north ndebele",
	"ndo": "ndonga",
	"nds": "low german: low saxon",
	"nep": "nepali",
	"new": "nepal bhasa: newari",
	"nia": "nias",
	"nic": "niger-kordofanian languages",
	"niu": "niuean",
	"nno": "norwegian nynorsk",
	"nob": "bokmål: norwegian",
	"nog": "nogai",
	"non": "norse: old",
	"nor": "norwegian",
	"nqo": "n ko",
	"nso": "pedi: sepedi: northern sotho",
	"nub": "nubian languages",
	"nwc": "classical newari: old newari",
	"nya": "chichewa: nyanja",
	"nym": "nyamwezi",
	"nyn": "nyankole",
	"nyo": "nyoro",
	"nzi": "nzima",
	"oci": "occitan post 1500",
	"oji": "ojibwa",
	"ori": "oriya",
	"orm": "oromo",
	"osa": "osage",
	"oss": "ossetian",
	"ota": "turkish: ottoman 1500-1928",
	"oto": "otomian languages",
	"paa": "papuan languages",
	"pag": "pangasinan",
	"pal": "pahlavi",
	"pam": "pampanga: kapampangan",
	"pan": "panjabi",
	"pap": "papiamento",
	"pau": "palauan",
	"peo": "persian: old ca.600-400 B.C.",
	"phi": "philippine languages",
	"phn": "phoenician",
	"pli": "pali",
	"pol": "polish",
	"pon": "pohnpeian",
	"por": "portuguese",
	"pra": "prakrit languages",
	"pro": "provençal: old to 1500:occitan: old to 1500",
	"pus": "pushto",
	"que": "quechua",
	"raj": "rajasthani",
	"rap": "rapanui",
	"rar": "rarotongan: cook islands maori",
	"roa": "romance languages",
	"roh": "romansh",
	"rom": "romany",
	"run": "rundi",
	"rup": "arumanian: macedo-romanian",
	"rus": "russian",
	"sad": "sandawe",
	"sag": "sango",
	"sah": "yakut",
	"sai": "south american indian languages",
	"sal": "salishan languages",
	"sam": "samaritan aramaic",
	"san": "sanskrit",
	"sas": "sasak",
	"sat": "santali",
	"scn": "sicilian",
	"sco": "scots",
	"sel": "selkup",
	"sem": "semitic languages",
	"sga": "irish: old to 900",
	"sgn": "sign languages",
	"shn": "shan",
	"sid": "sidamo",
	"sin": "sinhala: sinhalese",
	"sio": "siouan languages",
	"sit": "sino-tibetan languages",
	"sla": "slavic languages",
	"slv": "slovenian",
	"sma": "southern sami",
	"sme": "northern sami",
	"smi": "sami languages",
	"smj": "lule sami",
	"smn": "inari sami",
	"smo": "samoan",
	"sms": "skolt sami",
	"sna": "shona",
	"snd": "sindhi",
	"snk": "soninke",
	"sog": "sogdian",
	"som": "somali",
	"son": "songhai languages",
	"sot": "sotho: southern",
	"spa": "spanish: castilian",
	"srd": "sardinian",
	"srn": "sranan tongo",
	"srp": "serbian",
	"srr": "serer",
	"ssa": "nilo-saharan languages",
	"ssw": "swati",
	"suk": "sukuma",
	"sun": "sundanese",
	"sus": "susu",
	"sux": "sumerian",
	"swa": "swahili",
	"swe": "swedish",
	"syc": "classical syriac",
	"syr": "syriac",
	"tah": "tahitian",
	"tai": "tai languages",
	"tam": "tamil",
	"tat": "tatar",
	"tel": "telugu",
	"tem": "timne",
	"ter": "tereno",
	"tet": "tetum",
	"tgk": "tajik",
	"tgl": "tagalog",
	"tha": "thai",
	"tig": "tigre",
	"tir": "tigrinya",
	"tiv": "tiv",
	"tkl": "tokelau",
	"tlh": "klingon: tlhingan-hol",
	"tli": "tlingit",
	"tmh": "tamashek",
	"tog": "tonga nyasa",
	"ton": "tonga tonga islands",
	"tpi": "tok pisin",
	"tsi": "tsimshian",
	"tsn": "tswana",
	"tso": "tsonga",
	"tuk": "turkmen",
	"tum": "tumbuka",
	"tup": "tupi languages",
	"tur": "turkish",
	"tut": "altaic languages",
	"tvl": "tuvalu",
	"twi": "twi",
	"tyv": "tuvinian",
	"udm": "udmurt",
	"uga": "ugaritic",
	"uig": "uighur",
	"ukr": "ukrainian",
	"umb": "umbundu",
	"und": "undetermined",
	"urd": "urdu",
	"uzb": "uzbek",
	"vai": "vai",
	"ven": "venda",
	"vie": "vietnamese",
	"vol": "volapük",
	"vot": "votic",
	"wak": "wakashan languages",
	"wal": "wolaitta",
	"war": "waray",
	"was": "washo",
	"wen": "sorbian languages",
	"wln": "walloon",
	"wol": "wolof",
	"xal": "kalmyk: oirat",
	"xho": "xhosa",
	"yao": "yao",
	"yap": "yapese",
	"yid": "yiddish",
	"yor": "yoruba",
	"ypk": "yupik languages",
	"zap": "zapotec",
	"zbl": "blissymbols",
	"zen": "zenaga",
	"zgh": "standard moroccan tamazight",
	"zha": "zhuang: chuang",
	"znd": "zande languages",
	"zul": "zulu",
	"zun": "zuni",
	"zza": "zaza: dimili: dimli: kirdki",
	"alb": "albanian",
	"sqi": "albanian",
	"arm": "armenian",
	"hye": "armenian",
	"baq": "basque",
	"eus": "basque",
	"tib": "tibetan",
	"bod": "tibetan",
	"bur": "burmese",
	"mya": "burmese",
	"cze": "czech",
	"ces": "czech",
	"chi": "chinese",
	"zho": "chinese",
	"wel": "welsh",
	"cym": "welsh",
	"ger": "german",
	"deu": "german",
	"dut": "dutch, flemish",
	"nld": "dutch, flemish",
	"gre": "greek, modern (1453-)",
	"ell": "greek, modern (1453-)",
	"per": "persian",
	"fas": "persian",
	"fra": "french",
	"fre": "french",
	"geo": "georgian",
	"kat": "georgian",
	"ice": "icelandic",
	"isl": "icelandic",
	"mac": "macedonian",
	"mkd": "macedonian",
	"mao": "maori",
	"mri": "maori",
	"may": "malay",
	"msa": "malay",
	"ron": "romanian, moldavian",
	"rum": "romanian, moldavian",
	"slk": "slovak",
	"slo": "slovak",
	"aa":  "afar",
	"ab":  "abkhazian",
	"af":  "afrikaans",
	"ak":  "akan",
	"sq":  "albanian",
	"am":  "amharic",
	"ar":  "arabic",
	"an":  "aragonese",
	"hy":  "armenian",
	"as":  "assamese",
	"av":  "avaric",
	"ae":  "avestan",
	"ay":  "aymara",
	"az":  "azerbaijani",
	"ba":  "bashkir",
	"bm":  "bambara",
	"eu":  "basque",
	"be":  "belarusian",
	"bn":  "bengali",
	"bh":  "bihari languages",
	"bi":  "bislama",
	"bo":  "tibetan",
	"bs":  "bosnian",
	"br":  "breton",
	"bg":  "bulgarian",
	"my":  "burmese",
	"ca":  "catalan, valencian",
	"cs":  "czech",
	"ch":  "chamorro",
	"ce":  "chechen",
	"zh":  "chinese",
	"cu":  "church slavic, old slavonic",
	"cv":  "chuvash",
	"kw":  "cornish",
	"co":  "corsican",
	"cr":  "cree",
	"cy":  "welsh",
	"da":  "danish",
	"de":  "german",
	"dv":  "dhivehi, maldivian",
	"nl":  "dutch, flemish",
	"dz":  "dzongkha",
	"el":  "greek, modern (1453-)",
	"en":  "english",
	"eo":  "esperanto",
	"et":  "estonian",
	"ee":  "ewe",
	"fo":  "faroese",
	"fa":  "persian",
	"fj":  "fijian",
	"fi":  "finnish",
	"fr":  "french",
	"fy":  "western frisian",
	"ff":  "fulah",
	"ka":  "georgian",
	"gd":  "gaelic, scottish gaelic",
	"ga":  "irish",
	"gl":  "galician",
	"gv":  "manx",
	"gn":  "guarani",
	"gu":  "gujarati",
	"ht":  "haitian, haitian creole",
	"ha":  "hausa",
	"he":  "hebrew",
	"hz":  "herero",
	"hi":  "hindi",
	"ho":  "hiri motu",
	"hr":  "croatian",
	"hu":  "hungarian",
	"ig":  "igbo",
	"is":  "icelandic",
	"io":  "ido",
	"ii":  "sichuan yi, nuosu",
	"iu":  "inuktitut",
	"ie":  "interlingue, occidental",
	"ia":  "interlingua (international auxiliary language association)",
	"id":  "indonesian",
	"ik":  "inupiaq",
	"it":  "italian",
	"jv":  "javanese",
	"ja":  "japanese",
	"kl":  "kalaallisut, greenlandic",
	"kn":  "kannada",
	"ks":  "kashmiri",
	"kr":  "kanuri",
	"kk":  "kazakh",
	"km":  "central khmer",
	"ki":  "kikuyu, gikuyu",
	"rw":  "kinyarwanda",
	"ky":  "kirghiz",
	"kv":  "komi",
	"kg":  "kongo",
	"ko":  "korean",
	"kj":  "kuanyama, kwanyama",
	"ku":  "kurdish",
	"lo":  "lao",
	"la":  "latin",
	"lv":  "latvian",
	"li":  "limburgan",
	"ln":  "lingala",
	"lt":  "lithuanian",
	"lb":  "luxembourgish",
	"lu":  "luba-katanga",
	"lg":  "ganda",
	"mk":  "macedonian",
	"mh":  "marshallese",
	"ml":  "malayalam",
	"mi":  "maori",
	"mr":  "marathi",
	"ms":  "malay",
	"mg":  "malagasy",
	"mt":  "maltese",
	"mn":  "mongolian",
	"na":  "nauru",
	"nv":  "navajo",
	"nr":  "ndebele, south, south ndebele",
	"nd":  "ndebele, north, north ndebele",
	"ng":  "ndonga",
	"ne":  "nepali",
	"nn":  "norwegian nynorsk",
	"nb":  "bokmål, norwegian",
	"no":  "norwegian",
	"ny":  "chichewa, nyanja",
	"oc":  "occitan (post 1500)",
	"oj":  "ojibwa",
	"or":  "oriya",
	"om":  "oromo",
	"os":  "ossetian",
	"pa":  "panjabi",
	"pi":  "pali",
	"pl":  "polish",
	"pt":  "portuguese",
	"ps":  "pushto",
	"qu":  "quechua",
	"rm":  "romansh",
	"ro":  "romanian, moldavian",
	"rn":  "rundi",
	"ru":  "russian",
	"sg":  "sango",
	"sa":  "sanskrit",
	"si":  "sinhala, sinhalese",
	"sk":  "slovak",
	"sl":  "slovenian",
	"se":  "northern sami",
	"sm":  "samoan",
	"sn":  "shona",
	"sd":  "sindhi",
	"so":  "somali",
	"st":  "sotho, southern",
	"es":  "spanish, castilian",
	"sc":  "sardinian",
	"sr":  "serbian",
	"ss":  "swati",
	"su":  "sundanese",
	"sw":  "swahili",
	"sv":  "swedish",
	"ty":  "tahitian",
	"ta":  "tamil",
	"tt":  "tatar",
	"te":  "telugu",
	"tg":  "tajik",
	"tl":  "tagalog",
	"th":  "thai",
	"ti":  "tigrinya",
	"to":  "tonga (tonga islands)",
	"tn":  "tswana",
	"ts":  "tsonga",
	"tk":  "turkmen",
	"tr":  "turkish",
	"tw":  "twi",
	"ug":  "uighur",
	"uk":  "ukrainian",
	"ur":  "urdu",
	"uz":  "uzbek",
	"ve":  "venda",
	"vi":  "vietnamese",
	"vo":  "volapük",
	"wa":  "walloon",
	"wo":  "wolof",
	"xh":  "xhosa",
	"yi":  "yiddish",
	"yo":  "yoruba",
	"za":  "zhuang, chuang",
	"zu":  "zulu",
}

// LanguageFromString возвращает код языка.
func LanguageFromString(lang string) string {
	normLang := strings.ToLower(lang)
	for code, name := range Languages {
		if code == normLang || name == lang {
			return name
		}
	}
	return "und"
}
