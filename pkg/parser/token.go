package parser

type Token int

const (
	ILLEGAL Token = iota
	EOF
	WS

	// Literals
	IDENTIFIER

	// MISC chars
	COMMA         // ,
	OPEN_BRACE    // {
	CLOSING_BRACE // }
	EQUAL         // =
	QUOTE         // "

	RELATION // ->

	// Model Keywords
	WORKSPACE
	MODEL
	ENTERPRISE
	GROUP
	PERSON
	SOFTWARE_SYSTEM
	CONTAINER
	COMPONENT
	DEPLOYMENT_ENV
	DEPLOYMENT_GROUP
	DEPLOYMENT_NODE
	INFRASTRUCTURE_NODE
	SOFTWARE_SYSTEM_INSTANCE
	CONTAINER_INSTANCE
	ELEMENT

	// Views Keywords
	VIEWS
	SYSTEM_LANDSCAPE
	SYSTEM_CONTEXT
	FILTERED
	DYNAMIC
	DEPLOYMENT
	CUSTOM
	STYLES
	THEME
	THEMES
	BRANDING
	TERMINOLOGY

	CONFIGURATION
	USERS

	// Reference keywords and pragmas
	BANG_DOCS
	BANG_ADRS
	BANG_IDENTIFIERS
	BANG_IMPLIED_RELATIONSHIPS
	BANG_REF
	BANG_INCLUDE
	EXTENDS
)

var keywords map[string]Token = map[string]Token{
	"->": RELATION, // is contained here, as it has more then one character

	// Model Keywords
	"workspace":              WORKSPACE,
	"model":                  MODEL,
	"enterprise":             ENTERPRISE,
	"group":                  GROUP,
	"person":                 PERSON,
	"softwaresystem":         SOFTWARE_SYSTEM,
	"container":              CONTAINER,
	"component":              COMPONENT,
	"deploymentenvironment":  DEPLOYMENT_ENV,
	"deploymentgroup":        DEPLOYMENT_GROUP,
	"deploymentnode":         DEPLOYMENT_NODE,
	"infrastructurenode":     INFRASTRUCTURE_NODE,
	"softwaresysteminstance": SOFTWARE_SYSTEM_INSTANCE,
	"containerinstance":      CONTAINER_INSTANCE,
	"element":                ELEMENT,

	// Views Keywords
	"views":           VIEWS,
	"systemlandscape": SYSTEM_LANDSCAPE,
	"systemcontext":   SYSTEM_CONTEXT,
	"filtered":        FILTERED,
	"dynamic":         DYNAMIC,
	"deployment":      DEPLOYMENT,
	"custom":          CUSTOM,
	"styles":          STYLES,
	"theme":           THEME,
	"themes":          THEMES,
	"branding":        BRANDING,
	"terminology":     TERMINOLOGY,

	"configation": CONFIGURATION,
	"users":       USERS,

	// Reference keywords and pragmas
	"!docs":                 BANG_DOCS,
	"!adrs":                 BANG_ADRS,
	"!identifiers":          BANG_IDENTIFIERS,
	"!impliedrelationships": BANG_IMPLIED_RELATIONSHIPS,
	"!ref":                  BANG_REF,
	"!include":              BANG_INCLUDE,
	"extends":               EXTENDS,
}

// this basically reverse lookup to the keywords map, but extended with the
// single character tokens and formated in a canonical capitalisation
var readables map[Token]string = map[Token]string{
	ILLEGAL: "<illegal or unknown token>",

	EOF: "<end of file>",
	WS:  "<whitespace>",

	// Literals
	IDENTIFIER: "<an identifier>",

	// MISC chars
	COMMA:         ",",
	OPEN_BRACE:    "{",
	CLOSING_BRACE: "}",
	EQUAL:         "=",
	QUOTE:         "\"",

	RELATION: "->",

	// Model Keywords
	WORKSPACE:                "workspace",
	MODEL:                    "model",
	ENTERPRISE:               "enterprise",
	GROUP:                    "group",
	PERSON:                   "person",
	SOFTWARE_SYSTEM:          "softwareSystem",
	CONTAINER:                "container",
	COMPONENT:                "component",
	DEPLOYMENT_ENV:           "deploymentEnvironment",
	DEPLOYMENT_GROUP:         "deploymentGroup",
	DEPLOYMENT_NODE:          "deploymentNode",
	INFRASTRUCTURE_NODE:      "infrastructureNode",
	SOFTWARE_SYSTEM_INSTANCE: "softwaresystemInstance",
	CONTAINER_INSTANCE:       "containerInstance",
	ELEMENT:                  "element",

	// Views Keywords
	VIEWS:            "views",
	SYSTEM_LANDSCAPE: "systemLandscape",
	SYSTEM_CONTEXT:   "systemContext",
	FILTERED:         "filtered",
	DYNAMIC:          "dynamic",
	DEPLOYMENT:       "deployment",
	CUSTOM:           "custom",
	STYLES:           "styles",
	THEME:            "theme",
	THEMES:           "themes",
	BRANDING:         "branding",
	TERMINOLOGY:      "terminology",

	CONFIGURATION: "configation",
	USERS:         "users",

	// Reference keywords and pragmas
	BANG_DOCS:                  "!docs",
	BANG_ADRS:                  "!adrs",
	BANG_IDENTIFIERS:           "!identifiers",
	BANG_IMPLIED_RELATIONSHIPS: "!impliedRelationships",
	BANG_REF:                   "!ref",
	BANG_INCLUDE:               "!include",
	EXTENDS:                    "extends",
}

func (t Token) String() string {
	ret, ok := readables[t]
	if ok {
		return ret
	} else {
		return readables[ILLEGAL]
	}
}
