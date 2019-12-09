# img
golang image to gray,binaryzation,reverse color


### install

```code
go get github.com/no-bibi/img
```

### demo

```go
package main

import (
	"github.com/no-bibi/img"
)

func main() {

	file, _ := img.Read(`./enter.png`)

	//to make it gray
	grayed := img.Gray(file)
	img.Encode(grayed,`./enter_grap.png`)

	//want binaryzation ? not problem
	binaryzation := img.Binaryzation(grayed, 127)
	img.Encode(binaryzation,`./enter_binaryzation.png`)

	//want reverse ? go on with Reverse
	reverse := img.Reverse(binaryzation)
	img.Encode(reverse,`./enter_reverse.png`)
}
```

### work out

<img src="source/enter.png" /><br>original
<img src="build/enter_grap.png" /><br>gray
<img src="build/enter_binaryzation.png" /><br>binaryzation
<img src="build/enter_reverse.png" /><br>reverse

### licenses

[MIT](http://opensource.org/licenses/MIT)

