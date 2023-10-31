## [1.1.1](https://github.com/Randsw/kubeinfo/compare/1.1.0...1.1.1) (2023-10-31)


### 🛠 Fixes

* Fetch all repo to build container with nesessary variables ([a80cf41](https://github.com/Randsw/kubeinfo/commit/a80cf415a78062990b71b61be73b8deee61de53c))

## [1.1.0](https://github.com/Randsw/kubeinfo/compare/1.0.4...1.1.0) (2023-10-31)


### 🦊 CI/CD

* Chart release only on push. ([d4f19be](https://github.com/Randsw/kubeinfo/commit/d4f19be4ab4830def1d49f8b5d80ad49d022e4ae))


### 🚀 Features

* Display tag, hash and date of commit in /healthz endpoint ([6541d9d](https://github.com/Randsw/kubeinfo/commit/6541d9de52887396a1edda708a63f3b5d404df4e))


### 🛠 Fixes

* Fix app name in status ([13fc6ee](https://github.com/Randsw/kubeinfo/commit/13fc6ee871172b1ea1fb299f8ce9ea7d8ff216fe))


### Other

* **deps:** bump actions/setup-node from 3 to 4 ([9ea8a79](https://github.com/Randsw/kubeinfo/commit/9ea8a797c3b3f4cb9ef5f9eec95409bbbcb2e72d))
* **deps:** bump randsw/kubeinfo in /helm-chart/kubeinfo-backend ([29552e2](https://github.com/Randsw/kubeinfo/commit/29552e24350a4b0c96fcdbe9cb0583995c4f1ece))
* **deps:** bump tj-actions/changed-files from 39 to 40 ([02323f2](https://github.com/Randsw/kubeinfo/commit/02323f2e5ea5ff3e764a491e5095f36f955ffb5d))

## [1.0.4](https://github.com/Randsw/kubeinfo/compare/1.0.3...1.0.4) (2023-10-24)


### 📔 Docs

* **helm:** Edit chart description ([2f4209a](https://github.com/Randsw/kubeinfo/commit/2f4209ae843c43e1d9ee6f6de8a6dc99998feb25))


### 🦊 CI/CD

* **helm:** Bump helm chart patch version if not manually bumped ([8d0515f](https://github.com/Randsw/kubeinfo/commit/8d0515fa701b17c4b0813bd9100cb4ff8e5c0110))
* **helm:** Bump version if user forget to do it manually ([424e603](https://github.com/Randsw/kubeinfo/commit/424e603ee0623ef9f073f2dc9e9d6cd847fe4f13))
* **helm:** Check manual helm chart version update ([3f5d477](https://github.com/Randsw/kubeinfo/commit/3f5d4774fa9ab6d549fa975203ed3f4880f9aa18))
* **helm:** Fix gh-action bump arguments ([16e45cc](https://github.com/Randsw/kubeinfo/commit/16e45ccaa41438648b96bbd11d76203fb50cf1c7))
* **helm:** Fix missing git configuration ([dec3e16](https://github.com/Randsw/kubeinfo/commit/dec3e16334da493327a6950c522936a8ab17142a))
* **helm:** Fix wildcard for yaml files in helm chart directory ([d2fc285](https://github.com/Randsw/kubeinfo/commit/d2fc2850fcb131ce49e0f8dfcdf3a73d00b06b81))
* **helm:** Refactor github actions ([9e808df](https://github.com/Randsw/kubeinfo/commit/9e808df1fba4d486e705873b7be479f807986cd6))
* **helm:** Remove unused commands in github-actions ([af21915](https://github.com/Randsw/kubeinfo/commit/af21915c3df41e8d41ad0fa96b6928e6d04cd9d2))
* **helm:** Test change only in values.yaml ([729495e](https://github.com/Randsw/kubeinfo/commit/729495efd070d3d4ca51bfa4d80d22a0ed245d76))
* **helm:** Use action to push. When push with GITHUB_TOKEN action is not triggered. Use PAT instead ([0c6ec14](https://github.com/Randsw/kubeinfo/commit/0c6ec1458049d2cc6a1f5aa9894b3632eafa90a8))
* **helm:** Use explicit step to bump helm chart version ([dded5c8](https://github.com/Randsw/kubeinfo/commit/dded5c8ae90712729eb050048a97acc70d8ed74a))
* **helm:** Use PAT then checkout. Fix version display in commit message ([ea5a6db](https://github.com/Randsw/kubeinfo/commit/ea5a6dbb8b1cbc5bbb2ce033d06b946079a1849c))


### 🛠 Fixes

* **logger:** Fix logger caller func - skip support logger function ([14c9d28](https://github.com/Randsw/kubeinfo/commit/14c9d283c06233e7e64d5118c8746f8fe034fcee))


### Other

* **deps:** bump github.com/fluxcd/helm-controller/api ([44e4203](https://github.com/Randsw/kubeinfo/commit/44e4203cead93290c0ec8e27e7e9ea3860012384))
* **deps:** bump github.com/fluxcd/helm-controller/api ([50c7275](https://github.com/Randsw/kubeinfo/commit/50c72751c3babf5c897a8d3e563cbd96261de207))
* **deps:** bump github.com/fluxcd/kustomize-controller/api ([3f42809](https://github.com/Randsw/kubeinfo/commit/3f428099c20f66c3c9308b9b0190738b1760bcdd))
* **deps:** bump github.com/prometheus/client_golang ([ca60dea](https://github.com/Randsw/kubeinfo/commit/ca60dea2d80460e182ae4441a89cc32fd2be1407))
* **deps:** bump go.uber.org/zap from 1.25.0 to 1.26.0 ([2004267](https://github.com/Randsw/kubeinfo/commit/20042674d48ad1202d88a5fb2bebed269c70b6ed))
* **deps:** bump k8s.io/api from 0.28.1 to 0.28.2 ([6865d09](https://github.com/Randsw/kubeinfo/commit/6865d09cc14422b89279ff2fef9bdaf02ffa704e))
* **deps:** bump k8s.io/api from 0.28.2 to 0.28.3 ([3f2cda2](https://github.com/Randsw/kubeinfo/commit/3f2cda2158a0d4a78e55111b34a39fead76b5598))
* **deps:** bump k8s.io/client-go from 0.28.1 to 0.28.2 ([8a63fc1](https://github.com/Randsw/kubeinfo/commit/8a63fc1faf7543d88b5eca6a5849283254613958))
* **deps:** bump k8s.io/client-go from 0.28.2 to 0.28.3 ([141643f](https://github.com/Randsw/kubeinfo/commit/141643fbb0b3e466536898aff44e2470435a3dbd))
* **deps:** bump randsw/kubeinfo in /helm-chart/kubeinfo-backend ([1b6ed98](https://github.com/Randsw/kubeinfo/commit/1b6ed9872a5af69174e29cf546a27b47099d0fc2))
* **deps:** Update gh-action go 1.to 21 ([e9b6da8](https://github.com/Randsw/kubeinfo/commit/e9b6da8e27d0910e95d0adce4bcbdc6714d0cfa5))

## [1.0.3](https://github.com/Randsw/kubeinfo/compare/1.0.2...1.0.3) (2023-09-14)


### 🦊 CI/CD

* Add dependabot ([128d568](https://github.com/Randsw/kubeinfo/commit/128d568a0d49ac09ecd69227f055c3aefbbb4dbd))


### 🛠 Fixes

* Test dependabot update values.yaml in helm chart from chcr.io ([1218f7f](https://github.com/Randsw/kubeinfo/commit/1218f7f1c9b6f1435dd2725111146c072b4b2ef2))


### Other

* **deps:** bump actions/checkout from 3 to 4 ([ea2eb98](https://github.com/Randsw/kubeinfo/commit/ea2eb98cfff7d1600ba260e1fbbaddfc5e57e8fe))
* **deps:** bump docker/build-push-action from 4 to 5 ([b97ac01](https://github.com/Randsw/kubeinfo/commit/b97ac01aba999cc5e698e1e6d1913d5c8e52c4d5))
* **deps:** bump docker/login-action from 2 to 3 ([46a2c1c](https://github.com/Randsw/kubeinfo/commit/46a2c1c3d34e208cabb9f1218c4f92a634fe8d73))
* **deps:** bump docker/metadata-action from 4 to 5 ([2784a77](https://github.com/Randsw/kubeinfo/commit/2784a772b343ab183dcb3c7dc003c1f91fd66bf6))
* **deps:** bump docker/setup-buildx-action from 2 to 3 ([836a24f](https://github.com/Randsw/kubeinfo/commit/836a24f09aaac3428a63904a2b0adc2d14c94e85))
* **deps:** bump github.com/fluxcd/helm-controller/api ([fcbf932](https://github.com/Randsw/kubeinfo/commit/fcbf932fd1e461d613cef340f3bfbf55a5ed324d))
* **deps:** bump github.com/fluxcd/kustomize-controller/api ([c45a4f7](https://github.com/Randsw/kubeinfo/commit/c45a4f77e9911da8fc3d7efc8758353c82997f35))
* **deps:** bump github.com/prometheus/client_golang ([3a7a923](https://github.com/Randsw/kubeinfo/commit/3a7a9233c03426ffbaa977e17ed466fd8fd54bf5))
* **deps:** bump go.uber.org/zap from 1.24.0 to 1.25.0 ([3e6714c](https://github.com/Randsw/kubeinfo/commit/3e6714c4cf958e82930e12fc924235e72f9bf100))
* **deps:** bump golang from 1.19 to 1.21 ([99ba164](https://github.com/Randsw/kubeinfo/commit/99ba16475e245926119cff43cd8feb1f29f73533))
* **deps:** bump helm/chart-testing-action from 2.3.1 to 2.4.0 ([5cb2aac](https://github.com/Randsw/kubeinfo/commit/5cb2aac5df0315fdd70666bbb88bd883ef760ea1))
* **deps:** bump helm/kind-action from 1.5.0 to 1.8.0 ([67305ab](https://github.com/Randsw/kubeinfo/commit/67305aba13d4769ca82c1c7620ff15677daa60cb))
* **deps:** bump k8s.io/client-go from 0.26.3 to 0.28.1 ([6808ab2](https://github.com/Randsw/kubeinfo/commit/6808ab2f0a420c047234d03129626ae2dba06ad3))
* **deps:** bump tj-actions/changed-files from 35 to 39 ([3f9cd6c](https://github.com/Randsw/kubeinfo/commit/3f9cd6cc1156167be0c8d6c5ac5ab162a0de2db4))

## [1.0.2](https://github.com/Randsw/kubeinfo/compare/1.0.1...1.0.2) (2023-05-18)


### 🛠 Fixes

* **struct:** Add comment to NamespaceResponce struct ([4e51aec](https://github.com/Randsw/kubeinfo/commit/4e51aec1b4e281fac3ffadeb35b4882937e6eb12))


### Other

* **release:** 1.0.2-develop.1 ([2916f6b](https://github.com/Randsw/kubeinfo/commit/2916f6b28f50de166beba50a837eebcdc1f3d20f))

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
