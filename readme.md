## Saints Quotes API

A RESTful API that delivers inspirational quotes from Catholic and Orthodox saints and includes details about religious orders.
It features biographical information, spiritual attributes, patronages associated with each saint, as well as information on religious orders, their founding, and their global presence.

### Motivation

To inspire reflection, spiritual growth, and inner peace, this project preserves the timeless wisdom of the saints and the legacy of their religious orders. By sharing their teachings, we aim to bring light to the world through the examples of holiness and service.

### Getting Started

To get started with the API, you can make requests to the following endpoints:

#### Get A Random Quote

Retrieve a random inspirational quote from a saint.

```bash
api/v1/random-quotes
```

```json
{
    "success": true,
    "data": {
        "id": "67fe296cbd4db683f593dc8c",
        "quote": "Tuhan, buatlah aku alat perdamaian-Mu. Di mana ada kebencian, biarlah aku menaburkan cinta; di mana ada luka, biarlah aku membawa penyembuhan; di mana ada keraguan, biarlah aku membawa iman.",
        "author": {
            "name": "Santo Fransiskus Assisi",
            "biography": "Santo Fransiskus Assisi adalah seorang santo, diakon, biarawan, dan pendiri Ordo Saudara-Saudara Fransiskan (OFM). Ia dikenal karena hidup sederhana, cinta kepada alam, dan dedikasinya dalam membangun perdamaian. Lahir di Assisi, Italia, pada 10 Juli 1182, ia meninggal pada 3 Oktober 1226.",
            "birth_date": "1182-07-10",
            "death_date": "1226-10-03",
            "death_place": "Assisi, Italia",
            "congregation": "Ordo Saudara-Saudara Fransiskan",
            "denomination": [
                "Katolik Roma",
                "Anglikan"
            ],
            "veneration_date": "4 Oktober",
            "pilgrimage_site": "Basilika Santo Fransiskus dari Assisi",
            "festival_date": "4 Oktober",
            "attributes": [
                "Salib",
                "Burung Merpati",
                "Pax et Bonum",
                "Poor Franciscan habit",
                "Stigmata"
            ],
            "patronage": [
                "Binatang",
                "Lingkungan",
                "Saudagar"
            ]
        },
        "category": "Perdamaian, Kasih",
        "created_at": "2025-04-15T09:39:56.08Z",
        "updated_at": "2025-04-15T09:39:56.08Z"
    },
    "message": "Random quote successfully retrieved",
    "code": 200
}
```

#### Get Quotes by Saint or Category

Query :

- `name` (e.g., "Fransiskus Assisi")
- `category` (e.g., "Perdamaian, Kasih")

```bash
/api/v1/quotes/search?name={query}
```

```bash
/api/v1/quotes/search?category={query}

```

```json
{
   {
    "success": true,
    "data": [
        {
            "id": "67ff75ad54be8c0b466e81bc",
            "quote": "Tuhan, buatlah aku alat perdamaian-Mu. Di mana ada kebencian, biarlah aku menaburkan cinta; di mana ada luka, biarlah aku membawa penyembuhan; di mana ada keraguan, biarlah aku membawa iman.",
            "author": {
                "name": "Santo Fransiskus Assisi",
                "biography": "Santo Fransiskus Assisi adalah seorang santo, diakon, biarawan, dan pendiri Ordo Saudara-Saudara Fransiskan (OFM). Ia dikenal karena hidup sederhana, cinta kepada alam, dan dedikasinya dalam membangun perdamaian. Lahir di Assisi, Italia, pada 10 Juli 1182, ia meninggal pada 3 Oktober 1226.",
                "birth_date": "1182-07-10",
                "death_date": "1226-10-03",
                "death_place": "Assisi, Italia",
                "congregation": "Ordo Saudara-Saudara Fransiskan",
                "denomination": [
                    "Katolik Roma",
                    "Anglikan"
                ],
                "veneration_date": "4 Oktober",
                "pilgrimage_site": "Basilika Santo Fransiskus dari Assisi",
                "attributes": [
                    "Salib",
                    "Burung Merpati",
                    "Pax et Bonum",
                    "Poor Franciscan habit",
                    "Stigmata"
                ],
                "patronage": [
                    "Binatang",
                    "Lingkungan",
                    "Saudagar"
                ]
            },
            "category": "Perdamaian, Kasih",
            "created_at": "2025-04-15T09:39:56.08Z",
            "updated_at": "2025-04-16T09:17:33.826Z"
        }
    ],
    "message": "Quotes successfully retrieved for the saint",
    "code": 200
    }
}
```

#### Get A Ordos

Retrieve a ordo inspirational from chatolic church

Query

- `name` (e.g., "Ordo Kapusin")
- `nickname` (e.g., "Kapusin")
- `category` (e.g., "Imam")
- `title` (e.g., "Cap")

```bash
/api/v1/search?order={query}
```

```json
{
    "success": true,
    "data": {
        "order_name": "Ordo Kapusin",
        "latin_name": "Ordo Fratrum Minorum Capuccinorum",
        "nickname": "Kapusin",
        "title": "O.F.M. Cap. (Ordo Saudara Dina Kapusin)",
        "category": "Klerikal (Imam)",
        "denomination": "Gereja Katolik Roma",
        "founder_name": "Matteo da Bascio (terinspirasi oleh Santo Fransiskus dari Assisi)",
        "founded_year": 1528,
        "current_leader": "Saudara Roberto Genuin, O.F.M. Cap.",
        "address": {
            "head_quarters": "General Headquarters of the Order of Friars Minor Capuchin",
            "street": "Via Piemonte, 70",
            "city": "Rome",
            "postal_code": "00187",
            "country": "Italy",
            "phone": [
                "+39 06 42011710",
                "+39 335 1641820"
            ],
            "fax": "+39 06 4828267"
        },
        "description": "Ordo Kapusin adalah sebuah tarekat religius dalam Gereja Katolik yang anggotanya—imam, bruder, dan frater—hidup bersama sebagai saudara. Nama 'Kapusin' berasal dari ciri khas tudung panjang runcing (capuccio) yang dikenakan oleh anggotanya. Sebagai cabang reformasi dari Ordo Fransiskan, mereka didirikan pada tahun 1528 untuk hidup lebih sederhana dan meneladani kehidupan Santo Fransiskus dari Assisi. Para frater Kapusin hidup dalam kerendahan hati dan kesederhanaan, melayani kaum miskin dan terpinggirkan melalui doa, kontemplasi, dan pelayanan sosial. Kehidupan komunitas mereka ditandai dengan kasih, pengorbanan, dan kesaksian Injil melalui cara hidup mereka.",
        "slogan": "Fratres Minores – Saudara-saudara Dina",
        "related_orders": [
            "Ordo Saudara Dina (OFM)",
            "Ordo Saudara Dina Konventual (OFMConv)",
            "Ordo Santa Klara",
            "Ordo Ketiga Reguler (TOR)",
            "Ordo Fransiskan Sekular (OFS)"
        ],
        "coat_of_arms_url": "https://upload.wikimedia.org/wikipedia/commons/a/ab/Coat_of_Arms_of_the_Order_of_Friars_Minor.svg",
        "constitution_approval": {
            "approved_by_pope": "Paus Klemens VII",
            "approval_date": "1528"
        },
        "missionary_activity": {
            "initial_goals": "Hidup secara radikal sesuai Injil, melayani yang miskin dan terpinggirkan dengan kerendahan hati.",
            "global_regions": [
                "Eropa Barat",
                "Eropa Tengah dan Timur",
                "Asia dan Oseania",
                "Amerika Utara",
                "Amerika Selatan"
            ],
            "indonesia_regions": [
                "Provinsi Kapusin Medan",
                "Kustodi Jenderal Sibolga",
                "Kustodi Jenderal Kepulauan Nias",
                "Provinsi Kapusin Pontianak"
            ]
        },
        "principles": {
            "motto": "Pax et Bonum – Damai dan Kebaikan",
            "spiritual_practice": [
                "Doa kontemplatif",
                "Hidup sederhana dan rendah hati",
                "Pelayanan pastoral dan sosial",
                "Pewartaan Injil lewat teladan hidup"
            ]
        },
        "notable_saints_blesseds": [
            {
                "name": "Santo Laurensius dari Brindisi",
                "contribution": "Doktor Gereja, dikenal karena pewartaan dan kontribusi dalam Kontra-Reformasi."
            },
            {
                "name": "Santo Pio dari Pietrelcina (Padre Pio)",
                "contribution": "Dikenal karena stigmata, karisma penyembuhan, dan bimbingan rohani."
            }
        ],
        "rules": [
            "Dalam nama Tuhan dimulai hidup Saudara Dina",
            "Tentang mereka yang ingin menjalani hidup ini dan bagaimana mereka diterima",
            "Tentang Ibadat Ilahi, puasa, dan cara saudara berjalan di dunia",
            "Agar saudara tidak menerima uang",
            "Tuhan memberikan aku, Fransiskus, untuk mulai bertobat dengan cara ini",
            "Tentang cara bekerja",
            "Agar saudara tidak memiliki apa pun untuk diri sendiri: tentang mengemis dan merawat orang sakit",
            "Tentang koreksi dan hukuman bagi saudara yang berdosa",
            "Tentang pemilihan Menteri Jenderal dan Kapitel Pentekosta",
            "Tentang para pewarta",
            "Tentang nasihat dan koreksi persaudaraan",
            "Agar saudara tidak masuk ke biara-biara para suster",
            "Tentang pergi ke antara orang Saracen dan non-Kristen lainnya"
        ],
        "metadata": {
            "source": [
                "https://www.ofmcap.org/it/",
                "Provinsi Kapusin Medan",
                "Kustodi Jenderal Sibolga",
                "Kustodi Jenderal Kepulauan Nias",
                "Provinsi Kapusin Pontianak"
            ],
            "retrieved_date": "2025-04-17",
            "source_description": "Data ini diambil dari situs resmi global dan Indonesia dari Ordo Kapusin serta sumber lain yang terkait sejarah dan konstitusinya."
        },
        "created_at": "2025-04-17T12:00:00Z",
        "updated_at": "2025-04-17T06:42:13.955Z"
    },
    "message": "Order successfully retrieved",
    "code": 200
}
```

### Content Sources

The quotes from saints and holy figures in this application are gathered from various sources. If there are any errors or discrepancies, please contact us for corrections.

### License

The content provided in this application, including quotes from saints and holy figures, is available for free use under the MIT License. You are free to use, modify, and share the content as long as appropriate credit is given to the original sources.
