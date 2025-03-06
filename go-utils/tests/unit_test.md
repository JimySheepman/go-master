# Unit Test Guide

## Contents

- [Mock Generate](#mock-generate)
- [Create Local Mock](#create-local-mock)
- [Write a Setup Method](#write-a-setup-method)
- [Table-Driven Unit Test Examples](#table-driven-unit-test-examples)
  - [Explanation](#explanation)
  - [Panic Handler Test Example](#panic-handler-test-example)
  - [GRPC tag test Example](#grpc-tag-test-example)
  - [Async Method Test Example](#async-method-test-example)
  - [Set Args Test Example](#set-args-test-example)

## Mock Generate

_İlk olarak mocklarımızı oluşturmamız gerekiyor._

```Go
// Açıklama;
//go:generate mockgen -package=<oluşturulan dosyanın package name'i> -destination=<oluşturrulacak dosya yolu ve ismi> -source=<referans aldığı dosya> < proje modül path'i>

// Örnek;
//go:generate mockgen -package=mock_services -destination=./mocks/address_mock.go  -source=address.go example-project/services
```

## Create Local Mock

_Ikinci aşamada ise testleri yazmadan önce her zaman cağıracağımız ve ortak kullanılacak yapıları bir araya toplayıp kullanıma harı hale getirmemiz gerek._

```Go
// Açıklama;
// mock struct; mocklanacak olan struct ile aynı interfacelerin mockları yazılır
type local<yazdıpımız struct'ın adı>Mocks struct {
 <structun içerisindeki interfacelerin yazamız gerek. Başına 'mock' koyarak> <mock dosyamızdaki interface'i implement eden uygun struct yapısını vermemiz gerek>
}

// Örnek;
// mocklanan struct
type testService struct {
 ctrl         testController
 ctxExtractor contextExtract
}

// mock struct
type localtestServiceMocks struct {
 mockCtrl         *mock_services.MocktestController
 mockCtxExtractor *mock_services.MockcontextExtract
}
```

## Write a Setup Method

_Burada ise mock ve gerekli service oluşturulmalı ve bu setup dosyası ilgili testler için çağrılmalıdır._

```Go
// Açıklama;
// setup dosyası; mocks context ve test edilen struct oluşturulur.
// mocklar burada injected edilir.
func _setup<yazdıpımız struct ın adı>Test_(t *testing.T) (<mock repo>, <test edilen struct>,<opsiyonel olarak context>) {
 ctrl, ctx := gomock.WithContext(context.Background(), t) // mock nesneleri oluşturmak için gerekli

 mocks := <mock repo tanımmlanmalı ve mock objeleri oluştrulmalı>

 srv := <new lenen service dönülmeli>

 return  mocks, srv,ctx
}

// Örnek;
func _setuptestServiceTest_(t *testing.T) (context.Context, *localtestServiceMocks, *testService) {
 ctrl, ctx := gomock.WithContext(context.Background(), t)

 mocks := &localtestServiceMocks{
  mockCtrl:         mock_services.NewMocktestController(ctrl),
  mockCtxExtractor: mock_services.NewMockcontextExtract(ctrl),
 }

 repo := NewtestService(mocks.mockLoc, mocks.mockCtrl, mocks.mockCtxExtractor)

 return ctx, mocks, repo
}
```

## Table-Driven Unit Test Examples

_Son olarak artık test yazmaya hazırız test-table şeklinde senaryolarımızı kurgulayıp kendi testlerimizi yazabileceğimiz bir yapı kurgulamaıyız. Bu yapı hata case'lerine göre gitmelidir._

### Explanation

```Go
// genel bir test table test yapısıdır. Array struct içerisi
// gerekli parametreler ile değişmeli ve gerekli require lar
// ile kontrol edilmeli

func Test<yazdıpımız struct ın adı>_<test edilen method adı>(t *testing.T) {
 ctx, mocks, srv := _setupTransferAdminServiceTest_(t) // setup dosyamızı cağrıyoruz.

 tests := []struct {
        // test tablomuzun field larını yazıyoruz ihtiyaca göre çoğaltılabilinir
  name         string
  ctx          context.Context
  req          *admin.ListTransfersDownloadRequest
  expected     *admin.ListTransfersDownloadResponse
  isError      bool
  expectations func()
 }{
  {// burada ise yeni bir test case yazıyoruz
   name:     "db.NamedQueryContext failure",
   ctx:      ctx,
   req:      nil,
   expected: nil,
   isError:  true,
   expectations: func() { // gerekir ise mock reposundan method çağrısı yapıyoruz.
    //mocks.mockCtrl.EXPECT().(gomock.Any(),gomock.Any()).Return(nil,errors.New("creation new object error"))
   },
  },
 }

 for _, tt := range tests {
  t.Run(tt.name, func(t *testing.T) {
   tt.expectations() // mock method lar çağrılıyor 

   actual, err := srv.<method adı>(tt.ctx, tt.req) // method çağrılır

    // bu krılım testin hata alıp almamasına göre ayarlanır.
    // buradaki assert içerikleri test'e göre değişir.
 
   if tt.isError { 
    assert.Error(t, err)
   } else {
    assert.Nil(t, err)
   }
    assert.Equal(t, tt.expected, actual)
  })
 }
}
```

### Panic Handler Test Example

```Go
func TestCompleter_NewCompleter(t *testing.T) {
 ctx := context.Background()
 ctrl, _ := gomock.WithContext(ctx, t)

 tests := []struct {
  testName        string
  condition       interface{}
  testLister   testLister
 }{
  {
   testName:        "test should be testLister nil and throwing panic",
   condition:       nil,
   testLister:   nil,
  },
 }

 for _, tt := range tests {
  t.Run(tt.testName, func(t *testing.T) {
   defer func() {
    if r := recover(); r == tt.condition {
     t.Errorf("The code did not  panic")
    }
   }()

   s := NewCompleter(tt.testLister, tt.documentLister, tt.deliveryManager)
   _ = s
  })
 }
}
```

### GRPC tag test Example

```Go
func Test_TestExample(t *testing.T) {
 ctx, mocks, srv := _setupTransfersUserTest_(t)

 type Tags struct {
  setTestKey   string
  setTestValue int64

 }

 tests := []struct {
  name         string
  ctx          context.Context
  req          *twcollector.CheckTransferRequirementsRequest
  tags         Tags
  expected     *twcollector.CheckTransferRequirementsResponse
  isError      bool
  expectations func()
 }{
  {
   name: "grpc_ctxtags.Extract get testID failure",
   ctx:  ctx,
   req: &twcollector.CheckTransferRequirementsRequest{
    UserId: 0,
   },
   tags: Tags{
    setTestKey:   "",
    setTestValue: 0,
   },
   expected:     nil,
   isError:      true,
   expectations: func() {},
  },
  {
   name: "succeed",
   ctx:  ctx,
   req: &twcollector.CheckTransferRequirementsRequest{
    UserId: 0,
   },
   tags: Tags{
    setTestKey:   "testID",
    setTestValue: 0,
   },
   expected:     nil,
   isError:      true,
   expectations: func() {},
  },
 }

 for _, tt := range tests {
  t.Run(tt.name, func(t *testing.T) {
   tt.expectations()

   tag := grpc_ctxtags.NewTags()
   tag.Set(tt.tags.setTestKey, tt.tags.setTestValue)
   tt.ctx = grpc_ctxtags.SetInContext(tt.ctx, tag)

   actual, err := srv.TestExample(tt.ctx, tt.req)
   if tt.isError {
    assert.Error(t, err)
   } else {
    assert.Nil(t, err)
   }
   assert.Equal(t, tt.expected, actual)
  })
 }
}
```

### Async Method Test Example

```Go
func Test_TestExample(t *testing.T) {
 _, mocks, srv := _setupTransferJobTest_(t)

 wg := sync.WaitGroup{}

 tests := []struct {
  name         string
  req          *database.ExampleTransfer
  expectations func()
 }{
  {
   name: "complete Transfer succeed",
   req: &database.ExampleTransfer{},
   expectations: func() {
    wg.Add(1)
    mocks.mockAccountDb.EXPECT().GetTestByIBAN(gomock.Any(), gomock.Any()).Return(&database.ExampleAccount{}, nil)
    mocks.mockTransferAdminController.EXPECT().TestTransfer(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return("", "", nil)
    mocks.mockTransferLogRepo.EXPECT().LogOperation(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
    mocks.mockSlackNotifierClient.EXPECT().NotifySlack(gomock.Any()).DoAndReturn(func(t *database.ExampleTransfer) {
     wg.Done()
    })

    mocks.mockFormicaFraudClient.EXPECT().SendTestAsync(gomock.Any(), gomock.Any())
    mocks.mockScannerServiceClient.EXPECT().UpdateTestAsync(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any())
   },
  },
 }

 for _, tt := range tests {
  t.Run(tt.name, func(t *testing.T) {
   tt.expectations()

   srv.TestExample(tt.req)
   wg.Wait()
  })
 }
}
```

### Set Args Test Example

```Go
func Test_TestExample(t *testing.T) {
 mocks, app := _setupApplicationTest_(t)

 tests := []struct {
  name         string
  expected       []*ExampleApplication
  isError      bool
  expectations func()
 }{
  {
   name:    "Select creation new object failure",
   expected:  nil,
   isError: true,
   expectations: func() {
    mocks.mockDb.EXPECT().Select(gomock.Any(), gomock.Any()).Return(nil, errors.New("creation new object error"))
   },
  },
  {
   name:    "Select creation new object succeed",
   expected:  []*ExampleApplication{{}},
   isError: false,
   expectations: func() {
    apps := []*ExampleApplication{{}}
    mocks.mockDb.EXPECT().Select(gomock.Any(), gomock.Any()).SetArg(0, apps).Return(nil, nil)
   },
  },
 }

 for _, tt := range tests {
  t.Run(tt.name, func(t *testing.T) {
   tt.expectations()

   actual, err := app.FindApplications()
   if tt.isError {
    assert.Error(t, err)
   } else {
    assert.Nil(t, err)
   }
    assert.Equal(t, tt.expected, actual)
  })
 }
}
```
