### Database document
> 2022-01-01 14:50:37
#### sales_prices  
NO | KEY | COLUMN | COMMENT | DATA_TYPE | NOTNULL | REMARK
:---: | :---: | --- | --- | --- | :---: | ---
1|PRI|id| |BİGİNT|Y|
2| |customer_id| Müşteri id|BİGİNT|N|
3| |customer_no|Müşteri No|TEXT|N|
4| |customer_name|Müşterinin adı|TEXT|N|
5| |product_id| Ürün numarası|BİGİNT|N|
6| |product_code|Ürün Kodu|TEXT|N|
7| |product_name|Ürün adı|TEXT|N|
8| |unit_price| Ürün Fiyati|NUMERİC|N|
9| |currency_code|Döviz Kodu: USD,TL,EUR gibi|TEXT|N|
10| |minimum_quantity|Minimum alış miktarı|NUMERİC|N|
11| |unit_of_measure|Birim|TEXT|N|
12| |start_date|Fiyatın geçerlilik başlangıç tarihi: Bu tarihten itibaren bu fiyat geçerli olacak.|TİMESTAMP WİTH TİME ZONE|N|
13| |end_date|Fiyatın geçerlilik bitiş tarihi. Bu tarihe kadar bu fiyat geçerli olacak.|TİMESTAMP WİTH TİME ZONE|N|
