module github.com/MrHuxu/react-go-ssr-boilerplate

go 1.12

require (
	github.com/dlclark/regexp2 v1.1.6 // indirect
	github.com/dop251/goja v0.0.0-20190611173659-6d56e81e6bf4
	github.com/gin-gonic/gin v1.4.0
	github.com/go-sourcemap/sourcemap v2.1.2+incompatible // indirect
)

replace (
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2 => github.com/golang/crypto v0.0.0-20190611184440-5c40567a22f8
	golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3 => github.com/golang/net v0.0.0-20190611141213-3f473d35a33a
	golang.org/x/net v0.0.0-20190503192946-f4e77d36d62c => github.com/golang/net v0.0.0-20190611141213-3f473d35a33a
	golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a => github.com/golang/sys v0.0.0-20190613124609-5ed2794edfdc
	golang.org/x/sys v0.0.0-20190222072716-a9d3bda3a223 => github.com/golang/sys v0.0.0-20190613124609-5ed2794edfdc
	golang.org/x/sys v0.0.0-20190412213103-97732733099d => github.com/golang/sys v0.0.0-20190613124609-5ed2794edfdc
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.2
	golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e => github.com/golang/tools v0.0.0-20180917221912-90fa682c2a6e
)
