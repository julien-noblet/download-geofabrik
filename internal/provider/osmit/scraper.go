package osmit

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

var ErrFetchCatalog = errors.New("failed to fetch catalog")

const (
	ProviderName               = "osmit-estratti"
	DefaultConfigFile          = "osmit-estratti.yml"
	BaseURL                    = "https://osmit-estratti.wmcloud.org/output"
	StartURL                   = "https://osmit-estratti.wmcloud.org/"
	defaultTimeout             = 30 * time.Second
	defaultKeepAlive           = 30 * time.Second
	defaultIdleTimeout         = 90 * time.Second
	defaultMaxIdleConns        = 20
	defaultMaxIdleConnsPerHost = 10
	minCodeAndNameParts        = 2
)

// Provider implements provider.Provider for osmit-estratti.wmcloud.org (OSM Italy).
// Field alignment optimized.
type Provider struct {
	Client   *http.Client
	BaseURL  string
	StartURL string
}

// NewProvider creates a new OSM Italy provider.
func NewProvider() *Provider {
	return &Provider{
		BaseURL:  BaseURL,
		StartURL: StartURL,
		Client: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   defaultTimeout,
					KeepAlive: defaultKeepAlive,
				}).DialContext,
				MaxIdleConns:        defaultMaxIdleConns,
				MaxIdleConnsPerHost: defaultMaxIdleConnsPerHost,
				IdleConnTimeout:     defaultIdleTimeout,
				ForceAttemptHTTP2:   true,
			},
		},
	}
}

// Name returns the unique service name.
func (p *Provider) Name() string {
	return ProviderName
}

// Description returns a human-readable description.
func (p *Provider) Description() string {
	return "OpenStreetMap Italy extracts downloads service (osmit-estratti)"
}

// DefaultConfigFile returns the default configuration filename.
func (p *Provider) DefaultConfigFile() string {
	return DefaultConfigFile
}

// DefaultFormats returns format definitions supported by osmit-estratti.
func DefaultFormats() catalog.FormatDefinitions {
	return catalog.FormatDefinitions{
		catalog.FormatOsmPbf:    {ID: catalog.FormatOsmPbf, Loc: ".osm.pbf", BasePath: "pbf"},
		catalog.FormatGPKG:      {ID: catalog.FormatGPKG, Loc: ".gpkg", BasePath: "gpkg"},
		catalog.FormatOBF:       {ID: catalog.FormatOBF, Loc: ".obf", BasePath: "obf"},
		catalog.FormatGarminOSM: {ID: catalog.FormatGarminOSM, Loc: ".tar.gz", BasePath: "garmin"},
	}
}

const (
	slugPiemonte          = "piemonte"
	slugValleDAosta       = "valle-daosta"
	slugLombardia         = "lombardia"
	slugTrentinoAltoAdige = "trentino-alto-adige"
	slugVeneto            = "veneto"
	slugFriuliVG          = "friuli-venezia-giulia"
	slugLiguria           = "liguria"
	slugEmiliaRomagna     = "emilia-romagna"
	slugToscana           = "toscana"
	slugUmbria            = "umbria"
	slugMarche            = "marche"
	slugLazio             = "lazio"
	slugAbruzzo           = "abruzzo"
	slugMolise            = "molise"
	slugCampania          = "campania"
	slugPuglia            = "puglia"
	slugBasilicata        = "basilicata"
	slugCalabria          = "calabria"
	slugSicilia           = "sicilia"
	slugSardegna          = "sardegna"
)

type regionInfo struct {
	Code string
	Name string
	Slug string
}

type provinceInfo struct {
	Code       string
	Name       string
	Slug       string
	RegionSlug string
}

var italianRegions = []regionInfo{
	{Code: "01", Name: "Piemonte", Slug: slugPiemonte},
	{Code: "02", Name: "Valle d'Aosta", Slug: slugValleDAosta},
	{Code: "03", Name: "Lombardia", Slug: slugLombardia},
	{Code: "04", Name: "Trentino-Alto Adige", Slug: slugTrentinoAltoAdige},
	{Code: "05", Name: "Veneto", Slug: slugVeneto},
	{Code: "06", Name: "Friuli-Venezia Giulia", Slug: slugFriuliVG},
	{Code: "07", Name: "Liguria", Slug: slugLiguria},
	{Code: "08", Name: "Emilia-Romagna", Slug: slugEmiliaRomagna},
	{Code: "09", Name: "Toscana", Slug: slugToscana},
	{Code: "10", Name: "Umbria", Slug: slugUmbria},
	{Code: "11", Name: "Marche", Slug: slugMarche},
	{Code: "12", Name: "Lazio", Slug: slugLazio},
	{Code: "13", Name: "Abruzzo", Slug: slugAbruzzo},
	{Code: "14", Name: "Molise", Slug: slugMolise},
	{Code: "15", Name: "Campania", Slug: slugCampania},
	{Code: "16", Name: "Puglia", Slug: slugPuglia},
	{Code: "17", Name: "Basilicata", Slug: slugBasilicata},
	{Code: "18", Name: "Calabria", Slug: slugCalabria},
	{Code: "19", Name: "Sicilia", Slug: slugSicilia},
	{Code: "20", Name: "Sardegna", Slug: slugSardegna},
}

var italianProvinces = []provinceInfo{
	{Code: "001", Name: "Torino", Slug: "torino", RegionSlug: slugPiemonte},
	{Code: "002", Name: "Vercelli", Slug: "vercelli", RegionSlug: slugPiemonte},
	{Code: "003", Name: "Novara", Slug: "novara", RegionSlug: slugPiemonte},
	{Code: "004", Name: "Cuneo", Slug: "cuneo", RegionSlug: slugPiemonte},
	{Code: "005", Name: "Asti", Slug: "asti", RegionSlug: slugPiemonte},
	{Code: "006", Name: "Alessandria", Slug: "alessandria", RegionSlug: slugPiemonte},
	{Code: "007", Name: "Aosta", Slug: "aosta", RegionSlug: slugValleDAosta},
	{Code: "008", Name: "Imperia", Slug: "imperia", RegionSlug: slugLiguria},
	{Code: "009", Name: "Savona", Slug: "savona", RegionSlug: slugLiguria},
	{Code: "010", Name: "Genova", Slug: "genova", RegionSlug: slugLiguria},
	{Code: "011", Name: "La Spezia", Slug: "la-spezia", RegionSlug: slugLiguria},
	{Code: "012", Name: "Varese", Slug: "varese", RegionSlug: slugLombardia},
	{Code: "013", Name: "Como", Slug: "como", RegionSlug: slugLombardia},
	{Code: "014", Name: "Sondrio", Slug: "sondrio", RegionSlug: slugLombardia},
	{Code: "015", Name: "Milano", Slug: "milano", RegionSlug: slugLombardia},
	{Code: "016", Name: "Bergamo", Slug: "bergamo", RegionSlug: slugLombardia},
	{Code: "017", Name: "Brescia", Slug: "brescia", RegionSlug: slugLombardia},
	{Code: "018", Name: "Pavia", Slug: "pavia", RegionSlug: slugLombardia},
	{Code: "019", Name: "Cremona", Slug: "cremona", RegionSlug: slugLombardia},
	{Code: "020", Name: "Mantova", Slug: "mantova", RegionSlug: slugLombardia},
	{Code: "021", Name: "Bolzano", Slug: "bolzano", RegionSlug: slugTrentinoAltoAdige},
	{Code: "022", Name: "Trento", Slug: "trento", RegionSlug: slugTrentinoAltoAdige},
	{Code: "023", Name: "Verona", Slug: "verona", RegionSlug: slugVeneto},
	{Code: "024", Name: "Vicenza", Slug: "vicenza", RegionSlug: slugVeneto},
	{Code: "025", Name: "Belluno", Slug: "belluno", RegionSlug: slugVeneto},
	{Code: "026", Name: "Treviso", Slug: "treviso", RegionSlug: slugVeneto},
	{Code: "027", Name: "Venezia", Slug: "venezia", RegionSlug: slugVeneto},
	{Code: "028", Name: "Padova", Slug: "padova", RegionSlug: slugVeneto},
	{Code: "029", Name: "Rovigo", Slug: "rovigo", RegionSlug: slugVeneto},
	{Code: "030", Name: "Udine", Slug: "udine", RegionSlug: slugFriuliVG},
	{Code: "031", Name: "Gorizia", Slug: "gorizia", RegionSlug: slugFriuliVG},
	{Code: "032", Name: "Trieste", Slug: "trieste", RegionSlug: slugFriuliVG},
	{Code: "033", Name: "Piacenza", Slug: "piacenza", RegionSlug: slugEmiliaRomagna},
	{Code: "034", Name: "Parma", Slug: "parma", RegionSlug: slugEmiliaRomagna},
	{Code: "035", Name: "Reggio Emilia", Slug: "reggio-emilia", RegionSlug: slugEmiliaRomagna},
	{Code: "036", Name: "Modena", Slug: "modena", RegionSlug: slugEmiliaRomagna},
	{Code: "037", Name: "Bologna", Slug: "bologna", RegionSlug: slugEmiliaRomagna},
	{Code: "038", Name: "Ferrara", Slug: "ferrara", RegionSlug: slugEmiliaRomagna},
	{Code: "039", Name: "Ravenna", Slug: "ravenna", RegionSlug: slugEmiliaRomagna},
	{Code: "040", Name: "Forlì-Cesena", Slug: "forli-cesena", RegionSlug: slugEmiliaRomagna},
	{Code: "041", Name: "Pesaro e Urbino", Slug: "pesaro-urbino", RegionSlug: slugMarche},
	{Code: "042", Name: "Ancona", Slug: "ancona", RegionSlug: slugMarche},
	{Code: "043", Name: "Macerata", Slug: "macerata", RegionSlug: slugMarche},
	{Code: "044", Name: "Ascoli Piceno", Slug: "ascoli-piceno", RegionSlug: slugMarche},
	{Code: "045", Name: "Massa-Carrara", Slug: "massa-carrara", RegionSlug: slugToscana},
	{Code: "046", Name: "Lucca", Slug: "lucca", RegionSlug: slugToscana},
	{Code: "047", Name: "Pistoia", Slug: "pistoia", RegionSlug: slugToscana},
	{Code: "048", Name: "Firenze", Slug: "firenze", RegionSlug: slugToscana},
	{Code: "049", Name: "Livorno", Slug: "livorno", RegionSlug: slugToscana},
	{Code: "050", Name: "Pisa", Slug: "pisa", RegionSlug: slugToscana},
	{Code: "051", Name: "Arezzo", Slug: "arezzo", RegionSlug: slugToscana},
	{Code: "052", Name: "Siena", Slug: "siena", RegionSlug: slugToscana},
	{Code: "053", Name: "Grosseto", Slug: "grosseto", RegionSlug: slugToscana},
	{Code: "054", Name: "Perugia", Slug: "perugia", RegionSlug: slugUmbria},
	{Code: "055", Name: "Terni", Slug: "terni", RegionSlug: slugUmbria},
	{Code: "056", Name: "Viterbo", Slug: "viterbo", RegionSlug: slugLazio},
	{Code: "057", Name: "Rieti", Slug: "rieti", RegionSlug: slugLazio},
	{Code: "058", Name: "Roma", Slug: "roma", RegionSlug: slugLazio},
	{Code: "059", Name: "Latina", Slug: "latina", RegionSlug: slugLazio},
	{Code: "060", Name: "Frosinone", Slug: "frosinone", RegionSlug: slugLazio},
	{Code: "061", Name: "Caserta", Slug: "caserta", RegionSlug: slugCampania},
	{Code: "062", Name: "Benevento", Slug: "benevento", RegionSlug: slugCampania},
	{Code: "063", Name: "Napoli", Slug: "napoli", RegionSlug: slugCampania},
	{Code: "064", Name: "Avellino", Slug: "avellino", RegionSlug: slugCampania},
	{Code: "065", Name: "Salerno", Slug: "salerno", RegionSlug: slugCampania},
	{Code: "066", Name: "L'Aquila", Slug: "laquila", RegionSlug: slugAbruzzo},
	{Code: "067", Name: "Teramo", Slug: "teramo", RegionSlug: slugAbruzzo},
	{Code: "068", Name: "Pescara", Slug: "pescara", RegionSlug: slugAbruzzo},
	{Code: "069", Name: "Chieti", Slug: "chieti", RegionSlug: slugAbruzzo},
	{Code: "070", Name: "Campobasso", Slug: "campobasso", RegionSlug: slugMolise},
	{Code: "071", Name: "Foggia", Slug: "foggia", RegionSlug: slugPuglia},
	{Code: "072", Name: "Bari", Slug: "bari", RegionSlug: slugPuglia},
	{Code: "073", Name: "Taranto", Slug: "taranto", RegionSlug: slugPuglia},
	{Code: "074", Name: "Brindisi", Slug: "brindisi", RegionSlug: slugPuglia},
	{Code: "075", Name: "Lecce", Slug: "lecce", RegionSlug: slugPuglia},
	{Code: "076", Name: "Potenza", Slug: "potenza", RegionSlug: slugBasilicata},
	{Code: "077", Name: "Matera", Slug: "matera", RegionSlug: slugBasilicata},
	{Code: "078", Name: "Cosenza", Slug: "cosenza", RegionSlug: slugCalabria},
	{Code: "079", Name: "Catanzaro", Slug: "catanzaro", RegionSlug: slugCalabria},
	{Code: "080", Name: "Reggio Calabria", Slug: "reggio-calabria", RegionSlug: slugCalabria},
	{Code: "081", Name: "Trapani", Slug: "trapani", RegionSlug: slugSicilia},
	{Code: "082", Name: "Palermo", Slug: "palermo", RegionSlug: slugSicilia},
	{Code: "083", Name: "Messina", Slug: "messina", RegionSlug: slugSicilia},
	{Code: "084", Name: "Agrigento", Slug: "agrigento", RegionSlug: slugSicilia},
	{Code: "085", Name: "Caltanissetta", Slug: "caltanissetta", RegionSlug: slugSicilia},
	{Code: "086", Name: "Enna", Slug: "enna", RegionSlug: slugSicilia},
	{Code: "087", Name: "Catania", Slug: "catania", RegionSlug: slugSicilia},
	{Code: "088", Name: "Ragusa", Slug: "ragusa", RegionSlug: slugSicilia},
	{Code: "089", Name: "Siracusa", Slug: "siracusa", RegionSlug: slugSicilia},
	{Code: "090", Name: "Sassari", Slug: "sassari", RegionSlug: slugSardegna},
	{Code: "091", Name: "Nuoro", Slug: "nuoro", RegionSlug: slugSardegna},
	{Code: "092", Name: "Cagliari", Slug: "cagliari", RegionSlug: slugSardegna},
	{Code: "093", Name: "Pordenone", Slug: "pordenone", RegionSlug: slugFriuliVG},
	{Code: "094", Name: "Isernia", Slug: "isernia", RegionSlug: slugMolise},
	{Code: "095", Name: "Oristano", Slug: "oristano", RegionSlug: slugSardegna},
	{Code: "096", Name: "Biella", Slug: "biella", RegionSlug: slugPiemonte},
	{Code: "097", Name: "Lecco", Slug: "lecco", RegionSlug: slugLombardia},
	{Code: "098", Name: "Lodi", Slug: "lodi", RegionSlug: slugLombardia},
	{Code: "099", Name: "Rimini", Slug: "rimini", RegionSlug: slugEmiliaRomagna},
	{Code: "100", Name: "Prato", Slug: "prato", RegionSlug: slugToscana},
	{Code: "101", Name: "Crotone", Slug: "crotone", RegionSlug: slugCalabria},
	{Code: "102", Name: "Vibo Valentia", Slug: "vibo-valentia", RegionSlug: slugCalabria},
	{Code: "103", Name: "Verbano-Cusio-Ossola", Slug: "verbano-cusio-ossola", RegionSlug: slugPiemonte},
	{Code: "108", Name: "Monza e Brianza", Slug: "monza-brianza", RegionSlug: slugLombardia},
	{Code: "109", Name: "Fermo", Slug: "fermo", RegionSlug: slugMarche},
	{Code: "110", Name: "Barletta-Andria-Trani", Slug: "barletta-andria-trani", RegionSlug: slugPuglia},
	{Code: "111", Name: "Sud Sardegna", Slug: "sud-sardegna", RegionSlug: slugSardegna},
}

// FetchCatalog builds the OpenStreetMap Italy catalog after verifying server availability.
func (p *Provider) FetchCatalog(ctx context.Context) (*catalog.Catalog, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.StartURL, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := p.Client
	if client == nil {
		client = http.DefaultClient
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFetchCatalog, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: unexpected HTTP status %d", ErrFetchCatalog, resp.StatusCode)
	}

	cat := catalog.New()
	cat.BaseURL = p.BaseURL
	cat.Formats = DefaultFormats()

	// Add Regions
	for _, reg := range italianRegions {
		regElem := catalog.Element{
			ID:      reg.Slug,
			Name:    reg.Name,
			File:    "regioni/" + reg.Code + "_" + sanitizeFileName(reg.Name),
			Formats: catalog.Formats{catalog.FormatOsmPbf, catalog.FormatGPKG, catalog.FormatOBF, catalog.FormatGarminOSM},
		}
		_ = cat.MergeElement(&regElem)
	}

	// Add Provinces
	for _, prov := range italianProvinces {
		provElem := catalog.Element{
			ID:      prov.Slug,
			Name:    prov.Name,
			File:    "province/" + prov.Code + "_" + sanitizeFileName(prov.Name),
			Formats: catalog.Formats{catalog.FormatOsmPbf, catalog.FormatGPKG},
		}
		_ = cat.MergeElement(&provElem)
	}

	return cat, nil
}

func sanitizeFileName(name string) string {
	replacer := strings.NewReplacer(
		" ", "_",
		"-", "_",
		"'", "",
	)

	return replacer.Replace(name)
}
