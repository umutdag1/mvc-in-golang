# MVC Framework
## _HandMade New Generated Framework_

HandMade New Generated Framework is a new coded-framework for creating Rest-API Service easily using an architecture that is called as either MVC Design Pattern or Architectural Pattern, which consists of both Controller and Model.

(HandMade New Generated Framework MVC Design Pattern'ü kullanılarak Controller ve Model barındıran Rest-API Service için yazılmış yeni bir frameworktür.)

- Controller
- Model

## Features (Özellikler)

- Writing Your Own EndPoint (Yeni EndPoint Tanımlayabilme)
- Auto-mechanism For Checking Which Function Belongs to Which Controller (EndPoint Icin Olusturulan Handler Fonksiyonu Controllerda Tanımlı Olup Olmadığını ya da Controller'a ait olup olmadığının otomatik olarak kontrolü)
- Including an Example That Holds The Temporary Entered Key-Value Pair in a JSON File Automatically (Otomatik Olarak Geçici Key-Value ikilisi İçeren Data Tutan Bir Örnek İçermesi)
- A Few Written Library Examples <JSONHANDLER-FILEHANDLER-FOLDERHANDLER-LOGHANDLER> (Birkaç Örnek Olarak Yazılmış Kütüphane <JSONHANDLER-FILEHANDLER-FOLDERHANDLER-LOGHANDLER>)

## Kütüphanelerin Özellikleri (Libraries Features)
- JSONHANDLER : Includes JSON Handling Mechanism

  (JSONHANDLER : JSON ile ilgili olaylarla ilgilenen yapılar barındırmaktadır.)

- FILEHANDLER : Includes File Handling Mechanism

  (FILEHANDLER : Dosya ile ilgili olaylarla ilgilenen yapılar barındırmaktadır.)

- FOLDERHANDLER : Includes Folder and Files Under The Folder Handling Mechanism

  (FOLDERHANDLER : Dosya Klasörü ve altında yer alan Dosyalar ile ilgili olaylarla ilgilenen yapılar barındırmaktadır)

- LOGHANDLER : Includes Log-Handling Mechanism Such as INFO which gives information, WARNING which gives warning information, ERROR which gives error information

  (LOGHANDLER : Bilgilendirme, Uyarı ya da Hata mesajlarının nasıl ekrana bastırılacağı ile ilgili yapıları barındıran (Handle) yani kendimize göre nasıl bastırılacağını ayarlamamız sağlayan kütüphanedir.)

## Controller

- Like all other MVC Frameworks, HandMade MVC Framework includes both Controller and Model machanisms. Controller is used for handling HTTP Request called from client side. All Controller you've created must include a structure of that Controller module itself. So, if a request is done to a function of which is written in Controller you've created, that function will be fired and automatically execute statements what you've typed in because of Controller module structure. It's done for all your Controller's functions. If the function you've requested is not created, then the application is ended with throwing an error.

 (Her MVC Framework'te olduğu gibi HandMade MVC Framework'tede  aynı Controller yapısı bulunmaktadır. Controller eğer bir istek çağırılmışsa, o isteğin handler'ının tanımlandığı kısım Controller'da bulunur. Her Controller'ın kendine ait struct tipinde bir modül tanımlanması gerekmektedir. Böylelikle Controllerda bir dosyanın içerisinde bir endpointe ait fonksiyon tanımlanmışsa ve o dosyaya ait bir modül tanımlanmışsa, modül bulunduğu yerdeki tanımlanan fonksiyonla aynı yerdeyse, program çalıştırıldığında sorunsuz olarak server çalıştırılacaktır. Eşleşmez ise program hatayı basar ve server'i başlatmadan sonlandırır.)

## Model
- Model is responsible for making connection with or dealing with data on db by doing CRUD (CREATE-READ-UPDATE-DELETE) transaction. Depending on db behaviours, it will produce a result that is whether successful or failed for Controler, from where a request was done  
 (Model sadece database ya da database gibi davranan teknolojiler ile veri alışverişi yada CRUD (CREATE-READ-UPDATE-DELETE) işlemleri gibi istenilen komuta göre hata var ise hata, hata olmadığı durumda istenilen veriye yada yapılmak istenen işleme başarıyla gerçekleştirilir.)

## API Example (API Örnek)
| EndPoint | Method | Kullanım | Özellik |
| ------ | ------ | ------ | ------ |
| /api/getAll | GET | No Parameters | Getting All Data from InMemory |
| /api/get | GET  | /api/get?key="key" | Getting Data Depending On Specified "key" |
| /api/set | POST | JSONBody : {key : "key", value : "value" }  | A "value" is stored on Inmemory with its assigned "key" |
| /api/flushAll | DELETE | No Parameters  | Flush All Data On InMemory |


| EndPoint | Method | Kullanım | Özellik |
| ------ | ------ | ------ | ------ |
| /api/getAll | GET | Parametre Yok. | InMemory'deki bütün dataları alır |
| /api/get | GET  | /api/get?key="key" | Verilen "key"'e göre datayı getirir |
| /api/set | POST | JSONBody : {key : "key", value : "value" }  | Verilen "key" ve karşılığı olan "value" InMemory'e kayıt edilir |
| /api/flushAll | DELETE | Parametre Yok  | InMemory'deki bütün dataları sıfırlar |

## Download - (İndirme)
- It's designed on Golang within Golang Compiler v1.16. So, it runs on v1.16 or releases.

 (HandMade MVC Framework  Golang Programlama Dili Ile Geliştirilmiş Bir Rest Api Frameworktür. Bundan dolayı Golang Compiler v1.16 ve üzeri ile çalıştırılması gerekmektedir.)

- You can clone the application by running all commands on your local "Command Prompt" below.

 (Golang Compiler indirildikten sonra, github üzerinden "https://github.com/umutdag1/yemeksepeti-odev.git" bağlantısını kullanarak "Command Prompt" dan aşağıdaki kod satırları sırasıyla uygulanır.)

```sh
git clone https://github.com/umutdag1/mvc-in-golang.git
cd mvc-in-golang
go run main.go
```

## How to make a request to an endpoint ? - (Bir Endpoint'e Nasıl İstek Yapılır ?)

- GET request can be handled except POST by the browser. So, you should use a tool to test your endpoints by requesting them. POSTMAN is very powerful tool to make them all easily.

 (Browser üzerinden GET istekleri yapılabilir POST için bu söz konusu değildir. Bütün istekleri doğru yapmak ve geri dönüşlerini kontrol edebilmek için POSTMAN uygulamasının kullanılması daha uygundur.)

How to use POSTMAN ? - (POSTMAN Nasıl Kullanılır ?)
- "https://learning.postman.com/docs/publishing-your-api/documenting-your-api/" you can reach POSTMAN document by browsing the url above.

("https://learning.postman.com/docs/publishing-your-api/documenting-your-api/" urlsinden POSTMAN dökümantasyonuna ulaşabilirsiniz)


## License - (Lisans)

[MIT](https://github.com/umutdag1/mvc-in-golang/blob/main/LICENSE) License