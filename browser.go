package uaparser

var (
    ie = &itemSpec {
        name: "IE",
        mustContains: []string{"MSIE"},
        mustNotContains: []string{
            "360SE",
            "Maxthon",
            "qihu",
            "QIHU",
            "QQBrowser",
            "QQDownload",
            "SE ", "MetaSr",
            "TencentTraveler",
        },
        versionSplitters: [][]string{[]string{"MSIE ", ";"}},
    }

    firefox = &itemSpec {
        name: "Firefox",
        mustContains: []string{"Firefox"},
        mustNotContains: []string{"Seamonkey"},
        versionSplitters: [][]string{[]string{"Firefox/", " "}},
    }

    safari = &itemSpec {
        name: "Safari",
        mustContains: []string{"Safari"},
        mustNotContains: []string{
            "Chrome", "Chromium",
            "CoolNovo",
            "Maxthon",
            "LBBROWSER",
            "QIHU",
            "QQBrowser",
            "SE ", "MetaSr",
            "TaoBrowser",
        },
        versionSplitters: [][]string{
            []string{"Version/", " "},
        },
    }

    chrome = &itemSpec {
        name: "Chrome",
        mustContains: []string{"Chrome"},
        mustNotContains: []string{
            "Chromium",
            "CoolNovo",
            "LBBROWSER",
            "Maxthon",
            "QIHU",
            "QQBrowser",
            "SE ", "MetaSr",
            "TaoBrowser",
        },
        versionSplitters: [][]string{[]string{"Chrome/", " "}},
    }

    opera = &itemSpec {
        name: "Opera",
        mustContains: []string{"Opera"},
        mustNotContains: []string{},
        versionSplitters: [][]string{
            []string{"Version/", " "},
            []string{"Opera/", " "},
        },
    }

    _360se = &itemSpec {
        name: "360SE",
        mustContains: []string{"360SE", "qihu", "QIHU"},
        mustNotContains: []string{},
    }

    sougou = &itemSpec {
        name: "Sougou",
        mustContains: []string{"SE ", "MetaSr"},
        mustNotContains: []string{},
        versionSplitters: [][]string{[]string{"SE ", " "}},
    }

    tencent = &itemSpec {
        name: "Tencent Traveler",
        mustContains: []string{"TencentTraveler"},
        mustNotContains: []string{"SE ", "MetaSr"},
        versionSplitters: [][]string{[]string{"TencentTraveler ", " "}},
    }

    qq = &itemSpec {
        name: "QQ Browser",
        mustContains: []string{"QQBrowser"},
        mustNotContains: []string{},
        versionSplitters: [][]string{[]string{"QQBrowser/", " "}},
    }

    maxthon = &itemSpec {
        name: "Maxthon",
        mustContains: []string{"Maxthon"},
        mustNotContains: []string{},
        versionSplitters: [][]string{
            []string{"Maxthon/", " "},
            []string{"Maxthon ", " "},
        },
    }

    _BROWSERS = []*itemSpec {
        ie,
        firefox,
        safari,
        chrome,
        opera,
        _360se,
        sougou,
        tencent,
        qq,
        maxthon,
    }
)
