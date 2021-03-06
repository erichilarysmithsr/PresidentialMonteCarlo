//
// college.go
// 

package main

type College struct {
	votes   int
	dem2008 bool
}

var college = map[string]College{
	"AL": College{9, false},
	"AK": College{3, false},
	"AZ": College{11, false},
	"AR": College{6, false},
	"CA": College{55, true},
	"CO": College{9, true},
	"CT": College{7, true},
	"DE": College{3, true},
	"DC": College{3, true},
	"FL": College{29, true},
	"GA": College{16, false},
	"HI": College{4, true},
	"ID": College{4, false},
	"IL": College{20, true},
	"IN": College{11, true},
	"IA": College{6, true},
	"KS": College{6, false},
	"KY": College{8, false},
	"LA": College{8, false},
	"ME": College{4, true},
	"MD": College{10, true},
	"MA": College{11, true},
	"MI": College{16, true},
	"MN": College{10, true},
	"MS": College{6, false},
	"MO": College{10, false},
	"MT": College{3, false},
	"NE": College{5, false},
	"NV": College{6, true},
	"NH": College{4, true},
	"NJ": College{14, true},
	"NM": College{5, true},
	"NY": College{29, true},
	"NC": College{15, true},
	"ND": College{3, false},
	"OH": College{18, true},
	"OK": College{7, false},
	"OR": College{7, true},
	"PA": College{20, true},
	"RI": College{4, true},
	"SC": College{9, false},
	"SD": College{3, false},
	"TN": College{11, false},
	"TX": College{38, false},
	"UT": College{6, false},
	"VT": College{3, true},
	"VA": College{13, true},
	"WA": College{12, true},
	"WV": College{5, false},
	"WI": College{10, true},
	"WY": College{3, false},
}
