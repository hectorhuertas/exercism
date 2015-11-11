module Gigasecond
  VERSION = 1
  GIGASECOND = 1e9

  def self.from(time)
    time + GIGASECOND
  end
end
