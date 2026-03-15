package main

var allowlist = map[string][]string{
	"svg": {
		"width",
		"height",
		"viewBox",
		"xmlns",
		"fill",
	},

	"g": {
		"fill",
		"stroke",
		"stroke-width",
		"transform",
	},

	"path": {
		"d",
		"fill",
		"stroke",
		"stroke-width",
		"transform",
	},

	"circle": {
		"cx",
		"cy",
		"r",
		"fill",
		"stroke",
		"stroke-width",
	},

	"rect": {
		"x",
		"y",
		"width",
		"height",
		"rx",
		"ry",
		"fill",
		"stroke",
		"stroke-width",
	},

	"line": {
		"x1",
		"y1",
		"x2",
		"y2",
		"stroke",
		"stroke-width",
	},

	"polyline": {
		"points",
		"fill",
		"stroke",
		"stroke-width",
	},

	"polygon": {
		"points",
		"fill",
		"stroke",
		"stroke-width",
	},

	"ellipse": {
		"cx",
		"cy",
		"rx",
		"ry",
		"fill",
		"stroke",
		"stroke-width",
	},

	"text": {
		"x",
		"y",
		"fill",
		"font-size",
		"text-anchor",
	},

	"tspan": {
		"x",
		"y",
		"dx",
		"dy",
		"fill",
	},

	"defs":  {},
	"title": {},
	"desc":  {},
}
