package twelvefactorapp

// Twelve-Factor App, modern, ölçeklenebilir, taşınabilir web uygulamaları geliştirmek için yaygın olarak kabul edilen bir metodolojidir. 2011 yılında Heroku platformunun kurucularından Adam Wiggins tarafından ortaya atılmıştır. Twelve-Factor App, uygulamaların mikro hizmet mimarisine uygun olarak geliştirilmesini ve dağıtılmasını kolaylaştırır. İşte bu metodolojinin 12 faktörü:
//
// Codebase (Kod Tabanı)
// Tüm uygulamalar tek bir kod tabanına dayanmalı ve bu kod tabanı bir sürüm kontrol sisteminde (örn. Git) yönetilmelidir.
// Kod tabanı, birden fazla dağıtım olabilir, ancak her dağıtım aynı kod tabanını kullanmalıdır.
//
// Dependencies (Bağımlılıklar)
// Uygulamanın bağımlılıkları açıkça belirtilmeli ve uygulamadan izole edilmelidir.
// Bağımlılıklar, paket yöneticileri (örn. Maven, NPM) kullanılarak yönetilmelidir.
//
// Config (Konfigürasyon)
// Konfigürasyon, koddan ayrılmalıdır.
// Konfigürasyon, genellikle ortam değişkenleri kullanılarak yönetilmelidir.
//
// Backing Services (Destek Hizmetleri)
// Uygulama, veri tabanları, mesajlaşma sistemleri gibi destek hizmetleri ile bağlantılı olmalı ve bu hizmetler bir URL veya başka bir konfigürasyon yöntemiyle tanımlanmalıdır.
//
// Build, Release, Run
// Uygulama, üç ayrı aşamada çalışmalıdır: build (inşa), release (yayınlama) ve run (çalıştırma).
// Build aşaması, kodu ve bağımlılıkları bir araya getirir. Release aşaması, build çıktısını ve konfigürasyonu bir araya getirir. Run aşaması, release'ı çalıştırır.
//
// Processes (Süreçler)
// Uygulama, stateless (durumsuz) ve share-nothing (paylaşım yapmayan) süreçler olarak çalışmalıdır.
// Durum bilgisi, genellikle veri tabanları veya diğer destek hizmetleri kullanılarak yönetilmelidir.
//
// Port Binding (Port Bağlama)
// Uygulama, kendi web sunucusunu başlatarak bir port üzerinde çalışmalıdır.
// Uygulama, port üzerinde HTTP isteklerini kabul etmelidir.
//
// Concurrency (Eşzamanlılık)
// Uygulama, süreçleri küçük ve izole tutulmalı ve yatay olarak ölçeklenebilir olmalıdır.
// Uygulama, farklı süreç türlerini işleyebilmelidir.
//
// Disposability (Atılabilirlik)
// Uygulama süreçleri hızlı bir şekilde başlatılmalı ve durdurulmalıdır.
// Süreçler, sağlam bir şekilde kapanabilmeli ve gerektiğinde yeniden başlatılabilmelidir.
//
// Dev/Prod Parity (Geliştirme/Üretim Eşitliği)
// Geliştirme, test ve üretim ortamları mümkün olduğunca benzer olmalıdır.
// Kod, zaman ve kişiler arasındaki farklılıklar minimize edilmelidir.
//
// Logs (Loglar)
// Uygulama, logları olay akışı olarak ele almalı ve standard output (standart çıktı) aracılığıyla kaydetmelidir.
// Log yönetimi, dış hizmetler (örn. log toplama ve analiz araçları) tarafından yapılmalıdır.
//
// Admin Processes (Yönetim Süreçleri)
// Yönetim görevleri (örn. veritabanı migrasyonları, bakım işleri) uygulamanın normal süreçlerinden ayrı olarak çalıştırılmalıdır.
// Yönetim süreçleri, bir kerelik çalıştırılabilir komutlar olarak ele alınmalıdır.
//
// Twelve-Factor metodolojisi, özellikle bulut tabanlı uygulamalar ve mikro hizmetler için ideal bir çerçeve sağlar. Uygulamaların taşınabilirliğini, ölçeklenebilirliğini ve bakımını kolaylaştırır.
