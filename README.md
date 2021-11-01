# retro_scraper
Scraper for easyretro

# How To Install

Download the binary in https://github.com/carlqt/retro_scraper/blob/master/bin/retro_scraper and place it on any directory you like.

# How to Use

Execute the binary and pass the url of easyretro to scrape e.g.

```
 ~/Projects/retro_scraper/bin  master *1  ls                                                                                                                   
retro_scraper
 ~/Projects/retro_scraper/bin  master *1  ./retro_scraper https://easyretro.io/publicboard/d0LnKN92OGdwxnFNhJt7Pow4T2b2/817d0a10-07ef-4408-bfb0-cd2e28c961c0 
2021/11/01 10:30:04 Done
 ~/Projects/retro_scraper/bin  master *1 ?1  ls                                                                                                             
output        retro_scraper
 ~/Projects/retro_scraper/bin  master *1 ?1  ls output                                                                                                 
easyretro.csv
 ~/Projects/retro_scraper/bin  master *1 ?1   
 ```
 
It will store the CSV under the `output` directory
