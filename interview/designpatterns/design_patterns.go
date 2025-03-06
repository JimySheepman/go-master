// 1. Yaratımsal (Creational) Tasarım Kalıpları
//
// Singleton: Bir sınıfın yalnızca bir örneğinin oluşturulmasını sağlar ve bu örneğe global erişim noktası sunar.
// Factory Method: Bir sınıfın alt sınıflarının hangi nesneyi oluşturacağına karar vermesine izin verir.
// Abstract Factory: İlgili nesnelerin ailesini oluşturmak için bir arayüz sağlar, ancak alt sınıflar oluşturma işlemini belirler.
// Builder: Kompleks bir nesnenin oluşturulma sürecini adım adım yöneten bir nesne oluşturur.
// Prototype: Bir nesnenin mevcut bir örneğini klonlayarak yeni nesneler oluşturur.
//
// 2. Yapısal (Structural) Tasarım Kalıpları
//
// Adapter: Bir sınıfın arayüzünü, beklenen başka bir arayüze çevirir.
// Bridge: Bir arayüz ve onun implementasyonları arasındaki bağımsızlığı sağlar.
// Composite: Nesneleri ağaç yapısında organize ederek bireysel nesneler ve nesne gruplarını aynı şekilde işlemeyi sağlar.
// Decorator: Nesnelere dinamik olarak yeni davranışlar ekler.
// Facade: Bir sistemin alt sistemlerinin daha basit bir arayüzle kullanılmasını sağlar.
// Flyweight: Çok sayıda küçük nesneye sahip uygulamalarda bellek kullanımını azaltmak için nesne paylaşımını sağlar.
// Proxy: Bir nesneye erişimi kontrol eden bir aracı sağlar.
//
// 3. Davranışsal (Behavioral) Tasarım Kalıpları
//
// Chain of Responsibility: Bir istek üzerindeki işlemi zincirleme nesnelerle gerçekleştirir.
// Command: İşlemleri nesne olarak kapsüller ve işlemleri geri alabilir.
// Interpreter: Bir dilin sözdizimini tanımlar ve bu sözdizimlerini yorumlar.
// Iterator: Bir koleksiyonun elemanlarına ardışık olarak erişim sağlar.
// Mediator: Nesneler arasındaki etkileşimleri merkezi bir şekilde yönetir.
// Memento: Bir nesnenin iç durumunu kaydederek daha sonra bu durumu geri yüklemeyi sağlar.
// Observer: Bir nesnedeki değişiklikleri başka nesnelere bildirir.
// State: Bir nesnenin iç durumunu değiştirdiğinde davranışını da değiştirir.
// Strategy: Bir algoritmayı bir nesne içinde kapsüller ve algoritmanın seçimini çalışma zamanında yapmayı sağlar.
// Template Method: Bir işlemin iskeletini tanımlar ve adımlarının alt sınıflar tarafından geçersiz kılınmasını sağlar.
// Visitor: Bir nesne yapısının elemanları üzerinde işlem yapmak için yeni bir işlem ekler.

package designpatterns

import (
	"fmt"
	"github.com/JimySheepman/go-master/go-algorithm/designpatterns/behavioral"
	"github.com/JimySheepman/go-master/go-algorithm/designpatterns/creational"
	"github.com/JimySheepman/go-master/go-algorithm/designpatterns/structural"
)

type designPattern func()

var designPatterns = []designPattern{
	creational.Singleton,
	creational.FactoryMethod,
	creational.AbstractFactory,
	creational.Builder,
	creational.Prototype,

	structural.Adapter,
	structural.Bridge,
	structural.Composite,
	structural.Decorator,
	structural.Facade,
	structural.Flyweight,
	structural.Proxy,

	behavioral.ChainOfResponsibility,
	behavioral.Command,
	behavioral.Iterator,
	behavioral.Mediator,
	behavioral.Memento,
	behavioral.Observer,
	behavioral.State,
	behavioral.Strategy,
	behavioral.TemplateMethod,
	behavioral.Visitor,
}

var designPatternsName = map[int]string{
	0: "Creational Pattern Singleton",
	1: "Creational Pattern Factory Method",
	2: "Creational Pattern Abstract Factory",
	3: "Creational Pattern Builder",
	4: "Creational Pattern Prototype",

	5:  "Structural Pattern Adapter",
	6:  "Structural Pattern Bridge",
	7:  "Structural Pattern Composite",
	8:  "Structural Pattern Decorator",
	9:  "Structural Pattern Facade",
	10: "Structural Pattern Flyweight",
	11: "Structural Pattern Proxy",

	12: "Behavioral Pattern Chain of Responsibility",
	13: "Behavioral Pattern Command",
	14: "Behavioral Pattern Iterator",
	15: "Behavioral Pattern Mediator",
	16: "Behavioral Pattern Memento",
	17: "Behavioral Pattern Observer",
	18: "Behavioral Pattern State",
	19: "Behavioral Pattern Strategy",
	20: "Behavioral Pattern Template Method",
	21: "Behavioral Pattern Visitor",
}

func PrintDesignPatterns() {
	for i, pattern := range designPatterns {
		fmt.Println("Design Pattern Name:", designPatternsName[i])
		pattern()
	}
}
