package linkify

import (
	"reflect"
	"testing"
)

func TestLinks(t *testing.T) {
	type testCase struct {
		in   string
		want []Link
	}
	testCases := []testCase{
		{"", nil},

		// 1

		{"127.0.0.1:8080/path?query=xxx#fragment", []Link{{Schema: "", Start: 0, End: 38}}},
		{"127.0.0.1", []Link{{Schema: "", Start: 0, End: 9}}},
		{"xxx 127.0.0.1", []Link{{Schema: "", Start: 4, End: 13}}},
		{"1.0001.0.1", nil},
		{"1234.0.0", nil},
		{"1.2345.0.1", nil},
		{"1.2.3.4x", nil},
		{"_127.0.0.1", nil},
		{"127.0.0_1", nil},
		{"127.0.0.1/%", nil},
		{"127.0.0.1%", nil},
		{"127.0.0.1/\x00", nil},
		{"127.0.0.1\x00", nil},
		{"127.0.0._", nil},
		{"127.0.0", nil},
		{"127.0._.1", nil},
		{" _.127.0", nil},
		{" .127._", nil},
		{" .127.", nil},
		{" .127.X", nil},
		{":234.0.0", nil},
		{"256.0.0.1", nil},
		{" X.127.0", nil},
		{"x234.0.0", nil},

		// 2

		{"google.com!", []Link{{Schema: "", Start: 0, End: 10}}},
		{"google.com...", []Link{{Schema: "", Start: 0, End: 10}}},
		{"google.com.", []Link{{Schema: "", Start: 0, End: 10}}},
		{"google.com/?search", []Link{{Schema: "", Start: 0, End: 18}}},
		{"some.host.name.com.net.abracadabra", nil},
		{"some.host.name.com.net.org", []Link{{Schema: "", Start: 0, End: 26}}},
		{"some.host.name.com.net.org/", []Link{{Schema: "", Start: 0, End: 27}}},
		{"some.host.name.com.net.org/path?query=xxx#fragment", []Link{{Schema: "", Start: 0, End: 50}}},
		{"ya.ru", []Link{{Schema: "", Start: 0, End: 5}}},
		{"a.r", nil},
		{".com", nil},
		{".COM", nil},
		{"-google.com", nil},
		{"goo---gle.com", nil},
		{"google-.com", nil},
		{"google.com/%", nil},
		{"google.com/\x00", nil},
		{"google.com\x00", nil},
		{"google.comXXX", nil},
		{"GOOGLE.COMXXX", nil},
		{"images-.google.com", nil},
		{"\ufffdgoogle.com", nil},
		{"\x00.com", nil},
		{"\x00.COM", nil},
		{"\x00google.com", nil},
		{"xxx.com.abracadabra", nil},
		{"XXX.ZZ", nil},

		// 3

		{"a@a.ru", []Link{{Schema: "mailto:", Start: 0, End: 6}}},
		{"a.b@golang.org", []Link{{Schema: "mailto:", Start: 0, End: 14}}},
		{"a..b@golang.org", []Link{{Schema: "", Start: 5, End: 15}}},
		{"r@golang.org", []Link{{Schema: "mailto:", Start: 0, End: 12}}},
		{".r@golang.org", []Link{{Schema: "", Start: 3, End: 13}}},
		{"@r@golang.org", []Link{{Schema: "", Start: 3, End: 13}}},
		{"r.@golang.org", []Link{{Schema: "", Start: 3, End: 13}}},
		{"\ufffdr@golang.org", []Link{{Schema: "", Start: 5, End: 15}}},
		{"XXX r@golang.org", []Link{{Schema: "mailto:", Start: 4, End: 16}}},
		{"XXX .r@golang.org", []Link{{Schema: "", Start: 7, End: 17}}},
		{"a@a.r", nil},
		{"A@A.R", nil},
		{"@golang", nil},
		{"@GOOGLE", nil},
		{"r@golang.", nil},
		{"r@golang", nil},
		{"r@golang.zzz", nil},
		{"r@", nil},
		{"ROOT@", nil},
		{"r@\x00golang", nil},
		{"r\x00@golang", nil},
		{"R@\x00GOLANG", nil},
		{"R\x00@GOLANG", nil},

		// 4

		{"//127.0.0.1:80/", []Link{{Schema: "//", Start: 0, End: 15}}},
		{"//google.com.", []Link{{Schema: "//", Start: 0, End: 12}}},
		{"//ya.ru", []Link{{Schema: "//", Start: 0, End: 7}}},
		{"://google.com", nil},
		{"//google.com%", nil},
		{"//google.com\x00", nil},
		{"//google", nil},
		{"//google.zzz", nil},
		{"//\x00google", nil},
		{"x//google.com", nil},

		// 5

		{"mailto:a.b.c@golang.org", []Link{{Schema: "mailto:", Start: 0, End: 23}}},
		{"mailto:a..b.c@golang.org", []Link{{Schema: "", Start: 14, End: 24}}},
		{"mailto:r@golang.org", []Link{{Schema: "mailto:", Start: 0, End: 19}}},
		{"mailto:r.@golang.org", []Link{{Schema: "", Start: 10, End: 20}}},
		{"mailto:xxxЖ@golang.org", []Link{{Schema: "", Start: 13, End: 23}}},
		{"ailto:xxxx", nil},
		{"AILTO:XXXX", nil},
		{"mailto:a@a.a", nil},
		{"MAILTO:A@A.A", nil},
		{"mailto:a..b.c@golang", nil},
		{"mailto:a..b.cgolang.org", nil},
		{"mailto:r@golang", nil},
		{"mailto:rgolangorg", nil},
		{"mailto:r@golang.zzz", nil},
		{"mailto:r@gol", nil},
		{"mailto:r@", nil},
		{"mailto:r@xxx", nil},
		{"mailto:\x00myemail", nil},
		{"MAILTO:\x00MYEMAIL", nil},
		{"mailto:xxx@", nil},
		{"xmailto:myemail", nil},
		{"XMAILTO:MYEMAIL", nil},
		{"xxxxxo:myemail", nil},
		{"XXXXXO:MYEMAIL", nil},
		{"xailto:myemail", nil},

		// 6

		{"http://127.0.0.1:80/", []Link{{Schema: "http:", Start: 0, End: 20}}},
		{"http://google.com/[1(2", []Link{{Schema: "http:", Start: 0, End: 18}}},
		{"http://google.com/#[1(2", []Link{{Schema: "http:", Start: 0, End: 19}}},
		{"http://google.com/[1](2)", []Link{{Schema: "http:", Start: 0, End: 24}}},
		{"http://google.com/#[1](2)", []Link{{Schema: "http:", Start: 0, End: 25}}},
		{"http://google.com/[1)", []Link{{Schema: "http:", Start: 0, End: 18}}},
		{"http://google.com/#[1)", []Link{{Schema: "http:", Start: 0, End: 19}}},
		{"http://google.com ", []Link{{Schema: "http:", Start: 0, End: 17}}},
		{"http://google.com--", []Link{{Schema: "http:", Start: 0, End: 17}}},
		{"http://google.com-", []Link{{Schema: "http:", Start: 0, End: 17}}},
		{"http://google.com!", []Link{{Schema: "http:", Start: 0, End: 17}}},
		{"http://google.com...", []Link{{Schema: "http:", Start: 0, End: 17}}},
		{"http://google.com.", []Link{{Schema: "http:", Start: 0, End: 17}}},
		{"HtTp://gOoGlE.CoM", []Link{{Schema: "http:", Start: 0, End: 17}}},
		{"http://google.com/--", []Link{{Schema: "http:", Start: 0, End: 18}}},
		{"http://google.com/", []Link{{Schema: "http:", Start: 0, End: 18}}},
		{`"http://google.com"`, []Link{{Schema: "http:", Start: 1, End: 18}}},
		{"(http://google.com)", []Link{{Schema: "http:", Start: 1, End: 18}}},
		{"http://google.com\n", []Link{{Schema: "http:", Start: 0, End: 17}}},
		{"http://google.com/?query=[1(2", []Link{{Schema: "http:", Start: 0, End: 25}}},
		{"http://google.com/?query=[1](2)", []Link{{Schema: "http:", Start: 0, End: 31}}},
		{"http://google.com/?query=[1)", []Link{{Schema: "http:", Start: 0, End: 25}}},
		{"http://google.com/?query=)x(", []Link{{Schema: "http:", Start: 0, End: 25}}},
		{"http://google.com/?query=]xxx", []Link{{Schema: "http:", Start: 0, End: 25}}},
		{"http://google.com/search---", []Link{{Schema: "http:", Start: 0, End: 24}}},
		{"http://google.com/search?query=xxx---", []Link{{Schema: "http:", Start: 0, End: 34}}},
		{"http://google.com/#toc---", []Link{{Schema: "http:", Start: 0, End: 22}}},
		{"http://google.com/]xxx", []Link{{Schema: "http:", Start: 0, End: 18}}},
		{"http://google.com/#]xxx", []Link{{Schema: "http:", Start: 0, End: 19}}},
		{"http://www.youtube.com/watch?v=EIBRdBVkDHQ.", []Link{{Schema: "http:", Start: 0, End: 42}}},
		{"http://goo---gle.com", nil},
		{"http://google-.com", nil},
		{"http://google.com\\", nil},
		{"http:google.com", nil},
		{"http://google.com/search?q%0z", nil},
		{"http://google.com/search?q%1", nil},
		{"http://google.com/search?q%", nil},
		{"http://google.com/search?q\ufffd", nil},
		{"http://google.com/search?q=x%78x?/:@#toc\ufffd", nil},
		{"http://google.com/search?q%z0", nil},
		{"http://google.com\x00", nil},
		{"http://google", nil},
		{"http://goo\ufffd", nil},
		{"http://images.google.com.-net", nil},
		{"http://images.google.com.net.-org", nil},
		{"http://\x00google", nil},
		{"HTTP://\x00GOOGLE", nil},
		{"http://xxx.com.abracadabra", nil},
		{"http://xxx", nil},
		{"HTTP://XXX", nil},
		{"http:xxxxxx", nil},
		{"HTTP:XXXXXX", nil},
		{"TP:GOOGLE", nil},
		{"ttp://google", nil},
		{"TTP://GOOGLE", nil},
		{"xhttp://google", nil},
		{"XHTTP://GOOGLE", nil},
		{"xttp://google", nil},
		{"XTTP://GOOGLE", nil},

		// 7

		{"https://google.com/p%61th?%71uery#%66ragment", []Link{{Schema: "https:", Start: 0, End: 44}}},
		{"https://\x00google", nil},
		{"HTTPS://\x00GOOGLE", nil},
		{"https:XXgoogle", nil},
		{"HTTPS:XXGOOGLE", nil},
		{"https://xxx", nil},
		{"HTTPS://XXX", nil},
		{"ttps://google", nil},
		{"TTPS://GOOGLE", nil},
		{"xhttps://google", nil},
		{"XHTTPS://GOOGLE", nil},
		{"xttps://google", nil},
		{"XTTPS://GOOGLE", nil},

		// 8

		{"ftp://google.com", []Link{{Schema: "ftp:", Start: 0, End: 16}}},
		{"fXp://google", nil},
		{"FXP://GOOGLE", nil},
		{"tp:google", nil},
		{"xtp://google", nil},
		{"XTP://GOOGLE", nil},

		// 9

		{"http://localhost:80/", []Link{{Schema: "http:", Start: 0, End: 20}}},
		{"localhost:3128/path?query=xxx#fragment", []Link{{Schema: "", Start: 0, End: 38}}},
		{"//localhost:80/", []Link{{Schema: "//", Start: 0, End: 15}}},
		{"localhost:80/!$&'()*+,;=/:@", []Link{{Schema: "", Start: 0, End: 27}}},
		{"localhost:80/#?!$&'()*+,;=/:@", []Link{{Schema: "", Start: 0, End: 29}}},
		{"localhos", nil},
		{"localhost:0", nil},
		{"localhost:65536", nil},
		{"localhost:80/#fragment%0", nil},
		{"localhost:80/#fragment%0z", nil},
		{"localhost:80/#fragment%", nil},
		{"localhost:80/#fragment\x00", nil},
		{"localhost:80/#fragment%z0", nil},
		{"LOCALHOST:80", nil},
		{"localhost:80/path%0", nil},
		{"localhost:80/path%0z", nil},
		{"localhost:80/path%", nil},
		{"localhost:80/path%z0", nil},
		{"localhost:80/\ufffd", nil},
		{"localhost:80\ufffd", nil},
		{"localhost:", nil},
		{"localhost:x", nil},
		{"localhost", nil},
		{"localhostX:80", nil},
		{"Xlocalhost:80", nil},

		// 10

		{":00", nil},
		{"skype://nickname", nil},
		{"SKYPE://NICKNAME", nil},
		{"XXX http://google.com r@golang.org localhost:80 mailto:r@golang.org //google.com XXX", []Link{Link{Schema: "http:", Start: 4, End: 21}, Link{Schema: "mailto:", Start: 22, End: 34}, Link{Schema: "", Start: 35, End: 47}, Link{Schema: "mailto:", Start: 48, End: 67}, Link{Schema: "//", Start: 68, End: 80}}},
	}
	for _, tc := range testCases {
		got := Links(tc.in)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Links(%q):\n got %#v\nwant %#v", tc.in, got, tc.want)
		}
	}
}

func TestParseIPv4(t *testing.T) {
	type testCase struct {
		in     string
		length int
		ok     bool
	}
	testCases := []testCase{
		{"", 0, false},
		{"8.8.8.8", 7, true},
		{"8.8.8.", 0, false},
		{"8.8.8.xxx", 0, false},
		{"8.8.8.8xxx", 7, true},
		{"256.0.0.1", 0, false},
		{"001.001.001.001", 15, true},
	}
	for _, tc := range testCases {
		length, ok := skipIPv4(tc.in)
		if ok != tc.ok {
			s := "failed"
			if !tc.ok {
				s = "unexpectedly succeeded"
			}
			t.Errorf("skipIPv4(%q) %s", tc.in, s)
		} else if length != tc.length {
			t.Errorf("skipIPv4(%q) returned length %d, want %d", tc.in, length, tc.length)
		}
	}
}
