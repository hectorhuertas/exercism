class Complement
  VERSION = 2
  def self.of_dna(word)
    convert_word(word, :dna)
  end

  def self.of_rna(word)
    convert_word(word, :rna)
  end

  def self.convert_word(word, type)
    word.chars.map do |char|
      convert_char(char, type)
    end.join
  end

  def self.transformations
    { dna: { 'C' => 'G', 'G' => 'C', 'T' => 'A', 'A' => 'U' },
      rna: { 'G' => 'C', 'C' => 'G', 'A' => 'T', 'U' => 'A' } }
  end

  def self.convert_char(char, type)
    transformations[type][char] || (fail ArgumentError)
  end
end
