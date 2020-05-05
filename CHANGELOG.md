<a name="unreleased"></a>
## [Unreleased]


<a name="v0.5.0"></a>
## [v0.5.0] - 2020-05-05
### Bug Fixes
- property field alignment ([#72](https://github.com/victorkt/flaggio/issues/72))
- constraint type when editing was always string ([#61](https://github.com/victorkt/flaggio/issues/61))
- explainer labels on segment rules ([#60](https://github.com/victorkt/flaggio/issues/60))

### Dependencies
- bump uuid from 7.0.3 to 8.0.0 in /web ([#78](https://github.com/victorkt/flaggio/issues/78))
- bump github.com/uber/jaeger-client-go from 2.23.0+incompatible to 2.23.1+incompatible ([#76](https://github.com/victorkt/flaggio/issues/76))
- bump github.com/sirupsen/logrus from 1.5.0 to 1.6.0 ([#75](https://github.com/victorkt/flaggio/issues/75))
- bump @material-ui/core from 4.9.11 to 4.9.12 in /web ([#66](https://github.com/victorkt/flaggio/issues/66))
- bump github.com/uber/jaeger-client-go from 2.22.1+incompatible to 2.23.0+incompatible ([#65](https://github.com/victorkt/flaggio/issues/65))
- bump node-sass from 4.13.1 to 4.14.0 in /web ([#67](https://github.com/victorkt/flaggio/issues/67))
- bump @material-ui/core from 4.9.10 to 4.9.11 in /web ([#58](https://github.com/victorkt/flaggio/issues/58))
- bump @apollo/react-hooks from 3.1.4 to 3.1.5 in /web ([#57](https://github.com/victorkt/flaggio/issues/57))
- bump github.com/go-chi/chi from 4.1.0+incompatible to 4.1.1+incompatible ([#59](https://github.com/victorkt/flaggio/issues/59))
- bump prettier from 2.0.4 to 2.0.5 in /web ([#68](https://github.com/victorkt/flaggio/issues/68))

### Features
- add new users page ([#74](https://github.com/victorkt/flaggio/issues/74))
- store evaluations and users ([#54](https://github.com/victorkt/flaggio/issues/54))
- keep search results when coming back from the (new/edit) flag page ([#73](https://github.com/victorkt/flaggio/issues/73))


<a name="v0.4.1"></a>
## [v0.4.1] - 2020-04-16
### Bug Fixes
- return specific variant in rule ([#55](https://github.com/victorkt/flaggio/issues/55))

### Features
- add constraint explainer labels (if/and) ([#56](https://github.com/victorkt/flaggio/issues/56))


<a name="v0.4.0"></a>
## [v0.4.0] - 2020-04-13
### Bug Fixes
- nils from context should be treated as invalid ([#33](https://github.com/victorkt/flaggio/issues/33))

### Dependencies
- fix package-lock.json ([#53](https://github.com/victorkt/flaggio/issues/53))
- bump [@material](https://github.com/material)-ui/styles from 4.9.6 to 4.9.10 in /web ([#51](https://github.com/victorkt/flaggio/issues/51))
- bump [@material](https://github.com/material)-ui/core from 4.9.9 to 4.9.10 in /web ([#49](https://github.com/victorkt/flaggio/issues/49))
- bump go.mongodb.org/mongo-driver from 1.3.1 to 1.3.2 ([#52](https://github.com/victorkt/flaggio/issues/52))
- bump github.com/go-chi/chi from 4.0.4+incompatible to 4.1.0+incompatible ([#35](https://github.com/victorkt/flaggio/issues/35))

### Documentation
- add CHANGELOG

### Features
- new variant is consistent with previous ones ([#46](https://github.com/victorkt/flaggio/issues/46))
- implement distribution rollout on frontend ([#43](https://github.com/victorkt/flaggio/issues/43))
- case insensitive sort for find all flags ([#44](https://github.com/victorkt/flaggio/issues/44))
- add 'rows per page' = All ([#42](https://github.com/victorkt/flaggio/issues/42))
- add tracing support ([#41](https://github.com/victorkt/flaggio/issues/41))
- caching is optional ([#40](https://github.com/victorkt/flaggio/issues/40))
- add redis cache support ([#34](https://github.com/victorkt/flaggio/issues/34))

### Performance Improvements
- limit amount of variants/rules/constraints ([#45](https://github.com/victorkt/flaggio/issues/45))


<a name="v0.3.0"></a>
## [v0.3.0] - 2020-04-01
### Bug Fixes
- evaluateAll doesn't return nulls anymore ([#31](https://github.com/victorkt/flaggio/issues/31))
- remove default variant refs in newFlag func ([#30](https://github.com/victorkt/flaggio/issues/30))

### Features
- add pagination to flags table ([#29](https://github.com/victorkt/flaggio/issues/29))
- add flag search functionality ([#28](https://github.com/victorkt/flaggio/issues/28))
- default variants for new flags ([#25](https://github.com/victorkt/flaggio/issues/25))


<a name="v0.2.0"></a>
## v0.2.0 - 2020-03-28
### Bug Fixes
- gqlgen new server api ([#20](https://github.com/victorkt/flaggio/issues/20))

### Continuous Integration
- only release on tags
- build & release ([#24](https://github.com/victorkt/flaggio/issues/24))

### Dependencies
- bump clientip ([#23](https://github.com/victorkt/flaggio/issues/23))

### Work in Progress
- new cmd


[Unreleased]: https://github.com/victorkt/flaggio/compare/v0.5.0...HEAD
[v0.5.0]: https://github.com/victorkt/flaggio/compare/v0.4.1...v0.5.0
[v0.4.1]: https://github.com/victorkt/flaggio/compare/v0.4.0...v0.4.1
[v0.4.0]: https://github.com/victorkt/flaggio/compare/v0.3.0...v0.4.0
[v0.3.0]: https://github.com/victorkt/flaggio/compare/v0.2.0...v0.3.0
