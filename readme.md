## Saints Quotes API

A RESTful API that delivers inspirational quotes from Catholic and Orthodox saints.  
It includes biographical details, spiritual attributes, and patronages associated with each saint.


### Motivation

To inspire reflection, spiritual growth, and inner peace, this project preserves the timeless wisdom of the saints and shares their teachings with the world.

### Getting Started

To get started with the API, you can make requests to the following endpoints:

#### Get A Random Quote

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

#### Get Quotes by Saint Name

```bash
/api/v1/quotes/search?name={author-name}
```


#### Get Quotes by Category

```bash
/api/v1/quotes/search/category?category={category}
```


### Rate Limiting

To ensure fair usage and prevent abuse, this API applies the following rate limiting policy only 20 requests per 15-minute window.

```json
{
  "message": "Too many requests, please try again after 15 minutes",
  "code": 429
}
```

### License

Licensed under the MIT License.  
Free to use, modify, and distribute even for commercial purposes.  
