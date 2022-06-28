<a name="unreleased"></a>
## [Unreleased]

### Chore
- **deps:** update module yaml to v2.2.8 ([#231](https://github.com/julien-noblet/download-geofabrik/issues/231))
- **deps:** update module stretchr/testify to v1.5.1
- **deps:** update module spf13/viper to v1.6.2
- **deps:** update module puerkitobio/goquery to v1.5.1 ([#234](https://github.com/julien-noblet/download-geofabrik/issues/234))
- **deps:** update module cheggaaa/pb/v3 to v3.0.4 ([#229](https://github.com/julien-noblet/download-geofabrik/issues/229))


<a name="v2.6.0-rc1"></a>
## [v2.6.0-rc1] - 2020-03-14
### Chore
- **deps:** update module apex/log to v1.1.2
- **deps:** update module yaml to v2.2.7 ([#220](https://github.com/julien-noblet/download-geofabrik/issues/220))
- **deps:** update module olekukonko/tablewriter to v0.0.3 ([#221](https://github.com/julien-noblet/download-geofabrik/issues/221))


<a name="v2.5.1-rc1"></a>
## [v2.5.1-rc1] - 2019-11-26

<a name="v2.5.1-rc0"></a>
## [v2.5.1-rc0] - 2019-11-16
### Chore
- **deps:** update module olekukonko/tablewriter to v0.0.2

### Doc
- Updating Changelog


<a name="v2.5.0"></a>
## [v2.5.0] - 2019-11-05
### Actions
- Add benchmarking ([#212](https://github.com/julien-noblet/download-geofabrik/issues/212))

### Build
- add go mod tidy in .pre-commit
- Readme should be started after a Deployment
- compress using upx some binaries

### Chore
- **deps:** update module yaml to v2.2.5
- **deps:** update module spf13/viper to v1.5.0
- **deps:** update module cheggaaa/pb/v3 to v3.0.2

### Ci
- remove benchmarks from coverage
- cache build & tests
- tests don't need to be verbose
- coverall wasn't getting good coverage by mistake
- coverall wasn't getting good coverage by mistake
- **travis:** Fix build stage
- **travis:** download gocoverall only on gocoverall job
- **travis:** remove ymls runs on tag

### Docker
- better doc, use ENV for output dir

### Feat
- Add output-dir option, can also be send by ENV

### Fix
- progress should be quiet
- updating tj/assert to fix license issue detected by fossa.
- Add Docker stuff in .travis and fix Dokcerfile

### Gogenerate
- Fix Snap builds

### Goreleaser
- remove debug sybols from binaries
- Add Docker support
- Add arm64 to SnapCraft
- Fix rights for Snapcraft
- Don't generate yaml/readme on tag
- Use zip for Windows, rename Darwin=>MacOS
- Generate yaml files and Readme on build
- Use ldflags

### Licences
- Add Licence in all .go files

### Log
- use apex/log to get better logs

### Refactor
- Spliting into packages

### Test
- revert use of stretchr/testify
- add configureBool test
- add generator.write test
- removing duplicate test
- add Generate tests
- fix LoadConfig tests
- add meta in config.MergeElement() tests
- add parent mismatch in config.MergeElement() tests
- add config.MergeElement() tests
- add parseSubregion guatemala
- use GetDefault to create a scrapper.Scrapper object in tests
- add parseLi guatemala + fix


<a name="v2.4.1"></a>
## [v2.4.1] - 2019-10-19
### Chore
- **deps:** update module yaml to v2.2.4 ([#207](https://github.com/julien-noblet/download-geofabrik/issues/207))
- **deps:** update golang.org/x
- **deps:** update golang.org/x/tools commit hash to aed303c
- **deps:** update module google.golang.org/appengine to v1.6.2
- **deps:** update golang.org/x
- **deps:** update golang.org/x/tools commit hash to 95c3470
- **deps:** update golang.org/x/sys commit hash to fb81701
- **deps:** update golang.org/x/tools commit hash to d72b05d
- **deps:** update golang.org/x/tools commit hash to c4a336e
- **deps:** update golang.org/x/tools commit hash to aa644d2
- **deps:** update golang.org/x/tools commit hash to b1e2c8e
- **deps:** update golang.org/x/tools commit hash to fc82fb2
- **deps:** update golang.org/x
- **deps:** update golang.org/x/tools commit hash to 85edb9e
- **deps:** update golang.org/x/tools commit hash to 15fda70
- **deps:** update golang.org/x/tools commit hash to 6889da9
- **deps:** update module stretchr/testify to v1.4.0
- **deps:** update golang.org/x/tools commit hash to 9065c18
- **deps:** update golang.org/x/tools commit hash to ea41424
- **deps:** update golang.org/x/tools commit hash to 9fae7b2
- **deps:** update golang.org/x
- **deps:** update golang.org/x
- **yml:** update yml files

### Ci
- **travis:** Use GO111MODULE ([#181](https://github.com/julien-noblet/download-geofabrik/issues/181))
- **travis:** Don't build yaml files on every master commits
- **travis:** use goreleaser

### Feat
- **scrapper:** Add a timeout optional value for scrapper

### Fix
- **generator:** fix an error string


<a name="v2.4.0"></a>
## [v2.4.0] - 2019-07-06
### Build
- More details in Makefile
- **Travis:** Update to go 1.12 ([#82](https://github.com/julien-noblet/download-geofabrik/issues/82))

### Chore
- Removing yml from git repo, but keeping them into packages
- Adding lastest geofabrik.yml and openstreetmap.fr.yml
- Revert yml files must be in repo
- Updating geofabrik.yml
- fix .gitignore
- **Renovate:** auto merge update ([#84](https://github.com/julien-noblet/download-geofabrik/issues/84))
- **deps:** update golang.org/x
- **deps:** update golang.org/x
- **deps:** update golang.org/x/tools commit hash to aa740d4
- **deps:** update golang.org/x/net commit hash to afa5a82
- **deps:** update golang.org/x/tools commit hash to 685feca
- **deps:** update golang.org/x/tools commit hash to c39e774
- **deps:** update golang.org/x/sys commit hash to e8e3143
- **deps:** update golang.org/x
- **deps:** update golang.org/x
- **deps:** update golang.org/x/tools commit hash to 4ca4b55
- **deps:** update golang.org/x
- **deps:** update golang.org/x/sys commit hash to ebb4019
- **deps:** update golang.org/x
- **deps:** update golang.org/x
- **deps:** update golang.org/x/sys commit hash to 16da32b
- **deps:** update golang.org/x/sys commit hash to 9773273
- **deps:** update golang.org/x
- **deps:** update module mitchellh/gox to v1.0.1
- **deps:** update golang.org/x/tools commit hash to 7e5bf92
- **deps:** update golang.org/x/tools commit hash to 96f2e7e
- **deps:** update golang.org/x/tools commit hash to 9e54453
- **deps:** update golang.org/x/tools commit hash to e5b8258
- **deps:** update golang.org/x/tools commit hash to 12dd9f8
- **deps:** update golang.org/x/tools commit hash to 0fdf0c7
- **deps:** update golang.org/x/tools commit hash to 719e078
- **deps:** update golang.org/x/tools commit hash to 0d5674b
- **deps:** update golang.org/x
- **deps:** update golang.org/x
- **deps:** update golang.org/x
- **deps:** update golang.org/x
- **deps:** update golang.org/x/sys commit hash to e409398
- **deps:** update golang.org/x/tools commit hash to 4c644d7
- **deps:** update golang.org/x/tools commit hash to aef51cc
- **deps:** update golang.org/x/tools commit hash to a68386b
- **deps:** update golang.org/x/tools commit hash to 15bbd99
- **deps:** update golang.org/x/tools commit hash to a96101f
- **deps:** update golang.org/x/sys commit hash to 9eb1bfa
- **deps:** update golang.org/x
- **deps:** update golang.org/x
- **deps:** update golang.org/x/net commit hash to 710a502
- **deps:** update golang.org/x
- **deps:** update golang.org/x/net commit hash to 15845e8
- **deps:** update golang.org/x/tools commit hash to 1d95b17
- **deps:** update golang.org/x/tools commit hash to 5a8dccf
- **deps:** update golang.org/x/net commit hash to e3b2ff5
- **deps:** update golang.org/x
- **deps:** update golang.org/x/net commit hash to 1272bf9
- **deps:** update golang.org/x/tools commit hash to 3f1ed9e
- **deps:** update golang.org/x
- **deps:** update module kr/pty to v1.1.4
- **deps:** update golang.org/x/tools commit hash to 63e6ed9
- **deps:** update golang.org/x/sys commit hash to a2f829d
- **deps:** update golang.org/x/tools commit hash to f0bfdbf
- **deps:** update golang.org/x/tools commit hash to 8b67d36
- **deps:** update golang.org/x
- **deps:** update module mitchellh/gox to v1 ([#100](https://github.com/julien-noblet/download-geofabrik/issues/100))
- **deps:** update golang.org/x/tools commit hash to dbad8e9
- **deps:** update golang.org/x
- **deps:** update module mattn/go-isatty to v0.0.7
- **deps:** update golang.org/x
- **deps:** update golang.org/x/net commit hash to 56fb011
- **deps:** update golang.org/x
- **deps:** update golang.org/x/net commit hash to b774fd8
- **deps:** update golang.org/x/sys commit hash to 980fc43
- **deps:** update golang.org/x/tools commit hash to 00c44ba
- **deps:** update golang.org/x/sys commit hash to 584f3b1
- **deps:** update golang.org/x/tools commit hash to fd53dfa ([#87](https://github.com/julien-noblet/download-geofabrik/issues/87))
- **deps:** update golang.org/x/sys commit hash to e844e01 ([#86](https://github.com/julien-noblet/download-geofabrik/issues/86))
- **deps:** update golang.org/x/net commit hash to 16b79f2 ([#85](https://github.com/julien-noblet/download-geofabrik/issues/85))
- **deps:** update golang.org/x/sys commit hash to b688937 ([#83](https://github.com/julien-noblet/download-geofabrik/issues/83))
- **deps:** update golang.org/x/net commit hash to 92fc7df ([#76](https://github.com/julien-noblet/download-geofabrik/issues/76))
- **deps:** update golang.org/x
- **deps:** update golang.org/x/tools commit hash to 589c23e ([#75](https://github.com/julien-noblet/download-geofabrik/issues/75))
- **deps:** update golang.org/x/sys commit hash to 775f819 ([#74](https://github.com/julien-noblet/download-geofabrik/issues/74))
- **deps:** update module mattn/go-isatty to v0.0.6 ([#81](https://github.com/julien-noblet/download-geofabrik/issues/81))
- **deps:** update module mattn/go-colorable to v0.1.1 ([#79](https://github.com/julien-noblet/download-geofabrik/issues/79))
- **deps:** update module mattn/go-isatty to v0.0.5 ([#80](https://github.com/julien-noblet/download-geofabrik/issues/80))
- **deps:** update module puerkitobio/purell to v1.1.1 ([#78](https://github.com/julien-noblet/download-geofabrik/issues/78))
- **deps:** update module mattn/go-colorable to v0.1.0 ([#77](https://github.com/julien-noblet/download-geofabrik/issues/77))
- **deps:** update golang.org/x/tools commit hash to fc1d57b ([#72](https://github.com/julien-noblet/download-geofabrik/issues/72))
- **deps:** update golang.org/x/sys commit hash to 48ac38b ([#71](https://github.com/julien-noblet/download-geofabrik/issues/71))
- **deps:** update golang.org/x/net commit hash to 915654e ([#68](https://github.com/julien-noblet/download-geofabrik/issues/68))
- **deps:** update golang.org/x/net commit hash to 4829fb1
- **deps:** update module stretchr/testify to v1.3.0 ([#73](https://github.com/julien-noblet/download-geofabrik/issues/73))
- **deps:** update golang.org/x/tools commit hash to d00ac6d ([#70](https://github.com/julien-noblet/download-geofabrik/issues/70))
- **deps:** update golang.org/x/sys commit hash to b4a75ba ([#67](https://github.com/julien-noblet/download-geofabrik/issues/67))
- **deps:** update module mattn/go-runewidth to v0.0.4 ([#69](https://github.com/julien-noblet/download-geofabrik/issues/69))
- **deps:** update golang.org/x/tools commit hash to 3c39ce7 ([#66](https://github.com/julien-noblet/download-geofabrik/issues/66))
- **deps:** update golang.org/x/sys commit hash to 73d4af5 ([#65](https://github.com/julien-noblet/download-geofabrik/issues/65))
- **deps:** update golang.org/x/tools commit hash to 49db546 ([#64](https://github.com/julien-noblet/download-geofabrik/issues/64))
- **deps:** update golang.org/x/sys commit hash to b05ddf5 ([#62](https://github.com/julien-noblet/download-geofabrik/issues/62))
- **deps:** update golang.org/x/tools commit hash to 59cd96f ([#63](https://github.com/julien-noblet/download-geofabrik/issues/63))
- **deps:** update golang.org/x/tools commit hash to 6a3e9aa ([#61](https://github.com/julien-noblet/download-geofabrik/issues/61))
- **deps:** update golang.org/x/sys commit hash to 2a47403 ([#60](https://github.com/julien-noblet/download-geofabrik/issues/60))
- **deps:** update golang.org/x/sys commit hash to ad97f36 ([#59](https://github.com/julien-noblet/download-geofabrik/issues/59))
- **deps:** update golang.org/x/tools commit hash to 8634b1e ([#57](https://github.com/julien-noblet/download-geofabrik/issues/57))
- **deps:** update module cheggaaa/pb to v2 ([#43](https://github.com/julien-noblet/download-geofabrik/issues/43))
- **deps:** update golang.org/x/net commit hash to 6105869 ([#56](https://github.com/julien-noblet/download-geofabrik/issues/56))
- **deps:** update golang.org/x/tools commit hash to bcd4e47 ([#53](https://github.com/julien-noblet/download-geofabrik/issues/53))
- **deps:** update golang.org/x
- **deps:** update golang.org/x/sys commit hash to 0cf1ed9 ([#49](https://github.com/julien-noblet/download-geofabrik/issues/49))
- **deps:** update golang.org/x/sys commit hash to 93218de ([#48](https://github.com/julien-noblet/download-geofabrik/issues/48))
- **deps:** update module puerkitobio/goquery to v1.5.0 ([#47](https://github.com/julien-noblet/download-geofabrik/issues/47))
- **deps:** update golang.org/x/net commit hash to adae6a3 ([#46](https://github.com/julien-noblet/download-geofabrik/issues/46))
- **deps:** update golang.org/x/net commit hash to 03003ca ([#45](https://github.com/julien-noblet/download-geofabrik/issues/45))
- **deps:** update module olekukonko/tablewriter to v0.0.1 ([#42](https://github.com/julien-noblet/download-geofabrik/issues/42))
- **deps:** update module cheggaaa/pb to v1.0.26 ([#41](https://github.com/julien-noblet/download-geofabrik/issues/41))
- **deps:** update golang.org/x/sys commit hash to 66b7b13 ([#40](https://github.com/julien-noblet/download-geofabrik/issues/40))
- **deps:** update golang.org/x/net commit hash to 10aee18 ([#39](https://github.com/julien-noblet/download-geofabrik/issues/39))
- **deps:** update golang.org/x/sys commit hash to 18eb32c
- **deps:** update golang.org/x/tools commit hash to 9e44c1c
- **deps:** update golang.org/x/net commit hash to 351d144 ([#54](https://github.com/julien-noblet/download-geofabrik/issues/54))
- **generator:** Update progress numbers
- **go:** Go 1.7 is not longer supported.
- **packages:** Use Go 1.11's modules

### Ci
- Deploy only on lastest go version
- Generate new geofabrik.yml on each releases
- generate ymls before generating readme
- Add coveralls to check code coverage
- Add version flag, See [#27](https://github.com/julien-noblet/download-geofabrik/issues/27)
- **Coveralls.io:** add coveralls.io's badge to Readme.md
- **Travis:** Try to update travis and run default make rule
- **Travis:** Fix somes rules [no-test,yaml]
- **Travis:** Fix somes rules [no-test,yaml]
- **Travis:** fix target branch to deploy .yml files
- **Travis:** fix All branches can deploy .yml files
- **Travis:** All branches can deploy .yml files
- **Travis:** Don't generate .yml files on build
- **Travis:** Add cache ([#166](https://github.com/julien-noblet/download-geofabrik/issues/166))
- **Travis:** Use Jobs with a lot of conditional jobs
- **Travis:** Try to fix rules
- **Travis:** Don't generate .yml files for all commits!
- **Travis:** Deploy with go 1.11
- **Travis:** Try with travis stages
- **Travis:** Fix rules [no-test,yaml]
- **Travis:** All YAML can be deployed now!
- **Travis:** Don't generate .yml files for all commits!
- **Travis:** use stages
- **Travis:** Don't generate .yml files on build
- **Travis:** All branches can deploy .yml files
- **Travis:** Auto upload files from travis-ci
- **YAML:** Try to auto-update YAML files

### Coverage
- add some tests
- add some tests
- format.go -> 100%

### Docs
- Updating Readme
- Updating README.md
- Updating readme
- Generate README.md Add CHANGELOG.md in zip files
- Updating readme
- Updating Changelog using github-changelog-generator
- **Gislab:** Add gislab in help

### Feat
- Handle error from download, create file only after the download start without issue.
- Splitting gofiles ang adding generate to download in geofabrik
- Add kml in format that download-geofabrik can handle.
- handling all errors
- Add some help for users when a have a 404 error while downloading
- Support md5sum
- Add user/pass support for http proxy ([#8](https://github.com/julien-noblet/download-geofabrik/issues/8))
- Add proxy support should fix [#8](https://github.com/julien-noblet/download-geofabrik/issues/8)
- Add proxy support should fix [#8](https://github.com/julien-noblet/download-geofabrik/issues/8)
- Add user/pass support for http proxy ([#8](https://github.com/julien-noblet/download-geofabrik/issues/8))
- **Generate:** Use colly to improve speed and code review ([#165](https://github.com/julien-noblet/download-geofabrik/issues/165))
- **Generator:** Add progress bar for generator
- **Gislab:** Add Poly
- **Gislab:** Add gislab to pre-commit hook
- **Progressbar:** Add progressbar
- **Progressbar:** Add progressbar
- **openstreetmap.fr:** Add osm.pbf.md5 file support from openstreetmap.fr and add some new regions.
- **osmfr:** Use osm.pbf.md5
- **osmfr:** Add openstreetmap.fr.yml first version (not complete)
- **osmfr:** Add new service: openstreetmap.fr
- **osmfr:** Add openstreetmap.fr.yml firts version (not complete)

### Fix
- element2url when format have a basepath
- **Gislab:** basePath & baseURL for osm.pbf & osm.bz2

### Fix
- remove unused var
- [#30](https://github.com/julien-noblet/download-geofabrik/issues/30), Quick & dirty fix
- [#31](https://github.com/julien-noblet/download-geofabrik/issues/31)
- help command
- Better name for us -> United States of America
- update url

### Refactor
- split main()
- Upgrading generator
- flag var name url to Furl
- Use https instead of http
- flag var name url to Furl
- Better error handling
- Remove unused vars, Remove == true/false and prefer using string.Contains vs string.Index != -1
- Add verbose switch when generating openstreetmap.fr.yml & upgrading openstreetmap.fr.yml
- Changing verbose level
- Don't use pointer for anything
- if quiet flag is on, can't be verbose
- **Elements:** Use error and not print, in elements, better coverage
- **Elements:** optimization in elements, better coverage
- **Geofabrik:** Upgrading geofabrik
- **Progressbar:** Use a minimal file size for displaying progressbar
- **Progressbar:** Reduce cpu usage when using progressbar
- **Progressbar:** Fix progressbar display
- **Progressbar:** Use a minimal file size for displaying progressbar
- **Progressbar:** Reduce cpu usage when using progressbar
- **Progressbar:** Fix progressbar display
- **download:** Don't download file twice when checksum is required
- **download:** Speedup downloading checksum test
- **download:** Check hashsum by default
- **download:** Optimize addHash()
- **generator:** Handle bad responses codes from server
- **generator:** doc.Url == ctx.URL() but easyer to create a doc ^^
- **osmfr:** fix generator osmfr

### Style
- lint
- Lint
- gofmt
- linting
- lint
- linting
- **Lint:** removing unnecessary parentheses
- **Travis:** removing trailing spaces

### Test
- Better coverage for config.go
- Better coverage
- find a way to test check(err error)
- More tests for config, element and format
- Fix a test
- simplify tests
- Add comments and tests for parsers
- More tests for element
- More tests for download-geofabrik
- More tests for download-geofabrik
- More tests for download-geofabrik
- Better test
- Globalize element test data
- add benchmark for stringInSlice
- Better coverage for download-geofabrik.go
- add documents, better covering code, test for config
- **Progressbar:** Add one test to check progressbar
- **generator:** Add test and benchmarks for generator
- **generator:** More tests for generator
- **generator:** Add handle error and add some tests for generator
- **osmfr:** testing parser osmfr
- **osmfr:** More tests for parser osmfr
- **osmfr:** More tests and bench for parser osmfr

### Pull Requests
- Merge pull request [#25](https://github.com/julien-noblet/download-geofabrik/issues/25) from julien-noblet/coverage_error_handling
- Merge pull request [#23](https://github.com/julien-noblet/download-geofabrik/issues/23) from julien-noblet/checksum
- Merge pull request [#24](https://github.com/julien-noblet/download-geofabrik/issues/24) from julien-noblet/openstreetmap.fr
- Merge pull request [#22](https://github.com/julien-noblet/download-geofabrik/issues/22) from julien-noblet/hotfix-20
- Merge pull request [#18](https://github.com/julien-noblet/download-geofabrik/issues/18) from julien-noblet/fix-[#16](https://github.com/julien-noblet/download-geofabrik/issues/16)
- Merge pull request [#15](https://github.com/julien-noblet/download-geofabrik/issues/15) from julien-noblet/travis
- Merge pull request [#13](https://github.com/julien-noblet/download-geofabrik/issues/13) from julien-noblet/stable
- Merge pull request [#11](https://github.com/julien-noblet/download-geofabrik/issues/11) from julien-noblet/stable

### BREAKING CHANGE

Proxy settings have changed use environment

* ci(travis): removing go version 'tip' tests and build
  It add a lot of charge and time form travis. I'm not sure, we've to check if compile with
developpement go version.

* gislab: Remove gislab, they're no offering any free stuff

* (doc) Add some comments

Update yml files

ci(travis): build with go 1.7 is no longer supported


<a name="v2.0.1"></a>
## [v2.0.1] - 2018-12-08

<a name="v2.0.2"></a>
## [v2.0.2] - 2018-12-08

<a name="v2.0.3"></a>
## [v2.0.3] - 2018-12-08

<a name="v2.1.0-alpha"></a>
## [v2.1.0-alpha] - 2018-12-08

<a name="v2.1.0-pr"></a>
## [v2.1.0-pr] - 2018-12-08

<a name="v2.1.0"></a>
## [v2.1.0] - 2018-12-08

<a name="v2.2.0-rc0"></a>
## [v2.2.0-rc0] - 2018-12-08

<a name="v2.2.0"></a>
## [v2.2.0] - 2018-12-08

<a name="v2.2.1"></a>
## [v2.2.1] - 2018-12-08

<a name="v2.2.2"></a>
## [v2.2.2] - 2018-12-08

<a name="v2.2.3"></a>
## [v2.2.3] - 2018-12-08

<a name="v2.3.0-rc0"></a>
## [v2.3.0-rc0] - 2018-12-08

<a name="v2.3.0"></a>
## [v2.3.0] - 2018-12-08

<a name="v2.3.1"></a>
## [v2.3.1] - 2018-12-08
### Build
- More details in Makefile
- New Makefile using gox

### Chore
- Update geofabrik.yml
- Updating geofabrik.yml
- Updating geofabrik.yml
- Updating geofabrik.yml
- Removing yml from git repo, but keeping them into packages
- fix .gitignore
- Revert yml files must be in repo
- Adding lastest geofabrik.yml and openstreetmap.fr.yml
- **deps:** update golang.org/x/sys commit hash to 0cf1ed9 ([#49](https://github.com/julien-noblet/download-geofabrik/issues/49))
- **deps:** update golang.org/x/net commit hash to 10aee18 ([#39](https://github.com/julien-noblet/download-geofabrik/issues/39))
- **deps:** update golang.org/x/sys commit hash to 66b7b13 ([#40](https://github.com/julien-noblet/download-geofabrik/issues/40))
- **deps:** update module cheggaaa/pb to v1.0.26 ([#41](https://github.com/julien-noblet/download-geofabrik/issues/41))
- **deps:** update module olekukonko/tablewriter to v0.0.1 ([#42](https://github.com/julien-noblet/download-geofabrik/issues/42))
- **deps:** update golang.org/x/net commit hash to 03003ca ([#45](https://github.com/julien-noblet/download-geofabrik/issues/45))
- **deps:** update golang.org/x/net commit hash to adae6a3 ([#46](https://github.com/julien-noblet/download-geofabrik/issues/46))
- **deps:** update module puerkitobio/goquery to v1.5.0 ([#47](https://github.com/julien-noblet/download-geofabrik/issues/47))
- **deps:** update golang.org/x/sys commit hash to 93218de ([#48](https://github.com/julien-noblet/download-geofabrik/issues/48))
- **packages:** Use Go 1.11's modules

### Ci
- add goreportcard
- Deploy only on lastest go version
- Generate new geofabrik.yml on each releases
- generate ymls before generating readme
- Add coveralls to check code coverage
- Add version flag, See [#27](https://github.com/julien-noblet/download-geofabrik/issues/27)
- **Coveralls.io:** add coveralls.io's badge to Readme.md
- **Travis:** Deploy with go 1.11
- **Travis:** fix target branch to deploy .yml files
- **Travis:** fix All branches can deploy .yml files
- **Travis:** All branches can deploy .yml files
- **Travis:** Don't generate .yml files on build
- **Travis:** Update .travis.yml
- **Travis:** Use Jobs with a lot of conditional jobs
- **Travis:** Try to fix rules
- **Travis:** Don't generate .yml files for all commits!
- **Travis:** Fix somes rules [no-test,yaml]
- **Travis:** Try with travis stages
- **Travis:** Fix somes rules [no-test,yaml]
- **Travis:** Fix rules [no-test,yaml]
- **Travis:** All YAML can be deployed now!
- **Travis:** Don't generate .yml files for all commits!
- **Travis:** use stages
- **Travis:** Don't generate .yml files on build
- **Travis:** Auto upload files from travis-ci
- **Travis:** Try to update travis and run default make rule
- **Travis:** Update .travis.yml
- **Travis:** Update .travis.yml
- **Travis:** Update .travis.yml
- **Travis:** Update .travis.yml
- **Travis:** Update .travis.yml
- **Travis:** All branches can deploy .yml files
- **YAML:** Try to auto-update YAML files

### Coverage
- add some tests
- add some tests
- format.go -> 100%

### Docs
- Updating Readme
- Updating README.md
- Updating readme
- Generate README.md Add CHANGELOG.md in zip files
- Updating readme
- Updating Changelog using github-changelog-generator
- Updating Readme
- **Gislab:** Add gislab in help

### Feat
- Support md5sum
- handling all errors
- Add some help for users when a have a 404 error while downloading
- Add user/pass support for http proxy ([#8](https://github.com/julien-noblet/download-geofabrik/issues/8))
- Add proxy support should fix [#8](https://github.com/julien-noblet/download-geofabrik/issues/8)
- Add user/pass support for http proxy ([#8](https://github.com/julien-noblet/download-geofabrik/issues/8))
- Splitting gofiles ang adding generate to download in geofabrik
- Add kml in format that download-geofabrik can handle.
- Handle error from download, create file only after the download start without issue.
- Add osh.pbf file support
- Add osh.pbf support
- Add proxy support should fix [#8](https://github.com/julien-noblet/download-geofabrik/issues/8)
- **Generator:** Add progress bar for generator
- **Gislab:** Add Poly
- **Gislab:** Add gislab to pre-commit hook
- **List:** Add Markdown support for listing
- **Progressbar:** Add progressbar
- **Progressbar:** Add progressbar
- **openstreetmap.fr:** Add osm.pbf.md5 file support from openstreetmap.fr and add some new regions.
- **osmfr:** Add openstreetmap.fr.yml first version (not complete)
- **osmfr:** Add new service: openstreetmap.fr
- **osmfr:** Add openstreetmap.fr.yml firts version (not complete)

### Fix
- element2url when format have a basepath
- **Gislab:** basePath & baseURL for osm.pbf & osm.bz2

### Fix
- remove unused var
- [#30](https://github.com/julien-noblet/download-geofabrik/issues/30), Quick & dirty fix
- [#31](https://github.com/julien-noblet/download-geofabrik/issues/31)
- help command
- Better name for us -> United States of America
- update url

### Refactor
- split main()
- Add new shapes, useragent and safe crawler
- Upgrading generator
- flag var name url to Furl
- Use https instead of http
- flag var name url to Furl
- Better error handling
- Remove unused vars, Remove == true/false and prefer using string.Contains vs string.Index != -1
- Add verbose switch when generating openstreetmap.fr.yml & upgrading openstreetmap.fr.yml
- Changing verbose level
- Don't use pointer for anything
- if quiet flag is on, can't be verbose
- **Elements:** optimization in elements, better coverage
- **Elements:** Use error and not print, in elements, better coverage
- **Geofabrik:** Upgrading geofabrik
- **Progressbar:** Fix progressbar display
- **Progressbar:** Use a minimal file size for displaying progressbar
- **Progressbar:** Reduce cpu usage when using progressbar
- **Progressbar:** Fix progressbar display
- **Progressbar:** Use a minimal file size for displaying progressbar
- **Progressbar:** Reduce cpu usage when using progressbar
- **download:** Optimize addHash()
- **download:** Don't download file twice when checksum is required
- **download:** Speedup downloading checksum test
- **download:** Check hashsum by default
- **generator:** doc.Url == ctx.URL() but easyer to create a doc ^^
- **generator:** Handle bad responses codes from server
- **generator:** New generator
- **osmfr:** fix generator osmfr

### Style
- lint
- Lint
- gofmt
- linting
- lint
- linting
- Linting
- **Lint:** removing unnecessary parentheses
- **Travis:** removing trailing spaces

### Test
- Better coverage for config.go
- Better coverage
- find a way to test check(err error)
- More tests for config, element and format
- Fix a test
- simplify tests
- Add comments and tests for parsers
- More tests for element
- More tests for download-geofabrik
- More tests for download-geofabrik
- More tests for download-geofabrik
- Better test
- Globalize element test data
- add benchmark for stringInSlice
- Better coverage for download-geofabrik.go
- add documents, better covering code, test for config
- **Progressbar:** Add one test to check progressbar
- **generator:** Add test and benchmarks for generator
- **generator:** More tests for generator
- **generator:** Add handle error and add some tests for generator
- **osmfr:** testing parser osmfr
- **osmfr:** More tests for parser osmfr
- **osmfr:** More tests and bench for parser osmfr

### Pull Requests
- Merge pull request [#25](https://github.com/julien-noblet/download-geofabrik/issues/25) from julien-noblet/coverage_error_handling
- Merge pull request [#23](https://github.com/julien-noblet/download-geofabrik/issues/23) from julien-noblet/checksum
- Merge pull request [#24](https://github.com/julien-noblet/download-geofabrik/issues/24) from julien-noblet/openstreetmap.fr
- Merge pull request [#22](https://github.com/julien-noblet/download-geofabrik/issues/22) from julien-noblet/hotfix-20
- Merge pull request [#18](https://github.com/julien-noblet/download-geofabrik/issues/18) from julien-noblet/fix-[#16](https://github.com/julien-noblet/download-geofabrik/issues/16)
- Merge pull request [#15](https://github.com/julien-noblet/download-geofabrik/issues/15) from julien-noblet/travis
- Merge pull request [#13](https://github.com/julien-noblet/download-geofabrik/issues/13) from julien-noblet/stable
- Merge pull request [#11](https://github.com/julien-noblet/download-geofabrik/issues/11) from julien-noblet/stable
- Merge pull request [#7](https://github.com/julien-noblet/download-geofabrik/issues/7) from julien-noblet/generator-config


<a name="v2.0.0"></a>
## [v2.0.0] - 2018-12-08
### Build
- try without makefile and add arm
- add Go 1.6
- 1.4 was not working too :/
- Try cross compile without 1.3

### Ci
- **Travis:** Add Travis Badge
- **Travis:** First try travis

### Feat
- Support of multiple config
- Add more thing in config
- Add State file format support
- Add Kingpin, better command, V2 ready

### Style
- golint not working with ARM
- Fix golint?
- Fix golint? Add Badge for Travis
- Add Go 1.4, Add go lint
- clean

### Pull Requests
- Merge pull request [#5](https://github.com/julien-noblet/download-geofabrik/issues/5) from julien-noblet/v2
- Merge pull request [#6](https://github.com/julien-noblet/download-geofabrik/issues/6) from julien-noblet/More_Config
- Merge pull request [#4](https://github.com/julien-noblet/download-geofabrik/issues/4) from julien-noblet/travis


<a name="v0.0.2"></a>
## [v0.0.2] - 2018-12-08
### Ci
- Added Gitter badge

### Feat
- Add poly and state.txt files

### Pull Requests
- Merge pull request [#1](https://github.com/julien-noblet/download-geofabrik/issues/1) from gitter-badger/gitter-badge


<a name="v0.0.1"></a>
## [v0.0.1] - 2018-12-08

<a name="v2.0.4"></a>
## v2.0.4 - 2018-12-08
### Build
- More details in Makefile
- New Makefile using gox
- try without makefile and add arm
- add Go 1.6
- 1.4 was not working too :/
- Try cross compile without 1.3
- Makefile

### Chore
- Updating geofabrik.yml
- Updating geofabrik.yml
- Updating geofabrik.yml
- Update geofabrik.yml
- **Geofabrik:** add all missing
- **Geofabrik:** Add Canada
- **Geofabrik:** add all Europe
- **Geofabrik:** add Germany and Great-Britain
- **Geofabrik:** Add France
- **Geofabrik:** First commit

### Ci
- Generate new geofabrik.yml on each releases
- Deploy only on lastest go version
- Added Gitter badge
- add goreportcard
- Add new rules to gitignore
- **Travis:** Update .travis.yml
- **Travis:** Update .travis.yml
- **Travis:** Update .travis.yml
- **Travis:** Update .travis.yml
- **Travis:** Update .travis.yml
- **Travis:** Try to update travis and run default make rule
- **Travis:** Add Travis Badge
- **Travis:** First try travis
- **Travis:** Auto upload files from travis-ci
- **Travis:** Update .travis.yml

### Docs
- Generate README.md Add CHANGELOG.md in zip files
- Updating Changelog using github-changelog-generator
- Updating Readme
- Updating Readme

### Feat
- Add Kingpin, better command, V2 ready
- Add osh.pbf support
- Add osh.pbf file support
- Support of multiple config
- Add more thing in config
- Add State file format support
- Add poly and state.txt files
- **List:** Add Markdown support for listing

### Fix
- Better name for us -> United States of America
- update url
- georgia-eu and georgia-us

### Refactor
- Use https instead of http
- Upgrading generator
- Add new shapes, useragent and safe crawler
- **Geofabrik:** Upgrading geofabrik
- **generator:** New generator

### Style
- Lint
- Linting
- golint not working with ARM
- Fix golint?
- Fix golint? Add Badge for Travis
- Add Go 1.4, Add go lint
- clean

### Pull Requests
- Merge pull request [#15](https://github.com/julien-noblet/download-geofabrik/issues/15) from julien-noblet/travis
- Merge pull request [#13](https://github.com/julien-noblet/download-geofabrik/issues/13) from julien-noblet/stable
- Merge pull request [#11](https://github.com/julien-noblet/download-geofabrik/issues/11) from julien-noblet/stable
- Merge pull request [#7](https://github.com/julien-noblet/download-geofabrik/issues/7) from julien-noblet/generator-config
- Merge pull request [#5](https://github.com/julien-noblet/download-geofabrik/issues/5) from julien-noblet/v2
- Merge pull request [#6](https://github.com/julien-noblet/download-geofabrik/issues/6) from julien-noblet/More_Config
- Merge pull request [#4](https://github.com/julien-noblet/download-geofabrik/issues/4) from julien-noblet/travis
- Merge pull request [#1](https://github.com/julien-noblet/download-geofabrik/issues/1) from gitter-badger/gitter-badge


[Unreleased]: https://github.com/julien-noblet/download-geofabrik/compare/v2.6.0-rc1...HEAD
[v2.6.0-rc1]: https://github.com/julien-noblet/download-geofabrik/compare/v2.5.1-rc1...v2.6.0-rc1
[v2.5.1-rc1]: https://github.com/julien-noblet/download-geofabrik/compare/v2.5.1-rc0...v2.5.1-rc1
[v2.5.1-rc0]: https://github.com/julien-noblet/download-geofabrik/compare/v2.5.0...v2.5.1-rc0
[v2.5.0]: https://github.com/julien-noblet/download-geofabrik/compare/v2.4.1...v2.5.0
[v2.4.1]: https://github.com/julien-noblet/download-geofabrik/compare/v2.4.0...v2.4.1
[v2.4.0]: https://github.com/julien-noblet/download-geofabrik/compare/v2.0.1...v2.4.0
[v2.0.1]: https://github.com/julien-noblet/download-geofabrik/compare/v2.0.2...v2.0.1
[v2.0.2]: https://github.com/julien-noblet/download-geofabrik/compare/v2.0.3...v2.0.2
[v2.0.3]: https://github.com/julien-noblet/download-geofabrik/compare/v2.1.0-alpha...v2.0.3
[v2.1.0-alpha]: https://github.com/julien-noblet/download-geofabrik/compare/v2.1.0-pr...v2.1.0-alpha
[v2.1.0-pr]: https://github.com/julien-noblet/download-geofabrik/compare/v2.1.0...v2.1.0-pr
[v2.1.0]: https://github.com/julien-noblet/download-geofabrik/compare/v2.2.0-rc0...v2.1.0
[v2.2.0-rc0]: https://github.com/julien-noblet/download-geofabrik/compare/v2.2.0...v2.2.0-rc0
[v2.2.0]: https://github.com/julien-noblet/download-geofabrik/compare/v2.2.1...v2.2.0
[v2.2.1]: https://github.com/julien-noblet/download-geofabrik/compare/v2.2.2...v2.2.1
[v2.2.2]: https://github.com/julien-noblet/download-geofabrik/compare/v2.2.3...v2.2.2
[v2.2.3]: https://github.com/julien-noblet/download-geofabrik/compare/v2.3.0-rc0...v2.2.3
[v2.3.0-rc0]: https://github.com/julien-noblet/download-geofabrik/compare/v2.3.0...v2.3.0-rc0
[v2.3.0]: https://github.com/julien-noblet/download-geofabrik/compare/v2.3.1...v2.3.0
[v2.3.1]: https://github.com/julien-noblet/download-geofabrik/compare/v2.0.0...v2.3.1
[v2.0.0]: https://github.com/julien-noblet/download-geofabrik/compare/v0.0.2...v2.0.0
[v0.0.2]: https://github.com/julien-noblet/download-geofabrik/compare/v0.0.1...v0.0.2
[v0.0.1]: https://github.com/julien-noblet/download-geofabrik/compare/v2.0.4...v0.0.1
