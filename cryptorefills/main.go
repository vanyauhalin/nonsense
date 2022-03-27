package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/chromedp/chromedp"
)

var (
	target = &Target{
		Brands: TargetBrands{
			"Spotify",
		},
		// Given as constant in main.js.
		Countries: TargetCountries{
			"ad": "Andorra",
			"ae": "United Arab Emirates",
			"af": "Afghanistan",
			"ag": "Antigua and Barbuda",
			"ai": "Anguilla",
			"al": "Albania",
			"am": "Armenia",
			"ao": "Angola",
			"ar": "Argentina",
			"as": "American Samoa",
			"at": "Austria",
			"au": "Australia",
			"aw": "Aruba",
			"ax": "Åland Islands",
			"az": "Azerbaijan",
			"ba": "Bosnia and Herzegovina",
			"bb": "Barbados",
			"bd": "Bangladesh",
			"be": "Belgium",
			"bf": "Burkina Faso",
			"bg": "Bulgaria",
			"bh": "Bahrain",
			"bi": "Burundi",
			"bj": "Benin",
			"bl": "Saint Barthélemy",
			"bm": "Bermuda",
			"bn": "Brunei",
			"bo": "Bolivia",
			"bq": "Caribbean Netherlands",
			"br": "Brazil",
			"bs": "Bahamas",
			"bt": "Bhutan",
			"bw": "Botswana",
			"by": "Belarus",
			"bz": "Belize",
			"ca": "Canada",
			"cc": "Cocos Islands",
			"cd": "Congo (DRC)",
			"cf": "Central African Republic",
			"cg": "Congo (Republic)",
			"ch": "Switzerland",
			"ci": "Côte d’Ivoire",
			"ck": "Cook Islands",
			"cl": "Chile",
			"cm": "Cameroon",
			"cn": "China",
			"co": "Colombia",
			"cr": "Costa Rica",
			"cv": "Cape Verde",
			"cw": "Curaçao",
			"cx": "Christmas Island",
			"cy": "Cyprus",
			"cz": "Czech Republic",
			"de": "Germany",
			"dj": "Djibouti",
			"dk": "Denmark",
			"dm": "Dominica",
			"do": "Dominican Republic",
			"dz": "Algeria",
			"ec": "Ecuador",
			"ee": "Estonia",
			"eg": "Egypt",
			"eh": "Western Sahara",
			"er": "Eritrea",
			"es": "Spain",
			"et": "Ethiopia",
			"fi": "Finland",
			"fj": "Fiji",
			"fk": "Falkland Islands",
			"fm": "Micronesia",
			"fo": "Faroe Islands",
			"fr": "France",
			"ga": "Gabon",
			"gb": "United Kingdom",
			"gd": "Grenada",
			"ge": "Georgia",
			"gf": "French Guiana",
			"gg": "Guernsey",
			"gh": "Ghana",
			"gi": "Gibraltar",
			"gl": "Greenland",
			"gm": "Gambia",
			"gn": "Guinea",
			"gp": "Guadeloupe",
			"gq": "Equatorial Guinea",
			"gr": "Greece",
			"gt": "Guatemala",
			"gu": "Guam",
			"gw": "Guinea-Bissau",
			"gy": "Guyana",
			"hk": "Hong Kong",
			"hn": "Honduras",
			"hr": "Croatia",
			"ht": "Haiti",
			"hu": "Hungary",
			"id": "Indonesia",
			"ie": "Ireland",
			"il": "Israel",
			"im": "Isle of Man",
			"in": "India",
			"io": "British Indian Ocean Territory",
			"iq": "Iraq",
			"is": "Iceland",
			"it": "Italy",
			"je": "Jersey",
			"jm": "Jamaica",
			"jo": "Jordan",
			"jp": "Japan",
			"ke": "Kenya",
			"kg": "Kyrgyzstan",
			"kh": "Cambodia",
			"ki": "Kiribati",
			"km": "Comoros",
			"kn": "Saint Kitts and Nevis",
			"kr": "South Korea",
			"kw": "Kuwait",
			"ky": "Cayman Islands",
			"kz": "Kazakhstan",
			"la": "Laos",
			"lb": "Lebanon",
			"lc": "Saint Lucia",
			"li": "Liechtenstein",
			"lk": "Sri Lanka",
			"lr": "Liberia",
			"ls": "Lesotho",
			"lt": "Lithuania",
			"lu": "Luxembourg",
			"lv": "Latvia",
			"ly": "Libya",
			"ma": "Morocco",
			"mc": "Monaco",
			"md": "Moldova",
			"me": "Montenegro",
			"mf": "Saint Martin",
			"mg": "Madagascar",
			"mh": "Marshall Islands",
			"mk": "Macedonia",
			"ml": "Mali",
			"mm": "Myanmar",
			"mn": "Mongolia",
			"mo": "Macau",
			"mp": "Northern Mariana Islands",
			"mq": "Martinique",
			"mr": "Mauritania",
			"ms": "Montserrat",
			"mt": "Malta",
			"mu": "Mauritius",
			"mv": "Maldives",
			"mw": "Malawi",
			"mx": "Mexico",
			"my": "Malaysia",
			"mz": "Mozambique",
			"na": "Namibia",
			"nc": "New Caledonia",
			"ne": "Niger",
			"nf": "Norfolk Island",
			"ng": "Nigeria",
			"ni": "Nicaragua",
			"nl": "Netherlands",
			"no": "Norway",
			"np": "Nepal",
			"nr": "Nauru",
			"nu": "Niue",
			"nz": "New Zealand",
			"om": "Oman",
			"pa": "Panama",
			"pe": "Peru",
			"pf": "French Polynesia",
			"pg": "Papua New Guinea",
			"ph": "Philippines",
			"pk": "Pakistan",
			"pl": "Poland",
			"pm": "Saint Pierre and Miquelon",
			"pr": "Puerto Rico",
			"ps": "Palestine",
			"pt": "Portugal",
			"pw": "Palau",
			"py": "Paraguay",
			"qa": "Qatar",
			"re": "Réunion",
			"ro": "Romania",
			"rs": "Serbia",
			"ru": "Russia",
			"rw": "Rwanda",
			"sa": "Saudi Arabia",
			"sb": "Solomon Islands",
			"sc": "Seychelles",
			"se": "Sweden",
			"sg": "Singapore",
			"sh": "Saint Helena",
			"si": "Slovenia",
			"sj": "Svalbard and Jan Mayen",
			"sk": "Slovakia",
			"sl": "Sierra Leone",
			"sm": "San Marino",
			"sn": "Senegal",
			"so": "Somalia",
			"sr": "Suriname",
			"st": "São Tomé and Príncipe",
			"sv": "El Salvador",
			"sx": "Sint Maarten",
			"sz": "Swaziland",
			"tc": "Turks and Caicos Islands",
			"td": "Chad",
			"tg": "Togo",
			"th": "Thailand",
			"tj": "Tajikistan",
			"tk": "Tokelau",
			"tl": "Timor-Leste",
			"tm": "Turkmenistan",
			"tn": "Tunisia",
			"to": "Tonga",
			"tr": "Turkey",
			"tt": "Trinidad and Tobago",
			"tv": "Tuvalu",
			"tw": "Taiwan",
			"tz": "Tanzania",
			"ua": "Ukraine",
			"ug": "Uganda",
			"us": "United States",
			"uy": "Uruguay",
			"uz": "Uzbekistan",
			"va": "Vatican City",
			"vc": "Saint Vincent and the Grenadines",
			"vg": "British Virgin Islands",
			"vi": "U.S. Virgin Islands",
			"vn": "Vietnam",
			"vu": "Vanuatu",
			"wf": "Wallis and Futuna",
			"ws": "Samoa",
			"ye": "Yemen",
			"yt": "Mayotte",
			"za": "South Africa",
			"zm": "Zambia",
			"zw": "Zimbabwe",
		},
		// TODO: not all currencies are used.
		Currencies: TargetCurrencies{
			"BTC",
			"USDT",
		},
	}
	store = make(Store)
)
const (
	www = "https://www.cryptorefills.com"
	api = "https://api.cryptorefills.com"
)

type Target struct {
	Brands TargetBrands
	Countries TargetCountries
	Currencies TargetCurrencies
}
type TargetBrands = []BrandBrand
type TargetCountries = map[CountryCode]string
type TargetCurrencies = []CurrencyName

type Store = map[CountryCode]StoreCountry
type StoreCountry = map[BrandID]*StoreBrand
type StoreBrand struct {
	Brand Brand
	Products StoreProducts
}
type StoreProducts = map[CurrencyName]BrandProducts

type CountryCode = string

type BrandsResponse struct {
	All_brands Brands
}
type Brands = []Brand
type Brand struct {
	Brand BrandBrand
	Brand_id BrandID
	Category string
}
type BrandBrand = string
type BrandID = string

type CurrenciesResponce []Currency
type Currency struct {
	Name CurrencyName
}
type CurrencyName = string

type BrandResponce struct {
	Products BrandProducts
}
type BrandProducts = []BrandProduct
type BrandProduct struct {
	Coin string
	Coin_amount string
	Denomination string
}

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	err := addBrands(ctx)
	if err != nil {
		log.Fatal("countries handling:", err)
	}
	if len(store) == 0 {
		log.Fatal("parsed countries is empty")
	}

	err = addProducts(ctx)
	if err != nil {
		log.Fatal("brands handling:", err)
	}

	createReport()
}

func addBrands(ctx context.Context) error {
	return chromedp.Run(
		ctx,
		chromedp.Navigate(www),
		chromedp.ActionFunc(func(ctx context.Context) error {
			number := 0
			length := len(target.Countries)

			for code := range target.Countries {
				number += 1
				log.Println(fmt.Sprintf("[%d/%d] %s", number, length, code))

				res, err := requestBrands(code)
				if err != nil {
					continue
				}

				parsed := parseBrands(res.All_brands)
				if len(parsed) > 0 {
					for _, brand := range parsed {
						store[code] = make(StoreCountry)
						store[code][brand.Brand_id] = &StoreBrand{brand, nil}
					}
				}
			}

			return nil
		}),
	)
}

func requestBrands(code CountryCode) (BrandsResponse, error) {
	payload := fmt.Sprintf("?country_code=%s", strings.ToUpper(code))
	url := fmt.Sprintf("%s/v2/brands%s", api, payload)
	req, err := http.Get(url)
	if err != nil {
		return BrandsResponse{}, err
	}

	var res BrandsResponse
	err = json.NewDecoder(req.Body).Decode(&res)
	if err != nil {
		return BrandsResponse{}, err
	}

	return res, nil
}

func parseBrands(brands Brands) Brands {
	var parsed Brands
	for _, brand := range brands {
		for _, target := range target.Brands {
			if brand.Brand == target {
				parsed = append(parsed, brand)
			}
		}
	}
	return parsed
}

func addProducts(ctx context.Context) error {
	for code, country := range store {
		for brandID, brand := range country {
			url := createBrandURL(code, brand.Brand)
			if err := chromedp.Run(
				ctx,
				chromedp.Navigate(url),
				chromedp.ActionFunc(func(ctx context.Context) error {
					url := fmt.Sprintf("%s/v3/currencies", api)
					res, err := requestCurrencies(url)
					if err != nil {
						return err
					}

					parsed := parseCurrencies(res)
					if len(parsed) == 0 {
						return nil
					}

					for _, currency := range parsed {
						payload := fmt.Sprintf(
							"?country_code=%s&coin=%s",
							code,
							currency.Name,
						)
						url = fmt.Sprintf(
							"%s/v2/products/brand/%s%s",
							api,
							brand.Brand.Brand_id,
							payload,
						)
						res, err := requestBrand(url)
						if err != nil {
							continue
						}
						if len(res.Products) > 0 {
							store[code][brandID].Products = make(StoreProducts)
							store[code][brandID].Products[currency.Name] = res.Products
						}
					}

					return nil
				}),
			); err != nil {
				return err
			}
		}
	}
	return nil
}

func createBrandURL(code CountryCode, brand Brand) string {
	return createURL(strings.Join([]string{
		www,
		target.Countries[code],
		brand.Category,
		brand.Brand,
	}, "/"))
}

func createURL(str string) string {
	return strings.Replace(strings.ToLower(str), " ", "_", -1)
}

func requestCurrencies(url string) (CurrenciesResponce, error) {
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var res CurrenciesResponce
	err = json.NewDecoder(req.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func parseCurrencies(currencies CurrenciesResponce) CurrenciesResponce {
	var parsed CurrenciesResponce
	for _, currency := range currencies {
		for _, target := range target.Currencies {
			if currency.Name == target {
				parsed = append(parsed, currency)
			}
		}
	}
	return parsed
}

func requestBrand(url string) (BrandResponce, error) {
	req, err := http.Get(url)
	if err != nil {
		return BrandResponce{}, err
	}

	var res BrandResponce
	err = json.NewDecoder(req.Body).Decode(&res)
	if err != nil {
		return BrandResponce{}, err
	}

	return res, nil
}

// TODO: make more flexible.
func createReport() {
	report := "Report:\ncode\tbrand\t\tcurrency\tdenomination\t\tamount\n"
	for code, country := range store {
		report += fmt.Sprintf("%s\t", code)
		for _, brand := range country {
			report += fmt.Sprintf("%s\t\t", brand.Brand.Brand)
			for currency, products := range brand.Products {
				report += fmt.Sprintf("%s\t\t", currency)
				for index, product := range products {
					if len(product.Denomination) == 0 {
						continue
					}
					if index == 0 {
						report += fmt.Sprintf(
							"%s%s%s\n",
							product.Denomination,
							strings.Repeat(" ", 24 - len(product.Denomination)),
							product.Coin_amount,
						)
					} else {
						report += fmt.Sprintf(
							"\t\t\t\t\t%s%s%s\n",
							product.Denomination,
							strings.Repeat(" ", 24 - len(product.Denomination)),
							product.Coin_amount,
						)
					}
				}
			}
		}
	}
	log.Println(report)
}
