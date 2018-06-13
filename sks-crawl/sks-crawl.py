# TODO: this doesn't work yet, but is supposed to crawl WoT and collect
# everyone's keys in dearmored format.

import scrapy

class SKSSpider(scrapy.Spider):
    name = "sks-spider";

    start_urls = ['https://pgp.key-server.io/search/0x41259773973A612A']

    def parse(self, response):
        for href in response.css('a::attr(href)').extract():
            print(href)
