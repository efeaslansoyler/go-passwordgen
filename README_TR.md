# Go Parola Üreteci

[![en](https://img.shields.io/badge/lang-en-red.svg)](README.md)
[![tr](https://img.shields.io/badge/lang-tr-blue.svg)](README_TR.md)

Go ile yazılmış esnek ve güvenli bir komut satırı parola üreteci.

## Özellikler

- Özelleştirilebilir uzunlukta parola üretimi
- Özel karakterleri dahil etme/çıkarma seçeneği
- Sayıları dahil etme/çıkarma seçeneği
- Büyük harfleri dahil etme/çıkarma seçeneği
- Küçük harfleri dahil etme/çıkarma seçeneği
- Aynı anda birden fazla parola üretebilme
- Parola gücü analizi (Mükemmel, Güçlü, Orta, Zayıf)
- Entropi hesaplama ve görüntüleme
- Üretim süresi görüntüleme
- Sürüm bilgisi
- Kullanımı kolay komut satırı arayüzü

## Kurulum

```bash
go install github.com/efeaslansoyler/go-passwordgen@latest
```

## Kullanım

Temel kullanım:
```bash
go-passwordgen
```

Bu komut, tüm karakter tiplerini içeren 12 karakterlik bir parola üretecektir.

### Seçenekler

- `-l, --length`: Parola uzunluğunu belirler (varsayılan: 12)
- `-s, --special`: Özel karakterleri dahil eder (varsayılan: true)
- `-n, --numbers`: Sayıları dahil eder (varsayılan: true)
- `-u, --upper`: Büyük harfleri dahil eder (varsayılan: true)
- `-o, --lower`: Küçük harfleri dahil eder (varsayılan: true)
- `-c, --count`: Üretilecek parola sayısı (varsayılan: 1)
- `-q, --quiet`: Çıktıyı bastırır (sadece parolayı yazdırır) (varsayılan: false)
- `-v, --version`: Sürüm bilgisini görüntüler

### Örnekler

16 karakterlik bir parola üretmek için:
```bash
go-passwordgen -l 16
```

5 adet parola üretmek için:
```bash
go-passwordgen -c 5
```

Özel karakterler olmadan parola üretmek için:
```bash
go-passwordgen --special=false
```

## Lisans

Bu proje MIT Lisansı ile lisanslanmıştır - detaylar için [LICENSE](LICENSE) dosyasına bakınız.

## Yazar

Efe Aslan Söyler (efeaslan1703@gmail.com)
