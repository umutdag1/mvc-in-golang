# MVC Framework
## _HandMade New Generated Framework_

HandMade Framework MVC Design Pattern'ü kullanılarak Controller ve Model barındıran Rest-API Service için yazılmış özel bir frameworktür

- Controller
- Model

## Özellikler

- Yeni EndPoint Tanımlayabilme
- EndPoint Icin Olusturulan Handler Fonksiyonu Controllerda Tanımlı Olup Olmadığını ya da Controller'a ait olup olmadığının otomatik olarak kontrolü
- InMemory Olarak Geçici Olarak Set Edilen Dataları Key-Value Olarak Tutabilme
- Belli Süre Aralıklarıyla Programın Açılışından İtibaren InMemory'de Tutulan Dataları JSON Olarak Dosyaya Kayıt Etme Özelliği
- Elle Yazılmış Birkaç Kütüphane (JSONHANDLER-FILEHANDLER-FOLDERHANDLER-LOGHANDLER)

## Controller

Her MVC Framework'te olduğu gibi HandMade MVC Framework'tede  aynı Controller yapısı bulunmaktadır. Controller eğer bir istek çağırılmışsa, o isteğin handler'ının tanımlandığı kısım Controller'da bulunur. Her Controller'ın kendine ait struct tipinde bir modül tanımlanması gerekmektedir. Böylelikle Controllerda bir dosyanın içerisinde bir endpointe ait fonksiyon tanımlanmışsa ve o dosyaya ait bir modül tanımlanmışsa, modül bulunduğu yerdeki tanımlanan fonksiyonla aynı yerdeyse, program çalıştırıldığında sorunsuz olarak server çalıştırılacaktır. Eşleşmez ise program hatayı basar ve server'i başlatmadan sonlandırır.

## Model
Model sadece database ya da database gibi davranan teknolojiler ile veri alışverişi yada CRUD (CREATE-READ-UPDATE-DELETE) işlemleri gibi istenilen komuta göre hata var ise hata, hata olmadığı durumda istenilen veriye yada yapılmak istenen işleme başarıyla gerçekleştirilir.

 Projenin kullanımının açıklamasından sonra bunu daha iyi görebilmek için birkaç endpoint Controller içerisinde "data.go" adlı dosya içerisine eklendi. Bunlar :

## API

| EndPoint | Method | Kullanım |
| ------ | ------ | ------ |
| /api/getAll | GET | Parametre Yok. |
| /api/get | GET  | /api/get?key=<key> |
| /api/set | POST | JSONBody : {key : <key>, value : <value> }  |

## İndirme

HandMade MVC Framework  Golang Programlama Dili Ile Geliştirilmiş Bir Rest Api Frameworktür. Bundan dolayı Golang Compiler v1.16 ve üzeri ile çalıştırılması gerekmektedir.

Golang Compiler indirildikten sonra, github üzerinden "https://github.com/umutdag1/yemeksepeti-odev.git" bağlantısını kullanarak "Command Prompt" dan aşağıdaki kod satırları sırasıyla uygulanır.

```sh
git clone https://github.com/umutdag1/yemeksepeti-odev.git
cd yemeksepeti-odev
go run main.go
```

## Bir Endpoint'e Nasıl İstek Yapılır ?

Browser üzerinden GET istekleri yapılabilir POST için bu söz konusu değildir. Bütün istekleri doğru yapmak ve geri dönüşlerini kontrol edebilmek için POSTMAN uygulamasının kullanılması daha uygundur.

POSTMAN Nasıl Kullanılır ?
"https://learning.postman.com/docs/publishing-your-api/documenting-your-api/" urlsinden POSTMAN dökümantasyonuna ulaşabilirsiniz


## License

MIT License

Copyright (c) 2021 umutdag1

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.