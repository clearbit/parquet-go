package goparquet

import "testing"

func TestFuzzCrashDeltaBitPackDecoder64DivByZero(t *testing.T) {
	data := []byte("PAR1\x15\x00\x15B\x15B,\x15\x02\x15\n\x15\x06\x15\x15\x02" +
		"\x19\xf0)H!org.apache.impa" +
		"la.ComplexTypesTbl\x15\f" +
		"\x00\x15\x04%\x00\x18\x02id\x004\x02\x18\tint_ar" +
		"ray\x15\x02\x15\x06\x005\x04\x18\x04list\x15\x02\x00\x15" +
		"\x02%\x02\x18\aelement\x005\x02\x18\x0fint" +
		"_array_Array\x15\x02\x15\x06\x005\x04\x18" +
		"\x04list\x15\x02\x005\x02\x18\aelement\x15" +
		"\x02\x15\x06\x005\x04\x18\x04list\x15\x02\x00\x15\x02%\x02\x18" +
		"\aelement\x005\x02\x18\ain\x00\x00\x00\x01p" +
		"\x15\x02\x15\x02\x005\x04\x18\x03map\x15\x04\x15\x04\x00\x15\f%" +
		"\x00\x18\x03key%\x00\x00\x15\x02%\x02\x18\x05value" +
		"\x005\x02\x18\rint_Map_Array\x15\x02" +
		"\x15\x06\x005\x04\x18\x04list\x15\x02\x005\x02\x18\ael" +
		"ement\x15\x02\x15\x02\x005\x04\x18\x03map\x15\x04\x15" +
		"\x04\x00\x15\f%\x00\x18\x03key%\x00\x00\x15\x02%\x02\x18\x05" +
		"value\x005\x02\x18\rnested_str" +
		"uct\x15\b\x00\x15\x02%\x02\x18\x01A\x005\x02\x18\x01b\x15" +
		"\x02\x15\x06\x005\x04\x18\x04list\x15\x02\x00\x15\x02%\x02\x18" +
		"\aelement\x005\x02\x18\x01C\x15\x02\x005\x02\x18" +
		"\x01d\x15\x02\x15\x06\x005\x04\x18\x04list\x15\x02\x005\x02" +
		"\x18\aelement\x15\x02\x15\x06\x005\x04\x18\x04li" +
		"st\x15\x02\x005\x02\x18\aelement\x15\x04\x00\x15" +
		"\x02%\x02\x18\x01E\x00\x15\f%\x02\x18\x01F%\x00\x005\x02\x18" +
		"\x01g\x15\x02\x15\x02\x005\x04\x18\x03map\x15\x04\x15\x04\x00\x15" +
		"\f%\x00\x18\x03key%\x00\x005\x02\x18\x05value" +
		"\x15\x02\x005\x02\x18\x01H\x15\x02\x005\x02\x18\x01i\x15\x02\x15\x06" +
		"\x005\x04\x18\x04list\x15\x02\x00\x15\n%\x02\x18\ael" +
		"ement\x00\x16\x0e\x19\x1c\x19\xdc&\b\x1c\x15\x04\x195\b" +
		"\x00\x06\x19\x18\x02id\x15\x00\x16\x0e\x16\xce\x01\x16\xce\x01&\b<" +
		"\x18\b\a\x00\x00\x00\x00\x00\x00\x00\x18\b\x01\x00\x00\x00\x00\x00\x00\x00" +
		"\x16\x00\x00\x00\x00&\xd6\x01\x1c\x15\x02\x19%\x06\x04\x198\tin" +
		"t_array\x04list\aelement" +
		"\x15\x00\x16\x1c\x16\x9c\x01\x16\x9c\x01&\xd6\x01<\x18\x04\x03\x00\x00\x00" +
		"\x18\x04\x01\x00\x00\x00\x16\x10\x00\x00\x00&\xf2\x02\x1c\x15\x02\x19%\x06" +
		"\x04\x19X\x0fint_array_Array\x04" +
		"list\aelement\x04list\ael" +
		"ement\x15\x00\x16(\x16\xce\x01\x16\xce\x01&\xf2\x02<\x18" +
		"\x04\x06\x00\x00\x00\x18\x04\x01\x00\x00\x00\x16\x14\x00\x00\x00&\xc0\x04\x1c" +
		"\x15\f\x19%\x06\x04\x198\aint_map\x03map" +
		"\x03key\x15\x00\x16\x14\x16\xa0\x01\x16\xa0\x01&\xc0\x04<\x18\x02" +
		"k3\x18\x02k1\x16\b\x00\x00\x00&\xe0\x05\x1c\x15\x02\x19%\x00" +
		"\x06\x198\aint_map\x03map\x05valu" +
		"e\x15\x00\x16\x14\x16z\x16z&\xe0\x05<\x18\x04d\x00\x00\x00\x18" +
		"\x04\x01\x00\x00\x00\x16\x0e\x00\x00\x00&\xda\x06\x1c\x15\f\x19%\x06\x04" +
		"\x19X\rint_Map_Array\x04lis" +
		"t\aelement\x03map\x03key\x15\x00\x16" +
		"\x16\x16\x9a\x01\x16\x9a\x01&\xda\x06<\x18\x02k3\x18\x02k1\x16" +
		"\x10\x00\x00\x00&\xf4\a\x1c\x15\x02\x19%\x06\x04\x19X\rint" +
		"_Map_Array\x04list\aele-" +
		"ent\x03map\x05value\x15\x00\x16\x16\x16\x90\x01" +
		"\x16\x90\x01&\xf4\a<\x18\x04\x01\x00\x00\x00\x18\x04\x01\x00\x00\x00\x16" +
		"\x12\x00\x00\x00&\x84\t\x1c\x15\x02\x195\b\x00\x06\x19(\rne" +
		"sted_struct\x01A\x15\x00\x16\x0e\x16`\x16" +
		"`&\x84\t<\x18\x04\a\x00\x00\x00\x18\x04\x01\x00\x00\x00\x16\n\x00" +
		"\x00\x00&\xe4\t\x1c\x15\x02\x19%\x00\x06\x19H\rneste" +
		"d_struct\x01b\x04list\aelem" +
		"ent\x15\x00\x16\x12\x16~\x16~&\xe4\t<\x18\x04\x03\x00\x00" +
		"\x00\x18\x04\x01\x00\x00\x00\x16\f\x00\x00\x00&\xe2\n\x1c\x15\x02\x19%" +
		"\x06\x04\x19\x88\rnested_struct\x01C" +
		"\x01d\x04list\aelement\x04list" +
		"\ael\xecment\x01E\x15\x00\x16&\x16\xb4\x01\x16\xb4\x01" +
		"&\xe2\n<\x18\x04\v\x00\x00\x00\x18\x04\xf6\xff\xff\xff\x16\x1a\x00\x00" +
		"\x00&\x96\f\x1c\x15\f\x19%\x06\x04\x19\x88\rnested" +
		"_struct\x01C\x01d\x04list\aele" +
		"ment\x04list\aelement\x01F\x15" +
		"\x00\x16&\x16\xba\x01\x16\xba\x01&\x96\f<\x18\x01c\x18\x03aa" +
		"a\x16\x1a\x00\x00\x00&\xd0\r\x1c\x15\f\x19%\x06\x04\x19H\rn" +
		"ested_struct\x01g\x03map\x03k" +
		"ey\x15\x00\x16\x16\x16\xca\x01\x16\xca\x01&\xd0\r<\x18\x02g5" +
		"\x18\x03foo\x16\b\x00\x00\x00&\x9a\x0f\x1c\x15\n\x19%\x06\x04" +
		"\x19\x88\rnested_struct\x01g\x03m" +
		"ap\x05value\x01H\x01i\x04list\ael" +
		"ement\x15\x00\x16\x1a\x16\xd0\x01\x16\xd0\x01&\x9a\x0f<\x18" +
		"\bffffff\n@\x18\b\x9a\x99\x99\x99\x99\x99\xf1?\x16" +
		"\x12\x00\x00\x00\x16\xe2\x10\x16\x0e\x00\x19\x1c\x18\x13pa:que" +
		"t.avro.schema\x18\xc0\t{\"ty" +
		"pe\":\"record\",\"name\":" +
		"\"ComplexTypesTbl\",\"n" +
		"amespace\":\"org.apach" +
		"e.impala\",\"fields\":[" +
		"{\"name\":\"id\",\"type\":" +
		"[\"null\",\"long\"]},{\"n" +
		"ame\":\"int_array\",\"ty" +
		"pe\":[\"null\",{\"type\":" +
		"\"array\",\"ioems\":[\"nu" +
		"ll\",\"int\"]}]},{\"name" +
		"\":\"int_array_Array\"," +
		"\"type\":[\"null\",{\"typ" +
		"e\":\"array\",\"items\":[" +
		"\"null\",{\"type\":\"arra" +
		"y\",\"items\":[\"null\",\"" +
		"int\"]}]}]},{\"name\":\"" +
		"int_map\",\"type\":[\"nu" +
		"ll\",{\"type\":\"map\",\"v" +
		"alues\":[\"null\",\"int\"" +
		"]}]},{\"na\xfbe\":\"int_Ma" +
		"p_Array\",\"type\":[\"nu" +
		"ll\",{\"type\":\"array\"," +
		"\"items\":[\"null\",{\"ty" +
		"pe\":\"map\",\"values\":[" +
		"\"null\",\"int\"]}]}]},{" +
		"\"name\":\"nested_struc" +
		"t\",\"type\":[\"null\",{\"" +
		"type\":\"record\",\"name" +
		"\":\"r1\",\"fields\":[{\"n" +
		"ame\":\"A\",\"type\":[ nu" +
		"ll\",\"int\"]},{\"name\":" +
		"\"b\",\"type\":[\"null\",{" +
		"\"type\":\"array\",\"item" +
		"s\":[\"nulc\",\"int\"]}]}" +
		",{\"name\":\"C\",\"type\":" +
		"[\"null\",{\"type\":\"rec" +
		"ord\",\"name\":\"r2\",\"fi" +
		"elds\":[{\"name\":\"d\",\"" +
		"type\":[\"null\",{\"type" +
		"\":\"array\",\"items\":[\"" +
		"null\",{\"type\":\"array" +
		"\",\"items\":[\"null\",{\"" +
		"type\":\"record\",\"name" +
		"\":\"r3\",\"fields\":[{\"n" +
		"ame\":\"E\",\"type\":[\"nu" +
		"ll\",\"int\"]},{\"name\":" +
		"\"F\",\"type\":[\"null\",\"" +
		"string\"]}]}]}]}]}]}]" +
		"},{\"name\":\"g\",\"type\"" +
		":[\"null\",{\"type\":\"ma" +
		"p\",\"values\":[\"null\"," +
		"{\"type\":\"record\",\"na" +
		"me\":\"r4\",\"fields\":[{" +
		"\"name\":\"H\",\"type\":[\"" +
		"null\",{\"type\":\"recor" +
		"d\",\"name\":\"r5\",\"fiel" +
		"ds\":[{\"name\":\"i\",\"ty" +
		"pe\":[\"nul\xef\xff\xff\xff\xff\xff\xff\xffe\":" +
		"\"array\",\"items\":[\"nu" +
		"ll\",\"double\"]}]}]}]}" +
		"]}]}]}]}]}]}\x00\x18Iparqu" +
		"et-mr version 1.8.0 " +
		"(bui,d 0fda28af84b97" +
		"46396014ad6a415b9059" +
		"2a98b3b)\x00\xfb\n\x00\x00PAR1")

	readAllData(t, data)
}
