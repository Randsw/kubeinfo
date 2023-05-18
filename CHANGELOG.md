## [1.0.2-develop.1](https://github.com/Randsw/kubeinfo/compare/1.0.1...1.0.2-develop.1) (2023-05-18)


### 🛠 Fixes

* **struct:** Add comment to NamespaceResponce struct ([4e51aec](https://github.com/Randsw/kubeinfo/commit/4e51aec1b4e281fac3ffadeb35b4882937e6eb12))

## [1.0.1](https://github.com/Randsw/kubeinfo/compare/1.0.0...1.0.1) (2023-05-18)


### 🛠 Fixes

* **struct:** Add comment to NodeResponse struct ([be83bb1](https://github.com/Randsw/kubeinfo/commit/be83bb1c3e47cddebbbf1623f2aec20a48816566))


### Other

* **release:** 0.3.1-develop.5 ([e20803d](https://github.com/Randsw/kubeinfo/commit/e20803de238525c20aad0dccc5a1da78d8db35ee))

## [1.0.0](https://github.com/Randsw/kubeinfo/compare/0.2.2...1.0.0) (2023-05-17)


### 🦊 CI/CD

* **semver:** Delete option in checkout ([03da109](https://github.com/Randsw/kubeinfo/commit/03da109efa54bd83b5377afa3ede0e8526c22d3b))
* **semver:** Disable npm publish ([0b34961](https://github.com/Randsw/kubeinfo/commit/0b349613a3469c7b69be9f4771baeffeeced63f7))
* **semver:** Fix typo in semver config ([d301398](https://github.com/Randsw/kubeinfo/commit/d3013987693489ecb549476b25a1bedcc74d08c6))
* **semver:** Install npm plugin ([669b94b](https://github.com/Randsw/kubeinfo/commit/669b94b01357b60474c533035ceced94cc5b4d38))
* **semver:** Remove semantic release 3rd party action, explicit install and use it ([cd30c91](https://github.com/Randsw/kubeinfo/commit/cd30c9115cb5812315c97d89beadf59312d2a6d4))
* **semver:** Trigger release ([b64e949](https://github.com/Randsw/kubeinfo/commit/b64e949d2e4164a88b52d12cfbfaf9acabd09a7c))
* **semver:** Use another var ([b623c0a](https://github.com/Randsw/kubeinfo/commit/b623c0ae37ae9599cc436fad2e6ed185ba4e485c))


### 🚀 Features

* **ci:** Add automatic semantic release ([cfad72e](https://github.com/Randsw/kubeinfo/commit/cfad72e1149cef8758f4e1504af8be547db8c538))
* **gh-actions:** Run tests only after lint stage ([7fc218f](https://github.com/Randsw/kubeinfo/commit/7fc218ff5f291b633fd55fbcb5d205abdebd79ae))
* **helm:** Add prometheus rule. Bump chart version ([#23](https://github.com/Randsw/kubeinfo/issues/23)) ([0b8b92f](https://github.com/Randsw/kubeinfo/commit/0b8b92f71d7864163175a4ae32057a2177310290))


### 🛠 Fixes

* **app-structs:** Add struct definition ([6f97348](https://github.com/Randsw/kubeinfo/commit/6f97348ce3365dea73b81a0f9814d257ecc05a96)), closes [#37](https://github.com/Randsw/kubeinfo/issues/37)
* **app-structs:** Add struct definition ([1e2f109](https://github.com/Randsw/kubeinfo/commit/1e2f109b3140a4f2032aaf345229583e7cec70b2))
* **app:** Add comment ([e5ce2c1](https://github.com/Randsw/kubeinfo/commit/e5ce2c189c25e75a1a692a82865d3effcb16359d))
* **app:** Change gh action workflow structure. Release only after lint and test on develop branch ([7cf9d91](https://github.com/Randsw/kubeinfo/commit/7cf9d9154d9fb17ff61b8a6cb891371a307a457e))
* **app:** More comments ([c0209b3](https://github.com/Randsw/kubeinfo/commit/c0209b3897911e6f0337afdea2e3776c438bcc90))
* **app:** Test image build after tag ([9e63578](https://github.com/Randsw/kubeinfo/commit/9e63578089eacce06a58cb711450ec2dcd37bea6))
* **app:** Test version bump patch ([8f85679](https://github.com/Randsw/kubeinfo/commit/8f85679b0734a410a0644592e1f2ec1fb3d6e25a))
* **app:** Use new action for checking which branch belong commit with tag ([d7e0c84](https://github.com/Randsw/kubeinfo/commit/d7e0c84b99ddbe4375f90bdaae949f7419ec793d))
* **ci:** Downgrade Node verison ([bf4c748](https://github.com/Randsw/kubeinfo/commit/bf4c74809af81308611fa039f71440aa8e674276)), closes [#30](https://github.com/Randsw/kubeinfo/issues/30)
* **ci:** Downgrade Node verison ([d8974e8](https://github.com/Randsw/kubeinfo/commit/d8974e8b2115207a6c01a46558e38ea2fbda51b6))
* **ci:** Fix semantic release extra plugins ([f096060](https://github.com/Randsw/kubeinfo/commit/f0960601d8ce886822ff3510030f46461384d2ee))
* **ci:** Fix semantic release extra plugins ([2e81b8a](https://github.com/Randsw/kubeinfo/commit/2e81b8a4c327884534954c78fe2428b99429393d))
* **ci:** Start semantic release only if tests pass ([f14c055](https://github.com/Randsw/kubeinfo/commit/f14c0559223423a1ddce8c40733cf480935dd00d))
* **ci:** Update go file ([e1543e0](https://github.com/Randsw/kubeinfo/commit/e1543e0aac639a1b144dd8ad52205056190a8b96))
* **gh-action:** Release can be done only on push to main ([44d6a9c](https://github.com/Randsw/kubeinfo/commit/44d6a9cb253f4e6074f7c14448e926cf36dab0c4))
* **gh-actions:** Fix error in release if statement ([1137c09](https://github.com/Randsw/kubeinfo/commit/1137c09dd49f11ec6e2f1e4db011c7aaa0008bd0))
* **gh-actions:** Run release only on tag belong to main branch commit ([06c07d7](https://github.com/Randsw/kubeinfo/commit/06c07d75b5b37f2c7936f64c10a3d6ef3be120c1))
* **semver:** Use node version 14 ([4c2f50f](https://github.com/Randsw/kubeinfo/commit/4c2f50f5c44f4fa52093e76682cedbe54fdfa5d6))
* **semver:** Use node version 14 ([d76c450](https://github.com/Randsw/kubeinfo/commit/d76c450308cad0fe23adc5ba6c3cb27718cf0dbd))
* **semver:** Use node version 16 ([b3a6e1d](https://github.com/Randsw/kubeinfo/commit/b3a6e1dfcd1c60ccade610aca8613eafe9d5e40c))


### Other

* **release:** 0.3.0-develop.1 [skip ci] ([5b08b0f](https://github.com/Randsw/kubeinfo/commit/5b08b0f7a028450a3211814e9ff624423488b8e7))
* **release:** 0.3.1-develop.1 [skip ci] ([d5cc411](https://github.com/Randsw/kubeinfo/commit/d5cc4110de190009bf33e5e7d7c07c3284cc5dca))
* **release:** 0.3.1-develop.2 ([f59460f](https://github.com/Randsw/kubeinfo/commit/f59460fb449bc3031303f8f6ffae714c60fccc8a))
* **release:** 0.3.1-develop.3 ([bdb43c3](https://github.com/Randsw/kubeinfo/commit/bdb43c32386ad08766c31a6821c7e18c55875ad9))
* **release:** 0.3.1-develop.4 ([77a5f7d](https://github.com/Randsw/kubeinfo/commit/77a5f7d7901bb360a30609fb9ba4693354954980))
* **semver:** Bump semantic release verison ([55973e3](https://github.com/Randsw/kubeinfo/commit/55973e39e68fd064806a65c8b7baf9fb33da6614))


## [0.3.1-develop.4](https://github.com/Randsw/kubeinfo/compare/0.3.1-develop.3...0.3.1-develop.4) (2023-05-17)


### 🛠 Fixes

* **app-structs:** Add struct definition ([1e2f109](https://github.com/Randsw/kubeinfo/commit/1e2f109b3140a4f2032aaf345229583e7cec70b2))

## [0.3.1-develop.3](https://github.com/Randsw/kubeinfo/compare/0.3.1-develop.2...0.3.1-develop.3) (2023-05-17)


### 🛠 Fixes

* **app:** Use new action for checking which branch belong commit with tag ([d7e0c84](https://github.com/Randsw/kubeinfo/commit/d7e0c84b99ddbe4375f90bdaae949f7419ec793d))

## [0.3.1-develop.2](https://github.com/Randsw/kubeinfo/compare/0.3.1-develop.1...0.3.1-develop.2) (2023-05-17)


### 🛠 Fixes

* **app:** Test image build after tag ([9e63578](https://github.com/Randsw/kubeinfo/commit/9e63578089eacce06a58cb711450ec2dcd37bea6))

## [0.3.1-develop.1](https://github.com/Randsw/kubeinfo/compare/0.3.0-develop.1...0.3.1-develop.1) (2023-05-17)


### 🦊 CI/CD

* **semver:** Disable npm publish ([0b34961](https://github.com/Randsw/kubeinfo/commit/0b349613a3469c7b69be9f4771baeffeeced63f7))
* **semver:** Fix typo in semver config ([d301398](https://github.com/Randsw/kubeinfo/commit/d3013987693489ecb549476b25a1bedcc74d08c6))
* **semver:** Install npm plugin ([669b94b](https://github.com/Randsw/kubeinfo/commit/669b94b01357b60474c533035ceced94cc5b4d38))


### 🛠 Fixes

* **app:** Change gh action workflow structure. Release only after lint and test on develop branch ([7cf9d91](https://github.com/Randsw/kubeinfo/commit/7cf9d9154d9fb17ff61b8a6cb891371a307a457e))
* **app:** Test version bump patch ([8f85679](https://github.com/Randsw/kubeinfo/commit/8f85679b0734a410a0644592e1f2ec1fb3d6e25a))
* **ci:** Start semantic release only if tests pass ([f14c055](https://github.com/Randsw/kubeinfo/commit/f14c0559223423a1ddce8c40733cf480935dd00d))

## [0.3.0-develop.1](https://github.com/Randsw/kubeinfo/compare/0.2.0...0.3.0-develop.1) (2023-05-17)


### :scissors: Refactor

* **helm:** Remove useless var ([d17cde0](https://github.com/Randsw/kubeinfo/commit/d17cde0fa78636f91f343639a7953e3aa73ea250))


### 🦊 CI/CD

* **semver:** Delete option in checkout ([03da109](https://github.com/Randsw/kubeinfo/commit/03da109efa54bd83b5377afa3ede0e8526c22d3b))
* **semver:** Remove semantic release 3rd party action, explicit install and use it ([cd30c91](https://github.com/Randsw/kubeinfo/commit/cd30c9115cb5812315c97d89beadf59312d2a6d4))
* **semver:** Trigger release ([b64e949](https://github.com/Randsw/kubeinfo/commit/b64e949d2e4164a88b52d12cfbfaf9acabd09a7c))
* **semver:** Use another var ([b623c0a](https://github.com/Randsw/kubeinfo/commit/b623c0ae37ae9599cc436fad2e6ed185ba4e485c))


### 🚀 Features

* **app:** Allow set serving address by env var ([2ab123f](https://github.com/Randsw/kubeinfo/commit/2ab123f28e66bcf97a68f907a392202b98eb134c))
* **app:** Get server port value from env ([c6d23fb](https://github.com/Randsw/kubeinfo/commit/c6d23fb60015f744cc5c4d6f8d567126054cc5f9))
* **ci:** Add automatic semantic release ([cfad72e](https://github.com/Randsw/kubeinfo/commit/cfad72e1149cef8758f4e1504af8be547db8c538))
* **ge-actions:** Allow trigger workflow on pull-request to develop ([2ce1a76](https://github.com/Randsw/kubeinfo/commit/2ce1a7615848f95c2852616ed45b5a5a5ada7b26))
* **gh-actions:** Run tests only after lint stage ([7fc218f](https://github.com/Randsw/kubeinfo/commit/7fc218ff5f291b633fd55fbcb5d205abdebd79ae))
* **helm:** Add prometheus rule. Bump chart version ([#23](https://github.com/Randsw/kubeinfo/issues/23)) ([0b8b92f](https://github.com/Randsw/kubeinfo/commit/0b8b92f71d7864163175a4ae32057a2177310290))


### 🛠 Fixes

* **app:** Add comment ([e5ce2c1](https://github.com/Randsw/kubeinfo/commit/e5ce2c189c25e75a1a692a82865d3effcb16359d))
* **app:** Fix serving port default value ([#13](https://github.com/Randsw/kubeinfo/issues/13)) ([1a613d9](https://github.com/Randsw/kubeinfo/commit/1a613d99bf0071a467930db45d53afc794812780))
* **app:** More comments ([c0209b3](https://github.com/Randsw/kubeinfo/commit/c0209b3897911e6f0337afdea2e3776c438bcc90))
* **ci:** Downgrade Node verison ([bf4c748](https://github.com/Randsw/kubeinfo/commit/bf4c74809af81308611fa039f71440aa8e674276)), closes [#30](https://github.com/Randsw/kubeinfo/issues/30)
* **ci:** Downgrade Node verison ([d8974e8](https://github.com/Randsw/kubeinfo/commit/d8974e8b2115207a6c01a46558e38ea2fbda51b6))
* **ci:** Fix semantic release extra plugins ([f096060](https://github.com/Randsw/kubeinfo/commit/f0960601d8ce886822ff3510030f46461384d2ee))
* **ci:** Fix semantic release extra plugins ([2e81b8a](https://github.com/Randsw/kubeinfo/commit/2e81b8a4c327884534954c78fe2428b99429393d))
* **ci:** Update go file ([e1543e0](https://github.com/Randsw/kubeinfo/commit/e1543e0aac639a1b144dd8ad52205056190a8b96))
* **gh-action:** Release can be done only on push to main ([44d6a9c](https://github.com/Randsw/kubeinfo/commit/44d6a9cb253f4e6074f7c14448e926cf36dab0c4))
* **gh-action:** Remove old workflow ([86487f7](https://github.com/Randsw/kubeinfo/commit/86487f7e73d821785764ebafde8c2b82237a0348))
* **gh-action:** Remove old workflow ([95ea668](https://github.com/Randsw/kubeinfo/commit/95ea6685d9568db896f8ba8436864e9e0b6817ac))
* **gh-actions:** Change gh action trigger ([125a4e1](https://github.com/Randsw/kubeinfo/commit/125a4e13b387f7db0672ca4215fa965f40674ad3))
* **gh-actions:** Change path filter in gh actions ([79e6579](https://github.com/Randsw/kubeinfo/commit/79e6579d77b2af3ead91943731e29f2f20032a41))
* **gh-actions:** Change trigger to run on push on develop and  main. Aslo trigger test after pull request to develop ([1d2bad4](https://github.com/Randsw/kubeinfo/commit/1d2bad452c3753140c5cbea52d2408a8bea30fae))
* **gh-actions:** Dont run test on main branch ([d331a7d](https://github.com/Randsw/kubeinfo/commit/d331a7d12ef2981061ea82ad6e6e0db14da5b766))
* **gh-actions:** Fix - Trigger dont run test workflow ([f3a26e9](https://github.com/Randsw/kubeinfo/commit/f3a26e9e96a202f7da5fdaf89cd66684cec789ce))
* **gh-actions:** Fix error in release if statement ([1137c09](https://github.com/Randsw/kubeinfo/commit/1137c09dd49f11ec6e2f1e4db011c7aaa0008bd0))
* **gh-actions:** Run release only on tag belong to main branch commit ([06c07d7](https://github.com/Randsw/kubeinfo/commit/06c07d75b5b37f2c7936f64c10a3d6ef3be120c1))
* **gh-actions:** Run test only if *.go file is changed ([68a8a6b](https://github.com/Randsw/kubeinfo/commit/68a8a6b837c9edb178f65f108f69445ab0a1f5b9))
* **gh-actions:** Trigger image build on tag in main ([f5a0e13](https://github.com/Randsw/kubeinfo/commit/f5a0e1331731f118b3b04396e89abd9a710b8ba1))
* **gh-actions:** Trigger image build on tag in main ([3f584d4](https://github.com/Randsw/kubeinfo/commit/3f584d45d926cbe14af175c1ad8682dac43e9554))
* **helm:** Disable prometheus metrics by default ([8267766](https://github.com/Randsw/kubeinfo/commit/8267766bbf9f741fb272a7bdf47f053d1ee3eebd))
* **helm:** Disable prometheus metrics by default ([4c83aa0](https://github.com/Randsw/kubeinfo/commit/4c83aa0c61571357b8f23d052b2d552fda4c59e9))
* **semver:** Use node version 14 ([4c2f50f](https://github.com/Randsw/kubeinfo/commit/4c2f50f5c44f4fa52093e76682cedbe54fdfa5d6))
* **semver:** Use node version 14 ([d76c450](https://github.com/Randsw/kubeinfo/commit/d76c450308cad0fe23adc5ba6c3cb27718cf0dbd))
* **semver:** Use node version 16 ([b3a6e1d](https://github.com/Randsw/kubeinfo/commit/b3a6e1dfcd1c60ccade610aca8613eafe9d5e40c))


### Other

* **app:** Bump app version for gh-actions test ([9a7acc4](https://github.com/Randsw/kubeinfo/commit/9a7acc4d03ed8143e4d85ca808286197b5f527ee))
* **helm:** Bump chart version ([14aa3f3](https://github.com/Randsw/kubeinfo/commit/14aa3f3b64124c378413052fcf1e22a26b35c0f0))
* **semver:** Bump semantic release verison ([55973e3](https://github.com/Randsw/kubeinfo/commit/55973e39e68fd064806a65c8b7baf9fb33da6614))
