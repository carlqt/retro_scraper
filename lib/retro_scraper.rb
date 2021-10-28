# frozen_string_literal: true

require_relative "retro_scraper/version"
require 'pry'

module RetroScraper
  class Error < StandardError; end
  # Your code goes here...
  def self.run
    browser = Ferrum::Browser.new(timout: 30, headless: false)
    browser.cookies.clear

    browser.headers.set(
      'referrer': '*'
    )

    browser.go_to("https://easyretro.io/publicboard/d0LnKN92OGdwxnFNhJt7Pow4T2b2/817d0a10-07ef-4408-bfb0-cd2e28c961c0")

    browser.quit
  end
end
