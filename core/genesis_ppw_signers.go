package core

import (
	"bytes"
	//"fmt"
)


func getPpwSignStr() string {

	var buf bytes.Buffer
	buf.WriteString("0x")

	for _,str := range ppwSigAddr {
		buf.WriteString(str[2:])
	}

	//fmt.Print(buf.String())
	return buf.String()

}

var (
	ppwSigAddr []string = []string{
		"0xf9b32578b4420a36f132db32b56f3831a7cc1804",
		"0xdb05642eabc8347ec78e21bdf0d906ba579d423a",
		"0x0036805b6846f26ac35f2a7d7eda4a2a58f08e8e",
		"0xc38eb01bce9bcb61327532dc5a540da4cf484ae5",
		"0x3a46ef1eb55428b3b88a222d80d23531054ef51d",
		"0x810524175efa012446103d1a04c9f4263a962acc",
		"0xb5eb9bf02a924367ed9d4f86dfcb1c572cd9a4f8",
		"0xf073d4e52c506f3f288faa9db1c1e5ae0f1e70f8",
		"0x7e98bc5a465c1d2afa6b9376709a525981f53d49",
		"0xbd100cf8286136659a7d63a38a154e28dbf3e0fd",
		"0x68e0a2f8d881a0dbd815835b8f20e0c15aac05d0",
		"0x6830a4eeded48c8ef3a9526fb672582f211c28fd",
		"0x5fa6a1fc058146d4bb507214791382d96cce0f1d",
		"0x8493b1e6ee8a580c1efeb26582855bb7d0134352",
		"0x833241056ccbb60c58a6ebeb721e09c7d968a7b2",
		"0x6c1f579d414a070623b329cfe2b4960d53358aac",
		"0xa9f1cd430f0a789c67f66f374e711b3e6127c25e",
		"0xa5c1347f172cefca33d577b9af894b8b54f5e80b",
		"0xac8bfaa9d7832592b588d56b0e5e03ecad71c659",
		"0x65277ca65ccaca2423e12f2fa045c7b752406102",
		"0xd8408f099ed5d93c5fe7e92aa2152c8a062aabd9",
		"0xd1def5573896b1371ced4e0d6660dad0db3ed736",
		"0xe3e6a874af7402b57cbb3d16c7fff6846ae4e4a0",
		"0x801747afdff28748278a3b49517c6c1629a6ed9c",
		"0xa29f4cad1494fbf57bd3267fc577653f1e1d3877",
		"0x5597180afa35fb00023e492202c195c878a6e24a",
		"0x5be9ce19552749f2edf799adedb3ef492e2928d2",
		"0xb9fd5c1d7f3692a1a761066fa3ff3d4aa86dbd41",
		"0x1623ca4c17cc38df7bef9d9ff4ed4b987bf6c838",
		"0x730900e35821218fab5f517152abf1161c1857fe",
		"0x79d4790897f3c220a114c5885868145627a278f7",
		"0x27476cec5ef161a75abb950b52f8ee9cdd45b8cc",
		"0x0040675b5cf755ebfa8da05f87f349d170fa4a80",
		"0x216383c4698759f69b7fac0bf6474712d065e804",
		"0x7eb355450a3644e774fd2fd9901a08cf8a1f633e",
		"0xceede69d776c6cc34a33cf9bfa18f59d01b46865",
		"0x5454f88fa6e91ef90a40072c7bc240512b1f18e3",
		"0xe1b6e7c61ca228114c8a2257fe1131a534160f21",
		"0x1ba18c63f6f69d199fc69c2e443dfa137cc0c048",
		"0x3773e4ecb5c0aeccc3a4c109de62ebc4baba3cec",
		"0x38b6e52b2c95c740d1d02889fdc51ae3d7fd4e75",
		"0x19d4772b9bca7941b27b48cf953002228b34de13",
		"0x1864c2a7966ccda0b6673f832695439354dacd15",
		"0xce2200d3a52721cf4ca879594427ea10eac8a48d",
		"0x7132b8c6c61083d24530de543793ab11484e4c54",
		"0x0dd3d7bf09006f09d7db3464c43ae0d9173a513a",
		"0x7883f5b46dd47fd75fbd048a411d44097e5223db",
		"0x86544220bb734b76af5e6fb5bf50c93a60fd4e7d",
		"0x1356b30952ff70c0bf4cba5f4d2c4f6bea115eb8",
		"0xbc42c1ac014a6137c7a52a3b4140cbb6ab6f4187",
		"0xf96746f4fc4bd48ee41b12ccfa1687187e6e803d",
		"0xeb91c9c64d0e430351c028ab67a99af6b7c72e5f",
		"0x1bcb29a8326f9701490ccb605943db701f1ac3fb",
		"0x28b94fbd7da585374fb4ce6620cc17a1a820cd9b",
		"0x9cdba349d5b7a9e698d51284c870d4e170a0765c",
		"0x4e28fd02fb4d3b00cb4e99ff22852edfd95c8457",
		"0x5f1f4278158396737b2ed22ff9fa118b8e43989b",
		"0x9cced7d1d0588edbd53b14de498fd27b11014e80",
		"0xca960064b6cb4d94a1137439fe5cb0b49e204d65",
		"0x46e1ff40ebd8a6cebba7145b433329880839fddf",
		"0x5be0e0ec189abc7af7806b747d002b076a191378",
		"0xb65259a8dd6212a137deb600717f8801ff80eb13",
		"0x5c467138dd34af411403fd18eca9eda2260221b0",
		"0xe930d1e8fefe04f84733246827e528a84920a033",
		"0xc7a12064816a1487ed21168ad9097b8f15fd5f58",
		"0xd100d2b093cb1511a0c61af547c943185cdb279d",
		"0x0c7d24f5c1f563b3ca76ff33d451bd71d0d77495",
		"0xd5c91d4ad971379fd53766e1ddf54c649ce81ea3",
		"0xe85b2cab9feb4f65f131bf79b9448751afc3dd4a",
		"0x8786d7b6a8d8d69219a592a50986f6182a72cab9",
		"0x73217ac73c478eb1d86fc470842e32a7d0ba5caf",
		"0x9d97bafc67812f89c623be1e5f1c5748b12cf59f",
		"0x25912a4946c39d9070274f3b196a60026964f89c",
		"0x3ad255c35d7cdf89d202cf8b2de250d3c7125d91",
		"0x716484462d638e1bb1cadda90f15bcb23dfba086",
		"0x9291f404a2d49625ab78d8856783cb84d94ad8ad",
		"0xb8d1d15026bcb5393416d1a48185b66b973f30aa",
		"0x9d1164d9a90a942089f30cebb7dc3cad587142cb",
		"0xbe02a4a5b32fdd931eb488b5d8472dc27562fab1",
		"0x6a5f9f0b2762ec7303c111e41d6fc26e986e9007",
		"0x549a2599d7ec58c6c90ecafb9c2fe0d609c3f6f7",
		"0xc37e02e0cb45347d02ecb1a9a9382d8f04de887d",
		"0xcab94f9847d304bb5f10da056744049ebe69838e",
		"0x7d885d05507a376de1a25cef42ecea9f883e309f",
		"0xc5352b0ed51087c6ca8bcd3fa59823ed823a7474",
		"0x487e5fdcff164f0ef10dd9b675bd4bfcc4e89256",
		"0xc09ab495180da5892175311ef50bb33e046a1b10",
		"0xfe8bb113afda7c5b17341d899e5c9be58bcc99b4",
		"0x87093b58c9df499decc977fee7d31136e36fde77",
		"0x6b156753e059df03490a59e0fc940d620149b002",
		"0xb0da7fae9ef832a616377e79b50a289aa4900bb2",
		"0x563563ac0aeaf3f90c6d9fec74aea3cdf8868722",
		"0xcdc751b958d1ede4fce40e6f72bea7f30bb88b09",
		"0xe884e86d3ab76cd567860b16714665869ce92d32",
		"0xa03c3ea097c0f7f1194520db30798d829f035293",
		"0x5f5232b08b55bb16a3e295f2fbe12cdc632ccaaf",
		"0xe9f61bdca8e590a9a44ff9e969b0cbb8e8c33512",
		"0x7652a05cec7bbb4ab8563d3e5fc411ae9fa44798",
		"0xd9ee761b480682a86168b63db5fbe0a6e773fc34",
		"0x1847b62a4b2fdb554e347f424ec60dd4f929d778",
		"0x4c657f5c3e19ac90c056ed0d25f182210f671166",
		"0xff293478800e0e12d8f6023c7a769441872402c9",
		"0xb3950edb2a6fe40184baf91c1dc09a4335b5f777",
		"0x6d04860f41f32c79d4225d3737f269f755e2d71a",
		"0x864457950ae362e0833397dd643f973825073309",
		"0xdd9c3cabb1fa2fe25b65cb5786fcae5043843b8a",
		"0xd76ffd62bb92ded12df29e78a9d134a4039792cc",
		"0x27d456a99f1ed54e7d361c0bdc161b8bbf1896fd",
		"0x341ed2d018b8349b72a1c8eb071f6addb9117a8e",
		"0x50664c52d61c5b15442b30295a2e7e17e03bcadc",
		"0x423c48c42c80149ee93d39058d78ad9e424705c9",
		"0xe2e929a7d37874120705253a1c306d5d1a13c30f",
		"0x728f5d4487e83bd64d8c1fd4a1120c0822ee8cef",
		"0x7fe9c4ec2ce68b06324208bb46f82b83a8d172fe",
		"0x99fb83abce0f933ab840f0960ea1dd3008df053b",
		"0xb53b90967626fd52f68371ae195a95c5b6735708",
		"0x1519a2574b3de89b1b7e45d2356a6cb7c589ee04",
		"0x3a8b3bf2cb2e9a25f0edfa1d539cdce131d17bb8",
		"0x97b19c2097845659a5552567712b830d7ed8efc9",
		"0x53375981a78522b578c0435f0ce0d2b021c81f3e",
		"0xc0f771abe4c4fe88b43abe3d2808b2f13201e3ea",
		"0xa684513752aa606fdb61784c8710df8e12308bc9",
		"0x1126798bec8c094bcb56ea00fe1b6ceefce00e5a",
		"0x0e54d44df9fef28533b66635ffbf13cca3fa8e1c",
		"0x8bc4b76445c13d08815750223101932b88910e53",
		"0x615cbe220b015577d1216ea935916a753d0706ee",
		"0xbf195a33899d1ca0c6cda4dacb819a4fea14c15c",
		"0x77646ba7ec9ccabb8412b8f9ba9831e0dce2594c",
		"0xef0bddf659b7a4347849289ca9c73b50476eca29",
		"0x285c885bed6517a7c3f80b56310bfcaaf7cd9d29",
		"0x5f4d92070c9bc3f21eba7caaaa1e88aa58ebc023",
		"0xc49d809b6dda1e6fb3ae8a13e4a3adb4d084d77e",
		"0xd9a38dfba1745895dff6dd78769e9a3e8f19336e",
		"0xf84a1fda23362520e92ba097f0625e4a2ebd7a87",
		"0x07d2261aa76ce43e5eb9eec34ae05085fd740e41",
		"0x5b7365aab9d298a5cae76d31fa76d24bd8c4f6f7",
		"0xfce4246e1b8c2611bb6dd35419172663c9400ded",
		"0x9fe6a25d757f5bd04f1397f055bf62cc5140ea3b",
		"0xa2e2511834a5e8364b65a6c1c6f683b1b554309c",
		"0x39f8d5fbee034c63058364707c7615c195fb376c",
		"0xbf9590102a04693382cf76c50834fa2dd8b740f8",
		"0x2bf2bf65eb36e1099530dca24d820ffc1c3a4dfd",
		"0xd4fd2f14acd9e4a299271ded19b36fe1dc5dca1a",
		"0x51a223eb9bfbae74cd5dc74810e797b2b4e6f5e3",
		"0xb4f83050fb519b0137adcd55baf12fd061dc40eb",
		"0xffc63887bab57bda5dad631ebb44ab6eedeece2d",
		"0x4b32b099c771bcae23e98034d82168727b256a24",
		"0x1a8392c6a2d87994e70998851656166f17771925",
		"0x291d96a37aa219a6c3d34995b8f37509342c46b5",
		"0xd6d868f424ed67c911967fb444ad605c5273f530",
		"0xa959654866530d0f6b8c58a7810c6bf0d1e50508",
		"0x939bc8d050e2b50cbeeaa0db145fea0dd571ff28",
		"0x3016712371a14a03b13f5aa11d6de33ce1c2817a",
		"0x3f19a9a549228b0379117363eb9c1dbce8a856c7",
		"0xc95494b7f5726026b4fcccc18dd20ab77207be49",
		"0xb3827725f9c8047087de2e628b8ecddb7405dc61",
		"0x549a5681534df1378eb91803730978adbed53cf1",
		"0xe10f44dd61302743b4a424b061f18d5a06b0d9b0",
		"0xf44c8f6e762b7634919271086f6617572888107d",
		"0x1141e117a61d4374147bf45dbc79ce3852321683",
		"0xf40be49bc96195bb8a129e5dcd25463f5a6cbb9c",
		"0x4bacfd4761ded2a14c9149a8ed1bc1efed7999b7",
		"0xc01b2babb1445084164a2d7c6c5ca2f54b752186",
		"0x1bf2a8ba2b4fb938ca3a40ac765a38e998d5462b",
		"0x62ef9f9ea09f3e03cc3b35900088ebbec4e39a38",
		"0xf84f8017a220c75b87940b5222bb3388abe2a2ae",
		"0x4201a4ffbdc38b969bb1a16bb973ed85e2dae11b",
		"0x6a1758b73f0f3b11a6a0b51300e2e3cb5837149a",
		"0x27a3739cc1bc4f0bd2028d7c37d049685643a77f",
		"0x622fafe80ac9ef138ec209124fc93824f194f7ae",
		"0x2954da047c2ed2fe7a7222d3a831075da077801d",
		"0x0a56aa95d6ce57054bd0d7e08e4f3e3130dd293c",
		"0x8b9b3ecef7b61e30926e16781b59ea64c9ac724a",
		"0x6c930094f20eb714611b212a9eb9eaee59973422",
		"0x25e95e583fdcf4c9bcf13e01bb85aa941cafd0a4",
		"0xe820e5b31cdce176eefbfc1d893a7c74e7b3977f",
		"0x500ce5d32dc40bb4b63585500b94615907a207bf",
		"0x4ef9f0cb2e2fbfdd52c306ec178cd94312e46797",
		"0xa4b8216a6b2218a19c4499fe18448e08a9b994f3",
		"0x59995d3bf2c295fcf5c85a7ae1b64a86d7039580",
		"0x8ce53404e5b8c7e0eeddf9656a378c808618e45f",
		"0x321a210c019790308abb948360d144e7e00b7dc5",
		"0x6e6916746f5f56ce996ab909bdfc408d7e84b41d",
		"0xc9b7c965d7a5d30850c6dfa7ef87a0c6b08ec218",
		"0x55b2460ca55d123a30279eb4ce6cb5a0cfcbb584",
		"0x9d472e14830e0a239086a05786ff4612964b3b83",
		"0x3ecb7c39cf5d7b885b3e3f9704f6bb20a35db077",
		"0x69d37238330ffcafcba91adba8b86453f115357a",
		"0x2cc79fa3b80c5b9b02051facd02478ea88a78e2c",
	}
)

