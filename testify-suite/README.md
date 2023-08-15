# testify-suite

Use https://pkg.go.dev/github.com/stretchr/testify/suite and test its setup/teardown functions.

## Usage

Command:

```shell
go test
```

Output:

```shell
SetupSuite()
SetupTest()
BeforeTest(ExampleTestSuite, TestOne)
TestOne()
AfterTest(ExampleTestSuite, TestOne)
TearDownTest()
SetupTest()
BeforeTest(ExampleTestSuite, TestThree)
TestThree()
AfterTest(ExampleTestSuite, TestThree)
TearDownTest()
SetupTest()
BeforeTest(ExampleTestSuite, TestTwo)
TestTwo()
AfterTest(ExampleTestSuite, TestTwo)
TearDownTest()
TearDownSuite()
PASS
ok      testify_suite   0.181s
```
