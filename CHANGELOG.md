## [1.2.1](https://github.com/Randsw/kubeinfo/compare/1.2.0...1.2.1) (2026-09-11)


### 🦊 CI/CD

* **docker:** cross-compile and harden the image build ([bb1fbcd](https://github.com/Randsw/kubeinfo/commit/bb1fbcd49ee170be1e5a2d0ebd4b2790c28f26cd))
* **docker:** drop .git from the build context ([53603f6](https://github.com/Randsw/kubeinfo/commit/53603f6e4c1e77023d17053da6c315a2a875be84))
* **docker:** pass version metadata as build args ([ab2c413](https://github.com/Randsw/kubeinfo/commit/ab2c41330dbc7e0023a3d0a19e7d3d718e8c28df))
* First checkout code, then setup go ([8c0e6c6](https://github.com/Randsw/kubeinfo/commit/8c0e6c639280701f84ddfc7290038a1e597a699b))
* Fix linter errors ([43f9422](https://github.com/Randsw/kubeinfo/commit/43f9422ddfb9e574b8a6a2460d6ef94403adc2a6))
* No cache in setup go job ([15f284d](https://github.com/Randsw/kubeinfo/commit/15f284d9be0b808ac2d2f761f8d4a898e6bdc711))
* update dependency ([e50bcb4](https://github.com/Randsw/kubeinfo/commit/e50bcb41058976229832c4abf504bfc9792150ec))
* Update helm chart release trigger condition ([4c71023](https://github.com/Randsw/kubeinfo/commit/4c7102388992237c57ae2746d75a1f7e8d871b25))
* **helm:** Fix file change detection ([6bbfbcf](https://github.com/Randsw/kubeinfo/commit/6bbfbcfb5026cf873d3b7427610573850e9e25dc))


### 🛠 Fixes

* **ci:** restore original runner comment wording ([98770ca](https://github.com/Randsw/kubeinfo/commit/98770ca184c523f0139de25160f3ec0a98399503))


### Other

* Bump golang to 1.25.1 ([b239e25](https://github.com/Randsw/kubeinfo/commit/b239e2538d0b259be878a74507a70871e9abf70d))
* Bump golang version to 1.26 ([5e6e5fc](https://github.com/Randsw/kubeinfo/commit/5e6e5fc046679e28f9962a04c2e941f9b89ccebb))
* Replace deprecated NewSimpleClientSet by NewClientSet ([911d6d7](https://github.com/Randsw/kubeinfo/commit/911d6d78a4ab09e03126e075ebca64e1b4286918))
* Update client-go ([f50429c](https://github.com/Randsw/kubeinfo/commit/f50429c08918c13016196873b5908e9acfc1dc38))
* **deps:** bump actions/checkout from 4 to 5 ([660d127](https://github.com/Randsw/kubeinfo/commit/660d12703fa3404f604d31a408c65d89918ee771))
* **deps:** bump actions/checkout from 5 to 6 ([61c4448](https://github.com/Randsw/kubeinfo/commit/61c4448070337dfa0baf1ef77b0d85d16699a1a2))
* **deps:** bump actions/checkout from 6 to 7 ([e4a89b5](https://github.com/Randsw/kubeinfo/commit/e4a89b50c8d7e5120c2c49dfd4e368147a28cb40))
* **deps:** bump actions/setup-go from 4 to 5 ([d8a6bb1](https://github.com/Randsw/kubeinfo/commit/d8a6bb1f2ecc4acf481f9671cd056b94b790fef2))
* **deps:** bump actions/setup-go from 5 to 6 ([85f70fc](https://github.com/Randsw/kubeinfo/commit/85f70fc158fd8faf3b5d5402507baafaffda8410))
* **deps:** bump actions/setup-go from 6 to 7 ([8a5b50e](https://github.com/Randsw/kubeinfo/commit/8a5b50e17d823bda865b91850cce9f9abd2ff711))
* **deps:** bump actions/setup-node from 4 to 5 ([18c6833](https://github.com/Randsw/kubeinfo/commit/18c68338a51bfacaabb783bdcf66ac8a9e338035))
* **deps:** bump actions/setup-node from 5 to 6 ([f7afc51](https://github.com/Randsw/kubeinfo/commit/f7afc516fd182d75dccdeb910723442c6741c5ba))
* **deps:** bump actions/setup-node from 6 to 7 ([773c92f](https://github.com/Randsw/kubeinfo/commit/773c92f1ac6c9985cc7075fdd7945bb89052338d))
* **deps:** bump actions/setup-python from 4 to 5 ([f67148b](https://github.com/Randsw/kubeinfo/commit/f67148b46dfde40e1bd33323c0e5fa0ca434eb40))
* **deps:** bump actions/setup-python from 5 to 6 ([f0ce50d](https://github.com/Randsw/kubeinfo/commit/f0ce50d6ccd7fc0efa00bad9faa0624dfb78bc2c))
* **deps:** bump actions/setup-python from 6 to 7 ([697ef94](https://github.com/Randsw/kubeinfo/commit/697ef948065e6b22b30511f427271f5fad9b6176))
* **deps:** bump azure/setup-helm from 3 to 4 ([33fdf17](https://github.com/Randsw/kubeinfo/commit/33fdf1713e939d0d7e516dfa509eb80734cc9634))
* **deps:** bump azure/setup-helm from 4 to 5 ([4a1a844](https://github.com/Randsw/kubeinfo/commit/4a1a8444e5a08266d3f5771948d2e7c01f134da1))
* **deps:** bump docker/build-push-action from 5 to 6 ([a748796](https://github.com/Randsw/kubeinfo/commit/a7487968d8f715f2f904e89a19244d470187660a))
* **deps:** bump docker/build-push-action from 6 to 7 ([1eb2054](https://github.com/Randsw/kubeinfo/commit/1eb2054d5a04e0d5e7c4d662cf8f484a67e82c9e))
* **deps:** bump docker/login-action from 3 to 4 ([ac57de2](https://github.com/Randsw/kubeinfo/commit/ac57de2246f469b63abe16f0bce622ebdef54d07))
* **deps:** bump docker/login-action from 4 to 4.5.2 ([eef9987](https://github.com/Randsw/kubeinfo/commit/eef9987daf5e306e7503f72af3dd80b08eba9308))
* **deps:** bump docker/login-action from 4.5.2 to 4.6.0 ([36f060d](https://github.com/Randsw/kubeinfo/commit/36f060d36c96e693d01647f9bf7af4ddc64e34c4))
* **deps:** bump docker/metadata-action from 5 to 6 ([fb39fed](https://github.com/Randsw/kubeinfo/commit/fb39fed7a352ef4c0ab9d5a85aaad171417ba397))
* **deps:** bump docker/setup-buildx-action from 3 to 4 ([9376fed](https://github.com/Randsw/kubeinfo/commit/9376fedb1665829018cafc163f62302eea8bb97c))
* **deps:** bump github.com/fluxcd/helm-controller/api ([22a1e3c](https://github.com/Randsw/kubeinfo/commit/22a1e3c353b858fb2b05eedc7f95717523c4bf2c))
* **deps:** bump github.com/fluxcd/helm-controller/api ([92a4d3f](https://github.com/Randsw/kubeinfo/commit/92a4d3f21daffb210ab47e5ee40f2e83edfe462a))
* **deps:** bump github.com/fluxcd/helm-controller/api ([2225ec7](https://github.com/Randsw/kubeinfo/commit/2225ec72bda4350c6609f4c988e1e14cdb8f4865))
* **deps:** bump github.com/fluxcd/helm-controller/api ([2422317](https://github.com/Randsw/kubeinfo/commit/2422317202d01846d4200977d42fdce5d63451a5))
* **deps:** bump github.com/fluxcd/helm-controller/api ([12dad11](https://github.com/Randsw/kubeinfo/commit/12dad118d5bdc3955e65c126c974c6c30584bbb0))
* **deps:** bump github.com/fluxcd/helm-controller/api ([7748142](https://github.com/Randsw/kubeinfo/commit/77481427743582ed278d491204085ed56c89a1ab))
* **deps:** bump github.com/fluxcd/helm-controller/api ([80b173e](https://github.com/Randsw/kubeinfo/commit/80b173e75a9fc86610e1f59a054d40b07d1fc5fa))
* **deps:** bump github.com/fluxcd/helm-controller/api ([72cb489](https://github.com/Randsw/kubeinfo/commit/72cb489c6af43ef79969ba0a2cbbaeda365192a0))
* **deps:** bump github.com/fluxcd/helm-controller/api ([14b75fa](https://github.com/Randsw/kubeinfo/commit/14b75faa9d1cb86cbd3271cfbef67b8c004763a6))
* **deps:** bump github.com/fluxcd/helm-controller/api ([94adccc](https://github.com/Randsw/kubeinfo/commit/94adccc13b0b7ebe337c3606777a8f28b50cd454))
* **deps:** bump github.com/fluxcd/helm-controller/api ([7f0d204](https://github.com/Randsw/kubeinfo/commit/7f0d20458e5d666194a1f5d3d0bbec776eee98d1))
* **deps:** bump github.com/fluxcd/helm-controller/api ([6240a9d](https://github.com/Randsw/kubeinfo/commit/6240a9d2e64eef7ab88c74f3d970414df97e7fef))
* **deps:** bump github.com/fluxcd/helm-controller/api ([8f55e2e](https://github.com/Randsw/kubeinfo/commit/8f55e2e4a3e001ef1d8b6b852404d95c1b2c9903))
* **deps:** bump github.com/fluxcd/helm-controller/api ([be84288](https://github.com/Randsw/kubeinfo/commit/be842882c049b35b0e51b58e305747760aa6938b))
* **deps:** bump github.com/fluxcd/helm-controller/api ([b1fce0e](https://github.com/Randsw/kubeinfo/commit/b1fce0ec2ce856b72028ad491f5e4243e3cc2efe))
* **deps:** bump github.com/fluxcd/helm-controller/api ([855b161](https://github.com/Randsw/kubeinfo/commit/855b161e00c8399f4d7e2a2047465b5d84f604e5))
* **deps:** bump github.com/fluxcd/helm-controller/api ([b25fb2c](https://github.com/Randsw/kubeinfo/commit/b25fb2ccd0e820d71557b69bc375bb8dfcddb103))
* **deps:** bump github.com/fluxcd/helm-controller/api ([460aafa](https://github.com/Randsw/kubeinfo/commit/460aafa870159d0543580388c74ff460615806f0))
* **deps:** bump github.com/fluxcd/helm-controller/api ([c05865e](https://github.com/Randsw/kubeinfo/commit/c05865e65ef16e197dd44e33a8e20f6889935ecc))
* **deps:** bump github.com/fluxcd/helm-controller/api ([08c163b](https://github.com/Randsw/kubeinfo/commit/08c163bfcf873f37641c6b03b0e1990da1c135ad))
* **deps:** bump github.com/fluxcd/helm-controller/api ([7ebff76](https://github.com/Randsw/kubeinfo/commit/7ebff768941c447b12dbff034c7bc419a34dfa11))
* **deps:** bump github.com/fluxcd/helm-controller/api ([557eeec](https://github.com/Randsw/kubeinfo/commit/557eeec50bf04dcb2593ed70c59eb6139646cc56))
* **deps:** bump github.com/fluxcd/helm-controller/api ([9de25f0](https://github.com/Randsw/kubeinfo/commit/9de25f0e4b447aaea63e6a3ded894f0c8cea74d1))
* **deps:** bump github.com/fluxcd/helm-controller/api ([5c8cb1d](https://github.com/Randsw/kubeinfo/commit/5c8cb1d8e5a93070359a7c8f9ff2a0e91ee671a1))
* **deps:** bump github.com/fluxcd/helm-controller/api ([b720763](https://github.com/Randsw/kubeinfo/commit/b720763629daa33c34e6ee66f15dbccc922cb276))
* **deps:** bump github.com/fluxcd/helm-controller/api ([c1c3a64](https://github.com/Randsw/kubeinfo/commit/c1c3a64d3dbfd921a904df31828d04828dd4695e))
* **deps:** bump github.com/fluxcd/helm-controller/api ([cf964ea](https://github.com/Randsw/kubeinfo/commit/cf964ea6d1ce67b911d1365b5e943c29651990e2))
* **deps:** bump github.com/fluxcd/kustomize-controller/api ([a8d602d](https://github.com/Randsw/kubeinfo/commit/a8d602da596854ff649a3deba45fa12bae138296))
* **deps:** bump github.com/fluxcd/kustomize-controller/api ([0f6a89d](https://github.com/Randsw/kubeinfo/commit/0f6a89d140c527da2bd22f41d00a122674b54608))
* **deps:** bump github.com/fluxcd/kustomize-controller/api ([dc14c3d](https://github.com/Randsw/kubeinfo/commit/dc14c3d48b7d596883a8fcd9d2be5de40cf50c42))
* **deps:** bump github.com/fluxcd/kustomize-controller/api ([348c8fd](https://github.com/Randsw/kubeinfo/commit/348c8fddd7407dc00596eabad34d6b97224ee60a))
* **deps:** bump github.com/fluxcd/kustomize-controller/api ([d7f56a5](https://github.com/Randsw/kubeinfo/commit/d7f56a50825b72b83ff180ca3c5bee15f10492b0))
* **deps:** bump github.com/fluxcd/kustomize-controller/api ([2f0f109](https://github.com/Randsw/kubeinfo/commit/2f0f1095d00e95051b0910539ae3b3cf340cdca6))
* **deps:** bump github.com/fluxcd/kustomize-controller/api ([def5e9b](https://github.com/Randsw/kubeinfo/commit/def5e9b92c6eb49d836b693cdec0c283ec2da711))
* **deps:** bump github.com/fluxcd/kustomize-controller/api ([4f003d2](https://github.com/Randsw/kubeinfo/commit/4f003d20b8e30a25963dc401883d60d0fe4646dc))
* **deps:** bump github.com/fluxcd/kustomize-controller/api ([49075e9](https://github.com/Randsw/kubeinfo/commit/49075e9e996153e54ba4cf67cf477abf338c32ff))
* **deps:** bump github.com/fluxcd/kustomize-controller/api ([6476a21](https://github.com/Randsw/kubeinfo/commit/6476a211a45f18a6877aab85bbbc375de91cb3a6))
* **deps:** bump github.com/fluxcd/kustomize-controller/api ([8a53274](https://github.com/Randsw/kubeinfo/commit/8a53274b3df127cfa013d3d4a009c80af2174543))
* **deps:** bump github.com/fluxcd/kustomize-controller/api ([30493a5](https://github.com/Randsw/kubeinfo/commit/30493a5a68e93d46e4dbdbbea3779a27d49b1ea7))
* **deps:** bump github.com/fluxcd/kustomize-controller/api ([30aaaeb](https://github.com/Randsw/kubeinfo/commit/30aaaebf78e7326f5c14b56728168920c1202411))
* **deps:** bump github.com/fluxcd/kustomize-controller/api ([64648c3](https://github.com/Randsw/kubeinfo/commit/64648c31afaa15e5d462942a9f9e7e74b1072a9e))
* **deps:** bump github.com/fluxcd/kustomize-controller/api ([cac7700](https://github.com/Randsw/kubeinfo/commit/cac7700c6b9140178c590a45f8aaed0e3209c39c))
* **deps:** bump github.com/fluxcd/kustomize-controller/api ([357f9d6](https://github.com/Randsw/kubeinfo/commit/357f9d6886a4ae981d682dcfa8fb42e81d62cdaf))
* **deps:** bump github.com/fluxcd/kustomize-controller/api ([cc424ba](https://github.com/Randsw/kubeinfo/commit/cc424ba9d9d5e86a7a913a3e21ff3accd6351553))
* **deps:** bump github.com/fluxcd/kustomize-controller/api ([a8df5ab](https://github.com/Randsw/kubeinfo/commit/a8df5ab6ed0cd93039b827ef0c12d82d8b01d744))
* **deps:** bump github.com/fluxcd/kustomize-controller/api ([e49a8f4](https://github.com/Randsw/kubeinfo/commit/e49a8f4debd296d2e59fd37969602ac2f8e7564f))
* **deps:** bump github.com/fluxcd/kustomize-controller/api ([eafc492](https://github.com/Randsw/kubeinfo/commit/eafc492617a569d1d5048326d284b1a9712dc53a))
* **deps:** bump github.com/fluxcd/kustomize-controller/api ([36d398b](https://github.com/Randsw/kubeinfo/commit/36d398b14d9b74079e7c06f759003f56cc9d26c2))
* **deps:** bump github.com/fluxcd/kustomize-controller/api ([19fc8f8](https://github.com/Randsw/kubeinfo/commit/19fc8f89decba5673313ae6548ab776666f37164))
* **deps:** bump github.com/fluxcd/kustomize-controller/api ([c979dcd](https://github.com/Randsw/kubeinfo/commit/c979dcdd5f9a2c7d10b7ceae94f892c2f6b44919))
* **deps:** bump github.com/fluxcd/kustomize-controller/api ([045c2f1](https://github.com/Randsw/kubeinfo/commit/045c2f19190ca0f2f2716841d7208869089f1d1b))
* **deps:** bump github.com/prometheus/client_golang ([b33e57d](https://github.com/Randsw/kubeinfo/commit/b33e57d7316c4c8b162dc7efd06a3a45c6f185bd))
* **deps:** bump github.com/prometheus/client_golang ([8dba62c](https://github.com/Randsw/kubeinfo/commit/8dba62cb38758aa5950dba35652883d137d7d6e2))
* **deps:** bump github.com/prometheus/client_golang ([e8abec3](https://github.com/Randsw/kubeinfo/commit/e8abec3be512618e5e317cf41403642564349b66))
* **deps:** bump github.com/prometheus/client_golang ([8d1d109](https://github.com/Randsw/kubeinfo/commit/8d1d1097bb883da445bcd6b6ae4203a8723b13bf))
* **deps:** bump github.com/prometheus/client_golang ([5290628](https://github.com/Randsw/kubeinfo/commit/52906284290e8a2576759757566e43a59c917615))
* **deps:** bump github.com/prometheus/client_golang ([b609c2a](https://github.com/Randsw/kubeinfo/commit/b609c2a0274c1ebd3ac238d532b3f473377e9aa8))
* **deps:** bump github.com/prometheus/client_golang ([4f971ca](https://github.com/Randsw/kubeinfo/commit/4f971ca8b6eba304667af5d9520fd37200a45b16))
* **deps:** bump github.com/prometheus/client_golang ([53c6361](https://github.com/Randsw/kubeinfo/commit/53c6361b1baac1863030732a6f757bee5ac372e8))
* **deps:** bump github.com/prometheus/client_golang ([4c39092](https://github.com/Randsw/kubeinfo/commit/4c39092902e8f0ae826bec4eb398eddd6e587b53))
* **deps:** bump github.com/prometheus/client_golang ([dbf0eb6](https://github.com/Randsw/kubeinfo/commit/dbf0eb6172d1303861cf18cb0fe0cab5e1c88029))
* **deps:** bump github.com/prometheus/client_golang ([f725707](https://github.com/Randsw/kubeinfo/commit/f725707ee3b81cd62613e6df41763be86bfb4fb8))
* **deps:** bump github.com/prometheus/client_golang ([a1345c2](https://github.com/Randsw/kubeinfo/commit/a1345c2b72c46fb2bcbb52df38d0c1c087bb2961))
* **deps:** bump github.com/prometheus/client_golang ([73f2126](https://github.com/Randsw/kubeinfo/commit/73f21267b80b54f6298b40c530c1411ad8dfe425))
* **deps:** bump github.com/prometheus/client_golang ([43c5465](https://github.com/Randsw/kubeinfo/commit/43c5465bd5ea9b07a0ab27b411572ccec5e5020d))
* **deps:** bump github.com/prometheus/client_golang ([0e47a80](https://github.com/Randsw/kubeinfo/commit/0e47a80ad1d28e69972fd689a9ae83b303ce78e5))
* **deps:** bump github.com/prometheus/client_golang ([9450ae0](https://github.com/Randsw/kubeinfo/commit/9450ae02a002eb964349048697af960b6614cde5))
* **deps:** bump github.com/prometheus/client_golang ([e9c74bf](https://github.com/Randsw/kubeinfo/commit/e9c74bfbbc14aefa5d7b5af94dcd8ce4f45a84c0))
* **deps:** bump github.com/swaggo/swag from 1.16.2 to 1.16.3 ([aa68dcf](https://github.com/Randsw/kubeinfo/commit/aa68dcf6d5586098e9c5ee82d3b01c25c3857036))
* **deps:** bump github.com/swaggo/swag from 1.16.3 to 1.16.4 ([7da7d09](https://github.com/Randsw/kubeinfo/commit/7da7d098d7599f0e391047d7c105065e82a09d4e))
* **deps:** bump github.com/swaggo/swag from 1.16.4 to 1.16.5 ([5b89f2f](https://github.com/Randsw/kubeinfo/commit/5b89f2fc1d8b668242a87c653a96f67c68d805d1))
* **deps:** bump github.com/swaggo/swag from 1.16.5 to 1.16.6 ([65a434d](https://github.com/Randsw/kubeinfo/commit/65a434da08af0979eae71f393f51d99b2951bf6e))
* **deps:** bump github.com/swaggo/swag from 1.8.1 to 1.16.2 ([142a844](https://github.com/Randsw/kubeinfo/commit/142a8447df721448e89e70ba8994e20e1963b3dc))
* **deps:** bump Go to 1.22 ([f271eed](https://github.com/Randsw/kubeinfo/commit/f271eeddb2bba30201865c6ba6973c6cc26a7905))
* **deps:** bump go.uber.org/zap from 1.26.0 to 1.27.0 ([ff83e7f](https://github.com/Randsw/kubeinfo/commit/ff83e7f6c8be88f3f2f594b8bacb3f1da77bbb57))
* **deps:** bump go.uber.org/zap from 1.27.0 to 1.27.1 ([1eabffb](https://github.com/Randsw/kubeinfo/commit/1eabffbb8d50043a2a7bf5d4c13e5a772f068f59))
* **deps:** bump go.uber.org/zap from 1.27.1 to 1.28.0 ([9fc9883](https://github.com/Randsw/kubeinfo/commit/9fc9883903f69cd2dc7965e3708efc78d1647206))
* **deps:** bump golang from 1.21 to 1.22 ([0adbc46](https://github.com/Randsw/kubeinfo/commit/0adbc462d36a0a48f558b38192ea0669f0fb7492))
* **deps:** bump golang from 1.22 to 1.23 ([c0af939](https://github.com/Randsw/kubeinfo/commit/c0af939703fe0137dec6369a73f7b59d8c98d288))
* **deps:** bump golang from 1.23 to 1.24 ([68602fe](https://github.com/Randsw/kubeinfo/commit/68602fe005caeb13ed23f7864ee7bf092b29d63d))
* **deps:** bump golang from 1.24 to 1.25 ([0f0392d](https://github.com/Randsw/kubeinfo/commit/0f0392dac28ab34e09de104b603d6f44b43f10d6))
* **deps:** bump golang from 1.25 to 1.26 ([56e6198](https://github.com/Randsw/kubeinfo/commit/56e619843351d029b5383291a775464288844e04))
* **deps:** bump golang from 1.26 to 1.27 ([104f0bb](https://github.com/Randsw/kubeinfo/commit/104f0bb709606b7baf71bf2c3d0c5b4d8079f5dd))
* **deps:** Bump golang version ([0740a57](https://github.com/Randsw/kubeinfo/commit/0740a576be9ffbfbe6c5fdec8df32d3956f5f489))
* **deps:** bump golangci/golangci-lint-action from 3 to 4 ([c352b13](https://github.com/Randsw/kubeinfo/commit/c352b1320ede9d8620abe6221c6971cbafea8221))
* **deps:** bump golangci/golangci-lint-action from 4 to 5 ([34ebc7c](https://github.com/Randsw/kubeinfo/commit/34ebc7cc6511f701760ae9fbbc13ee8280bc51c9))
* **deps:** bump golangci/golangci-lint-action from 5 to 6 ([f9c5b3b](https://github.com/Randsw/kubeinfo/commit/f9c5b3b3e37fd81594ec65b4429bb04e64389b42))
* **deps:** bump golangci/golangci-lint-action from 6 to 7 ([33321d2](https://github.com/Randsw/kubeinfo/commit/33321d20739f28c2f20b45bc9e96a7dc812d6151))
* **deps:** bump golangci/golangci-lint-action from 7 to 8 ([05871c9](https://github.com/Randsw/kubeinfo/commit/05871c9f25423b7dd83b223fa823fe85e7a1f2ef))
* **deps:** bump golangci/golangci-lint-action from 8 to 9 ([092e573](https://github.com/Randsw/kubeinfo/commit/092e57372b4e5129b4e424b7da44761609859e98))
* **deps:** bump helm/chart-releaser-action from 1.6.0 to 1.7.0 ([3e23e5d](https://github.com/Randsw/kubeinfo/commit/3e23e5dd756e3ccaf7ca6ebe2a7fd9cc4bb925ca))
* **deps:** bump helm/chart-testing-action from 2.6.1 to 2.7.0 ([8e9cc75](https://github.com/Randsw/kubeinfo/commit/8e9cc7546a03c1e9620f719b1e45b2fb6f085849))
* **deps:** bump helm/chart-testing-action from 2.7.0 to 2.8.0 ([4242832](https://github.com/Randsw/kubeinfo/commit/4242832ce478664bd95cad4117616992ab6b4e23))
* **deps:** bump helm/kind-action from 1.10.0 to 1.11.0 ([5b4ecda](https://github.com/Randsw/kubeinfo/commit/5b4ecda5ad38add58bb2ba1cbc53258fb1171218))
* **deps:** bump helm/kind-action from 1.11.0 to 1.12.0 ([f6105d5](https://github.com/Randsw/kubeinfo/commit/f6105d59b40a60832331463931781285bab41c5c))
* **deps:** bump helm/kind-action from 1.12.0 to 1.13.0 ([bbcbb6d](https://github.com/Randsw/kubeinfo/commit/bbcbb6d7216bfd6413ddc7c6629858a98fabf932))
* **deps:** bump helm/kind-action from 1.13.0 to 1.14.0 ([85c5eb9](https://github.com/Randsw/kubeinfo/commit/85c5eb9e2bf84332a7edc4592a90eed358f086ab))
* **deps:** bump helm/kind-action from 1.14.0 to 1.15.0 ([c5ffefb](https://github.com/Randsw/kubeinfo/commit/c5ffefb74cc09d6b917086f75bf095f125726411))
* **deps:** bump helm/kind-action from 1.8.0 to 1.9.0 ([e0b8b4f](https://github.com/Randsw/kubeinfo/commit/e0b8b4fd1c776d2c41e730c36cb9e5b2806fe264))
* **deps:** bump helm/kind-action from 1.9.0 to 1.10.0 ([0a6ec51](https://github.com/Randsw/kubeinfo/commit/0a6ec516c46abc44333764941cdf1d04000a4b78))
* **deps:** bump k8s.io/api from 0.32.3 to 0.32.4 ([9f5ad5f](https://github.com/Randsw/kubeinfo/commit/9f5ad5f30f9bb321b1f090e05ae1e3947e8e560e))
* **deps:** bump k8s.io/api from 0.36.3 to 0.36.4 ([e3cfccd](https://github.com/Randsw/kubeinfo/commit/e3cfccd0203df1d331511907fec98ee4417e3fe4))
* **deps:** bump k8s.io/apimachinery from 0.32.3 to 0.32.4 ([b21d4dd](https://github.com/Randsw/kubeinfo/commit/b21d4dddb7c7ca9db128367abc846efbede2eef0))
* **deps:** bump k8s.io/client-go from 0.28.4 to 0.29.0 ([0fe1a00](https://github.com/Randsw/kubeinfo/commit/0fe1a000fb145579a6f7883705c69994984ff4f3))
* **deps:** bump k8s.io/client-go from 0.29.0 to 0.29.1 ([0688f57](https://github.com/Randsw/kubeinfo/commit/0688f5773bea83eae759511361bc5293a8eab0f3))
* **deps:** bump k8s.io/client-go from 0.29.1 to 0.29.2 ([39b0551](https://github.com/Randsw/kubeinfo/commit/39b0551b3b7985181b3993d3606308b7b49712d1))
* **deps:** bump k8s.io/client-go from 0.29.2 to 0.29.3 ([0cc898e](https://github.com/Randsw/kubeinfo/commit/0cc898e4291753c4000bb4669474d390301f60da))
* **deps:** bump k8s.io/client-go from 0.29.3 to 0.29.4 ([5fe3a00](https://github.com/Randsw/kubeinfo/commit/5fe3a008b99f7ee4a13e0365094bc63d79fd2f5d))
* **deps:** bump k8s.io/client-go from 0.29.4 to 0.30.0 ([f4a287e](https://github.com/Randsw/kubeinfo/commit/f4a287e16b4ce222fa9ce3aacf5ad00cd3504a77))
* **deps:** bump k8s.io/client-go from 0.30.0 to 0.30.1 ([dc133b0](https://github.com/Randsw/kubeinfo/commit/dc133b02eacc7c33e6410a99b59650d55dabf1c6))
* **deps:** bump k8s.io/client-go from 0.30.1 to 0.30.2 ([2afdc96](https://github.com/Randsw/kubeinfo/commit/2afdc9607a841caaa1d8367a6ec3122330324965))
* **deps:** bump k8s.io/client-go from 0.30.2 to 0.30.3 ([a2cda51](https://github.com/Randsw/kubeinfo/commit/a2cda51fe4e0f8856049ae0e2afc7f5c8cb3a157))
* **deps:** bump k8s.io/client-go from 0.30.3 to 0.31.0 ([a2ee575](https://github.com/Randsw/kubeinfo/commit/a2ee5756894b4de7f4ed69201232cd13246e36eb))
* **deps:** bump k8s.io/client-go from 0.31.0 to 0.31.1 ([bf35adf](https://github.com/Randsw/kubeinfo/commit/bf35adf1e16e26cabf276b30befa0fca20713967))
* **deps:** bump k8s.io/client-go from 0.31.1 to 0.31.2 ([ff2517a](https://github.com/Randsw/kubeinfo/commit/ff2517a48a4b2d35a22cf91e6ea96aaa755697be))
* **deps:** bump k8s.io/client-go from 0.31.2 to 0.31.3 ([7e2761a](https://github.com/Randsw/kubeinfo/commit/7e2761ad699f50f16ad9c46a8cc4606802d76ecd))
* **deps:** bump k8s.io/client-go from 0.31.3 to 0.31.4 ([ac45886](https://github.com/Randsw/kubeinfo/commit/ac458864d16b9da5fce52380a0a543a18592c3c8))
* **deps:** bump k8s.io/client-go from 0.31.4 to 0.32.0 ([a6dcfac](https://github.com/Randsw/kubeinfo/commit/a6dcfac4a26d8c972c0716239a5971ace42a5cb1))
* **deps:** bump k8s.io/client-go from 0.32.0 to 0.32.1 ([048f3d3](https://github.com/Randsw/kubeinfo/commit/048f3d3dbec5c69259b945aa0f459984db8236e5))
* **deps:** bump k8s.io/client-go from 0.32.1 to 0.32.2 ([3897f86](https://github.com/Randsw/kubeinfo/commit/3897f8608004126cd80d63572798f0ea7499b4ba))
* **deps:** bump k8s.io/client-go from 0.32.2 to 0.32.3 ([66e7cda](https://github.com/Randsw/kubeinfo/commit/66e7cdaec429459c2863207d477398d2a49c523f))
* **deps:** bump k8s.io/client-go from 0.32.3 to 0.32.4 ([de0510e](https://github.com/Randsw/kubeinfo/commit/de0510e011dae9f7be9ed9a79c161488be949278))
* **deps:** bump k8s.io/client-go from 0.32.4 to 0.33.0 ([9503f05](https://github.com/Randsw/kubeinfo/commit/9503f05ca22d6b310e638b09d0695f5883690bc8))
* **deps:** bump k8s.io/client-go from 0.33.0 to 0.33.1 ([2b7521c](https://github.com/Randsw/kubeinfo/commit/2b7521c78feb20a867d3dc0b3c0c0220ea308722))
* **deps:** bump k8s.io/client-go from 0.33.1 to 0.33.2 ([059eb3b](https://github.com/Randsw/kubeinfo/commit/059eb3ba8627b699c3dacc59718aafcd1dcc70dc))
* **deps:** bump k8s.io/client-go from 0.33.2 to 0.33.3 ([97a066e](https://github.com/Randsw/kubeinfo/commit/97a066e12974203d6e125e9ad3bb4a8c47cbdaca))
* **deps:** bump k8s.io/client-go from 0.33.3 to 0.33.4 ([92ef9c3](https://github.com/Randsw/kubeinfo/commit/92ef9c371039da2180ad6c19fadddd8c1831e4d5))
* **deps:** bump k8s.io/client-go from 0.33.4 to 0.34.0 ([345de98](https://github.com/Randsw/kubeinfo/commit/345de98102bdcc27a5d4affc9bfdb2f98ddb5799))
* **deps:** bump k8s.io/client-go from 0.34.0 to 0.34.1 ([a775b35](https://github.com/Randsw/kubeinfo/commit/a775b358b7a18dbc2ae4874dfdd7a1ad606f3b15))
* **deps:** bump k8s.io/client-go from 0.34.1 to 0.34.2 ([45e12ba](https://github.com/Randsw/kubeinfo/commit/45e12ba7afe3e047b9953367c6b6795dbb46174b))
* **deps:** bump k8s.io/client-go from 0.34.2 to 0.34.3 ([b6c74fa](https://github.com/Randsw/kubeinfo/commit/b6c74fa59a1d30999b6b98dbd3674e880525099c))
* **deps:** bump k8s.io/client-go from 0.34.3 to 0.35.0 ([0765aa3](https://github.com/Randsw/kubeinfo/commit/0765aa356adb98e1f2fe41444699b1736b591bf1))
* **deps:** bump k8s.io/client-go from 0.35.0 to 0.35.1 ([59e22f5](https://github.com/Randsw/kubeinfo/commit/59e22f5a204b9cd2fa101b5e6af98cd43f1760c7))
* **deps:** bump k8s.io/client-go from 0.35.1 to 0.35.2 ([1cc78ff](https://github.com/Randsw/kubeinfo/commit/1cc78ff9df3edaa4af4fef29df0b6f96e1cf5bb6))
* **deps:** bump k8s.io/client-go from 0.35.2 to 0.35.3 ([506dcfd](https://github.com/Randsw/kubeinfo/commit/506dcfdf71a8f5adba8a1bd8a7c163844957ccb1))
* **deps:** bump k8s.io/client-go from 0.35.3 to 0.35.4 ([7f062d6](https://github.com/Randsw/kubeinfo/commit/7f062d6ee41c216235b31a178660a8072379b6f2))
* **deps:** bump k8s.io/client-go from 0.35.4 to 0.36.0 ([459b6e7](https://github.com/Randsw/kubeinfo/commit/459b6e750c0564ea9f7c6b52066364458b15927d))
* **deps:** bump k8s.io/client-go from 0.36.0 to 0.36.1 ([8d9b1e7](https://github.com/Randsw/kubeinfo/commit/8d9b1e7e63c43301cde2553c0c746be50d391204))
* **deps:** bump k8s.io/client-go from 0.36.1 to 0.36.2 ([774a372](https://github.com/Randsw/kubeinfo/commit/774a3727be0163e6fffb0654c720e35bfae26c1b))
* **deps:** bump k8s.io/client-go from 0.36.2 to 0.36.3 ([5f65dc7](https://github.com/Randsw/kubeinfo/commit/5f65dc70e1ce1031ed4c08a9fb8690d05bfd9f1d))
* **deps:** bump k8s.io/client-go from 0.36.3 to 0.36.4 ([a52244d](https://github.com/Randsw/kubeinfo/commit/a52244dabc531ac7539224e2bd73855ce3b3128c))
* **deps:** bump k8s.io/client-go from 0.36.4 to 0.37.0 ([3f8772c](https://github.com/Randsw/kubeinfo/commit/3f8772c3c94bc4fc62832e92fd125adda0e96090))
* **deps:** bump softprops/action-gh-release from 1 to 2 ([ff06ad4](https://github.com/Randsw/kubeinfo/commit/ff06ad4183755a8f81796dde0712a5388cb1ce55))
* **deps:** bump softprops/action-gh-release from 2 to 3 ([e2f7e7a](https://github.com/Randsw/kubeinfo/commit/e2f7e7acedfcff73a18bb921aadde48b889621f7))
* **deps:** bump tj-actions/changed-files from 40 to 41 ([6f00269](https://github.com/Randsw/kubeinfo/commit/6f002691d2a18a2f81801bc92ccabfb91a6ac4b5))
* **deps:** bump tj-actions/changed-files from 41 to 42 ([f10c98e](https://github.com/Randsw/kubeinfo/commit/f10c98e46658489e8c99f23bcd60bcfc8b2e76da))
* **deps:** bump tj-actions/changed-files from 42 to 43 ([583e213](https://github.com/Randsw/kubeinfo/commit/583e21354fb83f1a71be1342d8bd308dd4b7f68b))
* **deps:** bump tj-actions/changed-files from 43 to 44 ([0d840b7](https://github.com/Randsw/kubeinfo/commit/0d840b70ed4f9fc8af0b4e6df9f77690002cd6c5))
* **deps:** bump tj-actions/changed-files from 44 to 45 ([5cedc30](https://github.com/Randsw/kubeinfo/commit/5cedc304565b50d51bd828717bb0245ce8dd6846))
* **deps:** bump tj-actions/changed-files from 45 to 46 ([c87708e](https://github.com/Randsw/kubeinfo/commit/c87708ee1af1e953dce2270768539038cb40ead1))
* **deps:** bump tj-actions/changed-files from 46 to 47 ([8427e9b](https://github.com/Randsw/kubeinfo/commit/8427e9bf8fc7a65988f2749b7d55eb9b4b8bdffa))
* **golang:** Bump go version ([c3ffc80](https://github.com/Randsw/kubeinfo/commit/c3ffc8024df40493f83c0a129d3e7e9e08662881))
* **modules:** Bump helm-controller version ([7e6c95e](https://github.com/Randsw/kubeinfo/commit/7e6c95e34420d0cb68034b18a51aaeab10cc5fd2))
* **modules:** Fix helm-controller tests ([736b819](https://github.com/Randsw/kubeinfo/commit/736b81910e8f2b6c7379fe7c419586971e4b260b))

## [1.2.0](https://github.com/Randsw/kubeinfo/compare/1.1.8...1.2.0) (2023-11-30)


### 📔 Docs

* Fix typo ([bd15b13](https://github.com/Randsw/kubeinfo/commit/bd15b134afbf50dbd16a071fa71709ef4b0a3271))


### 🦊 CI/CD

* use latest version of chart-test ([9b8bc65](https://github.com/Randsw/kubeinfo/commit/9b8bc65f8da83a420008dccc5ddabd1b6a583bcf))


### 🚀 Features

* **api:** Add swagger page ([e7b36ca](https://github.com/Randsw/kubeinfo/commit/e7b36ca671df68fb71fa4da1fcc240729496a6a7))


### 🛠 Fixes

* **api:** Add all handlers to swagger docs ([678b5eb](https://github.com/Randsw/kubeinfo/commit/678b5eb7757af9d0610377785a07bd974005328a))
* **api:** Add failure field to swagger docs ([3718307](https://github.com/Randsw/kubeinfo/commit/3718307c600ddd26531b008c935aa997c1fcf6e5))
* **api:** Fix api host to execute from container ([f01d4bc](https://github.com/Randsw/kubeinfo/commit/f01d4bcb093609b7f07f0f46ed5044c7b552a0f8))
* **api:** Fix router path in swagger ([1467f1a](https://github.com/Randsw/kubeinfo/commit/1467f1a5364f2c678961fe6079f6b879e03e58d1))
* **api:** Fix swagger doc import path ([bcfc455](https://github.com/Randsw/kubeinfo/commit/bcfc455c00b2b086671e76b440e6fe7200d64a2d))
* **api:** Update swagger files ([4352ffd](https://github.com/Randsw/kubeinfo/commit/4352ffd246c7dde7a8dd69d6b1efa5c149330071))


### Other

* **deps:** bump github.com/gorilla/mux from 1.8.0 to 1.8.1 ([490825c](https://github.com/Randsw/kubeinfo/commit/490825c3ea1d37d5d43a53d58c1978728b82e3cb))
* **deps:** bump helm/chart-releaser-action from 1.5.0 to 1.6.0 ([84649c8](https://github.com/Randsw/kubeinfo/commit/84649c8ee149d63bf8ab2a2461963dd2369a107e))
* **deps:** bump helm/chart-testing-action from 2.6.0 to 2.6.1 ([9293815](https://github.com/Randsw/kubeinfo/commit/929381522f7f91f66156a7a1ad13667544d2bf8d))
* **deps:** bump k8s.io/client-go from 0.28.3 to 0.28.4 ([6364b48](https://github.com/Randsw/kubeinfo/commit/6364b48aaa02274d8e58f361c0584bdd512a039d))
* **deps:** bump randsw/kubeinfo in /helm-chart/kubeinfo-backend ([92b28f9](https://github.com/Randsw/kubeinfo/commit/92b28f9d2031b3786875ca09c8ada6a0c87da84e))
* **release:** 1.2.0-develop.1 ([42b941e](https://github.com/Randsw/kubeinfo/commit/42b941e2f48d202683cf5236fcdced6242a53924))
* **release:** 1.2.0-develop.2 ([784ba16](https://github.com/Randsw/kubeinfo/commit/784ba164e3f0c29b4ec5688d9626c7a88ccfd260))
* **release:** 1.2.0-develop.3 ([205a064](https://github.com/Randsw/kubeinfo/commit/205a064f05ababd91906b64d2a046a8c348836a2))
* **release:** 1.2.0-develop.4 ([3bdacaf](https://github.com/Randsw/kubeinfo/commit/3bdacaf46cf7ca384a13e6cef53f380f9b1618a4))
* **release:** 1.2.0-develop.5 ([c9e9d7b](https://github.com/Randsw/kubeinfo/commit/c9e9d7b9e444ea1c31a746557a17c63022171e84))
* **release:** 1.2.0-develop.6 ([07f9c25](https://github.com/Randsw/kubeinfo/commit/07f9c25ddce81da1d7935acbddd265a11fedd836))
* **release:** 1.2.0-develop.7 ([ebf6506](https://github.com/Randsw/kubeinfo/commit/ebf65063d7a51b1d75efc53281c0ec28e52d8d6c))

## [1.2.0-develop.7](https://github.com/Randsw/kubeinfo/compare/1.2.0-develop.6...1.2.0-develop.7) (2023-11-30)


### 🛠 Fixes

* **api:** Update swagger files ([4352ffd](https://github.com/Randsw/kubeinfo/commit/4352ffd246c7dde7a8dd69d6b1efa5c149330071))

## [1.2.0-develop.6](https://github.com/Randsw/kubeinfo/compare/1.2.0-develop.5...1.2.0-develop.6) (2023-11-30)


### 🛠 Fixes

* **api:** Add failure field to swagger docs ([3718307](https://github.com/Randsw/kubeinfo/commit/3718307c600ddd26531b008c935aa997c1fcf6e5))

## [1.2.0-develop.5](https://github.com/Randsw/kubeinfo/compare/1.2.0-develop.4...1.2.0-develop.5) (2023-11-30)


### 🛠 Fixes

* **api:** Add all handlers to swagger docs ([678b5eb](https://github.com/Randsw/kubeinfo/commit/678b5eb7757af9d0610377785a07bd974005328a))

## [1.2.0-develop.4](https://github.com/Randsw/kubeinfo/compare/1.2.0-develop.3...1.2.0-develop.4) (2023-11-29)


### 🛠 Fixes

* **api:** Fix router path in swagger ([1467f1a](https://github.com/Randsw/kubeinfo/commit/1467f1a5364f2c678961fe6079f6b879e03e58d1))

## [1.2.0-develop.3](https://github.com/Randsw/kubeinfo/compare/1.2.0-develop.2...1.2.0-develop.3) (2023-11-29)


### 🛠 Fixes

* **api:** Fix api host to execute from container ([f01d4bc](https://github.com/Randsw/kubeinfo/commit/f01d4bcb093609b7f07f0f46ed5044c7b552a0f8))

## [1.2.0-develop.2](https://github.com/Randsw/kubeinfo/compare/1.2.0-develop.1...1.2.0-develop.2) (2023-11-29)


### 🛠 Fixes

* **api:** Fix swagger doc import path ([bcfc455](https://github.com/Randsw/kubeinfo/commit/bcfc455c00b2b086671e76b440e6fe7200d64a2d))

## [1.2.0-develop.1](https://github.com/Randsw/kubeinfo/compare/1.1.8...1.2.0-develop.1) (2023-11-29)


### 🦊 CI/CD

* use latest version of chart-test ([9b8bc65](https://github.com/Randsw/kubeinfo/commit/9b8bc65f8da83a420008dccc5ddabd1b6a583bcf))


### 🚀 Features

* **api:** Add swagger page ([e7b36ca](https://github.com/Randsw/kubeinfo/commit/e7b36ca671df68fb71fa4da1fcc240729496a6a7))


### Other

* **deps:** bump github.com/gorilla/mux from 1.8.0 to 1.8.1 ([490825c](https://github.com/Randsw/kubeinfo/commit/490825c3ea1d37d5d43a53d58c1978728b82e3cb))
* **deps:** bump helm/chart-releaser-action from 1.5.0 to 1.6.0 ([84649c8](https://github.com/Randsw/kubeinfo/commit/84649c8ee149d63bf8ab2a2461963dd2369a107e))
* **deps:** bump helm/chart-testing-action from 2.6.0 to 2.6.1 ([9293815](https://github.com/Randsw/kubeinfo/commit/929381522f7f91f66156a7a1ad13667544d2bf8d))
* **deps:** bump k8s.io/client-go from 0.28.3 to 0.28.4 ([6364b48](https://github.com/Randsw/kubeinfo/commit/6364b48aaa02274d8e58f361c0584bdd512a039d))
* **deps:** bump randsw/kubeinfo in /helm-chart/kubeinfo-backend ([92b28f9](https://github.com/Randsw/kubeinfo/commit/92b28f9d2031b3786875ca09c8ada6a0c87da84e))

## [1.1.8](https://github.com/Randsw/kubeinfo/compare/1.1.7...1.1.8) (2023-11-01)


### 💈 Style

* Remove debug part ([4f5e57c](https://github.com/Randsw/kubeinfo/commit/4f5e57ced63f8d9acefc00c8b15a9cddc21af403))


### 🦊 CI/CD

* Bump chart-testing-action to v2.6.0 ([3f6f678](https://github.com/Randsw/kubeinfo/commit/3f6f6782589a45c3ffc0359df8c0e3a28987837a))
* Downgrade chart-testing to v3.8.0 ([2e5eb81](https://github.com/Randsw/kubeinfo/commit/2e5eb81ffe8e26860792e2c9d53e5313beaec331))
* Downgrade chart-testing-action to v2.4.0 ([6880e27](https://github.com/Randsw/kubeinfo/commit/6880e2706b2b7bac2403d4b0a75c25335f885fec))


### 🛠 Fixes

* Add timeouts to http server ([a759aaf](https://github.com/Randsw/kubeinfo/commit/a759aaf7d8d39c9bb38452e2bfbd8f34c762b733))


### Other

* **deps:** bump helm/chart-testing-action from 2.4.0 to 2.5.0 ([9e2bc08](https://github.com/Randsw/kubeinfo/commit/9e2bc081d38bf1ccc8b2683a986cf03fce06fd66))
* **deps:** bump randsw/kubeinfo in /helm-chart/kubeinfo-backend ([37da241](https://github.com/Randsw/kubeinfo/commit/37da241039011c2c58a2280bad310b49f83a4c50))

## [1.1.7](https://github.com/Randsw/kubeinfo/compare/1.1.6...1.1.7) (2023-10-31)


### 🛠 Fixes

* Remove .git from .dockerignore ([07f1930](https://github.com/Randsw/kubeinfo/commit/07f193010b3b0ae8de96a6b60c1cbff9a5afb6dc))

## [1.1.6](https://github.com/Randsw/kubeinfo/compare/1.1.5...1.1.6) (2023-10-31)


### 🛠 Fixes

* Huge test ([25ea6b2](https://github.com/Randsw/kubeinfo/commit/25ea6b2e512c84e05b1cea6bb79f46bdfb3934b8))

## [1.1.5](https://github.com/Randsw/kubeinfo/compare/1.1.4...1.1.5) (2023-10-31)


### 🛠 Fixes

* Try to sovle problem with empty status field ([9d01c24](https://github.com/Randsw/kubeinfo/commit/9d01c249784e2108daae88124a911eda1c6f61db))

## [1.1.4](https://github.com/Randsw/kubeinfo/compare/1.1.3...1.1.4) (2023-10-31)


### 🛠 Fixes

* Downgrade checkout version in build ([9541d01](https://github.com/Randsw/kubeinfo/commit/9541d01845e8dbbfb0a75269e682d80667f061d2))

## [1.1.3](https://github.com/Randsw/kubeinfo/compare/1.1.2...1.1.3) (2023-10-31)


### 🛠 Fixes

* Try to get tag and hash ([8493642](https://github.com/Randsw/kubeinfo/commit/8493642bc13f3e17f8ac4e6c68617fca146ab688))

## [1.1.2](https://github.com/Randsw/kubeinfo/compare/1.1.1...1.1.2) (2023-10-31)


### 🛠 Fixes

* Configure git to trust the workspace despite the different owner ([0d96e5a](https://github.com/Randsw/kubeinfo/commit/0d96e5a4565356b3e4e151e58bb83fad80dc54cb))
* Rename field in info ([5397d8e](https://github.com/Randsw/kubeinfo/commit/5397d8e8f2ff60ee6dd0f7f60206be60e43fe8cf))

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
