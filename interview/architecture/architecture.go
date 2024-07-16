package architecture

// Backends for Frontends pattern
//
// Bu modeli şu durumlarda kullanın:
//
// Paylaşılan veya genel amaçlı bir arka uç hizmetinin, önemli miktarda geliştirme yüküyle birlikte sürdürülmesi gerekir.
// Belirli istemci arayüzlerinin gereksinimleri için arka ucu optimize etmek istiyorsunuz.
// Birden fazla arayüzü barındırmak için genel amaçlı bir arka uçta özelleştirmeler yapılır.
// Bir programlama dili, belirli bir kullanıcı arayüzünün arka ucu için daha uygundur, ancak tüm kullanıcı arayüzleri için uygun değildir.
//
// Bu model uygun olmayabilir:
//
// Arayüzler arka uçtan aynı veya benzer istekleri yaptığında.
// Arka uçla etkileşimde bulunmak için yalnızca bir arayüz kullanıldığında.

// Circuit Breaker
// Bu modeli kullanın:
//
// Bu işlemin başarısız olma ihtimali yüksekse, bir uygulamanın uzak bir hizmeti çağırmaya veya paylaşılan bir kaynağa erişmeye çalışmasını önlemek için.
//
// Bu model önerilmez:
//
// Bellek içi veri yapısı gibi bir uygulamadaki yerel özel kaynaklara erişimi yönetmek için. Bu ortamda devre kesici kullanmak sisteminize ek yük getirecektir.
// Uygulamalarınızın iş mantığındaki istisnaları ele almanın bir alternatifi olarak.
//
//
// CQRS
//
// Aşağıdaki senaryolar için CQRS'yi göz önünde bulundurun:
//
// Birçok kullanıcının aynı verilere paralel olarak eriştiği ortak çalışma alanları. CQRS, etki alanı düzeyinde birleştirme çakışmalarını en aza indirmek için
// yeterli ayrıntı düzeyine sahip komutları tanımlamanıza olanak tanır ve ortaya çıkan çakışmalar, komut tarafından birleştirilebilir.
//
// Kullanıcıların karmaşık bir süreç boyunca bir dizi adım halinde veya karmaşık etki alanı modelleriyle yönlendirildiği görev tabanlı kullanıcı arayüzleri.
// Yazma modeli, iş mantığı, giriş doğrulama ve iş doğrulamayı içeren tam bir komut işleme yığınına sahiptir. Yazma modeli, bir dizi ilişkili nesneyi veri değişiklikleri için
// tek bir birim (DDD terminolojisinde bir toplama) olarak ele alabilir ve bu nesnelerin her zaman tutarlı bir durumda olmasını sağlayabilir. Okuma modelinde iş mantığı veya
// doğrulama yığını yoktur ve yalnızca görünüm modelinde kullanılmak üzere bir DTO döndürür. Okuma modeli sonuçta yazma modeliyle tutarlıdır.
//
// Veri okuma performansının, özellikle okuma sayısı yazma sayısından çok daha fazla olduğunda, veri yazma performansından ayrı olarak ince ayar yapılması gereken senaryolar.
// Bu senaryoda okuma modelinin ölçeğini genişletebilirsiniz ancak yazma modelini yalnızca birkaç örnekte çalıştırabilirsiniz. Az sayıda yazma modeli örneği,
// birleştirme çakışmalarının oluşumunun en aza indirilmesine de yardımcı olur.
//
// Bir geliştirici ekibinin yazma modelinin parçası olan karmaşık etki alanı modeline, diğer ekibin ise okuma modeline ve kullanıcı arayüzlerine odaklanabileceği senaryolar.
//
// Sistemin zaman içinde gelişmesinin beklendiği ve modelin birden fazla versiyonunu içerebileceği veya iş kurallarının düzenli olarak değiştiği senaryolar.
//
// Diğer sistemlerle entegrasyon, özellikle bir alt sistemin geçici arızasının diğerlerinin kullanılabilirliğini etkilememesi gereken olay kaynağı kullanımıyla birlikte.
//
// Bu model şu durumlarda önerilmez:
//
// Etki alanı veya iş kuralları basittir.
//
// Basit CRUD tarzı bir kullanıcı arayüzü ve veri erişim işlemleri yeterlidir.
//
// CQRS'yi sisteminizin en değerli olacağı sınırlı bölümlerine uygulamayı düşünün.
//
//
// Saga
//
// Aşağıdakileri yapmanız gerektiğinde Saga modelini kullanın:
//
// Dağıtılmış bir sistemde sıkı bağlantı olmadan veri tutarlılığı sağlayın.
// Sıradaki işlemlerden biri başarısız olursa geri alın veya telafi edin.
//
// Saga modeli aşağıdakiler için daha az uygundur:
//
// Sıkıca bağlanmış işlemler.
// Daha önceki katılımcılarda meydana gelen telafi edici işlemler.
// Döngüsel bağımlılıklar.
//
// Event-Driven
//
// Birden fazla alt sistemin aynı olayları işlemesi gerekir.
// Minimum zaman gecikmesiyle gerçek zamanlı işleme.
// Zaman pencerelerinde desen eşleştirme veya toplama gibi karmaşık olay işleme.
// Nesnelerin İnterneti gibi yüksek hacimli ve yüksek hızlı veriler.
//
// Faydalar
//
// Üreticiler ve tüketiciler birbirinden ayrılmıştır.
// Noktadan noktaya entegrasyon yok. Sisteme yeni tüketiciler eklemek kolaydır.
// Tüketiciler olaylara anında yanıt verebilirler.
// Yüksek düzeyde ölçeklenebilir ve dağıtılabilir.
// Alt sistemler olay akışının bağımsız görünümlerine sahiptir.
//
// Zorluklar
//
// Garantili Teslimat. Bazı sistemlerde, özellikle IoT senaryolarında, olayların iletildiğini garanti etmek çok önemlidir.
// Olayların sırayla veya tam olarak bir kez işlenmesi. Her tüketici türü, esneklik ve ölçeklenebilirlik için genellikle birden çok örnekte çalışır.
// Olayların sırayla işlenmesi gerekiyorsa (tüketici türü dahilinde) veya bağımsız mesaj işleme mantığı uygulanmazsa bu durum bir zorluk yaratabilir.
// Hizmetler genelinde mesajların koordine edilmesi. İş süreçleri genellikle tüm iş yükünde tutarlı bir sonuç elde etmek için birden fazla hizmetin
// mesaj yayınlamasını ve mesajlara abone olmasını içerir. Koreografi modeli ve Saga Düzenleme gibi iş akışı modelleri, çeşitli hizmetler genelinde
// mesaj akışlarını güvenilir bir şekilde yönetmek için kullanılabilir.
//
// n-tier
// N katmanlı mimariler genellikle hizmet olarak altyapı (IaaS) uygulamaları olarak uygulanır ve her katman ayrı bir VM kümesinde çalışır. Ancak N katmanlı bir uygulamanın saf IaaS olması gerekmez. Çoğunlukla, mimarinin bazı bölümleri için (özellikle önbellekleme, mesajlaşma ve veri depolama) yönetilen hizmetleri kullanmak avantajlıdır.
//
// Aşağıdakiler için N katmanlı bir mimari düşünün:
//
// Basit web uygulamaları.
// Mimari gereksinimlerin henüz net olmadığı durumlarda iyi bir başlangıç ​​noktası.
// Şirket içi bir uygulamayı minimum düzeyde yeniden düzenlemeyle Azure'a geçirme.
// Şirket içi ve bulut uygulamalarının birleşik geliştirilmesi.
//
// N katmanlı mimariler, geleneksel şirket içi uygulamalarda çok yaygındır, dolayısıyla mevcut iş yüklerini Azure'a geçirmek için doğal bir uyum sağlar.
//
// Event-Sourcing
// Bu modeli aşağıdaki senaryolarda kullanın:
//
// Verilerdeki amacı, amacı veya nedeni yakalamak istediğinizde. Örneğin, bir müşteri varlığındaki değişiklikler, eve taşınma ,
// hesap kapanma veya vefat etme gibi bir dizi belirli olay türü olarak yakalanabilir .
//
// Verilerde çakışan güncellemelerin oluşmasını en aza indirmek veya tamamen önlemek hayati önem taşıdığında.
//
// Meydana gelen olayları kaydetmek, sistemin durumunu geri yüklemek, değişiklikleri geri almak veya geçmiş ve denetim günlüğü tutmak için
// bunları yeniden oynatmak istediğinizde. Örneğin, bir görev birden fazla adım içerdiğinde, güncellemeleri geri almak için eylemler gerçekleştirmeniz
// ve ardından verileri tekrar tutarlı bir duruma getirmek için bazı adımları tekrarlamanız gerekebilir.
//
// Olayları kullandığınızda. Bu, uygulamanın işleyişinin doğal bir özelliğidir ve çok az ekstra geliştirme veya uygulama çabası gerektirir.
//
// Veri girme veya güncelleme sürecini bu eylemleri uygulamak için gereken görevlerden ayırmanız gerektiğinde. Bu değişiklik,
// kullanıcı arayüzü performansını iyileştirmek veya olayları, olaylar meydana geldiğinde harekete geçecek diğer dinleyicilere dağıtmak olabilir.
// Örneğin, bir bordro sistemini gider gönderimi web sitesine entegre edebilirsiniz. Web sitesinde yapılan veri güncellemelerine yanıt olarak etkinlik
// mağazası tarafından gündeme getirilen etkinlikler, hem web sitesi hem de bordro sistemi tarafından tüketilecektir.
//
// Gereksinimlerin değişmesi durumunda materyalleştirilmiş modellerin ve varlık verilerinin formatını değiştirebilmek için esneklik istediğinizde veya
// CQRS ile birlikte kullanıldığında, verileri ortaya çıkaran bir okuma modelini veya görünümleri uyarlamanız gerekir.
//
// CQRS ile birlikte kullanıldığında, bir okuma modeli güncellenirken nihai tutarlılık kabul edilebilir veya bir olay akışından alınan varlıkların ve
// verilerin yeniden nemlendirilmesinin performans etkisi kabul edilebilir.
//
// Bu model aşağıdaki durumlarda kullanışlı olmayabilir:
//
// Küçük veya basit etki alanları, çok az iş mantığına sahip olan veya hiç olmayan sistemler veya geleneksel CRUD veri yönetimi mekanizmalarıyla doğal olarak iyi çalışan etki alanı olmayan sistemler.
//
// Veri görünümlerinde tutarlılığın ve gerçek zamanlı güncellemelerin gerekli olduğu sistemler.
//
// Denetim izlerinin, geçmişin ve geri alma ve yeniden yürütme eylemlerinin gerekli olmadığı sistemler.
//
// Temel verilerde çakışan güncellemelerin yalnızca düşük oranda görüldüğü sistemler. Örneğin, ağırlıklı olarak verileri güncellemek yerine ekleyen sistemler.
//
