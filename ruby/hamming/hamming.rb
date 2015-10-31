require 'pry'
class Hamming
  Hamming::VERSION = 1

  def self.compute(string1, string2)
    fail ArgumentError if string1.length != string2.length

    distance = 0
    string1.length.times do |pos|
      distance += 1 if string1[pos] != string2[pos]
    end
    distance
  end
end
