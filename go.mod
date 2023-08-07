module giao

go 1.19

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/common v0.4.0
	github.com/qianlnk/pgbar v0.0.0-20210208085217-8c19b9f2477e
	github.com/schollz/progressbar/v3 v3.12.1
	github.com/spf13/viper v1.15.0 // indirect
)

require (
	github.com/alecthomas/template v0.0.0-20160405071501-a0175ee3bccc // indirect
	github.com/alecthomas/units v0.0.0-20151022065526-2efee857e7cf // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/qianlnk/to v0.0.0-20191230085244-91e712717368 // indirect
	github.com/rivo/uniseg v0.4.3 // indirect
	github.com/rogpeppe/go-internal v1.8.0 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/spf13/afero v1.9.3 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/term v0.1.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	gopkg.in/alecthomas/kingpin.v2 v2.2.6 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require giao/pkg/util v0.0.0

replace giao/pkg/util => ../pkg/util
