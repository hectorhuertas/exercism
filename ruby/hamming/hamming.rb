require 'pry'
class Hamming
  Hamming::VERSION = 1

  #My solution, pairing with Lovisa
  def self.compute(string1, string2)
    fail ArgumentError if string1.length != string2.length

    distance = 0
    string1.length.times do |pos|
      distance += 1 if string1[pos] != string2[pos]
    end
    distance
  end

  # #Another solution, driven by Lovisa
  # def self.compute(string1, string2)
  #   fail ArgumentError if string1.length != string2.length
  #
  #   string1.chars.zip(string2.chars).reduce(0) do |sum,pair|
  #     sum += 1 if pair[0] != pair[1]
  #     sum
  #   end
  # end
end
