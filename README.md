# nippon-colors
Traditional Japan colors for 24bit terminal

## Overview

[Nippon colors](https://nipponcolors.com) is a collection of colors traditionally used in Japanese arts and crafts. This library exposes those colors for 24bit terminal output as a collection of wrappers of [gookit.Color](https://github.com/gookit/color) types. The color list and data taken from [nipponcolor](https://github.com/ssshooter/nippon-color) PWA repository.

## Usage

```
import (
	color "github.com/eiri/nippon-colors"
)

func main() {
    color.Tsutsuji.Print("Azalea")
}
```

## License

[MIT](https://github.com/eiri/nippon-colors/blob/main/LICENSE)
