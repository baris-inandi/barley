class Barley < Formula
  desc "Cross-platform Text Expander written in Rust"
  homepage "https://github.com/baris-inandi/barley"
  url "https://github.com/baris-inandi/barley"
  sha256 "2aa782df52939a5c2eb59248597cf24006d3b68e7e65960073a63ad33c0c0f29"
  version "1"

  def install
    bin.install "barley"
  end
end
